// Code generated by go-swagger; DO NOT EDIT.

package cronjob

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
)

// NewStopCronJobParams creates a new StopCronJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopCronJobParams() *StopCronJobParams {
	return &StopCronJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopCronJobParamsWithTimeout creates a new StopCronJobParams object
// with the ability to set a timeout on a request.
func NewStopCronJobParamsWithTimeout(timeout time.Duration) *StopCronJobParams {
	return &StopCronJobParams{
		timeout: timeout,
	}
}

// NewStopCronJobParamsWithContext creates a new StopCronJobParams object
// with the ability to set a context for a request.
func NewStopCronJobParamsWithContext(ctx context.Context) *StopCronJobParams {
	return &StopCronJobParams{
		Context: ctx,
	}
}

// NewStopCronJobParamsWithHTTPClient creates a new StopCronJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopCronJobParamsWithHTTPClient(client *http.Client) *StopCronJobParams {
	return &StopCronJobParams{
		HTTPClient: client,
	}
}

/* StopCronJobParams contains all the parameters to send to the API endpoint
   for the stop cron job operation.

   Typically these are written to a http.Request.
*/
type StopCronJobParams struct {

	/* ID.

	   Cronjob id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopCronJobParams) WithDefaults() *StopCronJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopCronJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop cron job params
func (o *StopCronJobParams) WithTimeout(timeout time.Duration) *StopCronJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop cron job params
func (o *StopCronJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop cron job params
func (o *StopCronJobParams) WithContext(ctx context.Context) *StopCronJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop cron job params
func (o *StopCronJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop cron job params
func (o *StopCronJobParams) WithHTTPClient(client *http.Client) *StopCronJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop cron job params
func (o *StopCronJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stop cron job params
func (o *StopCronJobParams) WithID(id string) *StopCronJobParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stop cron job params
func (o *StopCronJobParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StopCronJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
