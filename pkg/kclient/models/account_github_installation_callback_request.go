// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountGithubInstallationCallbackRequest account github installation callback request
// swagger:model accountGithubInstallationCallbackRequest
type AccountGithubInstallationCallbackRequest struct {

	// installation id
	InstallationID string `json:"installation_id,omitempty"`

	// setup action
	SetupAction string `json:"setup_action,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this account github installation callback request
func (m *AccountGithubInstallationCallbackRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountGithubInstallationCallbackRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountGithubInstallationCallbackRequest) UnmarshalBinary(b []byte) error {
	var res AccountGithubInstallationCallbackRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
