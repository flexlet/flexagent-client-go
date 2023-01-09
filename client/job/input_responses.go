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

// InputReader is a Reader for the Input structure.
type InputReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InputReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInputOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInputBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInputUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInputForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInputNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInputInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInputOK creates a InputOK with default headers values
func NewInputOK() *InputOK {
	return &InputOK{}
}

/* InputOK describes a response with status code 200, with default header values.

Input job succeeded
*/
type InputOK struct {
	Payload *models.Job
}

func (o *InputOK) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/input][%d] inputOK  %+v", 200, o.Payload)
}
func (o *InputOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *InputOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInputBadRequest creates a InputBadRequest with default headers values
func NewInputBadRequest() *InputBadRequest {
	return &InputBadRequest{}
}

/* InputBadRequest describes a response with status code 400, with default header values.

InputBadRequest input bad request
*/
type InputBadRequest struct {
	Payload *InputBadRequestBody
}

func (o *InputBadRequest) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/input][%d] inputBadRequest  %+v", 400, o.Payload)
}
func (o *InputBadRequest) GetPayload() *InputBadRequestBody {
	return o.Payload
}

func (o *InputBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InputBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInputUnauthorized creates a InputUnauthorized with default headers values
func NewInputUnauthorized() *InputUnauthorized {
	return &InputUnauthorized{}
}

/* InputUnauthorized describes a response with status code 401, with default header values.

InputUnauthorized input unauthorized
*/
type InputUnauthorized struct {
	Payload interface{}
}

func (o *InputUnauthorized) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/input][%d] inputUnauthorized  %+v", 401, o.Payload)
}
func (o *InputUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *InputUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInputForbidden creates a InputForbidden with default headers values
func NewInputForbidden() *InputForbidden {
	return &InputForbidden{}
}

/* InputForbidden describes a response with status code 403, with default header values.

InputForbidden input forbidden
*/
type InputForbidden struct {
	Payload *InputForbiddenBody
}

func (o *InputForbidden) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/input][%d] inputForbidden  %+v", 403, o.Payload)
}
func (o *InputForbidden) GetPayload() *InputForbiddenBody {
	return o.Payload
}

func (o *InputForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InputForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInputNotFound creates a InputNotFound with default headers values
func NewInputNotFound() *InputNotFound {
	return &InputNotFound{}
}

/* InputNotFound describes a response with status code 404, with default header values.

InputNotFound input not found
*/
type InputNotFound struct {
	Payload *InputNotFoundBody
}

func (o *InputNotFound) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/input][%d] inputNotFound  %+v", 404, o.Payload)
}
func (o *InputNotFound) GetPayload() *InputNotFoundBody {
	return o.Payload
}

func (o *InputNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InputNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInputInternalServerError creates a InputInternalServerError with default headers values
func NewInputInternalServerError() *InputInternalServerError {
	return &InputInternalServerError{}
}

/* InputInternalServerError describes a response with status code 500, with default header values.

InputInternalServerError input internal server error
*/
type InputInternalServerError struct {
	Payload interface{}
}

func (o *InputInternalServerError) Error() string {
	return fmt.Sprintf("[POST /jobs/{urn}/input][%d] inputInternalServerError  %+v", 500, o.Payload)
}
func (o *InputInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *InputInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*InputBadRequestBody input bad request body
swagger:model InputBadRequestBody
*/
type InputBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this input bad request body
func (o *InputBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InputBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("inputBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this input bad request body based on context it is used
func (o *InputBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InputBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InputBadRequestBody) UnmarshalBinary(b []byte) error {
	var res InputBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InputForbiddenBody input forbidden body
swagger:model InputForbiddenBody
*/
type InputForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this input forbidden body
func (o *InputForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InputForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("inputForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this input forbidden body based on context it is used
func (o *InputForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InputForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InputForbiddenBody) UnmarshalBinary(b []byte) error {
	var res InputForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InputNotFoundBody input not found body
swagger:model InputNotFoundBody
*/
type InputNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this input not found body
func (o *InputNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *InputNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("inputNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *InputNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("inputNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this input not found body based on context it is used
func (o *InputNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InputNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InputNotFoundBody) UnmarshalBinary(b []byte) error {
	var res InputNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
