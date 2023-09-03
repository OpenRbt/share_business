// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washBonus/internal/app"
)

// GetSettingsForOrganizationHandlerFunc turns a function with the right signature into a get settings for organization handler
type GetSettingsForOrganizationHandlerFunc func(GetSettingsForOrganizationParams, *app.Auth) GetSettingsForOrganizationResponder

// Handle executing the request and returning a response
func (fn GetSettingsForOrganizationHandlerFunc) Handle(params GetSettingsForOrganizationParams, principal *app.Auth) GetSettingsForOrganizationResponder {
	return fn(params, principal)
}

// GetSettingsForOrganizationHandler interface for that can handle valid get settings for organization params
type GetSettingsForOrganizationHandler interface {
	Handle(GetSettingsForOrganizationParams, *app.Auth) GetSettingsForOrganizationResponder
}

// NewGetSettingsForOrganization creates a new http.Handler for the get settings for organization operation
func NewGetSettingsForOrganization(ctx *middleware.Context, handler GetSettingsForOrganizationHandler) *GetSettingsForOrganization {
	return &GetSettingsForOrganization{Context: ctx, Handler: handler}
}

/*
	GetSettingsForOrganization swagger:route GET /organizations/{id}/settings organizations getSettingsForOrganization

GetSettingsForOrganization get settings for organization API
*/
type GetSettingsForOrganization struct {
	Context *middleware.Context
	Handler GetSettingsForOrganizationHandler
}

func (o *GetSettingsForOrganization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSettingsForOrganizationParams()
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
