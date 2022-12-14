// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new role API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for role API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddPermissionsRole(params *AddPermissionsRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddPermissionsRoleOK, error)

	AddRole(params *AddRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddRoleCreated, error)

	DeletePermissionsRole(params *DeletePermissionsRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePermissionsRoleOK, error)

	DeleteRole(params *DeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRoleNoContent, error)

	EditRole(params *EditRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditRoleOK, error)

	GetRole(params *GetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleOK, error)

	ListRole(params *ListRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRoleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddPermissionsRole add permissions role API
*/
func (a *Client) AddPermissionsRole(params *AddPermissionsRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddPermissionsRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPermissionsRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addPermissionsRole",
		Method:             "POST",
		PathPattern:        "/role/addPermissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddPermissionsRoleReader{formats: a.formats},
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
	success, ok := result.(*AddPermissionsRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddPermissionsRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddRole add role API
*/
func (a *Client) AddRole(params *AddRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddRoleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addRole",
		Method:             "POST",
		PathPattern:        "/role/add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddRoleReader{formats: a.formats},
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
	success, ok := result.(*AddRoleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeletePermissionsRole delete permissions role API
*/
func (a *Client) DeletePermissionsRole(params *DeletePermissionsRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePermissionsRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePermissionsRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePermissionsRole",
		Method:             "POST",
		PathPattern:        "/role/deletePermissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePermissionsRoleReader{formats: a.formats},
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
	success, ok := result.(*DeletePermissionsRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePermissionsRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRole delete role API
*/
func (a *Client) DeleteRole(params *DeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRoleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRole",
		Method:             "DELETE",
		PathPattern:        "/role/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRoleReader{formats: a.formats},
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
	success, ok := result.(*DeleteRoleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EditRole edit role API
*/
func (a *Client) EditRole(params *EditRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "editRole",
		Method:             "PUT",
		PathPattern:        "/role/edit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EditRoleReader{formats: a.formats},
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
	success, ok := result.(*EditRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EditRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRole get role API
*/
func (a *Client) GetRole(params *GetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRole",
		Method:             "POST",
		PathPattern:        "/role/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRoleReader{formats: a.formats},
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
	success, ok := result.(*GetRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListRole list role API
*/
func (a *Client) ListRole(params *ListRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listRole",
		Method:             "POST",
		PathPattern:        "/role/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListRoleReader{formats: a.formats},
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
	success, ok := result.(*ListRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
