package firebase

import (
	"context"
	"path/filepath"
	"time"
	"washbonus/internal/app"
	"washbonus/internal/config"
	"washbonus/internal/infrastructure/rabbit"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

const authTimeout = time.Second * 15

type FirebaseService struct {
	app  *firebase.App
	auth *auth.Client

	userSvc   app.UserService
	adminSvc  app.AdminService
	rabbitSvc rabbit.RabbitService
}

func New(cfg config.FirebaseConfig, userSvc app.UserService, adminSvc app.AdminService, rabbitSvc rabbit.RabbitService) (*FirebaseService, error) {
	keyFilePath, err := filepath.Abs(cfg.FirebaseKeyFilePath)
	if err != nil {
		panic("Unable to load service key")
	}
	opt := option.WithCredentialsFile(keyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Failed to load Firebase")
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Failed to load Firebase auth")
	}

	return &FirebaseService{
		app:  app,
		auth: auth,

		userSvc:   userSvc,
		adminSvc:  adminSvc,
		rabbitSvc: rabbitSvc,
	}, nil
}
