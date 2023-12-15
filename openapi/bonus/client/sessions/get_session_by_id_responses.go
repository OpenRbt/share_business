// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"washbonus/openapi/bonus/models"
)

// GetSessionByIDReader is a Reader for the GetSessionByID structure.
type GetSessionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSessionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSessionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSessionByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSessionByIDOK creates a GetSessionByIDOK with default headers values
func NewGetSessionByIDOK() *GetSessionByIDOK {
	return &GetSessionByIDOK{}
}

/*
GetSessionByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetSessionByIDOK struct {
	Payload *models.Session
}

// IsSuccess returns true when this get session by Id o k response has a 2xx status code
func (o *GetSessionByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get session by Id o k response has a 3xx status code
func (o *GetSessionByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session by Id o k response has a 4xx status code
func (o *GetSessionByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get session by Id o k response has a 5xx status code
func (o *GetSessionByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get session by Id o k response a status code equal to that given
func (o *GetSessionByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get session by Id o k response
func (o *GetSessionByIDOK) Code() int {
	return 200
}

func (o *GetSessionByIDOK) Error() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdOK  %+v", 200, o.Payload)
}

func (o *GetSessionByIDOK) String() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdOK  %+v", 200, o.Payload)
}

func (o *GetSessionByIDOK) GetPayload() *models.Session {
	return o.Payload
}

func (o *GetSessionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionByIDDefault creates a GetSessionByIDDefault with default headers values
func NewGetSessionByIDDefault(code int) *GetSessionByIDDefault {
	return &GetSessionByIDDefault{
		_statusCode: code,
	}
}

/*
GetSessionByIDDefault describes a response with status code -1, with default header values.

Generic error response
*/
type GetSessionByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get session by Id default response has a 2xx status code
func (o *GetSessionByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get session by Id default response has a 3xx status code
func (o *GetSessionByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get session by Id default response has a 4xx status code
func (o *GetSessionByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get session by Id default response has a 5xx status code
func (o *GetSessionByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get session by Id default response a status code equal to that given
func (o *GetSessionByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get session by Id default response
func (o *GetSessionByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSessionByIDDefault) Error() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionById default  %+v", o._statusCode, o.Payload)
}

func (o *GetSessionByIDDefault) String() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionById default  %+v", o._statusCode, o.Payload)
}

func (o *GetSessionByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSessionByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
