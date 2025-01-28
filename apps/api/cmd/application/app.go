package application

import (
	"github.com/jmoiron/sqlx"
	"github.com/vincentkdeli/vinance-backend/config"
	auth "github.com/vincentkdeli/vinance-backend/pkg/authenticator"
	"github.com/vincentkdeli/vinance-backend/pkg/service"
	"github.com/vincentkdeli/vinance-backend/pkg/validator"
	"github.com/vincentkdeli/vinance-backend/repository/db"

	_ "github.com/lib/pq"
)

type App struct {
	AuthService   service.IAuthService
	MemberService service.IMemberService

	Authenticator *auth.Authenticator
}

func NewApp() *App {
	authConfig := config.LoadAuth()
	authenticator := auth.NewAuthenticator(authConfig)

	// Init Database
	database := connectDatabase()

	// Init repository
	authRepo := db.NewAuth(database)
	memberRepo := db.NewMember(database)

	// Init service
	authService := service.NewAuthService(authRepo, *authenticator)
	memberService := service.NewMemberService(memberRepo)

	validator.Init()

	return &App{
		AuthService:   authService,
		MemberService: memberService,

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
