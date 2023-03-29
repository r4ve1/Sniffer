package app

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"

	"sniffer/internal/cache"
	"sniffer/internal/capture"
	"sniffer/internal/reader"
	"sniffer/internal/view"
)

type T struct {
	appCtx  context.Context
	log     logger.Logger
	v       view.I
	session session
}

type session struct {
	filename string
	capture  capture.I
	reader   reader.I
	cache    cache.I
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
	it.v = view.New(ctx, it.log)
}
