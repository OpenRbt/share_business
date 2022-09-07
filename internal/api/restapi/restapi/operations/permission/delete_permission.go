// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeletePermissionHandlerFunc turns a function with the right signature into a delete permission handler
type DeletePermissionHandlerFunc func(DeletePermissionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePermissionHandlerFunc) Handle(params DeletePermissionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePermissionHandler interface for that can handle valid delete permission params
type DeletePermissionHandler interface {
	Handle(DeletePermissionParams, interface{}) middleware.Responder
}

// NewDeletePermission creates a new http.Handler for the delete permission operation
func NewDeletePermission(ctx *middleware.Context, handler DeletePermissionHandler) *DeletePermission {
	return &DeletePermission{Context: ctx, Handler: handler}
}

/*
	DeletePermission swagger:route DELETE /permission/delete Permission deletePermission

DeletePermission delete permission API
*/
type DeletePermission struct {
	Context *middleware.Context
	Handler DeletePermissionHandler
}

func (o *DeletePermission) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePermissionParams()
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

// DeletePermissionBody delete permission body
//
// swagger:model DeletePermissionBody
type DeletePermissionBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this delete permission body
func (o *DeletePermissionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete permission body based on context it is used
func (o *DeletePermissionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePermissionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePermissionBody) UnmarshalBinary(b []byte) error {
	var res DeletePermissionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
