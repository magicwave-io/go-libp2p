# Datagram protocols with libp2p

This libp2p example is a form of readme-driven development for packet switching.
I'm using it to test out interfaces and conventions, and also to produce working code.

For now it's going to be the simplest possible echo server/client.

## New packages

- [ ] go-multigram
- [ ] go-multigram-select
- [ ] go-libp2p-wireguard
- [ ] go-udp-transport


## Affected packages

- [x] go-multiaddr-net  (manet.PacketConn)
- [ ] go-libp2p-transport  (tptiface.PacketTransport/Dialer/Conn)
- [ ] go-libp2p-swarm  (sw.AddPacketTransport(), sw.SetPacketHandler(), sw.WritePacket())
- [ ] go-libp2p-host  (hostiface.SetPacketHandler(), hostiface.WritePacket())
- [ ] go-libp2p-net  (netiface.PacketHandler)


## About wireguard-go

- Looks good, but way too tightly coupled
- main.go, device.go
- send.go, receive.go
- device routines/queues and peer routines/queues
- split up: packetconn, cli+uapi+tun+routing
