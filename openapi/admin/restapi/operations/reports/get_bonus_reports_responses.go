// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washbonus/openapi/admin/models"
)

// GetBonusReportsOKCode is the HTTP code returned for type GetBonusReportsOK
const GetBonusReportsOKCode int = 200

/*
GetBonusReportsOK OK

swagger:response getBonusReportsOK
*/
type GetBonusReportsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReportPage `json:"body,omitempty"`
}

// NewGetBonusReportsOK creates GetBonusReportsOK with default headers values
func NewGetBonusReportsOK() *GetBonusReportsOK {

	return &GetBonusReportsOK{}
}

// WithPayload adds the payload to the get bonus reports o k response
func (o *GetBonusReportsOK) WithPayload(payload *models.ReportPage) *GetBonusReportsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bonus reports o k response
func (o *GetBonusReportsOK) SetPayload(payload *models.ReportPage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBonusReportsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetBonusReportsOK) GetBonusReportsResponder() {}

// GetBonusReportsForbiddenCode is the HTTP code returned for type GetBonusReportsForbidden
const GetBonusReportsForbiddenCode int = 403

/*
GetBonusReportsForbidden Generic error response

swagger:response getBonusReportsForbidden
*/
type GetBonusReportsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBonusReportsForbidden creates GetBonusReportsForbidden with default headers values
func NewGetBonusReportsForbidden() *GetBonusReportsForbidden {

	return &GetBonusReportsForbidden{}
}

// WithPayload adds the payload to the get bonus reports forbidden response
func (o *GetBonusReportsForbidden) WithPayload(payload *models.Error) *GetBonusReportsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bonus reports forbidden response
func (o *GetBonusReportsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBonusReportsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetBonusReportsForbidden) GetBonusReportsResponder() {}

// GetBonusReportsNotFoundCode is the HTTP code returned for type GetBonusReportsNotFound
const GetBonusReportsNotFoundCode int = 404

/*
GetBonusReportsNotFound Generic error response

swagger:response getBonusReportsNotFound
*/
type GetBonusReportsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBonusReportsNotFound creates GetBonusReportsNotFound with default headers values
func NewGetBonusReportsNotFound() *GetBonusReportsNotFound {

	return &GetBonusReportsNotFound{}
}

// WithPayload adds the payload to the get bonus reports not found response
func (o *GetBonusReportsNotFound) WithPayload(payload *models.Error) *GetBonusReportsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bonus reports not found response
func (o *GetBonusReportsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBonusReportsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetBonusReportsNotFound) GetBonusReportsResponder() {}

/*
GetBonusReportsDefault Generic error response

swagger:response getBonusReportsDefault
*/
type GetBonusReportsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBonusReportsDefault creates GetBonusReportsDefault with default headers values
func NewGetBonusReportsDefault(code int) *GetBonusReportsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBonusReportsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get bonus reports default response
func (o *GetBonusReportsDefault) WithStatusCode(code int) *GetBonusReportsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get bonus reports default response
func (o *GetBonusReportsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get bonus reports default response
func (o *GetBonusReportsDefault) WithPayload(payload *models.Error) *GetBonusReportsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bonus reports default response
func (o *GetBonusReportsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBonusReportsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetBonusReportsDefault) GetBonusReportsResponder() {}

type GetBonusReportsNotImplementedResponder struct {
	middleware.Responder
}

func (*GetBonusReportsNotImplementedResponder) GetBonusReportsResponder() {}

func GetBonusReportsNotImplemented() GetBonusReportsResponder {
	return &GetBonusReportsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetBonusReports has not yet been implemented",
		),
	}
}

type GetBonusReportsResponder interface {
	middleware.Responder
	GetBonusReportsResponder()
}
