// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GocardlessAssociationAttributes gocardless association attributes
// swagger:model GocardlessAssociationAttributes
type GocardlessAssociationAttributes struct {

	// schemes
	Schemes []string `json:"schemes"`
}

// Validate validates this gocardless association attributes
func (m *GocardlessAssociationAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GocardlessAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GocardlessAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res GocardlessAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
