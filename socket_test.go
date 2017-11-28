package wpaclient

import (
	"os"
	"strings"
	"testing"
)

func TestConnDial(t *testing.T) {
	tests := []struct {
		name   string
		addr   string
		expErr bool
		clean  func(s *socket)
		skip   bool
	}{
		{name: "success0"},
		{name: "success1"},
		{name: "fail max socket count", expErr: true, skip: socketType == "UDP",
			clean: func(s *socket) {
				for i := 0; i < 2; i++ {
					os.Remove(localSocket(i))
				}
			}},
		{name: "check file", skip: socketType == "UDP",
			clean: func(s *socket) {
				s.close()
				if _, err := os.Stat(localSocket(0)); !os.IsNotExist(err) {
					t.Errorf("Expect Close() to remove socket file, it exists")
				}
			}},
		{name: "not exist", addr: "not_exist", expErr: true},
		{name: "closed", expErr: true},
	}

	ts, close := newTestServer(t)
	defer close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}

			if tt.addr == "" {
				tt.addr = ts.addr()
			}

			if tt.name == "closed" {
				if socketType == "UDP" {
					tt.addr = "127.0.0.1:1"
				} else {
					fn := "/tmp/dkfhnli23yr98ndshfw2asdfkajjfo3eiu"
					tt.addr = fn
					os.Create(fn)
					defer os.Remove(fn)
				}
			}

			soc, err := dial(tt.addr)
			if err != nil && !tt.expErr {
				t.Errorf("Dial not expect an error, got %v", err)
			}

			if tt.expErr && err == nil {
				t.Error("Dial expect an error, got <nil>")
			}

			if tt.clean != nil {
				tt.clean(soc)
			}
		})
	}
}

func TestSocketExecute(t *testing.T) {
	ts, close := newTestServer(t)
	defer close()

	soc, err := dial(ts.addr())
	if err != nil {
		t.Fatalf("Dial failed, %v", err)
	}
	defer soc.close()

	tests := []struct {
		name string
		fn   func() error
		err  string
	}{
		{name: "success", fn: func() error {
			_, err := soc.execute([]byte(CmdPing))
			return err
		}},
		{name: "receive fail", err: "use of closed network connection",
			fn: func() error {
				soc.close()
				_, err := soc.receive()
				return err
			}},
		{name: "send fail", err: "use of closed network connection",
			fn: func() error {
				_, err := soc.execute([]byte(CmdPing))
				return err
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			if (tt.err == "" && err != nil) || (err != nil && !strings.Contains(err.Error(), tt.err)) {
				t.Errorf("Execute expect error %v, got %v", tt.err, err)
			}
		})
	}
}
