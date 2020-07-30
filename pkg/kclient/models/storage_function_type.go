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

// StorageFunctionType storage function type
// swagger:model storageFunctionType
type StorageFunctionType string

const (

	// StorageFunctionTypeDOCKER captures enum value "DOCKER"
	StorageFunctionTypeDOCKER StorageFunctionType = "DOCKER"

	// StorageFunctionTypeCATALOG captures enum value "CATALOG"
	StorageFunctionTypeCATALOG StorageFunctionType = "CATALOG"

	// StorageFunctionTypeCODE captures enum value "CODE"
	StorageFunctionTypeCODE StorageFunctionType = "CODE"
)

// for schema
var storageFunctionTypeEnum []interface{}

func init() {
	var res []StorageFunctionType
	if err := json.Unmarshal([]byte(`["DOCKER","CATALOG","CODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageFunctionTypeEnum = append(storageFunctionTypeEnum, v)
	}
}

func (m StorageFunctionType) validateStorageFunctionTypeEnum(path, location string, value StorageFunctionType) error {
	if err := validate.Enum(path, location, value, storageFunctionTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this storage function type
func (m StorageFunctionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStorageFunctionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
