// Code generated by go-swagger; DO NOT EDIT.

package preparation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostPreparationIDSourceNameReader is a Reader for the PostPreparationIDSourceName structure.
type PostPreparationIDSourceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPreparationIDSourceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPreparationIDSourceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPreparationIDSourceNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPreparationIDSourceNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation/{id}/source/{name}] PostPreparationIDSourceName", response, response.Code())
	}
}

// NewPostPreparationIDSourceNameOK creates a PostPreparationIDSourceNameOK with default headers values
func NewPostPreparationIDSourceNameOK() *PostPreparationIDSourceNameOK {
	return &PostPreparationIDSourceNameOK{}
}

/*
PostPreparationIDSourceNameOK describes a response with status code 200, with default header values.

OK
*/
type PostPreparationIDSourceNameOK struct {
	Payload *models.ModelPreparation
}

// IsSuccess returns true when this post preparation Id source name o k response has a 2xx status code
func (o *PostPreparationIDSourceNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post preparation Id source name o k response has a 3xx status code
func (o *PostPreparationIDSourceNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id source name o k response has a 4xx status code
func (o *PostPreparationIDSourceNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post preparation Id source name o k response has a 5xx status code
func (o *PostPreparationIDSourceNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post preparation Id source name o k response a status code equal to that given
func (o *PostPreparationIDSourceNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post preparation Id source name o k response
func (o *PostPreparationIDSourceNameOK) Code() int {
	return 200
}

func (o *PostPreparationIDSourceNameOK) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}][%d] postPreparationIdSourceNameOK  %+v", 200, o.Payload)
}

func (o *PostPreparationIDSourceNameOK) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}][%d] postPreparationIdSourceNameOK  %+v", 200, o.Payload)
}

func (o *PostPreparationIDSourceNameOK) GetPayload() *models.ModelPreparation {
	return o.Payload
}

func (o *PostPreparationIDSourceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelPreparation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPreparationIDSourceNameBadRequest creates a PostPreparationIDSourceNameBadRequest with default headers values
func NewPostPreparationIDSourceNameBadRequest() *PostPreparationIDSourceNameBadRequest {
	return &PostPreparationIDSourceNameBadRequest{}
}

/*
PostPreparationIDSourceNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostPreparationIDSourceNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post preparation Id source name bad request response has a 2xx status code
func (o *PostPreparationIDSourceNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post preparation Id source name bad request response has a 3xx status code
func (o *PostPreparationIDSourceNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id source name bad request response has a 4xx status code
func (o *PostPreparationIDSourceNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post preparation Id source name bad request response has a 5xx status code
func (o *PostPreparationIDSourceNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post preparation Id source name bad request response a status code equal to that given
func (o *PostPreparationIDSourceNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post preparation Id source name bad request response
func (o *PostPreparationIDSourceNameBadRequest) Code() int {
	return 400
}

func (o *PostPreparationIDSourceNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}][%d] postPreparationIdSourceNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostPreparationIDSourceNameBadRequest) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}][%d] postPreparationIdSourceNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostPreparationIDSourceNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostPreparationIDSourceNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPreparationIDSourceNameInternalServerError creates a PostPreparationIDSourceNameInternalServerError with default headers values
func NewPostPreparationIDSourceNameInternalServerError() *PostPreparationIDSourceNameInternalServerError {
	return &PostPreparationIDSourceNameInternalServerError{}
}

/*
PostPreparationIDSourceNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostPreparationIDSourceNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post preparation Id source name internal server error response has a 2xx status code
func (o *PostPreparationIDSourceNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post preparation Id source name internal server error response has a 3xx status code
func (o *PostPreparationIDSourceNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id source name internal server error response has a 4xx status code
func (o *PostPreparationIDSourceNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post preparation Id source name internal server error response has a 5xx status code
func (o *PostPreparationIDSourceNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post preparation Id source name internal server error response a status code equal to that given
func (o *PostPreparationIDSourceNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post preparation Id source name internal server error response
func (o *PostPreparationIDSourceNameInternalServerError) Code() int {
	return 500
}

func (o *PostPreparationIDSourceNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}][%d] postPreparationIdSourceNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPreparationIDSourceNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}][%d] postPreparationIdSourceNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPreparationIDSourceNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostPreparationIDSourceNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
