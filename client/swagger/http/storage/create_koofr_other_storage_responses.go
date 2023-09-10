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

// CreateKoofrOtherStorageReader is a Reader for the CreateKoofrOtherStorage structure.
type CreateKoofrOtherStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateKoofrOtherStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateKoofrOtherStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateKoofrOtherStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateKoofrOtherStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/koofr/other] CreateKoofrOtherStorage", response, response.Code())
	}
}

// NewCreateKoofrOtherStorageOK creates a CreateKoofrOtherStorageOK with default headers values
func NewCreateKoofrOtherStorageOK() *CreateKoofrOtherStorageOK {
	return &CreateKoofrOtherStorageOK{}
}

/*
CreateKoofrOtherStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateKoofrOtherStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create koofr other storage o k response has a 2xx status code
func (o *CreateKoofrOtherStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create koofr other storage o k response has a 3xx status code
func (o *CreateKoofrOtherStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create koofr other storage o k response has a 4xx status code
func (o *CreateKoofrOtherStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create koofr other storage o k response has a 5xx status code
func (o *CreateKoofrOtherStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create koofr other storage o k response a status code equal to that given
func (o *CreateKoofrOtherStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create koofr other storage o k response
func (o *CreateKoofrOtherStorageOK) Code() int {
	return 200
}

func (o *CreateKoofrOtherStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] createKoofrOtherStorageOK  %+v", 200, o.Payload)
}

func (o *CreateKoofrOtherStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] createKoofrOtherStorageOK  %+v", 200, o.Payload)
}

func (o *CreateKoofrOtherStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateKoofrOtherStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKoofrOtherStorageBadRequest creates a CreateKoofrOtherStorageBadRequest with default headers values
func NewCreateKoofrOtherStorageBadRequest() *CreateKoofrOtherStorageBadRequest {
	return &CreateKoofrOtherStorageBadRequest{}
}

/*
CreateKoofrOtherStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateKoofrOtherStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create koofr other storage bad request response has a 2xx status code
func (o *CreateKoofrOtherStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create koofr other storage bad request response has a 3xx status code
func (o *CreateKoofrOtherStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create koofr other storage bad request response has a 4xx status code
func (o *CreateKoofrOtherStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create koofr other storage bad request response has a 5xx status code
func (o *CreateKoofrOtherStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create koofr other storage bad request response a status code equal to that given
func (o *CreateKoofrOtherStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create koofr other storage bad request response
func (o *CreateKoofrOtherStorageBadRequest) Code() int {
	return 400
}

func (o *CreateKoofrOtherStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] createKoofrOtherStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateKoofrOtherStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] createKoofrOtherStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateKoofrOtherStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateKoofrOtherStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKoofrOtherStorageInternalServerError creates a CreateKoofrOtherStorageInternalServerError with default headers values
func NewCreateKoofrOtherStorageInternalServerError() *CreateKoofrOtherStorageInternalServerError {
	return &CreateKoofrOtherStorageInternalServerError{}
}

/*
CreateKoofrOtherStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateKoofrOtherStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create koofr other storage internal server error response has a 2xx status code
func (o *CreateKoofrOtherStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create koofr other storage internal server error response has a 3xx status code
func (o *CreateKoofrOtherStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create koofr other storage internal server error response has a 4xx status code
func (o *CreateKoofrOtherStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create koofr other storage internal server error response has a 5xx status code
func (o *CreateKoofrOtherStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create koofr other storage internal server error response a status code equal to that given
func (o *CreateKoofrOtherStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create koofr other storage internal server error response
func (o *CreateKoofrOtherStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateKoofrOtherStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] createKoofrOtherStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateKoofrOtherStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] createKoofrOtherStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateKoofrOtherStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateKoofrOtherStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
