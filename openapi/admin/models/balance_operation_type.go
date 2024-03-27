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

// BalanceOperationType balance operation type
//
// swagger:model BalanceOperationType
type BalanceOperationType string

func NewBalanceOperationType(value BalanceOperationType) *BalanceOperationType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated BalanceOperationType.
func (m BalanceOperationType) Pointer() *BalanceOperationType {
	return &m
}

const (

	// BalanceOperationTypeDeposit captures enum value "deposit"
	BalanceOperationTypeDeposit BalanceOperationType = "deposit"

	// BalanceOperationTypeWithdrawal captures enum value "withdrawal"
	BalanceOperationTypeWithdrawal BalanceOperationType = "withdrawal"
)

// for schema
var balanceOperationTypeEnum []interface{}

func init() {
	var res []BalanceOperationType
	if err := json.Unmarshal([]byte(`["deposit","withdrawal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		balanceOperationTypeEnum = append(balanceOperationTypeEnum, v)
	}
}

func (m BalanceOperationType) validateBalanceOperationTypeEnum(path, location string, value BalanceOperationType) error {
	if err := validate.EnumCase(path, location, value, balanceOperationTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this balance operation type
func (m BalanceOperationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBalanceOperationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this balance operation type based on context it is used
func (m BalanceOperationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
