package wpaclient

import (
	"io"
	"net"

	"github.com/pkg/errors"
)

type socket struct {
	c net.Conn
	f string
}

func (s *socket) send(b []byte) error {
	n, err := s.c.Write(b)
	if err != nil {
		return errors.Wrap(err, "write to socket failed")
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
		return nil, errors.Wrap(err, "read from socket failed")
	}

	return b[:n], nil
}

func (s *socket) execute(cmd []byte) ([]byte, error) {
	err := s.send(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "send failed")
	}

	buf, err := s.receive()
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}

	return buf, nil
}
