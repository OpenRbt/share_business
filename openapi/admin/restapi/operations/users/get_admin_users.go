// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// GetAdminUsersHandlerFunc turns a function with the right signature into a get admin users handler
type GetAdminUsersHandlerFunc func(GetAdminUsersParams, *app.AdminAuth) GetAdminUsersResponder

// Handle executing the request and returning a response
func (fn GetAdminUsersHandlerFunc) Handle(params GetAdminUsersParams, principal *app.AdminAuth) GetAdminUsersResponder {
	return fn(params, principal)
}

// GetAdminUsersHandler interface for that can handle valid get admin users params
type GetAdminUsersHandler interface {
	Handle(GetAdminUsersParams, *app.AdminAuth) GetAdminUsersResponder
}

// NewGetAdminUsers creates a new http.Handler for the get admin users operation
func NewGetAdminUsers(ctx *middleware.Context, handler GetAdminUsersHandler) *GetAdminUsers {
	return &GetAdminUsers{Context: ctx, Handler: handler}
}

/*
	GetAdminUsers swagger:route GET /users users getAdminUsers

GetAdminUsers get admin users API
*/
type GetAdminUsers struct {
	Context *middleware.Context
	Handler GetAdminUsersHandler
}

func (o *GetAdminUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAdminUsersParams()
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
