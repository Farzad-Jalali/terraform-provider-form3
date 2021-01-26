// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SepaDDAssociationAttributes sepa d d association attributes
//
// swagger:model SepaDDAssociationAttributes
type SepaDDAssociationAttributes struct {

	// allow submissions
	AllowSubmissions bool `json:"allowSubmissions,omitempty"`

	// bic
	Bic string `json:"bic,omitempty"`

	// business user
	BusinessUser string `json:"businessUser,omitempty"`

	// local instrument
	LocalInstrument string `json:"localInstrument,omitempty"`

	// receiver business user
	ReceiverBusinessUser string `json:"receiverBusinessUser,omitempty"`
}

// Validate validates this sepa d d association attributes
func (m *SepaDDAssociationAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SepaDDAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaDDAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res SepaDDAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
