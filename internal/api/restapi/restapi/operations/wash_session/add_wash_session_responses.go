// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wash-bonus/internal/api/restapi/models"
)

// AddWashSessionCreatedCode is the HTTP code returned for type AddWashSessionCreated
const AddWashSessionCreatedCode int = 201

/*AddWashSessionCreated Created

swagger:response addWashSessionCreated
*/
type AddWashSessionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.WashSession `json:"body,omitempty"`
}

// NewAddWashSessionCreated creates AddWashSessionCreated with default headers values
func NewAddWashSessionCreated() *AddWashSessionCreated {

	return &AddWashSessionCreated{}
}

// WithPayload adds the payload to the add wash session created response
func (o *AddWashSessionCreated) WithPayload(payload *models.WashSession) *AddWashSessionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add wash session created response
func (o *AddWashSessionCreated) SetPayload(payload *models.WashSession) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddWashSessionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddWashSessionDefault error

swagger:response addWashSessionDefault
*/
type AddWashSessionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddWashSessionDefault creates AddWashSessionDefault with default headers values
func NewAddWashSessionDefault(code int) *AddWashSessionDefault {
	if code <= 0 {
		code = 500
	}

	return &AddWashSessionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add wash session default response
func (o *AddWashSessionDefault) WithStatusCode(code int) *AddWashSessionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add wash session default response
func (o *AddWashSessionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add wash session default response
func (o *AddWashSessionDefault) WithPayload(payload *models.Error) *AddWashSessionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add wash session default response
func (o *AddWashSessionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddWashSessionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
