package renderer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Wails struct {
	appCtx context.Context
	log    logger.Logger
	briefs []*Brief
	mutex  sync.Mutex
}

const emitInterval = time.Millisecond * 200

func NewWails(appCtx context.Context, log logger.Logger) I {
	it := &Wails{}
	it.log = log
	it.appCtx = appCtx

	go it.emitBriefs()
	return it
}

func (it *Wails) emitBriefs() {
	ticker := time.NewTicker(emitInterval)
	for _ = range ticker.C {
		if len(it.briefs) != 0 {
			it.mutex.Lock()
			it.log.Debug(fmt.Sprintf("Emitting %d packets", len(it.briefs)))
			runtime.EventsEmit(it.appCtx, "packets", it.briefs)
			it.briefs = []*Brief{}
			it.mutex.Unlock()
		}
	}
}

func (it *Wails) RenderBrief(packet *Brief) {
	it.mutex.Lock()
	defer it.mutex.Unlock()
	it.briefs = append(it.briefs, packet)
}

func (it *Wails) RenderDetail(packet *Detail) {
	it.log.Debug("Emitting detail")
	runtime.EventsEmit(it.appCtx, "detail", packet)
}

func (it *Wails) Reset() {
	it.mutex.Lock()
	defer it.mutex.Unlock()
	it.briefs = []*Brief{}
	it.log.Debug("wails renderer reset")
}
