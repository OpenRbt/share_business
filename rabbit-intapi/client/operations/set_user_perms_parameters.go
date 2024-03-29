// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"washbonus/rabbit-intapi/models"
)

// NewSetUserPermsParams creates a new SetUserPermsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetUserPermsParams() *SetUserPermsParams {
	return &SetUserPermsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetUserPermsParamsWithTimeout creates a new SetUserPermsParams object
// with the ability to set a timeout on a request.
func NewSetUserPermsParamsWithTimeout(timeout time.Duration) *SetUserPermsParams {
	return &SetUserPermsParams{
		timeout: timeout,
	}
}

// NewSetUserPermsParamsWithContext creates a new SetUserPermsParams object
// with the ability to set a context for a request.
func NewSetUserPermsParamsWithContext(ctx context.Context) *SetUserPermsParams {
	return &SetUserPermsParams{
		Context: ctx,
	}
}

// NewSetUserPermsParamsWithHTTPClient creates a new SetUserPermsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetUserPermsParamsWithHTTPClient(client *http.Client) *SetUserPermsParams {
	return &SetUserPermsParams{
		HTTPClient: client,
	}
}

/*
SetUserPermsParams contains all the parameters to send to the API endpoint

	for the set user perms operation.

	Typically these are written to a http.Request.
*/
type SetUserPermsParams struct {

	// Body.
	Body *models.ManagePermissions

	// UserID.
	UserID string

	// Vhost.
	Vhost string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set user perms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUserPermsParams) WithDefaults() *SetUserPermsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set user perms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUserPermsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set user perms params
func (o *SetUserPermsParams) WithTimeout(timeout time.Duration) *SetUserPermsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set user perms params
func (o *SetUserPermsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set user perms params
func (o *SetUserPermsParams) WithContext(ctx context.Context) *SetUserPermsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set user perms params
func (o *SetUserPermsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set user perms params
func (o *SetUserPermsParams) WithHTTPClient(client *http.Client) *SetUserPermsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set user perms params
func (o *SetUserPermsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set user perms params
func (o *SetUserPermsParams) WithBody(body *models.ManagePermissions) *SetUserPermsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set user perms params
func (o *SetUserPermsParams) SetBody(body *models.ManagePermissions) {
	o.Body = body
}

// WithUserID adds the userID to the set user perms params
func (o *SetUserPermsParams) WithUserID(userID string) *SetUserPermsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set user perms params
func (o *SetUserPermsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithVhost adds the vhost to the set user perms params
func (o *SetUserPermsParams) WithVhost(vhost string) *SetUserPermsParams {
	o.SetVhost(vhost)
	return o
}

// SetVhost adds the vhost to the set user perms params
func (o *SetUserPermsParams) SetVhost(vhost string) {
	o.Vhost = vhost
}

// WriteToRequest writes these params to a swagger request
func (o *SetUserPermsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	// path param vhost
	if err := r.SetPathParam("vhost", o.Vhost); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
