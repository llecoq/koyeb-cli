// Code generated by go-swagger; DO NOT EDIT.

package functions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// FunctionsInvokeFunctionReader is a Reader for the FunctionsInvokeFunction structure.
type FunctionsInvokeFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FunctionsInvokeFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFunctionsInvokeFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFunctionsInvokeFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFunctionsInvokeFunctionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFunctionsInvokeFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewFunctionsInvokeFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFunctionsInvokeFunctionOK creates a FunctionsInvokeFunctionOK with default headers values
func NewFunctionsInvokeFunctionOK() *FunctionsInvokeFunctionOK {
	return &FunctionsInvokeFunctionOK{}
}

/*FunctionsInvokeFunctionOK handles this case with default header values.

A successful response.
*/
type FunctionsInvokeFunctionOK struct {
	Payload *models.StorageInvokeFunctionReply
}

func (o *FunctionsInvokeFunctionOK) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/functions/{function}/invoke][%d] functionsInvokeFunctionOK  %+v", 200, o.Payload)
}

func (o *FunctionsInvokeFunctionOK) GetPayload() *models.StorageInvokeFunctionReply {
	return o.Payload
}

func (o *FunctionsInvokeFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageInvokeFunctionReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsInvokeFunctionBadRequest creates a FunctionsInvokeFunctionBadRequest with default headers values
func NewFunctionsInvokeFunctionBadRequest() *FunctionsInvokeFunctionBadRequest {
	return &FunctionsInvokeFunctionBadRequest{}
}

/*FunctionsInvokeFunctionBadRequest handles this case with default header values.

Validation error
*/
type FunctionsInvokeFunctionBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *FunctionsInvokeFunctionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/functions/{function}/invoke][%d] functionsInvokeFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *FunctionsInvokeFunctionBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *FunctionsInvokeFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsInvokeFunctionForbidden creates a FunctionsInvokeFunctionForbidden with default headers values
func NewFunctionsInvokeFunctionForbidden() *FunctionsInvokeFunctionForbidden {
	return &FunctionsInvokeFunctionForbidden{}
}

/*FunctionsInvokeFunctionForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type FunctionsInvokeFunctionForbidden struct {
	Payload *models.CommonError
}

func (o *FunctionsInvokeFunctionForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/functions/{function}/invoke][%d] functionsInvokeFunctionForbidden  %+v", 403, o.Payload)
}

func (o *FunctionsInvokeFunctionForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *FunctionsInvokeFunctionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsInvokeFunctionNotFound creates a FunctionsInvokeFunctionNotFound with default headers values
func NewFunctionsInvokeFunctionNotFound() *FunctionsInvokeFunctionNotFound {
	return &FunctionsInvokeFunctionNotFound{}
}

/*FunctionsInvokeFunctionNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type FunctionsInvokeFunctionNotFound struct {
	Payload *models.CommonError
}

func (o *FunctionsInvokeFunctionNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/functions/{function}/invoke][%d] functionsInvokeFunctionNotFound  %+v", 404, o.Payload)
}

func (o *FunctionsInvokeFunctionNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *FunctionsInvokeFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsInvokeFunctionDefault creates a FunctionsInvokeFunctionDefault with default headers values
func NewFunctionsInvokeFunctionDefault(code int) *FunctionsInvokeFunctionDefault {
	return &FunctionsInvokeFunctionDefault{
		_statusCode: code,
	}
}

/*FunctionsInvokeFunctionDefault handles this case with default header values.

An unexpected error response
*/
type FunctionsInvokeFunctionDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the functions invoke function default response
func (o *FunctionsInvokeFunctionDefault) Code() int {
	return o._statusCode
}

func (o *FunctionsInvokeFunctionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/functions/{function}/invoke][%d] Functions_InvokeFunction default  %+v", o._statusCode, o.Payload)
}

func (o *FunctionsInvokeFunctionDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *FunctionsInvokeFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
