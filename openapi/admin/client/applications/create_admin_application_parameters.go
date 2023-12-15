// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// NewCreateAdminApplicationParams creates a new CreateAdminApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAdminApplicationParams() *CreateAdminApplicationParams {
	return &CreateAdminApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAdminApplicationParamsWithTimeout creates a new CreateAdminApplicationParams object
// with the ability to set a timeout on a request.
func NewCreateAdminApplicationParamsWithTimeout(timeout time.Duration) *CreateAdminApplicationParams {
	return &CreateAdminApplicationParams{
		timeout: timeout,
	}
}

// NewCreateAdminApplicationParamsWithContext creates a new CreateAdminApplicationParams object
// with the ability to set a context for a request.
func NewCreateAdminApplicationParamsWithContext(ctx context.Context) *CreateAdminApplicationParams {
	return &CreateAdminApplicationParams{
		Context: ctx,
	}
}

// NewCreateAdminApplicationParamsWithHTTPClient creates a new CreateAdminApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAdminApplicationParamsWithHTTPClient(client *http.Client) *CreateAdminApplicationParams {
	return &CreateAdminApplicationParams{
		HTTPClient: client,
	}
}

/*
CreateAdminApplicationParams contains all the parameters to send to the API endpoint

	for the create admin application operation.

	Typically these are written to a http.Request.
*/
type CreateAdminApplicationParams struct {

	// Body.
	Body CreateAdminApplicationBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create admin application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAdminApplicationParams) WithDefaults() *CreateAdminApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create admin application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAdminApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create admin application params
func (o *CreateAdminApplicationParams) WithTimeout(timeout time.Duration) *CreateAdminApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create admin application params
func (o *CreateAdminApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create admin application params
func (o *CreateAdminApplicationParams) WithContext(ctx context.Context) *CreateAdminApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create admin application params
func (o *CreateAdminApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create admin application params
func (o *CreateAdminApplicationParams) WithHTTPClient(client *http.Client) *CreateAdminApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create admin application params
func (o *CreateAdminApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create admin application params
func (o *CreateAdminApplicationParams) WithBody(body CreateAdminApplicationBody) *CreateAdminApplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create admin application params
func (o *CreateAdminApplicationParams) SetBody(body CreateAdminApplicationBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAdminApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
