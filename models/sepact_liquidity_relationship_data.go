// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SepactLiquidityRelationshipData sepact liquidity relationship data
//
// swagger:model SepactLiquidityRelationshipData
type SepactLiquidityRelationshipData struct {

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// type
	// Required: true
	// Enum: [sepact_liquidity_associations]
	Type string `json:"type"`
}

// Validate validates this sepact liquidity relationship data
func (m *SepactLiquidityRelationshipData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepactLiquidityRelationshipData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var sepactLiquidityRelationshipDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sepact_liquidity_associations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sepactLiquidityRelationshipDataTypeTypePropEnum = append(sepactLiquidityRelationshipDataTypeTypePropEnum, v)
	}
}

const (

	// SepactLiquidityRelationshipDataTypeSepactLiquidityAssociations captures enum value "sepact_liquidity_associations"
	SepactLiquidityRelationshipDataTypeSepactLiquidityAssociations string = "sepact_liquidity_associations"
)

// prop value enum
func (m *SepactLiquidityRelationshipData) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sepactLiquidityRelationshipDataTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SepactLiquidityRelationshipData) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepactLiquidityRelationshipData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepactLiquidityRelationshipData) UnmarshalBinary(b []byte) error {
	var res SepactLiquidityRelationshipData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}