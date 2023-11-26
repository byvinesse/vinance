package application

import (
	"github.com/vincentkdeli/vinance-backend/pkg/validator"
)

type App struct {
}

func NewApp() *App {
	validator.Init()

	return &App{}
}
