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

// NewPostStorageS3IbmcosParams creates a new PostStorageS3IbmcosParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageS3IbmcosParams() *PostStorageS3IbmcosParams {
	return &PostStorageS3IbmcosParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageS3IbmcosParamsWithTimeout creates a new PostStorageS3IbmcosParams object
// with the ability to set a timeout on a request.
func NewPostStorageS3IbmcosParamsWithTimeout(timeout time.Duration) *PostStorageS3IbmcosParams {
	return &PostStorageS3IbmcosParams{
		timeout: timeout,
	}
}

// NewPostStorageS3IbmcosParamsWithContext creates a new PostStorageS3IbmcosParams object
// with the ability to set a context for a request.
func NewPostStorageS3IbmcosParamsWithContext(ctx context.Context) *PostStorageS3IbmcosParams {
	return &PostStorageS3IbmcosParams{
		Context: ctx,
	}
}

// NewPostStorageS3IbmcosParamsWithHTTPClient creates a new PostStorageS3IbmcosParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageS3IbmcosParamsWithHTTPClient(client *http.Client) *PostStorageS3IbmcosParams {
	return &PostStorageS3IbmcosParams{
		HTTPClient: client,
	}
}

/*
PostStorageS3IbmcosParams contains all the parameters to send to the API endpoint

	for the post storage s3 ibmcos operation.

	Typically these are written to a http.Request.
*/
type PostStorageS3IbmcosParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateS3IBMCOSStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage s3 ibmcos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3IbmcosParams) WithDefaults() *PostStorageS3IbmcosParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage s3 ibmcos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3IbmcosParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) WithTimeout(timeout time.Duration) *PostStorageS3IbmcosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) WithContext(ctx context.Context) *PostStorageS3IbmcosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) WithHTTPClient(client *http.Client) *PostStorageS3IbmcosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) WithRequest(request *models.StorageCreateS3IBMCOSStorageRequest) *PostStorageS3IbmcosParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage s3 ibmcos params
func (o *PostStorageS3IbmcosParams) SetRequest(request *models.StorageCreateS3IBMCOSStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageS3IbmcosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
