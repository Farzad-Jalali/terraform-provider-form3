// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReconciliationAssociationAttributes reconciliation association attributes
//
// swagger:model ReconciliationAssociationAttributes
type ReconciliationAssociationAttributes struct {

	// bank ids
	// Required: true
	// Min Items: 1
	// Unique: true
	BankIds []string `json:"bank_ids"`

	// name
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// scheme type
	// Required: true
	// Enum: [SEPAINSTANT FPS]
	SchemeType string `json:"scheme_type"`
}

// Validate validates this reconciliation association attributes
func (m *ReconciliationAssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReconciliationAssociationAttributes) validateBankIds(formats strfmt.Registry) error {

	if err := validate.Required("bank_ids", "body", m.BankIds); err != nil {
		return err
	}

	iBankIdsSize := int64(len(m.BankIds))

	if err := validate.MinItems("bank_ids", "body", iBankIdsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("bank_ids", "body", m.BankIds); err != nil {
		return err
	}

	for i := 0; i < len(m.BankIds); i++ {

		if err := validate.MinLength("bank_ids"+"."+strconv.Itoa(i), "body", string(m.BankIds[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *ReconciliationAssociationAttributes) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

var reconciliationAssociationAttributesTypeSchemeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEPAINSTANT","FPS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reconciliationAssociationAttributesTypeSchemeTypePropEnum = append(reconciliationAssociationAttributesTypeSchemeTypePropEnum, v)
	}
}

const (

	// ReconciliationAssociationAttributesSchemeTypeSEPAINSTANT captures enum value "SEPAINSTANT"
	ReconciliationAssociationAttributesSchemeTypeSEPAINSTANT string = "SEPAINSTANT"

	// ReconciliationAssociationAttributesSchemeTypeFPS captures enum value "FPS"
	ReconciliationAssociationAttributesSchemeTypeFPS string = "FPS"
)

// prop value enum
func (m *ReconciliationAssociationAttributes) validateSchemeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, reconciliationAssociationAttributesTypeSchemeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReconciliationAssociationAttributes) validateSchemeType(formats strfmt.Registry) error {

	if err := validate.RequiredString("scheme_type", "body", string(m.SchemeType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateSchemeTypeEnum("scheme_type", "body", m.SchemeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReconciliationAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReconciliationAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res ReconciliationAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
