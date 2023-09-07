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

// CreateS3IDriveStorageReader is a Reader for the CreateS3IDriveStorage structure.
type CreateS3IDriveStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3IDriveStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3IDriveStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3IDriveStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3IDriveStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/idrive] CreateS3IDriveStorage", response, response.Code())
	}
}

// NewCreateS3IDriveStorageOK creates a CreateS3IDriveStorageOK with default headers values
func NewCreateS3IDriveStorageOK() *CreateS3IDriveStorageOK {
	return &CreateS3IDriveStorageOK{}
}

/*
CreateS3IDriveStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3IDriveStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 i drive storage o k response has a 2xx status code
func (o *CreateS3IDriveStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 i drive storage o k response has a 3xx status code
func (o *CreateS3IDriveStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 i drive storage o k response has a 4xx status code
func (o *CreateS3IDriveStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 i drive storage o k response has a 5xx status code
func (o *CreateS3IDriveStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 i drive storage o k response a status code equal to that given
func (o *CreateS3IDriveStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 i drive storage o k response
func (o *CreateS3IDriveStorageOK) Code() int {
	return 200
}

func (o *CreateS3IDriveStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/idrive][%d] createS3IDriveStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3IDriveStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/idrive][%d] createS3IDriveStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3IDriveStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3IDriveStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3IDriveStorageBadRequest creates a CreateS3IDriveStorageBadRequest with default headers values
func NewCreateS3IDriveStorageBadRequest() *CreateS3IDriveStorageBadRequest {
	return &CreateS3IDriveStorageBadRequest{}
}

/*
CreateS3IDriveStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3IDriveStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 i drive storage bad request response has a 2xx status code
func (o *CreateS3IDriveStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 i drive storage bad request response has a 3xx status code
func (o *CreateS3IDriveStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 i drive storage bad request response has a 4xx status code
func (o *CreateS3IDriveStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 i drive storage bad request response has a 5xx status code
func (o *CreateS3IDriveStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 i drive storage bad request response a status code equal to that given
func (o *CreateS3IDriveStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 i drive storage bad request response
func (o *CreateS3IDriveStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3IDriveStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/idrive][%d] createS3IDriveStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3IDriveStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/idrive][%d] createS3IDriveStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3IDriveStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3IDriveStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3IDriveStorageInternalServerError creates a CreateS3IDriveStorageInternalServerError with default headers values
func NewCreateS3IDriveStorageInternalServerError() *CreateS3IDriveStorageInternalServerError {
	return &CreateS3IDriveStorageInternalServerError{}
}

/*
CreateS3IDriveStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3IDriveStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 i drive storage internal server error response has a 2xx status code
func (o *CreateS3IDriveStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 i drive storage internal server error response has a 3xx status code
func (o *CreateS3IDriveStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 i drive storage internal server error response has a 4xx status code
func (o *CreateS3IDriveStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 i drive storage internal server error response has a 5xx status code
func (o *CreateS3IDriveStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 i drive storage internal server error response a status code equal to that given
func (o *CreateS3IDriveStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 i drive storage internal server error response
func (o *CreateS3IDriveStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3IDriveStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/idrive][%d] createS3IDriveStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3IDriveStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/idrive][%d] createS3IDriveStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3IDriveStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3IDriveStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
