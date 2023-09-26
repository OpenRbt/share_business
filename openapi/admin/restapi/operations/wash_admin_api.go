// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"washbonus/internal/app"
	"washbonus/openapi/admin/restapi/operations/applications"
	"washbonus/openapi/admin/restapi/operations/organizations"
	"washbonus/openapi/admin/restapi/operations/server_groups"
	"washbonus/openapi/admin/restapi/operations/users"
	"washbonus/openapi/admin/restapi/operations/wash_servers"
)

// NewWashAdminAPI creates a new WashAdmin instance
func NewWashAdminAPI(spec *loads.Document) *WashAdminAPI {
	return &WashAdminAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		WashServersAssignServerToGroupHandler: wash_servers.AssignServerToGroupHandlerFunc(func(params wash_servers.AssignServerToGroupParams, principal *app.AdminAuth) wash_servers.AssignServerToGroupResponder {
			return wash_servers.AssignServerToGroupNotImplemented()
		}),
		OrganizationsAssignUserToOrganizationHandler: organizations.AssignUserToOrganizationHandlerFunc(func(params organizations.AssignUserToOrganizationParams, principal *app.AdminAuth) organizations.AssignUserToOrganizationResponder {
			return organizations.AssignUserToOrganizationNotImplemented()
		}),
		UsersBlockAdminUserHandler: users.BlockAdminUserHandlerFunc(func(params users.BlockAdminUserParams, principal *app.AdminAuth) users.BlockAdminUserResponder {
			return users.BlockAdminUserNotImplemented()
		}),
		ApplicationsCreateAdminApplicationHandler: applications.CreateAdminApplicationHandlerFunc(func(params applications.CreateAdminApplicationParams) applications.CreateAdminApplicationResponder {
			return applications.CreateAdminApplicationNotImplemented()
		}),
		OrganizationsCreateOrganizationHandler: organizations.CreateOrganizationHandlerFunc(func(params organizations.CreateOrganizationParams, principal *app.AdminAuth) organizations.CreateOrganizationResponder {
			return organizations.CreateOrganizationNotImplemented()
		}),
		ServerGroupsCreateServerGroupHandler: server_groups.CreateServerGroupHandlerFunc(func(params server_groups.CreateServerGroupParams, principal *app.AdminAuth) server_groups.CreateServerGroupResponder {
			return server_groups.CreateServerGroupNotImplemented()
		}),
		WashServersCreateWashServerHandler: wash_servers.CreateWashServerHandlerFunc(func(params wash_servers.CreateWashServerParams, principal *app.AdminAuth) wash_servers.CreateWashServerResponder {
			return wash_servers.CreateWashServerNotImplemented()
		}),
		OrganizationsDeleteOrganizationHandler: organizations.DeleteOrganizationHandlerFunc(func(params organizations.DeleteOrganizationParams, principal *app.AdminAuth) organizations.DeleteOrganizationResponder {
			return organizations.DeleteOrganizationNotImplemented()
		}),
		ServerGroupsDeleteServerGroupHandler: server_groups.DeleteServerGroupHandlerFunc(func(params server_groups.DeleteServerGroupParams, principal *app.AdminAuth) server_groups.DeleteServerGroupResponder {
			return server_groups.DeleteServerGroupNotImplemented()
		}),
		WashServersDeleteWashServerHandler: wash_servers.DeleteWashServerHandlerFunc(func(params wash_servers.DeleteWashServerParams, principal *app.AdminAuth) wash_servers.DeleteWashServerResponder {
			return wash_servers.DeleteWashServerNotImplemented()
		}),
		ApplicationsGetAdminApplicationByIDHandler: applications.GetAdminApplicationByIDHandlerFunc(func(params applications.GetAdminApplicationByIDParams, principal *app.AdminAuth) applications.GetAdminApplicationByIDResponder {
			return applications.GetAdminApplicationByIDNotImplemented()
		}),
		ApplicationsGetAdminApplicationsHandler: applications.GetAdminApplicationsHandlerFunc(func(params applications.GetAdminApplicationsParams, principal *app.AdminAuth) applications.GetAdminApplicationsResponder {
			return applications.GetAdminApplicationsNotImplemented()
		}),
		UsersGetAdminUserByIDHandler: users.GetAdminUserByIDHandlerFunc(func(params users.GetAdminUserByIDParams, principal *app.AdminAuth) users.GetAdminUserByIDResponder {
			return users.GetAdminUserByIDNotImplemented()
		}),
		UsersGetAdminUsersHandler: users.GetAdminUsersHandlerFunc(func(params users.GetAdminUsersParams, principal *app.AdminAuth) users.GetAdminUsersResponder {
			return users.GetAdminUsersNotImplemented()
		}),
		OrganizationsGetOrganizationByIDHandler: organizations.GetOrganizationByIDHandlerFunc(func(params organizations.GetOrganizationByIDParams, principal *app.AdminAuth) organizations.GetOrganizationByIDResponder {
			return organizations.GetOrganizationByIDNotImplemented()
		}),
		OrganizationsGetOrganizationsHandler: organizations.GetOrganizationsHandlerFunc(func(params organizations.GetOrganizationsParams, principal *app.AdminAuth) organizations.GetOrganizationsResponder {
			return organizations.GetOrganizationsNotImplemented()
		}),
		ServerGroupsGetServerGroupByIDHandler: server_groups.GetServerGroupByIDHandlerFunc(func(params server_groups.GetServerGroupByIDParams, principal *app.AdminAuth) server_groups.GetServerGroupByIDResponder {
			return server_groups.GetServerGroupByIDNotImplemented()
		}),
		ServerGroupsGetServerGroupsHandler: server_groups.GetServerGroupsHandlerFunc(func(params server_groups.GetServerGroupsParams, principal *app.AdminAuth) server_groups.GetServerGroupsResponder {
			return server_groups.GetServerGroupsNotImplemented()
		}),
		WashServersGetWashServerByIDHandler: wash_servers.GetWashServerByIDHandlerFunc(func(params wash_servers.GetWashServerByIDParams, principal *app.AdminAuth) wash_servers.GetWashServerByIDResponder {
			return wash_servers.GetWashServerByIDNotImplemented()
		}),
		WashServersGetWashServersHandler: wash_servers.GetWashServersHandlerFunc(func(params wash_servers.GetWashServersParams, principal *app.AdminAuth) wash_servers.GetWashServersResponder {
			return wash_servers.GetWashServersNotImplemented()
		}),
		OrganizationsRemoveUserFromOrganizationHandler: organizations.RemoveUserFromOrganizationHandlerFunc(func(params organizations.RemoveUserFromOrganizationParams, principal *app.AdminAuth) organizations.RemoveUserFromOrganizationResponder {
			return organizations.RemoveUserFromOrganizationNotImplemented()
		}),
		ApplicationsReviewAdminApplicationHandler: applications.ReviewAdminApplicationHandlerFunc(func(params applications.ReviewAdminApplicationParams, principal *app.AdminAuth) applications.ReviewAdminApplicationResponder {
			return applications.ReviewAdminApplicationNotImplemented()
		}),
		UsersUpdateAdminUserRoleHandler: users.UpdateAdminUserRoleHandlerFunc(func(params users.UpdateAdminUserRoleParams, principal *app.AdminAuth) users.UpdateAdminUserRoleResponder {
			return users.UpdateAdminUserRoleNotImplemented()
		}),
		OrganizationsUpdateOrganizationHandler: organizations.UpdateOrganizationHandlerFunc(func(params organizations.UpdateOrganizationParams, principal *app.AdminAuth) organizations.UpdateOrganizationResponder {
			return organizations.UpdateOrganizationNotImplemented()
		}),
		ServerGroupsUpdateServerGroupHandler: server_groups.UpdateServerGroupHandlerFunc(func(params server_groups.UpdateServerGroupParams, principal *app.AdminAuth) server_groups.UpdateServerGroupResponder {
			return server_groups.UpdateServerGroupNotImplemented()
		}),
		WashServersUpdateWashServerHandler: wash_servers.UpdateWashServerHandlerFunc(func(params wash_servers.UpdateWashServerParams, principal *app.AdminAuth) wash_servers.UpdateWashServerResponder {
			return wash_servers.UpdateWashServerNotImplemented()
		}),

		// Applies when the "Authorization" header is set
		AuthKeyAuth: func(token string) (*app.AdminAuth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WashAdminAPI Admin service for self-service car washes */
type WashAdminAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AuthKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	AuthKeyAuth func(string) (*app.AdminAuth, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// WashServersAssignServerToGroupHandler sets the operation handler for the assign server to group operation
	WashServersAssignServerToGroupHandler wash_servers.AssignServerToGroupHandler
	// OrganizationsAssignUserToOrganizationHandler sets the operation handler for the assign user to organization operation
	OrganizationsAssignUserToOrganizationHandler organizations.AssignUserToOrganizationHandler
	// UsersBlockAdminUserHandler sets the operation handler for the block admin user operation
	UsersBlockAdminUserHandler users.BlockAdminUserHandler
	// ApplicationsCreateAdminApplicationHandler sets the operation handler for the create admin application operation
	ApplicationsCreateAdminApplicationHandler applications.CreateAdminApplicationHandler
	// OrganizationsCreateOrganizationHandler sets the operation handler for the create organization operation
	OrganizationsCreateOrganizationHandler organizations.CreateOrganizationHandler
	// ServerGroupsCreateServerGroupHandler sets the operation handler for the create server group operation
	ServerGroupsCreateServerGroupHandler server_groups.CreateServerGroupHandler
	// WashServersCreateWashServerHandler sets the operation handler for the create wash server operation
	WashServersCreateWashServerHandler wash_servers.CreateWashServerHandler
	// OrganizationsDeleteOrganizationHandler sets the operation handler for the delete organization operation
	OrganizationsDeleteOrganizationHandler organizations.DeleteOrganizationHandler
	// ServerGroupsDeleteServerGroupHandler sets the operation handler for the delete server group operation
	ServerGroupsDeleteServerGroupHandler server_groups.DeleteServerGroupHandler
	// WashServersDeleteWashServerHandler sets the operation handler for the delete wash server operation
	WashServersDeleteWashServerHandler wash_servers.DeleteWashServerHandler
	// ApplicationsGetAdminApplicationByIDHandler sets the operation handler for the get admin application by Id operation
	ApplicationsGetAdminApplicationByIDHandler applications.GetAdminApplicationByIDHandler
	// ApplicationsGetAdminApplicationsHandler sets the operation handler for the get admin applications operation
	ApplicationsGetAdminApplicationsHandler applications.GetAdminApplicationsHandler
	// UsersGetAdminUserByIDHandler sets the operation handler for the get admin user by Id operation
	UsersGetAdminUserByIDHandler users.GetAdminUserByIDHandler
	// UsersGetAdminUsersHandler sets the operation handler for the get admin users operation
	UsersGetAdminUsersHandler users.GetAdminUsersHandler
	// OrganizationsGetOrganizationByIDHandler sets the operation handler for the get organization by Id operation
	OrganizationsGetOrganizationByIDHandler organizations.GetOrganizationByIDHandler
	// OrganizationsGetOrganizationsHandler sets the operation handler for the get organizations operation
	OrganizationsGetOrganizationsHandler organizations.GetOrganizationsHandler
	// ServerGroupsGetServerGroupByIDHandler sets the operation handler for the get server group by Id operation
	ServerGroupsGetServerGroupByIDHandler server_groups.GetServerGroupByIDHandler
	// ServerGroupsGetServerGroupsHandler sets the operation handler for the get server groups operation
	ServerGroupsGetServerGroupsHandler server_groups.GetServerGroupsHandler
	// WashServersGetWashServerByIDHandler sets the operation handler for the get wash server by Id operation
	WashServersGetWashServerByIDHandler wash_servers.GetWashServerByIDHandler
	// WashServersGetWashServersHandler sets the operation handler for the get wash servers operation
	WashServersGetWashServersHandler wash_servers.GetWashServersHandler
	// OrganizationsRemoveUserFromOrganizationHandler sets the operation handler for the remove user from organization operation
	OrganizationsRemoveUserFromOrganizationHandler organizations.RemoveUserFromOrganizationHandler
	// ApplicationsReviewAdminApplicationHandler sets the operation handler for the review admin application operation
	ApplicationsReviewAdminApplicationHandler applications.ReviewAdminApplicationHandler
	// UsersUpdateAdminUserRoleHandler sets the operation handler for the update admin user role operation
	UsersUpdateAdminUserRoleHandler users.UpdateAdminUserRoleHandler
	// OrganizationsUpdateOrganizationHandler sets the operation handler for the update organization operation
	OrganizationsUpdateOrganizationHandler organizations.UpdateOrganizationHandler
	// ServerGroupsUpdateServerGroupHandler sets the operation handler for the update server group operation
	ServerGroupsUpdateServerGroupHandler server_groups.UpdateServerGroupHandler
	// WashServersUpdateWashServerHandler sets the operation handler for the update wash server operation
	WashServersUpdateWashServerHandler wash_servers.UpdateWashServerHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *WashAdminAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *WashAdminAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *WashAdminAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WashAdminAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WashAdminAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WashAdminAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WashAdminAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WashAdminAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WashAdminAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WashAdminAPI
func (o *WashAdminAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AuthKeyAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.WashServersAssignServerToGroupHandler == nil {
		unregistered = append(unregistered, "wash_servers.AssignServerToGroupHandler")
	}
	if o.OrganizationsAssignUserToOrganizationHandler == nil {
		unregistered = append(unregistered, "organizations.AssignUserToOrganizationHandler")
	}
	if o.UsersBlockAdminUserHandler == nil {
		unregistered = append(unregistered, "users.BlockAdminUserHandler")
	}
	if o.ApplicationsCreateAdminApplicationHandler == nil {
		unregistered = append(unregistered, "applications.CreateAdminApplicationHandler")
	}
	if o.OrganizationsCreateOrganizationHandler == nil {
		unregistered = append(unregistered, "organizations.CreateOrganizationHandler")
	}
	if o.ServerGroupsCreateServerGroupHandler == nil {
		unregistered = append(unregistered, "server_groups.CreateServerGroupHandler")
	}
	if o.WashServersCreateWashServerHandler == nil {
		unregistered = append(unregistered, "wash_servers.CreateWashServerHandler")
	}
	if o.OrganizationsDeleteOrganizationHandler == nil {
		unregistered = append(unregistered, "organizations.DeleteOrganizationHandler")
	}
	if o.ServerGroupsDeleteServerGroupHandler == nil {
		unregistered = append(unregistered, "server_groups.DeleteServerGroupHandler")
	}
	if o.WashServersDeleteWashServerHandler == nil {
		unregistered = append(unregistered, "wash_servers.DeleteWashServerHandler")
	}
	if o.ApplicationsGetAdminApplicationByIDHandler == nil {
		unregistered = append(unregistered, "applications.GetAdminApplicationByIDHandler")
	}
	if o.ApplicationsGetAdminApplicationsHandler == nil {
		unregistered = append(unregistered, "applications.GetAdminApplicationsHandler")
	}
	if o.UsersGetAdminUserByIDHandler == nil {
		unregistered = append(unregistered, "users.GetAdminUserByIDHandler")
	}
	if o.UsersGetAdminUsersHandler == nil {
		unregistered = append(unregistered, "users.GetAdminUsersHandler")
	}
	if o.OrganizationsGetOrganizationByIDHandler == nil {
		unregistered = append(unregistered, "organizations.GetOrganizationByIDHandler")
	}
	if o.OrganizationsGetOrganizationsHandler == nil {
		unregistered = append(unregistered, "organizations.GetOrganizationsHandler")
	}
	if o.ServerGroupsGetServerGroupByIDHandler == nil {
		unregistered = append(unregistered, "server_groups.GetServerGroupByIDHandler")
	}
	if o.ServerGroupsGetServerGroupsHandler == nil {
		unregistered = append(unregistered, "server_groups.GetServerGroupsHandler")
	}
	if o.WashServersGetWashServerByIDHandler == nil {
		unregistered = append(unregistered, "wash_servers.GetWashServerByIDHandler")
	}
	if o.WashServersGetWashServersHandler == nil {
		unregistered = append(unregistered, "wash_servers.GetWashServersHandler")
	}
	if o.OrganizationsRemoveUserFromOrganizationHandler == nil {
		unregistered = append(unregistered, "organizations.RemoveUserFromOrganizationHandler")
	}
	if o.ApplicationsReviewAdminApplicationHandler == nil {
		unregistered = append(unregistered, "applications.ReviewAdminApplicationHandler")
	}
	if o.UsersUpdateAdminUserRoleHandler == nil {
		unregistered = append(unregistered, "users.UpdateAdminUserRoleHandler")
	}
	if o.OrganizationsUpdateOrganizationHandler == nil {
		unregistered = append(unregistered, "organizations.UpdateOrganizationHandler")
	}
	if o.ServerGroupsUpdateServerGroupHandler == nil {
		unregistered = append(unregistered, "server_groups.UpdateServerGroupHandler")
	}
	if o.WashServersUpdateWashServerHandler == nil {
		unregistered = append(unregistered, "wash_servers.UpdateWashServerHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WashAdminAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WashAdminAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "authKey":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.AuthKeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *WashAdminAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WashAdminAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *WashAdminAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WashAdminAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the wash admin API
func (o *WashAdminAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WashAdminAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/server-groups/{groupId}/wash-servers/{serverId}"] = wash_servers.NewAssignServerToGroup(o.context, o.WashServersAssignServerToGroupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/organizations/{organizationId}/users/{userId}"] = organizations.NewAssignUserToOrganization(o.context, o.OrganizationsAssignUserToOrganizationHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{userId}"] = users.NewBlockAdminUser(o.context, o.UsersBlockAdminUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/applications"] = applications.NewCreateAdminApplication(o.context, o.ApplicationsCreateAdminApplicationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/organizations"] = organizations.NewCreateOrganization(o.context, o.OrganizationsCreateOrganizationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/server-groups"] = server_groups.NewCreateServerGroup(o.context, o.ServerGroupsCreateServerGroupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/wash-servers"] = wash_servers.NewCreateWashServer(o.context, o.WashServersCreateWashServerHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/organizations/{organizationId}"] = organizations.NewDeleteOrganization(o.context, o.OrganizationsDeleteOrganizationHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/server-groups/{groupId}"] = server_groups.NewDeleteServerGroup(o.context, o.ServerGroupsDeleteServerGroupHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/wash-servers/{serverId}"] = wash_servers.NewDeleteWashServer(o.context, o.WashServersDeleteWashServerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/applications/{id}"] = applications.NewGetAdminApplicationByID(o.context, o.ApplicationsGetAdminApplicationByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/applications"] = applications.NewGetAdminApplications(o.context, o.ApplicationsGetAdminApplicationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{userId}"] = users.NewGetAdminUserByID(o.context, o.UsersGetAdminUserByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = users.NewGetAdminUsers(o.context, o.UsersGetAdminUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/organizations/{organizationId}"] = organizations.NewGetOrganizationByID(o.context, o.OrganizationsGetOrganizationByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/organizations"] = organizations.NewGetOrganizations(o.context, o.OrganizationsGetOrganizationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/server-groups/{groupId}"] = server_groups.NewGetServerGroupByID(o.context, o.ServerGroupsGetServerGroupByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/server-groups"] = server_groups.NewGetServerGroups(o.context, o.ServerGroupsGetServerGroupsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wash-servers/{serverId}"] = wash_servers.NewGetWashServerByID(o.context, o.WashServersGetWashServerByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wash-servers"] = wash_servers.NewGetWashServers(o.context, o.WashServersGetWashServersHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/organizations/{organizationId}/users/{userId}"] = organizations.NewRemoveUserFromOrganization(o.context, o.OrganizationsRemoveUserFromOrganizationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/applications/{id}"] = applications.NewReviewAdminApplication(o.context, o.ApplicationsReviewAdminApplicationHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/users/{userId}"] = users.NewUpdateAdminUserRole(o.context, o.UsersUpdateAdminUserRoleHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/organizations/{organizationId}"] = organizations.NewUpdateOrganization(o.context, o.OrganizationsUpdateOrganizationHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/server-groups/{groupId}"] = server_groups.NewUpdateServerGroup(o.context, o.ServerGroupsUpdateServerGroupHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/wash-servers/{serverId}"] = wash_servers.NewUpdateWashServer(o.context, o.WashServersUpdateWashServerHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WashAdminAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WashAdminAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WashAdminAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WashAdminAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WashAdminAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
