// Code generated by go-swagger; DO NOT EDIT.

package data_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostSourceSugarsyncDatasetDatasetNameReader is a Reader for the PostSourceSugarsyncDatasetDatasetName structure.
type PostSourceSugarsyncDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceSugarsyncDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceSugarsyncDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceSugarsyncDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceSugarsyncDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/sugarsync/dataset/{datasetName}] PostSourceSugarsyncDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceSugarsyncDatasetDatasetNameOK creates a PostSourceSugarsyncDatasetDatasetNameOK with default headers values
func NewPostSourceSugarsyncDatasetDatasetNameOK() *PostSourceSugarsyncDatasetDatasetNameOK {
	return &PostSourceSugarsyncDatasetDatasetNameOK{}
}

/*
PostSourceSugarsyncDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceSugarsyncDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source sugarsync dataset dataset name o k response has a 2xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source sugarsync dataset dataset name o k response has a 3xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source sugarsync dataset dataset name o k response has a 4xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source sugarsync dataset dataset name o k response has a 5xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source sugarsync dataset dataset name o k response a status code equal to that given
func (o *PostSourceSugarsyncDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source sugarsync dataset dataset name o k response
func (o *PostSourceSugarsyncDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceSugarsyncDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/sugarsync/dataset/{datasetName}][%d] postSourceSugarsyncDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceSugarsyncDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/sugarsync/dataset/{datasetName}][%d] postSourceSugarsyncDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceSugarsyncDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceSugarsyncDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceSugarsyncDatasetDatasetNameBadRequest creates a PostSourceSugarsyncDatasetDatasetNameBadRequest with default headers values
func NewPostSourceSugarsyncDatasetDatasetNameBadRequest() *PostSourceSugarsyncDatasetDatasetNameBadRequest {
	return &PostSourceSugarsyncDatasetDatasetNameBadRequest{}
}

/*
PostSourceSugarsyncDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceSugarsyncDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source sugarsync dataset dataset name bad request response has a 2xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source sugarsync dataset dataset name bad request response has a 3xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source sugarsync dataset dataset name bad request response has a 4xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source sugarsync dataset dataset name bad request response has a 5xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source sugarsync dataset dataset name bad request response a status code equal to that given
func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source sugarsync dataset dataset name bad request response
func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/sugarsync/dataset/{datasetName}][%d] postSourceSugarsyncDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/sugarsync/dataset/{datasetName}][%d] postSourceSugarsyncDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceSugarsyncDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceSugarsyncDatasetDatasetNameInternalServerError creates a PostSourceSugarsyncDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceSugarsyncDatasetDatasetNameInternalServerError() *PostSourceSugarsyncDatasetDatasetNameInternalServerError {
	return &PostSourceSugarsyncDatasetDatasetNameInternalServerError{}
}

/*
PostSourceSugarsyncDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceSugarsyncDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source sugarsync dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source sugarsync dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source sugarsync dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source sugarsync dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source sugarsync dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source sugarsync dataset dataset name internal server error response
func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/sugarsync/dataset/{datasetName}][%d] postSourceSugarsyncDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/sugarsync/dataset/{datasetName}][%d] postSourceSugarsyncDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceSugarsyncDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
