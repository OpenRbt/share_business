// Code generated by go-swagger; DO NOT EDIT.

package wash_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// EditWashServerHandlerFunc turns a function with the right signature into a edit wash server handler
type EditWashServerHandlerFunc func(EditWashServerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EditWashServerHandlerFunc) Handle(params EditWashServerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EditWashServerHandler interface for that can handle valid edit wash server params
type EditWashServerHandler interface {
	Handle(EditWashServerParams, interface{}) middleware.Responder
}

// NewEditWashServer creates a new http.Handler for the edit wash server operation
func NewEditWashServer(ctx *middleware.Context, handler EditWashServerHandler) *EditWashServer {
	return &EditWashServer{Context: ctx, Handler: handler}
}

/* EditWashServer swagger:route PUT /washServer/edit WashServer editWashServer

EditWashServer edit wash server API

*/
type EditWashServer struct {
	Context *middleware.Context
	Handler EditWashServerHandler
}

func (o *EditWashServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEditWashServerParams()
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

// EditWashServerBody edit wash server body
//
// swagger:model EditWashServerBody
type EditWashServerBody struct {

	// data
	Data *models.WashServerAdd `json:"data,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this edit wash server body
func (o *EditWashServerBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditWashServerBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edit wash server body based on the context it is used
func (o *EditWashServerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditWashServerBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EditWashServerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditWashServerBody) UnmarshalBinary(b []byte) error {
	var res EditWashServerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
