package firebase

import (
	"context"
	"path/filepath"
	"sync"
	"time"
	"washbonus/internal/app"
	"washbonus/internal/config"
	"washbonus/internal/infrastructure/rabbit"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

const authTimeout = time.Second * 15

type UserCache struct {
	sync.RWMutex
	Cache map[string]*app.AdminAuth
}

type FirebaseService struct {
	app  *firebase.App
	auth *auth.Client

	userSvc   app.UserService
	adminSvc  app.AdminService
	rabbitSvc rabbit.RabbitService

	adminCache *UserCache
}

func New(cfg config.FirebaseConfig, userSvc app.UserService, adminSvc app.AdminService, rabbitSvc rabbit.RabbitService) (*FirebaseService, error) {
	keyFilePath, err := filepath.Abs(cfg.FirebaseKeyFilePath)
	if err != nil {
		panic("Unable to load service key")
	}
	opt := option.WithCredentialsFile(keyFilePath)

	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Failed to load Firebase")
	}

	auth, err := fbApp.Auth(context.Background())
	if err != nil {
		panic("Failed to load Firebase auth")
	}

	return &FirebaseService{
		app:  fbApp,
		auth: auth,

		userSvc:   userSvc,
		adminSvc:  adminSvc,
		rabbitSvc: rabbitSvc,

		adminCache: &UserCache{Cache: make(map[string]*app.AdminAuth)},
	}, nil
}
