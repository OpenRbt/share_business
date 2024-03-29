// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// AssignServerToGroupHandlerFunc turns a function with the right signature into a assign server to group handler
type AssignServerToGroupHandlerFunc func(AssignServerToGroupParams, *app.AdminAuth) AssignServerToGroupResponder

// Handle executing the request and returning a response
func (fn AssignServerToGroupHandlerFunc) Handle(params AssignServerToGroupParams, principal *app.AdminAuth) AssignServerToGroupResponder {
	return fn(params, principal)
}

// AssignServerToGroupHandler interface for that can handle valid assign server to group params
type AssignServerToGroupHandler interface {
	Handle(AssignServerToGroupParams, *app.AdminAuth) AssignServerToGroupResponder
}

// NewAssignServerToGroup creates a new http.Handler for the assign server to group operation
func NewAssignServerToGroup(ctx *middleware.Context, handler AssignServerToGroupHandler) *AssignServerToGroup {
	return &AssignServerToGroup{Context: ctx, Handler: handler}
}

/*
	AssignServerToGroup swagger:route POST /server-groups/{groupId}/wash-servers/{serverId} washServers serverGroup assignServerToGroup

AssignServerToGroup assign server to group API
*/
type AssignServerToGroup struct {
	Context *middleware.Context
	Handler AssignServerToGroupHandler
}

func (o *AssignServerToGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAssignServerToGroupParams()
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
