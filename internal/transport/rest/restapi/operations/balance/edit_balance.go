// Code generated by go-swagger; DO NOT EDIT.

package balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/transport/rest/models"
)

// EditBalanceHandlerFunc turns a function with the right signature into a edit balance handler
type EditBalanceHandlerFunc func(EditBalanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EditBalanceHandlerFunc) Handle(params EditBalanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EditBalanceHandler interface for that can handle valid edit balance params
type EditBalanceHandler interface {
	Handle(EditBalanceParams, interface{}) middleware.Responder
}

// NewEditBalance creates a new http.Handler for the edit balance operation
func NewEditBalance(ctx *middleware.Context, handler EditBalanceHandler) *EditBalance {
	return &EditBalance{Context: ctx, Handler: handler}
}

/*
	EditBalance swagger:route PUT /balance/edit Balance editBalance

EditBalance edit balance API
*/
type EditBalance struct {
	Context *middleware.Context
	Handler EditBalanceHandler
}

func (o *EditBalance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEditBalanceParams()
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

// EditBalanceBody edit balance body
//
// swagger:model EditBalanceBody
type EditBalanceBody struct {

	// data
	Data *models.BalanceAdd `json:"data,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this edit balance body
func (o *EditBalanceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditBalanceBody) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this edit balance body based on the context it is used
func (o *EditBalanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditBalanceBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *EditBalanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditBalanceBody) UnmarshalBinary(b []byte) error {
	var res EditBalanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}