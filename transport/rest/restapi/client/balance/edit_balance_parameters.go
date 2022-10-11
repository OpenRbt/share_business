// Code generated by go-swagger; DO NOT EDIT.

package balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEditBalanceParams creates a new EditBalanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEditBalanceParams() *EditBalanceParams {
	return &EditBalanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEditBalanceParamsWithTimeout creates a new EditBalanceParams object
// with the ability to set a timeout on a request.
func NewEditBalanceParamsWithTimeout(timeout time.Duration) *EditBalanceParams {
	return &EditBalanceParams{
		timeout: timeout,
	}
}

// NewEditBalanceParamsWithContext creates a new EditBalanceParams object
// with the ability to set a context for a request.
func NewEditBalanceParamsWithContext(ctx context.Context) *EditBalanceParams {
	return &EditBalanceParams{
		Context: ctx,
	}
}

// NewEditBalanceParamsWithHTTPClient creates a new EditBalanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEditBalanceParamsWithHTTPClient(client *http.Client) *EditBalanceParams {
	return &EditBalanceParams{
		HTTPClient: client,
	}
}

/*
EditBalanceParams contains all the parameters to send to the API endpoint

	for the edit balance operation.

	Typically these are written to a http.Request.
*/
type EditBalanceParams struct {

	// Body.
	Body EditBalanceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edit balance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditBalanceParams) WithDefaults() *EditBalanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edit balance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditBalanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edit balance params
func (o *EditBalanceParams) WithTimeout(timeout time.Duration) *EditBalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit balance params
func (o *EditBalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit balance params
func (o *EditBalanceParams) WithContext(ctx context.Context) *EditBalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit balance params
func (o *EditBalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit balance params
func (o *EditBalanceParams) WithHTTPClient(client *http.Client) *EditBalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit balance params
func (o *EditBalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit balance params
func (o *EditBalanceParams) WithBody(body EditBalanceBody) *EditBalanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit balance params
func (o *EditBalanceParams) SetBody(body EditBalanceBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EditBalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
