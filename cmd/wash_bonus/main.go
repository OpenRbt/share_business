package main

import (
	"log"
	"wash_bonus/internal/dal/sessions"
	"wash_bonus/internal/infrastructure/rabbit"

	"go.uber.org/zap"

	session_svc "wash_bonus/internal/app/session"
	user_svc "wash_bonus/internal/app/user"
	wash_server_svc "wash_bonus/internal/app/wash_server"

	user_repo "wash_bonus/internal/dal/user"
	wash_server_repo "wash_bonus/internal/dal/wash_server"
	"wash_bonus/internal/infrastructure/firebase"
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

	authSvc := firebase.New(cfg.FirebaseConfig.FirebaseKeyFilePath)

	userRepo := user_repo.NewRepo(l, dbConn)
	washRepo := wash_server_repo.NewRepo(l, dbConn)
	sessionRepo := sessions.NewRepo(l, dbConn)

	userSvc := user_svc.New(l, userRepo)
	washServerSvc := wash_server_svc.New(l, washRepo)
	sessionSvc := session_svc.New(l, washRepo, userRepo, sessionRepo)

	rabbitSvc, err := rabbit.New(
		l,
		cfg.RabbitMQConfig.Url,
		cfg.RabbitMQConfig.Port,
		cfg.RabbitMQConfig.CertsPath,
		cfg.RabbitMQConfig.User,
		cfg.RabbitMQConfig.Password,
		washServerSvc,
		sessionSvc,
	)
	if err != nil {
		l.Fatalln("new rabbit conn: ", err)
	}
	l.Debug("connected to rabbit")

	sessionSvc.AssignRabbit(rabbitSvc.SendMessage)

	errc := make(chan error)

	go runHTTPServer(errc, l, cfg, authSvc, userSvc, sessionSvc)

	err = <-errc
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg *bootstrap.Config, authSvc firebase.Service, userSvc user_svc.Service, sessionSvc session_svc.Service) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("panic: ", r)
		}
	}()
	server, err := rest.NewServer(cfg, authSvc, l, userSvc, sessionSvc)
	if err != nil {
		l.Fatalln("init rest server:", err)
	}

	errc <- server.Serve()
}
