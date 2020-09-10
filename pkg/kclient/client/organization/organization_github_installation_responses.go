// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// OrganizationGithubInstallationReader is a Reader for the OrganizationGithubInstallation structure.
type OrganizationGithubInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationGithubInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationGithubInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationGithubInstallationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationGithubInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationGithubInstallationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrganizationGithubInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationGithubInstallationOK creates a OrganizationGithubInstallationOK with default headers values
func NewOrganizationGithubInstallationOK() *OrganizationGithubInstallationOK {
	return &OrganizationGithubInstallationOK{}
}

/*OrganizationGithubInstallationOK handles this case with default header values.

A successful response.
*/
type OrganizationGithubInstallationOK struct {
	Payload *models.AccountGithubInstallationReply
}

func (o *OrganizationGithubInstallationOK) Error() string {
	return fmt.Sprintf("[GET /v1/github/installation][%d] organizationGithubInstallationOK  %+v", 200, o.Payload)
}

func (o *OrganizationGithubInstallationOK) GetPayload() *models.AccountGithubInstallationReply {
	return o.Payload
}

func (o *OrganizationGithubInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountGithubInstallationReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGithubInstallationBadRequest creates a OrganizationGithubInstallationBadRequest with default headers values
func NewOrganizationGithubInstallationBadRequest() *OrganizationGithubInstallationBadRequest {
	return &OrganizationGithubInstallationBadRequest{}
}

/*OrganizationGithubInstallationBadRequest handles this case with default header values.

Validation error
*/
type OrganizationGithubInstallationBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *OrganizationGithubInstallationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/github/installation][%d] organizationGithubInstallationBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationGithubInstallationBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *OrganizationGithubInstallationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGithubInstallationForbidden creates a OrganizationGithubInstallationForbidden with default headers values
func NewOrganizationGithubInstallationForbidden() *OrganizationGithubInstallationForbidden {
	return &OrganizationGithubInstallationForbidden{}
}

/*OrganizationGithubInstallationForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type OrganizationGithubInstallationForbidden struct {
	Payload *models.CommonError
}

func (o *OrganizationGithubInstallationForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/github/installation][%d] organizationGithubInstallationForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationGithubInstallationForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationGithubInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGithubInstallationNotFound creates a OrganizationGithubInstallationNotFound with default headers values
func NewOrganizationGithubInstallationNotFound() *OrganizationGithubInstallationNotFound {
	return &OrganizationGithubInstallationNotFound{}
}

/*OrganizationGithubInstallationNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type OrganizationGithubInstallationNotFound struct {
	Payload *models.CommonError
}

func (o *OrganizationGithubInstallationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/github/installation][%d] organizationGithubInstallationNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationGithubInstallationNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationGithubInstallationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGithubInstallationDefault creates a OrganizationGithubInstallationDefault with default headers values
func NewOrganizationGithubInstallationDefault(code int) *OrganizationGithubInstallationDefault {
	return &OrganizationGithubInstallationDefault{
		_statusCode: code,
	}
}

/*OrganizationGithubInstallationDefault handles this case with default header values.

An unexpected error response
*/
type OrganizationGithubInstallationDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the organization github installation default response
func (o *OrganizationGithubInstallationDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationGithubInstallationDefault) Error() string {
	return fmt.Sprintf("[GET /v1/github/installation][%d] organization_GithubInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationGithubInstallationDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *OrganizationGithubInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
