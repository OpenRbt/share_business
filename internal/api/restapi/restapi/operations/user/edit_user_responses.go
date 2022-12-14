// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wash-bonus/internal/api/restapi/models"
)

// EditUserOKCode is the HTTP code returned for type EditUserOK
const EditUserOKCode int = 200

/*EditUserOK OK

swagger:response editUserOK
*/
type EditUserOK struct {
}

// NewEditUserOK creates EditUserOK with default headers values
func NewEditUserOK() *EditUserOK {

	return &EditUserOK{}
}

// WriteResponse to the client
func (o *EditUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*EditUserDefault error

swagger:response editUserDefault
*/
type EditUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEditUserDefault creates EditUserDefault with default headers values
func NewEditUserDefault(code int) *EditUserDefault {
	if code <= 0 {
		code = 500
	}

	return &EditUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the edit user default response
func (o *EditUserDefault) WithStatusCode(code int) *EditUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the edit user default response
func (o *EditUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the edit user default response
func (o *EditUserDefault) WithPayload(payload *models.Error) *EditUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit user default response
func (o *EditUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
