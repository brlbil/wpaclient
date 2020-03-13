package wpaclient

import (
	"fmt"
	"net"
)

var socketType = "UDP"

func localSocket(i int) string {
	return ""
}

func dial(addr string) (*socket, error) {
	if addr == "" {
		addr = "127.0.0.1:9878"
	}

	remote, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("could not resolve udp address: %w", err)
	}

	c, err := net.DialUDP("udp",
		&net.UDPAddr{},
		remote,
	)
	if err != nil {
		return nil, fmt.Errorf("dial failed: %w", err)
	}

	return &socket{c: c}, nil
}

func (s *Socket) close() error {
	return s.c.close()
}

func testServerConn() (*net.UDPConn, func(), error) {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		return nil, nil, err
	}
	con, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, func() {}, err
	}

	fn := func() {
		con.Close()
	}

	return con, fn, nil
}
