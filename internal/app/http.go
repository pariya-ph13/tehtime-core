package app

import (
	"github.com/TehranTime/tehtime-core/internal/config"
	"github.com/gofiber/fiber/v2"
	html "github.com/gofiber/template/html/v2"
)

func newHttp(cfg *config.Config) (*fiber.App, error) {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{Views: engine, ViewsLayout: "layouts/main"})
	_ = cfg
	return app, nil
}
