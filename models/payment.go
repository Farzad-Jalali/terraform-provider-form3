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

// Payment payment
//
// swagger:model Payment
type Payment struct {

	// attributes
	// Required: true
	Attributes *PaymentAttributes `json:"attributes"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *PaymentRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this payment
func (m *Payment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Payment) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *Payment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *Payment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Payment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Payment) UnmarshalBinary(b []byte) error {
	var res Payment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributes payment attributes
//
// swagger:model PaymentAttributes
type PaymentAttributes struct {

	// Amount of money moved between the instructing agent and instructed agent
	// Pattern: ^[0-9.]{0,20}$
	Amount string `json:"amount,omitempty"`

	// batch booking indicator
	BatchBookingIndicator string `json:"batch_booking_indicator,omitempty"`

	// batch id
	BatchID string `json:"batch_id,omitempty"`

	// batch type
	BatchType string `json:"batch_type,omitempty"`

	// beneficiary party
	BeneficiaryParty *PaymentAttributesBeneficiaryParty `json:"beneficiary_party,omitempty"`

	// Category purpose in proprietary form. Specifies the high level purpose of the instruction. Cannot be used at the same time as `category_purpose_coded`.
	CategoryPurpose string `json:"category_purpose,omitempty"`

	// Category purpose in a coded form. Specifies the high level purpose of the instruction. Cannot be used at the same time as `category_purpose`.
	CategoryPurposeCoded string `json:"category_purpose_coded,omitempty"`

	// charges information
	ChargesInformation *ChargesInformation `json:"charges_information,omitempty"`

	// Unique identifier for organisations collecting payments
	ClearingID string `json:"clearing_id,omitempty"`

	// Currency of the transaction amount. Currency code as defined in [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
	Currency string `json:"currency,omitempty"`

	// debtor party
	DebtorParty *PaymentAttributesDebtorParty `json:"debtor_party,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	EndToEndReference string `json:"end_to_end_reference,omitempty"`

	// file number
	FileNumber string `json:"file_number,omitempty"`

	// fx
	Fx *PaymentAttributesFx `json:"fx,omitempty"`

	// Unique identification, as assigned by the initiating party to unambigiously identify the transaction. This identification is an point-to-point reference and is passed on, unchanged, throughout the entire chain. Cannot includ leading, trailing or internal spaces.
	InstructionID string `json:"instruction_id,omitempty"`

	// intermediary bank
	IntermediaryBank *IntermediaryBankAccountHoldingEntity `json:"intermediary_bank,omitempty"`

	// Numeric reference field, see scheme specific descriptions for usage
	NumericReference string `json:"numeric_reference,omitempty"`

	// Timestamp of when the payment instruction meets the set processing conditions. Format: YYYY-MM-DDThh:mm:ss:mmm+hh:mm
	// Format: date-time
	PaymentAcceptanceDatetime *strfmt.DateTime `json:"payment_acceptance_datetime,omitempty"`

	// Purpose of the payment in a proprietary form
	PaymentPurpose string `json:"payment_purpose,omitempty"`

	// Purpose of the payment in a coded form
	PaymentPurposeCoded string `json:"payment_purpose_coded,omitempty"`

	// Clearing infrastructure through which the payment instruction is to be processed. Default for given organisation ID is used if left empty. Has to be a valid [scheme identifier](http://draft-api-docs.form3.tech/api.html#enumerations-schemes).
	PaymentScheme string `json:"payment_scheme,omitempty"`

	// payment type
	PaymentType string `json:"payment_type,omitempty"`

	// Date on which the payment is to be debited from the debtor account. Formatted according to ISO 8601 format: YYYY-MM-DD.
	// Format: date
	ProcessingDate strfmt.Date `json:"processing_date,omitempty"`

	// receivers correspondent
	ReceiversCorrespondent *ReceiversCorrespondentAccountHoldingEntity `json:"receivers_correspondent,omitempty"`

	// Payment reference for beneficiary use
	Reference string `json:"reference,omitempty"`

	// Regulatory reporting information
	RegulatoryReporting string `json:"regulatory_reporting,omitempty"`

	// reimbursement
	Reimbursement *ReimbursementAccountHoldingEntity `json:"reimbursement,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts receivable system provided by the debtor for the beneficiary.
	RemittanceInformation string `json:"remittance_information,omitempty"`

	// The scheme specific payment [sub type](http://api-docs.form3.tech/api.html#enumerations-scheme-specific-payment-sub-types)
	SchemePaymentSubType string `json:"scheme_payment_sub_type,omitempty"`

	// The [scheme-specific payment type](#enumerations-scheme-payment-types)
	SchemePaymentType string `json:"scheme_payment_type,omitempty"`

	// Date on which the payment is processed by the scheme. Only used if different from `processing_date`.
	// Format: date
	SchemeProcessingDate strfmt.Date `json:"scheme_processing_date,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	SchemeTransactionID string `json:"scheme_transaction_id,omitempty"`

	// senders correspondent
	SendersCorrespondent *SendersCorrespondentAccountHoldingEntity `json:"senders_correspondent,omitempty"`

	// structured reference
	StructuredReference *PaymentAttributesStructuredReference `json:"structured_reference,omitempty"`

	// swift
	Swift *PaymentAttributesSwift `json:"swift,omitempty"`

	// ultimate beneficiary
	UltimateBeneficiary *UltimateEntity `json:"ultimate_beneficiary,omitempty"`

	// ultimate debtor
	UltimateDebtor *UltimateEntity `json:"ultimate_debtor,omitempty"`

	// The scheme-specific unique transaction ID. Populated by the scheme.
	UniqueSchemeID string `json:"unique_scheme_id,omitempty"`
}

// Validate validates this payment attributes
func (m *PaymentAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargesInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntermediaryBank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAcceptanceDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiversCorrespondent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReimbursement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSendersCorrespondent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStructuredReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwift(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUltimateBeneficiary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUltimateDebtor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributes) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"amount", "body", string(m.Amount), `^[0-9.]{0,20}$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateBeneficiaryParty(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryParty) { // not required
		return nil
	}

	if m.BeneficiaryParty != nil {
		if err := m.BeneficiaryParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateChargesInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargesInformation) { // not required
		return nil
	}

	if m.ChargesInformation != nil {
		if err := m.ChargesInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "charges_information")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateDebtorParty(formats strfmt.Registry) error {

	if swag.IsZero(m.DebtorParty) { // not required
		return nil
	}

	if m.DebtorParty != nil {
		if err := m.DebtorParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "debtor_party")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateFx(formats strfmt.Registry) error {

	if swag.IsZero(m.Fx) { // not required
		return nil
	}

	if m.Fx != nil {
		if err := m.Fx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "fx")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateIntermediaryBank(formats strfmt.Registry) error {

	if swag.IsZero(m.IntermediaryBank) { // not required
		return nil
	}

	if m.IntermediaryBank != nil {
		if err := m.IntermediaryBank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "intermediary_bank")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validatePaymentAcceptanceDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAcceptanceDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"payment_acceptance_datetime", "body", "date-time", m.PaymentAcceptanceDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"processing_date", "body", "date", m.ProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateReceiversCorrespondent(formats strfmt.Registry) error {

	if swag.IsZero(m.ReceiversCorrespondent) { // not required
		return nil
	}

	if m.ReceiversCorrespondent != nil {
		if err := m.ReceiversCorrespondent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "receivers_correspondent")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateReimbursement(formats strfmt.Registry) error {

	if swag.IsZero(m.Reimbursement) { // not required
		return nil
	}

	if m.Reimbursement != nil {
		if err := m.Reimbursement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "reimbursement")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateSchemeProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"scheme_processing_date", "body", "date", m.SchemeProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateSendersCorrespondent(formats strfmt.Registry) error {

	if swag.IsZero(m.SendersCorrespondent) { // not required
		return nil
	}

	if m.SendersCorrespondent != nil {
		if err := m.SendersCorrespondent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "senders_correspondent")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateStructuredReference(formats strfmt.Registry) error {

	if swag.IsZero(m.StructuredReference) { // not required
		return nil
	}

	if m.StructuredReference != nil {
		if err := m.StructuredReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "structured_reference")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateSwift(formats strfmt.Registry) error {

	if swag.IsZero(m.Swift) { // not required
		return nil
	}

	if m.Swift != nil {
		if err := m.Swift.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "swift")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateUltimateBeneficiary(formats strfmt.Registry) error {

	if swag.IsZero(m.UltimateBeneficiary) { // not required
		return nil
	}

	if m.UltimateBeneficiary != nil {
		if err := m.UltimateBeneficiary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "ultimate_beneficiary")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateUltimateDebtor(formats strfmt.Registry) error {

	if swag.IsZero(m.UltimateDebtor) { // not required
		return nil
	}

	if m.UltimateDebtor != nil {
		if err := m.UltimateDebtor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "ultimate_debtor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesBeneficiaryParty payment attributes beneficiary party
//
// swagger:model PaymentAttributesBeneficiaryParty
type PaymentAttributesBeneficiaryParty struct {

	// Name of beneficiary as given with account
	AccountName string `json:"account_name,omitempty"`

	// Beneficiary account number
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// The type of the account given with `beneficiary_party.account_number`. Single digit number. Only required if requested by the beneficiary party. Defaults to 0.
	AccountType int64 `json:"account_type,omitempty"`

	// account with
	AccountWith *BeneficiaryDebtorAccountHoldingEntity `json:"account_with,omitempty"`

	// Beneficiary address
	Address []string `json:"address,omitempty"`

	// Beneficiary birth city
	BirthCity string `json:"birth_city,omitempty"`

	// Beneficiary birth country, ISO 3166 format country code
	BirthCountry string `json:"birth_country,omitempty"`

	// Beneficiary birth date. Formatted according to ISO 8601 format: YYYY-MM-DD
	// Format: date
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`

	// Beneficiary birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// Country of the beneficiary address, ISO 3166 format country code
	Country string `json:"country,omitempty"`

	// Beneficiary name
	Name string `json:"name,omitempty"`

	// Organisation identification of a beneficiary, used in the case that the beneficiary is an organisation and not a private person
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// The code that specifies the type of `organisation_identification`
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// Issuer of the organisation identification
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`

	// Beneficiary phone number
	TelephoneNumber string `json:"telephone_number,omitempty"`
}

// Validate validates this payment attributes beneficiary party
func (m *PaymentAttributesBeneficiaryParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributesBeneficiaryParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "account_number_code")
		}
		return err
	}

	return nil
}

func (m *PaymentAttributesBeneficiaryParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "account_with")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributesBeneficiaryParty) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"beneficiary_party"+"."+"birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesBeneficiaryParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesBeneficiaryParty) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesBeneficiaryParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesDebtorParty payment attributes debtor party
//
// swagger:model PaymentAttributesDebtorParty
type PaymentAttributesDebtorParty struct {

	// Name of debtor as given with account
	AccountName string `json:"account_name,omitempty"`

	// Debtor account number. Allows upper case and numeric characters.
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account with
	AccountWith *BeneficiaryDebtorAccountHoldingEntity `json:"account_with,omitempty"`

	// Debtor address
	Address []string `json:"address,omitempty"`

	// Debtor birth city
	BirthCity string `json:"birth_city,omitempty"`

	// Debtor birth country. ISO 3166 format country code
	BirthCountry string `json:"birth_country,omitempty"`

	// Debtor birth date. Formatted according to ISO 8601 format: YYYY-MM-DD
	// Format: date
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`

	// Debtor birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// Country of debtor address. ISO 3166 format country code"
	Country string `json:"country,omitempty"`

	// SWIFT BIC for ordering customer, either BIC8 or BIC11
	CustomerID string `json:"customer_id,omitempty"`

	// Code for `customer_id`
	CustomerIDCode string `json:"customer_id_code,omitempty"`

	// Debtor name
	Name string `json:"name,omitempty"`

	// Organisation identification of a debtor, in the case that the debtor is an organisation and not a private person
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// The code that specifies the type of `organisation_identification`
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// Issuer of the `organisation_identification`
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`
}

// Validate validates this payment attributes debtor party
func (m *PaymentAttributesDebtorParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributesDebtorParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "debtor_party" + "." + "account_number_code")
		}
		return err
	}

	return nil
}

func (m *PaymentAttributesDebtorParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "debtor_party" + "." + "account_with")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributesDebtorParty) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"debtor_party"+"."+"birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesDebtorParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesDebtorParty) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesDebtorParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesFx payment attributes fx
//
// swagger:model PaymentAttributesFx
type PaymentAttributesFx struct {

	// Reference to the foreign exchange contract associated with the transaction
	ContractReference string `json:"contract_reference,omitempty"`

	// Factor used to convert an amount from the instructed currency into the transaction currency: i.e. to convert the `fx.original_amount`, expressed in the `fx.original_currency`, to `amount` specified in `currency`. Decimal value, represented as a string, maximum length 12. Must be > 0.
	ExchangeRate string `json:"exchange_rate,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as instructed by the initiating party. Decimal value. Must be > 0.
	OriginalAmount string `json:"original_amount,omitempty"`

	// Currency of `orginal_amount`. Currency code as defined in ISO 4217.
	OriginalCurrency string `json:"original_currency,omitempty"`
}

// Validate validates this payment attributes fx
func (m *PaymentAttributesFx) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesFx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesFx) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesFx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesStructuredReference payment attributes structured reference
//
// swagger:model PaymentAttributesStructuredReference
type PaymentAttributesStructuredReference struct {

	// Issuer of remittance reference
	Issuer string `json:"issuer,omitempty"`

	// Unique reference to unambiguously refer to the payment originated by the creditor, this reference enables reconciliation by the creditor upon receipt of the amount of money.
	Reference string `json:"reference,omitempty"`
}

// Validate validates this payment attributes structured reference
func (m *PaymentAttributesStructuredReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesStructuredReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesStructuredReference) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesStructuredReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesSwift payment attributes swift
//
// swagger:model PaymentAttributesSwift
type PaymentAttributesSwift struct {

	// SWIFT service level
	BankOperationCode string `json:"bank_operation_code,omitempty"`

	// header
	Header *PaymentAttributesSwiftHeader `json:"header,omitempty"`

	// A SWIFT instruction code
	InstructionCode string `json:"instruction_code,omitempty"`

	// This field specifies additional information for the Receiver or other party specified.
	SenderReceiverInformation string `json:"sender_receiver_information,omitempty"`

	// This repetitive field specifies one or several time indication(s) related to the processing of the payment instruction.
	TimeIndication string `json:"time_indication,omitempty"`
}

// Validate validates this payment attributes swift
func (m *PaymentAttributesSwift) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributesSwift) validateHeader(formats strfmt.Registry) error {

	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "swift" + "." + "header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesSwift) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesSwift) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesSwift
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesSwiftHeader payment attributes swift header
//
// swagger:model PaymentAttributesSwiftHeader
type PaymentAttributesSwiftHeader struct {

	// Destination SWIFT logical terminal address. Complete 12-character SWIFT destination, including BIC (x8), logical terminal code (x1) and branch code (x).
	Destination string `json:"destination,omitempty"`

	// The message type of the SWIFT payment, has to match `[A-Z]{2}[0-9]{3}`. Currently `MT103` is the only supported value
	MessageType string `json:"message_type,omitempty"`

	// SWIFT priority. Either `Normal` or `Priority`.
	Priority string `json:"priority,omitempty"`

	// The destination SWIFT BIC for SWIFT MT messages being sent by Form3 client to SWIFT. Formatted as BIC8 or BIC11.
	Recipient string `json:"recipient,omitempty"`

	// The source SWIFT BIC for SWIFT MT messages being received by Form3 client from SWIFT. Formatted as BIC8 or BIC11.
	Source string `json:"source,omitempty"`

	// Message User Reference (MUR) value, which can be up to 16 characters, and will be returned in the ACK
	UserReference string `json:"user_reference,omitempty"`
}

// Validate validates this payment attributes swift header
func (m *PaymentAttributesSwiftHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesSwiftHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesSwiftHeader) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesSwiftHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
