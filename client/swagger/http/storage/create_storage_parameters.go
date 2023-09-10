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

// NewCreateStorageParams creates a new CreateStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStorageParams() *CreateStorageParams {
	return &CreateStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStorageParamsWithTimeout creates a new CreateStorageParams object
// with the ability to set a timeout on a request.
func NewCreateStorageParamsWithTimeout(timeout time.Duration) *CreateStorageParams {
	return &CreateStorageParams{
		timeout: timeout,
	}
}

// NewCreateStorageParamsWithContext creates a new CreateStorageParams object
// with the ability to set a context for a request.
func NewCreateStorageParamsWithContext(ctx context.Context) *CreateStorageParams {
	return &CreateStorageParams{
		Context: ctx,
	}
}

// NewCreateStorageParamsWithHTTPClient creates a new CreateStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStorageParamsWithHTTPClient(client *http.Client) *CreateStorageParams {
	return &CreateStorageParams{
		HTTPClient: client,
	}
}

/*
CreateStorageParams contains all the parameters to send to the API endpoint

	for the create storage operation.

	Typically these are written to a http.Request.
*/
type CreateStorageParams struct {

	/* Body.

	   Storage configuration
	*/
	Body *models.StorageCreateRequest

	/* StorageType.

	   Storage type
	*/
	StorageType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStorageParams) WithDefaults() *CreateStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create storage params
func (o *CreateStorageParams) WithTimeout(timeout time.Duration) *CreateStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create storage params
func (o *CreateStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create storage params
func (o *CreateStorageParams) WithContext(ctx context.Context) *CreateStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create storage params
func (o *CreateStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create storage params
func (o *CreateStorageParams) WithHTTPClient(client *http.Client) *CreateStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create storage params
func (o *CreateStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create storage params
func (o *CreateStorageParams) WithBody(body *models.StorageCreateRequest) *CreateStorageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create storage params
func (o *CreateStorageParams) SetBody(body *models.StorageCreateRequest) {
	o.Body = body
}

// WithStorageType adds the storageType to the create storage params
func (o *CreateStorageParams) WithStorageType(storageType string) *CreateStorageParams {
	o.SetStorageType(storageType)
	return o
}

// SetStorageType adds the storageType to the create storage params
func (o *CreateStorageParams) SetStorageType(storageType string) {
	o.StorageType = storageType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param storageType
	if err := r.SetPathParam("storageType", o.StorageType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
