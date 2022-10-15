// Code generated by go-swagger; DO NOT EDIT.

package standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wash-bonus/internal/api/restapi/models"
)

// AddTestDataOKCode is the HTTP code returned for type AddTestDataOK
const AddTestDataOKCode int = 200

/*AddTestDataOK OK

swagger:response addTestDataOK
*/
type AddTestDataOK struct {
}

// NewAddTestDataOK creates AddTestDataOK with default headers values
func NewAddTestDataOK() *AddTestDataOK {

	return &AddTestDataOK{}
}

// WriteResponse to the client
func (o *AddTestDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*AddTestDataDefault error

swagger:response addTestDataDefault
*/
type AddTestDataDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddTestDataDefault creates AddTestDataDefault with default headers values
func NewAddTestDataDefault(code int) *AddTestDataDefault {
	if code <= 0 {
		code = 500
	}

	return &AddTestDataDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add test data default response
func (o *AddTestDataDefault) WithStatusCode(code int) *AddTestDataDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add test data default response
func (o *AddTestDataDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add test data default response
func (o *AddTestDataDefault) WithPayload(payload *models.Error) *AddTestDataDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add test data default response
func (o *AddTestDataDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddTestDataDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}