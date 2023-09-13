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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewRenameStorageParams creates a new RenameStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenameStorageParams() *RenameStorageParams {
	return &RenameStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenameStorageParamsWithTimeout creates a new RenameStorageParams object
// with the ability to set a timeout on a request.
func NewRenameStorageParamsWithTimeout(timeout time.Duration) *RenameStorageParams {
	return &RenameStorageParams{
		timeout: timeout,
	}
}

// NewRenameStorageParamsWithContext creates a new RenameStorageParams object
// with the ability to set a context for a request.
func NewRenameStorageParamsWithContext(ctx context.Context) *RenameStorageParams {
	return &RenameStorageParams{
		Context: ctx,
	}
}

// NewRenameStorageParamsWithHTTPClient creates a new RenameStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenameStorageParamsWithHTTPClient(client *http.Client) *RenameStorageParams {
	return &RenameStorageParams{
		HTTPClient: client,
	}
}

/*
RenameStorageParams contains all the parameters to send to the API endpoint

	for the rename storage operation.

	Typically these are written to a http.Request.
*/
type RenameStorageParams struct {

	/* Name.

	   Storage ID or name
	*/
	Name string

	/* Request.

	   New storage name
	*/
	Request *models.StorageRenameRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rename storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenameStorageParams) WithDefaults() *RenameStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rename storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenameStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rename storage params
func (o *RenameStorageParams) WithTimeout(timeout time.Duration) *RenameStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rename storage params
func (o *RenameStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rename storage params
func (o *RenameStorageParams) WithContext(ctx context.Context) *RenameStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rename storage params
func (o *RenameStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rename storage params
func (o *RenameStorageParams) WithHTTPClient(client *http.Client) *RenameStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rename storage params
func (o *RenameStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the rename storage params
func (o *RenameStorageParams) WithName(name string) *RenameStorageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the rename storage params
func (o *RenameStorageParams) SetName(name string) {
	o.Name = name
}

// WithRequest adds the request to the rename storage params
func (o *RenameStorageParams) WithRequest(request *models.StorageRenameRequest) *RenameStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the rename storage params
func (o *RenameStorageParams) SetRequest(request *models.StorageRenameRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *RenameStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
