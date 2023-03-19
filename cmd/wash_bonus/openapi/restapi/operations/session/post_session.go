// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"wash_bonus/internal/app"
)

// PostSessionHandlerFunc turns a function with the right signature into a post session handler
type PostSessionHandlerFunc func(PostSessionParams, *app.Auth) PostSessionResponder

// Handle executing the request and returning a response
func (fn PostSessionHandlerFunc) Handle(params PostSessionParams, principal *app.Auth) PostSessionResponder {
	return fn(params, principal)
}

// PostSessionHandler interface for that can handle valid post session params
type PostSessionHandler interface {
	Handle(PostSessionParams, *app.Auth) PostSessionResponder
}

// NewPostSession creates a new http.Handler for the post session operation
func NewPostSession(ctx *middleware.Context, handler PostSessionHandler) *PostSession {
	return &PostSession{Context: ctx, Handler: handler}
}

/*
	PostSession swagger:route POST /session/{UID} session postSession

PostSession post session API
*/
type PostSession struct {
	Context *middleware.Context
	Handler PostSessionHandler
}

func (o *PostSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostSessionParams()
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