// Code generated by go-swagger; DO NOT EDIT.

package bonus_balance

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

// NewAddBonusBalanceParams creates a new AddBonusBalanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddBonusBalanceParams() *AddBonusBalanceParams {
	return &AddBonusBalanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddBonusBalanceParamsWithTimeout creates a new AddBonusBalanceParams object
// with the ability to set a timeout on a request.
func NewAddBonusBalanceParamsWithTimeout(timeout time.Duration) *AddBonusBalanceParams {
	return &AddBonusBalanceParams{
		timeout: timeout,
	}
}

// NewAddBonusBalanceParamsWithContext creates a new AddBonusBalanceParams object
// with the ability to set a context for a request.
func NewAddBonusBalanceParamsWithContext(ctx context.Context) *AddBonusBalanceParams {
	return &AddBonusBalanceParams{
		Context: ctx,
	}
}

// NewAddBonusBalanceParamsWithHTTPClient creates a new AddBonusBalanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddBonusBalanceParamsWithHTTPClient(client *http.Client) *AddBonusBalanceParams {
	return &AddBonusBalanceParams{
		HTTPClient: client,
	}
}

/*
AddBonusBalanceParams contains all the parameters to send to the API endpoint

	for the add bonus balance operation.

	Typically these are written to a http.Request.
*/
type AddBonusBalanceParams struct {

	// Body.
	Body *models.BalanceAdd

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add bonus balance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBonusBalanceParams) WithDefaults() *AddBonusBalanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add bonus balance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBonusBalanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add bonus balance params
func (o *AddBonusBalanceParams) WithTimeout(timeout time.Duration) *AddBonusBalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add bonus balance params
func (o *AddBonusBalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add bonus balance params
func (o *AddBonusBalanceParams) WithContext(ctx context.Context) *AddBonusBalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add bonus balance params
func (o *AddBonusBalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add bonus balance params
func (o *AddBonusBalanceParams) WithHTTPClient(client *http.Client) *AddBonusBalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add bonus balance params
func (o *AddBonusBalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add bonus balance params
func (o *AddBonusBalanceParams) WithBody(body *models.BalanceAdd) *AddBonusBalanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add bonus balance params
func (o *AddBonusBalanceParams) SetBody(body *models.BalanceAdd) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddBonusBalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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