package view

import (
	"context"
	"fmt"

	"github.com/google/gopacket"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type T struct {
	log               logger.Logger
	appCtx            context.Context
	cnt               int
	firstPktTimestamp int64
}

func New(appCtx context.Context, log logger.Logger) *T {
	it := &T{}
	it.appCtx = appCtx
	it.log = log
	return it
}

func (it *T) Display(packet gopacket.Packet) {
	runtime.EventsEmit(it.appCtx, "packet", it.toPacketBrief(packet))
}

func (it *T) Reset() {
	it.cnt = 0
}

type PacketBrief struct {
	No          int
	Timestamp   int64
	Length      int
	Source      string
	Destination string
	Protocol    string
}

func (it *T) toPacketBrief(packet gopacket.Packet) *PacketBrief {
	it.cnt++
	brief := &PacketBrief{}
	brief.No = it.cnt
	if brief.No == 1 {
		brief.Timestamp = 0
		it.firstPktTimestamp = packet.Metadata().Timestamp.UnixNano()
	} else {
		brief.Timestamp = packet.Metadata().Timestamp.UnixNano() - it.firstPktTimestamp
	}
	brief.Length = packet.Metadata().Length

	// determine flow
	var flow gopacket.Flow
	if packet.NetworkLayer() != nil {
		flow = packet.NetworkLayer().NetworkFlow()
	} else if packet.LinkLayer() != nil {
		flow = packet.LinkLayer().LinkFlow()
	} else {
		it.log.Error(fmt.Sprintf("No flow found, packet: %v", packet))
	}
	brief.Source = flow.Src().String()
	brief.Destination = flow.Dst().String()

	// determine protocol
	for i := len(packet.Layers()) - 1; i >= 0; i-- {
		t := packet.Layers()[i].LayerType()
		if t != gopacket.LayerTypePayload {
			brief.Protocol = t.String()
			break
		}
	}
	return brief
}
