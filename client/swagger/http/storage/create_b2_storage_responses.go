// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// CreateB2StorageReader is a Reader for the CreateB2Storage structure.
type CreateB2StorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateB2StorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateB2StorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateB2StorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateB2StorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/b2] CreateB2Storage", response, response.Code())
	}
}

// NewCreateB2StorageOK creates a CreateB2StorageOK with default headers values
func NewCreateB2StorageOK() *CreateB2StorageOK {
	return &CreateB2StorageOK{}
}

/*
CreateB2StorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateB2StorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create b2 storage o k response has a 2xx status code
func (o *CreateB2StorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create b2 storage o k response has a 3xx status code
func (o *CreateB2StorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create b2 storage o k response has a 4xx status code
func (o *CreateB2StorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create b2 storage o k response has a 5xx status code
func (o *CreateB2StorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create b2 storage o k response a status code equal to that given
func (o *CreateB2StorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create b2 storage o k response
func (o *CreateB2StorageOK) Code() int {
	return 200
}

func (o *CreateB2StorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/b2][%d] createB2StorageOK %s", 200, payload)
}

func (o *CreateB2StorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/b2][%d] createB2StorageOK %s", 200, payload)
}

func (o *CreateB2StorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateB2StorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateB2StorageBadRequest creates a CreateB2StorageBadRequest with default headers values
func NewCreateB2StorageBadRequest() *CreateB2StorageBadRequest {
	return &CreateB2StorageBadRequest{}
}

/*
CreateB2StorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateB2StorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create b2 storage bad request response has a 2xx status code
func (o *CreateB2StorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create b2 storage bad request response has a 3xx status code
func (o *CreateB2StorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create b2 storage bad request response has a 4xx status code
func (o *CreateB2StorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create b2 storage bad request response has a 5xx status code
func (o *CreateB2StorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create b2 storage bad request response a status code equal to that given
func (o *CreateB2StorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create b2 storage bad request response
func (o *CreateB2StorageBadRequest) Code() int {
	return 400
}

func (o *CreateB2StorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/b2][%d] createB2StorageBadRequest %s", 400, payload)
}

func (o *CreateB2StorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/b2][%d] createB2StorageBadRequest %s", 400, payload)
}

func (o *CreateB2StorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateB2StorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateB2StorageInternalServerError creates a CreateB2StorageInternalServerError with default headers values
func NewCreateB2StorageInternalServerError() *CreateB2StorageInternalServerError {
	return &CreateB2StorageInternalServerError{}
}

/*
CreateB2StorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateB2StorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create b2 storage internal server error response has a 2xx status code
func (o *CreateB2StorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create b2 storage internal server error response has a 3xx status code
func (o *CreateB2StorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create b2 storage internal server error response has a 4xx status code
func (o *CreateB2StorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create b2 storage internal server error response has a 5xx status code
func (o *CreateB2StorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create b2 storage internal server error response a status code equal to that given
func (o *CreateB2StorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create b2 storage internal server error response
func (o *CreateB2StorageInternalServerError) Code() int {
	return 500
}

func (o *CreateB2StorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/b2][%d] createB2StorageInternalServerError %s", 500, payload)
}

func (o *CreateB2StorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/b2][%d] createB2StorageInternalServerError %s", 500, payload)
}

func (o *CreateB2StorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateB2StorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
