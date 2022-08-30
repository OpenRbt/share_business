// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"wash-bonus/internal/api/restapi/restapi/operations"
	"wash-bonus/internal/api/restapi/restapi/operations/permission"
	"wash-bonus/internal/api/restapi/restapi/operations/role"
	"wash-bonus/internal/api/restapi/restapi/operations/standard"
	"wash-bonus/internal/api/restapi/restapi/operations/user"
	"wash-bonus/internal/api/restapi/restapi/operations/wash_server"
	"wash-bonus/internal/api/restapi/restapi/operations/wash_session"
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

	if api.PermissionAddPermissionHandler == nil {
		api.PermissionAddPermissionHandler = permission.AddPermissionHandlerFunc(func(params permission.AddPermissionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation permission.AddPermission has not yet been implemented")
		})
	}
	if api.RoleAddPermissionsRoleHandler == nil {
		api.RoleAddPermissionsRoleHandler = role.AddPermissionsRoleHandlerFunc(func(params role.AddPermissionsRoleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation role.AddPermissionsRole has not yet been implemented")
		})
	}
	if api.RoleAddRoleHandler == nil {
		api.RoleAddRoleHandler = role.AddRoleHandlerFunc(func(params role.AddRoleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation role.AddRole has not yet been implemented")
		})
	}
	if api.StandardAddTestDataHandler == nil {
		api.StandardAddTestDataHandler = standard.AddTestDataHandlerFunc(func(params standard.AddTestDataParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation standard.AddTestData has not yet been implemented")
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
	if api.WashSessionAddWashSessionHandler == nil {
		api.WashSessionAddWashSessionHandler = wash_session.AddWashSessionHandlerFunc(func(params wash_session.AddWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.AddWashSession has not yet been implemented")
		})
	}
	if api.PermissionDeletePermissionHandler == nil {
		api.PermissionDeletePermissionHandler = permission.DeletePermissionHandlerFunc(func(params permission.DeletePermissionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation permission.DeletePermission has not yet been implemented")
		})
	}
	if api.RoleDeletePermissionsRoleHandler == nil {
		api.RoleDeletePermissionsRoleHandler = role.DeletePermissionsRoleHandlerFunc(func(params role.DeletePermissionsRoleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation role.DeletePermissionsRole has not yet been implemented")
		})
	}
	if api.RoleDeleteRoleHandler == nil {
		api.RoleDeleteRoleHandler = role.DeleteRoleHandlerFunc(func(params role.DeleteRoleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation role.DeleteRole has not yet been implemented")
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
	if api.WashSessionDeleteWashSessionHandler == nil {
		api.WashSessionDeleteWashSessionHandler = wash_session.DeleteWashSessionHandlerFunc(func(params wash_session.DeleteWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.DeleteWashSession has not yet been implemented")
		})
	}
	if api.PermissionEditPermissionHandler == nil {
		api.PermissionEditPermissionHandler = permission.EditPermissionHandlerFunc(func(params permission.EditPermissionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation permission.EditPermission has not yet been implemented")
		})
	}
	if api.RoleEditRoleHandler == nil {
		api.RoleEditRoleHandler = role.EditRoleHandlerFunc(func(params role.EditRoleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation role.EditRole has not yet been implemented")
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
	if api.WashSessionEditWashSessionHandler == nil {
		api.WashSessionEditWashSessionHandler = wash_session.EditWashSessionHandlerFunc(func(params wash_session.EditWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.EditWashSession has not yet been implemented")
		})
	}
	if api.PermissionGetPermissionHandler == nil {
		api.PermissionGetPermissionHandler = permission.GetPermissionHandlerFunc(func(params permission.GetPermissionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation permission.GetPermission has not yet been implemented")
		})
	}
	if api.RoleGetRoleHandler == nil {
		api.RoleGetRoleHandler = role.GetRoleHandlerFunc(func(params role.GetRoleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation role.GetRole has not yet been implemented")
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
	if api.WashSessionGetWashSessionHandler == nil {
		api.WashSessionGetWashSessionHandler = wash_session.GetWashSessionHandlerFunc(func(params wash_session.GetWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.GetWashSession has not yet been implemented")
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation standard.HealthCheck has not yet been implemented")
		})
	}
	if api.PermissionListPermissionHandler == nil {
		api.PermissionListPermissionHandler = permission.ListPermissionHandlerFunc(func(params permission.ListPermissionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation permission.ListPermission has not yet been implemented")
		})
	}
	if api.RoleListRoleHandler == nil {
		api.RoleListRoleHandler = role.ListRoleHandlerFunc(func(params role.ListRoleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation role.ListRole has not yet been implemented")
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
	if api.WashSessionListWashSessionHandler == nil {
		api.WashSessionListWashSessionHandler = wash_session.ListWashSessionHandlerFunc(func(params wash_session.ListWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.ListWashSession has not yet been implemented")
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
