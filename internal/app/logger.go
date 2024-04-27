package app

import (
	"github.com/TehranTime/tehtime-core/internal/config"
)

func NewLogger(
	config config.Config[conf.Config],
	sentry ports.ErrorHandler,
) ports.LoggerWithTraceID {
	return zap.New(config.GetConfig().LogLevel, sentry)
}
