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

// NewRemoveStorageParams creates a new RemoveStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveStorageParams() *RemoveStorageParams {
	return &RemoveStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveStorageParamsWithTimeout creates a new RemoveStorageParams object
// with the ability to set a timeout on a request.
func NewRemoveStorageParamsWithTimeout(timeout time.Duration) *RemoveStorageParams {
	return &RemoveStorageParams{
		timeout: timeout,
	}
}

// NewRemoveStorageParamsWithContext creates a new RemoveStorageParams object
// with the ability to set a context for a request.
func NewRemoveStorageParamsWithContext(ctx context.Context) *RemoveStorageParams {
	return &RemoveStorageParams{
		Context: ctx,
	}
}

// NewRemoveStorageParamsWithHTTPClient creates a new RemoveStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveStorageParamsWithHTTPClient(client *http.Client) *RemoveStorageParams {
	return &RemoveStorageParams{
		HTTPClient: client,
	}
}

/*
RemoveStorageParams contains all the parameters to send to the API endpoint

	for the remove storage operation.

	Typically these are written to a http.Request.
*/
type RemoveStorageParams struct {

	/* Name.

	   Storage ID or name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveStorageParams) WithDefaults() *RemoveStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove storage params
func (o *RemoveStorageParams) WithTimeout(timeout time.Duration) *RemoveStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove storage params
func (o *RemoveStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove storage params
func (o *RemoveStorageParams) WithContext(ctx context.Context) *RemoveStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove storage params
func (o *RemoveStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove storage params
func (o *RemoveStorageParams) WithHTTPClient(client *http.Client) *RemoveStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove storage params
func (o *RemoveStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the remove storage params
func (o *RemoveStorageParams) WithName(name string) *RemoveStorageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the remove storage params
func (o *RemoveStorageParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
