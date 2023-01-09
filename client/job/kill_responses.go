// Code generated by go-swagger; DO NOT EDIT.

package job

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

// KillReader is a Reader for the Kill structure.
type KillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKillBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKillUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKillInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKillOK creates a KillOK with default headers values
func NewKillOK() *KillOK {
	return &KillOK{}
}

/* KillOK describes a response with status code 200, with default header values.

Kill job succeeded
*/
type KillOK struct {
	Payload *models.Job
}

func (o *KillOK) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/kill][%d] killOK  %+v", 200, o.Payload)
}
func (o *KillOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *KillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillBadRequest creates a KillBadRequest with default headers values
func NewKillBadRequest() *KillBadRequest {
	return &KillBadRequest{}
}

/* KillBadRequest describes a response with status code 400, with default header values.

KillBadRequest kill bad request
*/
type KillBadRequest struct {
	Payload *KillBadRequestBody
}

func (o *KillBadRequest) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/kill][%d] killBadRequest  %+v", 400, o.Payload)
}
func (o *KillBadRequest) GetPayload() *KillBadRequestBody {
	return o.Payload
}

func (o *KillBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KillBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillUnauthorized creates a KillUnauthorized with default headers values
func NewKillUnauthorized() *KillUnauthorized {
	return &KillUnauthorized{}
}

/* KillUnauthorized describes a response with status code 401, with default header values.

KillUnauthorized kill unauthorized
*/
type KillUnauthorized struct {
	Payload interface{}
}

func (o *KillUnauthorized) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/kill][%d] killUnauthorized  %+v", 401, o.Payload)
}
func (o *KillUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KillUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillForbidden creates a KillForbidden with default headers values
func NewKillForbidden() *KillForbidden {
	return &KillForbidden{}
}

/* KillForbidden describes a response with status code 403, with default header values.

KillForbidden kill forbidden
*/
type KillForbidden struct {
	Payload *KillForbiddenBody
}

func (o *KillForbidden) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/kill][%d] killForbidden  %+v", 403, o.Payload)
}
func (o *KillForbidden) GetPayload() *KillForbiddenBody {
	return o.Payload
}

func (o *KillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KillForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillNotFound creates a KillNotFound with default headers values
func NewKillNotFound() *KillNotFound {
	return &KillNotFound{}
}

/* KillNotFound describes a response with status code 404, with default header values.

KillNotFound kill not found
*/
type KillNotFound struct {
	Payload *KillNotFoundBody
}

func (o *KillNotFound) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/kill][%d] killNotFound  %+v", 404, o.Payload)
}
func (o *KillNotFound) GetPayload() *KillNotFoundBody {
	return o.Payload
}

func (o *KillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KillNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillInternalServerError creates a KillInternalServerError with default headers values
func NewKillInternalServerError() *KillInternalServerError {
	return &KillInternalServerError{}
}

/* KillInternalServerError describes a response with status code 500, with default header values.

KillInternalServerError kill internal server error
*/
type KillInternalServerError struct {
	Payload interface{}
}

func (o *KillInternalServerError) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/kill][%d] killInternalServerError  %+v", 500, o.Payload)
}
func (o *KillInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *KillInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*KillBadRequestBody kill bad request body
swagger:model KillBadRequestBody
*/
type KillBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this kill bad request body
func (o *KillBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KillBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("killBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kill bad request body based on context it is used
func (o *KillBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KillBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KillBadRequestBody) UnmarshalBinary(b []byte) error {
	var res KillBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KillForbiddenBody kill forbidden body
swagger:model KillForbiddenBody
*/
type KillForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this kill forbidden body
func (o *KillForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KillForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("killForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kill forbidden body based on context it is used
func (o *KillForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KillForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KillForbiddenBody) UnmarshalBinary(b []byte) error {
	var res KillForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KillNotFoundBody kill not found body
swagger:model KillNotFoundBody
*/
type KillNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this kill not found body
func (o *KillNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *KillNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("killNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *KillNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("killNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kill not found body based on context it is used
func (o *KillNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KillNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KillNotFoundBody) UnmarshalBinary(b []byte) error {
	var res KillNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
