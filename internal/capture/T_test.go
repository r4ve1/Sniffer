package capture

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

func TestT_Capture(t *testing.T) {
	capt, err := New("RZ608 Wi-Fi 6E 80MHz", logger.NewDefaultLogger())
	assert.Nil(t, err)
	err = capt.StartCapture()
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
	capt.PauseCapture()
	time.Sleep(1 * time.Second)
	err = capt.ResumeCapture()
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
}
