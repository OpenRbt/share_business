// Code generated by go-swagger; DO NOT EDIT.

package wash_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wash-bonus/internal/api/restapi/models"
)

// AddWashServerCreatedCode is the HTTP code returned for type AddWashServerCreated
const AddWashServerCreatedCode int = 201

/*AddWashServerCreated Created

swagger:response addWashServerCreated
*/
type AddWashServerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.WashServer `json:"body,omitempty"`
}

// NewAddWashServerCreated creates AddWashServerCreated with default headers values
func NewAddWashServerCreated() *AddWashServerCreated {

	return &AddWashServerCreated{}
}

// WithPayload adds the payload to the add wash server created response
func (o *AddWashServerCreated) WithPayload(payload *models.WashServer) *AddWashServerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add wash server created response
func (o *AddWashServerCreated) SetPayload(payload *models.WashServer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddWashServerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddWashServerDefault error

swagger:response addWashServerDefault
*/
type AddWashServerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddWashServerDefault creates AddWashServerDefault with default headers values
func NewAddWashServerDefault(code int) *AddWashServerDefault {
	if code <= 0 {
		code = 500
	}

	return &AddWashServerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add wash server default response
func (o *AddWashServerDefault) WithStatusCode(code int) *AddWashServerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add wash server default response
func (o *AddWashServerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add wash server default response
func (o *AddWashServerDefault) WithPayload(payload *models.Error) *AddWashServerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add wash server default response
func (o *AddWashServerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddWashServerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
