// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FunctionRunInfoState function run info state
// swagger:model FunctionRunInfoState
type FunctionRunInfoState string

const (

	// FunctionRunInfoStateUNKNOWN captures enum value "UNKNOWN"
	FunctionRunInfoStateUNKNOWN FunctionRunInfoState = "UNKNOWN"

	// FunctionRunInfoStateSTARTING captures enum value "STARTING"
	FunctionRunInfoStateSTARTING FunctionRunInfoState = "STARTING"

	// FunctionRunInfoStateRUNNING captures enum value "RUNNING"
	FunctionRunInfoStateRUNNING FunctionRunInfoState = "RUNNING"

	// FunctionRunInfoStateFAILED captures enum value "FAILED"
	FunctionRunInfoStateFAILED FunctionRunInfoState = "FAILED"

	// FunctionRunInfoStateSUCCEEDED captures enum value "SUCCEEDED"
	FunctionRunInfoStateSUCCEEDED FunctionRunInfoState = "SUCCEEDED"
)

// for schema
var functionRunInfoStateEnum []interface{}

func init() {
	var res []FunctionRunInfoState
	if err := json.Unmarshal([]byte(`["UNKNOWN","STARTING","RUNNING","FAILED","SUCCEEDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		functionRunInfoStateEnum = append(functionRunInfoStateEnum, v)
	}
}

func (m FunctionRunInfoState) validateFunctionRunInfoStateEnum(path, location string, value FunctionRunInfoState) error {
	if err := validate.Enum(path, location, value, functionRunInfoStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this function run info state
func (m FunctionRunInfoState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFunctionRunInfoStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
