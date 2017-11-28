package wpaclient

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"testing"
)

type testConn interface {
	LocalAddr() net.Addr
	ReadFrom(b []byte) (int, net.Addr, error)
	WriteTo(b []byte, addr net.Addr) (int, error)
	Close() error
}

type testServer struct {
	conn     testConn
	subAddr  map[string]net.Addr
	networks []Network
	scanned  bool
	cmdMap   map[string]string
	t        *testing.T
}

var setNetworkUsage = `set_network variables:
	ssid (network name, SSID)
	psk (WPA passphrase or pre-shared key)
	key_mgmt (key management protocol)
	identity (EAP identity)
	password (EAP password)
	...

Note: Values are entered in the same format as the configuration file is using,
i.e., strings values need to be inside double quotation marks.
For example: set_network 1 ssid "network name"

Please see wpa_supplicant.conf documentation for full list of
available variables.`

var setNetworkFail = `Invalid SET_NETWORK command: needs three arguments
(network id, variable name, and value)`

var rmNetworkFail = "Invalid REMOVE_NETWORK command - at least 1 argument is required."

var netheader = "network id / ssid / bssid / flags"

var scanRes = `d0:7a:b5:31:23:a0	2472	-30	[WPA2-PSK-CCMP][WPS][ESS]	AP0
00:1f:1f:37:42:d9	2442	-37	[WPA2-PSK-CCMP][ESS]	AP1
24:00:ba:f8:65:df	2412	-77	[WPA-PSK-CCMP+TKIP][WPA2-PSK-CCMP+TKIP][WPS][ESS]	AP2`

func unmarshalNetwork(ns []Network) string {
	nl := ""
	for _, n := range ns {
		nl = fmt.Sprintf("%s\n%d\t%s\t%s\t[%s]",
			nl, n.ID, n.SSID, n.BSSID, strings.Join(n.Flags, "]["))
	}
	return nl
}

func (ts *testServer) run() {
	scanheader := "bssid / frequency / signal level / flags / ssid"

	go func() {
		b := make([]byte, 4095)
		for {
			n, raddr, err := ts.conn.ReadFrom(b)
			if err != nil {
				if err != io.EOF {
					return
				}
				ts.t.Errorf("TestServer Read error: %s", err)
			}

			sc := strings.Split(string(b[:n]), " ")
			scmd := sc[0]
			args := []string{}
			if len(sc) > 1 && sc[1] != "" {
				args = sc[1:]
			}

			if res, ok := ts.cmdMap[scmd]; ok {
				ts.write(res, raddr)
				continue
			}

			switch scmd {
			case cmdAttach:
				if _, ok := ts.subAddr[raddr.String()]; !ok {
					ts.subAddr[raddr.String()] = raddr
					ts.write("OK", raddr)
				}
			case cmdDetach:
				if _, ok := ts.subAddr[raddr.String()]; ok {
					delete(ts.subAddr, raddr.String())
					ts.write("OK", raddr)
				}
			// not a standart wpa command
			case "CLOSE":
				if err := ts.conn.Close(); err != nil {
					ts.t.Errorf("Close error, %s", err)
				}
			// not a standart wpa command
			case "EVENTS":
				ts.write("OK", raddr)
				evs := []string{
					WpaEventAvoidFreq,
					WpaEventBeaconLoss,
					WpaEventBssAdded,
					WpaEventBssAdded,
					WpaEventChannelSwitch,
					WpaEventConnected,
					WpaEventDisconnected,
					WpaEventEapFailure,
					WpaEventEapNotification,
					WpaEventNetworkNotFound,
					WpaEventPasswordChanged,
				}
				ts.sendMsg(2, evs...)
			case CmdScanResults:
				if ts.scanned {
					ts.write(fmt.Sprintf("%s\n%s", scanheader, scanRes), raddr)
					continue
				}
				ts.write(scanheader, raddr)
			case CmdScan:
				msg := []string{
					WpaEventScanStarted,
					WpaEventScanResults,
					WpsEventApAvailable,
				}
				ts.write("OK", raddr)
				ts.sendMsg(3, msg...)
				ts.scanned = true
			case CmdAddNetwork:
				id := len(ts.networks)
				ts.networks = append(ts.networks, Network{
					ID:    id,
					BSSID: "any",
					Flags: []string{"DISABLED"},
				})
				ts.write(fmt.Sprint(id), raddr)
			case CmdRemoveNetwork:
				if len(args) == 0 {
					ts.write(rmNetworkFail, raddr)
					continue
				}
				id, err := strconv.Atoi(args[0])
				if err != nil || (id < 0 || id >= len(ts.networks)) {
					ts.write("FAIL", raddr)
					continue
				}
				ts.networks = ts.networks[:id]
				ts.write("OK", raddr)
			case CmdSetNetwork:
				if len(args) == 0 {
					ts.write(setNetworkUsage, raddr)
					continue
				}
				if len(args) != 3 {
					ts.write(setNetworkFail, raddr)
					continue
				}

				id, err := strconv.Atoi(args[0])
				if err != nil || args[1] != "ssid" {
					ts.write(setNetworkFail, raddr)
					continue
				}
				if id < 0 || id >= len(ts.networks) {
					ts.write("FAIL", raddr)
					continue
				}
				ts.networks[id].SSID = strings.Trim(args[2], "\"")
				ts.write("OK", raddr)
			case CmdListNetworks:
				ts.write(netheader+unmarshalNetwork(ts.networks), raddr)
			default:
				ts.write("UNKNOWN COMMAND", raddr)
			}
		}
	}()
}

func (ts *testServer) write(s string, addr net.Addr) {
	_, err := ts.conn.WriteTo([]byte(s+"\n"), addr)
	if err != nil {
		ts.t.Logf("write cmd: %s to addr: %s, error: %s", s, addr.String(), err)
		return
	}
}

func (ts *testServer) sendMsg(sev int, msg ...string) {
	go func() {
		for _, s := range msg {
			for _, addr := range ts.subAddr {
				if sev < 0 {
					ts.write(s, addr)
				}
				ts.write(fmt.Sprintf("<%d>%s", sev, s), addr)
			}
		}
	}()
}

func (ts *testServer) addr() string {
	return ts.conn.LocalAddr().String()
}

func newTestServer(t *testing.T) (*testServer, func()) {
	conn, fn, err := testServerConn()
	if err != nil {
		t.Fatalf("New Test Server failed with error: %s", err)
	}

	m := map[string]string{
		CmdPing: "PONG",
	}

	ts := &testServer{conn: conn, subAddr: make(map[string]net.Addr), networks: []Network{}, cmdMap: m, t: t}
	ts.run()

	return ts, fn
}
