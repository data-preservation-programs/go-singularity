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

// NewPostStorageSmbParams creates a new PostStorageSmbParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageSmbParams() *PostStorageSmbParams {
	return &PostStorageSmbParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageSmbParamsWithTimeout creates a new PostStorageSmbParams object
// with the ability to set a timeout on a request.
func NewPostStorageSmbParamsWithTimeout(timeout time.Duration) *PostStorageSmbParams {
	return &PostStorageSmbParams{
		timeout: timeout,
	}
}

// NewPostStorageSmbParamsWithContext creates a new PostStorageSmbParams object
// with the ability to set a context for a request.
func NewPostStorageSmbParamsWithContext(ctx context.Context) *PostStorageSmbParams {
	return &PostStorageSmbParams{
		Context: ctx,
	}
}

// NewPostStorageSmbParamsWithHTTPClient creates a new PostStorageSmbParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageSmbParamsWithHTTPClient(client *http.Client) *PostStorageSmbParams {
	return &PostStorageSmbParams{
		HTTPClient: client,
	}
}

/*
PostStorageSmbParams contains all the parameters to send to the API endpoint

	for the post storage smb operation.

	Typically these are written to a http.Request.
*/
type PostStorageSmbParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateSmbStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage smb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageSmbParams) WithDefaults() *PostStorageSmbParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage smb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageSmbParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage smb params
func (o *PostStorageSmbParams) WithTimeout(timeout time.Duration) *PostStorageSmbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage smb params
func (o *PostStorageSmbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage smb params
func (o *PostStorageSmbParams) WithContext(ctx context.Context) *PostStorageSmbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage smb params
func (o *PostStorageSmbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage smb params
func (o *PostStorageSmbParams) WithHTTPClient(client *http.Client) *PostStorageSmbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage smb params
func (o *PostStorageSmbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage smb params
func (o *PostStorageSmbParams) WithRequest(request *models.StorageCreateSmbStorageRequest) *PostStorageSmbParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage smb params
func (o *PostStorageSmbParams) SetRequest(request *models.StorageCreateSmbStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageSmbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
