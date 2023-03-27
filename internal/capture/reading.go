package capture

import (
	"context"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type reading struct {
	handle  *pcap.Handle
	packets []gopacket.Packet
	ctx     context.Context
	cancel  context.CancelFunc
}

func (it *T) StartReading(appCtx context.Context, filter string) error {
	src, err := it.openReader(filter)
	if err != nil {
		return err
	}
	it.r.packets = make([]gopacket.Packet, 0)
	it.r.ctx, it.r.cancel = context.WithCancel(context.Background())
	go func() {
		for pkt, err := src.NextPacket(); ; pkt, err = src.NextPacket() {
			select {
			case <-it.r.ctx.Done():
				return
			default:
				if err == nil && pkt != nil {
					it.r.packets = append(it.r.packets, pkt)
					layers := pkt.Layers()
					runtime.EventsEmit(appCtx, "packet", layers[len(layers)-1])
				}
			}
		}
	}()
	return nil
}

func (it *T) StopReading() {
	if it.r.cancel != nil {
		it.r.cancel()
		it.r.cancel = nil
	}
	if it.r.handle != nil {
		it.r.handle.Close()
	}
	it.log.Info("Stopped reading")
}

func (it *T) openReader(filter string) (*gopacket.PacketSource, error) {
	var err error
	it.log.Info("Opening file: " + it.filename)
	it.r.handle, err = pcap.OpenOffline(it.filename)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file: %s", err))
		return nil, err
	}
	err = it.r.handle.SetBPFFilter(filter)
	if err != nil {
		log.Error(fmt.Sprintf("Error setting filter: %s", err))
		return nil, err
	}
	return gopacket.NewPacketSource(it.r.handle, it.r.handle.LinkType()), nil
}
