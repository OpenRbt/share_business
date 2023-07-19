// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"washBonus/internal/app"
	"washBonus/openapi/restapi/operations"
	"washBonus/openapi/restapi/operations/sessions"
	"washBonus/openapi/restapi/operations/standard"
	"washBonus/openapi/restapi/operations/users"
	"washBonus/openapi/restapi/operations/wash_servers"
)

//go:generate swagger generate server --target ../../openapi --name WashBonus --spec ../swagger.yaml --principal washBonus/internal/app.Auth --exclude-main --strict-responders

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
	if api.WashServersCreateWashServerHandler == nil {
		api.WashServersCreateWashServerHandler = wash_servers.CreateWashServerHandlerFunc(func(params wash_servers.CreateWashServerParams, principal *app.Auth) wash_servers.CreateWashServerResponder {
			return wash_servers.CreateWashServerNotImplemented()
		})
	}
	if api.WashServersDeleteHandler == nil {
		api.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(func(params wash_servers.DeleteParams, principal *app.Auth) wash_servers.DeleteResponder {
			return wash_servers.DeleteNotImplemented()
		})
	}
	if api.UsersGetCurrentUserHandler == nil {
		api.UsersGetCurrentUserHandler = users.GetCurrentUserHandlerFunc(func(params users.GetCurrentUserParams, principal *app.Auth) users.GetCurrentUserResponder {
			return users.GetCurrentUserNotImplemented()
		})
	}
	if api.SessionsGetSessionByIDHandler == nil {
		api.SessionsGetSessionByIDHandler = sessions.GetSessionByIDHandlerFunc(func(params sessions.GetSessionByIDParams, principal *app.Auth) sessions.GetSessionByIDResponder {
			return sessions.GetSessionByIDNotImplemented()
		})
	}
	if api.UsersGetUserByIDHandler == nil {
		api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams, principal *app.Auth) users.GetUserByIDResponder {
			return users.GetUserByIDNotImplemented()
		})
	}
	if api.WashServersGetWashServerByIDHandler == nil {
		api.WashServersGetWashServerByIDHandler = wash_servers.GetWashServerByIDHandlerFunc(func(params wash_servers.GetWashServerByIDParams, principal *app.Auth) wash_servers.GetWashServerByIDResponder {
			return wash_servers.GetWashServerByIDNotImplemented()
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *app.Auth) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		})
	}
	if api.WashServersListHandler == nil {
		api.WashServersListHandler = wash_servers.ListHandlerFunc(func(params wash_servers.ListParams, principal *app.Auth) wash_servers.ListResponder {
			return wash_servers.ListNotImplemented()
		})
	}
	if api.UsersUpdateUserHandler == nil {
		api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(func(params users.UpdateUserParams, principal *app.Auth) users.UpdateUserResponder {
			return users.UpdateUserNotImplemented()
		})
	}
	if api.WashServersUpdateWashServerHandler == nil {
		api.WashServersUpdateWashServerHandler = wash_servers.UpdateWashServerHandlerFunc(func(params wash_servers.UpdateWashServerParams, principal *app.Auth) wash_servers.UpdateWashServerResponder {
			return wash_servers.UpdateWashServerNotImplemented()
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
