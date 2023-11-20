package main

import "api.default.indicoinnovation.pt/pkg/app"

func main() {
	app.ApplicationInit()
	defer app.Inst.DB.Close()

	app.Inst.Server = route()

	// Listening to Server
	app.Setup()
}
