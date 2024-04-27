package app

import (
	"github.com/TehranTime/tehtime-core/internal/config"
	conf "github.com/TehranTime/tehtime-core/internal/config"
)

func newHttp(cfg config.Config[conf.Config],
	sentry ports.ErrorHandler) (ports.HttpServer, error) {
	return fiber.New(
		cfg.GetConfig().Debug,
		cfg.GetConfig().HttpServer.Address,
		uuid.New(),
		sentry,
	), nil
}
