package main

import (
	"github.com/byvinesse/vinance-backend/cmd/application"
	"github.com/byvinesse/vinance-backend/cmd/server"
)

func main() {
	app := application.NewApp()

	server.Run(app)
}
