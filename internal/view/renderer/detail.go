package renderer

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Detail struct {
	Ethernet ethernet
	Ip       ip
	Tcp      tcp
}

type ethernet struct {
	SrcMAC       string
	DstMAC       string
	EthernetType string
}

type ip struct {
	Version uint8
	SrcIP   string
	DstIP   string
}

type tcp struct {
	SrcPort string
	DstPort string
	Seq     uint32
}

func (it *Detail) AddLayer(layer gopacket.Layer) {
	if layer == nil {
		return
	}
	switch layer.(type) {
	case *layers.Ethernet:
		eth := layer.(*layers.Ethernet)
		it.Ethernet.SrcMAC = eth.SrcMAC.String()
		it.Ethernet.DstMAC = eth.DstMAC.String()
		it.Ethernet.EthernetType = eth.EthernetType.String()
	case *layers.IPv4:
		ip := layer.(*layers.IPv4)
		it.Ip.Version = ip.Version
		it.Ip.SrcIP = ip.SrcIP.String()
		it.Ip.DstIP = ip.DstIP.String()
	case *layers.IPv6:
		ip := layer.(*layers.IPv6)
		it.Ip.Version = ip.Version
		it.Ip.SrcIP = ip.SrcIP.String()
		it.Ip.DstIP = ip.DstIP.String()
	case *layers.TCP:
		tcp := layer.(*layers.TCP)
		it.Tcp.SrcPort = tcp.SrcPort.String()
		it.Tcp.DstPort = tcp.DstPort.String()
		it.Tcp.Seq = tcp.Seq
	}
}
