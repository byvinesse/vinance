package application

import (
	"context"

	"github.com/vincentkdeli/vinance-backend/config"
	"github.com/vincentkdeli/vinance-backend/entity"
	auth "github.com/vincentkdeli/vinance-backend/pkg/authenticator"
	"github.com/vincentkdeli/vinance-backend/pkg/service"
	"github.com/vincentkdeli/vinance-backend/pkg/validator"
	"github.com/vincentkdeli/vinance-backend/repository/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	AuthService   service.IAuthService
	MemberService service.IMemberService

	Authenticator *auth.Authenticator
}

func NewApp() *App {
	ctx := context.Background()

	authConfig := config.LoadAuth()
	authenticator := auth.NewAuthenticator(authConfig)

	// Init Database
	database := connectDatabase(ctx)

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

func connectDatabase(ctx context.Context) *mongo.Database {
	databaseConfig := config.LoadDatabase()
	clientOptions := options.Client()
	clientOptions.ApplyURI(databaseConfig.URI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return client.Database(entity.DatabaseName)
}
