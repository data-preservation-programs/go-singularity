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

// CreateS3HuaweiOBSStorageReader is a Reader for the CreateS3HuaweiOBSStorage structure.
type CreateS3HuaweiOBSStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3HuaweiOBSStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3HuaweiOBSStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3HuaweiOBSStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3HuaweiOBSStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/huaweiobs] CreateS3HuaweiOBSStorage", response, response.Code())
	}
}

// NewCreateS3HuaweiOBSStorageOK creates a CreateS3HuaweiOBSStorageOK with default headers values
func NewCreateS3HuaweiOBSStorageOK() *CreateS3HuaweiOBSStorageOK {
	return &CreateS3HuaweiOBSStorageOK{}
}

/*
CreateS3HuaweiOBSStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3HuaweiOBSStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 huawei o b s storage o k response has a 2xx status code
func (o *CreateS3HuaweiOBSStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 huawei o b s storage o k response has a 3xx status code
func (o *CreateS3HuaweiOBSStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 huawei o b s storage o k response has a 4xx status code
func (o *CreateS3HuaweiOBSStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 huawei o b s storage o k response has a 5xx status code
func (o *CreateS3HuaweiOBSStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 huawei o b s storage o k response a status code equal to that given
func (o *CreateS3HuaweiOBSStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 huawei o b s storage o k response
func (o *CreateS3HuaweiOBSStorageOK) Code() int {
	return 200
}

func (o *CreateS3HuaweiOBSStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/huaweiobs][%d] createS3HuaweiOBSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3HuaweiOBSStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/huaweiobs][%d] createS3HuaweiOBSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3HuaweiOBSStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3HuaweiOBSStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3HuaweiOBSStorageBadRequest creates a CreateS3HuaweiOBSStorageBadRequest with default headers values
func NewCreateS3HuaweiOBSStorageBadRequest() *CreateS3HuaweiOBSStorageBadRequest {
	return &CreateS3HuaweiOBSStorageBadRequest{}
}

/*
CreateS3HuaweiOBSStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3HuaweiOBSStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 huawei o b s storage bad request response has a 2xx status code
func (o *CreateS3HuaweiOBSStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 huawei o b s storage bad request response has a 3xx status code
func (o *CreateS3HuaweiOBSStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 huawei o b s storage bad request response has a 4xx status code
func (o *CreateS3HuaweiOBSStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 huawei o b s storage bad request response has a 5xx status code
func (o *CreateS3HuaweiOBSStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 huawei o b s storage bad request response a status code equal to that given
func (o *CreateS3HuaweiOBSStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 huawei o b s storage bad request response
func (o *CreateS3HuaweiOBSStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3HuaweiOBSStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/huaweiobs][%d] createS3HuaweiOBSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3HuaweiOBSStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/huaweiobs][%d] createS3HuaweiOBSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3HuaweiOBSStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3HuaweiOBSStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3HuaweiOBSStorageInternalServerError creates a CreateS3HuaweiOBSStorageInternalServerError with default headers values
func NewCreateS3HuaweiOBSStorageInternalServerError() *CreateS3HuaweiOBSStorageInternalServerError {
	return &CreateS3HuaweiOBSStorageInternalServerError{}
}

/*
CreateS3HuaweiOBSStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3HuaweiOBSStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 huawei o b s storage internal server error response has a 2xx status code
func (o *CreateS3HuaweiOBSStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 huawei o b s storage internal server error response has a 3xx status code
func (o *CreateS3HuaweiOBSStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 huawei o b s storage internal server error response has a 4xx status code
func (o *CreateS3HuaweiOBSStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 huawei o b s storage internal server error response has a 5xx status code
func (o *CreateS3HuaweiOBSStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 huawei o b s storage internal server error response a status code equal to that given
func (o *CreateS3HuaweiOBSStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 huawei o b s storage internal server error response
func (o *CreateS3HuaweiOBSStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3HuaweiOBSStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/huaweiobs][%d] createS3HuaweiOBSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3HuaweiOBSStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/huaweiobs][%d] createS3HuaweiOBSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3HuaweiOBSStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3HuaweiOBSStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
