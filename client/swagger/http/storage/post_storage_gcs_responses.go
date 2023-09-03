// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostStorageGcsReader is a Reader for the PostStorageGcs structure.
type PostStorageGcsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageGcsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageGcsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageGcsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageGcsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/gcs] PostStorageGcs", response, response.Code())
	}
}

// NewPostStorageGcsOK creates a PostStorageGcsOK with default headers values
func NewPostStorageGcsOK() *PostStorageGcsOK {
	return &PostStorageGcsOK{}
}

/*
PostStorageGcsOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageGcsOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage gcs o k response has a 2xx status code
func (o *PostStorageGcsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage gcs o k response has a 3xx status code
func (o *PostStorageGcsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage gcs o k response has a 4xx status code
func (o *PostStorageGcsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage gcs o k response has a 5xx status code
func (o *PostStorageGcsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage gcs o k response a status code equal to that given
func (o *PostStorageGcsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage gcs o k response
func (o *PostStorageGcsOK) Code() int {
	return 200
}

func (o *PostStorageGcsOK) Error() string {
	return fmt.Sprintf("[POST /storage/gcs][%d] postStorageGcsOK  %+v", 200, o.Payload)
}

func (o *PostStorageGcsOK) String() string {
	return fmt.Sprintf("[POST /storage/gcs][%d] postStorageGcsOK  %+v", 200, o.Payload)
}

func (o *PostStorageGcsOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageGcsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageGcsBadRequest creates a PostStorageGcsBadRequest with default headers values
func NewPostStorageGcsBadRequest() *PostStorageGcsBadRequest {
	return &PostStorageGcsBadRequest{}
}

/*
PostStorageGcsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageGcsBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage gcs bad request response has a 2xx status code
func (o *PostStorageGcsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage gcs bad request response has a 3xx status code
func (o *PostStorageGcsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage gcs bad request response has a 4xx status code
func (o *PostStorageGcsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage gcs bad request response has a 5xx status code
func (o *PostStorageGcsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage gcs bad request response a status code equal to that given
func (o *PostStorageGcsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage gcs bad request response
func (o *PostStorageGcsBadRequest) Code() int {
	return 400
}

func (o *PostStorageGcsBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/gcs][%d] postStorageGcsBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageGcsBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/gcs][%d] postStorageGcsBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageGcsBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageGcsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageGcsInternalServerError creates a PostStorageGcsInternalServerError with default headers values
func NewPostStorageGcsInternalServerError() *PostStorageGcsInternalServerError {
	return &PostStorageGcsInternalServerError{}
}

/*
PostStorageGcsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageGcsInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage gcs internal server error response has a 2xx status code
func (o *PostStorageGcsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage gcs internal server error response has a 3xx status code
func (o *PostStorageGcsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage gcs internal server error response has a 4xx status code
func (o *PostStorageGcsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage gcs internal server error response has a 5xx status code
func (o *PostStorageGcsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage gcs internal server error response a status code equal to that given
func (o *PostStorageGcsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage gcs internal server error response
func (o *PostStorageGcsInternalServerError) Code() int {
	return 500
}

func (o *PostStorageGcsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/gcs][%d] postStorageGcsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageGcsInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/gcs][%d] postStorageGcsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageGcsInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageGcsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}