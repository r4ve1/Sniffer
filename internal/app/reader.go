package app

import (
	"fmt"

	"github.com/google/gopacket"

	"sniffer/internal/cache"
	"sniffer/internal/reader"
)

const (
	readerNotInitialized = "reader not initialized"
	cacheNotInitialized  = "cache not initialized"
)

func (it *T) StartReader(filter string) error {
	if it.session.filename == "" {
		return fmt.Errorf(captureNotInitialized)
	}
	// reset view
	it.v.Reset()
	it.session.reader = reader.New(it.log, it.session.filename, it.v, cache.NewInMem)
	var err error
	it.session.cache, err = it.session.reader.Start(filter)
	if err != nil {
		return err
	}
	return nil
}

func (it *T) StopReader() error {
	if it.session.reader == nil {
		return fmt.Errorf(readerNotInitialized)
	}
	err := it.session.reader.Stop()
	if err != nil {
		return err
	}
	it.log.Debug("Reader stopped")
	return nil
}

func (it *T) Get(i int) (gopacket.Packet, error) {
	if it.session.cache == nil {
		return nil, fmt.Errorf(cacheNotInitialized)
	}
	return it.session.cache.Get(i), nil
}
