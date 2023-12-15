// Code generated by go-swagger; DO NOT EDIT.

package wallets

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
	"github.com/go-openapi/swag"
)

// NewGetWalletsParams creates a new GetWalletsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWalletsParams() *GetWalletsParams {
	return &GetWalletsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWalletsParamsWithTimeout creates a new GetWalletsParams object
// with the ability to set a timeout on a request.
func NewGetWalletsParamsWithTimeout(timeout time.Duration) *GetWalletsParams {
	return &GetWalletsParams{
		timeout: timeout,
	}
}

// NewGetWalletsParamsWithContext creates a new GetWalletsParams object
// with the ability to set a context for a request.
func NewGetWalletsParamsWithContext(ctx context.Context) *GetWalletsParams {
	return &GetWalletsParams{
		Context: ctx,
	}
}

// NewGetWalletsParamsWithHTTPClient creates a new GetWalletsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWalletsParamsWithHTTPClient(client *http.Client) *GetWalletsParams {
	return &GetWalletsParams{
		HTTPClient: client,
	}
}

/*
GetWalletsParams contains all the parameters to send to the API endpoint

	for the get wallets operation.

	Typically these are written to a http.Request.
*/
type GetWalletsParams struct {

	/* Limit.

	   Maximum number of records to return

	   Format: int64
	   Default: 100
	*/
	Limit *int64

	/* Offset.

	   Number of records to skip for pagination

	   Format: int64
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wallets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWalletsParams) WithDefaults() *GetWalletsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wallets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWalletsParams) SetDefaults() {
	var (
		limitDefault = int64(100)

		offsetDefault = int64(0)
	)

	val := GetWalletsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get wallets params
func (o *GetWalletsParams) WithTimeout(timeout time.Duration) *GetWalletsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wallets params
func (o *GetWalletsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wallets params
func (o *GetWalletsParams) WithContext(ctx context.Context) *GetWalletsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wallets params
func (o *GetWalletsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wallets params
func (o *GetWalletsParams) WithHTTPClient(client *http.Client) *GetWalletsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wallets params
func (o *GetWalletsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get wallets params
func (o *GetWalletsParams) WithLimit(limit *int64) *GetWalletsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get wallets params
func (o *GetWalletsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get wallets params
func (o *GetWalletsParams) WithOffset(offset *int64) *GetWalletsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get wallets params
func (o *GetWalletsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetWalletsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
