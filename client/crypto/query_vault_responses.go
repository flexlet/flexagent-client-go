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

// QueryVaultReader is a Reader for the QueryVault structure.
type QueryVaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryVaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryVaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryVaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryVaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryVaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryVaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryVaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryVaultOK creates a QueryVaultOK with default headers values
func NewQueryVaultOK() *QueryVaultOK {
	return &QueryVaultOK{}
}

/* QueryVaultOK describes a response with status code 200, with default header values.

Get vault succeed
*/
type QueryVaultOK struct {
	Payload map[string]string
}

func (o *QueryVaultOK) Error() string {
	return fmt.Sprintf("[GET /crypto/vault/{name}][%d] queryVaultOK  %+v", 200, o.Payload)
}
func (o *QueryVaultOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *QueryVaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVaultBadRequest creates a QueryVaultBadRequest with default headers values
func NewQueryVaultBadRequest() *QueryVaultBadRequest {
	return &QueryVaultBadRequest{}
}

/* QueryVaultBadRequest describes a response with status code 400, with default header values.

QueryVaultBadRequest query vault bad request
*/
type QueryVaultBadRequest struct {
	Payload *QueryVaultBadRequestBody
}

func (o *QueryVaultBadRequest) Error() string {
	return fmt.Sprintf("[GET /crypto/vault/{name}][%d] queryVaultBadRequest  %+v", 400, o.Payload)
}
func (o *QueryVaultBadRequest) GetPayload() *QueryVaultBadRequestBody {
	return o.Payload
}

func (o *QueryVaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryVaultBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVaultUnauthorized creates a QueryVaultUnauthorized with default headers values
func NewQueryVaultUnauthorized() *QueryVaultUnauthorized {
	return &QueryVaultUnauthorized{}
}

/* QueryVaultUnauthorized describes a response with status code 401, with default header values.

QueryVaultUnauthorized query vault unauthorized
*/
type QueryVaultUnauthorized struct {
	Payload interface{}
}

func (o *QueryVaultUnauthorized) Error() string {
	return fmt.Sprintf("[GET /crypto/vault/{name}][%d] queryVaultUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryVaultUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *QueryVaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVaultForbidden creates a QueryVaultForbidden with default headers values
func NewQueryVaultForbidden() *QueryVaultForbidden {
	return &QueryVaultForbidden{}
}

/* QueryVaultForbidden describes a response with status code 403, with default header values.

QueryVaultForbidden query vault forbidden
*/
type QueryVaultForbidden struct {
	Payload *QueryVaultForbiddenBody
}

func (o *QueryVaultForbidden) Error() string {
	return fmt.Sprintf("[GET /crypto/vault/{name}][%d] queryVaultForbidden  %+v", 403, o.Payload)
}
func (o *QueryVaultForbidden) GetPayload() *QueryVaultForbiddenBody {
	return o.Payload
}

func (o *QueryVaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryVaultForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVaultNotFound creates a QueryVaultNotFound with default headers values
func NewQueryVaultNotFound() *QueryVaultNotFound {
	return &QueryVaultNotFound{}
}

/* QueryVaultNotFound describes a response with status code 404, with default header values.

QueryVaultNotFound query vault not found
*/
type QueryVaultNotFound struct {
	Payload *QueryVaultNotFoundBody
}

func (o *QueryVaultNotFound) Error() string {
	return fmt.Sprintf("[GET /crypto/vault/{name}][%d] queryVaultNotFound  %+v", 404, o.Payload)
}
func (o *QueryVaultNotFound) GetPayload() *QueryVaultNotFoundBody {
	return o.Payload
}

func (o *QueryVaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryVaultNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVaultInternalServerError creates a QueryVaultInternalServerError with default headers values
func NewQueryVaultInternalServerError() *QueryVaultInternalServerError {
	return &QueryVaultInternalServerError{}
}

/* QueryVaultInternalServerError describes a response with status code 500, with default header values.

QueryVaultInternalServerError query vault internal server error
*/
type QueryVaultInternalServerError struct {
	Payload interface{}
}

func (o *QueryVaultInternalServerError) Error() string {
	return fmt.Sprintf("[GET /crypto/vault/{name}][%d] queryVaultInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryVaultInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *QueryVaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QueryVaultBadRequestBody query vault bad request body
swagger:model QueryVaultBadRequestBody
*/
type QueryVaultBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query vault bad request body
func (o *QueryVaultBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryVaultBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryVaultBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query vault bad request body based on context it is used
func (o *QueryVaultBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryVaultBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryVaultBadRequestBody) UnmarshalBinary(b []byte) error {
	var res QueryVaultBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QueryVaultForbiddenBody query vault forbidden body
swagger:model QueryVaultForbiddenBody
*/
type QueryVaultForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query vault forbidden body
func (o *QueryVaultForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryVaultForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryVaultForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query vault forbidden body based on context it is used
func (o *QueryVaultForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryVaultForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryVaultForbiddenBody) UnmarshalBinary(b []byte) error {
	var res QueryVaultForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QueryVaultNotFoundBody query vault not found body
swagger:model QueryVaultNotFoundBody
*/
type QueryVaultNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this query vault not found body
func (o *QueryVaultNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *QueryVaultNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("queryVaultNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *QueryVaultNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("queryVaultNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query vault not found body based on context it is used
func (o *QueryVaultNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryVaultNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryVaultNotFoundBody) UnmarshalBinary(b []byte) error {
	var res QueryVaultNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
