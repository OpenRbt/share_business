// Code generated by go-swagger; DO NOT EDIT.

package bonus

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

	"wash_bonus/openapi/models"
)

// NewConfirmParams creates a new ConfirmParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfirmParams() *ConfirmParams {
	return &ConfirmParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfirmParamsWithTimeout creates a new ConfirmParams object
// with the ability to set a timeout on a request.
func NewConfirmParamsWithTimeout(timeout time.Duration) *ConfirmParams {
	return &ConfirmParams{
		timeout: timeout,
	}
}

// NewConfirmParamsWithContext creates a new ConfirmParams object
// with the ability to set a context for a request.
func NewConfirmParamsWithContext(ctx context.Context) *ConfirmParams {
	return &ConfirmParams{
		Context: ctx,
	}
}

// NewConfirmParamsWithHTTPClient creates a new ConfirmParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfirmParamsWithHTTPClient(client *http.Client) *ConfirmParams {
	return &ConfirmParams{
		HTTPClient: client,
	}
}

/*
ConfirmParams contains all the parameters to send to the API endpoint

	for the confirm operation.

	Typically these are written to a http.Request.
*/
type ConfirmParams struct {

	// Body.
	Body *models.BonusConsumeConfirm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the confirm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmParams) WithDefaults() *ConfirmParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the confirm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the confirm params
func (o *ConfirmParams) WithTimeout(timeout time.Duration) *ConfirmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the confirm params
func (o *ConfirmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the confirm params
func (o *ConfirmParams) WithContext(ctx context.Context) *ConfirmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the confirm params
func (o *ConfirmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the confirm params
func (o *ConfirmParams) WithHTTPClient(client *http.Client) *ConfirmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the confirm params
func (o *ConfirmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the confirm params
func (o *ConfirmParams) WithBody(body *models.BonusConsumeConfirm) *ConfirmParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the confirm params
func (o *ConfirmParams) SetBody(body *models.BonusConsumeConfirm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ConfirmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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