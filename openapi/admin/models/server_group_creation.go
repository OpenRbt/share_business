// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServerGroupCreation server group creation
//
// swagger:model ServerGroupCreation
type ServerGroupCreation struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization Id
	// Required: true
	// Format: uuid
	OrganizationID *strfmt.UUID `json:"organizationId"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *ServerGroupCreation) UnmarshalJSON(data []byte) error {
	var props struct {

		// description
		// Required: true
		Description *string `json:"description"`

		// name
		// Required: true
		Name *string `json:"name"`

		// organization Id
		// Required: true
		// Format: uuid
		OrganizationID *strfmt.UUID `json:"organizationId"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Description = props.Description
	m.Name = props.Name
	m.OrganizationID = props.OrganizationID
	return nil
}

// Validate validates this server group creation
func (m *ServerGroupCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerGroupCreation) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServerGroupCreation) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServerGroupCreation) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organizationId", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this server group creation based on context it is used
func (m *ServerGroupCreation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerGroupCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerGroupCreation) UnmarshalBinary(b []byte) error {
	var res ServerGroupCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}