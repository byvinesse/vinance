package main

import (
	"github.com/vincentkdeli/vinance-backend/cmd/application"
	"github.com/vincentkdeli/vinance-backend/cmd/server"
)

func main() {
	app := application.NewApp()

	server.Run(app)
}
