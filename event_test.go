package wpaclient

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestParseEvent(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
		ev   *Event
	}{
		{
			name: "short buf",
			buf:  []byte("MSG\n"),
			ev:   &Event{Err: errors.Errorf("message too short (MSG)")},
		},
		{
			name: "event severity",
			buf:  []byte("<a>Message"),
			ev:   &Event{Err: errors.Errorf("parse severity: strconv.Atoi: parsing \"a\": invalid syntax")},
		},
		{
			name: "event connected",
			buf:  []byte(fmt.Sprintf("<2>%s\n", WpaEventConnected)),
			ev:   &Event{Sev: 2, Message: WpaEventConnected},
		},
		{
			name: "event auth failed",
			buf:  []byte(fmt.Sprintf("<3>%sOTP-T:Challenge 1235663 needed for SSID foobar\n", WpaCtrlReq)),
			ev:   &Event{Err: errors.Errorf("parse networkID: strconv.Atoi: parsing \"T\": invalid syntax")},
		},
		{
			name: "event auth request",
			buf:  []byte(fmt.Sprintf("<3>%sPASSWORD-1:Password needed for SSID foobar\n", WpaCtrlReq)),
			ev: &Event{
				Sev:     3,
				Message: WpaCtrlReq,
				AuthReq: &AuthReq{
					ID:   1,
					Type: "PASSWORD",
					Text: "Password needed for SSID foobar",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := parseEvent(tt.buf)
			if tt.ev.Err != nil && tt.ev.Err.Error() != ev.Err.Error() {
				t.Errorf("Error expected %v, got %v", tt.ev.Err, ev.Err)
			}

			if tt.ev.Err == nil && !reflect.DeepEqual(tt.ev, ev) {
				t.Errorf("Event expected %#v, got %#v", tt.ev, ev)
			}
		})
	}
}
