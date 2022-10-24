// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"wash-bonus/internal/transport/rest/restapi/restapi/operations"
	balance2 "wash-bonus/internal/transport/rest/restapi/restapi/operations/balance"
	standard2 "wash-bonus/internal/transport/rest/restapi/restapi/operations/standard"
	user2 "wash-bonus/internal/transport/rest/restapi/restapi/operations/user"
	wash_server2 "wash-bonus/internal/transport/rest/restapi/restapi/operations/wash_server"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ../../restapi --name WashBonus --spec ../../../../swagger.yaml --principal interface{} --exclude-main

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
		api.BalanceAddBalanceHandler = balance2.AddBalanceHandlerFunc(func(params balance2.AddBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.AddBalance has not yet been implemented")
		})
	}
	if api.StandardAddTestDataHandler == nil {
		api.StandardAddTestDataHandler = standard2.AddTestDataHandlerFunc(func(params standard2.AddTestDataParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation standard.AddTestData has not yet been implemented")
		})
	}
	if api.UserAddUserHandler == nil {
		api.UserAddUserHandler = user2.AddUserHandlerFunc(func(params user2.AddUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.AddUser has not yet been implemented")
		})
	}
	if api.WashServerAddWashServerHandler == nil {
		api.WashServerAddWashServerHandler = wash_server2.AddWashServerHandlerFunc(func(params wash_server2.AddWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.AddWashServer has not yet been implemented")
		})
	}
	if api.BalanceDeleteBalanceHandler == nil {
		api.BalanceDeleteBalanceHandler = balance2.DeleteBalanceHandlerFunc(func(params balance2.DeleteBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.DeleteBalance has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user2.DeleteUserHandlerFunc(func(params user2.DeleteUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.WashServerDeleteWashServerHandler == nil {
		api.WashServerDeleteWashServerHandler = wash_server2.DeleteWashServerHandlerFunc(func(params wash_server2.DeleteWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.DeleteWashServer has not yet been implemented")
		})
	}
	if api.BalanceEditBalanceHandler == nil {
		api.BalanceEditBalanceHandler = balance2.EditBalanceHandlerFunc(func(params balance2.EditBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.EditBalance has not yet been implemented")
		})
	}
	if api.UserEditUserHandler == nil {
		api.UserEditUserHandler = user2.EditUserHandlerFunc(func(params user2.EditUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.EditUser has not yet been implemented")
		})
	}
	if api.WashServerEditWashServerHandler == nil {
		api.WashServerEditWashServerHandler = wash_server2.EditWashServerHandlerFunc(func(params wash_server2.EditWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.EditWashServer has not yet been implemented")
		})
	}
	if api.BalanceGetBalanceHandler == nil {
		api.BalanceGetBalanceHandler = balance2.GetBalanceHandlerFunc(func(params balance2.GetBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.GetBalance has not yet been implemented")
		})
	}
	if api.WashServerGenerateKeyWashServerHandler == nil {
		api.WashServerGenerateKeyWashServerHandler = wash_server2.GenerateKeyWashServerHandlerFunc(func(params wash_server2.GenerateKeyWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.GenerateKeyWashServer has not yet been implemented")
		})
	}
	if api.UserGetCurrentUserHandler == nil {
		api.UserGetCurrentUserHandler = user2.GetCurrentUserHandlerFunc(func(params user2.GetCurrentUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetCurrentUser has not yet been implemented")
		})
	}
	if api.UserGetUserHandler == nil {
		api.UserGetUserHandler = user2.GetUserHandlerFunc(func(params user2.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		})
	}
	if api.WashServerGetWashServerHandler == nil {
		api.WashServerGetWashServerHandler = wash_server2.GetWashServerHandlerFunc(func(params wash_server2.GetWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.GetWashServer has not yet been implemented")
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard2.HealthCheckHandlerFunc(func(params standard2.HealthCheckParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation standard.HealthCheck has not yet been implemented")
		})
	}
	if api.UserListUserHandler == nil {
		api.UserListUserHandler = user2.ListUserHandlerFunc(func(params user2.ListUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUser has not yet been implemented")
		})
	}
	if api.WashServerListWashServerHandler == nil {
		api.WashServerListWashServerHandler = wash_server2.ListWashServerHandlerFunc(func(params wash_server2.ListWashServerParams, principal interface{}) middleware.Responder {
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
