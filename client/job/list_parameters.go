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

// NewListParams creates a new ListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListParams() *ListParams {
	return &ListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListParamsWithTimeout creates a new ListParams object
// with the ability to set a timeout on a request.
func NewListParamsWithTimeout(timeout time.Duration) *ListParams {
	return &ListParams{
		timeout: timeout,
	}
}

// NewListParamsWithContext creates a new ListParams object
// with the ability to set a context for a request.
func NewListParamsWithContext(ctx context.Context) *ListParams {
	return &ListParams{
		Context: ctx,
	}
}

// NewListParamsWithHTTPClient creates a new ListParams object
// with the ability to set a custom HTTPClient for a request.
func NewListParamsWithHTTPClient(client *http.Client) *ListParams {
	return &ListParams{
		HTTPClient: client,
	}
}

/* ListParams contains all the parameters to send to the API endpoint
   for the list operation.

   Typically these are written to a http.Request.
*/
type ListParams struct {

	/* Operation.

	   Operation
	*/
	Operation *string

	/* Plugin.

	   Plugin
	*/
	Plugin *string

	/* StartTimeBegin.

	   Start time range begin

	   Format: int64
	*/
	StartTimeBegin *int64

	/* StartTimeEnd.

	   Start time range end

	   Format: int64
	*/
	StartTimeEnd *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListParams) WithDefaults() *ListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list params
func (o *ListParams) WithTimeout(timeout time.Duration) *ListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list params
func (o *ListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list params
func (o *ListParams) WithContext(ctx context.Context) *ListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list params
func (o *ListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list params
func (o *ListParams) WithHTTPClient(client *http.Client) *ListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list params
func (o *ListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperation adds the operation to the list params
func (o *ListParams) WithOperation(operation *string) *ListParams {
	o.SetOperation(operation)
	return o
}

// SetOperation adds the operation to the list params
func (o *ListParams) SetOperation(operation *string) {
	o.Operation = operation
}

// WithPlugin adds the plugin to the list params
func (o *ListParams) WithPlugin(plugin *string) *ListParams {
	o.SetPlugin(plugin)
	return o
}

// SetPlugin adds the plugin to the list params
func (o *ListParams) SetPlugin(plugin *string) {
	o.Plugin = plugin
}

// WithStartTimeBegin adds the startTimeBegin to the list params
func (o *ListParams) WithStartTimeBegin(startTimeBegin *int64) *ListParams {
	o.SetStartTimeBegin(startTimeBegin)
	return o
}

// SetStartTimeBegin adds the startTimeBegin to the list params
func (o *ListParams) SetStartTimeBegin(startTimeBegin *int64) {
	o.StartTimeBegin = startTimeBegin
}

// WithStartTimeEnd adds the startTimeEnd to the list params
func (o *ListParams) WithStartTimeEnd(startTimeEnd *int64) *ListParams {
	o.SetStartTimeEnd(startTimeEnd)
	return o
}

// SetStartTimeEnd adds the startTimeEnd to the list params
func (o *ListParams) SetStartTimeEnd(startTimeEnd *int64) {
	o.StartTimeEnd = startTimeEnd
}

// WriteToRequest writes these params to a swagger request
func (o *ListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Operation != nil {

		// query param operation
		var qrOperation string

		if o.Operation != nil {
			qrOperation = *o.Operation
		}
		qOperation := qrOperation
		if qOperation != "" {

			if err := r.SetQueryParam("operation", qOperation); err != nil {
				return err
			}
		}
	}

	if o.Plugin != nil {

		// query param plugin
		var qrPlugin string

		if o.Plugin != nil {
			qrPlugin = *o.Plugin
		}
		qPlugin := qrPlugin
		if qPlugin != "" {

			if err := r.SetQueryParam("plugin", qPlugin); err != nil {
				return err
			}
		}
	}

	if o.StartTimeBegin != nil {

		// query param startTimeBegin
		var qrStartTimeBegin int64

		if o.StartTimeBegin != nil {
			qrStartTimeBegin = *o.StartTimeBegin
		}
		qStartTimeBegin := swag.FormatInt64(qrStartTimeBegin)
		if qStartTimeBegin != "" {

			if err := r.SetQueryParam("startTimeBegin", qStartTimeBegin); err != nil {
				return err
			}
		}
	}

	if o.StartTimeEnd != nil {

		// query param startTimeEnd
		var qrStartTimeEnd int64

		if o.StartTimeEnd != nil {
			qrStartTimeEnd = *o.StartTimeEnd
		}
		qStartTimeEnd := swag.FormatInt64(qrStartTimeEnd)
		if qStartTimeEnd != "" {

			if err := r.SetQueryParam("startTimeEnd", qStartTimeEnd); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
