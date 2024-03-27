package rest

import (
	"net/http"
	"path"
	"strconv"
	"washbonus/internal/app"
	"washbonus/internal/config"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"github.com/rs/cors"
	"github.com/sebest/xff"
	"go.uber.org/zap"

	"washbonus/openapi/admin/restapi"
	admin "washbonus/openapi/admin/restapi"
	adminOp "washbonus/openapi/admin/restapi/operations"

	bonus "washbonus/openapi/bonus/restapi"
	bonusOp "washbonus/openapi/bonus/restapi/operations"
	standard "washbonus/openapi/bonus/restapi/operations/standard"
)

type firebaseService interface {
	BonusAuth(token string) (*app.Auth, error)
	AdminAuth(token string) (*app.AdminAuth, error)
}

type service struct {
	l        *zap.SugaredLogger
	firebase firebaseService

	sessionCtrl     app.SessionController
	userCtrl        app.UserController
	adminCtrl       app.AdminController
	washServerCtrl  app.WashServerController
	orgCtrl         app.OrganizationController
	serverGroupCtrl app.ServerGroupController
	walletCtrl      app.WalletController
	bonusReportCtrl app.BonusReportController
}

func NewServer(l *zap.SugaredLogger, cfg *config.Config, firebase app.FirebaseService, ctrls app.Controllers) (*bonus.Server, error) {
	svc := &service{
		l:        l,
		firebase: firebase,

		sessionCtrl:     ctrls.Session,
		adminCtrl:       ctrls.Admin,
		userCtrl:        ctrls.User,
		washServerCtrl:  ctrls.Wash,
		orgCtrl:         ctrls.Org,
		serverGroupCtrl: ctrls.Group,
		walletCtrl:      ctrls.Wallet,
		bonusReportCtrl: ctrls.BonusReport,
	}

	// Load Bonus Swagger Definition
	swaggerSpecBonus, err := loads.Embedded(bonus.SwaggerJSON, bonus.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load embedded swagger spec for bonus")
	}

	// Load Admin Swagger Definition
	swaggerSpecAdmin, err := loads.Embedded(admin.SwaggerJSON, admin.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load embedded swagger spec for admin")
	}

	// Create Bonus API
	bonusAPI := bonusOp.NewWashBonusAPI(swaggerSpecBonus)
	bonusAPI.Logger = structlog.New(structlog.KeyUnit, "swagger-bonus").Printf
	bonusAPI.AuthKeyAuth = svc.firebase.BonusAuth

	// Create Admin API
	adminAPI := adminOp.NewWashAdminAPI(swaggerSpecAdmin)
	adminAPI.Logger = structlog.New(structlog.KeyUnit, "swagger-admin").Printf
	adminAPI.AuthKeyAuth = svc.firebase.AdminAuth

	// Initialize Handlers
	svc.initAdminUserHandlers(adminAPI)
	svc.initOrganizationsHandlers(adminAPI)
	svc.initServerGroupHandlers(adminAPI)
	svc.initBonusReportHandlers(adminAPI)
	svc.initSessionHandlers(bonusAPI)
	svc.initWalletHandlers(bonusAPI)
	svc.initWashServerHandlers(adminAPI)

	bonusAPI.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(healthCheck)

	server := bonus.NewServer(nil) // No specific API attached directly
	server.Host = string(cfg.Host)
	port, err := strconv.ParseInt(cfg.HTTPPort, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse port")
	}
	server.Port = int(port)

	// The middleware executes before anything.
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(cfg.BasePath, svc.l)
		accesslog := makeAccessLog(cfg.BasePath, svc.l)
		redocOpts := middleware.RedocOpts{
			BasePath: cfg.BasePath,
			SpecURL:  path.Join(cfg.BasePath, "/swagger.json"),
		}
		return xffmw.Handler(logger(noCache(recovery(accesslog(
			middleware.Spec(cfg.BasePath, restapi.FlatSwaggerJSON,
				middleware.Redoc(redocOpts,
					handler))), svc.l))))
	}
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	//middlewares := func(handler http.Handler) http.Handler {
	//	safePath := map[string]bool{}
	//	isSafe := func(r *http.Request) bool { return safePath[r.URL.Path] }
	//	//forbidCSRF := makeForbidCSRF(isSafe)
	//
	//	return forbidCSRF(handler)
	//}

	newCORS := cors.New(cors.Options{
		AllowedOrigins:   splitCommaSeparatedStr(cfg.AllowedOrigins),
		AllowedMethods:   []string{"POST", "PUT", "PATCH", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	newCORS.Log = cors.Logger(structlog.New(structlog.KeyUnit, "CORS"))
	handleCORS := newCORS.Handler

	// Setup the combined mux
	mux := http.NewServeMux()
	mux.Handle("/admin/", adminAPI.Serve(globalMiddlewares))
	mux.Handle("/", bonusAPI.Serve(globalMiddlewares))

	server.SetHandler(handleCORS(mux))

	return server, nil
}

func healthCheck(params standard.HealthCheckParams, profile *app.Auth) standard.HealthCheckResponder {
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}
