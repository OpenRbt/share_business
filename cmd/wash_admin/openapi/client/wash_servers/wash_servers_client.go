// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new wash servers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wash servers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Add(params *AddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddNoContent, error)

	Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNoContent, error)

	Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error)

	Update(params *UpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Add add API
*/
func (a *Client) Add(params *AddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "add",
		Method:             "PUT",
		PathPattern:        "/wash-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for add: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Delete delete API
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete",
		Method:             "DELETE",
		PathPattern:        "/wash-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get get API
*/
func (a *Client) Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get",
		Method:             "GET",
		PathPattern:        "/wash-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Update update API
*/
func (a *Client) Update(params *UpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update",
		Method:             "PATCH",
		PathPattern:        "/wash-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}