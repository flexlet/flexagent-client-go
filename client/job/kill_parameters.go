// Code generated by go-swagger; DO NOT EDIT.

package job

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

// NewKillParams creates a new KillParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKillParams() *KillParams {
	return &KillParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKillParamsWithTimeout creates a new KillParams object
// with the ability to set a timeout on a request.
func NewKillParamsWithTimeout(timeout time.Duration) *KillParams {
	return &KillParams{
		timeout: timeout,
	}
}

// NewKillParamsWithContext creates a new KillParams object
// with the ability to set a context for a request.
func NewKillParamsWithContext(ctx context.Context) *KillParams {
	return &KillParams{
		Context: ctx,
	}
}

// NewKillParamsWithHTTPClient creates a new KillParams object
// with the ability to set a custom HTTPClient for a request.
func NewKillParamsWithHTTPClient(client *http.Client) *KillParams {
	return &KillParams{
		HTTPClient: client,
	}
}

/* KillParams contains all the parameters to send to the API endpoint
   for the kill operation.

   Typically these are written to a http.Request.
*/
type KillParams struct {

	/* Force.

	   Force kill
	*/
	Force *bool

	/* Urn.

	   Job urn
	*/
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KillParams) WithDefaults() *KillParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KillParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := KillParams{
		Force: &forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the kill params
func (o *KillParams) WithTimeout(timeout time.Duration) *KillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kill params
func (o *KillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kill params
func (o *KillParams) WithContext(ctx context.Context) *KillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kill params
func (o *KillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kill params
func (o *KillParams) WithHTTPClient(client *http.Client) *KillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kill params
func (o *KillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the kill params
func (o *KillParams) WithForce(force *bool) *KillParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the kill params
func (o *KillParams) SetForce(force *bool) {
	o.Force = force
}

// WithUrn adds the urn to the kill params
func (o *KillParams) WithUrn(urn string) *KillParams {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the kill params
func (o *KillParams) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *KillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param urn
	if err := r.SetPathParam("urn", o.Urn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
