// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListStoragesParams creates a new ListStoragesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListStoragesParams() *ListStoragesParams {
	return &ListStoragesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListStoragesParamsWithTimeout creates a new ListStoragesParams object
// with the ability to set a timeout on a request.
func NewListStoragesParamsWithTimeout(timeout time.Duration) *ListStoragesParams {
	return &ListStoragesParams{
		timeout: timeout,
	}
}

// NewListStoragesParamsWithContext creates a new ListStoragesParams object
// with the ability to set a context for a request.
func NewListStoragesParamsWithContext(ctx context.Context) *ListStoragesParams {
	return &ListStoragesParams{
		Context: ctx,
	}
}

// NewListStoragesParamsWithHTTPClient creates a new ListStoragesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListStoragesParamsWithHTTPClient(client *http.Client) *ListStoragesParams {
	return &ListStoragesParams{
		HTTPClient: client,
	}
}

/*
ListStoragesParams contains all the parameters to send to the API endpoint

	for the list storages operation.

	Typically these are written to a http.Request.
*/
type ListStoragesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list storages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListStoragesParams) WithDefaults() *ListStoragesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list storages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListStoragesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list storages params
func (o *ListStoragesParams) WithTimeout(timeout time.Duration) *ListStoragesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list storages params
func (o *ListStoragesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list storages params
func (o *ListStoragesParams) WithContext(ctx context.Context) *ListStoragesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list storages params
func (o *ListStoragesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list storages params
func (o *ListStoragesParams) WithHTTPClient(client *http.Client) *ListStoragesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list storages params
func (o *ListStoragesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListStoragesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}