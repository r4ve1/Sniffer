package app

import (
	"fmt"
	"time"

	"sniffer/internal/capture"
	"sniffer/internal/capture/plugins/writer"
)

const captureNotInitialized = "capture not initialized"

type DeviceInfo struct {
	Description string
	Addresses   []string
}

func (it *T) StartCapture(device string) error {
	it.session.filename = fmt.Sprintf("%s.pcap", time.Now().Format("2006-01-02_15-04-05"))
	wp, err := writer.New(it.log, it.session.filename)
	if err != nil {
		return err
	}
	it.session.capture = capture.New(it.log, wp)
	return it.session.capture.Start(device)
}

func (it *T) PauseCapture() error {
	if it.session.capture == nil {
		return fmt.Errorf(captureNotInitialized)
	}
	return it.session.capture.Pause()
}

func (it *T) ResumeCapture() error {
	if it.session.capture == nil {
		return fmt.Errorf(captureNotInitialized)
	}
	return it.session.capture.Resume()
}

func (it *T) StopCapture() error {
	if it.session.capture == nil {
		return fmt.Errorf(captureNotInitialized)
	}
	err := it.session.capture.Stop()
	if err != nil {
		return err
	}
	it.log.Debug("Capture stopped")
	return nil
}
