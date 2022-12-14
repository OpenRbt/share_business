// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// GetWashSessionReader is a Reader for the GetWashSession structure.
type GetWashSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWashSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWashSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWashSessionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWashSessionOK creates a GetWashSessionOK with default headers values
func NewGetWashSessionOK() *GetWashSessionOK {
	return &GetWashSessionOK{}
}

/* GetWashSessionOK describes a response with status code 200, with default header values.

OK
*/
type GetWashSessionOK struct {
	Payload *models.WashSession
}

func (o *GetWashSessionOK) Error() string {
	return fmt.Sprintf("[POST /washSession/get][%d] getWashSessionOK  %+v", 200, o.Payload)
}
func (o *GetWashSessionOK) GetPayload() *models.WashSession {
	return o.Payload
}

func (o *GetWashSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WashSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWashSessionDefault creates a GetWashSessionDefault with default headers values
func NewGetWashSessionDefault(code int) *GetWashSessionDefault {
	return &GetWashSessionDefault{
		_statusCode: code,
	}
}

/* GetWashSessionDefault describes a response with status code -1, with default header values.

error
*/
type GetWashSessionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wash session default response
func (o *GetWashSessionDefault) Code() int {
	return o._statusCode
}

func (o *GetWashSessionDefault) Error() string {
	return fmt.Sprintf("[POST /washSession/get][%d] getWashSession default  %+v", o._statusCode, o.Payload)
}
func (o *GetWashSessionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWashSessionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWashSessionBody get wash session body
swagger:model GetWashSessionBody
*/
type GetWashSessionBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this get wash session body
func (o *GetWashSessionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wash session body based on context it is used
func (o *GetWashSessionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWashSessionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWashSessionBody) UnmarshalBinary(b []byte) error {
	var res GetWashSessionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
