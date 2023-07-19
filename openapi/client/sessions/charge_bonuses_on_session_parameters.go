// Code generated by go-swagger; DO NOT EDIT.

package sessions

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

// NewChargeBonusesOnSessionParams creates a new ChargeBonusesOnSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChargeBonusesOnSessionParams() *ChargeBonusesOnSessionParams {
	return &ChargeBonusesOnSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChargeBonusesOnSessionParamsWithTimeout creates a new ChargeBonusesOnSessionParams object
// with the ability to set a timeout on a request.
func NewChargeBonusesOnSessionParamsWithTimeout(timeout time.Duration) *ChargeBonusesOnSessionParams {
	return &ChargeBonusesOnSessionParams{
		timeout: timeout,
	}
}

// NewChargeBonusesOnSessionParamsWithContext creates a new ChargeBonusesOnSessionParams object
// with the ability to set a context for a request.
func NewChargeBonusesOnSessionParamsWithContext(ctx context.Context) *ChargeBonusesOnSessionParams {
	return &ChargeBonusesOnSessionParams{
		Context: ctx,
	}
}

// NewChargeBonusesOnSessionParamsWithHTTPClient creates a new ChargeBonusesOnSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewChargeBonusesOnSessionParamsWithHTTPClient(client *http.Client) *ChargeBonusesOnSessionParams {
	return &ChargeBonusesOnSessionParams{
		HTTPClient: client,
	}
}

/*
ChargeBonusesOnSessionParams contains all the parameters to send to the API endpoint

	for the charge bonuses on session operation.

	Typically these are written to a http.Request.
*/
type ChargeBonusesOnSessionParams struct {

	// Body.
	Body *models.BonusCharge

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the charge bonuses on session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargeBonusesOnSessionParams) WithDefaults() *ChargeBonusesOnSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the charge bonuses on session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargeBonusesOnSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) WithTimeout(timeout time.Duration) *ChargeBonusesOnSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) WithContext(ctx context.Context) *ChargeBonusesOnSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) WithHTTPClient(client *http.Client) *ChargeBonusesOnSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) WithBody(body *models.BonusCharge) *ChargeBonusesOnSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) SetBody(body *models.BonusCharge) {
	o.Body = body
}

// WithID adds the id to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) WithID(id string) *ChargeBonusesOnSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the charge bonuses on session params
func (o *ChargeBonusesOnSessionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChargeBonusesOnSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
