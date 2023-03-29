package view

import (
	"github.com/google/gopacket"
)

type I interface {
	Display(packet gopacket.Packet)
	Reset()
}
