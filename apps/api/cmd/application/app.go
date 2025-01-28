package application

import (
	"github.com/byvinesse/vinance-backend/config"
	auth "github.com/byvinesse/vinance-backend/pkg/authenticator"
	"github.com/byvinesse/vinance-backend/pkg/service"
	"github.com/byvinesse/vinance-backend/pkg/validator"
	"github.com/byvinesse/vinance-backend/repository/db"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type App struct {
	UserService service.IUserService

	Authenticator *auth.Authenticator
}

func NewApp() *App {
	authConfig := config.LoadAuth()
	authenticator := auth.NewAuthenticator(authConfig)

	// Init Database
	database := connectDatabase()

	// Init repository
	userRepo := db.NewUser(database)

	// Init service
	userService := service.NewUserService(userRepo, *authenticator)

	validator.Init()

	return &App{
		UserService: userService,

		Authenticator: authenticator,
	}
}

func connectDatabase() *sqlx.DB {
	databaseConfig := config.LoadDatabase()

	database, err := sqlx.Open("postgres", databaseConfig.URI)
	if err != nil {
		panic(err)
	}

	return database
}
