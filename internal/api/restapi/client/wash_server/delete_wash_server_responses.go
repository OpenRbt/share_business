// Code generated by go-swagger; DO NOT EDIT.

package wash_server

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

// DeleteWashServerReader is a Reader for the DeleteWashServer structure.
type DeleteWashServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWashServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWashServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteWashServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWashServerNoContent creates a DeleteWashServerNoContent with default headers values
func NewDeleteWashServerNoContent() *DeleteWashServerNoContent {
	return &DeleteWashServerNoContent{}
}

/* DeleteWashServerNoContent describes a response with status code 204, with default header values.

Deleted
*/
type DeleteWashServerNoContent struct {
}

func (o *DeleteWashServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /washServer/delete][%d] deleteWashServerNoContent ", 204)
}

func (o *DeleteWashServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWashServerDefault creates a DeleteWashServerDefault with default headers values
func NewDeleteWashServerDefault(code int) *DeleteWashServerDefault {
	return &DeleteWashServerDefault{
		_statusCode: code,
	}
}

/* DeleteWashServerDefault describes a response with status code -1, with default header values.

error
*/
type DeleteWashServerDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete wash server default response
func (o *DeleteWashServerDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWashServerDefault) Error() string {
	return fmt.Sprintf("[DELETE /washServer/delete][%d] deleteWashServer default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteWashServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteWashServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteWashServerBody delete wash server body
swagger:model DeleteWashServerBody
*/
type DeleteWashServerBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this delete wash server body
func (o *DeleteWashServerBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete wash server body based on context it is used
func (o *DeleteWashServerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWashServerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWashServerBody) UnmarshalBinary(b []byte) error {
	var res DeleteWashServerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
