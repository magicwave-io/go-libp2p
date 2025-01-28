package conn

import (
	"io"
	"net"
	"time"

	ic "github.com/ipfs/go-libp2p-crypto"
	peer "github.com/ipfs/go-libp2p-peer"
	transport "github.com/ipfs/go-libp2p-transport"
	ma "github.com/jbenet/go-multiaddr"
	filter "github.com/libp2p/go-libp2p/p2p/net/filter"
)

type PeerConn interface {
	io.Closer

	// LocalPeer (this side) ID, PrivateKey, and Address
	LocalPeer() peer.ID
	LocalPrivateKey() ic.PrivKey
	LocalMultiaddr() ma.Multiaddr

	// RemotePeer ID, PublicKey, and Address
	RemotePeer() peer.ID
	RemotePublicKey() ic.PubKey
	RemoteMultiaddr() ma.Multiaddr
}

// Conn is a generic message-based Peer-to-Peer connection.
type Conn interface {
	PeerConn

	// ID is an identifier unique to this connection.
	ID() string

	// can't just say "net.Conn" cause we have duplicate methods.
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error

	io.Reader
	io.Writer
}

// Dialer is an object that can open connections. We could have a "convenience"
// Dial function as before, but it would have many arguments, as dialing is
// no longer simple (need a peerstore, a local peer, a context, a network, etc)
type Dialer struct {
	// LocalPeer is the identity of the local Peer.
	LocalPeer peer.ID

	// LocalAddrs is a set of local addresses to use.
	//LocalAddrs []ma.Multiaddr

	// Dialers are the sub-dialers usable by this dialer
	// selected in order based on the address being dialed
	Dialers []transport.Dialer

	// PrivateKey used to initialize a secure connection.
	// Warning: if PrivateKey is nil, connection will not be secured.
	PrivateKey ic.PrivKey

	// Wrapper to wrap the raw connection (optional)
	Wrapper WrapFunc

	fallback transport.Dialer
}

// Listener is an object that can accept connections. It matches net.Listener
type Listener interface {

	// Accept waits for and returns the next connection to the listener.
	Accept() (net.Conn, error)

	// Addr is the local address
	Addr() net.Addr

	// Multiaddr is the local multiaddr address
	Multiaddr() ma.Multiaddr

	// LocalPeer is the identity of the local Peer.
	LocalPeer() peer.ID

	SetAddrFilters(*filter.Filters)

	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error
}

// EncryptConnections is a global parameter because it should either be
// enabled or _completely disabled_. I.e. a node should only be able to talk
// to proper (encrypted) networks if it is encrypting all its transports.
// Running a node with disabled transport encryption is useful to debug the
// protocols, achieve implementation interop, or for private networks which
// -- for whatever reason -- _must_ run unencrypted.
var EncryptConnections = true
