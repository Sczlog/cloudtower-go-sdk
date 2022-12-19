// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SoftwareEdition software edition
//
// swagger:model SoftwareEdition
type SoftwareEdition string

func NewSoftwareEdition(value SoftwareEdition) *SoftwareEdition {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SoftwareEdition.
func (m SoftwareEdition) Pointer() *SoftwareEdition {
	return &m
}

const (

	// SoftwareEditionCOMMUNITY captures enum value "COMMUNITY"
	SoftwareEditionCOMMUNITY SoftwareEdition = "COMMUNITY"

	// SoftwareEditionENTERPRISE captures enum value "ENTERPRISE"
	SoftwareEditionENTERPRISE SoftwareEdition = "ENTERPRISE"

	// SoftwareEditionESSENTIAL captures enum value "ESSENTIAL"
	SoftwareEditionESSENTIAL SoftwareEdition = "ESSENTIAL"

	// SoftwareEditionEXPRESS captures enum value "EXPRESS"
	SoftwareEditionEXPRESS SoftwareEdition = "EXPRESS"

	// SoftwareEditionSTANDARD captures enum value "STANDARD"
	SoftwareEditionSTANDARD SoftwareEdition = "STANDARD"

	// SoftwareEditionTRIAL captures enum value "TRIAL"
	SoftwareEditionTRIAL SoftwareEdition = "TRIAL"
)

// for schema
var softwareEditionEnum []interface{}

func init() {
	var res []SoftwareEdition
	if err := json.Unmarshal([]byte(`["COMMUNITY","ENTERPRISE","ESSENTIAL","EXPRESS","STANDARD","TRIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareEditionEnum = append(softwareEditionEnum, v)
	}
}

func (m SoftwareEdition) validateSoftwareEditionEnum(path, location string, value SoftwareEdition) error {
	if err := validate.EnumCase(path, location, value, softwareEditionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this software edition
func (m SoftwareEdition) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSoftwareEditionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this software edition based on context it is used
func (m SoftwareEdition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
