package wpaclient

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	cmdAttach = "ATTACH"
	cmdDetach = "DETACH"
)

type handlers struct {
	sync.RWMutex
	cm  map[string]chan Event
	evm map[string]map[string]struct{}
}

// Client represends wpa_supplicant client
type Client struct {
	// socket connection address
	addr string

	// primary socket to run commands
	cmdsock *socket

	// socket to get events
	evsock *socket

	// channel to push events
	evch chan Event

	// holds subscribed channels to receive events
	hand *handlers

	// hold the status for event socket
	attached bool

	// mutex for protecting attached status
	amut *sync.RWMutex

	// mutex for protecting command execution
	cmdmut *sync.Mutex
}

// New returns a new Client object, returns error if dialing socket fails
func New(addr string) (*Client, error) {
	cs, err := dial(addr)
	if err != nil {
		return nil, err
	}

	return &Client{addr: addr, cmdsock: cs, evch: make(chan Event, 10),

		amut: &sync.RWMutex{}, cmdmut: &sync.Mutex{}, hand: &handlers{}}, nil
}

// Execute send a commad with its args to wpa_supplicant and reads the response
// returns ErrCmdFailed if FAIL returned, returns ErrUnknownCmd if "UNKNOWN COMMAND" returnred
// return InvalidCmdErr if Invalid <CMD> or usage message returned
func (c *Client) Execute(cmd string, args ...string) ([]byte, error) {
	c.cmdmut.Lock()
	defer c.cmdmut.Unlock()

	b := []byte(cmd)
	if len(args) > 0 {
		a := " " + strings.Join(args, " ")
		b = append(b, []byte(a)...)
	}

	buf, err := c.cmdsock.execute(b)

	if err != nil {
		return nil, err
	}

	if err := validate(cmd, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *Client) dispacher() {
	for ev := range c.evch {
		c.hand.RLock()
		for ca, ch := range c.hand.cm {
			evm := c.hand.evm[ca]
			_, ok := evm[ev.Message]
			if ev.Err != nil || len(evm) == 0 || ok {
				select {
				case ch <- ev:
				default:
				}
			}
		}
		c.hand.RUnlock()
	}

	// eventChannel closed so we need to close all channels listen for events
	c.hand.Lock()

	for _, ch := range c.hand.cm {
		close(ch)
	}
	c.hand.cm = nil
	c.hand.evm = nil

	c.hand.Unlock()
}

// Notify returns a receive only event channel.
// If no events are provided, all incoming events will be relayed to channel.
// Otherwise, just the provided events will.
func (c *Client) Notify(evs ...string) (<-chan Event, error) {
	c.hand.Lock()
	defer c.hand.Unlock()

	if c.hand.cm == nil {
		c.hand.cm = make(map[string]chan Event)
		c.hand.evm = make(map[string]map[string]struct{})
	}

	ch := make(chan Event, 5)
	ca := fmt.Sprint(ch)

	c.hand.cm[ca] = ch
	evm := map[string]struct{}{}
	for _, ev := range evs {
		evm[ev] = struct{}{}
	}
	c.hand.evm[ca] = evm

	c.amut.RLock()
	a := c.attached
	c.amut.RUnlock()

	if !a {
		if err := c.attach(); err != nil {
			return nil, errors.Wrap(err, "attach failed")
		}

		go c.dispacher()
	}

	return ch, nil
}

// Stop causes client to stop relaying incoming events to ch.
func (c *Client) Stop(ch <-chan Event) {
	c.hand.Lock()
	defer c.hand.Unlock()

	ca := fmt.Sprint(ch)

	chn := c.hand.cm[ca]
	if chn == nil {
		return
	}

	close(chn)
	delete(c.hand.cm, ca)
	delete(c.hand.evm, ca)
}

// attach attaches on second socket and receives events
// every subsequent call return a subscriber channel
func (c *Client) attach() error {
	c.amut.Lock()
	defer c.amut.Unlock()

	if c.attached {
		return nil
	}

	if c.evsock == nil {
		s, err := dial(c.addr)
		if err != nil {
			return err
		}

		c.evsock = s
	}

	res, err := c.evsock.execute([]byte(cmdAttach))
	if err != nil {
		return err
	}

	if err := validate(cmdAttach, res); err != nil {
		return err
	}

	c.evch = make(chan Event, 10)

	go func() {
		ucn := "use of closed network connection"
		for {
			b, err := c.evsock.receive()
			if err != nil {
				if err == io.EOF || strings.Contains(err.Error(), ucn) {
					return
				}
			}

			// Detach command received, terminate here
			if bytes.Equal(b, []byte("OK\n")) {
				return
			}

			select {
			case c.evch <- *parseEvent(b):
			default:
			}
		}
	}()

	c.attached = true

	return nil
}

// detach detaches from event socket
func (c *Client) detach() error {
	c.amut.Lock()
	defer c.amut.Unlock()

	if c.attached {
		c.attached = false
		close(c.evch)
		return c.evsock.send([]byte(cmdDetach))
	}

	return nil
}

// Close closes cmd and event sockets
func (c *Client) Close() error {
	var err error

	if c.cmdsock != nil {

		if e := c.cmdsock.close(); e != nil {

			if err != nil {
				err = errors.Wrap(e, err.Error())
			}
			err = e
		}
	}

	if c.evsock != nil {
		e := c.detach()
		if err != nil {
			err = errors.Wrap(e, err.Error())
		} else {
			err = e
		}

		if e := c.evsock.close(); e != nil {
			if err != nil {
				err = errors.Wrap(e, err.Error())
			}
			err = e
		}
	}

	return err
}

// Scan executes "SCAN" and "SCAN_RESULT" commands returns scanned list of Access Points
func (c *Client) Scan() ([]AP, error) {
	ch, err := c.Notify(WpsEventApAvailable)
	if err != nil {
		return nil, err
	}
	defer c.Stop(ch)

	scan := func() error {
		_, err := c.Execute(CmdScan)
		if err != nil {
			return err
		}

		select {
		case <-ch:
		case <-time.After(time.Second * 2):
			return errors.New("scan timed out")
		}

		return nil
	}

	scanRes := func() ([]AP, error) {
		res, err := c.Execute(CmdScanResults)
		if err != nil {
			return nil, err
		}

		aps, err := parseAP(res)
		if err != nil {
			return nil, err
		}

		return aps, nil
	}

	aps, err := scanRes()
	if err != nil {
		return nil, err
	}

	if len(aps) == 0 {
		if err := scan(); err != nil {
			return nil, err
		}
	}

	return scanRes()
}

// ListNetworks executes "LIST_NETWORK" command and returns Networks
func (c *Client) ListNetworks() ([]Network, error) {
	res, err := c.Execute(CmdListNetworks)
	if err != nil {
		return nil, err
	}

	return parseNetwork(res)
}
