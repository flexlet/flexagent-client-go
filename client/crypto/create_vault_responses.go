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

// CreateVaultReader is a Reader for the CreateVault structure.
type CreateVaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateVaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVaultOK creates a CreateVaultOK with default headers values
func NewCreateVaultOK() *CreateVaultOK {
	return &CreateVaultOK{}
}

/* CreateVaultOK describes a response with status code 200, with default header values.

Create vault succeeded
*/
type CreateVaultOK struct {
	Payload string
}

func (o *CreateVaultOK) Error() string {
	return fmt.Sprintf("[POST /crypto/vault/{name}][%d] createVaultOK  %+v", 200, o.Payload)
}
func (o *CreateVaultOK) GetPayload() string {
	return o.Payload
}

func (o *CreateVaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVaultBadRequest creates a CreateVaultBadRequest with default headers values
func NewCreateVaultBadRequest() *CreateVaultBadRequest {
	return &CreateVaultBadRequest{}
}

/* CreateVaultBadRequest describes a response with status code 400, with default header values.

CreateVaultBadRequest create vault bad request
*/
type CreateVaultBadRequest struct {
	Payload *CreateVaultBadRequestBody
}

func (o *CreateVaultBadRequest) Error() string {
	return fmt.Sprintf("[POST /crypto/vault/{name}][%d] createVaultBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVaultBadRequest) GetPayload() *CreateVaultBadRequestBody {
	return o.Payload
}

func (o *CreateVaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateVaultBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVaultUnauthorized creates a CreateVaultUnauthorized with default headers values
func NewCreateVaultUnauthorized() *CreateVaultUnauthorized {
	return &CreateVaultUnauthorized{}
}

/* CreateVaultUnauthorized describes a response with status code 401, with default header values.

CreateVaultUnauthorized create vault unauthorized
*/
type CreateVaultUnauthorized struct {
	Payload interface{}
}

func (o *CreateVaultUnauthorized) Error() string {
	return fmt.Sprintf("[POST /crypto/vault/{name}][%d] createVaultUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateVaultUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateVaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVaultForbidden creates a CreateVaultForbidden with default headers values
func NewCreateVaultForbidden() *CreateVaultForbidden {
	return &CreateVaultForbidden{}
}

/* CreateVaultForbidden describes a response with status code 403, with default header values.

CreateVaultForbidden create vault forbidden
*/
type CreateVaultForbidden struct {
	Payload *CreateVaultForbiddenBody
}

func (o *CreateVaultForbidden) Error() string {
	return fmt.Sprintf("[POST /crypto/vault/{name}][%d] createVaultForbidden  %+v", 403, o.Payload)
}
func (o *CreateVaultForbidden) GetPayload() *CreateVaultForbiddenBody {
	return o.Payload
}

func (o *CreateVaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateVaultForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVaultNotFound creates a CreateVaultNotFound with default headers values
func NewCreateVaultNotFound() *CreateVaultNotFound {
	return &CreateVaultNotFound{}
}

/* CreateVaultNotFound describes a response with status code 404, with default header values.

CreateVaultNotFound create vault not found
*/
type CreateVaultNotFound struct {
	Payload *CreateVaultNotFoundBody
}

func (o *CreateVaultNotFound) Error() string {
	return fmt.Sprintf("[POST /crypto/vault/{name}][%d] createVaultNotFound  %+v", 404, o.Payload)
}
func (o *CreateVaultNotFound) GetPayload() *CreateVaultNotFoundBody {
	return o.Payload
}

func (o *CreateVaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateVaultNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVaultInternalServerError creates a CreateVaultInternalServerError with default headers values
func NewCreateVaultInternalServerError() *CreateVaultInternalServerError {
	return &CreateVaultInternalServerError{}
}

/* CreateVaultInternalServerError describes a response with status code 500, with default header values.

CreateVaultInternalServerError create vault internal server error
*/
type CreateVaultInternalServerError struct {
	Payload interface{}
}

func (o *CreateVaultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /crypto/vault/{name}][%d] createVaultInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVaultInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateVaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateVaultBadRequestBody create vault bad request body
swagger:model CreateVaultBadRequestBody
*/
type CreateVaultBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create vault bad request body
func (o *CreateVaultBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVaultBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createVaultBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create vault bad request body based on context it is used
func (o *CreateVaultBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateVaultBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVaultBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateVaultBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateVaultForbiddenBody create vault forbidden body
swagger:model CreateVaultForbiddenBody
*/
type CreateVaultForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create vault forbidden body
func (o *CreateVaultForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVaultForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createVaultForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create vault forbidden body based on context it is used
func (o *CreateVaultForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateVaultForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVaultForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CreateVaultForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateVaultNotFoundBody create vault not found body
swagger:model CreateVaultNotFoundBody
*/
type CreateVaultNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this create vault not found body
func (o *CreateVaultNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateVaultNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("createVaultNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *CreateVaultNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("createVaultNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create vault not found body based on context it is used
func (o *CreateVaultNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateVaultNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVaultNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CreateVaultNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}