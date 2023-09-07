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

// NewCreatePutioStorageParams creates a new CreatePutioStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePutioStorageParams() *CreatePutioStorageParams {
	return &CreatePutioStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePutioStorageParamsWithTimeout creates a new CreatePutioStorageParams object
// with the ability to set a timeout on a request.
func NewCreatePutioStorageParamsWithTimeout(timeout time.Duration) *CreatePutioStorageParams {
	return &CreatePutioStorageParams{
		timeout: timeout,
	}
}

// NewCreatePutioStorageParamsWithContext creates a new CreatePutioStorageParams object
// with the ability to set a context for a request.
func NewCreatePutioStorageParamsWithContext(ctx context.Context) *CreatePutioStorageParams {
	return &CreatePutioStorageParams{
		Context: ctx,
	}
}

// NewCreatePutioStorageParamsWithHTTPClient creates a new CreatePutioStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePutioStorageParamsWithHTTPClient(client *http.Client) *CreatePutioStorageParams {
	return &CreatePutioStorageParams{
		HTTPClient: client,
	}
}

/*
CreatePutioStorageParams contains all the parameters to send to the API endpoint

	for the create putio storage operation.

	Typically these are written to a http.Request.
*/
type CreatePutioStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreatePutioStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create putio storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePutioStorageParams) WithDefaults() *CreatePutioStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create putio storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePutioStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create putio storage params
func (o *CreatePutioStorageParams) WithTimeout(timeout time.Duration) *CreatePutioStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create putio storage params
func (o *CreatePutioStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create putio storage params
func (o *CreatePutioStorageParams) WithContext(ctx context.Context) *CreatePutioStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create putio storage params
func (o *CreatePutioStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create putio storage params
func (o *CreatePutioStorageParams) WithHTTPClient(client *http.Client) *CreatePutioStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create putio storage params
func (o *CreatePutioStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create putio storage params
func (o *CreatePutioStorageParams) WithRequest(request *models.StorageCreatePutioStorageRequest) *CreatePutioStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create putio storage params
func (o *CreatePutioStorageParams) SetRequest(request *models.StorageCreatePutioStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePutioStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
