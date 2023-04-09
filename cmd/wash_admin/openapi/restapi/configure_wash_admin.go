// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"wash_admin/internal/app"
	"wash_admin/openapi/restapi/operations"
	"wash_admin/openapi/restapi/operations/standard"
	"wash_admin/openapi/restapi/operations/wash_servers"
)

//go:generate swagger generate server --target ../../openapi --name WashAdmin --spec ../swagger.yaml --principal wash_admin/internal/app.Auth --exclude-main --strict-responders

func configureFlags(api *operations.WashAdminAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WashAdminAPI) http.Handler {
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

	if api.WashServersAddHandler == nil {
		api.WashServersAddHandler = wash_servers.AddHandlerFunc(func(params wash_servers.AddParams, principal *app.Auth) wash_servers.AddResponder {
			return wash_servers.AddNotImplemented()
		})
	}
	if api.WashServersDeleteHandler == nil {
		api.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(func(params wash_servers.DeleteParams, principal *app.Auth) wash_servers.DeleteResponder {
			return wash_servers.DeleteNotImplemented()
		})
	}
	if api.WashServersGetWashServerHandler == nil {
		api.WashServersGetWashServerHandler = wash_servers.GetWashServerHandlerFunc(func(params wash_servers.GetWashServerParams, principal *app.Auth) wash_servers.GetWashServerResponder {
			return wash_servers.GetWashServerNotImplemented()
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
	if api.WashServersUpdateHandler == nil {
		api.WashServersUpdateHandler = wash_servers.UpdateHandlerFunc(func(params wash_servers.UpdateParams, principal *app.Auth) wash_servers.UpdateResponder {
			return wash_servers.UpdateNotImplemented()
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
