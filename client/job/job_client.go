// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new job API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for job API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Delete(params *DeleteParams, opts ...ClientOption) (*DeleteOK, error)

	Input(params *InputParams, opts ...ClientOption) (*InputOK, error)

	Kill(params *KillParams, opts ...ClientOption) (*KillOK, error)

	List(params *ListParams, opts ...ClientOption) (*ListOK, error)

	Query(params *QueryParams, opts ...ClientOption) (*QueryOK, error)

	Submit(params *SubmitParams, opts ...ClientOption) (*SubmitOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Delete Delete job
*/
func (a *Client) Delete(params *DeleteParams, opts ...ClientOption) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete",
		Method:             "DELETE",
		PathPattern:        "/jobs/{urn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
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
	success, ok := result.(*DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Input Expect and input
*/
func (a *Client) Input(params *InputParams, opts ...ClientOption) (*InputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInputParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "input",
		Method:             "POST",
		PathPattern:        "/jobs/{urn}/input",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InputReader{formats: a.formats},
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
	success, ok := result.(*InputOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for input: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Kill Kill job
*/
func (a *Client) Kill(params *KillParams, opts ...ClientOption) (*KillOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKillParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "kill",
		Method:             "POST",
		PathPattern:        "/jobs/{urn}/kill",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KillReader{formats: a.formats},
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
	success, ok := result.(*KillOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for kill: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  List List Jobs
*/
func (a *Client) List(params *ListParams, opts ...ClientOption) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list",
		Method:             "GET",
		PathPattern:        "/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
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
	success, ok := result.(*ListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Query Query job
*/
func (a *Client) Query(params *QueryParams, opts ...ClientOption) (*QueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "query",
		Method:             "GET",
		PathPattern:        "/jobs/{urn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QueryReader{formats: a.formats},
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
	success, ok := result.(*QueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for query: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Submit Submit jobs
*/
func (a *Client) Submit(params *SubmitParams, opts ...ClientOption) (*SubmitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "submit",
		Method:             "POST",
		PathPattern:        "/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubmitReader{formats: a.formats},
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
	success, ok := result.(*SubmitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for submit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
