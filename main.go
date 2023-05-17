package main

import (
	"muscurdig/app"
)

func main() {
	a := app.NewApp()
	a.Run()

	a.Cleanup()
}
