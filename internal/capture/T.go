package capture

import (
	"fmt"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"golang.org/x/exp/slices"
)

type T struct {
	log      logger.Logger
	dev      pcap.Interface
	filename string
	w        capture
	r        reading
}

func New(devName string, log logger.Logger) (*T, error) {
	it := &T{}
	it.log = log
	// make sure device exists
	devs, err := pcap.FindAllDevs()
	if err != nil {
		it.log.Error(fmt.Sprintf("Error listing devices: %s", err))
		return nil, err
	}
	i := slices.IndexFunc(devs, func(i pcap.Interface) bool {
		return i.Description == devName
	})
	if i == -1 {
		err := fmt.Errorf("device %s not found", devName)
		return nil, err
	}
	it.dev = devs[i]
	// set filename
	it.filename = fmt.Sprintf("%s.pcap", time.Now().Format("2006-01-02_15-04-05"))

	return it, nil
}
