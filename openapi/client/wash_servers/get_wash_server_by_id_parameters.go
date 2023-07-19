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
)

// NewGetWashServerByIDParams creates a new GetWashServerByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWashServerByIDParams() *GetWashServerByIDParams {
	return &GetWashServerByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWashServerByIDParamsWithTimeout creates a new GetWashServerByIDParams object
// with the ability to set a timeout on a request.
func NewGetWashServerByIDParamsWithTimeout(timeout time.Duration) *GetWashServerByIDParams {
	return &GetWashServerByIDParams{
		timeout: timeout,
	}
}

// NewGetWashServerByIDParamsWithContext creates a new GetWashServerByIDParams object
// with the ability to set a context for a request.
func NewGetWashServerByIDParamsWithContext(ctx context.Context) *GetWashServerByIDParams {
	return &GetWashServerByIDParams{
		Context: ctx,
	}
}

// NewGetWashServerByIDParamsWithHTTPClient creates a new GetWashServerByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWashServerByIDParamsWithHTTPClient(client *http.Client) *GetWashServerByIDParams {
	return &GetWashServerByIDParams{
		HTTPClient: client,
	}
}

/*
GetWashServerByIDParams contains all the parameters to send to the API endpoint

	for the get wash server by Id operation.

	Typically these are written to a http.Request.
*/
type GetWashServerByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wash server by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWashServerByIDParams) WithDefaults() *GetWashServerByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wash server by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWashServerByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wash server by Id params
func (o *GetWashServerByIDParams) WithTimeout(timeout time.Duration) *GetWashServerByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wash server by Id params
func (o *GetWashServerByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wash server by Id params
func (o *GetWashServerByIDParams) WithContext(ctx context.Context) *GetWashServerByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wash server by Id params
func (o *GetWashServerByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wash server by Id params
func (o *GetWashServerByIDParams) WithHTTPClient(client *http.Client) *GetWashServerByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wash server by Id params
func (o *GetWashServerByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get wash server by Id params
func (o *GetWashServerByIDParams) WithID(id string) *GetWashServerByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get wash server by Id params
func (o *GetWashServerByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetWashServerByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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