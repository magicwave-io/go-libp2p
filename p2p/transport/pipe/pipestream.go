package pipetransport

import (
	"fmt"
	"net"
	"time"

	mux "github.com/libp2p/go-libp2p-core/mux"
	streammux "github.com/libp2p/go-stream-muxer"
)

var ErrHalfClosed = fmt.Errorf("tried to write to stream that was half closed")
var ErrAlreadyHalfClosed = fmt.Errorf("tried to half close stream that was already half closed")

type PipeStream struct {
	inbound  net.Conn
	outbound net.Conn
	errchs   struct {
		in  <-chan error
		out chan<- error
	}
}

func NewPipeStream(inbound net.Conn, outbound net.Conn, errinch <-chan error, erroutch chan<- error) *PipeStream {
	return &PipeStream{
		inbound:  inbound,
		outbound: outbound,
		errchs: struct {
			in  <-chan error
			out chan<- error
		}{
			in:  errinch,
			out: erroutch,
		},
	}
}

func (s *PipeStream) Close() error {
	return s.outbound.Close()
}

func (s *PipeStream) Reset() error {
	s.errchs.out <- mux.ErrReset
	err1 := s.inbound.Close()
	err2 := s.outbound.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *PipeStream) Read(b []byte) (int, error) {
	n, err := s.inbound.Read(b)

	select {
	case reseterr := <-s.errchs.in:
		return n, reseterr
	default:
		return n, err
	}
}

func (s *PipeStream) Write(b []byte) (int, error) {
	n, err := s.outbound.Write(b)

	select {
	case reseterr := <-s.errchs.in:
		return n, reseterr
	default:
		fmt.Println("no probs here")
		return n, err
	}
}

func (s *PipeStream) SetDeadline(t time.Time) error {
	err := s.inbound.SetDeadline(t)
	if err != nil {
		return err
	}
	err = s.outbound.SetDeadline(t)
	return err
}

func (s *PipeStream) SetReadDeadline(t time.Time) error {
	return s.inbound.SetReadDeadline(t)
}

func (s *PipeStream) SetWriteDeadline(t time.Time) error {
	return s.outbound.SetWriteDeadline(t)
}

var _ streammux.Stream = (*PipeStream)(nil)
