// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddRoleHandlerFunc turns a function with the right signature into a add role handler
type AddRoleHandlerFunc func(AddRoleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddRoleHandlerFunc) Handle(params AddRoleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddRoleHandler interface for that can handle valid add role params
type AddRoleHandler interface {
	Handle(AddRoleParams, interface{}) middleware.Responder
}

// NewAddRole creates a new http.Handler for the add role operation
func NewAddRole(ctx *middleware.Context, handler AddRoleHandler) *AddRole {
	return &AddRole{Context: ctx, Handler: handler}
}

/* AddRole swagger:route POST /role/add Role addRole

AddRole add role API

*/
type AddRole struct {
	Context *middleware.Context
	Handler AddRoleHandler
}

func (o *AddRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddRoleParams()
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
