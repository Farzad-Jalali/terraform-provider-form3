// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CredentialSecret credential secret
//
// swagger:model CredentialSecret
type CredentialSecret struct {

	// client id
	// Format: uuid
	ClientID strfmt.UUID `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`
}

// Validate validates this credential secret
func (m *CredentialSecret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialSecret) validateClientID(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientID) { // not required
		return nil
	}

	if err := validate.FormatOf("client_id", "body", "uuid", m.ClientID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialSecret) UnmarshalBinary(b []byte) error {
	var res CredentialSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
