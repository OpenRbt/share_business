// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// DeleteOrganizationHandlerFunc turns a function with the right signature into a delete organization handler
type DeleteOrganizationHandlerFunc func(DeleteOrganizationParams, *app.AdminAuth) DeleteOrganizationResponder

// Handle executing the request and returning a response
func (fn DeleteOrganizationHandlerFunc) Handle(params DeleteOrganizationParams, principal *app.AdminAuth) DeleteOrganizationResponder {
	return fn(params, principal)
}

// DeleteOrganizationHandler interface for that can handle valid delete organization params
type DeleteOrganizationHandler interface {
	Handle(DeleteOrganizationParams, *app.AdminAuth) DeleteOrganizationResponder
}

// NewDeleteOrganization creates a new http.Handler for the delete organization operation
func NewDeleteOrganization(ctx *middleware.Context, handler DeleteOrganizationHandler) *DeleteOrganization {
	return &DeleteOrganization{Context: ctx, Handler: handler}
}

/*
	DeleteOrganization swagger:route DELETE /organizations/{organizationId} organizations deleteOrganization

DeleteOrganization delete organization API
*/
type DeleteOrganization struct {
	Context *middleware.Context
	Handler DeleteOrganizationHandler
}

func (o *DeleteOrganization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteOrganizationParams()
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
