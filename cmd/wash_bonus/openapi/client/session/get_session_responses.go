// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash_bonus/openapi/models"
)

// GetSessionReader is a Reader for the GetSession structure.
type GetSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSessionOK creates a GetSessionOK with default headers values
func NewGetSessionOK() *GetSessionOK {
	return &GetSessionOK{}
}

/*
GetSessionOK describes a response with status code 200, with default header values.

OK
*/
type GetSessionOK struct {
	Payload *models.Session
}

// IsSuccess returns true when this get session o k response has a 2xx status code
func (o *GetSessionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get session o k response has a 3xx status code
func (o *GetSessionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session o k response has a 4xx status code
func (o *GetSessionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get session o k response has a 5xx status code
func (o *GetSessionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get session o k response a status code equal to that given
func (o *GetSessionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get session o k response
func (o *GetSessionOK) Code() int {
	return 200
}

func (o *GetSessionOK) Error() string {
	return fmt.Sprintf("[GET /session/{UID}][%d] getSessionOK  %+v", 200, o.Payload)
}

func (o *GetSessionOK) String() string {
	return fmt.Sprintf("[GET /session/{UID}][%d] getSessionOK  %+v", 200, o.Payload)
}

func (o *GetSessionOK) GetPayload() *models.Session {
	return o.Payload
}

func (o *GetSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionNotFound creates a GetSessionNotFound with default headers values
func NewGetSessionNotFound() *GetSessionNotFound {
	return &GetSessionNotFound{}
}

/*
GetSessionNotFound describes a response with status code 404, with default header values.

Profile not exists
*/
type GetSessionNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get session not found response has a 2xx status code
func (o *GetSessionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session not found response has a 3xx status code
func (o *GetSessionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session not found response has a 4xx status code
func (o *GetSessionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get session not found response has a 5xx status code
func (o *GetSessionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get session not found response a status code equal to that given
func (o *GetSessionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get session not found response
func (o *GetSessionNotFound) Code() int {
	return 404
}

func (o *GetSessionNotFound) Error() string {
	return fmt.Sprintf("[GET /session/{UID}][%d] getSessionNotFound  %+v", 404, o.Payload)
}

func (o *GetSessionNotFound) String() string {
	return fmt.Sprintf("[GET /session/{UID}][%d] getSessionNotFound  %+v", 404, o.Payload)
}

func (o *GetSessionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionInternalServerError creates a GetSessionInternalServerError with default headers values
func NewGetSessionInternalServerError() *GetSessionInternalServerError {
	return &GetSessionInternalServerError{}
}

/*
GetSessionInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type GetSessionInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get session internal server error response has a 2xx status code
func (o *GetSessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session internal server error response has a 3xx status code
func (o *GetSessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session internal server error response has a 4xx status code
func (o *GetSessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get session internal server error response has a 5xx status code
func (o *GetSessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get session internal server error response a status code equal to that given
func (o *GetSessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get session internal server error response
func (o *GetSessionInternalServerError) Code() int {
	return 500
}

func (o *GetSessionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /session/{UID}][%d] getSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSessionInternalServerError) String() string {
	return fmt.Sprintf("[GET /session/{UID}][%d] getSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSessionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}