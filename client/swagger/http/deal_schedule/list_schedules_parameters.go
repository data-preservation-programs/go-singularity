// Code generated by go-swagger; DO NOT EDIT.

package deal_schedule

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

// NewListSchedulesParams creates a new ListSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSchedulesParams() *ListSchedulesParams {
	return &ListSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSchedulesParamsWithTimeout creates a new ListSchedulesParams object
// with the ability to set a timeout on a request.
func NewListSchedulesParamsWithTimeout(timeout time.Duration) *ListSchedulesParams {
	return &ListSchedulesParams{
		timeout: timeout,
	}
}

// NewListSchedulesParamsWithContext creates a new ListSchedulesParams object
// with the ability to set a context for a request.
func NewListSchedulesParamsWithContext(ctx context.Context) *ListSchedulesParams {
	return &ListSchedulesParams{
		Context: ctx,
	}
}

// NewListSchedulesParamsWithHTTPClient creates a new ListSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSchedulesParamsWithHTTPClient(client *http.Client) *ListSchedulesParams {
	return &ListSchedulesParams{
		HTTPClient: client,
	}
}

/*
ListSchedulesParams contains all the parameters to send to the API endpoint

	for the list schedules operation.

	Typically these are written to a http.Request.
*/
type ListSchedulesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSchedulesParams) WithDefaults() *ListSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list schedules params
func (o *ListSchedulesParams) WithTimeout(timeout time.Duration) *ListSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list schedules params
func (o *ListSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list schedules params
func (o *ListSchedulesParams) WithContext(ctx context.Context) *ListSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list schedules params
func (o *ListSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list schedules params
func (o *ListSchedulesParams) WithHTTPClient(client *http.Client) *ListSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list schedules params
func (o *ListSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
