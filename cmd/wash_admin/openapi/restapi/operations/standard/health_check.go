// Code generated by go-swagger; DO NOT EDIT.

package standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash_admin/internal/app"
)

// HealthCheckHandlerFunc turns a function with the right signature into a health check handler
type HealthCheckHandlerFunc func(HealthCheckParams, *app.Auth) HealthCheckResponder

// Handle executing the request and returning a response
func (fn HealthCheckHandlerFunc) Handle(params HealthCheckParams, principal *app.Auth) HealthCheckResponder {
	return fn(params, principal)
}

// HealthCheckHandler interface for that can handle valid health check params
type HealthCheckHandler interface {
	Handle(HealthCheckParams, *app.Auth) HealthCheckResponder
}

// NewHealthCheck creates a new http.Handler for the health check operation
func NewHealthCheck(ctx *middleware.Context, handler HealthCheckHandler) *HealthCheck {
	return &HealthCheck{Context: ctx, Handler: handler}
}

/*
	HealthCheck swagger:route GET /healthCheck Standard healthCheck

HealthCheck health check API
*/
type HealthCheck struct {
	Context *middleware.Context
	Handler HealthCheckHandler
}

func (o *HealthCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewHealthCheckParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *app.Auth
	if uprinc != nil {
		principal = uprinc.(*app.Auth) // this is really a app.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// HealthCheckOKBody health check o k body
//
// swagger:model HealthCheckOKBody
type HealthCheckOKBody struct {

	// ok
	Ok bool `json:"ok,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *HealthCheckOKBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// ok
		Ok bool `json:"ok,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Ok = props.Ok
	return nil
}

// Validate validates this health check o k body
func (o *HealthCheckOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health check o k body based on context it is used
func (o *HealthCheckOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HealthCheckOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HealthCheckOKBody) UnmarshalBinary(b []byte) error {
	var res HealthCheckOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
