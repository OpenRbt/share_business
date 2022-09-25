// Code generated by go-swagger; DO NOT EDIT.

package wash_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddWashServerHandlerFunc turns a function with the right signature into a add wash server handler
type AddWashServerHandlerFunc func(AddWashServerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddWashServerHandlerFunc) Handle(params AddWashServerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddWashServerHandler interface for that can handle valid add wash server params
type AddWashServerHandler interface {
	Handle(AddWashServerParams, interface{}) middleware.Responder
}

// NewAddWashServer creates a new http.Handler for the add wash server operation
func NewAddWashServer(ctx *middleware.Context, handler AddWashServerHandler) *AddWashServer {
	return &AddWashServer{Context: ctx, Handler: handler}
}

/*
	AddWashServer swagger:route POST /washServer/add WashServer addWashServer

AddWashServer add wash server API
*/
type AddWashServer struct {
	Context *middleware.Context
	Handler AddWashServerHandler
}

func (o *AddWashServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddWashServerParams()
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
