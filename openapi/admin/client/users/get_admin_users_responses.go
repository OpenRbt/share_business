// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"washbonus/openapi/admin/models"
)

// GetAdminUsersReader is a Reader for the GetAdminUsers structure.
type GetAdminUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAdminUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdminUsersOK creates a GetAdminUsersOK with default headers values
func NewGetAdminUsersOK() *GetAdminUsersOK {
	return &GetAdminUsersOK{}
}

/*
GetAdminUsersOK describes a response with status code 200, with default header values.

OK
*/
type GetAdminUsersOK struct {
	Payload []*models.AdminUser
}

// IsSuccess returns true when this get admin users o k response has a 2xx status code
func (o *GetAdminUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get admin users o k response has a 3xx status code
func (o *GetAdminUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get admin users o k response has a 4xx status code
func (o *GetAdminUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get admin users o k response has a 5xx status code
func (o *GetAdminUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get admin users o k response a status code equal to that given
func (o *GetAdminUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get admin users o k response
func (o *GetAdminUsersOK) Code() int {
	return 200
}

func (o *GetAdminUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] getAdminUsersOK  %+v", 200, o.Payload)
}

func (o *GetAdminUsersOK) String() string {
	return fmt.Sprintf("[GET /users][%d] getAdminUsersOK  %+v", 200, o.Payload)
}

func (o *GetAdminUsersOK) GetPayload() []*models.AdminUser {
	return o.Payload
}

func (o *GetAdminUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminUsersDefault creates a GetAdminUsersDefault with default headers values
func NewGetAdminUsersDefault(code int) *GetAdminUsersDefault {
	return &GetAdminUsersDefault{
		_statusCode: code,
	}
}

/*
GetAdminUsersDefault describes a response with status code -1, with default header values.

Generic error response
*/
type GetAdminUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get admin users default response has a 2xx status code
func (o *GetAdminUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get admin users default response has a 3xx status code
func (o *GetAdminUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get admin users default response has a 4xx status code
func (o *GetAdminUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get admin users default response has a 5xx status code
func (o *GetAdminUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get admin users default response a status code equal to that given
func (o *GetAdminUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get admin users default response
func (o *GetAdminUsersDefault) Code() int {
	return o._statusCode
}

func (o *GetAdminUsersDefault) Error() string {
	return fmt.Sprintf("[GET /users][%d] getAdminUsers default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminUsersDefault) String() string {
	return fmt.Sprintf("[GET /users][%d] getAdminUsers default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdminUsersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAdminUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
