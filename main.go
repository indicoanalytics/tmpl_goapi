package main

import (
	"api.default.indicoinnovation.pt/app/appinstance"
	"api.default.indicoinnovation.pt/pkg/app"
)

func main() {
	app.ApplicationInit()
	defer appinstance.Data.DB.Close()

	appinstance.Data.Server = route()

	// Listening to Server
	app.Setup()
}
