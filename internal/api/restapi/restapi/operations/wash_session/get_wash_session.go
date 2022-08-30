// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetWashSessionHandlerFunc turns a function with the right signature into a get wash session handler
type GetWashSessionHandlerFunc func(GetWashSessionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWashSessionHandlerFunc) Handle(params GetWashSessionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetWashSessionHandler interface for that can handle valid get wash session params
type GetWashSessionHandler interface {
	Handle(GetWashSessionParams, interface{}) middleware.Responder
}

// NewGetWashSession creates a new http.Handler for the get wash session operation
func NewGetWashSession(ctx *middleware.Context, handler GetWashSessionHandler) *GetWashSession {
	return &GetWashSession{Context: ctx, Handler: handler}
}

/* GetWashSession swagger:route POST /washSession/get WashSession getWashSession

GetWashSession get wash session API

*/
type GetWashSession struct {
	Context *middleware.Context
	Handler GetWashSessionHandler
}

func (o *GetWashSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWashSessionParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetWashSessionBody get wash session body
//
// swagger:model GetWashSessionBody
type GetWashSessionBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this get wash session body
func (o *GetWashSessionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wash session body based on context it is used
func (o *GetWashSessionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWashSessionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWashSessionBody) UnmarshalBinary(b []byte) error {
	var res GetWashSessionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
