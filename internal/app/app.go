package app

import (
	timehandler "github.com/TehranTime/tehtime-core/internal/handler/time"
	"github.com/TehranTime/tehtime-core/internal/router"
	timesvc "github.com/TehranTime/tehtime-core/internal/service/time"
	"go.uber.org/fx"
)

func Run() error {
	app := fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			newHttp,
			timesvc.NewService,
			timehandler.NewHandler,
		),
		fx.Invoke(
			router.Router,
		),
	)
	err := app.Err()
	if err != nil {
		panic(err)
	}
	app.Run()
	return nil
}
