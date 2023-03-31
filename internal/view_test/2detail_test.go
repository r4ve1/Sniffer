package view_test

import (
	"testing"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"

	"sniffer/internal/reader"
	"sniffer/internal/view"
	"sniffer/internal/view/renderer"
)

func TestView_RenderPacketDetail(t *testing.T) {
	const samplePath = "D:/Desktop/http-chunked-gzip.pcap"
	consoleRenderer := &renderer.Console{}
	log := logger.NewDefaultLogger()
	vi := view.New(consoleRenderer, log)
	r := reader.New(log, samplePath, vi)
	r.Start("")
	time.Sleep(time.Second)
	vi.RenderPacketDetail(3)
}
