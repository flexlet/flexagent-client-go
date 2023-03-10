// Code generated by go-swagger; DO NOT EDIT.

package cronjob

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/flexlet/flexagent-client-go/models"
)

// QueryCronJobReader is a Reader for the QueryCronJob structure.
type QueryCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCronJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCronJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCronJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCronJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryCronJobOK creates a QueryCronJobOK with default headers values
func NewQueryCronJobOK() *QueryCronJobOK {
	return &QueryCronJobOK{}
}

/* QueryCronJobOK describes a response with status code 200, with default header values.

Query cronjob succeeded
*/
type QueryCronJobOK struct {
	Payload *models.CronJob
}

func (o *QueryCronJobOK) Error() string {
	return fmt.Sprintf("[GET /cronjobs/{id}][%d] queryCronJobOK  %+v", 200, o.Payload)
}
func (o *QueryCronJobOK) GetPayload() *models.CronJob {
	return o.Payload
}

func (o *QueryCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCronJobBadRequest creates a QueryCronJobBadRequest with default headers values
func NewQueryCronJobBadRequest() *QueryCronJobBadRequest {
	return &QueryCronJobBadRequest{}
}

/* QueryCronJobBadRequest describes a response with status code 400, with default header values.

QueryCronJobBadRequest query cron job bad request
*/
type QueryCronJobBadRequest struct {
	Payload *QueryCronJobBadRequestBody
}

func (o *QueryCronJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /cronjobs/{id}][%d] queryCronJobBadRequest  %+v", 400, o.Payload)
}
func (o *QueryCronJobBadRequest) GetPayload() *QueryCronJobBadRequestBody {
	return o.Payload
}

func (o *QueryCronJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryCronJobBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCronJobUnauthorized creates a QueryCronJobUnauthorized with default headers values
func NewQueryCronJobUnauthorized() *QueryCronJobUnauthorized {
	return &QueryCronJobUnauthorized{}
}

/* QueryCronJobUnauthorized describes a response with status code 401, with default header values.

QueryCronJobUnauthorized query cron job unauthorized
*/
type QueryCronJobUnauthorized struct {
	Payload interface{}
}

func (o *QueryCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cronjobs/{id}][%d] queryCronJobUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryCronJobUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *QueryCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCronJobForbidden creates a QueryCronJobForbidden with default headers values
func NewQueryCronJobForbidden() *QueryCronJobForbidden {
	return &QueryCronJobForbidden{}
}

/* QueryCronJobForbidden describes a response with status code 403, with default header values.

QueryCronJobForbidden query cron job forbidden
*/
type QueryCronJobForbidden struct {
	Payload *QueryCronJobForbiddenBody
}

func (o *QueryCronJobForbidden) Error() string {
	return fmt.Sprintf("[GET /cronjobs/{id}][%d] queryCronJobForbidden  %+v", 403, o.Payload)
}
func (o *QueryCronJobForbidden) GetPayload() *QueryCronJobForbiddenBody {
	return o.Payload
}

func (o *QueryCronJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryCronJobForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCronJobNotFound creates a QueryCronJobNotFound with default headers values
func NewQueryCronJobNotFound() *QueryCronJobNotFound {
	return &QueryCronJobNotFound{}
}

/* QueryCronJobNotFound describes a response with status code 404, with default header values.

QueryCronJobNotFound query cron job not found
*/
type QueryCronJobNotFound struct {
	Payload *QueryCronJobNotFoundBody
}

func (o *QueryCronJobNotFound) Error() string {
	return fmt.Sprintf("[GET /cronjobs/{id}][%d] queryCronJobNotFound  %+v", 404, o.Payload)
}
func (o *QueryCronJobNotFound) GetPayload() *QueryCronJobNotFoundBody {
	return o.Payload
}

func (o *QueryCronJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryCronJobNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCronJobInternalServerError creates a QueryCronJobInternalServerError with default headers values
func NewQueryCronJobInternalServerError() *QueryCronJobInternalServerError {
	return &QueryCronJobInternalServerError{}
}

/* QueryCronJobInternalServerError describes a response with status code 500, with default header values.

QueryCronJobInternalServerError query cron job internal server error
*/
type QueryCronJobInternalServerError struct {
	Payload interface{}
}

func (o *QueryCronJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cronjobs/{id}][%d] queryCronJobInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryCronJobInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *QueryCronJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QueryCronJobBadRequestBody query cron job bad request body
swagger:model QueryCronJobBadRequestBody
*/
type QueryCronJobBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query cron job bad request body
func (o *QueryCronJobBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryCronJobBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query cron job bad request body based on context it is used
func (o *QueryCronJobBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryCronJobBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryCronJobBadRequestBody) UnmarshalBinary(b []byte) error {
	var res QueryCronJobBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QueryCronJobForbiddenBody query cron job forbidden body
swagger:model QueryCronJobForbiddenBody
*/
type QueryCronJobForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query cron job forbidden body
func (o *QueryCronJobForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryCronJobForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query cron job forbidden body based on context it is used
func (o *QueryCronJobForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryCronJobForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryCronJobForbiddenBody) UnmarshalBinary(b []byte) error {
	var res QueryCronJobForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QueryCronJobNotFoundBody query cron job not found body
swagger:model QueryCronJobNotFoundBody
*/
type QueryCronJobNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this query cron job not found body
func (o *QueryCronJobNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryCronJobNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *QueryCronJobNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query cron job not found body based on context it is used
func (o *QueryCronJobNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryCronJobNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryCronJobNotFoundBody) UnmarshalBinary(b []byte) error {
	var res QueryCronJobNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
