package wpaclient

import (
	"net"
	"reflect"
	"testing"
)

func TestUnmarshalNetwork(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []Network
		err    bool
	}{
		{
			name: "list no ssid",
			input: `network id / ssid / bssid / flags
0		any	[DISABLED]
1		any	[DISABLED]
2		any	[DISABLED]`,
			expect: []Network{
				{ID: 0, SSID: "", BSSID: "any", Flags: []string{"DISABLED"}},
				{ID: 1, SSID: "", BSSID: "any", Flags: []string{"DISABLED"}},
				{ID: 2, SSID: "", BSSID: "any", Flags: []string{"DISABLED"}},
			},
		},
		{
			name: "list with ssid",
			input: `network id / ssid / bssid / flags
0	AP0	any	[DISABLED]
1		any	[DISABLED]
2		any	[DISABLED]`,
			expect: []Network{
				{ID: 0, SSID: "AP0", BSSID: "any", Flags: []string{"DISABLED"}},
				{ID: 1, SSID: "", BSSID: "any", Flags: []string{"DISABLED"}},
				{ID: 2, SSID: "", BSSID: "any", Flags: []string{"DISABLED"}},
			},
		},
		{
			name:   "no results",
			input:  "network id / ssid / bssid / flags\n",
			expect: []Network{},
		},
		{
			name:   "empty",
			expect: []Network{},
		},
		{
			name: "id error",
			input: "a		any	[DISABLED]",
			err: true,
		},
		{
			name: "read error",
			input: "0		any	[DISABLED]	A",
			err: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nts, err := parseNetwork([]byte(tt.input))
			if err != nil && !tt.err {
				t.Errorf("Expected error nil, got %v", err)
			}

			if !reflect.DeepEqual(nts, tt.expect) {
				t.Errorf("Expected %#v\ngot %#v", tt.expect, nts)
			}
		})
	}
}

func TestUnmarshalAP(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []AP
		err    bool
	}{
		{
			name: "scan",
			input: `bssid / frequency / signal level / flags / ssid
d0:7a:b5:31:23:a0	2472	-30	[WPA2-PSK-CCMP][WPS][ESS]	AP0
00:1f:1f:37:42:d9	2442	-37	[WPA2-PSK-CCMP][ESS]	AP1
24:00:ba:f8:65:df	2412	-77	[WPA-PSK-CCMP+TKIP][WPA2-PSK-CCMP+TKIP][WPS][ESS]	AP2
`,
			expect: []AP{
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
			},
		},
		{
			name: "no flags",
			input: "d0:7a:b5:31:23:a0	2472	-30		AP0",
			expect: []AP{
				{
					BSSID:          net.HardwareAddr{0xd0, 0x7a, 0xb5, 0x31, 0x23, 0xa0},
					SSID:           "AP0",
					Frequency:      2472,
					SignalStrength: -30,
					Flags:          []string{""},
				},
			},
		},
		{
			name:   "no results",
			input:  "bssid / frequency / signal level / flags / ssid\n",
			expect: []AP{},
		},
		{
			name:   "empty",
			expect: []AP{},
		},
		{
			name: "mac error",
			input: "d0:7a:b5	2472	-30	[WPA2-PSK-CCMP][WPS][ESS]	AP0",
			err: true,
		},
		{
			name: "frequency error",
			input: "d0:7a:b5:31:23:a0	abc	-30	[WPA2-PSK-CCMP][WPS][ESS]	AP0",
			err: true,
		},
		{
			name: "signal error",
			input: "d0:7a:b5:31:23:a0	2472	-ab	[WPA2-PSK-CCMP][WPS][ESS]	AP0",
			err: true,
		},
		{
			name: "read error",
			input: "d0:7a:b5:31:23:a0	2472	-ab	[WPA2-PSK-CCMP][WPS][ESS]	AP0	ERROR",
			err: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aps, err := parseAP([]byte(tt.input))
			if err != nil && !tt.err {
				t.Errorf("Expected error nil, got %v", err)
			}

			if !reflect.DeepEqual(aps, tt.expect) {
				t.Errorf("Expected %#v\ngot %#v", tt.expect, aps)
			}
		})
	}

}
