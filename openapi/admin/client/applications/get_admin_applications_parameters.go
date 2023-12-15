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
	"github.com/go-openapi/swag"
)

// NewGetAdminApplicationsParams creates a new GetAdminApplicationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAdminApplicationsParams() *GetAdminApplicationsParams {
	return &GetAdminApplicationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdminApplicationsParamsWithTimeout creates a new GetAdminApplicationsParams object
// with the ability to set a timeout on a request.
func NewGetAdminApplicationsParamsWithTimeout(timeout time.Duration) *GetAdminApplicationsParams {
	return &GetAdminApplicationsParams{
		timeout: timeout,
	}
}

// NewGetAdminApplicationsParamsWithContext creates a new GetAdminApplicationsParams object
// with the ability to set a context for a request.
func NewGetAdminApplicationsParamsWithContext(ctx context.Context) *GetAdminApplicationsParams {
	return &GetAdminApplicationsParams{
		Context: ctx,
	}
}

// NewGetAdminApplicationsParamsWithHTTPClient creates a new GetAdminApplicationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAdminApplicationsParamsWithHTTPClient(client *http.Client) *GetAdminApplicationsParams {
	return &GetAdminApplicationsParams{
		HTTPClient: client,
	}
}

/*
GetAdminApplicationsParams contains all the parameters to send to the API endpoint

	for the get admin applications operation.

	Typically these are written to a http.Request.
*/
type GetAdminApplicationsParams struct {

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

	// Status.
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get admin applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdminApplicationsParams) WithDefaults() *GetAdminApplicationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get admin applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdminApplicationsParams) SetDefaults() {
	var (
		limitDefault = int64(100)

		offsetDefault = int64(0)
	)

	val := GetAdminApplicationsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get admin applications params
func (o *GetAdminApplicationsParams) WithTimeout(timeout time.Duration) *GetAdminApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get admin applications params
func (o *GetAdminApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get admin applications params
func (o *GetAdminApplicationsParams) WithContext(ctx context.Context) *GetAdminApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get admin applications params
func (o *GetAdminApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get admin applications params
func (o *GetAdminApplicationsParams) WithHTTPClient(client *http.Client) *GetAdminApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get admin applications params
func (o *GetAdminApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get admin applications params
func (o *GetAdminApplicationsParams) WithLimit(limit *int64) *GetAdminApplicationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get admin applications params
func (o *GetAdminApplicationsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get admin applications params
func (o *GetAdminApplicationsParams) WithOffset(offset *int64) *GetAdminApplicationsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get admin applications params
func (o *GetAdminApplicationsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithStatus adds the status to the get admin applications params
func (o *GetAdminApplicationsParams) WithStatus(status *string) *GetAdminApplicationsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get admin applications params
func (o *GetAdminApplicationsParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdminApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
