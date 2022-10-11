// Code generated by go-swagger; DO NOT EDIT.

package balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteBalanceHandlerFunc turns a function with the right signature into a delete balance handler
type DeleteBalanceHandlerFunc func(DeleteBalanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBalanceHandlerFunc) Handle(params DeleteBalanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteBalanceHandler interface for that can handle valid delete balance params
type DeleteBalanceHandler interface {
	Handle(DeleteBalanceParams, interface{}) middleware.Responder
}

// NewDeleteBalance creates a new http.Handler for the delete balance operation
func NewDeleteBalance(ctx *middleware.Context, handler DeleteBalanceHandler) *DeleteBalance {
	return &DeleteBalance{Context: ctx, Handler: handler}
}

/*
	DeleteBalance swagger:route DELETE /balance/deleted Balance deleteBalance

DeleteBalance delete balance API
*/
type DeleteBalance struct {
	Context *middleware.Context
	Handler DeleteBalanceHandler
}

func (o *DeleteBalance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteBalanceParams()
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

// DeleteBalanceBody delete balance body
//
// swagger:model DeleteBalanceBody
type DeleteBalanceBody struct {

	// id
	ID string `json:"id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this delete balance body
func (o *DeleteBalanceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete balance body based on context it is used
func (o *DeleteBalanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBalanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBalanceBody) UnmarshalBinary(b []byte) error {
	var res DeleteBalanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
