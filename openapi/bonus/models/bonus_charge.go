// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BonusCharge bonus amount for use in session
//
// swagger:model BonusCharge
type BonusCharge struct {

	// amount
	Amount int64 `json:"amount,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *BonusCharge) UnmarshalJSON(data []byte) error {
	var props struct {

		// amount
		Amount int64 `json:"amount,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Amount = props.Amount
	return nil
}

// Validate validates this bonus charge
func (m *BonusCharge) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bonus charge based on context it is used
func (m *BonusCharge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BonusCharge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BonusCharge) UnmarshalBinary(b []byte) error {
	var res BonusCharge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
