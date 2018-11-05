// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SettlementCycleAttributes settlement cycle attributes
// swagger:model settlementCycleAttributes
type SettlementCycleAttributes struct {

	// gateway
	// Pattern: ^[A-Za-z_\-]*$
	Gateway string `json:"gateway,omitempty"`

	// settlement cycle number
	// Minimum: 1
	SettlementCycleNumber int64 `json:"settlement_cycle_number,omitempty"`

	// settlement cycle type
	// Pattern: ^[A-Za-z_\-]*$
	SettlementCycleType string `json:"settlement_cycle_type,omitempty"`
}

// Validate validates this settlement cycle attributes
func (m *SettlementCycleAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSettlementCycleNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSettlementCycleType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettlementCycleAttributes) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if err := validate.Pattern("gateway", "body", string(m.Gateway), `^[A-Za-z_\-]*$`); err != nil {
		return err
	}

	return nil
}

func (m *SettlementCycleAttributes) validateSettlementCycleNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycleNumber) { // not required
		return nil
	}

	if err := validate.MinimumInt("settlement_cycle_number", "body", int64(m.SettlementCycleNumber), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *SettlementCycleAttributes) validateSettlementCycleType(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycleType) { // not required
		return nil
	}

	if err := validate.Pattern("settlement_cycle_type", "body", string(m.SettlementCycleType), `^[A-Za-z_\-]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettlementCycleAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettlementCycleAttributes) UnmarshalBinary(b []byte) error {
	var res SettlementCycleAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
