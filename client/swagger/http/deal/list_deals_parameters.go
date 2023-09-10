// Code generated by go-swagger; DO NOT EDIT.

package deal

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

// NewListDealsParams creates a new ListDealsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDealsParams() *ListDealsParams {
	return &ListDealsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDealsParamsWithTimeout creates a new ListDealsParams object
// with the ability to set a timeout on a request.
func NewListDealsParamsWithTimeout(timeout time.Duration) *ListDealsParams {
	return &ListDealsParams{
		timeout: timeout,
	}
}

// NewListDealsParamsWithContext creates a new ListDealsParams object
// with the ability to set a context for a request.
func NewListDealsParamsWithContext(ctx context.Context) *ListDealsParams {
	return &ListDealsParams{
		Context: ctx,
	}
}

// NewListDealsParamsWithHTTPClient creates a new ListDealsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDealsParamsWithHTTPClient(client *http.Client) *ListDealsParams {
	return &ListDealsParams{
		HTTPClient: client,
	}
}

/*
ListDealsParams contains all the parameters to send to the API endpoint

	for the list deals operation.

	Typically these are written to a http.Request.
*/
type ListDealsParams struct {

	/* Request.

	   ListDealRequest
	*/
	Request *models.DealListDealRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list deals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDealsParams) WithDefaults() *ListDealsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list deals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDealsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list deals params
func (o *ListDealsParams) WithTimeout(timeout time.Duration) *ListDealsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list deals params
func (o *ListDealsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list deals params
func (o *ListDealsParams) WithContext(ctx context.Context) *ListDealsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list deals params
func (o *ListDealsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list deals params
func (o *ListDealsParams) WithHTTPClient(client *http.Client) *ListDealsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list deals params
func (o *ListDealsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the list deals params
func (o *ListDealsParams) WithRequest(request *models.DealListDealRequest) *ListDealsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the list deals params
func (o *ListDealsParams) SetRequest(request *models.DealListDealRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ListDealsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
