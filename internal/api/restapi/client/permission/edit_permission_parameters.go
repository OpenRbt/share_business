// Code generated by go-swagger; DO NOT EDIT.

package permission

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

// NewEditPermissionParams creates a new EditPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEditPermissionParams() *EditPermissionParams {
	return &EditPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEditPermissionParamsWithTimeout creates a new EditPermissionParams object
// with the ability to set a timeout on a request.
func NewEditPermissionParamsWithTimeout(timeout time.Duration) *EditPermissionParams {
	return &EditPermissionParams{
		timeout: timeout,
	}
}

// NewEditPermissionParamsWithContext creates a new EditPermissionParams object
// with the ability to set a context for a request.
func NewEditPermissionParamsWithContext(ctx context.Context) *EditPermissionParams {
	return &EditPermissionParams{
		Context: ctx,
	}
}

// NewEditPermissionParamsWithHTTPClient creates a new EditPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewEditPermissionParamsWithHTTPClient(client *http.Client) *EditPermissionParams {
	return &EditPermissionParams{
		HTTPClient: client,
	}
}

/*
EditPermissionParams contains all the parameters to send to the API endpoint

	for the edit permission operation.

	Typically these are written to a http.Request.
*/
type EditPermissionParams struct {

	// Body.
	Body EditPermissionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edit permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditPermissionParams) WithDefaults() *EditPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edit permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edit permission params
func (o *EditPermissionParams) WithTimeout(timeout time.Duration) *EditPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit permission params
func (o *EditPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit permission params
func (o *EditPermissionParams) WithContext(ctx context.Context) *EditPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit permission params
func (o *EditPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit permission params
func (o *EditPermissionParams) WithHTTPClient(client *http.Client) *EditPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit permission params
func (o *EditPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit permission params
func (o *EditPermissionParams) WithBody(body EditPermissionBody) *EditPermissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit permission params
func (o *EditPermissionParams) SetBody(body EditPermissionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EditPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
