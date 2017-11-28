package wpaclient

import (
	"bytes"
	"fmt"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestNewClient(t *testing.T) {
	if _, err := New("not_exist"); err == nil {
		t.Error("New expected an error, got <nil>")
	}

	c := Client{}
	if err := c.Close(); err != nil {
		t.Errorf("Close expect an error <nil>, got %v", err)
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		cmd  string
		args []string
		exp  string
		err  error
	}{
		{cmd: CmdPing, exp: "PONG\n"},
		{cmd: "NOPE", args: []string{"arg0"}, err: ErrUnknownCmd},
		{cmd: CmdPing, err: ErrCmdFailed},
	}

	ts, close := newTestServer(t)
	defer close()

	c, _ := New(ts.addr())
	defer c.Close()

	for i, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			if i == 1 {
				ts.cmdMap[CmdPing] = "BOOM"
			}

			b, err := c.Execute(tt.cmd, tt.args...)
			if tt.err != errors.Cause(err) {
				t.Errorf("Execute expect error %v, got %v", tt.err, err)
			}

			if !bytes.Equal(b, []byte(tt.exp)) {
				t.Errorf("Output expected %s, got %s", tt.exp, string(b))
			}
		})
	}
}

func TestAttachDetach(t *testing.T) {
	ts, close := newTestServer(t)
	defer close()

	c, err := New(ts.addr())
	if err != nil {
		t.Fatalf("New Client failed, %v", err)
	}
	defer c.Close()

	if err := c.detach(); err != nil {
		t.Errorf("Detach not expect an error, got %v", err)
	}

	ts.cmdMap[cmdAttach] = "FAIL"
	if err := c.attach(); err == nil {
		t.Errorf("Attach expect an error, got %v", err)
	}
	delete(ts.cmdMap, cmdAttach)

	for i := 0; i < 2; i++ {
		if err := c.attach(); err != nil {
			t.Errorf("Attach not expect an error, got %v", err)
		}
	}

	// detach first
	if err := c.detach(); err != nil {
		t.Errorf("Detach not expect an error, got %v", err)
	}

	// check if attached, evsock set
	if c.attached || c.evsock == nil {
		t.Fatalf("Detach did not work as expected")
	}

	// lets wait for go ruting to exit
	time.Sleep(time.Millisecond)

	c.evsock.close()

	if err := c.attach(); err == nil {
		t.Errorf("Attach expect an error, got %v", err)
	}

	c.evsock = nil
	close()

	if err := c.attach(); err == nil {
		t.Errorf("Attach expect an error, got %v", err)
	}

}

func TestNotifyStop(t *testing.T) {
	tests := []struct {
		name string
		evs  []string
		exp  []string
		stop bool
	}{
		{
			name: "three registered, four events got",
			evs:  []string{WpaEventBssAdded, WpaEventDisconnected, WpaEventNetworkNotFound},
			exp:  []string{WpaEventBssAdded, WpaEventBssAdded, WpaEventDisconnected, WpaEventNetworkNotFound},
		},
		{
			name: "all",
			exp: []string{
				WpaEventAvoidFreq,
				WpaEventBeaconLoss,
				WpaEventBssAdded,
				WpaEventBssAdded,
				WpaEventChannelSwitch,
			},
		},
		{
			name: "last one",
			evs:  []string{WpaEventPasswordChanged},
			exp:  []string{WpaEventPasswordChanged},
			stop: true,
		},
	}

	ts, close := newTestServer(t)
	defer close()

	c, err := New(ts.addr())
	if err != nil {
		t.Fatalf("New Client failed, %v", err)
	}
	defer c.Close()

	chs := []<-chan Event{}

	for _, tt := range tests {
		ch, err := c.Notify(tt.evs...)
		if err != nil {
			t.Fatalf("Notify not expect an error, got %v", err)
		}
		chs = append(chs, ch)

		if ch, ok := c.hand.cm[fmt.Sprint(ch)]; !ok || ch == nil {
			t.Fatalf("Handler channel element not set: %v, channel is nil: %v", ok, ch == nil)
		}

		if evm, ok := c.hand.evm[fmt.Sprint(ch)]; !ok || (len(evm) != len(tt.evs)) {
			t.Fatalf("Handler event filter not set: %v, len not equal: %d == %d", ok, len(evm), len(tt.evs))
		}
	}

	if _, err := c.Execute("EVENTS"); err != nil {
		t.Errorf("Execute not expect an error, got %v", err)
	}
	time.Sleep(time.Millisecond)

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evs := []string{}

			for len(chs[i]) > 0 {
				ev := <-chs[i]
				evs = append(evs, ev.Message)
			}

			if !reflect.DeepEqual(evs, tt.exp) {
				t.Errorf("Expect to get event %v, got %v", tt.exp, evs)
			}

			if tt.stop {
				c.Stop(chs[i])

				if _, ok := <-chs[i]; ok {
					t.Errorf("expect channel %d to be closed", i)
				}
			}
		})
	}

	if err := c.detach(); err != nil {
		t.Errorf("Detach not expect an error, got %v", err)
	}

	for i, ch := range chs {
		if _, ok := <-ch; ok {
			t.Errorf("expect channel %d to be closed", i)
		}

		c.Stop(ch)
	}
}

func TestScan(t *testing.T) {
	ts, close := newTestServer(t)
	defer close()

	c, err := New(ts.addr())
	if err != nil {
		t.Fatalf("New Client failed, %v", err)
	}
	defer c.Close()

	expAps := []AP{
		{
			BSSID:          net.HardwareAddr{0xd0, 0x7a, 0xb5, 0x31, 0x23, 0xa0},
			SSID:           "AP0",
			Frequency:      2472,
			SignalStrength: -30,
			Flags:          []string{"WPA2-PSK-CCMP", "WPS", "ESS"},
		},
		{
			BSSID:          net.HardwareAddr{0x0, 0x1f, 0x1f, 0x37, 0x42, 0xd9},
			SSID:           "AP1",
			Frequency:      2442,
			SignalStrength: -37,
			Flags:          []string{"WPA2-PSK-CCMP", "ESS"},
		},
		{
			BSSID:          net.HardwareAddr{0x24, 0x0, 0xba, 0xf8, 0x65, 0xdf},
			SSID:           "AP2",
			Frequency:      2412,
			SignalStrength: -77,
			Flags:          []string{"WPA-PSK-CCMP+TKIP", "WPA2-PSK-CCMP+TKIP", "WPS", "ESS"},
		},
	}

	for i := 0; i < 2; i++ {
		aps, err := c.Scan()
		if err != nil {
			t.Errorf("Scan not expect an error, got %v", err)
		}

		if !reflect.DeepEqual(aps, expAps) {
			t.Errorf("Expected %#v\ngot %#v", expAps, aps)
		}
	}

}

func TestNetworkCmd(t *testing.T) {
	tests := []struct {
		name   string
		cmd    string
		ncount int
		args   string
		exp    string
		err    error
	}{
		{name: "list empty", cmd: CmdListNetworks},
		{name: "add 0", cmd: CmdAddNetwork, exp: "0"},
		{name: "add 1", cmd: CmdAddNetwork, exp: "1"},
		{name: "add 2", cmd: CmdAddNetwork, exp: "2"},
		{name: "rm 2", cmd: CmdRemoveNetwork, args: "2", exp: "OK"},
		{name: "rm 3", cmd: CmdRemoveNetwork, args: "2", err: ErrCmdFailed},
		{name: "rm", cmd: CmdRemoveNetwork,
			err: &InvalidCmdError{Cmd: CmdRemoveNetwork, Err: "at least 1 argument is required."}},
		{name: "set", cmd: CmdSetNetwork, err: &InvalidCmdError{Cmd: CmdSetNetwork}},
		{name: "set 0", cmd: CmdSetNetwork, args: "0 ssid \"test\"", exp: "OK"},
		{name: "set fail", cmd: CmdSetNetwork, args: "0 aaa bbb",
			err: &InvalidCmdError{Cmd: CmdSetNetwork, Err: "needs three arguments (network id, variable name, and value)"}},
		{name: "set 2", cmd: CmdSetNetwork, args: "2 ssid \"test\"", err: ErrCmdFailed},
		{name: "list", cmd: CmdListNetworks, ncount: 2},
	}

	ts, close := newTestServer(t)
	defer close()

	c, err := New(ts.addr())
	if err != nil {
		t.Fatalf("New Client failed, %v", err)
	}
	defer c.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cmd == CmdListNetworks {
				ns, err := c.ListNetworks()
				if tt.err != nil && tt.err != err {
					t.Errorf("ListNetworks expects error %v, got %v", tt.err, err)
				}

				if len(ns) != tt.ncount {
					t.Errorf("ListNetworks expects %d networks, got %d", tt.ncount, len(ns))
				}

				return
			}

			res, err := c.Execute(tt.cmd, strings.Split(tt.args, " ")...)
			if tt.err != err && (tt.err != nil && err != nil && tt.err.Error() != err.Error()) {
				t.Errorf("Execute(%s, %s) expects error %v, got %v", tt.cmd, tt.args, tt.err, err)
			}

			if tt.exp != "" && string(res) != tt.exp+"\n" {
				t.Errorf("Execute(%s, %s) expects res %v, got %v", tt.cmd, tt.args, tt.exp, string(res))
			}
		})
	}

}
