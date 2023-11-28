package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vincentkdeli/vinance-backend/cmd/application"
	"github.com/vincentkdeli/vinance-backend/handler/server"
	"github.com/vincentkdeli/vinance-backend/pkg/errors"
)

func Run(app *application.App) {
	port := os.Getenv("PORT")
	if port == "" {
		panic("no PORT set")
	}

	e := echo.New()

	// Middlewares
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "Content-Length", "Origin"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))
	e.HTTPErrorHandler = errors.Middleware

	// Init Handler
	h := server.NewHandler(app)

	initRoutes(e, app, h)

	e.Logger.Fatal(e.Start(":" + port))
}

func initRoutes(e *echo.Echo, app *application.App, h *server.Handler) {
	// Healthcheck
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	v1 := e.Group("/v1")

	// Auth
	v1.POST("/register", h.Register)
	v1.POST("/login", h.Login)
}
