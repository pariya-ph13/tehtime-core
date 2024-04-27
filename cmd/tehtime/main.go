package main

import (
	"github.com/TehranTime/tehtime-core/internal/app"
	"os"
)

func main() {
	err := app.Run()
	if err != nil {
		os.Exit(1)
	}
}
