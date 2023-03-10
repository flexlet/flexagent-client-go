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

	"github.com/flexlet/flexagent-client-go/models"
)

// NewSubmitParams creates a new SubmitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitParams() *SubmitParams {
	return &SubmitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitParamsWithTimeout creates a new SubmitParams object
// with the ability to set a timeout on a request.
func NewSubmitParamsWithTimeout(timeout time.Duration) *SubmitParams {
	return &SubmitParams{
		timeout: timeout,
	}
}

// NewSubmitParamsWithContext creates a new SubmitParams object
// with the ability to set a context for a request.
func NewSubmitParamsWithContext(ctx context.Context) *SubmitParams {
	return &SubmitParams{
		Context: ctx,
	}
}

// NewSubmitParamsWithHTTPClient creates a new SubmitParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitParamsWithHTTPClient(client *http.Client) *SubmitParams {
	return &SubmitParams{
		HTTPClient: client,
	}
}

/* SubmitParams contains all the parameters to send to the API endpoint
   for the submit operation.

   Typically these are written to a http.Request.
*/
type SubmitParams struct {

	/* Spec.

	   Job spec
	*/
	Spec []*models.JobSpec

	/* Wait.

	   Wait until job finished
	*/
	Wait *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitParams) WithDefaults() *SubmitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitParams) SetDefaults() {
	var (
		waitDefault = bool(false)
	)

	val := SubmitParams{
		Wait: &waitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the submit params
func (o *SubmitParams) WithTimeout(timeout time.Duration) *SubmitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit params
func (o *SubmitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit params
func (o *SubmitParams) WithContext(ctx context.Context) *SubmitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit params
func (o *SubmitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit params
func (o *SubmitParams) WithHTTPClient(client *http.Client) *SubmitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit params
func (o *SubmitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSpec adds the spec to the submit params
func (o *SubmitParams) WithSpec(spec []*models.JobSpec) *SubmitParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the submit params
func (o *SubmitParams) SetSpec(spec []*models.JobSpec) {
	o.Spec = spec
}

// WithWait adds the wait to the submit params
func (o *SubmitParams) WithWait(wait *bool) *SubmitParams {
	o.SetWait(wait)
	return o
}

// SetWait adds the wait to the submit params
func (o *SubmitParams) SetWait(wait *bool) {
	o.Wait = wait
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Spec != nil {
		if err := r.SetBodyParam(o.Spec); err != nil {
			return err
		}
	}

	if o.Wait != nil {

		// query param wait
		var qrWait bool

		if o.Wait != nil {
			qrWait = *o.Wait
		}
		qWait := swag.FormatBool(qrWait)
		if qWait != "" {

			if err := r.SetQueryParam("wait", qWait); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
