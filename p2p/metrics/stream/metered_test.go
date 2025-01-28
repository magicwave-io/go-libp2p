package meterstream

import (
	"io"
	"io/ioutil"
	"testing"

	randbo "github.com/dustin/randbo"
	peer "github.com/ipfs/go-libp2p-peer"
	inet "github.com/ipfs/go-libp2p/p2p/net"
	protocol "github.com/ipfs/go-libp2p/p2p/protocol"
)

type FakeStream struct {
	ReadBuf io.Reader
	inet.Stream
}

func (fs *FakeStream) Read(b []byte) (int, error) {
	return fs.ReadBuf.Read(b)
}

func (fs *FakeStream) Write(b []byte) (int, error) {
	return len(b), nil
}

func TestCallbacksWork(t *testing.T) {
	fake := new(FakeStream)

	var sent int64
	var recv int64

	sentCB := func(n int64, proto protocol.ID, p peer.ID) {
		sent += n
	}

	recvCB := func(n int64, proto protocol.ID, p peer.ID) {
		recv += n
	}

	ms := newMeteredStream(fake, protocol.ID("TEST"), peer.ID("PEER"), recvCB, sentCB)

	toWrite := int64(100000)
	toRead := int64(100000)

	fake.ReadBuf = io.LimitReader(randbo.New(), toRead)
	writeData := io.LimitReader(randbo.New(), toWrite)

	n, err := io.Copy(ms, writeData)
	if err != nil {
		t.Fatal(err)
	}

	if n != toWrite {
		t.Fatal("incorrect write amount")
	}

	if toWrite != sent {
		t.Fatal("incorrectly reported writes", toWrite, sent)
	}

	n, err = io.Copy(ioutil.Discard, ms)
	if err != nil {
		t.Fatal(err)
	}

	if n != toRead {
		t.Fatal("incorrect read amount")
	}

	if toRead != recv {
		t.Fatal("incorrectly reported reads")
	}
}
