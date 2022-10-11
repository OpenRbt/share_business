// Code generated by go-swagger; DO NOT EDIT.

package standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddTestDataHandlerFunc turns a function with the right signature into a add test data handler
type AddTestDataHandlerFunc func(AddTestDataParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTestDataHandlerFunc) Handle(params AddTestDataParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddTestDataHandler interface for that can handle valid add test data params
type AddTestDataHandler interface {
	Handle(AddTestDataParams, interface{}) middleware.Responder
}

// NewAddTestData creates a new http.Handler for the add test data operation
func NewAddTestData(ctx *middleware.Context, handler AddTestDataHandler) *AddTestData {
	return &AddTestData{Context: ctx, Handler: handler}
}

/* AddTestData swagger:route POST /addTestData Standard addTestData

AddTestData add test data API

*/
type AddTestData struct {
	Context *middleware.Context
	Handler AddTestDataHandler
}

func (o *AddTestData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddTestDataParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
