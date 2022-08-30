// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash-bonus/internal/api/restapi/models"
)

// AddRoleReader is a Reader for the AddRole structure.
type AddRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddRoleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRoleCreated creates a AddRoleCreated with default headers values
func NewAddRoleCreated() *AddRoleCreated {
	return &AddRoleCreated{}
}

/* AddRoleCreated describes a response with status code 201, with default header values.

Created
*/
type AddRoleCreated struct {
	Payload *models.Role
}

func (o *AddRoleCreated) Error() string {
	return fmt.Sprintf("[POST /role/add][%d] addRoleCreated  %+v", 201, o.Payload)
}
func (o *AddRoleCreated) GetPayload() *models.Role {
	return o.Payload
}

func (o *AddRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRoleDefault creates a AddRoleDefault with default headers values
func NewAddRoleDefault(code int) *AddRoleDefault {
	return &AddRoleDefault{
		_statusCode: code,
	}
}

/* AddRoleDefault describes a response with status code -1, with default header values.

error
*/
type AddRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add role default response
func (o *AddRoleDefault) Code() int {
	return o._statusCode
}

func (o *AddRoleDefault) Error() string {
	return fmt.Sprintf("[POST /role/add][%d] addRole default  %+v", o._statusCode, o.Payload)
}
func (o *AddRoleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
