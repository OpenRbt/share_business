// Code generated by go-swagger; DO NOT EDIT.

package server_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// DeleteServerGroupHandlerFunc turns a function with the right signature into a delete server group handler
type DeleteServerGroupHandlerFunc func(DeleteServerGroupParams, *app.AdminAuth) DeleteServerGroupResponder

// Handle executing the request and returning a response
func (fn DeleteServerGroupHandlerFunc) Handle(params DeleteServerGroupParams, principal *app.AdminAuth) DeleteServerGroupResponder {
	return fn(params, principal)
}

// DeleteServerGroupHandler interface for that can handle valid delete server group params
type DeleteServerGroupHandler interface {
	Handle(DeleteServerGroupParams, *app.AdminAuth) DeleteServerGroupResponder
}

// NewDeleteServerGroup creates a new http.Handler for the delete server group operation
func NewDeleteServerGroup(ctx *middleware.Context, handler DeleteServerGroupHandler) *DeleteServerGroup {
	return &DeleteServerGroup{Context: ctx, Handler: handler}
}

/* DeleteServerGroup swagger:route DELETE /server-groups/{groupId} serverGroups deleteServerGroup

DeleteServerGroup delete server group API

*/
type DeleteServerGroup struct {
	Context *middleware.Context
	Handler DeleteServerGroupHandler
}

func (o *DeleteServerGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteServerGroupParams()
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
