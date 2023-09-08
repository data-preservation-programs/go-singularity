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

// PostStorageQingstorReader is a Reader for the PostStorageQingstor structure.
type PostStorageQingstorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageQingstorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageQingstorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageQingstorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageQingstorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/qingstor] PostStorageQingstor", response, response.Code())
	}
}

// NewPostStorageQingstorOK creates a PostStorageQingstorOK with default headers values
func NewPostStorageQingstorOK() *PostStorageQingstorOK {
	return &PostStorageQingstorOK{}
}

/*
PostStorageQingstorOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageQingstorOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage qingstor o k response has a 2xx status code
func (o *PostStorageQingstorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage qingstor o k response has a 3xx status code
func (o *PostStorageQingstorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage qingstor o k response has a 4xx status code
func (o *PostStorageQingstorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage qingstor o k response has a 5xx status code
func (o *PostStorageQingstorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage qingstor o k response a status code equal to that given
func (o *PostStorageQingstorOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage qingstor o k response
func (o *PostStorageQingstorOK) Code() int {
	return 200
}

func (o *PostStorageQingstorOK) Error() string {
	return fmt.Sprintf("[POST /storage/qingstor][%d] postStorageQingstorOK  %+v", 200, o.Payload)
}

func (o *PostStorageQingstorOK) String() string {
	return fmt.Sprintf("[POST /storage/qingstor][%d] postStorageQingstorOK  %+v", 200, o.Payload)
}

func (o *PostStorageQingstorOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageQingstorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageQingstorBadRequest creates a PostStorageQingstorBadRequest with default headers values
func NewPostStorageQingstorBadRequest() *PostStorageQingstorBadRequest {
	return &PostStorageQingstorBadRequest{}
}

/*
PostStorageQingstorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageQingstorBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage qingstor bad request response has a 2xx status code
func (o *PostStorageQingstorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage qingstor bad request response has a 3xx status code
func (o *PostStorageQingstorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage qingstor bad request response has a 4xx status code
func (o *PostStorageQingstorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage qingstor bad request response has a 5xx status code
func (o *PostStorageQingstorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage qingstor bad request response a status code equal to that given
func (o *PostStorageQingstorBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage qingstor bad request response
func (o *PostStorageQingstorBadRequest) Code() int {
	return 400
}

func (o *PostStorageQingstorBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/qingstor][%d] postStorageQingstorBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageQingstorBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/qingstor][%d] postStorageQingstorBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageQingstorBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageQingstorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageQingstorInternalServerError creates a PostStorageQingstorInternalServerError with default headers values
func NewPostStorageQingstorInternalServerError() *PostStorageQingstorInternalServerError {
	return &PostStorageQingstorInternalServerError{}
}

/*
PostStorageQingstorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageQingstorInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage qingstor internal server error response has a 2xx status code
func (o *PostStorageQingstorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage qingstor internal server error response has a 3xx status code
func (o *PostStorageQingstorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage qingstor internal server error response has a 4xx status code
func (o *PostStorageQingstorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage qingstor internal server error response has a 5xx status code
func (o *PostStorageQingstorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage qingstor internal server error response a status code equal to that given
func (o *PostStorageQingstorInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage qingstor internal server error response
func (o *PostStorageQingstorInternalServerError) Code() int {
	return 500
}

func (o *PostStorageQingstorInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/qingstor][%d] postStorageQingstorInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageQingstorInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/qingstor][%d] postStorageQingstorInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageQingstorInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageQingstorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
