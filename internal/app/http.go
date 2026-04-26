package app

import (
	"github.com/TehranTime/tehtime-core/internal/config"
	"github.com/gofiber/fiber/v2"
)

func newHttp(cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()
	_ = cfg
	return app, nil
}
