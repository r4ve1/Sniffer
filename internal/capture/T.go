package capture

import (
	"context"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"golang.org/x/exp/slices"
)

const SnapLen = 0x10000

type t struct {
	log     logger.Logger
	plugins []Plugin
	live    *pcap.Handle
	cancel  context.CancelFunc
}

func New(log logger.Logger, plugins ...Plugin) I {
	it := &t{}
	it.log = log
	it.plugins = plugins
	return it
}

func (it *t) Start(dev string) error {
	// make sure device exists
	devs, err := pcap.FindAllDevs()
	if err != nil {
		it.log.Error(fmt.Sprintf("Error listing devices: %s", err))
		return err
	}
	i := slices.IndexFunc(devs, func(i pcap.Interface) bool {
		return i.Description == dev
	})
	if i == -1 {
		err := fmt.Errorf("device %s not found", dev)
		return err
	}
	it.live, err = pcap.OpenLive(devs[i].Name, SnapLen, true, pcap.BlockForever)
	if err != nil {
		it.log.Error(fmt.Sprintf("Error opening device: %s", err))
		return err
	}
	for i := range it.plugins {
		it.plugins[i].OnLiveOpen(it.live)
	}
	return it.Resume()
}

func (it *t) Pause() error {
	if it.cancel != nil {
		it.cancel()
		it.cancel = nil
	}
	return nil
}

func (it *t) Resume() error {
	var ctx context.Context
	ctx, it.cancel = context.WithCancel(context.Background())

	// until context is done
	go func() {
		pcapSrc := gopacket.NewPacketSource(it.live, it.live.LinkType())
		for {
			select {
			case <-ctx.Done():
				return
			case pkt, ok := <-pcapSrc.Packets():
				if !ok {
					it.log.Error(fmt.Sprintf("Error capturing packet"))
					return
				}
				for i := range it.plugins {
					it.plugins[i].OnPacket(pkt)
				}
			}
		}
	}()
	return nil
}

func (it *t) Stop() error {
	err := it.Pause()
	if err != nil {
		return err
	}
	// close live
	it.live.Close()
	it.live = nil
	return nil
}
