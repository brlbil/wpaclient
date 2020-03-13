package wpaclient

import (
	"fmt"
	"io"
	"net"
)

type socket struct {
	c net.Conn
	f string
}

func (s *socket) send(b []byte) error {
	n, err := s.c.Write(b)
	if err != nil {
		return fmt.Errorf("write to socket failed: %w", err)
	}

	if n != len(b) {
		return io.ErrShortWrite
	}

	return nil
}

func (s *socket) receive() ([]byte, error) {
	b := make([]byte, 4095)

	n, err := s.c.Read(b[:])
	if err != nil {
		return nil, fmt.Errorf("read from socket failed: %w", err)
	}

	return b[:n], nil
}

func (s *socket) execute(cmd []byte) ([]byte, error) {
	err := s.send(cmd)
	if err != nil {
		return nil, fmt.Errorf("send failed: %w", err)
	}

	buf, err := s.receive()
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}

	return buf, nil
}
