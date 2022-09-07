// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// EditPermissionReader is a Reader for the EditPermission structure.
type EditPermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditPermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditPermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEditPermissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEditPermissionOK creates a EditPermissionOK with default headers values
func NewEditPermissionOK() *EditPermissionOK {
	return &EditPermissionOK{}
}

/*
EditPermissionOK describes a response with status code 200, with default header values.

OK
*/
type EditPermissionOK struct {
}

// IsSuccess returns true when this edit permission o k response has a 2xx status code
func (o *EditPermissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edit permission o k response has a 3xx status code
func (o *EditPermissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit permission o k response has a 4xx status code
func (o *EditPermissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edit permission o k response has a 5xx status code
func (o *EditPermissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edit permission o k response a status code equal to that given
func (o *EditPermissionOK) IsCode(code int) bool {
	return code == 200
}

func (o *EditPermissionOK) Error() string {
	return fmt.Sprintf("[PUT /permission/edit][%d] editPermissionOK ", 200)
}

func (o *EditPermissionOK) String() string {
	return fmt.Sprintf("[PUT /permission/edit][%d] editPermissionOK ", 200)
}

func (o *EditPermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditPermissionDefault creates a EditPermissionDefault with default headers values
func NewEditPermissionDefault(code int) *EditPermissionDefault {
	return &EditPermissionDefault{
		_statusCode: code,
	}
}

/*
EditPermissionDefault describes a response with status code -1, with default header values.

error
*/
type EditPermissionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the edit permission default response
func (o *EditPermissionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edit permission default response has a 2xx status code
func (o *EditPermissionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edit permission default response has a 3xx status code
func (o *EditPermissionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edit permission default response has a 4xx status code
func (o *EditPermissionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edit permission default response has a 5xx status code
func (o *EditPermissionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edit permission default response a status code equal to that given
func (o *EditPermissionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EditPermissionDefault) Error() string {
	return fmt.Sprintf("[PUT /permission/edit][%d] editPermission default  %+v", o._statusCode, o.Payload)
}

func (o *EditPermissionDefault) String() string {
	return fmt.Sprintf("[PUT /permission/edit][%d] editPermission default  %+v", o._statusCode, o.Payload)
}

func (o *EditPermissionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EditPermissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
EditPermissionBody edit permission body
swagger:model EditPermissionBody
*/
type EditPermissionBody struct {

	// data
	Data *models.PermissionAdd `json:"data,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this edit permission body
func (o *EditPermissionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditPermissionBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edit permission body based on the context it is used
func (o *EditPermissionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditPermissionBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EditPermissionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditPermissionBody) UnmarshalBinary(b []byte) error {
	var res EditPermissionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
