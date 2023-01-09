// Code generated by go-swagger; DO NOT EDIT.

package crypto

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new crypto API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for crypto API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVault(params *CreateVaultParams, opts ...ClientOption) (*CreateVaultOK, error)

	Decrypt(params *DecryptParams, opts ...ClientOption) (*DecryptOK, error)

	DeleteVault(params *DeleteVaultParams, opts ...ClientOption) (*DeleteVaultOK, error)

	Encrypt(params *EncryptParams, opts ...ClientOption) (*EncryptOK, error)

	ListVaults(params *ListVaultsParams, opts ...ClientOption) (*ListVaultsOK, error)

	QueryVault(params *QueryVaultParams, opts ...ClientOption) (*QueryVaultOK, error)

	SecretDecrypt(params *SecretDecryptParams, opts ...ClientOption) (*SecretDecryptOK, error)

	SecretEncrypt(params *SecretEncryptParams, opts ...ClientOption) (*SecretEncryptOK, error)

	UpdateVault(params *UpdateVaultParams, opts ...ClientOption) (*UpdateVaultOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVault Create vault
*/
func (a *Client) CreateVault(params *CreateVaultParams, opts ...ClientOption) (*CreateVaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createVault",
		Method:             "POST",
		PathPattern:        "/crypto/vault/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateVaultReader{formats: a.formats},
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
	success, ok := result.(*CreateVaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createVault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Decrypt Decrypt data
*/
func (a *Client) Decrypt(params *DecryptParams, opts ...ClientOption) (*DecryptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecryptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "decrypt",
		Method:             "POST",
		PathPattern:        "/crypto/decrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DecryptReader{formats: a.formats},
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
	success, ok := result.(*DecryptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for decrypt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVault Delete vault
*/
func (a *Client) DeleteVault(params *DeleteVaultParams, opts ...ClientOption) (*DeleteVaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVault",
		Method:             "DELETE",
		PathPattern:        "/crypto/vault/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteVaultReader{formats: a.formats},
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
	success, ok := result.(*DeleteVaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Encrypt Encrypt data
*/
func (a *Client) Encrypt(params *EncryptParams, opts ...ClientOption) (*EncryptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEncryptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "encrypt",
		Method:             "POST",
		PathPattern:        "/crypto/encrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EncryptReader{formats: a.formats},
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
	success, ok := result.(*EncryptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for encrypt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListVaults List vaults
*/
func (a *Client) ListVaults(params *ListVaultsParams, opts ...ClientOption) (*ListVaultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVaultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listVaults",
		Method:             "GET",
		PathPattern:        "/crypto/vaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListVaultsReader{formats: a.formats},
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
	success, ok := result.(*ListVaultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listVaults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryVault Query vault
*/
func (a *Client) QueryVault(params *QueryVaultParams, opts ...ClientOption) (*QueryVaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryVaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryVault",
		Method:             "GET",
		PathPattern:        "/crypto/vault/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QueryVaultReader{formats: a.formats},
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
	success, ok := result.(*QueryVaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryVault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretDecrypt Decrypt k8s secret
*/
func (a *Client) SecretDecrypt(params *SecretDecryptParams, opts ...ClientOption) (*SecretDecryptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretDecryptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "secretDecrypt",
		Method:             "POST",
		PathPattern:        "/crypto/secret/decrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SecretDecryptReader{formats: a.formats},
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
	success, ok := result.(*SecretDecryptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secretDecrypt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretEncrypt Encrypt k8s secret
*/
func (a *Client) SecretEncrypt(params *SecretEncryptParams, opts ...ClientOption) (*SecretEncryptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretEncryptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "secretEncrypt",
		Method:             "POST",
		PathPattern:        "/crypto/secret/encrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SecretEncryptReader{formats: a.formats},
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
	success, ok := result.(*SecretEncryptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secretEncrypt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVault Update vault
*/
func (a *Client) UpdateVault(params *UpdateVaultParams, opts ...ClientOption) (*UpdateVaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVault",
		Method:             "PUT",
		PathPattern:        "/crypto/vault/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateVaultReader{formats: a.formats},
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
	success, ok := result.(*UpdateVaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}