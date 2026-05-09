package server

import (
	"net/http"
	"os"

	"github.com/byvinesse/vinance-backend/cmd/application"
	"github.com/byvinesse/vinance-backend/handler/server"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	appmiddleware "github.com/byvinesse/vinance-backend/pkg/middleware"
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
	withAuth := appmiddleware.Authentication(
		app.Authenticator,
		appmiddleware.AuthConfig{
			AllowAuthorizationHeader: true,
		},
	)

	// Healthcheck
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/protected", func(c echo.Context) error {
		return response.Ok(c, true)
	}, withAuth)

	// Auth
	authRoute := e.Group("/auth")

	// Auth V1
	authV1Route := authRoute.Group("/v1")
	authV1Route.POST("/register", h.Register)
	authV1Route.POST("/login", h.Login)

	// Accounts
	accountsRoute := e.Group("/accounts")

	// Accounts V1
	accountsV1Route := accountsRoute.Group("/v1")
	accountsV1Route.POST("/_create", h.CreateAccount, withAuth)

	// Users
	usersRoute := e.Group("/users")

	// Users V1
	usersV1Route := usersRoute.Group("/v1")
	usersV1Route.GET("/profile", h.GetProfile, withAuth)

	// Categories
	categoriesRoute := e.Group("/categories")

	// Categories V1
	categoriesV1Route := categoriesRoute.Group("/v1")
	categoriesV1Route.GET("/complete", h.GetCompleteCategories, withAuth)

	// Records
	recordsRoute := e.Group("/records")

	// Records V1
	recordsV1Route := recordsRoute.Group("/v1")
	recordsV1Route.POST("/_create", h.CreateRecord, withAuth)
	recordsV1Route.GET("", h.GetRecords, withAuth)
}
