package reader

import (
	"context"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/wailsapp/wails/v2/pkg/logger"

	"sniffer/internal/cache"
	"sniffer/internal/view"
)

type cacheFactory func() cache.I

type t struct {
	filename string
	log      logger.Logger
	v        view.I
	cf       cacheFactory

	cancel context.CancelFunc
}

func New(log logger.Logger, filename string, v view.I, cf cacheFactory) I {
	it := &t{}
	it.log = log
	it.filename = filename
	it.v = v
	it.cf = cf
	return it
}

func (it *t) Start(filter string) (cache.I, error) {
	it.log.Info("Opening file: " + it.filename)
	h, err := pcap.OpenOffline(it.filename)
	if err != nil {
		it.log.Error("Error opening pcap file: " + err.Error())
		return nil, err
	}
	err = h.SetBPFFilter(filter)
	if err != nil {
		it.log.Error("Error setting filter: " + err.Error())
		return nil, err
	} else {
		it.log.Info("Filter set: " + filter)
	}
	src := gopacket.NewPacketSource(h, h.LinkType())
	c := it.cf()
	var ctx context.Context
	ctx, it.cancel = context.WithCancel(context.Background())
	go func() {
		for pkt, err := src.NextPacket(); ; pkt, err = src.NextPacket() {
			select {
			case <-ctx.Done():
				return
			default:
				if err == nil && pkt != nil {
					c.Add(pkt)
					it.v.Display(pkt)
				}
			}
		}
	}()
	return c, nil
}

func (it *t) Stop() error {
	if it.cancel != nil {
		it.cancel()
		it.cancel = nil
	}
	it.log.Info("Stopped reading")
	return nil
}
