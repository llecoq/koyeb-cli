// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// SecretsUpdateSecret2Reader is a Reader for the SecretsUpdateSecret2 structure.
type SecretsUpdateSecret2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsUpdateSecret2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsUpdateSecret2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsUpdateSecret2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsUpdateSecret2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSecretsUpdateSecret2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSecretsUpdateSecret2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsUpdateSecret2OK creates a SecretsUpdateSecret2OK with default headers values
func NewSecretsUpdateSecret2OK() *SecretsUpdateSecret2OK {
	return &SecretsUpdateSecret2OK{}
}

/*SecretsUpdateSecret2OK handles this case with default header values.

A successful response.
*/
type SecretsUpdateSecret2OK struct {
	Payload *models.StorageUpdateSecretReply
}

func (o *SecretsUpdateSecret2OK) Error() string {
	return fmt.Sprintf("[PATCH /v1/secrets/{id}][%d] secretsUpdateSecret2OK  %+v", 200, o.Payload)
}

func (o *SecretsUpdateSecret2OK) GetPayload() *models.StorageUpdateSecretReply {
	return o.Payload
}

func (o *SecretsUpdateSecret2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageUpdateSecretReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsUpdateSecret2BadRequest creates a SecretsUpdateSecret2BadRequest with default headers values
func NewSecretsUpdateSecret2BadRequest() *SecretsUpdateSecret2BadRequest {
	return &SecretsUpdateSecret2BadRequest{}
}

/*SecretsUpdateSecret2BadRequest handles this case with default header values.

Validation error
*/
type SecretsUpdateSecret2BadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *SecretsUpdateSecret2BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/secrets/{id}][%d] secretsUpdateSecret2BadRequest  %+v", 400, o.Payload)
}

func (o *SecretsUpdateSecret2BadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *SecretsUpdateSecret2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsUpdateSecret2Forbidden creates a SecretsUpdateSecret2Forbidden with default headers values
func NewSecretsUpdateSecret2Forbidden() *SecretsUpdateSecret2Forbidden {
	return &SecretsUpdateSecret2Forbidden{}
}

/*SecretsUpdateSecret2Forbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type SecretsUpdateSecret2Forbidden struct {
	Payload *models.CommonError
}

func (o *SecretsUpdateSecret2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/secrets/{id}][%d] secretsUpdateSecret2Forbidden  %+v", 403, o.Payload)
}

func (o *SecretsUpdateSecret2Forbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *SecretsUpdateSecret2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsUpdateSecret2NotFound creates a SecretsUpdateSecret2NotFound with default headers values
func NewSecretsUpdateSecret2NotFound() *SecretsUpdateSecret2NotFound {
	return &SecretsUpdateSecret2NotFound{}
}

/*SecretsUpdateSecret2NotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type SecretsUpdateSecret2NotFound struct {
	Payload *models.CommonError
}

func (o *SecretsUpdateSecret2NotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/secrets/{id}][%d] secretsUpdateSecret2NotFound  %+v", 404, o.Payload)
}

func (o *SecretsUpdateSecret2NotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *SecretsUpdateSecret2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsUpdateSecret2Default creates a SecretsUpdateSecret2Default with default headers values
func NewSecretsUpdateSecret2Default(code int) *SecretsUpdateSecret2Default {
	return &SecretsUpdateSecret2Default{
		_statusCode: code,
	}
}

/*SecretsUpdateSecret2Default handles this case with default header values.

An unexpected error response
*/
type SecretsUpdateSecret2Default struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the secrets update secret2 default response
func (o *SecretsUpdateSecret2Default) Code() int {
	return o._statusCode
}

func (o *SecretsUpdateSecret2Default) Error() string {
	return fmt.Sprintf("[PATCH /v1/secrets/{id}][%d] Secrets_UpdateSecret2 default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsUpdateSecret2Default) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *SecretsUpdateSecret2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
