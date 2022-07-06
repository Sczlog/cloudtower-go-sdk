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

// VMDiskOrderByInput Vm disk order by input
//
// swagger:model VmDiskOrderByInput
type VMDiskOrderByInput string

func NewVMDiskOrderByInput(value VMDiskOrderByInput) *VMDiskOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VMDiskOrderByInput.
func (m VMDiskOrderByInput) Pointer() *VMDiskOrderByInput {
	return &m
}

const (

	// VMDiskOrderByInputBootASC captures enum value "boot_ASC"
	VMDiskOrderByInputBootASC VMDiskOrderByInput = "boot_ASC"

	// VMDiskOrderByInputBootDESC captures enum value "boot_DESC"
	VMDiskOrderByInputBootDESC VMDiskOrderByInput = "boot_DESC"

	// VMDiskOrderByInputBusASC captures enum value "bus_ASC"
	VMDiskOrderByInputBusASC VMDiskOrderByInput = "bus_ASC"

	// VMDiskOrderByInputBusDESC captures enum value "bus_DESC"
	VMDiskOrderByInputBusDESC VMDiskOrderByInput = "bus_DESC"

	// VMDiskOrderByInputCloudInitImageNameASC captures enum value "cloud_init_image_name_ASC"
	VMDiskOrderByInputCloudInitImageNameASC VMDiskOrderByInput = "cloud_init_image_name_ASC"

	// VMDiskOrderByInputCloudInitImageNameDESC captures enum value "cloud_init_image_name_DESC"
	VMDiskOrderByInputCloudInitImageNameDESC VMDiskOrderByInput = "cloud_init_image_name_DESC"

	// VMDiskOrderByInputCloudInitImagePathASC captures enum value "cloud_init_image_path_ASC"
	VMDiskOrderByInputCloudInitImagePathASC VMDiskOrderByInput = "cloud_init_image_path_ASC"

	// VMDiskOrderByInputCloudInitImagePathDESC captures enum value "cloud_init_image_path_DESC"
	VMDiskOrderByInputCloudInitImagePathDESC VMDiskOrderByInput = "cloud_init_image_path_DESC"

	// VMDiskOrderByInputDeviceASC captures enum value "device_ASC"
	VMDiskOrderByInputDeviceASC VMDiskOrderByInput = "device_ASC"

	// VMDiskOrderByInputDeviceDESC captures enum value "device_DESC"
	VMDiskOrderByInputDeviceDESC VMDiskOrderByInput = "device_DESC"

	// VMDiskOrderByInputDisabledASC captures enum value "disabled_ASC"
	VMDiskOrderByInputDisabledASC VMDiskOrderByInput = "disabled_ASC"

	// VMDiskOrderByInputDisabledDESC captures enum value "disabled_DESC"
	VMDiskOrderByInputDisabledDESC VMDiskOrderByInput = "disabled_DESC"

	// VMDiskOrderByInputIDASC captures enum value "id_ASC"
	VMDiskOrderByInputIDASC VMDiskOrderByInput = "id_ASC"

	// VMDiskOrderByInputIDDESC captures enum value "id_DESC"
	VMDiskOrderByInputIDDESC VMDiskOrderByInput = "id_DESC"

	// VMDiskOrderByInputKeyASC captures enum value "key_ASC"
	VMDiskOrderByInputKeyASC VMDiskOrderByInput = "key_ASC"

	// VMDiskOrderByInputKeyDESC captures enum value "key_DESC"
	VMDiskOrderByInputKeyDESC VMDiskOrderByInput = "key_DESC"

	// VMDiskOrderByInputMaxBandwidthASC captures enum value "max_bandwidth_ASC"
	VMDiskOrderByInputMaxBandwidthASC VMDiskOrderByInput = "max_bandwidth_ASC"

	// VMDiskOrderByInputMaxBandwidthDESC captures enum value "max_bandwidth_DESC"
	VMDiskOrderByInputMaxBandwidthDESC VMDiskOrderByInput = "max_bandwidth_DESC"

	// VMDiskOrderByInputMaxBandwidthPolicyASC captures enum value "max_bandwidth_policy_ASC"
	VMDiskOrderByInputMaxBandwidthPolicyASC VMDiskOrderByInput = "max_bandwidth_policy_ASC"

	// VMDiskOrderByInputMaxBandwidthPolicyDESC captures enum value "max_bandwidth_policy_DESC"
	VMDiskOrderByInputMaxBandwidthPolicyDESC VMDiskOrderByInput = "max_bandwidth_policy_DESC"

	// VMDiskOrderByInputMaxIopsASC captures enum value "max_iops_ASC"
	VMDiskOrderByInputMaxIopsASC VMDiskOrderByInput = "max_iops_ASC"

	// VMDiskOrderByInputMaxIopsDESC captures enum value "max_iops_DESC"
	VMDiskOrderByInputMaxIopsDESC VMDiskOrderByInput = "max_iops_DESC"

	// VMDiskOrderByInputMaxIopsPolicyASC captures enum value "max_iops_policy_ASC"
	VMDiskOrderByInputMaxIopsPolicyASC VMDiskOrderByInput = "max_iops_policy_ASC"

	// VMDiskOrderByInputMaxIopsPolicyDESC captures enum value "max_iops_policy_DESC"
	VMDiskOrderByInputMaxIopsPolicyDESC VMDiskOrderByInput = "max_iops_policy_DESC"

	// VMDiskOrderByInputSerialASC captures enum value "serial_ASC"
	VMDiskOrderByInputSerialASC VMDiskOrderByInput = "serial_ASC"

	// VMDiskOrderByInputSerialDESC captures enum value "serial_DESC"
	VMDiskOrderByInputSerialDESC VMDiskOrderByInput = "serial_DESC"

	// VMDiskOrderByInputTypeASC captures enum value "type_ASC"
	VMDiskOrderByInputTypeASC VMDiskOrderByInput = "type_ASC"

	// VMDiskOrderByInputTypeDESC captures enum value "type_DESC"
	VMDiskOrderByInputTypeDESC VMDiskOrderByInput = "type_DESC"

	// VMDiskOrderByInputUnsafeImagePathASC captures enum value "unsafe_image_path_ASC"
	VMDiskOrderByInputUnsafeImagePathASC VMDiskOrderByInput = "unsafe_image_path_ASC"

	// VMDiskOrderByInputUnsafeImagePathDESC captures enum value "unsafe_image_path_DESC"
	VMDiskOrderByInputUnsafeImagePathDESC VMDiskOrderByInput = "unsafe_image_path_DESC"

	// VMDiskOrderByInputUnsafeImageUUIDASC captures enum value "unsafe_image_uuid_ASC"
	VMDiskOrderByInputUnsafeImageUUIDASC VMDiskOrderByInput = "unsafe_image_uuid_ASC"

	// VMDiskOrderByInputUnsafeImageUUIDDESC captures enum value "unsafe_image_uuid_DESC"
	VMDiskOrderByInputUnsafeImageUUIDDESC VMDiskOrderByInput = "unsafe_image_uuid_DESC"

	// VMDiskOrderByInputUnsafeProvisionASC captures enum value "unsafe_provision_ASC"
	VMDiskOrderByInputUnsafeProvisionASC VMDiskOrderByInput = "unsafe_provision_ASC"

	// VMDiskOrderByInputUnsafeProvisionDESC captures enum value "unsafe_provision_DESC"
	VMDiskOrderByInputUnsafeProvisionDESC VMDiskOrderByInput = "unsafe_provision_DESC"
)

// for schema
var vmDiskOrderByInputEnum []interface{}

func init() {
	var res []VMDiskOrderByInput
	if err := json.Unmarshal([]byte(`["boot_ASC","boot_DESC","bus_ASC","bus_DESC","cloud_init_image_name_ASC","cloud_init_image_name_DESC","cloud_init_image_path_ASC","cloud_init_image_path_DESC","device_ASC","device_DESC","disabled_ASC","disabled_DESC","id_ASC","id_DESC","key_ASC","key_DESC","max_bandwidth_ASC","max_bandwidth_DESC","max_bandwidth_policy_ASC","max_bandwidth_policy_DESC","max_iops_ASC","max_iops_DESC","max_iops_policy_ASC","max_iops_policy_DESC","serial_ASC","serial_DESC","type_ASC","type_DESC","unsafe_image_path_ASC","unsafe_image_path_DESC","unsafe_image_uuid_ASC","unsafe_image_uuid_DESC","unsafe_provision_ASC","unsafe_provision_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmDiskOrderByInputEnum = append(vmDiskOrderByInputEnum, v)
	}
}

func (m VMDiskOrderByInput) validateVMDiskOrderByInputEnum(path, location string, value VMDiskOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vmDiskOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm disk order by input
func (m VMDiskOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMDiskOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm disk order by input based on context it is used
func (m VMDiskOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
