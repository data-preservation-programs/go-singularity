// Code generated by go-swagger; DO NOT EDIT.

package deal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// SendManualReader is a Reader for the SendManual structure.
type SendManualReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendManualReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendManualOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendManualBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSendManualInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /send_deal] SendManual", response, response.Code())
	}
}

// NewSendManualOK creates a SendManualOK with default headers values
func NewSendManualOK() *SendManualOK {
	return &SendManualOK{}
}

/*
SendManualOK describes a response with status code 200, with default header values.

OK
*/
type SendManualOK struct {
	Payload *models.ModelDeal
}

// IsSuccess returns true when this send manual o k response has a 2xx status code
func (o *SendManualOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this send manual o k response has a 3xx status code
func (o *SendManualOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send manual o k response has a 4xx status code
func (o *SendManualOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this send manual o k response has a 5xx status code
func (o *SendManualOK) IsServerError() bool {
	return false
}

// IsCode returns true when this send manual o k response a status code equal to that given
func (o *SendManualOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the send manual o k response
func (o *SendManualOK) Code() int {
	return 200
}

func (o *SendManualOK) Error() string {
	return fmt.Sprintf("[POST /send_deal][%d] sendManualOK  %+v", 200, o.Payload)
}

func (o *SendManualOK) String() string {
	return fmt.Sprintf("[POST /send_deal][%d] sendManualOK  %+v", 200, o.Payload)
}

func (o *SendManualOK) GetPayload() *models.ModelDeal {
	return o.Payload
}

func (o *SendManualOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelDeal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendManualBadRequest creates a SendManualBadRequest with default headers values
func NewSendManualBadRequest() *SendManualBadRequest {
	return &SendManualBadRequest{}
}

/*
SendManualBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SendManualBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this send manual bad request response has a 2xx status code
func (o *SendManualBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send manual bad request response has a 3xx status code
func (o *SendManualBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send manual bad request response has a 4xx status code
func (o *SendManualBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this send manual bad request response has a 5xx status code
func (o *SendManualBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this send manual bad request response a status code equal to that given
func (o *SendManualBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the send manual bad request response
func (o *SendManualBadRequest) Code() int {
	return 400
}

func (o *SendManualBadRequest) Error() string {
	return fmt.Sprintf("[POST /send_deal][%d] sendManualBadRequest  %+v", 400, o.Payload)
}

func (o *SendManualBadRequest) String() string {
	return fmt.Sprintf("[POST /send_deal][%d] sendManualBadRequest  %+v", 400, o.Payload)
}

func (o *SendManualBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *SendManualBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendManualInternalServerError creates a SendManualInternalServerError with default headers values
func NewSendManualInternalServerError() *SendManualInternalServerError {
	return &SendManualInternalServerError{}
}

/*
SendManualInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SendManualInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this send manual internal server error response has a 2xx status code
func (o *SendManualInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send manual internal server error response has a 3xx status code
func (o *SendManualInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send manual internal server error response has a 4xx status code
func (o *SendManualInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this send manual internal server error response has a 5xx status code
func (o *SendManualInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this send manual internal server error response a status code equal to that given
func (o *SendManualInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the send manual internal server error response
func (o *SendManualInternalServerError) Code() int {
	return 500
}

func (o *SendManualInternalServerError) Error() string {
	return fmt.Sprintf("[POST /send_deal][%d] sendManualInternalServerError  %+v", 500, o.Payload)
}

func (o *SendManualInternalServerError) String() string {
	return fmt.Sprintf("[POST /send_deal][%d] sendManualInternalServerError  %+v", 500, o.Payload)
}

func (o *SendManualInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *SendManualInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
