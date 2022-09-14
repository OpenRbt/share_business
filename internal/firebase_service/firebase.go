package firebase_service

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/powerman/structlog"
	"google.golang.org/api/option"
	"path/filepath"
	"time"

	"net/http"
)

type UID string

const authTimeout = time.Second

var log = structlog.New()

type FirebaseService interface {
	AuthMiddleware(handler http.Handler) http.Handler
	GetFirebaseProfile(token string) (interface{}, error)
}

type Service struct {
	app  *firebase.App
	auth *auth.Client
}

func New(keyfileLocation string) FirebaseService {
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
		panic("Failed to load Firebase")
	}

	return &Service{
		app:  app,
		auth: auth,
	}
}
