package main

import "api.default.indicoinnovation.pt/pkg/app"

func main() {
	app.ApplicationInit()
	app.Inst.Server = route()

	// Listening to Server
	app.Setup()
}
