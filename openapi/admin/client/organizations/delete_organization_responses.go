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

// DeleteOrganizationReader is a Reader for the DeleteOrganization structure.
type DeleteOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrganizationNoContent creates a DeleteOrganizationNoContent with default headers values
func NewDeleteOrganizationNoContent() *DeleteOrganizationNoContent {
	return &DeleteOrganizationNoContent{}
}

/*
DeleteOrganizationNoContent describes a response with status code 204, with default header values.

OK
*/
type DeleteOrganizationNoContent struct {
}

// IsSuccess returns true when this delete organization no content response has a 2xx status code
func (o *DeleteOrganizationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization no content response has a 3xx status code
func (o *DeleteOrganizationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization no content response has a 4xx status code
func (o *DeleteOrganizationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization no content response has a 5xx status code
func (o *DeleteOrganizationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization no content response a status code equal to that given
func (o *DeleteOrganizationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization no content response
func (o *DeleteOrganizationNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}][%d] deleteOrganizationNoContent ", 204)
}

func (o *DeleteOrganizationNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}][%d] deleteOrganizationNoContent ", 204)
}

func (o *DeleteOrganizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationDefault creates a DeleteOrganizationDefault with default headers values
func NewDeleteOrganizationDefault(code int) *DeleteOrganizationDefault {
	return &DeleteOrganizationDefault{
		_statusCode: code,
	}
}

/*
DeleteOrganizationDefault describes a response with status code -1, with default header values.

Generic error response
*/
type DeleteOrganizationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete organization default response has a 2xx status code
func (o *DeleteOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete organization default response has a 3xx status code
func (o *DeleteOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete organization default response has a 4xx status code
func (o *DeleteOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete organization default response has a 5xx status code
func (o *DeleteOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete organization default response a status code equal to that given
func (o *DeleteOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete organization default response
func (o *DeleteOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrganizationDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}][%d] deleteOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationDefault) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}][%d] deleteOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
