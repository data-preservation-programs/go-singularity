// Code generated by go-swagger; DO NOT EDIT.

package wallet_association

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new wallet association API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wallet association API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AttachWallet(params *AttachWalletParams, opts ...ClientOption) (*AttachWalletOK, error)

	DetachWallet(params *DetachWalletParams, opts ...ClientOption) (*DetachWalletOK, error)

	ListAttachedWallets(params *ListAttachedWalletsParams, opts ...ClientOption) (*ListAttachedWalletsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AttachWallet attaches a new wallet with a preparation
*/
func (a *Client) AttachWallet(params *AttachWalletParams, opts ...ClientOption) (*AttachWalletOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachWalletParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachWallet",
		Method:             "POST",
		PathPattern:        "/preparation/{id}/wallet/{wallet}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachWalletReader{formats: a.formats},
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
	success, ok := result.(*AttachWalletOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AttachWallet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetachWallet detaches a new wallet from a preparation
*/
func (a *Client) DetachWallet(params *DetachWalletParams, opts ...ClientOption) (*DetachWalletOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachWalletParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DetachWallet",
		Method:             "DELETE",
		PathPattern:        "/preparation/{id}/wallet/{wallet}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DetachWalletReader{formats: a.formats},
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
	success, ok := result.(*DetachWalletOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DetachWallet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListAttachedWallets lists all wallets of a preparation
*/
func (a *Client) ListAttachedWallets(params *ListAttachedWalletsParams, opts ...ClientOption) (*ListAttachedWalletsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAttachedWalletsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListAttachedWallets",
		Method:             "POST",
		PathPattern:        "/preparation/{id}/wallet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAttachedWalletsReader{formats: a.formats},
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
	success, ok := result.(*ListAttachedWalletsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListAttachedWallets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
