// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washbonus/openapi/admin/models"
)

// UpdateAdminUserRoleNoContentCode is the HTTP code returned for type UpdateAdminUserRoleNoContent
const UpdateAdminUserRoleNoContentCode int = 204

/*
UpdateAdminUserRoleNoContent OK

swagger:response updateAdminUserRoleNoContent
*/
type UpdateAdminUserRoleNoContent struct {
}

// NewUpdateAdminUserRoleNoContent creates UpdateAdminUserRoleNoContent with default headers values
func NewUpdateAdminUserRoleNoContent() *UpdateAdminUserRoleNoContent {

	return &UpdateAdminUserRoleNoContent{}
}

// WriteResponse to the client
func (o *UpdateAdminUserRoleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *UpdateAdminUserRoleNoContent) UpdateAdminUserRoleResponder() {}

/*
UpdateAdminUserRoleDefault Generic error response

swagger:response updateAdminUserRoleDefault
*/
type UpdateAdminUserRoleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateAdminUserRoleDefault creates UpdateAdminUserRoleDefault with default headers values
func NewUpdateAdminUserRoleDefault(code int) *UpdateAdminUserRoleDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateAdminUserRoleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update admin user role default response
func (o *UpdateAdminUserRoleDefault) WithStatusCode(code int) *UpdateAdminUserRoleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update admin user role default response
func (o *UpdateAdminUserRoleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update admin user role default response
func (o *UpdateAdminUserRoleDefault) WithPayload(payload *models.Error) *UpdateAdminUserRoleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update admin user role default response
func (o *UpdateAdminUserRoleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAdminUserRoleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateAdminUserRoleDefault) UpdateAdminUserRoleResponder() {}

type UpdateAdminUserRoleNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateAdminUserRoleNotImplementedResponder) UpdateAdminUserRoleResponder() {}

func UpdateAdminUserRoleNotImplemented() UpdateAdminUserRoleResponder {
	return &UpdateAdminUserRoleNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateAdminUserRole has not yet been implemented",
		),
	}
}

type UpdateAdminUserRoleResponder interface {
	middleware.Responder
	UpdateAdminUserRoleResponder()
}
