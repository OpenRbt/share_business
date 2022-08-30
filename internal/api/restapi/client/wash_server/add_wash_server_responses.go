// Code generated by go-swagger; DO NOT EDIT.

package wash_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash-bonus/internal/api/restapi/models"
)

// AddWashServerReader is a Reader for the AddWashServer structure.
type AddWashServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWashServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddWashServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddWashServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddWashServerCreated creates a AddWashServerCreated with default headers values
func NewAddWashServerCreated() *AddWashServerCreated {
	return &AddWashServerCreated{}
}

/* AddWashServerCreated describes a response with status code 201, with default header values.

Created
*/
type AddWashServerCreated struct {
	Payload *models.WashServer
}

func (o *AddWashServerCreated) Error() string {
	return fmt.Sprintf("[POST /washServer/add][%d] addWashServerCreated  %+v", 201, o.Payload)
}
func (o *AddWashServerCreated) GetPayload() *models.WashServer {
	return o.Payload
}

func (o *AddWashServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WashServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddWashServerDefault creates a AddWashServerDefault with default headers values
func NewAddWashServerDefault(code int) *AddWashServerDefault {
	return &AddWashServerDefault{
		_statusCode: code,
	}
}

/* AddWashServerDefault describes a response with status code -1, with default header values.

error
*/
type AddWashServerDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add wash server default response
func (o *AddWashServerDefault) Code() int {
	return o._statusCode
}

func (o *AddWashServerDefault) Error() string {
	return fmt.Sprintf("[POST /washServer/add][%d] addWashServer default  %+v", o._statusCode, o.Payload)
}
func (o *AddWashServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddWashServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
