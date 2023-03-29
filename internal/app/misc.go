package app

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

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
