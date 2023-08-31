package main

import (
	"log"
	"washBonus/internal/app"
	"washBonus/internal/config"
	"washBonus/internal/controllers"
	organizationRepo "washBonus/internal/dal/organizations"
	serverGroupRepo "washBonus/internal/dal/serverGroups"
	sessionsRepo "washBonus/internal/dal/sessions"
	userRepo "washBonus/internal/dal/user"
	walletRepo "washBonus/internal/dal/wallets"
	washRepo "washBonus/internal/dal/washServer"
	"washBonus/internal/infrastructure/firebase"
	rabbitMQ "washBonus/internal/infrastructure/rabbit"
	"washBonus/internal/services/organizations"
	"washBonus/internal/services/rabbit"
	"washBonus/internal/services/schedule"
	"washBonus/internal/services/serverGroups"
	"washBonus/internal/services/session"
	"washBonus/internal/services/user"
	"washBonus/internal/services/wallets"
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

	userRepo := userRepo.NewRepo(l, dbConn)
	washRepo := washRepo.NewRepo(l, dbConn)
	sessionRepo := sessionsRepo.NewRepo(l, dbConn)
	orgRepo := organizationRepo.NewRepo(l, dbConn)
	groupRepo := serverGroupRepo.NewRepo(l, dbConn)
	walletRepo := walletRepo.NewRepo(l, dbConn)

	userSvc := user.New(l, userRepo, orgRepo)
	authSvc := firebase.New(cfg.FirebaseConfig.FirebaseKeyFilePath, userSvc)
	washServerSvc := washServer.New(l, washRepo, groupRepo)
	sessionSvc := session.New(l, userRepo, sessionRepo, washRepo, walletRepo, cfg.SessionsConfig.ReportsProcessingDelayInMinutes, cfg.SessionsConfig.MoneyReportRewardPercentDefault)
	orgSvc := organizations.New(l, orgRepo, userRepo)
	groupSvc := serverGroups.New(l, groupRepo, orgRepo)
	walletSvc := wallets.New(l, walletRepo, orgRepo)

	rabbitSvc := rabbit.New(l, sessionSvc, userSvc, washServerSvc, walletSvc)

	rabbitMQ, err := rabbitMQ.New(l, cfg.RabbitMQConfig.Url, cfg.RabbitMQConfig.Port, cfg.RabbitMQConfig.User, cfg.RabbitMQConfig.Password, rabbitSvc)
	if err != nil {
		l.Fatalln("new rabbit conn: ", err)
	}
	l.Debug("connected to rabbit")

	sessionCtrl := controllers.NewSessionController(l, sessionSvc, userSvc, washServerSvc, rabbitMQ)
	userCtrl := controllers.NewUserController(l, userSvc, sessionSvc)
	washServerCtrl := controllers.NewWashServerController(l, washServerSvc, userSvc, groupSvc, orgSvc, rabbitMQ)
	orgCtrl := controllers.NewOrganizationController(l, orgSvc)
	groupCtrl := controllers.NewServerGroupController(l, groupSvc, orgSvc)
	walletCtrl := controllers.NewWalletController(l, walletSvc, sessionSvc)

	schedulerSvc := schedule.New(l, sessionSvc)
	runScheduler(schedulerSvc, cfg.SchedulerConfig)

	errc := make(chan error)

	go runHTTPServer(errc, l, cfg, authSvc, sessionCtrl, userCtrl, washServerCtrl, orgCtrl, groupCtrl, walletCtrl)

	err = <-errc
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg *config.Config, authSvc firebase.Service, sessionCtrl app.SessionController, userCtrl app.UserController, washServerCtrl app.WashServerController, orgCtrl app.OrganizationController, groupCtrl app.ServerGroupController, walletCtrl app.WalletController) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("panic: ", r)
		}
	}()
	server, err := rest.NewServer(cfg, authSvc, l, sessionCtrl, userCtrl, washServerCtrl, orgCtrl, groupCtrl, walletCtrl)
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
