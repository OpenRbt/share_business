// Code generated by go-swagger; DO NOT EDIT.

package balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wash-bonus/internal/transport/rest/models"
)

// AddBalanceCreatedCode is the HTTP code returned for type AddBalanceCreated
const AddBalanceCreatedCode int = 201

/*
AddBalanceCreated Created

swagger:response addBalanceCreated
*/
type AddBalanceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Balance `json:"body,omitempty"`
}

// NewAddBalanceCreated creates AddBalanceCreated with default headers values
func NewAddBalanceCreated() *AddBalanceCreated {

	return &AddBalanceCreated{}
}

// WithPayload adds the payload to the add balance created response
func (o *AddBalanceCreated) WithPayload(payload *models.Balance) *AddBalanceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add balance created response
func (o *AddBalanceCreated) SetPayload(payload *models.Balance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBalanceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
AddBalanceDefault error

swagger:response addBalanceDefault
*/
type AddBalanceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddBalanceDefault creates AddBalanceDefault with default headers values
func NewAddBalanceDefault(code int) *AddBalanceDefault {
	if code <= 0 {
		code = 500
	}

	return &AddBalanceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add balance default response
func (o *AddBalanceDefault) WithStatusCode(code int) *AddBalanceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add balance default response
func (o *AddBalanceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add balance default response
func (o *AddBalanceDefault) WithPayload(payload *models.Error) *AddBalanceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add balance default response
func (o *AddBalanceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBalanceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}