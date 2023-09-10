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

// CreateS3TencentCOSStorageReader is a Reader for the CreateS3TencentCOSStorage structure.
type CreateS3TencentCOSStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3TencentCOSStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3TencentCOSStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3TencentCOSStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3TencentCOSStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/tencentcos] CreateS3TencentCOSStorage", response, response.Code())
	}
}

// NewCreateS3TencentCOSStorageOK creates a CreateS3TencentCOSStorageOK with default headers values
func NewCreateS3TencentCOSStorageOK() *CreateS3TencentCOSStorageOK {
	return &CreateS3TencentCOSStorageOK{}
}

/*
CreateS3TencentCOSStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3TencentCOSStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 tencent c o s storage o k response has a 2xx status code
func (o *CreateS3TencentCOSStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 tencent c o s storage o k response has a 3xx status code
func (o *CreateS3TencentCOSStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 tencent c o s storage o k response has a 4xx status code
func (o *CreateS3TencentCOSStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 tencent c o s storage o k response has a 5xx status code
func (o *CreateS3TencentCOSStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 tencent c o s storage o k response a status code equal to that given
func (o *CreateS3TencentCOSStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 tencent c o s storage o k response
func (o *CreateS3TencentCOSStorageOK) Code() int {
	return 200
}

func (o *CreateS3TencentCOSStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/tencentcos][%d] createS3TencentCOSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3TencentCOSStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/tencentcos][%d] createS3TencentCOSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3TencentCOSStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3TencentCOSStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3TencentCOSStorageBadRequest creates a CreateS3TencentCOSStorageBadRequest with default headers values
func NewCreateS3TencentCOSStorageBadRequest() *CreateS3TencentCOSStorageBadRequest {
	return &CreateS3TencentCOSStorageBadRequest{}
}

/*
CreateS3TencentCOSStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3TencentCOSStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 tencent c o s storage bad request response has a 2xx status code
func (o *CreateS3TencentCOSStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 tencent c o s storage bad request response has a 3xx status code
func (o *CreateS3TencentCOSStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 tencent c o s storage bad request response has a 4xx status code
func (o *CreateS3TencentCOSStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 tencent c o s storage bad request response has a 5xx status code
func (o *CreateS3TencentCOSStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 tencent c o s storage bad request response a status code equal to that given
func (o *CreateS3TencentCOSStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 tencent c o s storage bad request response
func (o *CreateS3TencentCOSStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3TencentCOSStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/tencentcos][%d] createS3TencentCOSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3TencentCOSStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/tencentcos][%d] createS3TencentCOSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3TencentCOSStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3TencentCOSStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3TencentCOSStorageInternalServerError creates a CreateS3TencentCOSStorageInternalServerError with default headers values
func NewCreateS3TencentCOSStorageInternalServerError() *CreateS3TencentCOSStorageInternalServerError {
	return &CreateS3TencentCOSStorageInternalServerError{}
}

/*
CreateS3TencentCOSStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3TencentCOSStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 tencent c o s storage internal server error response has a 2xx status code
func (o *CreateS3TencentCOSStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 tencent c o s storage internal server error response has a 3xx status code
func (o *CreateS3TencentCOSStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 tencent c o s storage internal server error response has a 4xx status code
func (o *CreateS3TencentCOSStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 tencent c o s storage internal server error response has a 5xx status code
func (o *CreateS3TencentCOSStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 tencent c o s storage internal server error response a status code equal to that given
func (o *CreateS3TencentCOSStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 tencent c o s storage internal server error response
func (o *CreateS3TencentCOSStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3TencentCOSStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/tencentcos][%d] createS3TencentCOSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3TencentCOSStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/tencentcos][%d] createS3TencentCOSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3TencentCOSStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3TencentCOSStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}