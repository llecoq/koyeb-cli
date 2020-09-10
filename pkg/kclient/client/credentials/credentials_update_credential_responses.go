// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// CredentialsUpdateCredentialReader is a Reader for the CredentialsUpdateCredential structure.
type CredentialsUpdateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsUpdateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialsUpdateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCredentialsUpdateCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCredentialsUpdateCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCredentialsUpdateCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCredentialsUpdateCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCredentialsUpdateCredentialOK creates a CredentialsUpdateCredentialOK with default headers values
func NewCredentialsUpdateCredentialOK() *CredentialsUpdateCredentialOK {
	return &CredentialsUpdateCredentialOK{}
}

/*CredentialsUpdateCredentialOK handles this case with default header values.

A successful response.
*/
type CredentialsUpdateCredentialOK struct {
	Payload *models.AccountCredentialReply
}

func (o *CredentialsUpdateCredentialOK) Error() string {
	return fmt.Sprintf("[PUT /v1/credentials/{id}][%d] credentialsUpdateCredentialOK  %+v", 200, o.Payload)
}

func (o *CredentialsUpdateCredentialOK) GetPayload() *models.AccountCredentialReply {
	return o.Payload
}

func (o *CredentialsUpdateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountCredentialReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCredentialsUpdateCredentialBadRequest creates a CredentialsUpdateCredentialBadRequest with default headers values
func NewCredentialsUpdateCredentialBadRequest() *CredentialsUpdateCredentialBadRequest {
	return &CredentialsUpdateCredentialBadRequest{}
}

/*CredentialsUpdateCredentialBadRequest handles this case with default header values.

Validation error
*/
type CredentialsUpdateCredentialBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *CredentialsUpdateCredentialBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/credentials/{id}][%d] credentialsUpdateCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *CredentialsUpdateCredentialBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *CredentialsUpdateCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCredentialsUpdateCredentialForbidden creates a CredentialsUpdateCredentialForbidden with default headers values
func NewCredentialsUpdateCredentialForbidden() *CredentialsUpdateCredentialForbidden {
	return &CredentialsUpdateCredentialForbidden{}
}

/*CredentialsUpdateCredentialForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type CredentialsUpdateCredentialForbidden struct {
	Payload *models.CommonError
}

func (o *CredentialsUpdateCredentialForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/credentials/{id}][%d] credentialsUpdateCredentialForbidden  %+v", 403, o.Payload)
}

func (o *CredentialsUpdateCredentialForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *CredentialsUpdateCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCredentialsUpdateCredentialNotFound creates a CredentialsUpdateCredentialNotFound with default headers values
func NewCredentialsUpdateCredentialNotFound() *CredentialsUpdateCredentialNotFound {
	return &CredentialsUpdateCredentialNotFound{}
}

/*CredentialsUpdateCredentialNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type CredentialsUpdateCredentialNotFound struct {
	Payload *models.CommonError
}

func (o *CredentialsUpdateCredentialNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/credentials/{id}][%d] credentialsUpdateCredentialNotFound  %+v", 404, o.Payload)
}

func (o *CredentialsUpdateCredentialNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *CredentialsUpdateCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCredentialsUpdateCredentialDefault creates a CredentialsUpdateCredentialDefault with default headers values
func NewCredentialsUpdateCredentialDefault(code int) *CredentialsUpdateCredentialDefault {
	return &CredentialsUpdateCredentialDefault{
		_statusCode: code,
	}
}

/*CredentialsUpdateCredentialDefault handles this case with default header values.

An unexpected error response
*/
type CredentialsUpdateCredentialDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the credentials update credential default response
func (o *CredentialsUpdateCredentialDefault) Code() int {
	return o._statusCode
}

func (o *CredentialsUpdateCredentialDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/credentials/{id}][%d] credentials_UpdateCredential default  %+v", o._statusCode, o.Payload)
}

func (o *CredentialsUpdateCredentialDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *CredentialsUpdateCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
