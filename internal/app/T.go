package app

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/logger"

	"sniffer/internal/capture"
)

type T struct {
	appCtx context.Context
	log    logger.Logger
	cp     *capture.T
}

func New(log logger.Logger) *T {
	it := &T{}
	it.log = log
	return it
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (it *T) Startup(ctx context.Context) {
	it.appCtx = ctx
}

// Greet returns a greeting for the given DeviceInfo
func (it *T) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
