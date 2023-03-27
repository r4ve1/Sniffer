package capture

import (
	"context"
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

type capture struct {
	live   *pcap.Handle
	ctx    context.Context
	cancel context.CancelFunc
	file   *os.File
	pcap   *pcapgo.Writer
}

func (it *T) StartCapture() (err error) {
	// open device
	it.w.live, err = pcap.OpenLive(it.dev.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		it.log.Error(fmt.Sprintf("Error opening device: %s", err))
		return err
	}

	// create pcap file
	it.w.file, err = os.Create(it.filename)
	if err != nil {
		it.log.Error(fmt.Sprintf("Error creating file: %s", err))
		return err
	}

	// write file header
	it.w.pcap = pcapgo.NewWriter(it.w.file)
	err = it.w.pcap.WriteFileHeader(65536, it.w.live.LinkType())
	if err != nil {
		it.log.Error(fmt.Sprintf("Error writing file header: %s", err))
		return err
	}

	return it.ResumeCapture()
}

func (it *T) PauseCapture() {
	if it.w.cancel != nil {
		it.w.cancel()
		it.w.cancel = nil
	}
}

func (it *T) ResumeCapture() error {
	it.w.ctx, it.w.cancel = context.WithCancel(context.Background())

	// until context is done
	go func() {
		pcapSrc := gopacket.NewPacketSource(it.w.live, it.w.live.LinkType())
		for {
			select {
			case <-it.w.ctx.Done():
				return
			default:
				// read packet
				pkt, err := pcapSrc.NextPacket()
				if err != nil {
					it.log.Error(fmt.Sprintf("Error reading packet: %s", err))
					return
				}
				// write packet
				err = it.w.pcap.WritePacket(pkt.Metadata().CaptureInfo, pkt.Data())
				if err != nil {
					it.log.Error(fmt.Sprintf("Error writing packet: %s", err))
					return
				}
			}
		}
	}()
	return nil
}

func (it *T) StopCapture() error {
	it.PauseCapture()
	// close live
	it.w.live.Close()
	it.w.live = nil
	// close writer
	err := it.w.file.Close()
	if err != nil {
		it.log.Error(fmt.Sprintf("Error closing file: %s", err))
		return err
	}
	it.log.Info(fmt.Sprintf("Capture stopped and saved to %s", it.filename))
	return nil
}
