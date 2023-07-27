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

// NewPostSourceKoofrDatasetDatasetNameParams creates a new PostSourceKoofrDatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceKoofrDatasetDatasetNameParams() *PostSourceKoofrDatasetDatasetNameParams {
	return &PostSourceKoofrDatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceKoofrDatasetDatasetNameParamsWithTimeout creates a new PostSourceKoofrDatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceKoofrDatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceKoofrDatasetDatasetNameParams {
	return &PostSourceKoofrDatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceKoofrDatasetDatasetNameParamsWithContext creates a new PostSourceKoofrDatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceKoofrDatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceKoofrDatasetDatasetNameParams {
	return &PostSourceKoofrDatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceKoofrDatasetDatasetNameParamsWithHTTPClient creates a new PostSourceKoofrDatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceKoofrDatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceKoofrDatasetDatasetNameParams {
	return &PostSourceKoofrDatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceKoofrDatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source koofr dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceKoofrDatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceKoofrRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source koofr dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceKoofrDatasetDatasetNameParams) WithDefaults() *PostSourceKoofrDatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source koofr dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceKoofrDatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceKoofrDatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceKoofrDatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceKoofrDatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceKoofrDatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) WithRequest(request *models.DatasourceKoofrRequest) *PostSourceKoofrDatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source koofr dataset dataset name params
func (o *PostSourceKoofrDatasetDatasetNameParams) SetRequest(request *models.DatasourceKoofrRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceKoofrDatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
