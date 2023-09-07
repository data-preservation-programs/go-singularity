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

// CreatePreparationReader is a Reader for the CreatePreparation structure.
type CreatePreparationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePreparationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePreparationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePreparationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePreparationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation] CreatePreparation", response, response.Code())
	}
}

// NewCreatePreparationOK creates a CreatePreparationOK with default headers values
func NewCreatePreparationOK() *CreatePreparationOK {
	return &CreatePreparationOK{}
}

/*
CreatePreparationOK describes a response with status code 200, with default header values.

OK
*/
type CreatePreparationOK struct {
	Payload *models.ModelPreparation
}

// IsSuccess returns true when this create preparation o k response has a 2xx status code
func (o *CreatePreparationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create preparation o k response has a 3xx status code
func (o *CreatePreparationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preparation o k response has a 4xx status code
func (o *CreatePreparationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create preparation o k response has a 5xx status code
func (o *CreatePreparationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create preparation o k response a status code equal to that given
func (o *CreatePreparationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create preparation o k response
func (o *CreatePreparationOK) Code() int {
	return 200
}

func (o *CreatePreparationOK) Error() string {
	return fmt.Sprintf("[POST /preparation][%d] createPreparationOK  %+v", 200, o.Payload)
}

func (o *CreatePreparationOK) String() string {
	return fmt.Sprintf("[POST /preparation][%d] createPreparationOK  %+v", 200, o.Payload)
}

func (o *CreatePreparationOK) GetPayload() *models.ModelPreparation {
	return o.Payload
}

func (o *CreatePreparationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelPreparation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePreparationBadRequest creates a CreatePreparationBadRequest with default headers values
func NewCreatePreparationBadRequest() *CreatePreparationBadRequest {
	return &CreatePreparationBadRequest{}
}

/*
CreatePreparationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreatePreparationBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create preparation bad request response has a 2xx status code
func (o *CreatePreparationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create preparation bad request response has a 3xx status code
func (o *CreatePreparationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preparation bad request response has a 4xx status code
func (o *CreatePreparationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create preparation bad request response has a 5xx status code
func (o *CreatePreparationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create preparation bad request response a status code equal to that given
func (o *CreatePreparationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create preparation bad request response
func (o *CreatePreparationBadRequest) Code() int {
	return 400
}

func (o *CreatePreparationBadRequest) Error() string {
	return fmt.Sprintf("[POST /preparation][%d] createPreparationBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePreparationBadRequest) String() string {
	return fmt.Sprintf("[POST /preparation][%d] createPreparationBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePreparationBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreatePreparationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePreparationInternalServerError creates a CreatePreparationInternalServerError with default headers values
func NewCreatePreparationInternalServerError() *CreatePreparationInternalServerError {
	return &CreatePreparationInternalServerError{}
}

/*
CreatePreparationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreatePreparationInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create preparation internal server error response has a 2xx status code
func (o *CreatePreparationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create preparation internal server error response has a 3xx status code
func (o *CreatePreparationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preparation internal server error response has a 4xx status code
func (o *CreatePreparationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create preparation internal server error response has a 5xx status code
func (o *CreatePreparationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create preparation internal server error response a status code equal to that given
func (o *CreatePreparationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create preparation internal server error response
func (o *CreatePreparationInternalServerError) Code() int {
	return 500
}

func (o *CreatePreparationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /preparation][%d] createPreparationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePreparationInternalServerError) String() string {
	return fmt.Sprintf("[POST /preparation][%d] createPreparationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePreparationInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreatePreparationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
