package server

import "github.com/vincentkdeli/vinance-backend/cmd/application"

type Handler struct {
	app *application.App
}

func NewHandler(app *application.App) *Handler {
	return &Handler{app}
}
