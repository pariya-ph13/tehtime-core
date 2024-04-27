package app

import (
	"github.com/TehranTime/tehtime-core/internal/router"
	"go.uber.org/fx"
)

func Run() error {
	app := fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			newSentry,
			newHttp,
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
