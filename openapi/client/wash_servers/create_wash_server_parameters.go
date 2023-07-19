// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

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

	"washBonus/openapi/models"
)

// NewCreateWashServerParams creates a new CreateWashServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateWashServerParams() *CreateWashServerParams {
	return &CreateWashServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWashServerParamsWithTimeout creates a new CreateWashServerParams object
// with the ability to set a timeout on a request.
func NewCreateWashServerParamsWithTimeout(timeout time.Duration) *CreateWashServerParams {
	return &CreateWashServerParams{
		timeout: timeout,
	}
}

// NewCreateWashServerParamsWithContext creates a new CreateWashServerParams object
// with the ability to set a context for a request.
func NewCreateWashServerParamsWithContext(ctx context.Context) *CreateWashServerParams {
	return &CreateWashServerParams{
		Context: ctx,
	}
}

// NewCreateWashServerParamsWithHTTPClient creates a new CreateWashServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateWashServerParamsWithHTTPClient(client *http.Client) *CreateWashServerParams {
	return &CreateWashServerParams{
		HTTPClient: client,
	}
}

/*
CreateWashServerParams contains all the parameters to send to the API endpoint

	for the create wash server operation.

	Typically these are written to a http.Request.
*/
type CreateWashServerParams struct {

	// Body.
	Body *models.WashServerCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create wash server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWashServerParams) WithDefaults() *CreateWashServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create wash server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWashServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create wash server params
func (o *CreateWashServerParams) WithTimeout(timeout time.Duration) *CreateWashServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create wash server params
func (o *CreateWashServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create wash server params
func (o *CreateWashServerParams) WithContext(ctx context.Context) *CreateWashServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create wash server params
func (o *CreateWashServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create wash server params
func (o *CreateWashServerParams) WithHTTPClient(client *http.Client) *CreateWashServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create wash server params
func (o *CreateWashServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create wash server params
func (o *CreateWashServerParams) WithBody(body *models.WashServerCreation) *CreateWashServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create wash server params
func (o *CreateWashServerParams) SetBody(body *models.WashServerCreation) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWashServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}