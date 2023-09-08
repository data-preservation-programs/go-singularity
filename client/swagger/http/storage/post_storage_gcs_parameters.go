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

// NewPostStorageGcsParams creates a new PostStorageGcsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageGcsParams() *PostStorageGcsParams {
	return &PostStorageGcsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageGcsParamsWithTimeout creates a new PostStorageGcsParams object
// with the ability to set a timeout on a request.
func NewPostStorageGcsParamsWithTimeout(timeout time.Duration) *PostStorageGcsParams {
	return &PostStorageGcsParams{
		timeout: timeout,
	}
}

// NewPostStorageGcsParamsWithContext creates a new PostStorageGcsParams object
// with the ability to set a context for a request.
func NewPostStorageGcsParamsWithContext(ctx context.Context) *PostStorageGcsParams {
	return &PostStorageGcsParams{
		Context: ctx,
	}
}

// NewPostStorageGcsParamsWithHTTPClient creates a new PostStorageGcsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageGcsParamsWithHTTPClient(client *http.Client) *PostStorageGcsParams {
	return &PostStorageGcsParams{
		HTTPClient: client,
	}
}

/*
PostStorageGcsParams contains all the parameters to send to the API endpoint

	for the post storage gcs operation.

	Typically these are written to a http.Request.
*/
type PostStorageGcsParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateGcsStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage gcs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageGcsParams) WithDefaults() *PostStorageGcsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage gcs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageGcsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage gcs params
func (o *PostStorageGcsParams) WithTimeout(timeout time.Duration) *PostStorageGcsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage gcs params
func (o *PostStorageGcsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage gcs params
func (o *PostStorageGcsParams) WithContext(ctx context.Context) *PostStorageGcsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage gcs params
func (o *PostStorageGcsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage gcs params
func (o *PostStorageGcsParams) WithHTTPClient(client *http.Client) *PostStorageGcsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage gcs params
func (o *PostStorageGcsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage gcs params
func (o *PostStorageGcsParams) WithRequest(request *models.StorageCreateGcsStorageRequest) *PostStorageGcsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage gcs params
func (o *PostStorageGcsParams) SetRequest(request *models.StorageCreateGcsStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageGcsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
