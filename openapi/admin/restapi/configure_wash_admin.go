// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"washbonus/internal/app"
	"washbonus/openapi/admin/restapi/operations"
	"washbonus/openapi/admin/restapi/operations/applications"
	"washbonus/openapi/admin/restapi/operations/organizations"
	"washbonus/openapi/admin/restapi/operations/server_groups"
	"washbonus/openapi/admin/restapi/operations/users"
	"washbonus/openapi/admin/restapi/operations/wash_servers"
)

//go:generate swagger generate server --target ../../admin --name WashAdmin --spec ../../admin.swagger.yaml --principal washbonus/internal/app.AdminAuth --exclude-main --strict-responders

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
		api.AuthKeyAuth = func(token string) (*app.AdminAuth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.WashServersAssignServerToGroupHandler == nil {
		api.WashServersAssignServerToGroupHandler = wash_servers.AssignServerToGroupHandlerFunc(func(params wash_servers.AssignServerToGroupParams, principal *app.AdminAuth) wash_servers.AssignServerToGroupResponder {
			return wash_servers.AssignServerToGroupNotImplemented()
		})
	}
	if api.OrganizationsAssignUserToOrganizationHandler == nil {
		api.OrganizationsAssignUserToOrganizationHandler = organizations.AssignUserToOrganizationHandlerFunc(func(params organizations.AssignUserToOrganizationParams, principal *app.AdminAuth) organizations.AssignUserToOrganizationResponder {
			return organizations.AssignUserToOrganizationNotImplemented()
		})
	}
	if api.ApplicationsCreateAdminApplicationHandler == nil {
		api.ApplicationsCreateAdminApplicationHandler = applications.CreateAdminApplicationHandlerFunc(func(params applications.CreateAdminApplicationParams) applications.CreateAdminApplicationResponder {
			return applications.CreateAdminApplicationNotImplemented()
		})
	}
	if api.OrganizationsCreateOrganizationHandler == nil {
		api.OrganizationsCreateOrganizationHandler = organizations.CreateOrganizationHandlerFunc(func(params organizations.CreateOrganizationParams, principal *app.AdminAuth) organizations.CreateOrganizationResponder {
			return organizations.CreateOrganizationNotImplemented()
		})
	}
	if api.ServerGroupsCreateServerGroupHandler == nil {
		api.ServerGroupsCreateServerGroupHandler = server_groups.CreateServerGroupHandlerFunc(func(params server_groups.CreateServerGroupParams, principal *app.AdminAuth) server_groups.CreateServerGroupResponder {
			return server_groups.CreateServerGroupNotImplemented()
		})
	}
	if api.WashServersCreateWashServerHandler == nil {
		api.WashServersCreateWashServerHandler = wash_servers.CreateWashServerHandlerFunc(func(params wash_servers.CreateWashServerParams, principal *app.AdminAuth) wash_servers.CreateWashServerResponder {
			return wash_servers.CreateWashServerNotImplemented()
		})
	}
	if api.UsersDeleteAdminUserHandler == nil {
		api.UsersDeleteAdminUserHandler = users.DeleteAdminUserHandlerFunc(func(params users.DeleteAdminUserParams, principal *app.AdminAuth) users.DeleteAdminUserResponder {
			return users.DeleteAdminUserNotImplemented()
		})
	}
	if api.OrganizationsDeleteOrganizationHandler == nil {
		api.OrganizationsDeleteOrganizationHandler = organizations.DeleteOrganizationHandlerFunc(func(params organizations.DeleteOrganizationParams, principal *app.AdminAuth) organizations.DeleteOrganizationResponder {
			return organizations.DeleteOrganizationNotImplemented()
		})
	}
	if api.ServerGroupsDeleteServerGroupHandler == nil {
		api.ServerGroupsDeleteServerGroupHandler = server_groups.DeleteServerGroupHandlerFunc(func(params server_groups.DeleteServerGroupParams, principal *app.AdminAuth) server_groups.DeleteServerGroupResponder {
			return server_groups.DeleteServerGroupNotImplemented()
		})
	}
	if api.WashServersDeleteWashServerHandler == nil {
		api.WashServersDeleteWashServerHandler = wash_servers.DeleteWashServerHandlerFunc(func(params wash_servers.DeleteWashServerParams, principal *app.AdminAuth) wash_servers.DeleteWashServerResponder {
			return wash_servers.DeleteWashServerNotImplemented()
		})
	}
	if api.ApplicationsGetAdminApplicationsHandler == nil {
		api.ApplicationsGetAdminApplicationsHandler = applications.GetAdminApplicationsHandlerFunc(func(params applications.GetAdminApplicationsParams, principal *app.AdminAuth) applications.GetAdminApplicationsResponder {
			return applications.GetAdminApplicationsNotImplemented()
		})
	}
	if api.UsersGetAdminUserByIDHandler == nil {
		api.UsersGetAdminUserByIDHandler = users.GetAdminUserByIDHandlerFunc(func(params users.GetAdminUserByIDParams, principal *app.AdminAuth) users.GetAdminUserByIDResponder {
			return users.GetAdminUserByIDNotImplemented()
		})
	}
	if api.UsersGetAdminUsersHandler == nil {
		api.UsersGetAdminUsersHandler = users.GetAdminUsersHandlerFunc(func(params users.GetAdminUsersParams, principal *app.AdminAuth) users.GetAdminUsersResponder {
			return users.GetAdminUsersNotImplemented()
		})
	}
	if api.OrganizationsGetOrganizationByIDHandler == nil {
		api.OrganizationsGetOrganizationByIDHandler = organizations.GetOrganizationByIDHandlerFunc(func(params organizations.GetOrganizationByIDParams, principal *app.AdminAuth) organizations.GetOrganizationByIDResponder {
			return organizations.GetOrganizationByIDNotImplemented()
		})
	}
	if api.OrganizationsGetOrganizationsHandler == nil {
		api.OrganizationsGetOrganizationsHandler = organizations.GetOrganizationsHandlerFunc(func(params organizations.GetOrganizationsParams, principal *app.AdminAuth) organizations.GetOrganizationsResponder {
			return organizations.GetOrganizationsNotImplemented()
		})
	}
	if api.ServerGroupsGetServerGroupByIDHandler == nil {
		api.ServerGroupsGetServerGroupByIDHandler = server_groups.GetServerGroupByIDHandlerFunc(func(params server_groups.GetServerGroupByIDParams, principal *app.AdminAuth) server_groups.GetServerGroupByIDResponder {
			return server_groups.GetServerGroupByIDNotImplemented()
		})
	}
	if api.ServerGroupsGetServerGroupsHandler == nil {
		api.ServerGroupsGetServerGroupsHandler = server_groups.GetServerGroupsHandlerFunc(func(params server_groups.GetServerGroupsParams, principal *app.AdminAuth) server_groups.GetServerGroupsResponder {
			return server_groups.GetServerGroupsNotImplemented()
		})
	}
	if api.WashServersGetWashServerByIDHandler == nil {
		api.WashServersGetWashServerByIDHandler = wash_servers.GetWashServerByIDHandlerFunc(func(params wash_servers.GetWashServerByIDParams, principal *app.AdminAuth) wash_servers.GetWashServerByIDResponder {
			return wash_servers.GetWashServerByIDNotImplemented()
		})
	}
	if api.WashServersGetWashServersHandler == nil {
		api.WashServersGetWashServersHandler = wash_servers.GetWashServersHandlerFunc(func(params wash_servers.GetWashServersParams, principal *app.AdminAuth) wash_servers.GetWashServersResponder {
			return wash_servers.GetWashServersNotImplemented()
		})
	}
	if api.OrganizationsRemoveUserFromOrganizationHandler == nil {
		api.OrganizationsRemoveUserFromOrganizationHandler = organizations.RemoveUserFromOrganizationHandlerFunc(func(params organizations.RemoveUserFromOrganizationParams, principal *app.AdminAuth) organizations.RemoveUserFromOrganizationResponder {
			return organizations.RemoveUserFromOrganizationNotImplemented()
		})
	}
	if api.ApplicationsReviewAdminApplicationHandler == nil {
		api.ApplicationsReviewAdminApplicationHandler = applications.ReviewAdminApplicationHandlerFunc(func(params applications.ReviewAdminApplicationParams, principal *app.AdminAuth) applications.ReviewAdminApplicationResponder {
			return applications.ReviewAdminApplicationNotImplemented()
		})
	}
	if api.UsersUpdateAdminUserRoleHandler == nil {
		api.UsersUpdateAdminUserRoleHandler = users.UpdateAdminUserRoleHandlerFunc(func(params users.UpdateAdminUserRoleParams, principal *app.AdminAuth) users.UpdateAdminUserRoleResponder {
			return users.UpdateAdminUserRoleNotImplemented()
		})
	}
	if api.OrganizationsUpdateOrganizationHandler == nil {
		api.OrganizationsUpdateOrganizationHandler = organizations.UpdateOrganizationHandlerFunc(func(params organizations.UpdateOrganizationParams, principal *app.AdminAuth) organizations.UpdateOrganizationResponder {
			return organizations.UpdateOrganizationNotImplemented()
		})
	}
	if api.ServerGroupsUpdateServerGroupHandler == nil {
		api.ServerGroupsUpdateServerGroupHandler = server_groups.UpdateServerGroupHandlerFunc(func(params server_groups.UpdateServerGroupParams, principal *app.AdminAuth) server_groups.UpdateServerGroupResponder {
			return server_groups.UpdateServerGroupNotImplemented()
		})
	}
	if api.WashServersUpdateWashServerHandler == nil {
		api.WashServersUpdateWashServerHandler = wash_servers.UpdateWashServerHandlerFunc(func(params wash_servers.UpdateWashServerParams, principal *app.AdminAuth) wash_servers.UpdateWashServerResponder {
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
