// Code generated by go-swagger; DO NOT EDIT.

package wallets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"washbonus/openapi/bonus/models"
)

// GetWalletByOrganizationIDReader is a Reader for the GetWalletByOrganizationID structure.
type GetWalletByOrganizationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletByOrganizationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWalletByOrganizationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWalletByOrganizationIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWalletByOrganizationIDOK creates a GetWalletByOrganizationIDOK with default headers values
func NewGetWalletByOrganizationIDOK() *GetWalletByOrganizationIDOK {
	return &GetWalletByOrganizationIDOK{}
}

/*
GetWalletByOrganizationIDOK describes a response with status code 200, with default header values.

OK
*/
type GetWalletByOrganizationIDOK struct {
	Payload *models.Wallet
}

// IsSuccess returns true when this get wallet by organization Id o k response has a 2xx status code
func (o *GetWalletByOrganizationIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get wallet by organization Id o k response has a 3xx status code
func (o *GetWalletByOrganizationIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get wallet by organization Id o k response has a 4xx status code
func (o *GetWalletByOrganizationIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get wallet by organization Id o k response has a 5xx status code
func (o *GetWalletByOrganizationIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get wallet by organization Id o k response a status code equal to that given
func (o *GetWalletByOrganizationIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get wallet by organization Id o k response
func (o *GetWalletByOrganizationIDOK) Code() int {
	return 200
}

func (o *GetWalletByOrganizationIDOK) Error() string {
	return fmt.Sprintf("[GET /wallets/by-organization/{id}][%d] getWalletByOrganizationIdOK  %+v", 200, o.Payload)
}

func (o *GetWalletByOrganizationIDOK) String() string {
	return fmt.Sprintf("[GET /wallets/by-organization/{id}][%d] getWalletByOrganizationIdOK  %+v", 200, o.Payload)
}

func (o *GetWalletByOrganizationIDOK) GetPayload() *models.Wallet {
	return o.Payload
}

func (o *GetWalletByOrganizationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Wallet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWalletByOrganizationIDDefault creates a GetWalletByOrganizationIDDefault with default headers values
func NewGetWalletByOrganizationIDDefault(code int) *GetWalletByOrganizationIDDefault {
	return &GetWalletByOrganizationIDDefault{
		_statusCode: code,
	}
}

/*
GetWalletByOrganizationIDDefault describes a response with status code -1, with default header values.

Generic error response
*/
type GetWalletByOrganizationIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get wallet by organization Id default response has a 2xx status code
func (o *GetWalletByOrganizationIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get wallet by organization Id default response has a 3xx status code
func (o *GetWalletByOrganizationIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get wallet by organization Id default response has a 4xx status code
func (o *GetWalletByOrganizationIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get wallet by organization Id default response has a 5xx status code
func (o *GetWalletByOrganizationIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get wallet by organization Id default response a status code equal to that given
func (o *GetWalletByOrganizationIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get wallet by organization Id default response
func (o *GetWalletByOrganizationIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWalletByOrganizationIDDefault) Error() string {
	return fmt.Sprintf("[GET /wallets/by-organization/{id}][%d] getWalletByOrganizationId default  %+v", o._statusCode, o.Payload)
}

func (o *GetWalletByOrganizationIDDefault) String() string {
	return fmt.Sprintf("[GET /wallets/by-organization/{id}][%d] getWalletByOrganizationId default  %+v", o._statusCode, o.Payload)
}

func (o *GetWalletByOrganizationIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWalletByOrganizationIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
