// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WashServer washServer object
//
// swagger:model washServer
type WashServer struct {

	// created at
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// last update at
	// Format: date-time
	LastUpdateAt *strfmt.DateTime `json:"lastUpdateAt,omitempty"`

	// modified at
	// Format: date-time
	ModifiedAt *strfmt.DateTime `json:"modifiedAt,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this wash server
func (m *WashServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WashServer) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WashServer) validateLastUpdateAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdateAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdateAt", "body", "date-time", m.LastUpdateAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WashServer) validateModifiedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedAt", "body", "date-time", m.ModifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wash server based on context it is used
func (m *WashServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WashServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WashServer) UnmarshalBinary(b []byte) error {
	var res WashServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
