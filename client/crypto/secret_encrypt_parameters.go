// Code generated by go-swagger; DO NOT EDIT.

package crypto

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/flexlet/flexagent-client-go/models"
)

// NewSecretEncryptParams creates a new SecretEncryptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretEncryptParams() *SecretEncryptParams {
	return &SecretEncryptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretEncryptParamsWithTimeout creates a new SecretEncryptParams object
// with the ability to set a timeout on a request.
func NewSecretEncryptParamsWithTimeout(timeout time.Duration) *SecretEncryptParams {
	return &SecretEncryptParams{
		timeout: timeout,
	}
}

// NewSecretEncryptParamsWithContext creates a new SecretEncryptParams object
// with the ability to set a context for a request.
func NewSecretEncryptParamsWithContext(ctx context.Context) *SecretEncryptParams {
	return &SecretEncryptParams{
		Context: ctx,
	}
}

// NewSecretEncryptParamsWithHTTPClient creates a new SecretEncryptParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretEncryptParamsWithHTTPClient(client *http.Client) *SecretEncryptParams {
	return &SecretEncryptParams{
		HTTPClient: client,
	}
}

/* SecretEncryptParams contains all the parameters to send to the API endpoint
   for the secret encrypt operation.

   Typically these are written to a http.Request.
*/
type SecretEncryptParams struct {

	/* Secret.

	   Secret to encrypt
	*/
	Secret *models.KubeSecret

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret encrypt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretEncryptParams) WithDefaults() *SecretEncryptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret encrypt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretEncryptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret encrypt params
func (o *SecretEncryptParams) WithTimeout(timeout time.Duration) *SecretEncryptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret encrypt params
func (o *SecretEncryptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret encrypt params
func (o *SecretEncryptParams) WithContext(ctx context.Context) *SecretEncryptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret encrypt params
func (o *SecretEncryptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret encrypt params
func (o *SecretEncryptParams) WithHTTPClient(client *http.Client) *SecretEncryptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret encrypt params
func (o *SecretEncryptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSecret adds the secret to the secret encrypt params
func (o *SecretEncryptParams) WithSecret(secret *models.KubeSecret) *SecretEncryptParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the secret encrypt params
func (o *SecretEncryptParams) SetSecret(secret *models.KubeSecret) {
	o.Secret = secret
}

// WriteToRequest writes these params to a swagger request
func (o *SecretEncryptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Secret != nil {
		if err := r.SetBodyParam(o.Secret); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
