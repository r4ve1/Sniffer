package cache

import (
	"github.com/google/gopacket"
)

type I interface {
	Get(i int) gopacket.Packet
	Clear()
	Add(packet gopacket.Packet)
}
