package capture

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Plugin interface {
	OnLiveOpen(handle *pcap.Handle)
	OnPacket(packet gopacket.Packet)
	OnStopped()
}
