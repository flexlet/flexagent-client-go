// Code generated by go-swagger; DO NOT EDIT.

package crypto

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
)

// UpdateVaultReader is a Reader for the UpdateVault structure.
type UpdateVaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateVaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateVaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVaultOK creates a UpdateVaultOK with default headers values
func NewUpdateVaultOK() *UpdateVaultOK {
	return &UpdateVaultOK{}
}

/* UpdateVaultOK describes a response with status code 200, with default header values.

Update vault succeeded
*/
type UpdateVaultOK struct {
	Payload string
}

func (o *UpdateVaultOK) Error() string {
	return fmt.Sprintf("[PUT /crypto/vault/{name}][%d] updateVaultOK  %+v", 200, o.Payload)
}
func (o *UpdateVaultOK) GetPayload() string {
	return o.Payload
}

func (o *UpdateVaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVaultBadRequest creates a UpdateVaultBadRequest with default headers values
func NewUpdateVaultBadRequest() *UpdateVaultBadRequest {
	return &UpdateVaultBadRequest{}
}

/* UpdateVaultBadRequest describes a response with status code 400, with default header values.

UpdateVaultBadRequest update vault bad request
*/
type UpdateVaultBadRequest struct {
	Payload *UpdateVaultBadRequestBody
}

func (o *UpdateVaultBadRequest) Error() string {
	return fmt.Sprintf("[PUT /crypto/vault/{name}][%d] updateVaultBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVaultBadRequest) GetPayload() *UpdateVaultBadRequestBody {
	return o.Payload
}

func (o *UpdateVaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateVaultBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVaultUnauthorized creates a UpdateVaultUnauthorized with default headers values
func NewUpdateVaultUnauthorized() *UpdateVaultUnauthorized {
	return &UpdateVaultUnauthorized{}
}

/* UpdateVaultUnauthorized describes a response with status code 401, with default header values.

UpdateVaultUnauthorized update vault unauthorized
*/
type UpdateVaultUnauthorized struct {
	Payload interface{}
}

func (o *UpdateVaultUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /crypto/vault/{name}][%d] updateVaultUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateVaultUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateVaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVaultForbidden creates a UpdateVaultForbidden with default headers values
func NewUpdateVaultForbidden() *UpdateVaultForbidden {
	return &UpdateVaultForbidden{}
}

/* UpdateVaultForbidden describes a response with status code 403, with default header values.

UpdateVaultForbidden update vault forbidden
*/
type UpdateVaultForbidden struct {
	Payload *UpdateVaultForbiddenBody
}

func (o *UpdateVaultForbidden) Error() string {
	return fmt.Sprintf("[PUT /crypto/vault/{name}][%d] updateVaultForbidden  %+v", 403, o.Payload)
}
func (o *UpdateVaultForbidden) GetPayload() *UpdateVaultForbiddenBody {
	return o.Payload
}

func (o *UpdateVaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateVaultForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVaultNotFound creates a UpdateVaultNotFound with default headers values
func NewUpdateVaultNotFound() *UpdateVaultNotFound {
	return &UpdateVaultNotFound{}
}

/* UpdateVaultNotFound describes a response with status code 404, with default header values.

UpdateVaultNotFound update vault not found
*/
type UpdateVaultNotFound struct {
	Payload *UpdateVaultNotFoundBody
}

func (o *UpdateVaultNotFound) Error() string {
	return fmt.Sprintf("[PUT /crypto/vault/{name}][%d] updateVaultNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVaultNotFound) GetPayload() *UpdateVaultNotFoundBody {
	return o.Payload
}

func (o *UpdateVaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateVaultNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVaultInternalServerError creates a UpdateVaultInternalServerError with default headers values
func NewUpdateVaultInternalServerError() *UpdateVaultInternalServerError {
	return &UpdateVaultInternalServerError{}
}

/* UpdateVaultInternalServerError describes a response with status code 500, with default header values.

UpdateVaultInternalServerError update vault internal server error
*/
type UpdateVaultInternalServerError struct {
	Payload interface{}
}

func (o *UpdateVaultInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /crypto/vault/{name}][%d] updateVaultInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVaultInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateVaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateVaultBadRequestBody update vault bad request body
swagger:model UpdateVaultBadRequestBody
*/
type UpdateVaultBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update vault bad request body
func (o *UpdateVaultBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVaultBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update vault bad request body based on context it is used
func (o *UpdateVaultBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVaultBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVaultBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateVaultBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateVaultForbiddenBody update vault forbidden body
swagger:model UpdateVaultForbiddenBody
*/
type UpdateVaultForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update vault forbidden body
func (o *UpdateVaultForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVaultForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update vault forbidden body based on context it is used
func (o *UpdateVaultForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVaultForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVaultForbiddenBody) UnmarshalBinary(b []byte) error {
	var res UpdateVaultForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateVaultNotFoundBody update vault not found body
swagger:model UpdateVaultNotFoundBody
*/
type UpdateVaultNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this update vault not found body
func (o *UpdateVaultNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateVaultNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *UpdateVaultNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update vault not found body based on context it is used
func (o *UpdateVaultNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVaultNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVaultNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateVaultNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}