// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// DeleteWalletAddressReader is a Reader for the DeleteWalletAddress structure.
type DeleteWalletAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWalletAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWalletAddressNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWalletAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWalletAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /wallet/{address}] DeleteWalletAddress", response, response.Code())
	}
}

// NewDeleteWalletAddressNoContent creates a DeleteWalletAddressNoContent with default headers values
func NewDeleteWalletAddressNoContent() *DeleteWalletAddressNoContent {
	return &DeleteWalletAddressNoContent{}
}

/*
DeleteWalletAddressNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteWalletAddressNoContent struct {
}

// IsSuccess returns true when this delete wallet address no content response has a 2xx status code
func (o *DeleteWalletAddressNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete wallet address no content response has a 3xx status code
func (o *DeleteWalletAddressNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete wallet address no content response has a 4xx status code
func (o *DeleteWalletAddressNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete wallet address no content response has a 5xx status code
func (o *DeleteWalletAddressNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete wallet address no content response a status code equal to that given
func (o *DeleteWalletAddressNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete wallet address no content response
func (o *DeleteWalletAddressNoContent) Code() int {
	return 204
}

func (o *DeleteWalletAddressNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] deleteWalletAddressNoContent ", 204)
}

func (o *DeleteWalletAddressNoContent) String() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] deleteWalletAddressNoContent ", 204)
}

func (o *DeleteWalletAddressNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWalletAddressBadRequest creates a DeleteWalletAddressBadRequest with default headers values
func NewDeleteWalletAddressBadRequest() *DeleteWalletAddressBadRequest {
	return &DeleteWalletAddressBadRequest{}
}

/*
DeleteWalletAddressBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteWalletAddressBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this delete wallet address bad request response has a 2xx status code
func (o *DeleteWalletAddressBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete wallet address bad request response has a 3xx status code
func (o *DeleteWalletAddressBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete wallet address bad request response has a 4xx status code
func (o *DeleteWalletAddressBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete wallet address bad request response has a 5xx status code
func (o *DeleteWalletAddressBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete wallet address bad request response a status code equal to that given
func (o *DeleteWalletAddressBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete wallet address bad request response
func (o *DeleteWalletAddressBadRequest) Code() int {
	return 400
}

func (o *DeleteWalletAddressBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] deleteWalletAddressBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWalletAddressBadRequest) String() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] deleteWalletAddressBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWalletAddressBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *DeleteWalletAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWalletAddressInternalServerError creates a DeleteWalletAddressInternalServerError with default headers values
func NewDeleteWalletAddressInternalServerError() *DeleteWalletAddressInternalServerError {
	return &DeleteWalletAddressInternalServerError{}
}

/*
DeleteWalletAddressInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteWalletAddressInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this delete wallet address internal server error response has a 2xx status code
func (o *DeleteWalletAddressInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete wallet address internal server error response has a 3xx status code
func (o *DeleteWalletAddressInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete wallet address internal server error response has a 4xx status code
func (o *DeleteWalletAddressInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete wallet address internal server error response has a 5xx status code
func (o *DeleteWalletAddressInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete wallet address internal server error response a status code equal to that given
func (o *DeleteWalletAddressInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete wallet address internal server error response
func (o *DeleteWalletAddressInternalServerError) Code() int {
	return 500
}

func (o *DeleteWalletAddressInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] deleteWalletAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWalletAddressInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] deleteWalletAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWalletAddressInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *DeleteWalletAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
