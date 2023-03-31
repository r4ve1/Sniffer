package view

import (
	"github.com/google/gopacket"

	"sniffer/internal/view/renderer"
)

func (it *View) packet2Detail(packet gopacket.Packet) *renderer.Detail {
	detail := &renderer.Detail{}
	for _, layer := range packet.Layers() {
		detail.AddLayer(layer)
	}
	return detail
}
