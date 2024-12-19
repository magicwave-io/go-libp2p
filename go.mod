module github.com/libp2p/go-libp2p

go 1.12

require (
	github.com/gogo/protobuf v1.3.1
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.4
	github.com/ipfs/go-detect-race v0.0.1
	github.com/ipfs/go-ipfs-util v0.0.2
	github.com/ipfs/go-log v1.0.4
	github.com/jbenet/go-cienv v0.1.0
	github.com/jbenet/goprocess v0.1.4
	github.com/libp2p/go-addr-util v0.0.2
	github.com/libp2p/go-conn-security-multistream v0.2.0
	github.com/libp2p/go-eventbus v0.2.1
	github.com/libp2p/go-libp2p-autonat v0.4.0
	github.com/libp2p/go-libp2p-blankhost v0.2.0
	github.com/libp2p/go-libp2p-circuit v0.4.0
	github.com/libp2p/go-libp2p-core v0.8.0
	github.com/libp2p/go-libp2p-discovery v0.5.0
	github.com/libp2p/go-libp2p-loggables v0.1.0
	github.com/libp2p/go-libp2p-mplex v0.4.1
	github.com/libp2p/go-libp2p-nat v0.0.6
	github.com/libp2p/go-libp2p-netutil v0.1.0
	github.com/libp2p/go-libp2p-noise v0.1.1
	github.com/libp2p/go-libp2p-peerstore v0.2.6
	github.com/libp2p/go-libp2p-swarm v0.4.0
	github.com/libp2p/go-libp2p-testing v0.4.0
	github.com/libp2p/go-libp2p-tls v0.1.3
	github.com/libp2p/go-libp2p-transport-upgrader v0.4.0
	github.com/libp2p/go-libp2p-yamux v0.5.1
	github.com/libp2p/go-msgio v0.0.6
	github.com/libp2p/go-netroute v0.1.3
	github.com/libp2p/go-stream-muxer-multistream v0.3.0
	github.com/libp2p/go-tcp-transport v0.2.1
	github.com/libp2p/go-ws-transport v0.4.0
	github.com/miekg/dns v1.1.31 // indirect
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multiaddr-dns v0.2.0
	github.com/multiformats/go-multistream v0.2.0
	github.com/onsi/ginkgo v1.12.1 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/whyrusleeping/mdns v0.0.0-20190826153040-b9b60ed33aa9
	golang.org/x/net v0.0.0-20200519113804-d87ec0cfa476 // indirect
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	google.golang.org/protobuf v1.23.0 // indirect
	paidpiper.com/go-libp2p-onion-transport v0.0.0-00010101000000-000000000000
)

replace github.com/libp2p/go-libp2p-core => ../go-libp2p-core

replace github.com/multiformats/go-multiaddr => ../go-multiaddr

replace github.com/multiformats/go-multiaddr-net => ../go-multiaddr-net

replace paidpiper.com/go-libp2p-onion-transport => ../go-libp2p-onion-transport

replace github.com/libp2p/go-eventbus => ../go-eventbus

replace github.com/libp2p/go-libp2p-swarm => ../go-libp2p-swarm
