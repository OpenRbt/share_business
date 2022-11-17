// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BalanceAdd bonus model for add and edit methods
//
// swagger:model balanceAdd
type BalanceAdd struct {

	// balance
	Balance string `json:"balance,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this balance add
func (m *BalanceAdd) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this balance add based on context it is used
func (m *BalanceAdd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BalanceAdd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BalanceAdd) UnmarshalBinary(b []byte) error {
	var res BalanceAdd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}