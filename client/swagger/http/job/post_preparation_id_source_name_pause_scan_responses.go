// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostPreparationIDSourceNamePauseScanReader is a Reader for the PostPreparationIDSourceNamePauseScan structure.
type PostPreparationIDSourceNamePauseScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPreparationIDSourceNamePauseScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPreparationIDSourceNamePauseScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPreparationIDSourceNamePauseScanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPreparationIDSourceNamePauseScanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation/{id}/source/{name}/pause-scan] PostPreparationIDSourceNamePauseScan", response, response.Code())
	}
}

// NewPostPreparationIDSourceNamePauseScanOK creates a PostPreparationIDSourceNamePauseScanOK with default headers values
func NewPostPreparationIDSourceNamePauseScanOK() *PostPreparationIDSourceNamePauseScanOK {
	return &PostPreparationIDSourceNamePauseScanOK{}
}

/*
PostPreparationIDSourceNamePauseScanOK describes a response with status code 200, with default header values.

OK
*/
type PostPreparationIDSourceNamePauseScanOK struct {
	Payload *models.ModelJob
}

// IsSuccess returns true when this post preparation Id source name pause scan o k response has a 2xx status code
func (o *PostPreparationIDSourceNamePauseScanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post preparation Id source name pause scan o k response has a 3xx status code
func (o *PostPreparationIDSourceNamePauseScanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id source name pause scan o k response has a 4xx status code
func (o *PostPreparationIDSourceNamePauseScanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post preparation Id source name pause scan o k response has a 5xx status code
func (o *PostPreparationIDSourceNamePauseScanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post preparation Id source name pause scan o k response a status code equal to that given
func (o *PostPreparationIDSourceNamePauseScanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post preparation Id source name pause scan o k response
func (o *PostPreparationIDSourceNamePauseScanOK) Code() int {
	return 200
}

func (o *PostPreparationIDSourceNamePauseScanOK) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-scan][%d] postPreparationIdSourceNamePauseScanOK  %+v", 200, o.Payload)
}

func (o *PostPreparationIDSourceNamePauseScanOK) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-scan][%d] postPreparationIdSourceNamePauseScanOK  %+v", 200, o.Payload)
}

func (o *PostPreparationIDSourceNamePauseScanOK) GetPayload() *models.ModelJob {
	return o.Payload
}

func (o *PostPreparationIDSourceNamePauseScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPreparationIDSourceNamePauseScanBadRequest creates a PostPreparationIDSourceNamePauseScanBadRequest with default headers values
func NewPostPreparationIDSourceNamePauseScanBadRequest() *PostPreparationIDSourceNamePauseScanBadRequest {
	return &PostPreparationIDSourceNamePauseScanBadRequest{}
}

/*
PostPreparationIDSourceNamePauseScanBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostPreparationIDSourceNamePauseScanBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post preparation Id source name pause scan bad request response has a 2xx status code
func (o *PostPreparationIDSourceNamePauseScanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post preparation Id source name pause scan bad request response has a 3xx status code
func (o *PostPreparationIDSourceNamePauseScanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id source name pause scan bad request response has a 4xx status code
func (o *PostPreparationIDSourceNamePauseScanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post preparation Id source name pause scan bad request response has a 5xx status code
func (o *PostPreparationIDSourceNamePauseScanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post preparation Id source name pause scan bad request response a status code equal to that given
func (o *PostPreparationIDSourceNamePauseScanBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post preparation Id source name pause scan bad request response
func (o *PostPreparationIDSourceNamePauseScanBadRequest) Code() int {
	return 400
}

func (o *PostPreparationIDSourceNamePauseScanBadRequest) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-scan][%d] postPreparationIdSourceNamePauseScanBadRequest  %+v", 400, o.Payload)
}

func (o *PostPreparationIDSourceNamePauseScanBadRequest) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-scan][%d] postPreparationIdSourceNamePauseScanBadRequest  %+v", 400, o.Payload)
}

func (o *PostPreparationIDSourceNamePauseScanBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostPreparationIDSourceNamePauseScanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPreparationIDSourceNamePauseScanInternalServerError creates a PostPreparationIDSourceNamePauseScanInternalServerError with default headers values
func NewPostPreparationIDSourceNamePauseScanInternalServerError() *PostPreparationIDSourceNamePauseScanInternalServerError {
	return &PostPreparationIDSourceNamePauseScanInternalServerError{}
}

/*
PostPreparationIDSourceNamePauseScanInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostPreparationIDSourceNamePauseScanInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post preparation Id source name pause scan internal server error response has a 2xx status code
func (o *PostPreparationIDSourceNamePauseScanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post preparation Id source name pause scan internal server error response has a 3xx status code
func (o *PostPreparationIDSourceNamePauseScanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id source name pause scan internal server error response has a 4xx status code
func (o *PostPreparationIDSourceNamePauseScanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post preparation Id source name pause scan internal server error response has a 5xx status code
func (o *PostPreparationIDSourceNamePauseScanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post preparation Id source name pause scan internal server error response a status code equal to that given
func (o *PostPreparationIDSourceNamePauseScanInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post preparation Id source name pause scan internal server error response
func (o *PostPreparationIDSourceNamePauseScanInternalServerError) Code() int {
	return 500
}

func (o *PostPreparationIDSourceNamePauseScanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-scan][%d] postPreparationIdSourceNamePauseScanInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPreparationIDSourceNamePauseScanInternalServerError) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-scan][%d] postPreparationIdSourceNamePauseScanInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPreparationIDSourceNamePauseScanInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostPreparationIDSourceNamePauseScanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
