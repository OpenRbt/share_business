// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewAddPermissionsRoleParams creates a new AddPermissionsRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPermissionsRoleParams() *AddPermissionsRoleParams {
	return &AddPermissionsRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPermissionsRoleParamsWithTimeout creates a new AddPermissionsRoleParams object
// with the ability to set a timeout on a request.
func NewAddPermissionsRoleParamsWithTimeout(timeout time.Duration) *AddPermissionsRoleParams {
	return &AddPermissionsRoleParams{
		timeout: timeout,
	}
}

// NewAddPermissionsRoleParamsWithContext creates a new AddPermissionsRoleParams object
// with the ability to set a context for a request.
func NewAddPermissionsRoleParamsWithContext(ctx context.Context) *AddPermissionsRoleParams {
	return &AddPermissionsRoleParams{
		Context: ctx,
	}
}

// NewAddPermissionsRoleParamsWithHTTPClient creates a new AddPermissionsRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPermissionsRoleParamsWithHTTPClient(client *http.Client) *AddPermissionsRoleParams {
	return &AddPermissionsRoleParams{
		HTTPClient: client,
	}
}

/* AddPermissionsRoleParams contains all the parameters to send to the API endpoint
   for the add permissions role operation.

   Typically these are written to a http.Request.
*/
type AddPermissionsRoleParams struct {

	// Body.
	Body AddPermissionsRoleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add permissions role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPermissionsRoleParams) WithDefaults() *AddPermissionsRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add permissions role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPermissionsRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add permissions role params
func (o *AddPermissionsRoleParams) WithTimeout(timeout time.Duration) *AddPermissionsRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add permissions role params
func (o *AddPermissionsRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add permissions role params
func (o *AddPermissionsRoleParams) WithContext(ctx context.Context) *AddPermissionsRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add permissions role params
func (o *AddPermissionsRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add permissions role params
func (o *AddPermissionsRoleParams) WithHTTPClient(client *http.Client) *AddPermissionsRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add permissions role params
func (o *AddPermissionsRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add permissions role params
func (o *AddPermissionsRoleParams) WithBody(body AddPermissionsRoleBody) *AddPermissionsRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add permissions role params
func (o *AddPermissionsRoleParams) SetBody(body AddPermissionsRoleBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddPermissionsRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
