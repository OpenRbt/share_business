// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// GetPermissionReader is a Reader for the GetPermission structure.
type GetPermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPermissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPermissionOK creates a GetPermissionOK with default headers values
func NewGetPermissionOK() *GetPermissionOK {
	return &GetPermissionOK{}
}

/*
GetPermissionOK describes a response with status code 200, with default header values.

OK
*/
type GetPermissionOK struct {
	Payload *models.Permission
}

// IsSuccess returns true when this get permission o k response has a 2xx status code
func (o *GetPermissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get permission o k response has a 3xx status code
func (o *GetPermissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permission o k response has a 4xx status code
func (o *GetPermissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permission o k response has a 5xx status code
func (o *GetPermissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get permission o k response a status code equal to that given
func (o *GetPermissionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPermissionOK) Error() string {
	return fmt.Sprintf("[POST /permission/get][%d] getPermissionOK  %+v", 200, o.Payload)
}

func (o *GetPermissionOK) String() string {
	return fmt.Sprintf("[POST /permission/get][%d] getPermissionOK  %+v", 200, o.Payload)
}

func (o *GetPermissionOK) GetPayload() *models.Permission {
	return o.Payload
}

func (o *GetPermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionDefault creates a GetPermissionDefault with default headers values
func NewGetPermissionDefault(code int) *GetPermissionDefault {
	return &GetPermissionDefault{
		_statusCode: code,
	}
}

/*
GetPermissionDefault describes a response with status code -1, with default header values.

error
*/
type GetPermissionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get permission default response
func (o *GetPermissionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get permission default response has a 2xx status code
func (o *GetPermissionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get permission default response has a 3xx status code
func (o *GetPermissionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get permission default response has a 4xx status code
func (o *GetPermissionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get permission default response has a 5xx status code
func (o *GetPermissionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get permission default response a status code equal to that given
func (o *GetPermissionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPermissionDefault) Error() string {
	return fmt.Sprintf("[POST /permission/get][%d] getPermission default  %+v", o._statusCode, o.Payload)
}

func (o *GetPermissionDefault) String() string {
	return fmt.Sprintf("[POST /permission/get][%d] getPermission default  %+v", o._statusCode, o.Payload)
}

func (o *GetPermissionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPermissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetPermissionBody get permission body
swagger:model GetPermissionBody
*/
type GetPermissionBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this get permission body
func (o *GetPermissionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get permission body based on context it is used
func (o *GetPermissionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPermissionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPermissionBody) UnmarshalBinary(b []byte) error {
	var res GetPermissionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
