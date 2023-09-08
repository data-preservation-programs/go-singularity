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

// GetStorageReader is a Reader for the GetStorage structure.
type GetStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /storage] GetStorage", response, response.Code())
	}
}

// NewGetStorageOK creates a GetStorageOK with default headers values
func NewGetStorageOK() *GetStorageOK {
	return &GetStorageOK{}
}

/*
GetStorageOK describes a response with status code 200, with default header values.

OK
*/
type GetStorageOK struct {
	Payload []*models.ModelStorage
}

// IsSuccess returns true when this get storage o k response has a 2xx status code
func (o *GetStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get storage o k response has a 3xx status code
func (o *GetStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage o k response has a 4xx status code
func (o *GetStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get storage o k response has a 5xx status code
func (o *GetStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get storage o k response a status code equal to that given
func (o *GetStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get storage o k response
func (o *GetStorageOK) Code() int {
	return 200
}

func (o *GetStorageOK) Error() string {
	return fmt.Sprintf("[GET /storage][%d] getStorageOK  %+v", 200, o.Payload)
}

func (o *GetStorageOK) String() string {
	return fmt.Sprintf("[GET /storage][%d] getStorageOK  %+v", 200, o.Payload)
}

func (o *GetStorageOK) GetPayload() []*models.ModelStorage {
	return o.Payload
}

func (o *GetStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageBadRequest creates a GetStorageBadRequest with default headers values
func NewGetStorageBadRequest() *GetStorageBadRequest {
	return &GetStorageBadRequest{}
}

/*
GetStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get storage bad request response has a 2xx status code
func (o *GetStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get storage bad request response has a 3xx status code
func (o *GetStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage bad request response has a 4xx status code
func (o *GetStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get storage bad request response has a 5xx status code
func (o *GetStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get storage bad request response a status code equal to that given
func (o *GetStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get storage bad request response
func (o *GetStorageBadRequest) Code() int {
	return 400
}

func (o *GetStorageBadRequest) Error() string {
	return fmt.Sprintf("[GET /storage][%d] getStorageBadRequest  %+v", 400, o.Payload)
}

func (o *GetStorageBadRequest) String() string {
	return fmt.Sprintf("[GET /storage][%d] getStorageBadRequest  %+v", 400, o.Payload)
}

func (o *GetStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageInternalServerError creates a GetStorageInternalServerError with default headers values
func NewGetStorageInternalServerError() *GetStorageInternalServerError {
	return &GetStorageInternalServerError{}
}

/*
GetStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get storage internal server error response has a 2xx status code
func (o *GetStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get storage internal server error response has a 3xx status code
func (o *GetStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage internal server error response has a 4xx status code
func (o *GetStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get storage internal server error response has a 5xx status code
func (o *GetStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get storage internal server error response a status code equal to that given
func (o *GetStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get storage internal server error response
func (o *GetStorageInternalServerError) Code() int {
	return 500
}

func (o *GetStorageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /storage][%d] getStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStorageInternalServerError) String() string {
	return fmt.Sprintf("[GET /storage][%d] getStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
