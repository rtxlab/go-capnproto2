package rpc

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"zombiezen.com/go/capnproto2"
	rpccapnp "zombiezen.com/go/capnproto2/std/capnp/rpc"
)

// A Sender delivers Cap'n Proto RPC messages to another vat.
type Sender interface {
	// NewMessage allocates a new message to be sent over the transport.
	// The caller must call one of send or cancel.  send will take its
	// cancelation and deadline from ctx.  The returned message should not
	// be used after calling send or cancel.
	//
	// The behavior of calling NewMessage or CloseSend before calling send
	// or cancel on the previous call of NewMessage is undefined.
	NewMessage(ctx context.Context) (_ rpccapnp.Message, send func() error, cancel func(), _ error)

	// CloseSend releases any resources associated with the sender.
	CloseSend() error
}

// A Receiver receives Cap'n Proto RPC messages from another vat.
type Receiver interface {
	// RecvMessage waits to receive a message and returns it.
	// The returned message is only valid until the next call to
	// RecvMessage or CloseRecv.
	RecvMessage(ctx context.Context) (rpccapnp.Message, error)

	// CloseRecv releases any resources associated with the receiver and
	// ends any unfinished RecvMessage call.
	CloseRecv() error
}

// StreamTransport serializes and deserializes unpacked Cap'n Proto
// messages on a byte stream.  StreamTransport adds no buffering beyond
// what its underlying stream has.
//
// Sender methods on StreamTransport cannot be called concurrently with
// each other and Receiver methods on StreamTransport cannot be called
// concurrently with each other.  However, it is safe to call Sender
// methods concurrently with Receiver methods.
type StreamTransport struct {
	// Send
	enc      *capnp.Encoder
	deadline writeDeadlineSetter
	// Receive
	recv Receiver
	// Close
	c  io.Closer
	cw writeCloser

	mu     sync.Mutex
	closes uint8
}

// NewStreamTransport creates a new transport that reads and writes to rwc.
// Closing the transport will close rwc.
//
// If rwc has a SetWriteDeadline method, it will be used when a message
// is sent.  If rwc has CloseRead/CloseWrite methods, those will be used
// during CloseRecv/CloseSend.  Regardless, Close will be called once
// CloseRecv and CloseSend have both been called.
func NewStreamTransport(rwc io.ReadWriteCloser) *StreamTransport {
	d, _ := rwc.(writeDeadlineSetter)
	cw, _ := rwc.(writeCloser)
	s := &StreamTransport{
		enc:      capnp.NewEncoder(rwc),
		deadline: d,
		c:        rwc,
		cw:       cw,
	}
	dec := capnp.NewDecoder(rwc)
	dec.ReuseBuffer()
	if c, ok := rwc.(readCloser); ok {
		s.recv = closerReceiver{dec, c}
	} else {
		s.recv = signalReceiver{dec, make(chan struct{})}
	}
	return s
}

// NewMessage allocates a new message to be sent.  The send function may
// make multiple calls to Write on the underlying writer.
func (s *StreamTransport) NewMessage(ctx context.Context) (_ rpccapnp.Message, send func() error, cancel func(), _ error) {
	// TODO(soon): reuse memory
	msg, seg, _ := capnp.NewMessage(capnp.MultiSegment(nil))
	rmsg, _ := rpccapnp.NewRootMessage(seg)
	send = func() error {
		if s.deadline != nil {
			// TODO(someday): log errors
			if d, ok := ctx.Deadline(); ok {
				s.deadline.SetWriteDeadline(d)
			} else {
				s.deadline.SetWriteDeadline(time.Time{})
			}
		}
		return s.enc.Encode(msg)
	}
	cancel = func() {}
	return rmsg, send, cancel, nil
}

// CloseSend calls CloseWrite, if present, on the underlying
// io.ReadWriteCloser.  If CloseRecv was called before this function,
// then the underlying io.ReadWriteCloser is closed.
func (s *StreamTransport) CloseSend() error {
	s.mu.Lock()
	if s.closes&1 == 1 {
		s.mu.Unlock()
		return errors.New("rpc stream transport: send already closed")
	}
	s.closes |= 1
	done := s.closes == 3
	s.mu.Unlock()

	werr := s.cw.CloseWrite()
	if !done {
		return werr
	}
	cerr := s.c.Close()
	if cerr != nil {
		return cerr
	}
	if werr != nil {
		return werr
	}
	return nil
}

// RecvMessage reads the next message from the underlying reader.
// The cancelation and deadline from ctx is ignored, but RecvMessage
// will return early if CloseRecv is called.
func (s *StreamTransport) RecvMessage(ctx context.Context) (rpccapnp.Message, error) {
	return s.recv.RecvMessage(ctx)
}

// CloseRecv calls CloseRead, if present, on the underlying
// io.ReadWriteCloser.  If CloseSend was called before this function,
// then the underlying io.ReadWriteCloser is closed.
func (s *StreamTransport) CloseRecv() error {
	s.mu.Lock()
	if s.closes&2 == 2 {
		s.mu.Unlock()
		return errors.New("rpc stream transport: receive already closed")
	}
	s.closes |= 2
	done := s.closes == 3
	s.mu.Unlock()

	rerr := s.recv.CloseRecv()
	if !done {
		return rerr
	}
	cerr := s.c.Close()
	if cerr != nil {
		return cerr
	}
	if rerr != nil {
		return rerr
	}
	return nil
}

// closerReceiver receives messages from a decoder, relying on a
// readCloser to interrupt the underlying io.Reader.
type closerReceiver struct {
	dec    *capnp.Decoder
	closer readCloser
}

func (cr closerReceiver) RecvMessage(ctx context.Context) (rpccapnp.Message, error) {
	msg, err := cr.dec.Decode()
	if err != nil {
		return rpccapnp.Message{}, err
	}
	return rpccapnp.ReadRootMessage(msg)
}

func (cr closerReceiver) CloseRecv() error {
	return cr.closer.CloseRead()
}

// signalReceiver receives messages from a decoder, abandoning a Decode
// once CloseRecv is called.  It is assumed that the caller will then
// eventually interrupt the read, usually by calling Close on the
// underlying io.ReadCloser.
type signalReceiver struct {
	dec   *capnp.Decoder
	close chan struct{}
}

func (sr signalReceiver) RecvMessage(ctx context.Context) (rpccapnp.Message, error) {
	select {
	case <-sr.close:
		return rpccapnp.Message{}, errors.New("RPC stream transport: receive on closed receiver")
	default:
	}
	var msg *capnp.Message
	var err error
	read := make(chan struct{})
	go func() {
		msg, err = sr.dec.Decode()
		close(read)
	}()
	select {
	case <-read:
	case <-sr.close:
		return rpccapnp.Message{}, errors.New("RPC stream transport: receive on closed receiver")
	}
	if err != nil {
		return rpccapnp.Message{}, err
	}
	return rpccapnp.ReadRootMessage(msg)
}

func (sr signalReceiver) CloseRecv() error {
	close(sr.close)
	return nil
}

// Optional interfaces that io.ReadWriteClosers could implement.
// See net.TCPConn for docs.

type writeDeadlineSetter interface {
	SetWriteDeadline(t time.Time) error
}

type readCloser interface {
	CloseRead() error
}

type writeCloser interface {
	CloseWrite() error
}
