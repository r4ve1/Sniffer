package reader

import (
	"context"
	"strconv"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/wailsapp/wails/v2/pkg/logger"

	"sniffer/internal/view"
)

type Reader struct {
	filename  string
	log       logger.Logger
	v         *view.View
	cancel    context.CancelFunc
	httpPorts map[uint16]bool
}

func New(log logger.Logger, filename string, v *view.View) I {
	it := &Reader{}
	it.log = log
	it.filename = filename
	it.v = v
	it.httpPorts = map[uint16]bool{
		80:   true,
		8080: true,
	}
	return it
}

func (it *Reader) Start(filter string) error {
	it.log.Info("Opening file: " + it.filename)
	h, err := pcap.OpenOffline(it.filename)
	if err != nil {
		it.log.Error("Error opening pcap file: " + err.Error())
		return err
	}
	err = h.SetBPFFilter(filter)
	if err != nil {
		it.log.Error("Error setting filter: " + err.Error())
		return err
	} else {
		it.log.Info("Filter set: " + filter)
	}

	// Set up assembly
	streamFactory := &httpStreamFactory{
		v: it.v,
		isHttpReq: func(flow gopacket.Flow) bool {
			dst := flow.Dst().String()
			dstPort, _ := strconv.Atoi(dst)
			return it.httpPorts[uint16(dstPort)]
		},
	}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)
	src := gopacket.NewPacketSource(h, h.LinkType())
	var ctx context.Context
	ctx, it.cancel = context.WithCancel(context.Background())
	go func() {
		for pkt, err := src.NextPacket(); ; pkt, err = src.NextPacket() {
			select {
			case <-ctx.Done():
				return
			default:
				if err == nil && pkt != nil && pkt.ErrorLayer() == nil {
					it.v.RenderPacketBrief(pkt)
					// http
					if pkt.TransportLayer() != nil {
						tcp, ok := pkt.TransportLayer().(*layers.TCP)
						if ok && it.IsHttp(tcp) {
							assembler.AssembleWithTimestamp(pkt.NetworkLayer().NetworkFlow(), tcp, pkt.Metadata().Timestamp)
						}
					}
				}
			}
		}
	}()
	return nil
}

func (it *Reader) Stop() error {
	if it.cancel != nil {
		it.cancel()
		it.cancel = nil
	}
	it.log.Info("Stopped reading")
	return nil
}

func (it *Reader) IsHttp(tcp *layers.TCP) bool {
	src, dst := uint16(tcp.SrcPort), uint16(tcp.DstPort)
	return it.httpPorts[src] || it.httpPorts[dst]
}
