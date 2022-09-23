// Code generated by go-swagger; DO NOT EDIT.

package bonus_balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bonus balance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bonus balance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddBonusBalance(params *AddBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddBonusBalanceCreated, error)

	DeleteBonusBalance(params *DeleteBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBonusBalanceNoContent, error)

	EditBonusBalance(params *EditBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditBonusBalanceOK, error)

	GetBonusBalance(params *GetBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBonusBalanceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddBonusBalance add bonus balance API
*/
func (a *Client) AddBonusBalance(params *AddBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddBonusBalanceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddBonusBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addBonusBalance",
		Method:             "POST",
		PathPattern:        "/balance/add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddBonusBalanceReader{formats: a.formats},
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
	success, ok := result.(*AddBonusBalanceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddBonusBalanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteBonusBalance delete bonus balance API
*/
func (a *Client) DeleteBonusBalance(params *DeleteBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBonusBalanceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBonusBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBonusBalance",
		Method:             "DELETE",
		PathPattern:        "/balance/deleted",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBonusBalanceReader{formats: a.formats},
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
	success, ok := result.(*DeleteBonusBalanceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBonusBalanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EditBonusBalance edit bonus balance API
*/
func (a *Client) EditBonusBalance(params *EditBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditBonusBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditBonusBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "editBonusBalance",
		Method:             "PUT",
		PathPattern:        "/balance/edit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EditBonusBalanceReader{formats: a.formats},
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
	success, ok := result.(*EditBonusBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EditBonusBalanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBonusBalance get bonus balance API
*/
func (a *Client) GetBonusBalance(params *GetBonusBalanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBonusBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBonusBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBonusBalance",
		Method:             "POST",
		PathPattern:        "/balance/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBonusBalanceReader{formats: a.formats},
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
	success, ok := result.(*GetBonusBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBonusBalanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
