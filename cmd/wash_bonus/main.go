package main

import (
	"log"
	"wash_bonus/internal/app"
	"wash_bonus/internal/app/balance"
	"wash_bonus/internal/dal"
	"wash_bonus/internal/firebase_authorization"
	"wash_bonus/internal/transport/rest"
	"wash_bonus/pkg/bootstrap"
)

func main() {
	cfg, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln("new config: ", err)
	}

	l, err := bootstrap.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("new logger: ", err)
	}

	dbConn, err := bootstrap.NewDbConn(cfg.DB)
	if err != nil {
		l.Fatalln("new db conn: ", err)
	}
	defer dbConn.Close()

	l.Debug("connected to db")

	err = bootstrap.UpMigrations(dbConn.DB, cfg.DB.Database, "migrations")
	if err != nil {
		l.Fatalln("up migrations: ", err)
	}

	l.Debug("applied migrations")

	authSvc := firebase_authorization.New(cfg.FirebaseConfig.FirebaseKeyFilePath)

	repo := dal.New(dbConn, l)

	userSvc := app.NewUserService(l, repo)
	balanceSvc := balance.New(l, repo)

	server, err := rest.NewServer(cfg, authSvc, l, userSvc, balanceSvc)
	if err != nil {
		l.Fatalln("init rest server:", err)
	}

	err = server.Serve()
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}
