// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewRemoveUserFromOrganizationParams creates a new RemoveUserFromOrganizationParams object
//
// There are no default values defined in the spec.
func NewRemoveUserFromOrganizationParams() RemoveUserFromOrganizationParams {

	return RemoveUserFromOrganizationParams{}
}

// RemoveUserFromOrganizationParams contains all the bound params for the remove user from organization operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeUserFromOrganization
type RemoveUserFromOrganizationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	OrganizationID strfmt.UUID
	/*
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveUserFromOrganizationParams() beforehand.
func (o *RemoveUserFromOrganizationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rOrganizationID, rhkOrganizationID, _ := route.Params.GetOK("organizationId")
	if err := o.bindOrganizationID(rOrganizationID, rhkOrganizationID, route.Formats); err != nil {
		res = append(res, err)
	}

	rUserID, rhkUserID, _ := route.Params.GetOK("userId")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOrganizationID binds and validates parameter OrganizationID from path.
func (o *RemoveUserFromOrganizationParams) bindOrganizationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("organizationId", "path", "strfmt.UUID", raw)
	}
	o.OrganizationID = *(value.(*strfmt.UUID))

	if err := o.validateOrganizationID(formats); err != nil {
		return err
	}

	return nil
}

// validateOrganizationID carries on validations for parameter OrganizationID
func (o *RemoveUserFromOrganizationParams) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.FormatOf("organizationId", "path", "uuid", o.OrganizationID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *RemoveUserFromOrganizationParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UserID = raw

	return nil
}