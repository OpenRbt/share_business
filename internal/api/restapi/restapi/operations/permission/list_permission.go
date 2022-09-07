// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// ListPermissionHandlerFunc turns a function with the right signature into a list permission handler
type ListPermissionHandlerFunc func(ListPermissionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPermissionHandlerFunc) Handle(params ListPermissionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListPermissionHandler interface for that can handle valid list permission params
type ListPermissionHandler interface {
	Handle(ListPermissionParams, interface{}) middleware.Responder
}

// NewListPermission creates a new http.Handler for the list permission operation
func NewListPermission(ctx *middleware.Context, handler ListPermissionHandler) *ListPermission {
	return &ListPermission{Context: ctx, Handler: handler}
}

/*
	ListPermission swagger:route POST /permission/list Permission listPermission

ListPermission list permission API
*/
type ListPermission struct {
	Context *middleware.Context
	Handler ListPermissionHandler
}

func (o *ListPermission) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListPermissionParams()
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

// ListPermissionOKBody list permission o k body
//
// swagger:model ListPermissionOKBody
type ListPermissionOKBody struct {

	// items
	Items []*models.Permission `json:"items"`

	// warnings
	Warnings []string `json:"warnings"`
}

// Validate validates this list permission o k body
func (o *ListPermissionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPermissionOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listPermissionOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listPermissionOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list permission o k body based on the context it is used
func (o *ListPermissionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPermissionOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listPermissionOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listPermissionOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPermissionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPermissionOKBody) UnmarshalBinary(b []byte) error {
	var res ListPermissionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
