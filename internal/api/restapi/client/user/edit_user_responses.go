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

// EditUserReader is a Reader for the EditUser structure.
type EditUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEditUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEditUserOK creates a EditUserOK with default headers values
func NewEditUserOK() *EditUserOK {
	return &EditUserOK{}
}

/* EditUserOK describes a response with status code 200, with default header values.

OK
*/
type EditUserOK struct {
}

func (o *EditUserOK) Error() string {
	return fmt.Sprintf("[PUT /user/{id}/edit][%d] editUserOK ", 200)
}

func (o *EditUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditUserDefault creates a EditUserDefault with default headers values
func NewEditUserDefault(code int) *EditUserDefault {
	return &EditUserDefault{
		_statusCode: code,
	}
}

/* EditUserDefault describes a response with status code -1, with default header values.

error
*/
type EditUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the edit user default response
func (o *EditUserDefault) Code() int {
	return o._statusCode
}

func (o *EditUserDefault) Error() string {
	return fmt.Sprintf("[PUT /user/{id}/edit][%d] editUser default  %+v", o._statusCode, o.Payload)
}
func (o *EditUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EditUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}