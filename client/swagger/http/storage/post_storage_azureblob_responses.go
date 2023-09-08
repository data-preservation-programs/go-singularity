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

// PostStorageAzureblobReader is a Reader for the PostStorageAzureblob structure.
type PostStorageAzureblobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageAzureblobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageAzureblobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageAzureblobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageAzureblobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/azureblob] PostStorageAzureblob", response, response.Code())
	}
}

// NewPostStorageAzureblobOK creates a PostStorageAzureblobOK with default headers values
func NewPostStorageAzureblobOK() *PostStorageAzureblobOK {
	return &PostStorageAzureblobOK{}
}

/*
PostStorageAzureblobOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageAzureblobOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage azureblob o k response has a 2xx status code
func (o *PostStorageAzureblobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage azureblob o k response has a 3xx status code
func (o *PostStorageAzureblobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage azureblob o k response has a 4xx status code
func (o *PostStorageAzureblobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage azureblob o k response has a 5xx status code
func (o *PostStorageAzureblobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage azureblob o k response a status code equal to that given
func (o *PostStorageAzureblobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage azureblob o k response
func (o *PostStorageAzureblobOK) Code() int {
	return 200
}

func (o *PostStorageAzureblobOK) Error() string {
	return fmt.Sprintf("[POST /storage/azureblob][%d] postStorageAzureblobOK  %+v", 200, o.Payload)
}

func (o *PostStorageAzureblobOK) String() string {
	return fmt.Sprintf("[POST /storage/azureblob][%d] postStorageAzureblobOK  %+v", 200, o.Payload)
}

func (o *PostStorageAzureblobOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageAzureblobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageAzureblobBadRequest creates a PostStorageAzureblobBadRequest with default headers values
func NewPostStorageAzureblobBadRequest() *PostStorageAzureblobBadRequest {
	return &PostStorageAzureblobBadRequest{}
}

/*
PostStorageAzureblobBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageAzureblobBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage azureblob bad request response has a 2xx status code
func (o *PostStorageAzureblobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage azureblob bad request response has a 3xx status code
func (o *PostStorageAzureblobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage azureblob bad request response has a 4xx status code
func (o *PostStorageAzureblobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage azureblob bad request response has a 5xx status code
func (o *PostStorageAzureblobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage azureblob bad request response a status code equal to that given
func (o *PostStorageAzureblobBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage azureblob bad request response
func (o *PostStorageAzureblobBadRequest) Code() int {
	return 400
}

func (o *PostStorageAzureblobBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/azureblob][%d] postStorageAzureblobBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageAzureblobBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/azureblob][%d] postStorageAzureblobBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageAzureblobBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageAzureblobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageAzureblobInternalServerError creates a PostStorageAzureblobInternalServerError with default headers values
func NewPostStorageAzureblobInternalServerError() *PostStorageAzureblobInternalServerError {
	return &PostStorageAzureblobInternalServerError{}
}

/*
PostStorageAzureblobInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageAzureblobInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage azureblob internal server error response has a 2xx status code
func (o *PostStorageAzureblobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage azureblob internal server error response has a 3xx status code
func (o *PostStorageAzureblobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage azureblob internal server error response has a 4xx status code
func (o *PostStorageAzureblobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage azureblob internal server error response has a 5xx status code
func (o *PostStorageAzureblobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage azureblob internal server error response a status code equal to that given
func (o *PostStorageAzureblobInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage azureblob internal server error response
func (o *PostStorageAzureblobInternalServerError) Code() int {
	return 500
}

func (o *PostStorageAzureblobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/azureblob][%d] postStorageAzureblobInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageAzureblobInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/azureblob][%d] postStorageAzureblobInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageAzureblobInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageAzureblobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
