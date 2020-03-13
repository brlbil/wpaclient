package wpaclient

import (
	"fmt"
	"net"
	"os"
	"path"
)

var socketType = "UNIX"

func localSocket(i int) string {
	return fmt.Sprintf("/tmp/wpa_ctrl_%d-%d", os.Getpid(), i)
}

func dial(addr string) (*socket, error) {
	var (
		ad  string
		err error
		lf  string
	)

	for _, a := range []string{
		addr,
		path.Join("/var/wpa_supplicant", addr),
		path.Join("/var/run/wpa_supplicant", addr),
	} {
		if _, err = os.Stat(a); !os.IsNotExist(err) {
			ad = a
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("socket not found: %w", err)
	}

	// ResolveUnixAddr will not return an error,
	// unless network is not unix, unixgram or unixpacket
	ra, _ := net.ResolveUnixAddr("unixgram", ad)

	for i := 0; i <= 2; i++ {
		l := localSocket(i)
		if i == 2 {
			return nil, fmt.Errorf("reached max socket file limit: %s", l)
		}

		_, err := os.Stat(l)
		if os.IsNotExist(err) {
			lf = l
			break
		}
	}

	c, err := net.DialUnix(
		"unixgram",
		&net.UnixAddr{Name: lf, Net: "unixgram"},
		ra,
	)

	if err != nil {
		os.Remove(lf)
		return nil, fmt.Errorf("dial failed: %w", err)
	}

	return &socket{c: c, f: lf}, nil
}

func (s *socket) close() error {
	defer os.Remove(s.f)
	return s.c.Close()
}

func testServerConn() (*net.UnixConn, func(), error) {
	lf := fmt.Sprintf("/tmp/wpa_test_listen_%d", os.Getpid())
	addr, err := net.ResolveUnixAddr("unixgram", lf)
	if err != nil {
		return nil, func() {}, err
	}

	con, err := net.ListenUnixgram("unixgram", addr)
	if err != nil {
		return nil, func() {}, err
	}

	fn := func() {
		con.Close()
		os.Remove(lf)
	}

	return con, fn, nil
}
