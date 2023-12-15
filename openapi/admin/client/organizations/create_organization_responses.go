// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"washbonus/openapi/admin/models"
)

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrganizationOK creates a CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {
	return &CreateOrganizationOK{}
}

/*
CreateOrganizationOK describes a response with status code 200, with default header values.

Successfull created
*/
type CreateOrganizationOK struct {
	Payload *models.Organization
}

// IsSuccess returns true when this create organization o k response has a 2xx status code
func (o *CreateOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization o k response has a 3xx status code
func (o *CreateOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization o k response has a 4xx status code
func (o *CreateOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization o k response has a 5xx status code
func (o *CreateOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization o k response a status code equal to that given
func (o *CreateOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create organization o k response
func (o *CreateOrganizationOK) Code() int {
	return 200
}

func (o *CreateOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) String() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) GetPayload() *models.Organization {
	return o.Payload
}

func (o *CreateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationDefault creates a CreateOrganizationDefault with default headers values
func NewCreateOrganizationDefault(code int) *CreateOrganizationDefault {
	return &CreateOrganizationDefault{
		_statusCode: code,
	}
}

/*
CreateOrganizationDefault describes a response with status code -1, with default header values.

Generic error response
*/
type CreateOrganizationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create organization default response has a 2xx status code
func (o *CreateOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create organization default response has a 3xx status code
func (o *CreateOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create organization default response has a 4xx status code
func (o *CreateOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create organization default response has a 5xx status code
func (o *CreateOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create organization default response a status code equal to that given
func (o *CreateOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create organization default response
func (o *CreateOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrganizationDefault) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrganizationDefault) String() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrganizationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
