// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/petuhovskiy/grpc-hydra-bench/hydracli/models"
)

// AcceptConsentRequestReader is a Reader for the AcceptConsentRequest structure.
type AcceptConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptConsentRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAcceptConsentRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAcceptConsentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAcceptConsentRequestOK creates a AcceptConsentRequestOK with default headers values
func NewAcceptConsentRequestOK() *AcceptConsentRequestOK {
	return &AcceptConsentRequestOK{}
}

/*AcceptConsentRequestOK handles this case with default header values.

completedRequest
*/
type AcceptConsentRequestOK struct {
	Payload *models.RequestHandlerResponse
}

func (o *AcceptConsentRequestOK) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/consent/accept][%d] acceptConsentRequestOK  %+v", 200, o.Payload)
}

func (o *AcceptConsentRequestOK) GetPayload() *models.RequestHandlerResponse {
	return o.Payload
}

func (o *AcceptConsentRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestHandlerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptConsentRequestNotFound creates a AcceptConsentRequestNotFound with default headers values
func NewAcceptConsentRequestNotFound() *AcceptConsentRequestNotFound {
	return &AcceptConsentRequestNotFound{}
}

/*AcceptConsentRequestNotFound handles this case with default header values.

genericError
*/
type AcceptConsentRequestNotFound struct {
	Payload *models.GenericError
}

func (o *AcceptConsentRequestNotFound) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/consent/accept][%d] acceptConsentRequestNotFound  %+v", 404, o.Payload)
}

func (o *AcceptConsentRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *AcceptConsentRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptConsentRequestInternalServerError creates a AcceptConsentRequestInternalServerError with default headers values
func NewAcceptConsentRequestInternalServerError() *AcceptConsentRequestInternalServerError {
	return &AcceptConsentRequestInternalServerError{}
}

/*AcceptConsentRequestInternalServerError handles this case with default header values.

genericError
*/
type AcceptConsentRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *AcceptConsentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/consent/accept][%d] acceptConsentRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *AcceptConsentRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *AcceptConsentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
