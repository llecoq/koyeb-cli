// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// S3CredentialsNewS3CredentialReader is a Reader for the S3CredentialsNewS3Credential structure.
type S3CredentialsNewS3CredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3CredentialsNewS3CredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3CredentialsNewS3CredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS3CredentialsNewS3CredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS3CredentialsNewS3CredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS3CredentialsNewS3CredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewS3CredentialsNewS3CredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3CredentialsNewS3CredentialOK creates a S3CredentialsNewS3CredentialOK with default headers values
func NewS3CredentialsNewS3CredentialOK() *S3CredentialsNewS3CredentialOK {
	return &S3CredentialsNewS3CredentialOK{}
}

/*S3CredentialsNewS3CredentialOK handles this case with default header values.

A successful response.
*/
type S3CredentialsNewS3CredentialOK struct {
	Payload *models.AccountS3CredentialReply
}

func (o *S3CredentialsNewS3CredentialOK) Error() string {
	return fmt.Sprintf("[POST /v1/s3_credentials][%d] s3CredentialsNewS3CredentialOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsNewS3CredentialOK) GetPayload() *models.AccountS3CredentialReply {
	return o.Payload
}

func (o *S3CredentialsNewS3CredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountS3CredentialReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsNewS3CredentialBadRequest creates a S3CredentialsNewS3CredentialBadRequest with default headers values
func NewS3CredentialsNewS3CredentialBadRequest() *S3CredentialsNewS3CredentialBadRequest {
	return &S3CredentialsNewS3CredentialBadRequest{}
}

/*S3CredentialsNewS3CredentialBadRequest handles this case with default header values.

Validation error
*/
type S3CredentialsNewS3CredentialBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *S3CredentialsNewS3CredentialBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/s3_credentials][%d] s3CredentialsNewS3CredentialBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsNewS3CredentialBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *S3CredentialsNewS3CredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsNewS3CredentialForbidden creates a S3CredentialsNewS3CredentialForbidden with default headers values
func NewS3CredentialsNewS3CredentialForbidden() *S3CredentialsNewS3CredentialForbidden {
	return &S3CredentialsNewS3CredentialForbidden{}
}

/*S3CredentialsNewS3CredentialForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type S3CredentialsNewS3CredentialForbidden struct {
	Payload *models.CommonError
}

func (o *S3CredentialsNewS3CredentialForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/s3_credentials][%d] s3CredentialsNewS3CredentialForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsNewS3CredentialForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *S3CredentialsNewS3CredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsNewS3CredentialNotFound creates a S3CredentialsNewS3CredentialNotFound with default headers values
func NewS3CredentialsNewS3CredentialNotFound() *S3CredentialsNewS3CredentialNotFound {
	return &S3CredentialsNewS3CredentialNotFound{}
}

/*S3CredentialsNewS3CredentialNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type S3CredentialsNewS3CredentialNotFound struct {
	Payload *models.CommonError
}

func (o *S3CredentialsNewS3CredentialNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/s3_credentials][%d] s3CredentialsNewS3CredentialNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsNewS3CredentialNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *S3CredentialsNewS3CredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsNewS3CredentialDefault creates a S3CredentialsNewS3CredentialDefault with default headers values
func NewS3CredentialsNewS3CredentialDefault(code int) *S3CredentialsNewS3CredentialDefault {
	return &S3CredentialsNewS3CredentialDefault{
		_statusCode: code,
	}
}

/*S3CredentialsNewS3CredentialDefault handles this case with default header values.

An unexpected error response
*/
type S3CredentialsNewS3CredentialDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the s3 credentials new s3 credential default response
func (o *S3CredentialsNewS3CredentialDefault) Code() int {
	return o._statusCode
}

func (o *S3CredentialsNewS3CredentialDefault) Error() string {
	return fmt.Sprintf("[POST /v1/s3_credentials][%d] s3Credentials_NewS3Credential default  %+v", o._statusCode, o.Payload)
}

func (o *S3CredentialsNewS3CredentialDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *S3CredentialsNewS3CredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
