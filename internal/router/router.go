package router

import (
	"context"
	"fmt"

	timehandler "github.com/TehranTime/tehtime-core/internal/handler/time"
	web "github.com/TehranTime/tehtime-core/internal/handler/web"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// Router function to set up routes and manage server lifecycle
func Router(
	lc fx.Lifecycle,
	app *fiber.App,
	th *timehandler.Handler,
	wh *web.Handler,
) {
	// Define a route group and add a route
	app.Get("/time", th.GetTime)
	app.Get("/", wh.Index)
	app.Get("/calendar", wh.CalendarPartial)

	// Manage the server lifecycle
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := app.Listen(":3000"); err != nil {
						fmt.Println("error starting server:", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return app.Shutdown()
			},
		},
	)
}
