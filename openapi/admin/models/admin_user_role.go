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

// AdminUserRole admin user role
//
// swagger:model AdminUserRole
type AdminUserRole string

func NewAdminUserRole(value AdminUserRole) *AdminUserRole {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AdminUserRole.
func (m AdminUserRole) Pointer() *AdminUserRole {
	return &m
}

const (

	// AdminUserRoleSystemManager captures enum value "systemManager"
	AdminUserRoleSystemManager AdminUserRole = "systemManager"

	// AdminUserRoleAdmin captures enum value "admin"
	AdminUserRoleAdmin AdminUserRole = "admin"

	// AdminUserRoleNoAccess captures enum value "noAccess"
	AdminUserRoleNoAccess AdminUserRole = "noAccess"
)

// for schema
var adminUserRoleEnum []interface{}

func init() {
	var res []AdminUserRole
	if err := json.Unmarshal([]byte(`["systemManager","admin","noAccess"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		adminUserRoleEnum = append(adminUserRoleEnum, v)
	}
}

func (m AdminUserRole) validateAdminUserRoleEnum(path, location string, value AdminUserRole) error {
	if err := validate.EnumCase(path, location, value, adminUserRoleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this admin user role
func (m AdminUserRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAdminUserRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this admin user role based on context it is used
func (m AdminUserRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}