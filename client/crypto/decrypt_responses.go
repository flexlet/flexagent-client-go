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

	"github.com/flexlet/flexagent-client-go/models"
)

// DecryptReader is a Reader for the Decrypt structure.
type DecryptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecryptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDecryptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDecryptBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDecryptUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDecryptForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDecryptNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDecryptInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDecryptOK creates a DecryptOK with default headers values
func NewDecryptOK() *DecryptOK {
	return &DecryptOK{}
}

/* DecryptOK describes a response with status code 200, with default header values.

Data decrypted
*/
type DecryptOK struct {
	Payload *models.CryptoData
}

func (o *DecryptOK) Error() string {
	return fmt.Sprintf("[POST /crypto/decrypt][%d] decryptOK  %+v", 200, o.Payload)
}
func (o *DecryptOK) GetPayload() *models.CryptoData {
	return o.Payload
}

func (o *DecryptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CryptoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptBadRequest creates a DecryptBadRequest with default headers values
func NewDecryptBadRequest() *DecryptBadRequest {
	return &DecryptBadRequest{}
}

/* DecryptBadRequest describes a response with status code 400, with default header values.

DecryptBadRequest decrypt bad request
*/
type DecryptBadRequest struct {
	Payload *DecryptBadRequestBody
}

func (o *DecryptBadRequest) Error() string {
	return fmt.Sprintf("[POST /crypto/decrypt][%d] decryptBadRequest  %+v", 400, o.Payload)
}
func (o *DecryptBadRequest) GetPayload() *DecryptBadRequestBody {
	return o.Payload
}

func (o *DecryptBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DecryptBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptUnauthorized creates a DecryptUnauthorized with default headers values
func NewDecryptUnauthorized() *DecryptUnauthorized {
	return &DecryptUnauthorized{}
}

/* DecryptUnauthorized describes a response with status code 401, with default header values.

DecryptUnauthorized decrypt unauthorized
*/
type DecryptUnauthorized struct {
	Payload interface{}
}

func (o *DecryptUnauthorized) Error() string {
	return fmt.Sprintf("[POST /crypto/decrypt][%d] decryptUnauthorized  %+v", 401, o.Payload)
}
func (o *DecryptUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DecryptUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptForbidden creates a DecryptForbidden with default headers values
func NewDecryptForbidden() *DecryptForbidden {
	return &DecryptForbidden{}
}

/* DecryptForbidden describes a response with status code 403, with default header values.

DecryptForbidden decrypt forbidden
*/
type DecryptForbidden struct {
	Payload *DecryptForbiddenBody
}

func (o *DecryptForbidden) Error() string {
	return fmt.Sprintf("[POST /crypto/decrypt][%d] decryptForbidden  %+v", 403, o.Payload)
}
func (o *DecryptForbidden) GetPayload() *DecryptForbiddenBody {
	return o.Payload
}

func (o *DecryptForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DecryptForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptNotFound creates a DecryptNotFound with default headers values
func NewDecryptNotFound() *DecryptNotFound {
	return &DecryptNotFound{}
}

/* DecryptNotFound describes a response with status code 404, with default header values.

DecryptNotFound decrypt not found
*/
type DecryptNotFound struct {
	Payload *DecryptNotFoundBody
}

func (o *DecryptNotFound) Error() string {
	return fmt.Sprintf("[POST /crypto/decrypt][%d] decryptNotFound  %+v", 404, o.Payload)
}
func (o *DecryptNotFound) GetPayload() *DecryptNotFoundBody {
	return o.Payload
}

func (o *DecryptNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DecryptNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptInternalServerError creates a DecryptInternalServerError with default headers values
func NewDecryptInternalServerError() *DecryptInternalServerError {
	return &DecryptInternalServerError{}
}

/* DecryptInternalServerError describes a response with status code 500, with default header values.

DecryptInternalServerError decrypt internal server error
*/
type DecryptInternalServerError struct {
	Payload interface{}
}

func (o *DecryptInternalServerError) Error() string {
	return fmt.Sprintf("[POST /crypto/decrypt][%d] decryptInternalServerError  %+v", 500, o.Payload)
}
func (o *DecryptInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *DecryptInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DecryptBadRequestBody decrypt bad request body
swagger:model DecryptBadRequestBody
*/
type DecryptBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this decrypt bad request body
func (o *DecryptBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DecryptBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("decryptBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this decrypt bad request body based on context it is used
func (o *DecryptBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DecryptBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DecryptBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DecryptBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DecryptForbiddenBody decrypt forbidden body
swagger:model DecryptForbiddenBody
*/
type DecryptForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this decrypt forbidden body
func (o *DecryptForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DecryptForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("decryptForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this decrypt forbidden body based on context it is used
func (o *DecryptForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DecryptForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DecryptForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DecryptForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DecryptNotFoundBody decrypt not found body
swagger:model DecryptNotFoundBody
*/
type DecryptNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this decrypt not found body
func (o *DecryptNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *DecryptNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("decryptNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *DecryptNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("decryptNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this decrypt not found body based on context it is used
func (o *DecryptNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DecryptNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DecryptNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DecryptNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}