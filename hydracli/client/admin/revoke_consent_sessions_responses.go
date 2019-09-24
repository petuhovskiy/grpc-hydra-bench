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

// RevokeConsentSessionsReader is a Reader for the RevokeConsentSessions structure.
type RevokeConsentSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeConsentSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRevokeConsentSessionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeConsentSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeConsentSessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRevokeConsentSessionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRevokeConsentSessionsNoContent creates a RevokeConsentSessionsNoContent with default headers values
func NewRevokeConsentSessionsNoContent() *RevokeConsentSessionsNoContent {
	return &RevokeConsentSessionsNoContent{}
}

/*RevokeConsentSessionsNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type RevokeConsentSessionsNoContent struct {
}

func (o *RevokeConsentSessionsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/consent][%d] revokeConsentSessionsNoContent ", 204)
}

func (o *RevokeConsentSessionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeConsentSessionsBadRequest creates a RevokeConsentSessionsBadRequest with default headers values
func NewRevokeConsentSessionsBadRequest() *RevokeConsentSessionsBadRequest {
	return &RevokeConsentSessionsBadRequest{}
}

/*RevokeConsentSessionsBadRequest handles this case with default header values.

genericError
*/
type RevokeConsentSessionsBadRequest struct {
	Payload *models.GenericError
}

func (o *RevokeConsentSessionsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/consent][%d] revokeConsentSessionsBadRequest  %+v", 400, o.Payload)
}

func (o *RevokeConsentSessionsBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RevokeConsentSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeConsentSessionsNotFound creates a RevokeConsentSessionsNotFound with default headers values
func NewRevokeConsentSessionsNotFound() *RevokeConsentSessionsNotFound {
	return &RevokeConsentSessionsNotFound{}
}

/*RevokeConsentSessionsNotFound handles this case with default header values.

genericError
*/
type RevokeConsentSessionsNotFound struct {
	Payload *models.GenericError
}

func (o *RevokeConsentSessionsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/consent][%d] revokeConsentSessionsNotFound  %+v", 404, o.Payload)
}

func (o *RevokeConsentSessionsNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RevokeConsentSessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeConsentSessionsInternalServerError creates a RevokeConsentSessionsInternalServerError with default headers values
func NewRevokeConsentSessionsInternalServerError() *RevokeConsentSessionsInternalServerError {
	return &RevokeConsentSessionsInternalServerError{}
}

/*RevokeConsentSessionsInternalServerError handles this case with default header values.

genericError
*/
type RevokeConsentSessionsInternalServerError struct {
	Payload *models.GenericError
}

func (o *RevokeConsentSessionsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/consent][%d] revokeConsentSessionsInternalServerError  %+v", 500, o.Payload)
}

func (o *RevokeConsentSessionsInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RevokeConsentSessionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
