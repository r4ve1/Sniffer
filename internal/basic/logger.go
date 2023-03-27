package basic

import (
	"github.com/wailsapp/wails/v2/pkg/logger"
)

func NewLogger() logger.Logger {
	return logger.NewDefaultLogger()
}
