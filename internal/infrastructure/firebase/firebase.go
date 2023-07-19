package firebase

import (
	"context"
	"path/filepath"
	"time"
	"washBonus/internal/app"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type UID string

const authTimeout = time.Second

type Service interface {
	Auth(token string) (*app.Auth, error)
}

type FirebaseService struct {
	app  *firebase.App
	auth *auth.Client
}

func New(keyfileLocation string) Service {
	keyFilePath, err := filepath.Abs(keyfileLocation)
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
	}
}
