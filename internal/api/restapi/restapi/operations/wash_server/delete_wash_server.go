// Code generated by go-swagger; DO NOT EDIT.

package wash_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteWashServerHandlerFunc turns a function with the right signature into a delete wash server handler
type DeleteWashServerHandlerFunc func(DeleteWashServerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteWashServerHandlerFunc) Handle(params DeleteWashServerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteWashServerHandler interface for that can handle valid delete wash server params
type DeleteWashServerHandler interface {
	Handle(DeleteWashServerParams, interface{}) middleware.Responder
}

// NewDeleteWashServer creates a new http.Handler for the delete wash server operation
func NewDeleteWashServer(ctx *middleware.Context, handler DeleteWashServerHandler) *DeleteWashServer {
	return &DeleteWashServer{Context: ctx, Handler: handler}
}

/* DeleteWashServer swagger:route DELETE /washServer/{id} WashServer deleteWashServer

DeleteWashServer delete wash server API

*/
type DeleteWashServer struct {
	Context *middleware.Context
	Handler DeleteWashServerHandler
}

func (o *DeleteWashServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteWashServerParams()
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
