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
	"washbonus/openapi/bonus/restapi/operations/sessions"
	"washbonus/openapi/bonus/restapi/operations/standard"
	"washbonus/openapi/bonus/restapi/operations/wallets"
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

		SessionsAssignUserToSessionHandler: sessions.AssignUserToSessionHandlerFunc(func(params sessions.AssignUserToSessionParams, principal *app.Auth) sessions.AssignUserToSessionResponder {
			return sessions.AssignUserToSessionNotImplemented()
		}),
		SessionsChargeBonusesOnSessionHandler: sessions.ChargeBonusesOnSessionHandlerFunc(func(params sessions.ChargeBonusesOnSessionParams, principal *app.Auth) sessions.ChargeBonusesOnSessionResponder {
			return sessions.ChargeBonusesOnSessionNotImplemented()
		}),
		SessionsGetSessionByIDHandler: sessions.GetSessionByIDHandlerFunc(func(params sessions.GetSessionByIDParams, principal *app.Auth) sessions.GetSessionByIDResponder {
			return sessions.GetSessionByIDNotImplemented()
		}),
		WalletsGetWalletByOrganizationIDHandler: wallets.GetWalletByOrganizationIDHandlerFunc(func(params wallets.GetWalletByOrganizationIDParams, principal *app.Auth) wallets.GetWalletByOrganizationIDResponder {
			return wallets.GetWalletByOrganizationIDNotImplemented()
		}),
		WalletsGetWalletsHandler: wallets.GetWalletsHandlerFunc(func(params wallets.GetWalletsParams, principal *app.Auth) wallets.GetWalletsResponder {
			return wallets.GetWalletsNotImplemented()
		}),
		StandardHealthCheckHandler: standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *app.Auth) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		}),

		// Applies when the "Authorization" header is set
		AuthKeyAuth: func(token string) (*app.Auth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WashBonusAPI Bonus system service for self-service car washes */
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
	AuthKeyAuth func(string) (*app.Auth, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// SessionsAssignUserToSessionHandler sets the operation handler for the assign user to session operation
	SessionsAssignUserToSessionHandler sessions.AssignUserToSessionHandler
	// SessionsChargeBonusesOnSessionHandler sets the operation handler for the charge bonuses on session operation
	SessionsChargeBonusesOnSessionHandler sessions.ChargeBonusesOnSessionHandler
	// SessionsGetSessionByIDHandler sets the operation handler for the get session by Id operation
	SessionsGetSessionByIDHandler sessions.GetSessionByIDHandler
	// WalletsGetWalletByOrganizationIDHandler sets the operation handler for the get wallet by organization Id operation
	WalletsGetWalletByOrganizationIDHandler wallets.GetWalletByOrganizationIDHandler
	// WalletsGetWalletsHandler sets the operation handler for the get wallets operation
	WalletsGetWalletsHandler wallets.GetWalletsHandler
	// StandardHealthCheckHandler sets the operation handler for the health check operation
	StandardHealthCheckHandler standard.HealthCheckHandler

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

	if o.SessionsAssignUserToSessionHandler == nil {
		unregistered = append(unregistered, "sessions.AssignUserToSessionHandler")
	}
	if o.SessionsChargeBonusesOnSessionHandler == nil {
		unregistered = append(unregistered, "sessions.ChargeBonusesOnSessionHandler")
	}
	if o.SessionsGetSessionByIDHandler == nil {
		unregistered = append(unregistered, "sessions.GetSessionByIDHandler")
	}
	if o.WalletsGetWalletByOrganizationIDHandler == nil {
		unregistered = append(unregistered, "wallets.GetWalletByOrganizationIDHandler")
	}
	if o.WalletsGetWalletsHandler == nil {
		unregistered = append(unregistered, "wallets.GetWalletsHandler")
	}
	if o.StandardHealthCheckHandler == nil {
		unregistered = append(unregistered, "standard.HealthCheckHandler")
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
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.AuthKeyAuth(token)
			})

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
	o.handlers["POST"]["/sessions/{sessionId}/assign-user"] = sessions.NewAssignUserToSession(o.context, o.SessionsAssignUserToSessionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sessions/{sessionId}/bonuses"] = sessions.NewChargeBonusesOnSession(o.context, o.SessionsChargeBonusesOnSessionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sessions/{sessionId}"] = sessions.NewGetSessionByID(o.context, o.SessionsGetSessionByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wallets/by-organization/{id}"] = wallets.NewGetWalletByOrganizationID(o.context, o.WalletsGetWalletByOrganizationIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wallets"] = wallets.NewGetWallets(o.context, o.WalletsGetWalletsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthCheck"] = standard.NewHealthCheck(o.context, o.StandardHealthCheckHandler)
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
		o.handlers[um][path] = builder(h)
	}
}