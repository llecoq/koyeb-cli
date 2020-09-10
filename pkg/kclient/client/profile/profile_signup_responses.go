// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// ProfileSignupReader is a Reader for the ProfileSignup structure.
type ProfileSignupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfileSignupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProfileSignupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProfileSignupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProfileSignupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProfileSignupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProfileSignupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProfileSignupOK creates a ProfileSignupOK with default headers values
func NewProfileSignupOK() *ProfileSignupOK {
	return &ProfileSignupOK{}
}

/*ProfileSignupOK handles this case with default header values.

A successful response.
*/
type ProfileSignupOK struct {
	Payload *models.AccountLoginReply
}

func (o *ProfileSignupOK) Error() string {
	return fmt.Sprintf("[POST /v1/account/signup][%d] profileSignupOK  %+v", 200, o.Payload)
}

func (o *ProfileSignupOK) GetPayload() *models.AccountLoginReply {
	return o.Payload
}

func (o *ProfileSignupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountLoginReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileSignupBadRequest creates a ProfileSignupBadRequest with default headers values
func NewProfileSignupBadRequest() *ProfileSignupBadRequest {
	return &ProfileSignupBadRequest{}
}

/*ProfileSignupBadRequest handles this case with default header values.

Validation error
*/
type ProfileSignupBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ProfileSignupBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/account/signup][%d] profileSignupBadRequest  %+v", 400, o.Payload)
}

func (o *ProfileSignupBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ProfileSignupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileSignupForbidden creates a ProfileSignupForbidden with default headers values
func NewProfileSignupForbidden() *ProfileSignupForbidden {
	return &ProfileSignupForbidden{}
}

/*ProfileSignupForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ProfileSignupForbidden struct {
	Payload *models.CommonError
}

func (o *ProfileSignupForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/account/signup][%d] profileSignupForbidden  %+v", 403, o.Payload)
}

func (o *ProfileSignupForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ProfileSignupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileSignupNotFound creates a ProfileSignupNotFound with default headers values
func NewProfileSignupNotFound() *ProfileSignupNotFound {
	return &ProfileSignupNotFound{}
}

/*ProfileSignupNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ProfileSignupNotFound struct {
	Payload *models.CommonError
}

func (o *ProfileSignupNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/account/signup][%d] profileSignupNotFound  %+v", 404, o.Payload)
}

func (o *ProfileSignupNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ProfileSignupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileSignupDefault creates a ProfileSignupDefault with default headers values
func NewProfileSignupDefault(code int) *ProfileSignupDefault {
	return &ProfileSignupDefault{
		_statusCode: code,
	}
}

/*ProfileSignupDefault handles this case with default header values.

An unexpected error response
*/
type ProfileSignupDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the profile signup default response
func (o *ProfileSignupDefault) Code() int {
	return o._statusCode
}

func (o *ProfileSignupDefault) Error() string {
	return fmt.Sprintf("[POST /v1/account/signup][%d] profile_Signup default  %+v", o._statusCode, o.Payload)
}

func (o *ProfileSignupDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *ProfileSignupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
