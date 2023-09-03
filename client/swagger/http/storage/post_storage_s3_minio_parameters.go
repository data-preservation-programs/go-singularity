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

// NewPostStorageS3MinioParams creates a new PostStorageS3MinioParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageS3MinioParams() *PostStorageS3MinioParams {
	return &PostStorageS3MinioParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageS3MinioParamsWithTimeout creates a new PostStorageS3MinioParams object
// with the ability to set a timeout on a request.
func NewPostStorageS3MinioParamsWithTimeout(timeout time.Duration) *PostStorageS3MinioParams {
	return &PostStorageS3MinioParams{
		timeout: timeout,
	}
}

// NewPostStorageS3MinioParamsWithContext creates a new PostStorageS3MinioParams object
// with the ability to set a context for a request.
func NewPostStorageS3MinioParamsWithContext(ctx context.Context) *PostStorageS3MinioParams {
	return &PostStorageS3MinioParams{
		Context: ctx,
	}
}

// NewPostStorageS3MinioParamsWithHTTPClient creates a new PostStorageS3MinioParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageS3MinioParamsWithHTTPClient(client *http.Client) *PostStorageS3MinioParams {
	return &PostStorageS3MinioParams{
		HTTPClient: client,
	}
}

/*
PostStorageS3MinioParams contains all the parameters to send to the API endpoint

	for the post storage s3 minio operation.

	Typically these are written to a http.Request.
*/
type PostStorageS3MinioParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateS3MinioStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage s3 minio params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3MinioParams) WithDefaults() *PostStorageS3MinioParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage s3 minio params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3MinioParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage s3 minio params
func (o *PostStorageS3MinioParams) WithTimeout(timeout time.Duration) *PostStorageS3MinioParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage s3 minio params
func (o *PostStorageS3MinioParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage s3 minio params
func (o *PostStorageS3MinioParams) WithContext(ctx context.Context) *PostStorageS3MinioParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage s3 minio params
func (o *PostStorageS3MinioParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage s3 minio params
func (o *PostStorageS3MinioParams) WithHTTPClient(client *http.Client) *PostStorageS3MinioParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage s3 minio params
func (o *PostStorageS3MinioParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage s3 minio params
func (o *PostStorageS3MinioParams) WithRequest(request *models.StorageCreateS3MinioStorageRequest) *PostStorageS3MinioParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage s3 minio params
func (o *PostStorageS3MinioParams) SetRequest(request *models.StorageCreateS3MinioStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageS3MinioParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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