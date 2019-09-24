// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/sdk/go/hydra/models"
)

// RejectLogoutRequestReader is a Reader for the RejectLogoutRequest structure.
type RejectLogoutRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectLogoutRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRejectLogoutRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRejectLogoutRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRejectLogoutRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRejectLogoutRequestNoContent creates a RejectLogoutRequestNoContent with default headers values
func NewRejectLogoutRequestNoContent() *RejectLogoutRequestNoContent {
	return &RejectLogoutRequestNoContent{}
}

/*RejectLogoutRequestNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type RejectLogoutRequestNoContent struct {
}

func (o *RejectLogoutRequestNoContent) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/logout/reject][%d] rejectLogoutRequestNoContent ", 204)
}

func (o *RejectLogoutRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRejectLogoutRequestNotFound creates a RejectLogoutRequestNotFound with default headers values
func NewRejectLogoutRequestNotFound() *RejectLogoutRequestNotFound {
	return &RejectLogoutRequestNotFound{}
}

/*RejectLogoutRequestNotFound handles this case with default header values.

genericError
*/
type RejectLogoutRequestNotFound struct {
	Payload *models.GenericError
}

func (o *RejectLogoutRequestNotFound) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/logout/reject][%d] rejectLogoutRequestNotFound  %+v", 404, o.Payload)
}

func (o *RejectLogoutRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RejectLogoutRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectLogoutRequestInternalServerError creates a RejectLogoutRequestInternalServerError with default headers values
func NewRejectLogoutRequestInternalServerError() *RejectLogoutRequestInternalServerError {
	return &RejectLogoutRequestInternalServerError{}
}

/*RejectLogoutRequestInternalServerError handles this case with default header values.

genericError
*/
type RejectLogoutRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *RejectLogoutRequestInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/logout/reject][%d] rejectLogoutRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *RejectLogoutRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RejectLogoutRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
