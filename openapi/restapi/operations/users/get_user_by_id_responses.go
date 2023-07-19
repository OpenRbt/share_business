// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washBonus/openapi/models"
)

// GetUserByIDOKCode is the HTTP code returned for type GetUserByIDOK
const GetUserByIDOKCode int = 200

/*
GetUserByIDOK OK

swagger:response getUserByIdOK
*/
type GetUserByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserByIDOK creates GetUserByIDOK with default headers values
func NewGetUserByIDOK() *GetUserByIDOK {

	return &GetUserByIDOK{}
}

// WithPayload adds the payload to the get user by Id o k response
func (o *GetUserByIDOK) WithPayload(payload *models.User) *GetUserByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id o k response
func (o *GetUserByIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetUserByIDOK) GetUserByIDResponder() {}

// GetUserByIDForbiddenCode is the HTTP code returned for type GetUserByIDForbidden
const GetUserByIDForbiddenCode int = 403

/*
GetUserByIDForbidden Forbidden

swagger:response getUserByIdForbidden
*/
type GetUserByIDForbidden struct {
}

// NewGetUserByIDForbidden creates GetUserByIDForbidden with default headers values
func NewGetUserByIDForbidden() *GetUserByIDForbidden {

	return &GetUserByIDForbidden{}
}

// WriteResponse to the client
func (o *GetUserByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

func (o *GetUserByIDForbidden) GetUserByIDResponder() {}

// GetUserByIDNotFoundCode is the HTTP code returned for type GetUserByIDNotFound
const GetUserByIDNotFoundCode int = 404

/*
GetUserByIDNotFound Not found

swagger:response getUserByIdNotFound
*/
type GetUserByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserByIDNotFound creates GetUserByIDNotFound with default headers values
func NewGetUserByIDNotFound() *GetUserByIDNotFound {

	return &GetUserByIDNotFound{}
}

// WithPayload adds the payload to the get user by Id not found response
func (o *GetUserByIDNotFound) WithPayload(payload *models.Error) *GetUserByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id not found response
func (o *GetUserByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetUserByIDNotFound) GetUserByIDResponder() {}

// GetUserByIDInternalServerErrorCode is the HTTP code returned for type GetUserByIDInternalServerError
const GetUserByIDInternalServerErrorCode int = 500

/*
GetUserByIDInternalServerError Internal error

swagger:response getUserByIdInternalServerError
*/
type GetUserByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserByIDInternalServerError creates GetUserByIDInternalServerError with default headers values
func NewGetUserByIDInternalServerError() *GetUserByIDInternalServerError {

	return &GetUserByIDInternalServerError{}
}

// WithPayload adds the payload to the get user by Id internal server error response
func (o *GetUserByIDInternalServerError) WithPayload(payload *models.Error) *GetUserByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id internal server error response
func (o *GetUserByIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetUserByIDInternalServerError) GetUserByIDResponder() {}

type GetUserByIDNotImplementedResponder struct {
	middleware.Responder
}

func (*GetUserByIDNotImplementedResponder) GetUserByIDResponder() {}

func GetUserByIDNotImplemented() GetUserByIDResponder {
	return &GetUserByIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetUserByID has not yet been implemented",
		),
	}
}

type GetUserByIDResponder interface {
	middleware.Responder
	GetUserByIDResponder()
}