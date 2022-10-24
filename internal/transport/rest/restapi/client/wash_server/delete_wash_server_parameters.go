// Code generated by go-swagger; DO NOT EDIT.

package wash_server

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

// NewDeleteWashServerParams creates a new DeleteWashServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWashServerParams() *DeleteWashServerParams {
	return &DeleteWashServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWashServerParamsWithTimeout creates a new DeleteWashServerParams object
// with the ability to set a timeout on a request.
func NewDeleteWashServerParamsWithTimeout(timeout time.Duration) *DeleteWashServerParams {
	return &DeleteWashServerParams{
		timeout: timeout,
	}
}

// NewDeleteWashServerParamsWithContext creates a new DeleteWashServerParams object
// with the ability to set a context for a request.
func NewDeleteWashServerParamsWithContext(ctx context.Context) *DeleteWashServerParams {
	return &DeleteWashServerParams{
		Context: ctx,
	}
}

// NewDeleteWashServerParamsWithHTTPClient creates a new DeleteWashServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWashServerParamsWithHTTPClient(client *http.Client) *DeleteWashServerParams {
	return &DeleteWashServerParams{
		HTTPClient: client,
	}
}

/*
DeleteWashServerParams contains all the parameters to send to the API endpoint

	for the delete wash server operation.

	Typically these are written to a http.Request.
*/
type DeleteWashServerParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete wash server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWashServerParams) WithDefaults() *DeleteWashServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete wash server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWashServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete wash server params
func (o *DeleteWashServerParams) WithTimeout(timeout time.Duration) *DeleteWashServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete wash server params
func (o *DeleteWashServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete wash server params
func (o *DeleteWashServerParams) WithContext(ctx context.Context) *DeleteWashServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete wash server params
func (o *DeleteWashServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete wash server params
func (o *DeleteWashServerParams) WithHTTPClient(client *http.Client) *DeleteWashServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete wash server params
func (o *DeleteWashServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete wash server params
func (o *DeleteWashServerParams) WithID(id string) *DeleteWashServerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete wash server params
func (o *DeleteWashServerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWashServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}