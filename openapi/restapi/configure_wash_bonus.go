// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"wash-bonus/openapi/restapi/operations"
	"wash-bonus/openapi/restapi/operations/balance"
	"wash-bonus/openapi/restapi/operations/standard"
	"wash-bonus/openapi/restapi/operations/user"
	"wash-bonus/openapi/restapi/operations/wash_server"
)

//go:generate swagger generate server --target ../../openapi --name WashBonus --spec ../../swagger.yaml --principal interface{} --exclude-main

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
		api.AuthKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.BalanceAddBalanceHandler == nil {
		api.BalanceAddBalanceHandler = balance.AddBalanceHandlerFunc(func(params balance.AddBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation balance.AddBalance has not yet been implemented")
		})
	}
	if api.UserAddUserHandler == nil {
		api.UserAddUserHandler = user.AddUserHandlerFunc(func(params user.AddUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.AddUser has not yet been implemented")
		})
	}
	if api.WashServerAddWashServerHandler == nil {
		api.WashServerAddWashServerHandler = wash_server.AddWashServerHandlerFunc(func(params wash_server.AddWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.AddWashServer has not yet been implemented")
		})
	}
	if api.BalanceDeleteBalanceHandler == nil {
		api.BalanceDeleteBalanceHandler = balance.DeleteBalanceHandlerFunc(func(params balance.DeleteBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation balance.DeleteBalance has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.WashServerDeleteWashServerHandler == nil {
		api.WashServerDeleteWashServerHandler = wash_server.DeleteWashServerHandlerFunc(func(params wash_server.DeleteWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.DeleteWashServer has not yet been implemented")
		})
	}
	if api.BalanceEditBalanceHandler == nil {
		api.BalanceEditBalanceHandler = balance.EditBalanceHandlerFunc(func(params balance.EditBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation balance.EditBalance has not yet been implemented")
		})
	}
	if api.UserEditUserHandler == nil {
		api.UserEditUserHandler = user.EditUserHandlerFunc(func(params user.EditUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.EditUser has not yet been implemented")
		})
	}
	if api.WashServerEditWashServerHandler == nil {
		api.WashServerEditWashServerHandler = wash_server.EditWashServerHandlerFunc(func(params wash_server.EditWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.EditWashServer has not yet been implemented")
		})
	}
	if api.WashServerGenerateKeyWashServerHandler == nil {
		api.WashServerGenerateKeyWashServerHandler = wash_server.GenerateKeyWashServerHandlerFunc(func(params wash_server.GenerateKeyWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.GenerateKeyWashServer has not yet been implemented")
		})
	}
	if api.BalanceGetBalanceHandler == nil {
		api.BalanceGetBalanceHandler = balance.GetBalanceHandlerFunc(func(params balance.GetBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation balance.GetBalance has not yet been implemented")
		})
	}
	if api.UserGetCurrentUserHandler == nil {
		api.UserGetCurrentUserHandler = user.GetCurrentUserHandlerFunc(func(params user.GetCurrentUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetCurrentUser has not yet been implemented")
		})
	}
	if api.UserGetUserHandler == nil {
		api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		})
	}
	if api.WashServerGetWashServerHandler == nil {
		api.WashServerGetWashServerHandler = wash_server.GetWashServerHandlerFunc(func(params wash_server.GetWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.GetWashServer has not yet been implemented")
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation standard.HealthCheck has not yet been implemented")
		})
	}
	if api.UserListUserHandler == nil {
		api.UserListUserHandler = user.ListUserHandlerFunc(func(params user.ListUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUser has not yet been implemented")
		})
	}
	if api.WashServerListWashServerHandler == nil {
		api.WashServerListWashServerHandler = wash_server.ListWashServerHandlerFunc(func(params wash_server.ListWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.ListWashServer has not yet been implemented")
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
