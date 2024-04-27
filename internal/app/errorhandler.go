package app

import (
	"github.com/TehranTime/tehtime-core/internal/config"
	conf "github.com/TehranTime/tehtime-core/internal/config"
)

func newSentry(cfg config.Config[conf.Config]) ports.ErrorHandler {
	return sentry.New(sentry.Config{
		Dsn:   cfg.GetConfig().Sentry.Dsn,
		Debug: cfg.GetConfig().Debug,
	})
}
