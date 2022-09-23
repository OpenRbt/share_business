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

	"wash-bonus/internal/api/restapi/restapi/operations/bonus_balance"
	"wash-bonus/internal/api/restapi/restapi/operations/standard"
	"wash-bonus/internal/api/restapi/restapi/operations/user"
	"wash-bonus/internal/api/restapi/restapi/operations/wash_server"
	"wash-bonus/internal/api/restapi/restapi/operations/wash_session"
)

// NewWashBonusAPI creates a new WashBonus instance
func NewWashBonusAPI(spec *loads.Document) *WashBonusAPI {
	return &WashBonusAPI{
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

		BonusBalanceAddBonusBalanceHandler: bonus_balance.AddBonusBalanceHandlerFunc(func(params bonus_balance.AddBonusBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.AddBonusBalance has not yet been implemented")
		}),
		StandardAddTestDataHandler: standard.AddTestDataHandlerFunc(func(params standard.AddTestDataParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation standard.AddTestData has not yet been implemented")
		}),
		UserAddUserHandler: user.AddUserHandlerFunc(func(params user.AddUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.AddUser has not yet been implemented")
		}),
		WashServerAddWashServerHandler: wash_server.AddWashServerHandlerFunc(func(params wash_server.AddWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.AddWashServer has not yet been implemented")
		}),
		WashSessionAddWashSessionHandler: wash_session.AddWashSessionHandlerFunc(func(params wash_session.AddWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.AddWashSession has not yet been implemented")
		}),
		BonusBalanceDeleteBonusBalanceHandler: bonus_balance.DeleteBonusBalanceHandlerFunc(func(params bonus_balance.DeleteBonusBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.DeleteBonusBalance has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		}),
		WashServerDeleteWashServerHandler: wash_server.DeleteWashServerHandlerFunc(func(params wash_server.DeleteWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.DeleteWashServer has not yet been implemented")
		}),
		WashSessionDeleteWashSessionHandler: wash_session.DeleteWashSessionHandlerFunc(func(params wash_session.DeleteWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.DeleteWashSession has not yet been implemented")
		}),
		BonusBalanceEditBonusBalanceHandler: bonus_balance.EditBonusBalanceHandlerFunc(func(params bonus_balance.EditBonusBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.EditBonusBalance has not yet been implemented")
		}),
		UserEditUserHandler: user.EditUserHandlerFunc(func(params user.EditUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.EditUser has not yet been implemented")
		}),
		WashServerEditWashServerHandler: wash_server.EditWashServerHandlerFunc(func(params wash_server.EditWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.EditWashServer has not yet been implemented")
		}),
		WashSessionEditWashSessionHandler: wash_session.EditWashSessionHandlerFunc(func(params wash_session.EditWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.EditWashSession has not yet been implemented")
		}),
		BonusBalanceGetBonusBalanceHandler: bonus_balance.GetBonusBalanceHandlerFunc(func(params bonus_balance.GetBonusBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bonus_balance.GetBonusBalance has not yet been implemented")
		}),
		UserGetUserHandler: user.GetUserHandlerFunc(func(params user.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		}),
		WashServerGetWashServerHandler: wash_server.GetWashServerHandlerFunc(func(params wash_server.GetWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.GetWashServer has not yet been implemented")
		}),
		WashSessionGetWashSessionHandler: wash_session.GetWashSessionHandlerFunc(func(params wash_session.GetWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.GetWashSession has not yet been implemented")
		}),
		StandardHealthCheckHandler: standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation standard.HealthCheck has not yet been implemented")
		}),
		UserListUserHandler: user.ListUserHandlerFunc(func(params user.ListUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUser has not yet been implemented")
		}),
		WashServerListWashServerHandler: wash_server.ListWashServerHandlerFunc(func(params wash_server.ListWashServerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_server.ListWashServer has not yet been implemented")
		}),
		WashSessionListWashSessionHandler: wash_session.ListWashSessionHandlerFunc(func(params wash_session.ListWashSessionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wash_session.ListWashSession has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		AuthKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WashBonusAPI microservice for the bonus system of self-service car washes */
type WashBonusAPI struct {
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
	AuthKeyAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// BonusBalanceAddBonusBalanceHandler sets the operation handler for the add bonus balance operation
	BonusBalanceAddBonusBalanceHandler bonus_balance.AddBonusBalanceHandler
	// StandardAddTestDataHandler sets the operation handler for the add test data operation
	StandardAddTestDataHandler standard.AddTestDataHandler
	// UserAddUserHandler sets the operation handler for the add user operation
	UserAddUserHandler user.AddUserHandler
	// WashServerAddWashServerHandler sets the operation handler for the add wash server operation
	WashServerAddWashServerHandler wash_server.AddWashServerHandler
	// WashSessionAddWashSessionHandler sets the operation handler for the add wash session operation
	WashSessionAddWashSessionHandler wash_session.AddWashSessionHandler
	// BonusBalanceDeleteBonusBalanceHandler sets the operation handler for the delete bonus balance operation
	BonusBalanceDeleteBonusBalanceHandler bonus_balance.DeleteBonusBalanceHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// WashServerDeleteWashServerHandler sets the operation handler for the delete wash server operation
	WashServerDeleteWashServerHandler wash_server.DeleteWashServerHandler
	// WashSessionDeleteWashSessionHandler sets the operation handler for the delete wash session operation
	WashSessionDeleteWashSessionHandler wash_session.DeleteWashSessionHandler
	// BonusBalanceEditBonusBalanceHandler sets the operation handler for the edit bonus balance operation
	BonusBalanceEditBonusBalanceHandler bonus_balance.EditBonusBalanceHandler
	// UserEditUserHandler sets the operation handler for the edit user operation
	UserEditUserHandler user.EditUserHandler
	// WashServerEditWashServerHandler sets the operation handler for the edit wash server operation
	WashServerEditWashServerHandler wash_server.EditWashServerHandler
	// WashSessionEditWashSessionHandler sets the operation handler for the edit wash session operation
	WashSessionEditWashSessionHandler wash_session.EditWashSessionHandler
	// BonusBalanceGetBonusBalanceHandler sets the operation handler for the get bonus balance operation
	BonusBalanceGetBonusBalanceHandler bonus_balance.GetBonusBalanceHandler
	// UserGetUserHandler sets the operation handler for the get user operation
	UserGetUserHandler user.GetUserHandler
	// WashServerGetWashServerHandler sets the operation handler for the get wash server operation
	WashServerGetWashServerHandler wash_server.GetWashServerHandler
	// WashSessionGetWashSessionHandler sets the operation handler for the get wash session operation
	WashSessionGetWashSessionHandler wash_session.GetWashSessionHandler
	// StandardHealthCheckHandler sets the operation handler for the health check operation
	StandardHealthCheckHandler standard.HealthCheckHandler
	// UserListUserHandler sets the operation handler for the list user operation
	UserListUserHandler user.ListUserHandler
	// WashServerListWashServerHandler sets the operation handler for the list wash server operation
	WashServerListWashServerHandler wash_server.ListWashServerHandler
	// WashSessionListWashSessionHandler sets the operation handler for the list wash session operation
	WashSessionListWashSessionHandler wash_session.ListWashSessionHandler

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
func (o *WashBonusAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *WashBonusAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *WashBonusAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WashBonusAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WashBonusAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WashBonusAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WashBonusAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WashBonusAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WashBonusAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WashBonusAPI
func (o *WashBonusAPI) Validate() error {
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

	if o.BonusBalanceAddBonusBalanceHandler == nil {
		unregistered = append(unregistered, "bonus_balance.AddBonusBalanceHandler")
	}
	if o.StandardAddTestDataHandler == nil {
		unregistered = append(unregistered, "standard.AddTestDataHandler")
	}
	if o.UserAddUserHandler == nil {
		unregistered = append(unregistered, "user.AddUserHandler")
	}
	if o.WashServerAddWashServerHandler == nil {
		unregistered = append(unregistered, "wash_server.AddWashServerHandler")
	}
	if o.WashSessionAddWashSessionHandler == nil {
		unregistered = append(unregistered, "wash_session.AddWashSessionHandler")
	}
	if o.BonusBalanceDeleteBonusBalanceHandler == nil {
		unregistered = append(unregistered, "bonus_balance.DeleteBonusBalanceHandler")
	}
	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}
	if o.WashServerDeleteWashServerHandler == nil {
		unregistered = append(unregistered, "wash_server.DeleteWashServerHandler")
	}
	if o.WashSessionDeleteWashSessionHandler == nil {
		unregistered = append(unregistered, "wash_session.DeleteWashSessionHandler")
	}
	if o.BonusBalanceEditBonusBalanceHandler == nil {
		unregistered = append(unregistered, "bonus_balance.EditBonusBalanceHandler")
	}
	if o.UserEditUserHandler == nil {
		unregistered = append(unregistered, "user.EditUserHandler")
	}
	if o.WashServerEditWashServerHandler == nil {
		unregistered = append(unregistered, "wash_server.EditWashServerHandler")
	}
	if o.WashSessionEditWashSessionHandler == nil {
		unregistered = append(unregistered, "wash_session.EditWashSessionHandler")
	}
	if o.BonusBalanceGetBonusBalanceHandler == nil {
		unregistered = append(unregistered, "bonus_balance.GetBonusBalanceHandler")
	}
	if o.UserGetUserHandler == nil {
		unregistered = append(unregistered, "user.GetUserHandler")
	}
	if o.WashServerGetWashServerHandler == nil {
		unregistered = append(unregistered, "wash_server.GetWashServerHandler")
	}
	if o.WashSessionGetWashSessionHandler == nil {
		unregistered = append(unregistered, "wash_session.GetWashSessionHandler")
	}
	if o.StandardHealthCheckHandler == nil {
		unregistered = append(unregistered, "standard.HealthCheckHandler")
	}
	if o.UserListUserHandler == nil {
		unregistered = append(unregistered, "user.ListUserHandler")
	}
	if o.WashServerListWashServerHandler == nil {
		unregistered = append(unregistered, "wash_server.ListWashServerHandler")
	}
	if o.WashSessionListWashSessionHandler == nil {
		unregistered = append(unregistered, "wash_session.ListWashSessionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WashBonusAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WashBonusAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "authKey":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.AuthKeyAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *WashBonusAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WashBonusAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *WashBonusAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *WashBonusAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the wash bonus API
func (o *WashBonusAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WashBonusAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/balance/add"] = bonus_balance.NewAddBonusBalance(o.context, o.BonusBalanceAddBonusBalanceHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/addTestData"] = standard.NewAddTestData(o.context, o.StandardAddTestDataHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/add"] = user.NewAddUser(o.context, o.UserAddUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/washServer/add"] = wash_server.NewAddWashServer(o.context, o.WashServerAddWashServerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/washSession/add"] = wash_session.NewAddWashSession(o.context, o.WashSessionAddWashSessionHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/balance/deleted"] = bonus_balance.NewDeleteBonusBalance(o.context, o.BonusBalanceDeleteBonusBalanceHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/user/delete"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/washServer/delete"] = wash_server.NewDeleteWashServer(o.context, o.WashServerDeleteWashServerHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/washSession/delete"] = wash_session.NewDeleteWashSession(o.context, o.WashSessionDeleteWashSessionHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/balance/edit"] = bonus_balance.NewEditBonusBalance(o.context, o.BonusBalanceEditBonusBalanceHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/edit"] = user.NewEditUser(o.context, o.UserEditUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/washServer/edit"] = wash_server.NewEditWashServer(o.context, o.WashServerEditWashServerHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/washSession/edit"] = wash_session.NewEditWashSession(o.context, o.WashSessionEditWashSessionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/balance/get"] = bonus_balance.NewGetBonusBalance(o.context, o.BonusBalanceGetBonusBalanceHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/get"] = user.NewGetUser(o.context, o.UserGetUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/washServer/get"] = wash_server.NewGetWashServer(o.context, o.WashServerGetWashServerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/washSession/get"] = wash_session.NewGetWashSession(o.context, o.WashSessionGetWashSessionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthCheck"] = standard.NewHealthCheck(o.context, o.StandardHealthCheckHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/list"] = user.NewListUser(o.context, o.UserListUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/washServer/list"] = wash_server.NewListWashServer(o.context, o.WashServerListWashServerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/washSession/list"] = wash_session.NewListWashSession(o.context, o.WashSessionListWashSessionHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WashBonusAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *WashBonusAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WashBonusAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WashBonusAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WashBonusAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
