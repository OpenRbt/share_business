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
)

// Session session
//
// swagger:model Session
type Session struct {

	// post balance
	PostBalance int64 `json:"postBalance,omitempty"`

	// post ID
	PostID int64 `json:"postID,omitempty"`

	// wash server
	WashServer *WashServer `json:"washServer,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Session) UnmarshalJSON(data []byte) error {
	var props struct {

		// post balance
		PostBalance int64 `json:"postBalance,omitempty"`

		// post ID
		PostID int64 `json:"postID,omitempty"`

		// wash server
		WashServer *WashServer `json:"washServer,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.PostBalance = props.PostBalance
	m.PostID = props.PostID
	m.WashServer = props.WashServer
	return nil
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWashServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Session) validateWashServer(formats strfmt.Registry) error {
	if swag.IsZero(m.WashServer) { // not required
		return nil
	}

	if m.WashServer != nil {
		if err := m.WashServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("washServer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("washServer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this session based on the context it is used
func (m *Session) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWashServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Session) contextValidateWashServer(ctx context.Context, formats strfmt.Registry) error {

	if m.WashServer != nil {

		if swag.IsZero(m.WashServer) { // not required
			return nil
		}

		if err := m.WashServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("washServer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("washServer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Session) UnmarshalBinary(b []byte) error {
	var res Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
