// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washBonus/internal/app"
)

// UpdateSettingForOrganizationHandlerFunc turns a function with the right signature into a update setting for organization handler
type UpdateSettingForOrganizationHandlerFunc func(UpdateSettingForOrganizationParams, *app.Auth) UpdateSettingForOrganizationResponder

// Handle executing the request and returning a response
func (fn UpdateSettingForOrganizationHandlerFunc) Handle(params UpdateSettingForOrganizationParams, principal *app.Auth) UpdateSettingForOrganizationResponder {
	return fn(params, principal)
}

// UpdateSettingForOrganizationHandler interface for that can handle valid update setting for organization params
type UpdateSettingForOrganizationHandler interface {
	Handle(UpdateSettingForOrganizationParams, *app.Auth) UpdateSettingForOrganizationResponder
}

// NewUpdateSettingForOrganization creates a new http.Handler for the update setting for organization operation
func NewUpdateSettingForOrganization(ctx *middleware.Context, handler UpdateSettingForOrganizationHandler) *UpdateSettingForOrganization {
	return &UpdateSettingForOrganization{Context: ctx, Handler: handler}
}

/*
	UpdateSettingForOrganization swagger:route PATCH /organizations/{id}/settings organizations updateSettingForOrganization

UpdateSettingForOrganization update setting for organization API
*/
type UpdateSettingForOrganization struct {
	Context *middleware.Context
	Handler UpdateSettingForOrganizationHandler
}

func (o *UpdateSettingForOrganization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateSettingForOrganizationParams()
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