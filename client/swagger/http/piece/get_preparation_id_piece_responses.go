// Code generated by go-swagger; DO NOT EDIT.

package piece

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// GetPreparationIDPieceReader is a Reader for the GetPreparationIDPiece structure.
type GetPreparationIDPieceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPreparationIDPieceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPreparationIDPieceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPreparationIDPieceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPreparationIDPieceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /preparation/{id}/piece] GetPreparationIDPiece", response, response.Code())
	}
}

// NewGetPreparationIDPieceOK creates a GetPreparationIDPieceOK with default headers values
func NewGetPreparationIDPieceOK() *GetPreparationIDPieceOK {
	return &GetPreparationIDPieceOK{}
}

/*
GetPreparationIDPieceOK describes a response with status code 200, with default header values.

OK
*/
type GetPreparationIDPieceOK struct {
	Payload []*models.DataprepPieceList
}

// IsSuccess returns true when this get preparation Id piece o k response has a 2xx status code
func (o *GetPreparationIDPieceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get preparation Id piece o k response has a 3xx status code
func (o *GetPreparationIDPieceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preparation Id piece o k response has a 4xx status code
func (o *GetPreparationIDPieceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get preparation Id piece o k response has a 5xx status code
func (o *GetPreparationIDPieceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get preparation Id piece o k response a status code equal to that given
func (o *GetPreparationIDPieceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get preparation Id piece o k response
func (o *GetPreparationIDPieceOK) Code() int {
	return 200
}

func (o *GetPreparationIDPieceOK) Error() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] getPreparationIdPieceOK  %+v", 200, o.Payload)
}

func (o *GetPreparationIDPieceOK) String() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] getPreparationIdPieceOK  %+v", 200, o.Payload)
}

func (o *GetPreparationIDPieceOK) GetPayload() []*models.DataprepPieceList {
	return o.Payload
}

func (o *GetPreparationIDPieceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreparationIDPieceBadRequest creates a GetPreparationIDPieceBadRequest with default headers values
func NewGetPreparationIDPieceBadRequest() *GetPreparationIDPieceBadRequest {
	return &GetPreparationIDPieceBadRequest{}
}

/*
GetPreparationIDPieceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPreparationIDPieceBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get preparation Id piece bad request response has a 2xx status code
func (o *GetPreparationIDPieceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get preparation Id piece bad request response has a 3xx status code
func (o *GetPreparationIDPieceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preparation Id piece bad request response has a 4xx status code
func (o *GetPreparationIDPieceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get preparation Id piece bad request response has a 5xx status code
func (o *GetPreparationIDPieceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get preparation Id piece bad request response a status code equal to that given
func (o *GetPreparationIDPieceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get preparation Id piece bad request response
func (o *GetPreparationIDPieceBadRequest) Code() int {
	return 400
}

func (o *GetPreparationIDPieceBadRequest) Error() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] getPreparationIdPieceBadRequest  %+v", 400, o.Payload)
}

func (o *GetPreparationIDPieceBadRequest) String() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] getPreparationIdPieceBadRequest  %+v", 400, o.Payload)
}

func (o *GetPreparationIDPieceBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetPreparationIDPieceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreparationIDPieceInternalServerError creates a GetPreparationIDPieceInternalServerError with default headers values
func NewGetPreparationIDPieceInternalServerError() *GetPreparationIDPieceInternalServerError {
	return &GetPreparationIDPieceInternalServerError{}
}

/*
GetPreparationIDPieceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPreparationIDPieceInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get preparation Id piece internal server error response has a 2xx status code
func (o *GetPreparationIDPieceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get preparation Id piece internal server error response has a 3xx status code
func (o *GetPreparationIDPieceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preparation Id piece internal server error response has a 4xx status code
func (o *GetPreparationIDPieceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get preparation Id piece internal server error response has a 5xx status code
func (o *GetPreparationIDPieceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get preparation Id piece internal server error response a status code equal to that given
func (o *GetPreparationIDPieceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get preparation Id piece internal server error response
func (o *GetPreparationIDPieceInternalServerError) Code() int {
	return 500
}

func (o *GetPreparationIDPieceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] getPreparationIdPieceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreparationIDPieceInternalServerError) String() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] getPreparationIdPieceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreparationIDPieceInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetPreparationIDPieceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
