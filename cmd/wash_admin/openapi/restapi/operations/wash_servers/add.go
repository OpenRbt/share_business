// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"wash_admin/internal/app"
)

// AddHandlerFunc turns a function with the right signature into a add handler
type AddHandlerFunc func(AddParams, *app.Auth) AddResponder

// Handle executing the request and returning a response
func (fn AddHandlerFunc) Handle(params AddParams, principal *app.Auth) AddResponder {
	return fn(params, principal)
}

// AddHandler interface for that can handle valid add params
type AddHandler interface {
	Handle(AddParams, *app.Auth) AddResponder
}

// NewAdd creates a new http.Handler for the add operation
func NewAdd(ctx *middleware.Context, handler AddHandler) *Add {
	return &Add{Context: ctx, Handler: handler}
}

/*
	Add swagger:route PUT /wash-server/ wash_servers add

Add add API
*/
type Add struct {
	Context *middleware.Context
	Handler AddHandler
}

func (o *Add) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddParams()
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
