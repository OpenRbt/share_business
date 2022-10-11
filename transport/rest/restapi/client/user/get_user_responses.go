// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash-bonus/transport/rest/restapi/models"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*
GetUserOK describes a response with status code 200, with default header values.

OK
*/
type GetUserOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get user o k response has a 2xx status code
func (o *GetUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user o k response has a 3xx status code
func (o *GetUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user o k response has a 4xx status code
func (o *GetUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user o k response has a 5xx status code
func (o *GetUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user o k response a status code equal to that given
func (o *GetUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /user/{id}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) String() string {
	return fmt.Sprintf("[GET /user/{id}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDefault creates a GetUserDefault with default headers values
func NewGetUserDefault(code int) *GetUserDefault {
	return &GetUserDefault{
		_statusCode: code,
	}
}

/*
GetUserDefault describes a response with status code -1, with default header values.

error
*/
type GetUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get user default response
func (o *GetUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get user default response has a 2xx status code
func (o *GetUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get user default response has a 3xx status code
func (o *GetUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get user default response has a 4xx status code
func (o *GetUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get user default response has a 5xx status code
func (o *GetUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get user default response a status code equal to that given
func (o *GetUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetUserDefault) Error() string {
	return fmt.Sprintf("[GET /user/{id}][%d] getUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDefault) String() string {
	return fmt.Sprintf("[GET /user/{id}][%d] getUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}