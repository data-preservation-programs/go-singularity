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

// PostSourceUptoboxDatasetDatasetNameReader is a Reader for the PostSourceUptoboxDatasetDatasetName structure.
type PostSourceUptoboxDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceUptoboxDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceUptoboxDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceUptoboxDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceUptoboxDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/uptobox/dataset/{datasetName}] PostSourceUptoboxDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceUptoboxDatasetDatasetNameOK creates a PostSourceUptoboxDatasetDatasetNameOK with default headers values
func NewPostSourceUptoboxDatasetDatasetNameOK() *PostSourceUptoboxDatasetDatasetNameOK {
	return &PostSourceUptoboxDatasetDatasetNameOK{}
}

/*
PostSourceUptoboxDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceUptoboxDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source uptobox dataset dataset name o k response has a 2xx status code
func (o *PostSourceUptoboxDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source uptobox dataset dataset name o k response has a 3xx status code
func (o *PostSourceUptoboxDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source uptobox dataset dataset name o k response has a 4xx status code
func (o *PostSourceUptoboxDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source uptobox dataset dataset name o k response has a 5xx status code
func (o *PostSourceUptoboxDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source uptobox dataset dataset name o k response a status code equal to that given
func (o *PostSourceUptoboxDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source uptobox dataset dataset name o k response
func (o *PostSourceUptoboxDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceUptoboxDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/uptobox/dataset/{datasetName}][%d] postSourceUptoboxDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceUptoboxDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/uptobox/dataset/{datasetName}][%d] postSourceUptoboxDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceUptoboxDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceUptoboxDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceUptoboxDatasetDatasetNameBadRequest creates a PostSourceUptoboxDatasetDatasetNameBadRequest with default headers values
func NewPostSourceUptoboxDatasetDatasetNameBadRequest() *PostSourceUptoboxDatasetDatasetNameBadRequest {
	return &PostSourceUptoboxDatasetDatasetNameBadRequest{}
}

/*
PostSourceUptoboxDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceUptoboxDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source uptobox dataset dataset name bad request response has a 2xx status code
func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source uptobox dataset dataset name bad request response has a 3xx status code
func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source uptobox dataset dataset name bad request response has a 4xx status code
func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source uptobox dataset dataset name bad request response has a 5xx status code
func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source uptobox dataset dataset name bad request response a status code equal to that given
func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source uptobox dataset dataset name bad request response
func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/uptobox/dataset/{datasetName}][%d] postSourceUptoboxDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/uptobox/dataset/{datasetName}][%d] postSourceUptoboxDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceUptoboxDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceUptoboxDatasetDatasetNameInternalServerError creates a PostSourceUptoboxDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceUptoboxDatasetDatasetNameInternalServerError() *PostSourceUptoboxDatasetDatasetNameInternalServerError {
	return &PostSourceUptoboxDatasetDatasetNameInternalServerError{}
}

/*
PostSourceUptoboxDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceUptoboxDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source uptobox dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source uptobox dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source uptobox dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source uptobox dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source uptobox dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source uptobox dataset dataset name internal server error response
func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/uptobox/dataset/{datasetName}][%d] postSourceUptoboxDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/uptobox/dataset/{datasetName}][%d] postSourceUptoboxDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceUptoboxDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
