package rest

import (
	"net/http"
	"path"
	"strconv"
	"washBonus/internal/app"
	"washBonus/internal/config"
	"washBonus/internal/infrastructure/firebase"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"github.com/rs/cors"
	"github.com/sebest/xff"
	"go.uber.org/zap"

	"washBonus/openapi/restapi"
	"washBonus/openapi/restapi/operations"
	"washBonus/openapi/restapi/operations/standard"
)

type service struct {
	l    *zap.SugaredLogger
	auth firebase.Service

	sessionCtrl     app.SessionController
	userCtrl        app.UserController
	washServerCtrl  app.WashServerController
	orgCtrl         app.OrganizationController
	serverGroupCtrl app.ServerGroupController
	walletCtrl      app.WalletController
}

func NewServer(cfg *config.Config, auth firebase.Service, l *zap.SugaredLogger, sessionCtrl app.SessionController, userCtrl app.UserController, washServerCtrl app.WashServerController, orgCtrl app.OrganizationController, serverGroupCtrl app.ServerGroupController, walletCtrl app.WalletController) (*restapi.Server, error) {
	svc := &service{
		l:    l,
		auth: auth,

		sessionCtrl:     sessionCtrl,
		userCtrl:        userCtrl,
		washServerCtrl:  washServerCtrl,
		orgCtrl:         orgCtrl,
		serverGroupCtrl: serverGroupCtrl,
		walletCtrl:      walletCtrl,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load embedded swagger spec")
	}
	if cfg.BasePath == "" {
		cfg.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.BasePath

	api := operations.NewWashBonusAPI(swaggerSpec)

	api.Logger = structlog.New(structlog.KeyUnit, "swagger").Printf
	api.AuthKeyAuth = svc.auth.Auth

	api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(healthCheck)

	svc.initUserHandlers(api)
	svc.initSessionHandlers(api)
	svc.initWashServerHandlers(api)
	svc.initOrganizationsHandlers(api)
	svc.initServerGroupHandlers(api)
	svc.initWalletHandlers(api)

	server := restapi.NewServer(api)
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
	middlewares := func(handler http.Handler) http.Handler {
		return handler
	}

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

	server.SetHandler(handleCORS(globalMiddlewares(api.Serve(middlewares))))

	return server, nil
}

func healthCheck(params standard.HealthCheckParams, profile *app.Auth) standard.HealthCheckResponder {
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}
