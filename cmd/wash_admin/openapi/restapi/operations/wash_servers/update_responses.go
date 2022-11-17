// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"wash_admin/openapi/models"
)

// UpdateNoContentCode is the HTTP code returned for type UpdateNoContent
const UpdateNoContentCode int = 204

/*
UpdateNoContent Success update

swagger:response updateNoContent
*/
type UpdateNoContent struct {
}

// NewUpdateNoContent creates UpdateNoContent with default headers values
func NewUpdateNoContent() *UpdateNoContent {

	return &UpdateNoContent{}
}

// WriteResponse to the client
func (o *UpdateNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *UpdateNoContent) UpdateResponder() {}

// UpdateNotFoundCode is the HTTP code returned for type UpdateNotFound
const UpdateNotFoundCode int = 404

/*
UpdateNotFound WashServer not exists

swagger:response updateNotFound
*/
type UpdateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateNotFound creates UpdateNotFound with default headers values
func NewUpdateNotFound() *UpdateNotFound {

	return &UpdateNotFound{}
}

// WithPayload adds the payload to the update not found response
func (o *UpdateNotFound) WithPayload(payload *models.Error) *UpdateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update not found response
func (o *UpdateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateNotFound) UpdateResponder() {}

// UpdateInternalServerErrorCode is the HTTP code returned for type UpdateInternalServerError
const UpdateInternalServerErrorCode int = 500

/*
UpdateInternalServerError Internal error

swagger:response updateInternalServerError
*/
type UpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateInternalServerError creates UpdateInternalServerError with default headers values
func NewUpdateInternalServerError() *UpdateInternalServerError {

	return &UpdateInternalServerError{}
}

// WithPayload adds the payload to the update internal server error response
func (o *UpdateInternalServerError) WithPayload(payload *models.Error) *UpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update internal server error response
func (o *UpdateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateInternalServerError) UpdateResponder() {}

type UpdateNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateNotImplementedResponder) UpdateResponder() {}

func UpdateNotImplemented() UpdateResponder {
	return &UpdateNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Update has not yet been implemented",
		),
	}
}

type UpdateResponder interface {
	middleware.Responder
	UpdateResponder()
}
