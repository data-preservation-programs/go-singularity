// Code generated by go-swagger; DO NOT EDIT.

package file

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

// GetFileDealsReader is a Reader for the GetFileDeals structure.
type GetFileDealsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileDealsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileDealsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetFileDealsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /file/{id}/deals] GetFileDeals", response, response.Code())
	}
}

// NewGetFileDealsOK creates a GetFileDealsOK with default headers values
func NewGetFileDealsOK() *GetFileDealsOK {
	return &GetFileDealsOK{}
}

/*
GetFileDealsOK describes a response with status code 200, with default header values.

OK
*/
type GetFileDealsOK struct {
	Payload []*models.FileDealsForFileRange
}

// IsSuccess returns true when this get file deals o k response has a 2xx status code
func (o *GetFileDealsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file deals o k response has a 3xx status code
func (o *GetFileDealsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file deals o k response has a 4xx status code
func (o *GetFileDealsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file deals o k response has a 5xx status code
func (o *GetFileDealsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file deals o k response a status code equal to that given
func (o *GetFileDealsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file deals o k response
func (o *GetFileDealsOK) Code() int {
	return 200
}

func (o *GetFileDealsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/deals][%d] getFileDealsOK %s", 200, payload)
}

func (o *GetFileDealsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/deals][%d] getFileDealsOK %s", 200, payload)
}

func (o *GetFileDealsOK) GetPayload() []*models.FileDealsForFileRange {
	return o.Payload
}

func (o *GetFileDealsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileDealsInternalServerError creates a GetFileDealsInternalServerError with default headers values
func NewGetFileDealsInternalServerError() *GetFileDealsInternalServerError {
	return &GetFileDealsInternalServerError{}
}

/*
GetFileDealsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFileDealsInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get file deals internal server error response has a 2xx status code
func (o *GetFileDealsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file deals internal server error response has a 3xx status code
func (o *GetFileDealsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file deals internal server error response has a 4xx status code
func (o *GetFileDealsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file deals internal server error response has a 5xx status code
func (o *GetFileDealsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get file deals internal server error response a status code equal to that given
func (o *GetFileDealsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get file deals internal server error response
func (o *GetFileDealsInternalServerError) Code() int {
	return 500
}

func (o *GetFileDealsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/deals][%d] getFileDealsInternalServerError %s", 500, payload)
}

func (o *GetFileDealsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/deals][%d] getFileDealsInternalServerError %s", 500, payload)
}

func (o *GetFileDealsInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetFileDealsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
