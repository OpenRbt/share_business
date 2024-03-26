// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// GetWashServersHandlerFunc turns a function with the right signature into a get wash servers handler
type GetWashServersHandlerFunc func(GetWashServersParams, *app.AdminAuth) GetWashServersResponder

// Handle executing the request and returning a response
func (fn GetWashServersHandlerFunc) Handle(params GetWashServersParams, principal *app.AdminAuth) GetWashServersResponder {
	return fn(params, principal)
}

// GetWashServersHandler interface for that can handle valid get wash servers params
type GetWashServersHandler interface {
	Handle(GetWashServersParams, *app.AdminAuth) GetWashServersResponder
}

// NewGetWashServers creates a new http.Handler for the get wash servers operation
func NewGetWashServers(ctx *middleware.Context, handler GetWashServersHandler) *GetWashServers {
	return &GetWashServers{Context: ctx, Handler: handler}
}

/*
	GetWashServers swagger:route GET /wash-servers/ washServers getWashServers

GetWashServers get wash servers API
*/
type GetWashServers struct {
	Context *middleware.Context
	Handler GetWashServersHandler
}

func (o *GetWashServers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWashServersParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *app.AdminAuth
	if uprinc != nil {
		principal = uprinc.(*app.AdminAuth) // this is really a app.AdminAuth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
