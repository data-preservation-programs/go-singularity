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

// CreateHdfsStorageReader is a Reader for the CreateHdfsStorage structure.
type CreateHdfsStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHdfsStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateHdfsStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHdfsStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateHdfsStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/hdfs] CreateHdfsStorage", response, response.Code())
	}
}

// NewCreateHdfsStorageOK creates a CreateHdfsStorageOK with default headers values
func NewCreateHdfsStorageOK() *CreateHdfsStorageOK {
	return &CreateHdfsStorageOK{}
}

/*
CreateHdfsStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateHdfsStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create hdfs storage o k response has a 2xx status code
func (o *CreateHdfsStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create hdfs storage o k response has a 3xx status code
func (o *CreateHdfsStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create hdfs storage o k response has a 4xx status code
func (o *CreateHdfsStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create hdfs storage o k response has a 5xx status code
func (o *CreateHdfsStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create hdfs storage o k response a status code equal to that given
func (o *CreateHdfsStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create hdfs storage o k response
func (o *CreateHdfsStorageOK) Code() int {
	return 200
}

func (o *CreateHdfsStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/hdfs][%d] createHdfsStorageOK %s", 200, payload)
}

func (o *CreateHdfsStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/hdfs][%d] createHdfsStorageOK %s", 200, payload)
}

func (o *CreateHdfsStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateHdfsStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHdfsStorageBadRequest creates a CreateHdfsStorageBadRequest with default headers values
func NewCreateHdfsStorageBadRequest() *CreateHdfsStorageBadRequest {
	return &CreateHdfsStorageBadRequest{}
}

/*
CreateHdfsStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateHdfsStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create hdfs storage bad request response has a 2xx status code
func (o *CreateHdfsStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create hdfs storage bad request response has a 3xx status code
func (o *CreateHdfsStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create hdfs storage bad request response has a 4xx status code
func (o *CreateHdfsStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create hdfs storage bad request response has a 5xx status code
func (o *CreateHdfsStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create hdfs storage bad request response a status code equal to that given
func (o *CreateHdfsStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create hdfs storage bad request response
func (o *CreateHdfsStorageBadRequest) Code() int {
	return 400
}

func (o *CreateHdfsStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/hdfs][%d] createHdfsStorageBadRequest %s", 400, payload)
}

func (o *CreateHdfsStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/hdfs][%d] createHdfsStorageBadRequest %s", 400, payload)
}

func (o *CreateHdfsStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateHdfsStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHdfsStorageInternalServerError creates a CreateHdfsStorageInternalServerError with default headers values
func NewCreateHdfsStorageInternalServerError() *CreateHdfsStorageInternalServerError {
	return &CreateHdfsStorageInternalServerError{}
}

/*
CreateHdfsStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateHdfsStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create hdfs storage internal server error response has a 2xx status code
func (o *CreateHdfsStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create hdfs storage internal server error response has a 3xx status code
func (o *CreateHdfsStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create hdfs storage internal server error response has a 4xx status code
func (o *CreateHdfsStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create hdfs storage internal server error response has a 5xx status code
func (o *CreateHdfsStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create hdfs storage internal server error response a status code equal to that given
func (o *CreateHdfsStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create hdfs storage internal server error response
func (o *CreateHdfsStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateHdfsStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/hdfs][%d] createHdfsStorageInternalServerError %s", 500, payload)
}

func (o *CreateHdfsStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/hdfs][%d] createHdfsStorageInternalServerError %s", 500, payload)
}

func (o *CreateHdfsStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateHdfsStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
