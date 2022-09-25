// Code generated by mtgroup-generator.
package api

import (
	"context"
	"net/http"
	"path"

	"wash-bonus/internal/api/restapi/restapi"
	"wash-bonus/internal/api/restapi/restapi/operations"
	"wash-bonus/internal/api/restapi/restapi/operations/standard"

	user "wash-bonus/internal/api/restapi/restapi/operations/user"
	washServer "wash-bonus/internal/api/restapi/restapi/operations/wash_server"

	"wash-bonus/internal/api/restapi/models"
	"wash-bonus/internal/app"
	"wash-bonus/internal/def"

	extauthapi "github.com/mtgroupit/mt-mock-extauthapi"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"github.com/rs/cors"
	"github.com/sebest/xff"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Log is a synonym for convenience.
type Log = *structlog.Logger

type Config struct {
	Host           string
	Port           int
	BasePath       string
	AllowedOrigins string
}

type service struct {
	app     app.App
	extAuth AuthSvc
}

func NewServer(appl app.App, extAuth AuthSvc, cfg Config) (*restapi.Server, error) {
	svc := &service{
		app:     appl,
		extAuth: extAuth,
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
	api.AuthKeyAuth = svc.checkerAuth

	api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(healthCheck)
	api.StandardAddTestDataHandler = standard.AddTestDataHandlerFunc(svc.addTestData)
	api.UserGetUserHandler = user.GetUserHandlerFunc(svc.GetUser)
	api.UserAddUserHandler = user.AddUserHandlerFunc(svc.AddUser)
	api.UserEditUserHandler = user.EditUserHandlerFunc(svc.EditUser)
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(svc.DeleteUser)
	api.UserListUserHandler = user.ListUserHandlerFunc(svc.ListUser)
	api.WashServerGetWashServerHandler = washServer.GetWashServerHandlerFunc(svc.GetWashServer)
	api.WashServerAddWashServerHandler = washServer.AddWashServerHandlerFunc(svc.AddWashServer)
	api.WashServerEditWashServerHandler = washServer.EditWashServerHandlerFunc(svc.EditWashServer)
	api.WashServerDeleteWashServerHandler = washServer.DeleteWashServerHandlerFunc(svc.DeleteWashServer)
	api.WashServerListWashServerHandler = washServer.ListWashServerHandlerFunc(svc.ListWashServer)

	server := restapi.NewServer(api)
	server.Host = string(cfg.Host)
	server.Port = int(cfg.Port)

	// The middleware executes before anything.
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(cfg.BasePath)
		accesslog := makeAccessLog(cfg.BasePath)
		redocOpts := middleware.RedocOpts{
			BasePath: cfg.BasePath,
			SpecURL:  path.Join(cfg.BasePath, "/swagger.json"),
		}
		return xffmw.Handler(logger(noCache(recovery(accesslog(
			middleware.Spec(cfg.BasePath, restapi.FlatSwaggerJSON,
				middleware.Redoc(redocOpts,
					handler)))))))
	}
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	middlewares := func(handler http.Handler) http.Handler {
		safePath := map[string]bool{}
		isSafe := func(r *http.Request) bool { return safePath[r.URL.Path] }
		forbidCSRF := makeForbidCSRF(isSafe)

		withValidatePath := map[string]bool{}
		needValidate := func(r *http.Request) bool { return withValidatePath[r.URL.Path] }
		validateToken := svc.makeValidateToken(needValidate)

		return validateToken(forbidCSRF(handler))
	}

	newCORS := cors.New(cors.Options{
		AllowedOrigins:   splitCommaSeparatedStr(cfg.AllowedOrigins),
		AllowedMethods:   []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
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

func healthCheck(params standard.HealthCheckParams, profile interface{}) middleware.Responder {
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}
func (svc *service) addTestData(params standard.AddTestDataParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	err := svc.app.AddTestData(toAppProfile(prof))
	switch {
	default:
		log.PrintErr("AddTestData server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return standard.NewAddTestDataDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("AddTestData client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return standard.NewAddTestDataDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("AddTestData ok")
		return standard.NewAddTestDataOK()
	}
}
