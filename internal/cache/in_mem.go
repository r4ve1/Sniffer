package cache

import (
	"github.com/google/gopacket"
)

type InMem struct {
	packets []gopacket.Packet
}

func NewInMem() I {
	it := &InMem{}
	it.packets = make([]gopacket.Packet, 0)
	return it
}

func (it *InMem) Get(i int) gopacket.Packet {
	return it.packets[i]
}

func (it *InMem) Clear() {
	it.packets = make([]gopacket.Packet, 0)
}

func (it *InMem) Add(packet gopacket.Packet) {
	it.packets = append(it.packets, packet)
}
