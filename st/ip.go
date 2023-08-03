package st

import (
	"context"
	"errors"
	"net"
	"time"
)

type SendCloser interface {
	Send(ctx context.Context, command string) (string, error)
	Close() error
}

type ST_IP struct {
	socket net.Conn
}

func newIpComm(ctx context.Context, ipAddress string, timeout time.Duration) (*ST_IP, error) {
	d := net.Dialer{
		Timeout: timeout,
	}
	socket, err := d.DialContext(ctx, "udp", ipAddress+":7775")
	return &ST_IP{socket: socket}, err
}

func (s *ST_IP) Send(ctx context.Context, command string) (string, error) {
	b := make([]byte, 2+len(command))
	b[0] = 0
	b[1] = 7
	for i, v := range command {
		b[i+2] = byte(v)
	}
	nWritten, err := s.socket.Write(b)
	if err != nil {
		return "", err
	}
	if nWritten != 2+len(command) {
		return "", errors.New("failed to write all bytes")
	}
	nRead, err := s.socket.Read(b)
	if err != nil {
		return "", err
	}
	retString := string(b[:nRead])
	return retString, nil
}

func (s *ST_IP) Close() error {
	return s.socket.Close()
}

type ST_RS485 struct {
}
