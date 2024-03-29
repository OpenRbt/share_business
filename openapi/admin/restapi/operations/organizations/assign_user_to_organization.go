// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// AssignUserToOrganizationHandlerFunc turns a function with the right signature into a assign user to organization handler
type AssignUserToOrganizationHandlerFunc func(AssignUserToOrganizationParams, *app.AdminAuth) AssignUserToOrganizationResponder

// Handle executing the request and returning a response
func (fn AssignUserToOrganizationHandlerFunc) Handle(params AssignUserToOrganizationParams, principal *app.AdminAuth) AssignUserToOrganizationResponder {
	return fn(params, principal)
}

// AssignUserToOrganizationHandler interface for that can handle valid assign user to organization params
type AssignUserToOrganizationHandler interface {
	Handle(AssignUserToOrganizationParams, *app.AdminAuth) AssignUserToOrganizationResponder
}

// NewAssignUserToOrganization creates a new http.Handler for the assign user to organization operation
func NewAssignUserToOrganization(ctx *middleware.Context, handler AssignUserToOrganizationHandler) *AssignUserToOrganization {
	return &AssignUserToOrganization{Context: ctx, Handler: handler}
}

/*
	AssignUserToOrganization swagger:route POST /organizations/{organizationId}/users/{userId} organizations assignUserToOrganization

AssignUserToOrganization assign user to organization API
*/
type AssignUserToOrganization struct {
	Context *middleware.Context
	Handler AssignUserToOrganizationHandler
}

func (o *AssignUserToOrganization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAssignUserToOrganizationParams()
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
