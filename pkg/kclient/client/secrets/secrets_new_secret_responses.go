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

// SecretsNewSecretReader is a Reader for the SecretsNewSecret structure.
type SecretsNewSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsNewSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsNewSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsNewSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsNewSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSecretsNewSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSecretsNewSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsNewSecretOK creates a SecretsNewSecretOK with default headers values
func NewSecretsNewSecretOK() *SecretsNewSecretOK {
	return &SecretsNewSecretOK{}
}

/*SecretsNewSecretOK handles this case with default header values.

A successful response.
*/
type SecretsNewSecretOK struct {
	Payload *models.StorageNewSecretReply
}

func (o *SecretsNewSecretOK) Error() string {
	return fmt.Sprintf("[POST /v1/secrets][%d] secretsNewSecretOK  %+v", 200, o.Payload)
}

func (o *SecretsNewSecretOK) GetPayload() *models.StorageNewSecretReply {
	return o.Payload
}

func (o *SecretsNewSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageNewSecretReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsNewSecretBadRequest creates a SecretsNewSecretBadRequest with default headers values
func NewSecretsNewSecretBadRequest() *SecretsNewSecretBadRequest {
	return &SecretsNewSecretBadRequest{}
}

/*SecretsNewSecretBadRequest handles this case with default header values.

Validation error
*/
type SecretsNewSecretBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *SecretsNewSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/secrets][%d] secretsNewSecretBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsNewSecretBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *SecretsNewSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsNewSecretForbidden creates a SecretsNewSecretForbidden with default headers values
func NewSecretsNewSecretForbidden() *SecretsNewSecretForbidden {
	return &SecretsNewSecretForbidden{}
}

/*SecretsNewSecretForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type SecretsNewSecretForbidden struct {
	Payload *models.CommonError
}

func (o *SecretsNewSecretForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/secrets][%d] secretsNewSecretForbidden  %+v", 403, o.Payload)
}

func (o *SecretsNewSecretForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *SecretsNewSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsNewSecretNotFound creates a SecretsNewSecretNotFound with default headers values
func NewSecretsNewSecretNotFound() *SecretsNewSecretNotFound {
	return &SecretsNewSecretNotFound{}
}

/*SecretsNewSecretNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type SecretsNewSecretNotFound struct {
	Payload *models.CommonError
}

func (o *SecretsNewSecretNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/secrets][%d] secretsNewSecretNotFound  %+v", 404, o.Payload)
}

func (o *SecretsNewSecretNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *SecretsNewSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsNewSecretDefault creates a SecretsNewSecretDefault with default headers values
func NewSecretsNewSecretDefault(code int) *SecretsNewSecretDefault {
	return &SecretsNewSecretDefault{
		_statusCode: code,
	}
}

/*SecretsNewSecretDefault handles this case with default header values.

An unexpected error response
*/
type SecretsNewSecretDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the secrets new secret default response
func (o *SecretsNewSecretDefault) Code() int {
	return o._statusCode
}

func (o *SecretsNewSecretDefault) Error() string {
	return fmt.Sprintf("[POST /v1/secrets][%d] Secrets_NewSecret default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsNewSecretDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *SecretsNewSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
