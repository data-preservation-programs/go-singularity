// Code generated by go-swagger; DO NOT EDIT.

package job

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

// NewPostPreparationIDSourceNamePauseScanParams creates a new PostPreparationIDSourceNamePauseScanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPreparationIDSourceNamePauseScanParams() *PostPreparationIDSourceNamePauseScanParams {
	return &PostPreparationIDSourceNamePauseScanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPreparationIDSourceNamePauseScanParamsWithTimeout creates a new PostPreparationIDSourceNamePauseScanParams object
// with the ability to set a timeout on a request.
func NewPostPreparationIDSourceNamePauseScanParamsWithTimeout(timeout time.Duration) *PostPreparationIDSourceNamePauseScanParams {
	return &PostPreparationIDSourceNamePauseScanParams{
		timeout: timeout,
	}
}

// NewPostPreparationIDSourceNamePauseScanParamsWithContext creates a new PostPreparationIDSourceNamePauseScanParams object
// with the ability to set a context for a request.
func NewPostPreparationIDSourceNamePauseScanParamsWithContext(ctx context.Context) *PostPreparationIDSourceNamePauseScanParams {
	return &PostPreparationIDSourceNamePauseScanParams{
		Context: ctx,
	}
}

// NewPostPreparationIDSourceNamePauseScanParamsWithHTTPClient creates a new PostPreparationIDSourceNamePauseScanParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPreparationIDSourceNamePauseScanParamsWithHTTPClient(client *http.Client) *PostPreparationIDSourceNamePauseScanParams {
	return &PostPreparationIDSourceNamePauseScanParams{
		HTTPClient: client,
	}
}

/*
PostPreparationIDSourceNamePauseScanParams contains all the parameters to send to the API endpoint

	for the post preparation ID source name pause scan operation.

	Typically these are written to a http.Request.
*/
type PostPreparationIDSourceNamePauseScanParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID string

	/* Name.

	   Storage ID or name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post preparation ID source name pause scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPreparationIDSourceNamePauseScanParams) WithDefaults() *PostPreparationIDSourceNamePauseScanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post preparation ID source name pause scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPreparationIDSourceNamePauseScanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) WithTimeout(timeout time.Duration) *PostPreparationIDSourceNamePauseScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) WithContext(ctx context.Context) *PostPreparationIDSourceNamePauseScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) WithHTTPClient(client *http.Client) *PostPreparationIDSourceNamePauseScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) WithID(id string) *PostPreparationIDSourceNamePauseScanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) WithName(name string) *PostPreparationIDSourceNamePauseScanParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post preparation ID source name pause scan params
func (o *PostPreparationIDSourceNamePauseScanParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostPreparationIDSourceNamePauseScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
