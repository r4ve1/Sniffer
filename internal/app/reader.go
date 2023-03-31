package app

import (
	"fmt"

	"sniffer/internal/reader"
)

const (
	readerNotInitialized = "reader not initialized"
)

func (it *T) StartReader(filter string) error {
	if it.session.filename == "" {
		return fmt.Errorf(captureNotInitialized)
	}
	// reset view
	it.v.Reset()
	it.session.reader = reader.New(it.log, it.session.filename, it.v)
	var err error
	err = it.session.reader.Start(filter)
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

func (it *T) GetDetail(i int) error {
	it.v.RenderPacketDetail(i)
	return nil
}
