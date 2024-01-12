// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ApplicationStatusEnum application status enum
//
// swagger:model ApplicationStatusEnum
type ApplicationStatusEnum string

func NewApplicationStatusEnum(value ApplicationStatusEnum) *ApplicationStatusEnum {
	v := value
	return &v
}

const (

	// ApplicationStatusEnumAccepted captures enum value "accepted"
	ApplicationStatusEnumAccepted ApplicationStatusEnum = "accepted"

	// ApplicationStatusEnumRejected captures enum value "rejected"
	ApplicationStatusEnumRejected ApplicationStatusEnum = "rejected"

	// ApplicationStatusEnumPending captures enum value "pending"
	ApplicationStatusEnumPending ApplicationStatusEnum = "pending"
)

// for schema
var applicationStatusEnumEnum []interface{}

func init() {
	var res []ApplicationStatusEnum
	if err := json.Unmarshal([]byte(`["accepted","rejected","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationStatusEnumEnum = append(applicationStatusEnumEnum, v)
	}
}

func (m ApplicationStatusEnum) validateApplicationStatusEnumEnum(path, location string, value ApplicationStatusEnum) error {
	if err := validate.EnumCase(path, location, value, applicationStatusEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this application status enum
func (m ApplicationStatusEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateApplicationStatusEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this application status enum based on context it is used
func (m ApplicationStatusEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
