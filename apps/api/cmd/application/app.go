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
	UserService     service.IUserService
	AccountService  service.IAccountService
	CategoryService service.ICategoryService

	Authenticator *auth.Authenticator
}

func NewApp() *App {
	authConfig := config.LoadAuth()
	authenticator := auth.NewAuthenticator(authConfig)

	// Init Database
	database := connectDatabase()

	// Init repository
	userRepo := db.NewUser(database)
	accountRepo := db.NewAccount(database)
	categoryRepo := db.NewCategory(database)

	// Init service
	userService := service.NewUserService(userRepo, *authenticator)
	accountService := service.NewAccountService(accountRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	validator.Init()

	return &App{
		UserService:     userService,
		AccountService:  accountService,
		CategoryService: categoryService,

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
