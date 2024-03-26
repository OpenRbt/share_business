// Code generated by go-swagger; DO NOT EDIT.

package server_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// GetServerGroupByIDHandlerFunc turns a function with the right signature into a get server group by Id handler
type GetServerGroupByIDHandlerFunc func(GetServerGroupByIDParams, *app.AdminAuth) GetServerGroupByIDResponder

// Handle executing the request and returning a response
func (fn GetServerGroupByIDHandlerFunc) Handle(params GetServerGroupByIDParams, principal *app.AdminAuth) GetServerGroupByIDResponder {
	return fn(params, principal)
}

// GetServerGroupByIDHandler interface for that can handle valid get server group by Id params
type GetServerGroupByIDHandler interface {
	Handle(GetServerGroupByIDParams, *app.AdminAuth) GetServerGroupByIDResponder
}

// NewGetServerGroupByID creates a new http.Handler for the get server group by Id operation
func NewGetServerGroupByID(ctx *middleware.Context, handler GetServerGroupByIDHandler) *GetServerGroupByID {
	return &GetServerGroupByID{Context: ctx, Handler: handler}
}

/*
	GetServerGroupByID swagger:route GET /server-groups/{groupId} serverGroups getServerGroupById

GetServerGroupByID get server group by Id API
*/
type GetServerGroupByID struct {
	Context *middleware.Context
	Handler GetServerGroupByIDHandler
}

func (o *GetServerGroupByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetServerGroupByIDParams()
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
