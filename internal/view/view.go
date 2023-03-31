package view

import (
	"net/http"

	"github.com/google/gopacket"
	"github.com/wailsapp/wails/v2/pkg/logger"

	"sniffer/internal/cache"
	"sniffer/internal/view/renderer"
)

type View struct {
	log               logger.Logger
	renderer          renderer.I
	last              gopacket.Packet
	cnt               int
	cache             cache.I
	firstPktTimestamp int64
}

func New(renderer renderer.I, log logger.Logger) *View {
	it := &View{}
	it.log = log
	it.renderer = renderer
	it.cache = cache.NewInMem()
	return it
}

func (it *View) RenderPacketBrief(packet gopacket.Packet) {
	it.renderer.RenderBrief(it.packet2Brief(packet))
	it.cache.Add(packet)
	it.last = packet
}

func (it *View) RenderHttpReqBrief(flow gopacket.Flow, req *http.Request) {
	it.renderer.RenderBrief(it.httpReq2Brief(flow, req))
}

func (it *View) RenderHttpRespBrief(flow gopacket.Flow, resp *http.Response) {
	it.renderer.RenderBrief(it.httpResp2Brief(flow, resp))
}

func (it *View) RenderPacketDetail(i int) {
	pkt := it.cache.Get(i)
	if pkt == nil {
		it.log.Error("Packet not found in cache")
		return
	}
	it.renderer.RenderDetail(it.packet2Detail(pkt))
}

func (it *View) Reset() {
	it.cnt = 0
	it.cache.Clear()
	it.renderer.Reset()
}
