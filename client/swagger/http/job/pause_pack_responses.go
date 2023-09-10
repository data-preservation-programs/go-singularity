// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PausePackReader is a Reader for the PausePack structure.
type PausePackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PausePackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPausePackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPausePackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPausePackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation/{id}/source/{name}/pause-pack/{job_id}] PausePack", response, response.Code())
	}
}

// NewPausePackOK creates a PausePackOK with default headers values
func NewPausePackOK() *PausePackOK {
	return &PausePackOK{}
}

/*
PausePackOK describes a response with status code 200, with default header values.

OK
*/
type PausePackOK struct {
	Payload []*models.ModelJob
}

// IsSuccess returns true when this pause pack o k response has a 2xx status code
func (o *PausePackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pause pack o k response has a 3xx status code
func (o *PausePackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause pack o k response has a 4xx status code
func (o *PausePackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause pack o k response has a 5xx status code
func (o *PausePackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pause pack o k response a status code equal to that given
func (o *PausePackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pause pack o k response
func (o *PausePackOK) Code() int {
	return 200
}

func (o *PausePackOK) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-pack/{job_id}][%d] pausePackOK  %+v", 200, o.Payload)
}

func (o *PausePackOK) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-pack/{job_id}][%d] pausePackOK  %+v", 200, o.Payload)
}

func (o *PausePackOK) GetPayload() []*models.ModelJob {
	return o.Payload
}

func (o *PausePackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPausePackBadRequest creates a PausePackBadRequest with default headers values
func NewPausePackBadRequest() *PausePackBadRequest {
	return &PausePackBadRequest{}
}

/*
PausePackBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PausePackBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this pause pack bad request response has a 2xx status code
func (o *PausePackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause pack bad request response has a 3xx status code
func (o *PausePackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause pack bad request response has a 4xx status code
func (o *PausePackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pause pack bad request response has a 5xx status code
func (o *PausePackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pause pack bad request response a status code equal to that given
func (o *PausePackBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pause pack bad request response
func (o *PausePackBadRequest) Code() int {
	return 400
}

func (o *PausePackBadRequest) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-pack/{job_id}][%d] pausePackBadRequest  %+v", 400, o.Payload)
}

func (o *PausePackBadRequest) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-pack/{job_id}][%d] pausePackBadRequest  %+v", 400, o.Payload)
}

func (o *PausePackBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PausePackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPausePackInternalServerError creates a PausePackInternalServerError with default headers values
func NewPausePackInternalServerError() *PausePackInternalServerError {
	return &PausePackInternalServerError{}
}

/*
PausePackInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PausePackInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this pause pack internal server error response has a 2xx status code
func (o *PausePackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause pack internal server error response has a 3xx status code
func (o *PausePackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause pack internal server error response has a 4xx status code
func (o *PausePackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause pack internal server error response has a 5xx status code
func (o *PausePackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pause pack internal server error response a status code equal to that given
func (o *PausePackInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pause pack internal server error response
func (o *PausePackInternalServerError) Code() int {
	return 500
}

func (o *PausePackInternalServerError) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-pack/{job_id}][%d] pausePackInternalServerError  %+v", 500, o.Payload)
}

func (o *PausePackInternalServerError) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/pause-pack/{job_id}][%d] pausePackInternalServerError  %+v", 500, o.Payload)
}

func (o *PausePackInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PausePackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}