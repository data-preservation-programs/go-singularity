// Code generated by go-swagger; DO NOT EDIT.

package preparation

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

// NewPostPreparationIDOutputNameParams creates a new PostPreparationIDOutputNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPreparationIDOutputNameParams() *PostPreparationIDOutputNameParams {
	return &PostPreparationIDOutputNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPreparationIDOutputNameParamsWithTimeout creates a new PostPreparationIDOutputNameParams object
// with the ability to set a timeout on a request.
func NewPostPreparationIDOutputNameParamsWithTimeout(timeout time.Duration) *PostPreparationIDOutputNameParams {
	return &PostPreparationIDOutputNameParams{
		timeout: timeout,
	}
}

// NewPostPreparationIDOutputNameParamsWithContext creates a new PostPreparationIDOutputNameParams object
// with the ability to set a context for a request.
func NewPostPreparationIDOutputNameParamsWithContext(ctx context.Context) *PostPreparationIDOutputNameParams {
	return &PostPreparationIDOutputNameParams{
		Context: ctx,
	}
}

// NewPostPreparationIDOutputNameParamsWithHTTPClient creates a new PostPreparationIDOutputNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPreparationIDOutputNameParamsWithHTTPClient(client *http.Client) *PostPreparationIDOutputNameParams {
	return &PostPreparationIDOutputNameParams{
		HTTPClient: client,
	}
}

/*
PostPreparationIDOutputNameParams contains all the parameters to send to the API endpoint

	for the post preparation ID output name operation.

	Typically these are written to a http.Request.
*/
type PostPreparationIDOutputNameParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID string

	/* Name.

	   Output storage ID or name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post preparation ID output name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPreparationIDOutputNameParams) WithDefaults() *PostPreparationIDOutputNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post preparation ID output name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPreparationIDOutputNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) WithTimeout(timeout time.Duration) *PostPreparationIDOutputNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) WithContext(ctx context.Context) *PostPreparationIDOutputNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) WithHTTPClient(client *http.Client) *PostPreparationIDOutputNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) WithID(id string) *PostPreparationIDOutputNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) WithName(name string) *PostPreparationIDOutputNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post preparation ID output name params
func (o *PostPreparationIDOutputNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostPreparationIDOutputNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
