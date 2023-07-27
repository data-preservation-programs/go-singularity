// Code generated by go-swagger; DO NOT EDIT.

package data_source

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

// NewPostSourceNetstorageDatasetDatasetNameParams creates a new PostSourceNetstorageDatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceNetstorageDatasetDatasetNameParams() *PostSourceNetstorageDatasetDatasetNameParams {
	return &PostSourceNetstorageDatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceNetstorageDatasetDatasetNameParamsWithTimeout creates a new PostSourceNetstorageDatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceNetstorageDatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceNetstorageDatasetDatasetNameParams {
	return &PostSourceNetstorageDatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceNetstorageDatasetDatasetNameParamsWithContext creates a new PostSourceNetstorageDatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceNetstorageDatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceNetstorageDatasetDatasetNameParams {
	return &PostSourceNetstorageDatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceNetstorageDatasetDatasetNameParamsWithHTTPClient creates a new PostSourceNetstorageDatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceNetstorageDatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceNetstorageDatasetDatasetNameParams {
	return &PostSourceNetstorageDatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceNetstorageDatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source netstorage dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceNetstorageDatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceNetstorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source netstorage dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceNetstorageDatasetDatasetNameParams) WithDefaults() *PostSourceNetstorageDatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source netstorage dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceNetstorageDatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceNetstorageDatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceNetstorageDatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceNetstorageDatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceNetstorageDatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) WithRequest(request *models.DatasourceNetstorageRequest) *PostSourceNetstorageDatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source netstorage dataset dataset name params
func (o *PostSourceNetstorageDatasetDatasetNameParams) SetRequest(request *models.DatasourceNetstorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceNetstorageDatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasetName
	if err := r.SetPathParam("datasetName", o.DatasetName); err != nil {
		return err
	}
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
