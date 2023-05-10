package main

import (
	"log"
	"wash_bonus/internal/app/scheduler"
	sessionSvc "wash_bonus/internal/app/session"
	userSvc "wash_bonus/internal/app/user"
	washserverSvc "wash_bonus/internal/app/wash_server"
	sessionsRepo "wash_bonus/internal/dal/sessions"
	userRepo "wash_bonus/internal/dal/user"
	washRepo "wash_bonus/internal/dal/wash_server"
	"wash_bonus/internal/infrastructure/firebase"
	"wash_bonus/internal/infrastructure/rabbit"
	"wash_bonus/internal/transport/rest"
	rabbitUseCase "wash_bonus/internal/usecase/rabbit"
	sessionUseCase "wash_bonus/internal/usecase/session"
	userUseCase "wash_bonus/internal/usecase/user"
	"wash_bonus/pkg/bootstrap"

	"go.uber.org/zap"
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

	userRepo := userRepo.NewRepo(l, dbConn)
	washRepo := washRepo.NewRepo(l, dbConn)
	sessionRepo := sessionsRepo.NewRepo(l, dbConn)

	userSvc := userSvc.New(l, userRepo)
	washServerSvc := washserverSvc.New(l, washRepo)
	sessionSvc := sessionSvc.New(l, userRepo, sessionRepo)

	rabbitUseCase := rabbitUseCase.New(l, sessionSvc, userSvc, washServerSvc)

	rabbitSvc, err := rabbit.New(l, cfg.RabbitMQConfig.Url, cfg.RabbitMQConfig.Port, cfg.RabbitMQConfig.User, cfg.RabbitMQConfig.Password, rabbitUseCase)
	if err != nil {
		l.Fatalln("new rabbit conn: ", err)
	}
	l.Debug("connected to rabbit")

	sessionUseCase := sessionUseCase.New(l, sessionSvc, userSvc, washServerSvc, rabbitSvc)
	userUseCase := userUseCase.New(l, userSvc)

	schedulerSvc := scheduler.New(l, sessionSvc)
	schedulerSvc.Run(cfg.SchedulerConfig.DelayMinutes)

	errc := make(chan error)

	go runHTTPServer(errc, l, cfg, authSvc, sessionUseCase, userUseCase)

	err = <-errc
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg *bootstrap.Config, authSvc firebase.Service, sessionUseCase sessionUseCase.UseCase, userUseCase userUseCase.UseCase) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("panic: ", r)
		}
	}()
	server, err := rest.NewServer(cfg, authSvc, l, sessionUseCase, userUseCase)
	if err != nil {
		l.Fatalln("init rest server:", err)
	}

	errc <- server.Serve()
}
