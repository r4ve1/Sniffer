package app

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"

	"sniffer/internal/capture"
	"sniffer/internal/reader"
	"sniffer/internal/view"
	"sniffer/internal/view/renderer"
)

type T struct {
	appCtx  context.Context
	log     logger.Logger
	v       *view.View
	session session
}

type session struct {
	filename string
	capture  capture.I
	reader   reader.I
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
	wailsRenderer := renderer.NewWails(ctx, it.log)
	it.v = view.New(wailsRenderer, it.log)
}
