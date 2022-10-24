// Code generated by go-swagger; DO NOT EDIT.

package bonus_balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	models2 "wash-bonus/internal/transport/rest/restapi/models"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddBonusBalanceReader is a Reader for the AddBonusBalance structure.
type AddBonusBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddBonusBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddBonusBalanceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddBonusBalanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddBonusBalanceCreated creates a AddBonusBalanceCreated with default headers values
func NewAddBonusBalanceCreated() *AddBonusBalanceCreated {
	return &AddBonusBalanceCreated{}
}

/*
AddBonusBalanceCreated describes a response with status code 201, with default header values.

Created
*/
type AddBonusBalanceCreated struct {
	Payload *models2.Balance
}

// IsSuccess returns true when this add bonus balance created response has a 2xx status code
func (o *AddBonusBalanceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add bonus balance created response has a 3xx status code
func (o *AddBonusBalanceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add bonus balance created response has a 4xx status code
func (o *AddBonusBalanceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add bonus balance created response has a 5xx status code
func (o *AddBonusBalanceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add bonus balance created response a status code equal to that given
func (o *AddBonusBalanceCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddBonusBalanceCreated) Error() string {
	return fmt.Sprintf("[POST /balance/add][%d] addBonusBalanceCreated  %+v", 201, o.Payload)
}

func (o *AddBonusBalanceCreated) String() string {
	return fmt.Sprintf("[POST /balance/add][%d] addBonusBalanceCreated  %+v", 201, o.Payload)
}

func (o *AddBonusBalanceCreated) GetPayload() *models2.Balance {
	return o.Payload
}

func (o *AddBonusBalanceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.Balance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddBonusBalanceDefault creates a AddBonusBalanceDefault with default headers values
func NewAddBonusBalanceDefault(code int) *AddBonusBalanceDefault {
	return &AddBonusBalanceDefault{
		_statusCode: code,
	}
}

/*
AddBonusBalanceDefault describes a response with status code -1, with default header values.

error
*/
type AddBonusBalanceDefault struct {
	_statusCode int

	Payload *models2.Error
}

// Code gets the status code for the add bonus balance default response
func (o *AddBonusBalanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add bonus balance default response has a 2xx status code
func (o *AddBonusBalanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add bonus balance default response has a 3xx status code
func (o *AddBonusBalanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add bonus balance default response has a 4xx status code
func (o *AddBonusBalanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add bonus balance default response has a 5xx status code
func (o *AddBonusBalanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add bonus balance default response a status code equal to that given
func (o *AddBonusBalanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddBonusBalanceDefault) Error() string {
	return fmt.Sprintf("[POST /balance/add][%d] addBonusBalance default  %+v", o._statusCode, o.Payload)
}

func (o *AddBonusBalanceDefault) String() string {
	return fmt.Sprintf("[POST /balance/add][%d] addBonusBalance default  %+v", o._statusCode, o.Payload)
}

func (o *AddBonusBalanceDefault) GetPayload() *models2.Error {
	return o.Payload
}

func (o *AddBonusBalanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}