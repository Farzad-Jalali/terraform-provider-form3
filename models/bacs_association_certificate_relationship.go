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

// BacsAssociationCertificateRelationship bacs association certificate relationship
//
// swagger:model BacsAssociationCertificateRelationship
type BacsAssociationCertificateRelationship struct {

	// data
	Data *BacsAssociationCertificateRelationshipData `json:"data,omitempty"`
}

// Validate validates this bacs association certificate relationship
func (m *BacsAssociationCertificateRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BacsAssociationCertificateRelationship) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BacsAssociationCertificateRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BacsAssociationCertificateRelationship) UnmarshalBinary(b []byte) error {
	var res BacsAssociationCertificateRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BacsAssociationCertificateRelationshipData bacs association certificate relationship data
//
// swagger:model BacsAssociationCertificateRelationshipData
type BacsAssociationCertificateRelationshipData struct {

	// certificate id
	// Format: uuid
	CertificateID strfmt.UUID `json:"certificate_id,omitempty"`

	// key id
	// Format: uuid
	KeyID strfmt.UUID `json:"key_id,omitempty"`

	// tsu number
	// Pattern: ^[0-9A-Z]{6}$
	TsuNumber string `json:"tsu_number,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this bacs association certificate relationship data
func (m *BacsAssociationCertificateRelationshipData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTsuNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BacsAssociationCertificateRelationshipData) validateCertificateID(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificateID) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"certificate_id", "body", "uuid", m.CertificateID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BacsAssociationCertificateRelationshipData) validateKeyID(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyID) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"key_id", "body", "uuid", m.KeyID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BacsAssociationCertificateRelationshipData) validateTsuNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.TsuNumber) { // not required
		return nil
	}

	if err := validate.Pattern("data"+"."+"tsu_number", "body", string(m.TsuNumber), `^[0-9A-Z]{6}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BacsAssociationCertificateRelationshipData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BacsAssociationCertificateRelationshipData) UnmarshalBinary(b []byte) error {
	var res BacsAssociationCertificateRelationshipData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
