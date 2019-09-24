// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/petuhovskiy/grpc-hydra-bench/hydracli/models"
)

// NewAcceptConsentRequestParams creates a new AcceptConsentRequestParams object
// with the default values initialized.
func NewAcceptConsentRequestParams() *AcceptConsentRequestParams {
	var ()
	return &AcceptConsentRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptConsentRequestParamsWithTimeout creates a new AcceptConsentRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAcceptConsentRequestParamsWithTimeout(timeout time.Duration) *AcceptConsentRequestParams {
	var ()
	return &AcceptConsentRequestParams{

		timeout: timeout,
	}
}

// NewAcceptConsentRequestParamsWithContext creates a new AcceptConsentRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewAcceptConsentRequestParamsWithContext(ctx context.Context) *AcceptConsentRequestParams {
	var ()
	return &AcceptConsentRequestParams{

		Context: ctx,
	}
}

// NewAcceptConsentRequestParamsWithHTTPClient creates a new AcceptConsentRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAcceptConsentRequestParamsWithHTTPClient(client *http.Client) *AcceptConsentRequestParams {
	var ()
	return &AcceptConsentRequestParams{
		HTTPClient: client,
	}
}

/*AcceptConsentRequestParams contains all the parameters to send to the API endpoint
for the accept consent request operation typically these are written to a http.Request
*/
type AcceptConsentRequestParams struct {

	/*Body*/
	Body *models.HandledConsentRequest
	/*ConsentChallenge*/
	ConsentChallenge string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the accept consent request params
func (o *AcceptConsentRequestParams) WithTimeout(timeout time.Duration) *AcceptConsentRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept consent request params
func (o *AcceptConsentRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept consent request params
func (o *AcceptConsentRequestParams) WithContext(ctx context.Context) *AcceptConsentRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept consent request params
func (o *AcceptConsentRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept consent request params
func (o *AcceptConsentRequestParams) WithHTTPClient(client *http.Client) *AcceptConsentRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept consent request params
func (o *AcceptConsentRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the accept consent request params
func (o *AcceptConsentRequestParams) WithBody(body *models.HandledConsentRequest) *AcceptConsentRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the accept consent request params
func (o *AcceptConsentRequestParams) SetBody(body *models.HandledConsentRequest) {
	o.Body = body
}

// WithConsentChallenge adds the consentChallenge to the accept consent request params
func (o *AcceptConsentRequestParams) WithConsentChallenge(consentChallenge string) *AcceptConsentRequestParams {
	o.SetConsentChallenge(consentChallenge)
	return o
}

// SetConsentChallenge adds the consentChallenge to the accept consent request params
func (o *AcceptConsentRequestParams) SetConsentChallenge(consentChallenge string) {
	o.ConsentChallenge = consentChallenge
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptConsentRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param consent_challenge
	qrConsentChallenge := o.ConsentChallenge
	qConsentChallenge := qrConsentChallenge
	if qConsentChallenge != "" {
		if err := r.SetQueryParam("consent_challenge", qConsentChallenge); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
