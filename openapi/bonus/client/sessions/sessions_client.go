// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sessions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sessions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssignUserToSession(params *AssignUserToSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssignUserToSessionNoContent, error)

	ChargeBonusesOnSession(params *ChargeBonusesOnSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargeBonusesOnSessionOK, error)

	GetSessionByID(params *GetSessionByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSessionByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AssignUserToSession assign user to session API
*/
func (a *Client) AssignUserToSession(params *AssignUserToSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssignUserToSessionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignUserToSessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assignUserToSession",
		Method:             "POST",
		PathPattern:        "/sessions/{sessionId}/assign-user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AssignUserToSessionReader{formats: a.formats},
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
	success, ok := result.(*AssignUserToSessionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AssignUserToSessionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChargeBonusesOnSession charge bonuses on session API
*/
func (a *Client) ChargeBonusesOnSession(params *ChargeBonusesOnSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargeBonusesOnSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargeBonusesOnSessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "chargeBonusesOnSession",
		Method:             "POST",
		PathPattern:        "/sessions/{sessionId}/bonuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChargeBonusesOnSessionReader{formats: a.formats},
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
	success, ok := result.(*ChargeBonusesOnSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargeBonusesOnSessionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSessionByID get session by Id API
*/
func (a *Client) GetSessionByID(params *GetSessionByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSessionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSessionByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSessionById",
		Method:             "GET",
		PathPattern:        "/sessions/{sessionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSessionByIDReader{formats: a.formats},
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
	success, ok := result.(*GetSessionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSessionByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
