// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washbonus/openapi/bonus/models"
)

// GetSessionByIDOKCode is the HTTP code returned for type GetSessionByIDOK
const GetSessionByIDOKCode int = 200

/*GetSessionByIDOK OK

swagger:response getSessionByIdOK
*/
type GetSessionByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Session `json:"body,omitempty"`
}

// NewGetSessionByIDOK creates GetSessionByIDOK with default headers values
func NewGetSessionByIDOK() *GetSessionByIDOK {

	return &GetSessionByIDOK{}
}

// WithPayload adds the payload to the get session by Id o k response
func (o *GetSessionByIDOK) WithPayload(payload *models.Session) *GetSessionByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get session by Id o k response
func (o *GetSessionByIDOK) SetPayload(payload *models.Session) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSessionByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSessionByIDOK) GetSessionByIDResponder() {}

/*GetSessionByIDDefault Generic error response

swagger:response getSessionByIdDefault
*/
type GetSessionByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSessionByIDDefault creates GetSessionByIDDefault with default headers values
func NewGetSessionByIDDefault(code int) *GetSessionByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSessionByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get session by Id default response
func (o *GetSessionByIDDefault) WithStatusCode(code int) *GetSessionByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get session by Id default response
func (o *GetSessionByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get session by Id default response
func (o *GetSessionByIDDefault) WithPayload(payload *models.Error) *GetSessionByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get session by Id default response
func (o *GetSessionByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSessionByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetSessionByIDDefault) GetSessionByIDResponder() {}

type GetSessionByIDNotImplementedResponder struct {
	middleware.Responder
}

func (*GetSessionByIDNotImplementedResponder) GetSessionByIDResponder() {}

func GetSessionByIDNotImplemented() GetSessionByIDResponder {
	return &GetSessionByIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetSessionByID has not yet been implemented",
		),
	}
}

type GetSessionByIDResponder interface {
	middleware.Responder
	GetSessionByIDResponder()
}
