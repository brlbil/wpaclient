package wpaclient

import (
	"fmt"
	"strconv"
	"strings"
)

// AuthReq represents data received with "CTRL-REQ-" requests
type AuthReq struct {
	ID   int
	Type string
	Text string
}

// Event represends events received from wpa_supplicant
type Event struct {
	Sev     int
	Message string
	AuthReq *AuthReq
	Err     error
}

func parseEvent(b []byte) *Event {
	if len(b) < 5 {
		msg := strings.TrimSuffix(string(b), "\n")
		return &Event{Err: fmt.Errorf("message too short: %s", msg)}
	}

	sb, err := strconv.Atoi(string(b[1]))
	if err != nil {
		return &Event{Err: fmt.Errorf("parse severity: %w", err)}
	}

	msg := strings.TrimSuffix(string(b[3:]), "\n")
	if strings.HasPrefix(msg, WpaCtrlReq) {
		msg = strings.TrimPrefix(msg, WpaCtrlReq)

		i := strings.Index(msg, "-")
		j := strings.Index(msg, ":")

		id, err := strconv.Atoi(msg[i+1 : j])
		if err != nil {
			return &Event{Err: fmt.Errorf("parse networkID: %w", err)}
		}

		return &Event{Sev: sb, Message: WpaCtrlReq,
			AuthReq: &AuthReq{ID: id, Type: msg[:i], Text: msg[j+1:]}}
	}

	return &Event{Sev: sb, Message: msg}
}
