// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostAdminResetReader is a Reader for the PostAdminReset structure.
type PostAdminResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAdminResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostAdminResetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostAdminResetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/reset] PostAdminReset", response, response.Code())
	}
}

// NewPostAdminResetNoContent creates a PostAdminResetNoContent with default headers values
func NewPostAdminResetNoContent() *PostAdminResetNoContent {
	return &PostAdminResetNoContent{}
}

/*
PostAdminResetNoContent describes a response with status code 204, with default header values.

No Content
*/
type PostAdminResetNoContent struct {
}

// IsSuccess returns true when this post admin reset no content response has a 2xx status code
func (o *PostAdminResetNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post admin reset no content response has a 3xx status code
func (o *PostAdminResetNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post admin reset no content response has a 4xx status code
func (o *PostAdminResetNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post admin reset no content response has a 5xx status code
func (o *PostAdminResetNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post admin reset no content response a status code equal to that given
func (o *PostAdminResetNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post admin reset no content response
func (o *PostAdminResetNoContent) Code() int {
	return 204
}

func (o *PostAdminResetNoContent) Error() string {
	return fmt.Sprintf("[POST /admin/reset][%d] postAdminResetNoContent ", 204)
}

func (o *PostAdminResetNoContent) String() string {
	return fmt.Sprintf("[POST /admin/reset][%d] postAdminResetNoContent ", 204)
}

func (o *PostAdminResetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAdminResetInternalServerError creates a PostAdminResetInternalServerError with default headers values
func NewPostAdminResetInternalServerError() *PostAdminResetInternalServerError {
	return &PostAdminResetInternalServerError{}
}

/*
PostAdminResetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAdminResetInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post admin reset internal server error response has a 2xx status code
func (o *PostAdminResetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post admin reset internal server error response has a 3xx status code
func (o *PostAdminResetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post admin reset internal server error response has a 4xx status code
func (o *PostAdminResetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post admin reset internal server error response has a 5xx status code
func (o *PostAdminResetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post admin reset internal server error response a status code equal to that given
func (o *PostAdminResetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post admin reset internal server error response
func (o *PostAdminResetInternalServerError) Code() int {
	return 500
}

func (o *PostAdminResetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/reset][%d] postAdminResetInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAdminResetInternalServerError) String() string {
	return fmt.Sprintf("[POST /admin/reset][%d] postAdminResetInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAdminResetInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostAdminResetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
