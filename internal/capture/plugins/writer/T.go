package writer

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

type T struct {
	filename string
	log      logger.Logger
	file     *os.File
	pcap     *pcapgo.Writer
}

func New(log logger.Logger, filename string) (*T, error) {
	it := &T{}
	it.log = log
	it.filename = filename
	// create pcap file
	var err error
	it.file, err = os.Create(it.filename)
	if err != nil {
		it.log.Error(fmt.Sprintf("Error creating file: %s", err))
		return nil, err
	}
	return it, nil
}

func (it *T) OnLiveOpen(live *pcap.Handle) {
	if it.pcap != nil {
		// only do once
		return
	}
	it.pcap = pcapgo.NewWriter(it.file)
	// write file header
	err := it.pcap.WriteFileHeader(65536, live.LinkType())
	if err != nil {
		it.log.Error(fmt.Sprintf("Error writing file header: %s", err))
	}
}

func (it *T) OnPacket(packet gopacket.Packet) {
	err := it.pcap.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	if err != nil {
		it.log.Error(fmt.Sprintf("Error writing packet: %s", err))
	}
}

func (it *T) OnStopped() {
	// close writer
	err := it.file.Close()
	if err != nil {
		it.log.Error(fmt.Sprintf("Error closing file: %s", err))
		return
	}
	it.log.Info(fmt.Sprintf("Capture stopped and saved to %s", it.filename))
}
