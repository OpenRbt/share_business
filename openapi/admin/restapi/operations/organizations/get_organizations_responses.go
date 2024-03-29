// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washbonus/openapi/admin/models"
)

// GetOrganizationsOKCode is the HTTP code returned for type GetOrganizationsOK
const GetOrganizationsOKCode int = 200

/*
GetOrganizationsOK OK

swagger:response getOrganizationsOK
*/
type GetOrganizationsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Organization `json:"body,omitempty"`
}

// NewGetOrganizationsOK creates GetOrganizationsOK with default headers values
func NewGetOrganizationsOK() *GetOrganizationsOK {

	return &GetOrganizationsOK{}
}

// WithPayload adds the payload to the get organizations o k response
func (o *GetOrganizationsOK) WithPayload(payload []*models.Organization) *GetOrganizationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get organizations o k response
func (o *GetOrganizationsOK) SetPayload(payload []*models.Organization) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrganizationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Organization, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *GetOrganizationsOK) GetOrganizationsResponder() {}

/*
GetOrganizationsDefault Generic error response

swagger:response getOrganizationsDefault
*/
type GetOrganizationsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOrganizationsDefault creates GetOrganizationsDefault with default headers values
func NewGetOrganizationsDefault(code int) *GetOrganizationsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetOrganizationsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get organizations default response
func (o *GetOrganizationsDefault) WithStatusCode(code int) *GetOrganizationsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get organizations default response
func (o *GetOrganizationsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get organizations default response
func (o *GetOrganizationsDefault) WithPayload(payload *models.Error) *GetOrganizationsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get organizations default response
func (o *GetOrganizationsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrganizationsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetOrganizationsDefault) GetOrganizationsResponder() {}

type GetOrganizationsNotImplementedResponder struct {
	middleware.Responder
}

func (*GetOrganizationsNotImplementedResponder) GetOrganizationsResponder() {}

func GetOrganizationsNotImplemented() GetOrganizationsResponder {
	return &GetOrganizationsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetOrganizations has not yet been implemented",
		),
	}
}

type GetOrganizationsResponder interface {
	middleware.Responder
	GetOrganizationsResponder()
}
