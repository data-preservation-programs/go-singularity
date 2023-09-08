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

// PostStorageKoofrKoofrReader is a Reader for the PostStorageKoofrKoofr structure.
type PostStorageKoofrKoofrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageKoofrKoofrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageKoofrKoofrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageKoofrKoofrBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageKoofrKoofrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/koofr/koofr] PostStorageKoofrKoofr", response, response.Code())
	}
}

// NewPostStorageKoofrKoofrOK creates a PostStorageKoofrKoofrOK with default headers values
func NewPostStorageKoofrKoofrOK() *PostStorageKoofrKoofrOK {
	return &PostStorageKoofrKoofrOK{}
}

/*
PostStorageKoofrKoofrOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageKoofrKoofrOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage koofr koofr o k response has a 2xx status code
func (o *PostStorageKoofrKoofrOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage koofr koofr o k response has a 3xx status code
func (o *PostStorageKoofrKoofrOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage koofr koofr o k response has a 4xx status code
func (o *PostStorageKoofrKoofrOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage koofr koofr o k response has a 5xx status code
func (o *PostStorageKoofrKoofrOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage koofr koofr o k response a status code equal to that given
func (o *PostStorageKoofrKoofrOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage koofr koofr o k response
func (o *PostStorageKoofrKoofrOK) Code() int {
	return 200
}

func (o *PostStorageKoofrKoofrOK) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] postStorageKoofrKoofrOK  %+v", 200, o.Payload)
}

func (o *PostStorageKoofrKoofrOK) String() string {
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] postStorageKoofrKoofrOK  %+v", 200, o.Payload)
}

func (o *PostStorageKoofrKoofrOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageKoofrKoofrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageKoofrKoofrBadRequest creates a PostStorageKoofrKoofrBadRequest with default headers values
func NewPostStorageKoofrKoofrBadRequest() *PostStorageKoofrKoofrBadRequest {
	return &PostStorageKoofrKoofrBadRequest{}
}

/*
PostStorageKoofrKoofrBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageKoofrKoofrBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage koofr koofr bad request response has a 2xx status code
func (o *PostStorageKoofrKoofrBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage koofr koofr bad request response has a 3xx status code
func (o *PostStorageKoofrKoofrBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage koofr koofr bad request response has a 4xx status code
func (o *PostStorageKoofrKoofrBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage koofr koofr bad request response has a 5xx status code
func (o *PostStorageKoofrKoofrBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage koofr koofr bad request response a status code equal to that given
func (o *PostStorageKoofrKoofrBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage koofr koofr bad request response
func (o *PostStorageKoofrKoofrBadRequest) Code() int {
	return 400
}

func (o *PostStorageKoofrKoofrBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] postStorageKoofrKoofrBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageKoofrKoofrBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] postStorageKoofrKoofrBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageKoofrKoofrBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageKoofrKoofrBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageKoofrKoofrInternalServerError creates a PostStorageKoofrKoofrInternalServerError with default headers values
func NewPostStorageKoofrKoofrInternalServerError() *PostStorageKoofrKoofrInternalServerError {
	return &PostStorageKoofrKoofrInternalServerError{}
}

/*
PostStorageKoofrKoofrInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageKoofrKoofrInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage koofr koofr internal server error response has a 2xx status code
func (o *PostStorageKoofrKoofrInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage koofr koofr internal server error response has a 3xx status code
func (o *PostStorageKoofrKoofrInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage koofr koofr internal server error response has a 4xx status code
func (o *PostStorageKoofrKoofrInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage koofr koofr internal server error response has a 5xx status code
func (o *PostStorageKoofrKoofrInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage koofr koofr internal server error response a status code equal to that given
func (o *PostStorageKoofrKoofrInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage koofr koofr internal server error response
func (o *PostStorageKoofrKoofrInternalServerError) Code() int {
	return 500
}

func (o *PostStorageKoofrKoofrInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] postStorageKoofrKoofrInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageKoofrKoofrInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] postStorageKoofrKoofrInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageKoofrKoofrInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageKoofrKoofrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
