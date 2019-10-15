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

// BacsAssociationAttributes bacs association attributes
// swagger:model BacsAssociationAttributes
type BacsAssociationAttributes struct {

	// account number
	// Pattern: ^[0-9]{8}$
	AccountNumber string `json:"account_number,omitempty"`

	// account type
	AccountType *int64 `json:"account_type,omitempty"`

	// bank code
	// Pattern: ^[0-9A-Z]{4}$
	BankCode string `json:"bank_code,omitempty"`

	// centre number
	// Pattern: ^[0-9A-Z]{2}$
	CentreNumber string `json:"centre_number,omitempty"`

	// service user number
	// Pattern: ^[0-9A-Z]{6}$
	ServiceUserNumber string `json:"service_user_number,omitempty"`

	// sorting code
	// Pattern: ^[0-9]{6}$
	SortingCode string `json:"sorting_code,omitempty"`

	// use test file submission
	UseTestFileSubmission bool `json:"use_test_file_submission,omitempty"`
}

// Validate validates this bacs association attributes
func (m *BacsAssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCentreNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceUserNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortingCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BacsAssociationAttributes) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("account_number", "body", string(m.AccountNumber), `^[0-9]{8}$`); err != nil {
		return err
	}

	return nil
}

func (m *BacsAssociationAttributes) validateBankCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankCode) { // not required
		return nil
	}

	if err := validate.Pattern("bank_code", "body", string(m.BankCode), `^[0-9A-Z]{4}$`); err != nil {
		return err
	}

	return nil
}

func (m *BacsAssociationAttributes) validateCentreNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.CentreNumber) { // not required
		return nil
	}

	if err := validate.Pattern("centre_number", "body", string(m.CentreNumber), `^[0-9A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *BacsAssociationAttributes) validateServiceUserNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceUserNumber) { // not required
		return nil
	}

	if err := validate.Pattern("service_user_number", "body", string(m.ServiceUserNumber), `^[0-9A-Z]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *BacsAssociationAttributes) validateSortingCode(formats strfmt.Registry) error {

	if swag.IsZero(m.SortingCode) { // not required
		return nil
	}

	if err := validate.Pattern("sorting_code", "body", string(m.SortingCode), `^[0-9]{6}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BacsAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BacsAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res BacsAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
