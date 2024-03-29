// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// AssignUserToSessionHandlerFunc turns a function with the right signature into a assign user to session handler
type AssignUserToSessionHandlerFunc func(AssignUserToSessionParams, *app.Auth) AssignUserToSessionResponder

// Handle executing the request and returning a response
func (fn AssignUserToSessionHandlerFunc) Handle(params AssignUserToSessionParams, principal *app.Auth) AssignUserToSessionResponder {
	return fn(params, principal)
}

// AssignUserToSessionHandler interface for that can handle valid assign user to session params
type AssignUserToSessionHandler interface {
	Handle(AssignUserToSessionParams, *app.Auth) AssignUserToSessionResponder
}

// NewAssignUserToSession creates a new http.Handler for the assign user to session operation
func NewAssignUserToSession(ctx *middleware.Context, handler AssignUserToSessionHandler) *AssignUserToSession {
	return &AssignUserToSession{Context: ctx, Handler: handler}
}

/*
	AssignUserToSession swagger:route POST /sessions/{sessionId}/assign-user sessions assignUserToSession

AssignUserToSession assign user to session API
*/
type AssignUserToSession struct {
	Context *middleware.Context
	Handler AssignUserToSessionHandler
}

func (o *AssignUserToSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAssignUserToSessionParams()
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
