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

// StartCronJobReader is a Reader for the StartCronJob structure.
type StartCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartCronJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStartCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStartCronJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartCronJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartCronJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartCronJobOK creates a StartCronJobOK with default headers values
func NewStartCronJobOK() *StartCronJobOK {
	return &StartCronJobOK{}
}

/* StartCronJobOK describes a response with status code 200, with default header values.

Start cronjob succeeded
*/
type StartCronJobOK struct {
	Payload *models.CronJob
}

func (o *StartCronJobOK) Error() string {
	return fmt.Sprintf("[POST /cronjobs/{id}/start][%d] startCronJobOK  %+v", 200, o.Payload)
}
func (o *StartCronJobOK) GetPayload() *models.CronJob {
	return o.Payload
}

func (o *StartCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartCronJobBadRequest creates a StartCronJobBadRequest with default headers values
func NewStartCronJobBadRequest() *StartCronJobBadRequest {
	return &StartCronJobBadRequest{}
}

/* StartCronJobBadRequest describes a response with status code 400, with default header values.

StartCronJobBadRequest start cron job bad request
*/
type StartCronJobBadRequest struct {
	Payload *StartCronJobBadRequestBody
}

func (o *StartCronJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /cronjobs/{id}/start][%d] startCronJobBadRequest  %+v", 400, o.Payload)
}
func (o *StartCronJobBadRequest) GetPayload() *StartCronJobBadRequestBody {
	return o.Payload
}

func (o *StartCronJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartCronJobBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartCronJobUnauthorized creates a StartCronJobUnauthorized with default headers values
func NewStartCronJobUnauthorized() *StartCronJobUnauthorized {
	return &StartCronJobUnauthorized{}
}

/* StartCronJobUnauthorized describes a response with status code 401, with default header values.

StartCronJobUnauthorized start cron job unauthorized
*/
type StartCronJobUnauthorized struct {
	Payload interface{}
}

func (o *StartCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cronjobs/{id}/start][%d] startCronJobUnauthorized  %+v", 401, o.Payload)
}
func (o *StartCronJobUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StartCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartCronJobForbidden creates a StartCronJobForbidden with default headers values
func NewStartCronJobForbidden() *StartCronJobForbidden {
	return &StartCronJobForbidden{}
}

/* StartCronJobForbidden describes a response with status code 403, with default header values.

StartCronJobForbidden start cron job forbidden
*/
type StartCronJobForbidden struct {
	Payload *StartCronJobForbiddenBody
}

func (o *StartCronJobForbidden) Error() string {
	return fmt.Sprintf("[POST /cronjobs/{id}/start][%d] startCronJobForbidden  %+v", 403, o.Payload)
}
func (o *StartCronJobForbidden) GetPayload() *StartCronJobForbiddenBody {
	return o.Payload
}

func (o *StartCronJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartCronJobForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartCronJobNotFound creates a StartCronJobNotFound with default headers values
func NewStartCronJobNotFound() *StartCronJobNotFound {
	return &StartCronJobNotFound{}
}

/* StartCronJobNotFound describes a response with status code 404, with default header values.

StartCronJobNotFound start cron job not found
*/
type StartCronJobNotFound struct {
	Payload *StartCronJobNotFoundBody
}

func (o *StartCronJobNotFound) Error() string {
	return fmt.Sprintf("[POST /cronjobs/{id}/start][%d] startCronJobNotFound  %+v", 404, o.Payload)
}
func (o *StartCronJobNotFound) GetPayload() *StartCronJobNotFoundBody {
	return o.Payload
}

func (o *StartCronJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartCronJobNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartCronJobInternalServerError creates a StartCronJobInternalServerError with default headers values
func NewStartCronJobInternalServerError() *StartCronJobInternalServerError {
	return &StartCronJobInternalServerError{}
}

/* StartCronJobInternalServerError describes a response with status code 500, with default header values.

StartCronJobInternalServerError start cron job internal server error
*/
type StartCronJobInternalServerError struct {
	Payload interface{}
}

func (o *StartCronJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cronjobs/{id}/start][%d] startCronJobInternalServerError  %+v", 500, o.Payload)
}
func (o *StartCronJobInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *StartCronJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartCronJobBadRequestBody start cron job bad request body
swagger:model StartCronJobBadRequestBody
*/
type StartCronJobBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this start cron job bad request body
func (o *StartCronJobBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartCronJobBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start cron job bad request body based on context it is used
func (o *StartCronJobBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartCronJobBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartCronJobBadRequestBody) UnmarshalBinary(b []byte) error {
	var res StartCronJobBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartCronJobForbiddenBody start cron job forbidden body
swagger:model StartCronJobForbiddenBody
*/
type StartCronJobForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this start cron job forbidden body
func (o *StartCronJobForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartCronJobForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start cron job forbidden body based on context it is used
func (o *StartCronJobForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartCronJobForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartCronJobForbiddenBody) UnmarshalBinary(b []byte) error {
	var res StartCronJobForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartCronJobNotFoundBody start cron job not found body
swagger:model StartCronJobNotFoundBody
*/
type StartCronJobNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this start cron job not found body
func (o *StartCronJobNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *StartCronJobNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *StartCronJobNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start cron job not found body based on context it is used
func (o *StartCronJobNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartCronJobNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartCronJobNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StartCronJobNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
