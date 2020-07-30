// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// LogsTailStackRevisionLogsForFunctionReader is a Reader for the LogsTailStackRevisionLogsForFunction structure.
type LogsTailStackRevisionLogsForFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogsTailStackRevisionLogsForFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogsTailStackRevisionLogsForFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLogsTailStackRevisionLogsForFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLogsTailStackRevisionLogsForFunctionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLogsTailStackRevisionLogsForFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLogsTailStackRevisionLogsForFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogsTailStackRevisionLogsForFunctionOK creates a LogsTailStackRevisionLogsForFunctionOK with default headers values
func NewLogsTailStackRevisionLogsForFunctionOK() *LogsTailStackRevisionLogsForFunctionOK {
	return &LogsTailStackRevisionLogsForFunctionOK{}
}

/*LogsTailStackRevisionLogsForFunctionOK handles this case with default header values.

A successful response.(streaming responses)
*/
type LogsTailStackRevisionLogsForFunctionOK struct {
	Payload *LogsTailStackRevisionLogsForFunctionOKBody
}

func (o *LogsTailStackRevisionLogsForFunctionOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{stack_id}/revisions/{sha}/functions/{fn_name}/logs/tail][%d] logsTailStackRevisionLogsForFunctionOK  %+v", 200, o.Payload)
}

func (o *LogsTailStackRevisionLogsForFunctionOK) GetPayload() *LogsTailStackRevisionLogsForFunctionOKBody {
	return o.Payload
}

func (o *LogsTailStackRevisionLogsForFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LogsTailStackRevisionLogsForFunctionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogsTailStackRevisionLogsForFunctionBadRequest creates a LogsTailStackRevisionLogsForFunctionBadRequest with default headers values
func NewLogsTailStackRevisionLogsForFunctionBadRequest() *LogsTailStackRevisionLogsForFunctionBadRequest {
	return &LogsTailStackRevisionLogsForFunctionBadRequest{}
}

/*LogsTailStackRevisionLogsForFunctionBadRequest handles this case with default header values.

Validation error
*/
type LogsTailStackRevisionLogsForFunctionBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *LogsTailStackRevisionLogsForFunctionBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{stack_id}/revisions/{sha}/functions/{fn_name}/logs/tail][%d] logsTailStackRevisionLogsForFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *LogsTailStackRevisionLogsForFunctionBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *LogsTailStackRevisionLogsForFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogsTailStackRevisionLogsForFunctionForbidden creates a LogsTailStackRevisionLogsForFunctionForbidden with default headers values
func NewLogsTailStackRevisionLogsForFunctionForbidden() *LogsTailStackRevisionLogsForFunctionForbidden {
	return &LogsTailStackRevisionLogsForFunctionForbidden{}
}

/*LogsTailStackRevisionLogsForFunctionForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type LogsTailStackRevisionLogsForFunctionForbidden struct {
	Payload *models.CommonError
}

func (o *LogsTailStackRevisionLogsForFunctionForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{stack_id}/revisions/{sha}/functions/{fn_name}/logs/tail][%d] logsTailStackRevisionLogsForFunctionForbidden  %+v", 403, o.Payload)
}

func (o *LogsTailStackRevisionLogsForFunctionForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *LogsTailStackRevisionLogsForFunctionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogsTailStackRevisionLogsForFunctionNotFound creates a LogsTailStackRevisionLogsForFunctionNotFound with default headers values
func NewLogsTailStackRevisionLogsForFunctionNotFound() *LogsTailStackRevisionLogsForFunctionNotFound {
	return &LogsTailStackRevisionLogsForFunctionNotFound{}
}

/*LogsTailStackRevisionLogsForFunctionNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type LogsTailStackRevisionLogsForFunctionNotFound struct {
	Payload *models.CommonError
}

func (o *LogsTailStackRevisionLogsForFunctionNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{stack_id}/revisions/{sha}/functions/{fn_name}/logs/tail][%d] logsTailStackRevisionLogsForFunctionNotFound  %+v", 404, o.Payload)
}

func (o *LogsTailStackRevisionLogsForFunctionNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *LogsTailStackRevisionLogsForFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogsTailStackRevisionLogsForFunctionDefault creates a LogsTailStackRevisionLogsForFunctionDefault with default headers values
func NewLogsTailStackRevisionLogsForFunctionDefault(code int) *LogsTailStackRevisionLogsForFunctionDefault {
	return &LogsTailStackRevisionLogsForFunctionDefault{
		_statusCode: code,
	}
}

/*LogsTailStackRevisionLogsForFunctionDefault handles this case with default header values.

An unexpected error response
*/
type LogsTailStackRevisionLogsForFunctionDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the logs tail stack revision logs for function default response
func (o *LogsTailStackRevisionLogsForFunctionDefault) Code() int {
	return o._statusCode
}

func (o *LogsTailStackRevisionLogsForFunctionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{stack_id}/revisions/{sha}/functions/{fn_name}/logs/tail][%d] logs_TailStackRevisionLogsForFunction default  %+v", o._statusCode, o.Payload)
}

func (o *LogsTailStackRevisionLogsForFunctionDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *LogsTailStackRevisionLogsForFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LogsTailStackRevisionLogsForFunctionOKBody Stream result of storageLogEntry
swagger:model LogsTailStackRevisionLogsForFunctionOKBody
*/
type LogsTailStackRevisionLogsForFunctionOKBody struct {

	// error
	Error *models.RuntimeStreamError `json:"error,omitempty"`

	// result
	Result *models.StorageLogEntry `json:"result,omitempty"`
}

// Validate validates this logs tail stack revision logs for function o k body
func (o *LogsTailStackRevisionLogsForFunctionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LogsTailStackRevisionLogsForFunctionOKBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logsTailStackRevisionLogsForFunctionOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *LogsTailStackRevisionLogsForFunctionOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logsTailStackRevisionLogsForFunctionOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LogsTailStackRevisionLogsForFunctionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LogsTailStackRevisionLogsForFunctionOKBody) UnmarshalBinary(b []byte) error {
	var res LogsTailStackRevisionLogsForFunctionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
