// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// KeyAttributes key attributes
// swagger:model KeyAttributes
type KeyAttributes struct {

	// certificate signing request
	CertificateSigningRequest string `json:"certificate_signing_request,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// private key
	PrivateKey string `json:"private_key,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this key attributes
func (m *KeyAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeyAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyAttributes) UnmarshalBinary(b []byte) error {
	var res KeyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
