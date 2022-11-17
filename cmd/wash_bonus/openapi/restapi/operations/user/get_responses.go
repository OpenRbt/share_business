// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"wash_bonus/openapi/models"
)

// GetOKCode is the HTTP code returned for type GetOK
const GetOKCode int = 200

/*
GetOK OK

swagger:response getOK
*/
type GetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewGetOK creates GetOK with default headers values
func NewGetOK() *GetOK {

	return &GetOK{}
}

// WithPayload adds the payload to the get o k response
func (o *GetOK) WithPayload(payload *models.Profile) *GetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get o k response
func (o *GetOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetOK) GetResponder() {}

// GetNotFoundCode is the HTTP code returned for type GetNotFound
const GetNotFoundCode int = 404

/*
GetNotFound Profile not exists

swagger:response getNotFound
*/
type GetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNotFound creates GetNotFound with default headers values
func NewGetNotFound() *GetNotFound {

	return &GetNotFound{}
}

// WithPayload adds the payload to the get not found response
func (o *GetNotFound) WithPayload(payload *models.Error) *GetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get not found response
func (o *GetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetNotFound) GetResponder() {}

// GetInternalServerErrorCode is the HTTP code returned for type GetInternalServerError
const GetInternalServerErrorCode int = 500

/*
GetInternalServerError Internal error

swagger:response getInternalServerError
*/
type GetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInternalServerError creates GetInternalServerError with default headers values
func NewGetInternalServerError() *GetInternalServerError {

	return &GetInternalServerError{}
}

// WithPayload adds the payload to the get internal server error response
func (o *GetInternalServerError) WithPayload(payload *models.Error) *GetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get internal server error response
func (o *GetInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetInternalServerError) GetResponder() {}

type GetNotImplementedResponder struct {
	middleware.Responder
}

func (*GetNotImplementedResponder) GetResponder() {}

func GetNotImplemented() GetResponder {
	return &GetNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Get has not yet been implemented",
		),
	}
}

type GetResponder interface {
	middleware.Responder
	GetResponder()
}
