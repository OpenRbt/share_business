// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// EditWashSessionReader is a Reader for the EditWashSession structure.
type EditWashSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditWashSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditWashSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEditWashSessionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEditWashSessionOK creates a EditWashSessionOK with default headers values
func NewEditWashSessionOK() *EditWashSessionOK {
	return &EditWashSessionOK{}
}

/* EditWashSessionOK describes a response with status code 200, with default header values.

OK
*/
type EditWashSessionOK struct {
}

func (o *EditWashSessionOK) Error() string {
	return fmt.Sprintf("[PUT /washSession/edit][%d] editWashSessionOK ", 200)
}

func (o *EditWashSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditWashSessionDefault creates a EditWashSessionDefault with default headers values
func NewEditWashSessionDefault(code int) *EditWashSessionDefault {
	return &EditWashSessionDefault{
		_statusCode: code,
	}
}

/* EditWashSessionDefault describes a response with status code -1, with default header values.

error
*/
type EditWashSessionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the edit wash session default response
func (o *EditWashSessionDefault) Code() int {
	return o._statusCode
}

func (o *EditWashSessionDefault) Error() string {
	return fmt.Sprintf("[PUT /washSession/edit][%d] editWashSession default  %+v", o._statusCode, o.Payload)
}
func (o *EditWashSessionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EditWashSessionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EditWashSessionBody edit wash session body
swagger:model EditWashSessionBody
*/
type EditWashSessionBody struct {

	// data
	Data *models.WashSessionAdd `json:"data,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this edit wash session body
func (o *EditWashSessionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditWashSessionBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edit wash session body based on the context it is used
func (o *EditWashSessionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditWashSessionBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EditWashSessionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditWashSessionBody) UnmarshalBinary(b []byte) error {
	var res EditWashSessionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
