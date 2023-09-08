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

// PostStorageS3OtherReader is a Reader for the PostStorageS3Other structure.
type PostStorageS3OtherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageS3OtherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageS3OtherOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageS3OtherBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageS3OtherInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/other] PostStorageS3Other", response, response.Code())
	}
}

// NewPostStorageS3OtherOK creates a PostStorageS3OtherOK with default headers values
func NewPostStorageS3OtherOK() *PostStorageS3OtherOK {
	return &PostStorageS3OtherOK{}
}

/*
PostStorageS3OtherOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageS3OtherOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage s3 other o k response has a 2xx status code
func (o *PostStorageS3OtherOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage s3 other o k response has a 3xx status code
func (o *PostStorageS3OtherOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 other o k response has a 4xx status code
func (o *PostStorageS3OtherOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 other o k response has a 5xx status code
func (o *PostStorageS3OtherOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 other o k response a status code equal to that given
func (o *PostStorageS3OtherOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage s3 other o k response
func (o *PostStorageS3OtherOK) Code() int {
	return 200
}

func (o *PostStorageS3OtherOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/other][%d] postStorageS3OtherOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3OtherOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/other][%d] postStorageS3OtherOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3OtherOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageS3OtherOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3OtherBadRequest creates a PostStorageS3OtherBadRequest with default headers values
func NewPostStorageS3OtherBadRequest() *PostStorageS3OtherBadRequest {
	return &PostStorageS3OtherBadRequest{}
}

/*
PostStorageS3OtherBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageS3OtherBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 other bad request response has a 2xx status code
func (o *PostStorageS3OtherBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 other bad request response has a 3xx status code
func (o *PostStorageS3OtherBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 other bad request response has a 4xx status code
func (o *PostStorageS3OtherBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage s3 other bad request response has a 5xx status code
func (o *PostStorageS3OtherBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 other bad request response a status code equal to that given
func (o *PostStorageS3OtherBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage s3 other bad request response
func (o *PostStorageS3OtherBadRequest) Code() int {
	return 400
}

func (o *PostStorageS3OtherBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/other][%d] postStorageS3OtherBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3OtherBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/other][%d] postStorageS3OtherBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3OtherBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3OtherBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3OtherInternalServerError creates a PostStorageS3OtherInternalServerError with default headers values
func NewPostStorageS3OtherInternalServerError() *PostStorageS3OtherInternalServerError {
	return &PostStorageS3OtherInternalServerError{}
}

/*
PostStorageS3OtherInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageS3OtherInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 other internal server error response has a 2xx status code
func (o *PostStorageS3OtherInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 other internal server error response has a 3xx status code
func (o *PostStorageS3OtherInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 other internal server error response has a 4xx status code
func (o *PostStorageS3OtherInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 other internal server error response has a 5xx status code
func (o *PostStorageS3OtherInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage s3 other internal server error response a status code equal to that given
func (o *PostStorageS3OtherInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage s3 other internal server error response
func (o *PostStorageS3OtherInternalServerError) Code() int {
	return 500
}

func (o *PostStorageS3OtherInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/other][%d] postStorageS3OtherInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3OtherInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/other][%d] postStorageS3OtherInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3OtherInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3OtherInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
