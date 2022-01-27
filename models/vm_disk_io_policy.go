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

// VMDiskIoPolicy Vm disk io policy
//
// swagger:model VmDiskIoPolicy
type VMDiskIoPolicy string

func NewVMDiskIoPolicy(value VMDiskIoPolicy) *VMDiskIoPolicy {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VMDiskIoPolicy.
func (m VMDiskIoPolicy) Pointer() *VMDiskIoPolicy {
	return &m
}

const (

	// VMDiskIoPolicyRESTRICTEACHDISK captures enum value "RESTRICT_EACH_DISK"
	VMDiskIoPolicyRESTRICTEACHDISK VMDiskIoPolicy = "RESTRICT_EACH_DISK"

	// VMDiskIoPolicyRESTRICTWHOLEVM captures enum value "RESTRICT_WHOLE_VM"
	VMDiskIoPolicyRESTRICTWHOLEVM VMDiskIoPolicy = "RESTRICT_WHOLE_VM"
)

// for schema
var vmDiskIoPolicyEnum []interface{}

func init() {
	var res []VMDiskIoPolicy
	if err := json.Unmarshal([]byte(`["RESTRICT_EACH_DISK","RESTRICT_WHOLE_VM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmDiskIoPolicyEnum = append(vmDiskIoPolicyEnum, v)
	}
}

func (m VMDiskIoPolicy) validateVMDiskIoPolicyEnum(path, location string, value VMDiskIoPolicy) error {
	if err := validate.EnumCase(path, location, value, vmDiskIoPolicyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm disk io policy
func (m VMDiskIoPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMDiskIoPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm disk io policy based on context it is used
func (m VMDiskIoPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
