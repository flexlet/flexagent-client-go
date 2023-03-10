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

// ListCronJobsReader is a Reader for the ListCronJobs structure.
type ListCronJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCronJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCronJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCronJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCronJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCronJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCronJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCronJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCronJobsOK creates a ListCronJobsOK with default headers values
func NewListCronJobsOK() *ListCronJobsOK {
	return &ListCronJobsOK{}
}

/* ListCronJobsOK describes a response with status code 200, with default header values.

List cronjobs succeeded
*/
type ListCronJobsOK struct {
	Payload []*models.CronJob
}

func (o *ListCronJobsOK) Error() string {
	return fmt.Sprintf("[GET /cronjobs][%d] listCronJobsOK  %+v", 200, o.Payload)
}
func (o *ListCronJobsOK) GetPayload() []*models.CronJob {
	return o.Payload
}

func (o *ListCronJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCronJobsBadRequest creates a ListCronJobsBadRequest with default headers values
func NewListCronJobsBadRequest() *ListCronJobsBadRequest {
	return &ListCronJobsBadRequest{}
}

/* ListCronJobsBadRequest describes a response with status code 400, with default header values.

ListCronJobsBadRequest list cron jobs bad request
*/
type ListCronJobsBadRequest struct {
	Payload *ListCronJobsBadRequestBody
}

func (o *ListCronJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /cronjobs][%d] listCronJobsBadRequest  %+v", 400, o.Payload)
}
func (o *ListCronJobsBadRequest) GetPayload() *ListCronJobsBadRequestBody {
	return o.Payload
}

func (o *ListCronJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListCronJobsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCronJobsUnauthorized creates a ListCronJobsUnauthorized with default headers values
func NewListCronJobsUnauthorized() *ListCronJobsUnauthorized {
	return &ListCronJobsUnauthorized{}
}

/* ListCronJobsUnauthorized describes a response with status code 401, with default header values.

ListCronJobsUnauthorized list cron jobs unauthorized
*/
type ListCronJobsUnauthorized struct {
	Payload interface{}
}

func (o *ListCronJobsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cronjobs][%d] listCronJobsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListCronJobsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ListCronJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCronJobsForbidden creates a ListCronJobsForbidden with default headers values
func NewListCronJobsForbidden() *ListCronJobsForbidden {
	return &ListCronJobsForbidden{}
}

/* ListCronJobsForbidden describes a response with status code 403, with default header values.

ListCronJobsForbidden list cron jobs forbidden
*/
type ListCronJobsForbidden struct {
	Payload *ListCronJobsForbiddenBody
}

func (o *ListCronJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /cronjobs][%d] listCronJobsForbidden  %+v", 403, o.Payload)
}
func (o *ListCronJobsForbidden) GetPayload() *ListCronJobsForbiddenBody {
	return o.Payload
}

func (o *ListCronJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListCronJobsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCronJobsNotFound creates a ListCronJobsNotFound with default headers values
func NewListCronJobsNotFound() *ListCronJobsNotFound {
	return &ListCronJobsNotFound{}
}

/* ListCronJobsNotFound describes a response with status code 404, with default header values.

ListCronJobsNotFound list cron jobs not found
*/
type ListCronJobsNotFound struct {
	Payload *ListCronJobsNotFoundBody
}

func (o *ListCronJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /cronjobs][%d] listCronJobsNotFound  %+v", 404, o.Payload)
}
func (o *ListCronJobsNotFound) GetPayload() *ListCronJobsNotFoundBody {
	return o.Payload
}

func (o *ListCronJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListCronJobsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCronJobsInternalServerError creates a ListCronJobsInternalServerError with default headers values
func NewListCronJobsInternalServerError() *ListCronJobsInternalServerError {
	return &ListCronJobsInternalServerError{}
}

/* ListCronJobsInternalServerError describes a response with status code 500, with default header values.

ListCronJobsInternalServerError list cron jobs internal server error
*/
type ListCronJobsInternalServerError struct {
	Payload interface{}
}

func (o *ListCronJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cronjobs][%d] listCronJobsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListCronJobsInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ListCronJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListCronJobsBadRequestBody list cron jobs bad request body
swagger:model ListCronJobsBadRequestBody
*/
type ListCronJobsBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this list cron jobs bad request body
func (o *ListCronJobsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListCronJobsBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("listCronJobsBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list cron jobs bad request body based on context it is used
func (o *ListCronJobsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListCronJobsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListCronJobsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ListCronJobsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListCronJobsForbiddenBody list cron jobs forbidden body
swagger:model ListCronJobsForbiddenBody
*/
type ListCronJobsForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this list cron jobs forbidden body
func (o *ListCronJobsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListCronJobsForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("listCronJobsForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list cron jobs forbidden body based on context it is used
func (o *ListCronJobsForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListCronJobsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListCronJobsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ListCronJobsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListCronJobsNotFoundBody list cron jobs not found body
swagger:model ListCronJobsNotFoundBody
*/
type ListCronJobsNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this list cron jobs not found body
func (o *ListCronJobsNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *ListCronJobsNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("listCronJobsNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *ListCronJobsNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("listCronJobsNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list cron jobs not found body based on context it is used
func (o *ListCronJobsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListCronJobsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListCronJobsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListCronJobsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
