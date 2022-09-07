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

	"wash-bonus/internal/api/restapi/models"
)

// NewListPermissionParams creates a new ListPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPermissionParams() *ListPermissionParams {
	return &ListPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPermissionParamsWithTimeout creates a new ListPermissionParams object
// with the ability to set a timeout on a request.
func NewListPermissionParamsWithTimeout(timeout time.Duration) *ListPermissionParams {
	return &ListPermissionParams{
		timeout: timeout,
	}
}

// NewListPermissionParamsWithContext creates a new ListPermissionParams object
// with the ability to set a context for a request.
func NewListPermissionParamsWithContext(ctx context.Context) *ListPermissionParams {
	return &ListPermissionParams{
		Context: ctx,
	}
}

// NewListPermissionParamsWithHTTPClient creates a new ListPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPermissionParamsWithHTTPClient(client *http.Client) *ListPermissionParams {
	return &ListPermissionParams{
		HTTPClient: client,
	}
}

/*
ListPermissionParams contains all the parameters to send to the API endpoint

	for the list permission operation.

	Typically these are written to a http.Request.
*/
type ListPermissionParams struct {

	// Body.
	Body *models.ListParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPermissionParams) WithDefaults() *ListPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list permission params
func (o *ListPermissionParams) WithTimeout(timeout time.Duration) *ListPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list permission params
func (o *ListPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list permission params
func (o *ListPermissionParams) WithContext(ctx context.Context) *ListPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list permission params
func (o *ListPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list permission params
func (o *ListPermissionParams) WithHTTPClient(client *http.Client) *ListPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list permission params
func (o *ListPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list permission params
func (o *ListPermissionParams) WithBody(body *models.ListParams) *ListPermissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list permission params
func (o *ListPermissionParams) SetBody(body *models.ListParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
