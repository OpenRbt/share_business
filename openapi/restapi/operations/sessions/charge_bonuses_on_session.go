// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"washBonus/internal/app"
)

// ChargeBonusesOnSessionHandlerFunc turns a function with the right signature into a charge bonuses on session handler
type ChargeBonusesOnSessionHandlerFunc func(ChargeBonusesOnSessionParams, *app.Auth) ChargeBonusesOnSessionResponder

// Handle executing the request and returning a response
func (fn ChargeBonusesOnSessionHandlerFunc) Handle(params ChargeBonusesOnSessionParams, principal *app.Auth) ChargeBonusesOnSessionResponder {
	return fn(params, principal)
}

// ChargeBonusesOnSessionHandler interface for that can handle valid charge bonuses on session params
type ChargeBonusesOnSessionHandler interface {
	Handle(ChargeBonusesOnSessionParams, *app.Auth) ChargeBonusesOnSessionResponder
}

// NewChargeBonusesOnSession creates a new http.Handler for the charge bonuses on session operation
func NewChargeBonusesOnSession(ctx *middleware.Context, handler ChargeBonusesOnSessionHandler) *ChargeBonusesOnSession {
	return &ChargeBonusesOnSession{Context: ctx, Handler: handler}
}

/*
	ChargeBonusesOnSession swagger:route POST /sessions/{id}/bonuses sessions chargeBonusesOnSession

ChargeBonusesOnSession charge bonuses on session API
*/
type ChargeBonusesOnSession struct {
	Context *middleware.Context
	Handler ChargeBonusesOnSessionHandler
}

func (o *ChargeBonusesOnSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewChargeBonusesOnSessionParams()
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