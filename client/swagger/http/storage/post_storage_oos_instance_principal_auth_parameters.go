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

// NewPostStorageOosInstancePrincipalAuthParams creates a new PostStorageOosInstancePrincipalAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageOosInstancePrincipalAuthParams() *PostStorageOosInstancePrincipalAuthParams {
	return &PostStorageOosInstancePrincipalAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageOosInstancePrincipalAuthParamsWithTimeout creates a new PostStorageOosInstancePrincipalAuthParams object
// with the ability to set a timeout on a request.
func NewPostStorageOosInstancePrincipalAuthParamsWithTimeout(timeout time.Duration) *PostStorageOosInstancePrincipalAuthParams {
	return &PostStorageOosInstancePrincipalAuthParams{
		timeout: timeout,
	}
}

// NewPostStorageOosInstancePrincipalAuthParamsWithContext creates a new PostStorageOosInstancePrincipalAuthParams object
// with the ability to set a context for a request.
func NewPostStorageOosInstancePrincipalAuthParamsWithContext(ctx context.Context) *PostStorageOosInstancePrincipalAuthParams {
	return &PostStorageOosInstancePrincipalAuthParams{
		Context: ctx,
	}
}

// NewPostStorageOosInstancePrincipalAuthParamsWithHTTPClient creates a new PostStorageOosInstancePrincipalAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageOosInstancePrincipalAuthParamsWithHTTPClient(client *http.Client) *PostStorageOosInstancePrincipalAuthParams {
	return &PostStorageOosInstancePrincipalAuthParams{
		HTTPClient: client,
	}
}

/*
PostStorageOosInstancePrincipalAuthParams contains all the parameters to send to the API endpoint

	for the post storage oos instance principal auth operation.

	Typically these are written to a http.Request.
*/
type PostStorageOosInstancePrincipalAuthParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateOosInstancePrincipalAuthStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage oos instance principal auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageOosInstancePrincipalAuthParams) WithDefaults() *PostStorageOosInstancePrincipalAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage oos instance principal auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageOosInstancePrincipalAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) WithTimeout(timeout time.Duration) *PostStorageOosInstancePrincipalAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) WithContext(ctx context.Context) *PostStorageOosInstancePrincipalAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) WithHTTPClient(client *http.Client) *PostStorageOosInstancePrincipalAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) WithRequest(request *models.StorageCreateOosInstancePrincipalAuthStorageRequest) *PostStorageOosInstancePrincipalAuthParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage oos instance principal auth params
func (o *PostStorageOosInstancePrincipalAuthParams) SetRequest(request *models.StorageCreateOosInstancePrincipalAuthStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageOosInstancePrincipalAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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