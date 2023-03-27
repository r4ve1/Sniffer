package app

import (
	"fmt"

	"github.com/google/gopacket/pcap"

	"sniffer/internal/capture"
)

type DeviceInfo struct {
	Description string
	Addresses   []string
}

func (it *T) ListDevices() ([]DeviceInfo, error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		it.log.Error(fmt.Sprintf("Error listing devices: %s", err))
		return nil, err
	}
	dis := make([]DeviceInfo, len(devices))
	for i, dev := range devices {
		addresses := make([]string, len(dev.Addresses))
		for j, addr := range dev.Addresses {
			addresses[j] = addr.IP.String()
		}
		dis[i] = DeviceInfo{
			Description: dev.Description,
			Addresses:   addresses,
		}
	}
	return dis, nil
}

func (it *T) Start(device, filter string) error {
	var err error
	it.cp, err = capture.New(device, it.log)
	if err != nil {
		return err
	}
	err = it.cp.StartCapture()
	if err != nil {
		return err
	}
	err = it.cp.StartReading(it.appCtx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (it *T) Stop() error {
	if it.cp == nil {
		return nil
	}
	it.cp.StopReading()
	err := it.cp.StopCapture()
	if err != nil {
		return err
	}
	it.cp = nil
	return nil
}

func (it *T) PauseCapture() error {
	it.cp.PauseCapture()
	return nil
}

func (it *T) ResumeCapture() error {
	if err := it.cp.ResumeCapture(); err != nil {
		return err
	}
	return nil
}

func (it *T) RestartReading(filter string) error {
	if err := it.cp.StartReading(it.appCtx, filter); err != nil {
		return err
	}
	return nil
}
