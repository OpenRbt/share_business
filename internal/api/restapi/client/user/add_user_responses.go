// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash-bonus/internal/api/restapi/models"
)

// AddUserReader is a Reader for the AddUser structure.
type AddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddUserCreated creates a AddUserCreated with default headers values
func NewAddUserCreated() *AddUserCreated {
	return &AddUserCreated{}
}

/* AddUserCreated describes a response with status code 201, with default header values.

Created
*/
type AddUserCreated struct {
	Payload *models.User
}

func (o *AddUserCreated) Error() string {
	return fmt.Sprintf("[POST /user/add][%d] addUserCreated  %+v", 201, o.Payload)
}
func (o *AddUserCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *AddUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserDefault creates a AddUserDefault with default headers values
func NewAddUserDefault(code int) *AddUserDefault {
	return &AddUserDefault{
		_statusCode: code,
	}
}

/* AddUserDefault describes a response with status code -1, with default header values.

error
*/
type AddUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add user default response
func (o *AddUserDefault) Code() int {
	return o._statusCode
}

func (o *AddUserDefault) Error() string {
	return fmt.Sprintf("[POST /user/add][%d] addUser default  %+v", o._statusCode, o.Payload)
}
func (o *AddUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
