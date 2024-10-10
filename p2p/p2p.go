package p2p

import (
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
)

var P2PProtocol = ma.Protocol{
	Code:  420,
	Name:  "p2p",
	VCode: ma.CodeToVarint(420),
}
var IpfsProtocol = ma.Protocol{
	Code:  421,
	Name:  "ipfs",
	VCode: ma.CodeToVarint(421),
}

var P2PCodec = &manet.NetCodec{
	NetAddrNetworks:  []string{"libp2p"},
	ProtocolName:     "p2p",
	ConvertMultiaddr: ConvertP2PMultiaddr,
	ParseNetAddr:     ParseP2PNetAddr,
}
var IpfsCodec = &manet.NetCodec{
	NetAddrNetworks:  []string{"libp2p+ipfs"},
	ProtocolName:     "ipfs",
	ConvertMultiaddr: ConvertIpfsMultiaddr,
	ParseNetAddr:     ParseIpfsNetAddr,
}

func init() {
	err := ma.AddProtocol(P2PProtocol)
	if err != nil {
		panic(fmt.Errorf("error registering libp2p protocol: %s", err))
	}
	err = ma.AddProtocol(IpfsProtocol)
	if err != nil {
		panic(fmt.Errorf("error registering libp2p+ipfs protocol: %s", err))
	}

	manet.RegisterNetCodec(P2PCodec)
	manet.RegisterNetCodec(IpfsCodec)
}
