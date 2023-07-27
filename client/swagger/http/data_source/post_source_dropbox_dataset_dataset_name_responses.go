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

// PostSourceDropboxDatasetDatasetNameReader is a Reader for the PostSourceDropboxDatasetDatasetName structure.
type PostSourceDropboxDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceDropboxDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceDropboxDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceDropboxDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceDropboxDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/dropbox/dataset/{datasetName}] PostSourceDropboxDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceDropboxDatasetDatasetNameOK creates a PostSourceDropboxDatasetDatasetNameOK with default headers values
func NewPostSourceDropboxDatasetDatasetNameOK() *PostSourceDropboxDatasetDatasetNameOK {
	return &PostSourceDropboxDatasetDatasetNameOK{}
}

/*
PostSourceDropboxDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceDropboxDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source dropbox dataset dataset name o k response has a 2xx status code
func (o *PostSourceDropboxDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source dropbox dataset dataset name o k response has a 3xx status code
func (o *PostSourceDropboxDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source dropbox dataset dataset name o k response has a 4xx status code
func (o *PostSourceDropboxDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source dropbox dataset dataset name o k response has a 5xx status code
func (o *PostSourceDropboxDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source dropbox dataset dataset name o k response a status code equal to that given
func (o *PostSourceDropboxDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source dropbox dataset dataset name o k response
func (o *PostSourceDropboxDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceDropboxDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/dropbox/dataset/{datasetName}][%d] postSourceDropboxDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceDropboxDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/dropbox/dataset/{datasetName}][%d] postSourceDropboxDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceDropboxDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceDropboxDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceDropboxDatasetDatasetNameBadRequest creates a PostSourceDropboxDatasetDatasetNameBadRequest with default headers values
func NewPostSourceDropboxDatasetDatasetNameBadRequest() *PostSourceDropboxDatasetDatasetNameBadRequest {
	return &PostSourceDropboxDatasetDatasetNameBadRequest{}
}

/*
PostSourceDropboxDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceDropboxDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source dropbox dataset dataset name bad request response has a 2xx status code
func (o *PostSourceDropboxDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source dropbox dataset dataset name bad request response has a 3xx status code
func (o *PostSourceDropboxDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source dropbox dataset dataset name bad request response has a 4xx status code
func (o *PostSourceDropboxDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source dropbox dataset dataset name bad request response has a 5xx status code
func (o *PostSourceDropboxDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source dropbox dataset dataset name bad request response a status code equal to that given
func (o *PostSourceDropboxDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source dropbox dataset dataset name bad request response
func (o *PostSourceDropboxDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceDropboxDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/dropbox/dataset/{datasetName}][%d] postSourceDropboxDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceDropboxDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/dropbox/dataset/{datasetName}][%d] postSourceDropboxDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceDropboxDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceDropboxDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceDropboxDatasetDatasetNameInternalServerError creates a PostSourceDropboxDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceDropboxDatasetDatasetNameInternalServerError() *PostSourceDropboxDatasetDatasetNameInternalServerError {
	return &PostSourceDropboxDatasetDatasetNameInternalServerError{}
}

/*
PostSourceDropboxDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceDropboxDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source dropbox dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source dropbox dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source dropbox dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source dropbox dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source dropbox dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source dropbox dataset dataset name internal server error response
func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/dropbox/dataset/{datasetName}][%d] postSourceDropboxDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/dropbox/dataset/{datasetName}][%d] postSourceDropboxDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceDropboxDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
