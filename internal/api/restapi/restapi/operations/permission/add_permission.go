// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddPermissionHandlerFunc turns a function with the right signature into a add permission handler
type AddPermissionHandlerFunc func(AddPermissionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPermissionHandlerFunc) Handle(params AddPermissionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddPermissionHandler interface for that can handle valid add permission params
type AddPermissionHandler interface {
	Handle(AddPermissionParams, interface{}) middleware.Responder
}

// NewAddPermission creates a new http.Handler for the add permission operation
func NewAddPermission(ctx *middleware.Context, handler AddPermissionHandler) *AddPermission {
	return &AddPermission{Context: ctx, Handler: handler}
}

/*
	AddPermission swagger:route POST /permission/add Permission addPermission

AddPermission add permission API
*/
type AddPermission struct {
	Context *middleware.Context
	Handler AddPermissionHandler
}

func (o *AddPermission) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddPermissionParams()
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
