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

// NewGetPreparationIDSourceNameExplorePathParams creates a new GetPreparationIDSourceNameExplorePathParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPreparationIDSourceNameExplorePathParams() *GetPreparationIDSourceNameExplorePathParams {
	return &GetPreparationIDSourceNameExplorePathParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPreparationIDSourceNameExplorePathParamsWithTimeout creates a new GetPreparationIDSourceNameExplorePathParams object
// with the ability to set a timeout on a request.
func NewGetPreparationIDSourceNameExplorePathParamsWithTimeout(timeout time.Duration) *GetPreparationIDSourceNameExplorePathParams {
	return &GetPreparationIDSourceNameExplorePathParams{
		timeout: timeout,
	}
}

// NewGetPreparationIDSourceNameExplorePathParamsWithContext creates a new GetPreparationIDSourceNameExplorePathParams object
// with the ability to set a context for a request.
func NewGetPreparationIDSourceNameExplorePathParamsWithContext(ctx context.Context) *GetPreparationIDSourceNameExplorePathParams {
	return &GetPreparationIDSourceNameExplorePathParams{
		Context: ctx,
	}
}

// NewGetPreparationIDSourceNameExplorePathParamsWithHTTPClient creates a new GetPreparationIDSourceNameExplorePathParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPreparationIDSourceNameExplorePathParamsWithHTTPClient(client *http.Client) *GetPreparationIDSourceNameExplorePathParams {
	return &GetPreparationIDSourceNameExplorePathParams{
		HTTPClient: client,
	}
}

/*
GetPreparationIDSourceNameExplorePathParams contains all the parameters to send to the API endpoint

	for the get preparation ID source name explore path operation.

	Typically these are written to a http.Request.
*/
type GetPreparationIDSourceNameExplorePathParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID string

	/* Name.

	   Source storage ID or name
	*/
	Name string

	/* Path.

	   Directory path
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get preparation ID source name explore path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPreparationIDSourceNameExplorePathParams) WithDefaults() *GetPreparationIDSourceNameExplorePathParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get preparation ID source name explore path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPreparationIDSourceNameExplorePathParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) WithTimeout(timeout time.Duration) *GetPreparationIDSourceNameExplorePathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) WithContext(ctx context.Context) *GetPreparationIDSourceNameExplorePathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) WithHTTPClient(client *http.Client) *GetPreparationIDSourceNameExplorePathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) WithID(id string) *GetPreparationIDSourceNameExplorePathParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) WithName(name string) *GetPreparationIDSourceNameExplorePathParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) SetName(name string) {
	o.Name = name
}

// WithPath adds the path to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) WithPath(path string) *GetPreparationIDSourceNameExplorePathParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get preparation ID source name explore path params
func (o *GetPreparationIDSourceNameExplorePathParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *GetPreparationIDSourceNameExplorePathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
