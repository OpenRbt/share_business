// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"washbonus/internal/app"
	"washbonus/openapi/bonus/restapi/operations"
	"washbonus/openapi/bonus/restapi/operations/sessions"
	"washbonus/openapi/bonus/restapi/operations/standard"
	"washbonus/openapi/bonus/restapi/operations/wallets"
)

//go:generate swagger generate server --target ../../bonus --name WashBonus --spec ../../bonus.swagger.yaml --principal washbonus/internal/app.Auth --exclude-main --strict-responders

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

	if api.SessionsAssignUserToSessionHandler == nil {
		api.SessionsAssignUserToSessionHandler = sessions.AssignUserToSessionHandlerFunc(func(params sessions.AssignUserToSessionParams, principal *app.Auth) sessions.AssignUserToSessionResponder {
			return sessions.AssignUserToSessionNotImplemented()
		})
	}
	if api.SessionsChargeBonusesOnSessionHandler == nil {
		api.SessionsChargeBonusesOnSessionHandler = sessions.ChargeBonusesOnSessionHandlerFunc(func(params sessions.ChargeBonusesOnSessionParams, principal *app.Auth) sessions.ChargeBonusesOnSessionResponder {
			return sessions.ChargeBonusesOnSessionNotImplemented()
		})
	}
	if api.SessionsGetSessionByIDHandler == nil {
		api.SessionsGetSessionByIDHandler = sessions.GetSessionByIDHandlerFunc(func(params sessions.GetSessionByIDParams, principal *app.Auth) sessions.GetSessionByIDResponder {
			return sessions.GetSessionByIDNotImplemented()
		})
	}
	if api.WalletsGetWalletByOrganizationIDHandler == nil {
		api.WalletsGetWalletByOrganizationIDHandler = wallets.GetWalletByOrganizationIDHandlerFunc(func(params wallets.GetWalletByOrganizationIDParams, principal *app.Auth) wallets.GetWalletByOrganizationIDResponder {
			return wallets.GetWalletByOrganizationIDNotImplemented()
		})
	}
	if api.WalletsGetWalletsHandler == nil {
		api.WalletsGetWalletsHandler = wallets.GetWalletsHandlerFunc(func(params wallets.GetWalletsParams, principal *app.Auth) wallets.GetWalletsResponder {
			return wallets.GetWalletsNotImplemented()
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *app.Auth) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
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
