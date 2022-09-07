// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wash-bonus/internal/api/restapi/models"
)

// DeleteRoleNoContentCode is the HTTP code returned for type DeleteRoleNoContent
const DeleteRoleNoContentCode int = 204

/*
DeleteRoleNoContent Deleted

swagger:response deleteRoleNoContent
*/
type DeleteRoleNoContent struct {
}

// NewDeleteRoleNoContent creates DeleteRoleNoContent with default headers values
func NewDeleteRoleNoContent() *DeleteRoleNoContent {

	return &DeleteRoleNoContent{}
}

// WriteResponse to the client
func (o *DeleteRoleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
DeleteRoleDefault error

swagger:response deleteRoleDefault
*/
type DeleteRoleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteRoleDefault creates DeleteRoleDefault with default headers values
func NewDeleteRoleDefault(code int) *DeleteRoleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteRoleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete role default response
func (o *DeleteRoleDefault) WithStatusCode(code int) *DeleteRoleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete role default response
func (o *DeleteRoleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete role default response
func (o *DeleteRoleDefault) WithPayload(payload *models.Error) *DeleteRoleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete role default response
func (o *DeleteRoleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRoleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
