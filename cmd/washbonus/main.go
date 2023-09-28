package main

import (
	"log"
	"washbonus/internal/app"
	"washbonus/internal/config"
	ctrls "washbonus/internal/controllers"
	adminRepo "washbonus/internal/dal/adminusers"
	orgRepo "washbonus/internal/dal/organizations"
	groupRepo "washbonus/internal/dal/servergroups"
	sessionRepo "washbonus/internal/dal/sessions"
	userRepo "washbonus/internal/dal/users"
	walletRepo "washbonus/internal/dal/wallets"
	washRepo "washbonus/internal/dal/washservers"
	"washbonus/internal/infrastructure/firebase"
	rabbitMQ "washbonus/internal/infrastructure/rabbit"
	"washbonus/internal/services/adminuser"
	"washbonus/internal/services/organization"
	"washbonus/internal/services/rabbit"
	"washbonus/internal/services/schedule"
	"washbonus/internal/services/servergroup"
	"washbonus/internal/services/session"
	"washbonus/internal/services/user"
	"washbonus/internal/services/wallet"
	"washbonus/internal/services/washserver"
	"washbonus/internal/transport/rest"

	"washbonus/pkg/bootstrap"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

func main() {
	cfg := loadConfig()
	l := setupLogger(cfg)
	dbConn := setupDatabase(cfg, l)
	defer dbConn.Close()

	repos := setupRepositories(l, dbConn)
	services := setupServices(l, repos)

	rabbitMQ := setupRabbit(l, cfg, services)
	firebase := setupFirebase(l, cfg, services)
	scheduler := schedule.New(l, services.Session)
	runSchedulerTasks(scheduler, cfg.SchedulerConfig)

	controllers := setupControllers(l, services, rabbitMQ)

	serveRestApi(l, cfg, firebase, controllers)
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg *config.Config, firebase app.FirebaseService, ctrls app.Controllers) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("failed to recover: ", r)
		}
	}()

	server, err := rest.NewServer(l, cfg, firebase, ctrls)
	if err != nil {
		l.Fatalln("failed to init rest server:", err)
	}

	errc <- server.Serve()
}

func runSchedulerTasks(schedulerSvc app.ScheduleService, schedulerCfg config.SchedulerConfig) {
	reportsDelay := schedulerCfg.ReportsDelayMinutes
	sessionsDelay := schedulerCfg.SessionsDelayMinutes

	sessionRetentionDays := schedulerCfg.SessionRetentionDays
	schedulerSvc.Run(reportsDelay, sessionsDelay, sessionRetentionDays)
}

func loadConfig() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("new config: ", err)
	}
	return cfg
}

func setupLogger(cfg *config.Config) *zap.SugaredLogger {
	l, err := bootstrap.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("new logger: ", err)
	}
	return l
}

func setupDatabase(cfg *config.Config, l *zap.SugaredLogger) *dbr.Connection {
	dbConn, err := bootstrap.NewDbConn(cfg.DB)
	if err != nil {
		l.Fatalln("new db conn: ", err)
	}
	l.Debug("connected to db")

	err = bootstrap.UpMigrations(dbConn.DB, cfg.DB.Database, "internal/migrations")
	if err != nil {
		l.Fatalln("failed to perform migrations: ", err)
	}

	l.Debug("applied migrations")

	return dbConn
}

func setupRepositories(l *zap.SugaredLogger, dbConn *dbr.Connection) app.Repositories {
	return app.Repositories{
		Admin:   adminRepo.NewRepo(l, dbConn),
		Org:     orgRepo.NewRepo(l, dbConn),
		Group:   groupRepo.NewRepo(l, dbConn),
		Wash:    washRepo.NewRepo(l, dbConn),
		User:    userRepo.NewRepo(l, dbConn),
		Session: sessionRepo.NewRepo(l, dbConn),
		Wallet:  walletRepo.NewRepo(l, dbConn),
	}
}

func setupServices(l *zap.SugaredLogger, repos app.Repositories) app.Services {
	return app.Services{
		Admin:   adminuser.New(l, repos.Admin, repos.Org),
		Org:     organization.New(l, repos.Org, repos.Admin),
		Group:   servergroup.New(l, repos.Group, repos.Org),
		Wash:    washserver.New(l, repos.Wash, repos.Group),
		User:    user.New(l, repos.User, repos.Org),
		Session: session.New(l, repos.User, repos.Session, repos.Wash, repos.Wallet),
		Wallet:  wallet.New(l, repos.Wallet, repos.Org),
	}
}

func setupRabbit(l *zap.SugaredLogger, cfg *config.Config, services app.Services) *rabbitMQ.Service {
	rabbitSvc := rabbit.New(l, services.Session, services.User, services.Wash, services.Wallet)

	rabbitMQ, err := rabbitMQ.New(l, cfg.RabbitMQConfig, rabbitSvc)
	if err != nil {
		l.Fatalln("failed to init rabbit: ", err)
	}
	l.Debug("connected to rabbit")

	return rabbitMQ
}

func setupFirebase(l *zap.SugaredLogger, cfg *config.Config, services app.Services) *firebase.FirebaseService {
	firebase, err := firebase.New(cfg.FirebaseConfig, services.User, services.Admin)
	if err != nil {
		l.Fatalln("failed to init firebase: ", err)
	}

	return firebase
}

func setupControllers(l *zap.SugaredLogger, services app.Services, rabbit rabbitMQ.RabbitService) app.Controllers {
	return app.Controllers{
		Admin:   ctrls.NewAdminUserController(l, services.Admin),
		Org:     ctrls.NewOrganizationController(l, services.Org),
		Group:   ctrls.NewServerGroupController(l, services.Group, services.Org),
		Wash:    ctrls.NewWashServerController(l, services.Wash, services.Group, services.Org, rabbit),
		User:    ctrls.NewUserController(l, services.User, services.Session),
		Session: ctrls.NewSessionController(l, services.Session, services.User, services.Wash, rabbit),
		Wallet:  ctrls.NewWalletController(l, services.Wallet, services.Session),
	}
}

func serveRestApi(l *zap.SugaredLogger, cfg *config.Config, firebase app.FirebaseService, controllers app.Controllers) {
	errc := make(chan error)
	go runHTTPServer(errc, l, cfg, firebase, controllers)

	err := <-errc
	if err != nil {
		l.Fatalln("Failed to serve REST API:", err)
	}
	close(errc)
}
