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

// MandateAttributesBeneficiaryParty mandate attributes beneficiary party
// swagger:model MandateAttributesBeneficiaryParty
type MandateAttributesBeneficiaryParty struct {

	// account name
	// Pattern: ^[A-Za-z0-9 \/\-?:\(\)\.,’\+\#\=\!\"%&\*\<\>;\{@\r\n]*$
	AccountName string `json:"account_name,omitempty"`

	// account number
	// Pattern: ^[A-Z0-9]{6,34}$
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account type
	AccountType int64 `json:"account_type,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// address
	Address []string `json:"address"`

	// country
	Country string `json:"country,omitempty"`
}

// Validate validates this mandate attributes beneficiary party
func (m *MandateAttributesBeneficiaryParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAttributesBeneficiaryParty) validateAccountName(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountName) { // not required
		return nil
	}

	if err := validate.Pattern("account_name", "body", string(m.AccountName), `^[A-Za-z0-9 \/\-?:\(\)\.,’\+\#\=\!\"%&\*\<\>;\{@\r\n]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateAttributesBeneficiaryParty) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{6,34}$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateAttributesBeneficiaryParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_number_code")
		}
		return err
	}

	return nil
}

func (m *MandateAttributesBeneficiaryParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_with")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAttributesBeneficiaryParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAttributesBeneficiaryParty) UnmarshalBinary(b []byte) error {
	var res MandateAttributesBeneficiaryParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
