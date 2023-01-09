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
	"github.com/go-openapi/swag"
)

// NewQueryVaultParams creates a new QueryVaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryVaultParams() *QueryVaultParams {
	return &QueryVaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryVaultParamsWithTimeout creates a new QueryVaultParams object
// with the ability to set a timeout on a request.
func NewQueryVaultParamsWithTimeout(timeout time.Duration) *QueryVaultParams {
	return &QueryVaultParams{
		timeout: timeout,
	}
}

// NewQueryVaultParamsWithContext creates a new QueryVaultParams object
// with the ability to set a context for a request.
func NewQueryVaultParamsWithContext(ctx context.Context) *QueryVaultParams {
	return &QueryVaultParams{
		Context: ctx,
	}
}

// NewQueryVaultParamsWithHTTPClient creates a new QueryVaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryVaultParamsWithHTTPClient(client *http.Client) *QueryVaultParams {
	return &QueryVaultParams{
		HTTPClient: client,
	}
}

/* QueryVaultParams contains all the parameters to send to the API endpoint
   for the query vault operation.

   Typically these are written to a http.Request.
*/
type QueryVaultParams struct {

	/* Keys.

	   Key filter
	*/
	Keys []string

	/* Name.

	   Vault name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query vault params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryVaultParams) WithDefaults() *QueryVaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query vault params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryVaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query vault params
func (o *QueryVaultParams) WithTimeout(timeout time.Duration) *QueryVaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query vault params
func (o *QueryVaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query vault params
func (o *QueryVaultParams) WithContext(ctx context.Context) *QueryVaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query vault params
func (o *QueryVaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query vault params
func (o *QueryVaultParams) WithHTTPClient(client *http.Client) *QueryVaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query vault params
func (o *QueryVaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeys adds the keys to the query vault params
func (o *QueryVaultParams) WithKeys(keys []string) *QueryVaultParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the query vault params
func (o *QueryVaultParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithName adds the name to the query vault params
func (o *QueryVaultParams) WithName(name string) *QueryVaultParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the query vault params
func (o *QueryVaultParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *QueryVaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Keys != nil {

		// binding items for keys
		joinedKeys := o.bindParamKeys(reg)

		// query array param keys
		if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamQueryVault binds the parameter keys
func (o *QueryVaultParams) bindParamKeys(formats strfmt.Registry) []string {
	keysIR := o.Keys

	var keysIC []string
	for _, keysIIR := range keysIR { // explode []string

		keysIIV := keysIIR // string as string
		keysIC = append(keysIC, keysIIV)
	}

	// items.CollectionFormat: ""
	keysIS := swag.JoinByFormat(keysIC, "")

	return keysIS
}
