package main

import (
	"log"
	"washBonus/internal/app"
	"washBonus/internal/config"
	"washBonus/internal/controllers"
	sessionsRepo "washBonus/internal/dal/sessions"
	userRepo "washBonus/internal/dal/user"
	washRepo "washBonus/internal/dal/washServer"
	"washBonus/internal/infrastructure/firebase"
	rabbitMQ "washBonus/internal/infrastructure/rabbit"
	"washBonus/internal/services/rabbit"
	"washBonus/internal/services/schedule"
	"washBonus/internal/services/session"
	"washBonus/internal/services/user"
	"washBonus/internal/services/washServer"
	"washBonus/internal/transport/rest"

	"washBonus/pkg/bootstrap"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.NewConfig()
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

	err = bootstrap.UpMigrations(dbConn.DB, cfg.DB.Database, "internal/migrations")
	if err != nil {
		l.Fatalln("up migrations: ", err)
	}

	l.Debug("applied migrations")

	authSvc := firebase.New(cfg.FirebaseConfig.FirebaseKeyFilePath)

	userRepo := userRepo.NewRepo(l, dbConn)
	washRepo := washRepo.NewRepo(l, dbConn)
	sessionRepo := sessionsRepo.NewRepo(l, dbConn)

	userSvc := user.New(l, userRepo)
	washServerSvc := washServer.New(l, washRepo)
	sessionSvc := session.New(l, userRepo, sessionRepo, washRepo, cfg.SessionsConfig.ReportsProcessingDelayInMinutes, cfg.SessionsConfig.MoneyReportRewardPercentDefault)

	rabbitSvc := rabbit.New(l, sessionSvc, userSvc, washServerSvc)

	rabbitMQ, err := rabbitMQ.New(l, cfg.RabbitMQConfig.Url, cfg.RabbitMQConfig.Port, cfg.RabbitMQConfig.User, cfg.RabbitMQConfig.Password, rabbitSvc)
	if err != nil {
		l.Fatalln("new rabbit conn: ", err)
	}
	l.Debug("connected to rabbit")

	sessionCtrl := controllers.NewSessionController(l, sessionSvc, userSvc, washServerSvc, rabbitMQ)
	userCtrl := controllers.NewUserController(l, userSvc, sessionSvc)
	washServerCtrl := controllers.NewWashServerController(l, washServerSvc, userSvc, rabbitMQ)

	schedulerSvc := schedule.New(l, sessionSvc)
	runScheduler(schedulerSvc, cfg.SchedulerConfig)

	errc := make(chan error)

	go runHTTPServer(errc, l, cfg, authSvc, sessionCtrl, userCtrl, washServerCtrl)

	err = <-errc
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg *config.Config, authSvc firebase.Service, sessionCtrl app.SessionController, userCtrl app.UserController, washServerCtrl app.WashServerController) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("panic: ", r)
		}
	}()
	server, err := rest.NewServer(cfg, authSvc, l, sessionCtrl, userCtrl, washServerCtrl)
	if err != nil {
		l.Fatalln("init rest server:", err)
	}

	errc <- server.Serve()
}

func runScheduler(schedulerSvc app.ScheduleService, schedulerCfg config.SchedulerConfig) {
	reportsDelay := schedulerCfg.ReportsDelayMinutes
	sessionsDelay := schedulerCfg.SessionsDelayMinutes

	sessionRetentionDays := schedulerCfg.SessionRetentionDays
	schedulerSvc.Run(reportsDelay, sessionsDelay, sessionRetentionDays)
}
