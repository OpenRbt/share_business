// Code generated by go-swagger; DO NOT EDIT.

package wallets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washbonus/internal/app"
)

// GetWalletsHandlerFunc turns a function with the right signature into a get wallets handler
type GetWalletsHandlerFunc func(GetWalletsParams, *app.Auth) GetWalletsResponder

// Handle executing the request and returning a response
func (fn GetWalletsHandlerFunc) Handle(params GetWalletsParams, principal *app.Auth) GetWalletsResponder {
	return fn(params, principal)
}

// GetWalletsHandler interface for that can handle valid get wallets params
type GetWalletsHandler interface {
	Handle(GetWalletsParams, *app.Auth) GetWalletsResponder
}

// NewGetWallets creates a new http.Handler for the get wallets operation
func NewGetWallets(ctx *middleware.Context, handler GetWalletsHandler) *GetWallets {
	return &GetWallets{Context: ctx, Handler: handler}
}

/*
	GetWallets swagger:route GET /wallets wallets getWallets

GetWallets get wallets API
*/
type GetWallets struct {
	Context *middleware.Context
	Handler GetWalletsHandler
}

func (o *GetWallets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWalletsParams()
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
