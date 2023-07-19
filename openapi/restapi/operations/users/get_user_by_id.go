// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washBonus/internal/app"
)

// GetUserByIDHandlerFunc turns a function with the right signature into a get user by Id handler
type GetUserByIDHandlerFunc func(GetUserByIDParams, *app.Auth) GetUserByIDResponder

// Handle executing the request and returning a response
func (fn GetUserByIDHandlerFunc) Handle(params GetUserByIDParams, principal *app.Auth) GetUserByIDResponder {
	return fn(params, principal)
}

// GetUserByIDHandler interface for that can handle valid get user by Id params
type GetUserByIDHandler interface {
	Handle(GetUserByIDParams, *app.Auth) GetUserByIDResponder
}

// NewGetUserByID creates a new http.Handler for the get user by Id operation
func NewGetUserByID(ctx *middleware.Context, handler GetUserByIDHandler) *GetUserByID {
	return &GetUserByID{Context: ctx, Handler: handler}
}

/*
	GetUserByID swagger:route GET /users/{id} users getUserById

GetUserByID get user by Id API
*/
type GetUserByID struct {
	Context *middleware.Context
	Handler GetUserByIDHandler
}

func (o *GetUserByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserByIDParams()
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