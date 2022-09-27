package firebase_auth

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/powerman/structlog"
	"google.golang.org/api/option"
	"path/filepath"
	"time"
)

type UID string
type FirebaseProfile auth.UserRecord

const authTimeout = time.Second

var log = structlog.New()

type Service interface {
	VerifyClaims(user FirebaseProfile, requiredClaims ...string) error
	Auth(token string) (interface{}, error)
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
