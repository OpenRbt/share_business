// Code generated by go-swagger; DO NOT EDIT.

package server_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washbonus/openapi/admin/models"
)

// CreateServerGroupOKCode is the HTTP code returned for type CreateServerGroupOK
const CreateServerGroupOKCode int = 200

/*
CreateServerGroupOK Successfull created

swagger:response createServerGroupOK
*/
type CreateServerGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServerGroup `json:"body,omitempty"`
}

// NewCreateServerGroupOK creates CreateServerGroupOK with default headers values
func NewCreateServerGroupOK() *CreateServerGroupOK {

	return &CreateServerGroupOK{}
}

// WithPayload adds the payload to the create server group o k response
func (o *CreateServerGroupOK) WithPayload(payload *models.ServerGroup) *CreateServerGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server group o k response
func (o *CreateServerGroupOK) SetPayload(payload *models.ServerGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateServerGroupOK) CreateServerGroupResponder() {}

/*
CreateServerGroupDefault Generic error response

swagger:response createServerGroupDefault
*/
type CreateServerGroupDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerGroupDefault creates CreateServerGroupDefault with default headers values
func NewCreateServerGroupDefault(code int) *CreateServerGroupDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateServerGroupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create server group default response
func (o *CreateServerGroupDefault) WithStatusCode(code int) *CreateServerGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create server group default response
func (o *CreateServerGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create server group default response
func (o *CreateServerGroupDefault) WithPayload(payload *models.Error) *CreateServerGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server group default response
func (o *CreateServerGroupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateServerGroupDefault) CreateServerGroupResponder() {}

type CreateServerGroupNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateServerGroupNotImplementedResponder) CreateServerGroupResponder() {}

func CreateServerGroupNotImplemented() CreateServerGroupResponder {
	return &CreateServerGroupNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CreateServerGroup has not yet been implemented",
		),
	}
}

type CreateServerGroupResponder interface {
	middleware.Responder
	CreateServerGroupResponder()
}
