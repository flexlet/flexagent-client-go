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

// NewEncryptParams creates a new EncryptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEncryptParams() *EncryptParams {
	return &EncryptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEncryptParamsWithTimeout creates a new EncryptParams object
// with the ability to set a timeout on a request.
func NewEncryptParamsWithTimeout(timeout time.Duration) *EncryptParams {
	return &EncryptParams{
		timeout: timeout,
	}
}

// NewEncryptParamsWithContext creates a new EncryptParams object
// with the ability to set a context for a request.
func NewEncryptParamsWithContext(ctx context.Context) *EncryptParams {
	return &EncryptParams{
		Context: ctx,
	}
}

// NewEncryptParamsWithHTTPClient creates a new EncryptParams object
// with the ability to set a custom HTTPClient for a request.
func NewEncryptParamsWithHTTPClient(client *http.Client) *EncryptParams {
	return &EncryptParams{
		HTTPClient: client,
	}
}

/* EncryptParams contains all the parameters to send to the API endpoint
   for the encrypt operation.

   Typically these are written to a http.Request.
*/
type EncryptParams struct {

	/* Data.

	   Data to encrypt
	*/
	Data *models.CryptoData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the encrypt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EncryptParams) WithDefaults() *EncryptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the encrypt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EncryptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the encrypt params
func (o *EncryptParams) WithTimeout(timeout time.Duration) *EncryptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the encrypt params
func (o *EncryptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the encrypt params
func (o *EncryptParams) WithContext(ctx context.Context) *EncryptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the encrypt params
func (o *EncryptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the encrypt params
func (o *EncryptParams) WithHTTPClient(client *http.Client) *EncryptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the encrypt params
func (o *EncryptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the encrypt params
func (o *EncryptParams) WithData(data *models.CryptoData) *EncryptParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the encrypt params
func (o *EncryptParams) SetData(data *models.CryptoData) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *EncryptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
