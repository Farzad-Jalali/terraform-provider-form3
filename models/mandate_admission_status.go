// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MandateAdmissionStatus mandate admission status
//
// swagger:model MandateAdmissionStatus
type MandateAdmissionStatus string

const (

	// MandateAdmissionStatusConfirmed captures enum value "confirmed"
	MandateAdmissionStatusConfirmed MandateAdmissionStatus = "confirmed"

	// MandateAdmissionStatusFailed captures enum value "failed"
	MandateAdmissionStatusFailed MandateAdmissionStatus = "failed"
)

// for schema
var mandateAdmissionStatusEnum []interface{}

func init() {
	var res []MandateAdmissionStatus
	if err := json.Unmarshal([]byte(`["confirmed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mandateAdmissionStatusEnum = append(mandateAdmissionStatusEnum, v)
	}
}

func (m MandateAdmissionStatus) validateMandateAdmissionStatusEnum(path, location string, value MandateAdmissionStatus) error {
	if err := validate.Enum(path, location, value, mandateAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this mandate admission status
func (m MandateAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMandateAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
