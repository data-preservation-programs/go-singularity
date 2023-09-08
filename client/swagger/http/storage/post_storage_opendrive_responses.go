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

// PostStorageOpendriveReader is a Reader for the PostStorageOpendrive structure.
type PostStorageOpendriveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageOpendriveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageOpendriveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageOpendriveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageOpendriveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/opendrive] PostStorageOpendrive", response, response.Code())
	}
}

// NewPostStorageOpendriveOK creates a PostStorageOpendriveOK with default headers values
func NewPostStorageOpendriveOK() *PostStorageOpendriveOK {
	return &PostStorageOpendriveOK{}
}

/*
PostStorageOpendriveOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageOpendriveOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage opendrive o k response has a 2xx status code
func (o *PostStorageOpendriveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage opendrive o k response has a 3xx status code
func (o *PostStorageOpendriveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage opendrive o k response has a 4xx status code
func (o *PostStorageOpendriveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage opendrive o k response has a 5xx status code
func (o *PostStorageOpendriveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage opendrive o k response a status code equal to that given
func (o *PostStorageOpendriveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage opendrive o k response
func (o *PostStorageOpendriveOK) Code() int {
	return 200
}

func (o *PostStorageOpendriveOK) Error() string {
	return fmt.Sprintf("[POST /storage/opendrive][%d] postStorageOpendriveOK  %+v", 200, o.Payload)
}

func (o *PostStorageOpendriveOK) String() string {
	return fmt.Sprintf("[POST /storage/opendrive][%d] postStorageOpendriveOK  %+v", 200, o.Payload)
}

func (o *PostStorageOpendriveOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageOpendriveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageOpendriveBadRequest creates a PostStorageOpendriveBadRequest with default headers values
func NewPostStorageOpendriveBadRequest() *PostStorageOpendriveBadRequest {
	return &PostStorageOpendriveBadRequest{}
}

/*
PostStorageOpendriveBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageOpendriveBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage opendrive bad request response has a 2xx status code
func (o *PostStorageOpendriveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage opendrive bad request response has a 3xx status code
func (o *PostStorageOpendriveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage opendrive bad request response has a 4xx status code
func (o *PostStorageOpendriveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage opendrive bad request response has a 5xx status code
func (o *PostStorageOpendriveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage opendrive bad request response a status code equal to that given
func (o *PostStorageOpendriveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage opendrive bad request response
func (o *PostStorageOpendriveBadRequest) Code() int {
	return 400
}

func (o *PostStorageOpendriveBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/opendrive][%d] postStorageOpendriveBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageOpendriveBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/opendrive][%d] postStorageOpendriveBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageOpendriveBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageOpendriveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageOpendriveInternalServerError creates a PostStorageOpendriveInternalServerError with default headers values
func NewPostStorageOpendriveInternalServerError() *PostStorageOpendriveInternalServerError {
	return &PostStorageOpendriveInternalServerError{}
}

/*
PostStorageOpendriveInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageOpendriveInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage opendrive internal server error response has a 2xx status code
func (o *PostStorageOpendriveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage opendrive internal server error response has a 3xx status code
func (o *PostStorageOpendriveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage opendrive internal server error response has a 4xx status code
func (o *PostStorageOpendriveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage opendrive internal server error response has a 5xx status code
func (o *PostStorageOpendriveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage opendrive internal server error response a status code equal to that given
func (o *PostStorageOpendriveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage opendrive internal server error response
func (o *PostStorageOpendriveInternalServerError) Code() int {
	return 500
}

func (o *PostStorageOpendriveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/opendrive][%d] postStorageOpendriveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageOpendriveInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/opendrive][%d] postStorageOpendriveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageOpendriveInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageOpendriveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
