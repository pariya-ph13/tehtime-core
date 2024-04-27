package router

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// Router function to set up routes and manage server lifecycle
func Router(
	lc fx.Lifecycle,
	app *fiber.App,
) {
	// Define a route group and add a route
	app.Get("/v2/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "here we are",
			"data":    "data",
		})
	})

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

//func Router(
//	lc fx.Lifecycle,
//	logger ports.LoggerWithTraceID,
//	httpService ports.HttpServer,
//) {
//
//	httpService.SetRouteGroups("v2", nil, []ports.Route{
//		{
//			Method: http.MethodGet,
//			Path:   "/test",
//			Handler: func(c *ports.HttpContext) error {
//				logger.Info(c.Context(), "just a test log")
//				return c.JSON("here we are", "data")
//			},
//		},
//	})
//	lc.Append(
//		fx.Hook{
//			OnStart: func(ctx context.Context) error {
//				go func() {
//					err := httpService.Listen()
//					if err != nil {
//						log.Error(ctx, err.Error())
//					}
//				}()
//				return nil
//			},
//		})
//}
