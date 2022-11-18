// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"wash_bonus/internal/app"
	"wash_bonus/openapi/restapi/operations"
	"wash_bonus/openapi/restapi/operations/bonus"
	"wash_bonus/openapi/restapi/operations/standard"
	"wash_bonus/openapi/restapi/operations/user"
)

//go:generate swagger generate server --target ../../openapi --name WashBonus --spec ../swagger.yaml --principal wash_bonus/internal/app.Auth --exclude-main --strict-responders

func configureFlags(api *operations.WashBonusAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WashBonusAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.AuthKeyAuth == nil {
		api.AuthKeyAuth = func(token string) (*app.Auth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.BonusCancelHandler == nil {
		api.BonusCancelHandler = bonus.CancelHandlerFunc(func(params bonus.CancelParams, principal *app.Auth) bonus.CancelResponder {
			return bonus.CancelNotImplemented()
		})
	}
	if api.BonusConfirmHandler == nil {
		api.BonusConfirmHandler = bonus.ConfirmHandlerFunc(func(params bonus.ConfirmParams, principal *app.Auth) bonus.ConfirmResponder {
			return bonus.ConfirmNotImplemented()
		})
	}
	if api.UserGetHandler == nil {
		api.UserGetHandler = user.GetHandlerFunc(func(params user.GetParams, principal *app.Auth) user.GetResponder {
			return user.GetNotImplemented()
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *app.Auth) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		})
	}
	if api.BonusUseHandler == nil {
		api.BonusUseHandler = bonus.UseHandlerFunc(func(params bonus.UseParams, principal *app.Auth) bonus.UseResponder {
			return bonus.UseNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
