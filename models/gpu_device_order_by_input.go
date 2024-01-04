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

// GpuDeviceOrderByInput gpu device order by input
//
// swagger:model GpuDeviceOrderByInput
type GpuDeviceOrderByInput string

func NewGpuDeviceOrderByInput(value GpuDeviceOrderByInput) *GpuDeviceOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GpuDeviceOrderByInput.
func (m GpuDeviceOrderByInput) Pointer() *GpuDeviceOrderByInput {
	return &m
}

const (

	// GpuDeviceOrderByInputAssignedVgpusNumASC captures enum value "assigned_vgpus_num_ASC"
	GpuDeviceOrderByInputAssignedVgpusNumASC GpuDeviceOrderByInput = "assigned_vgpus_num_ASC"

	// GpuDeviceOrderByInputAssignedVgpusNumDESC captures enum value "assigned_vgpus_num_DESC"
	GpuDeviceOrderByInputAssignedVgpusNumDESC GpuDeviceOrderByInput = "assigned_vgpus_num_DESC"

	// GpuDeviceOrderByInputAvailableVgpusNumASC captures enum value "available_vgpus_num_ASC"
	GpuDeviceOrderByInputAvailableVgpusNumASC GpuDeviceOrderByInput = "available_vgpus_num_ASC"

	// GpuDeviceOrderByInputAvailableVgpusNumDESC captures enum value "available_vgpus_num_DESC"
	GpuDeviceOrderByInputAvailableVgpusNumDESC GpuDeviceOrderByInput = "available_vgpus_num_DESC"

	// GpuDeviceOrderByInputBrandASC captures enum value "brand_ASC"
	GpuDeviceOrderByInputBrandASC GpuDeviceOrderByInput = "brand_ASC"

	// GpuDeviceOrderByInputBrandDESC captures enum value "brand_DESC"
	GpuDeviceOrderByInputBrandDESC GpuDeviceOrderByInput = "brand_DESC"

	// GpuDeviceOrderByInputBusLocationASC captures enum value "bus_location_ASC"
	GpuDeviceOrderByInputBusLocationASC GpuDeviceOrderByInput = "bus_location_ASC"

	// GpuDeviceOrderByInputBusLocationDESC captures enum value "bus_location_DESC"
	GpuDeviceOrderByInputBusLocationDESC GpuDeviceOrderByInput = "bus_location_DESC"

	// GpuDeviceOrderByInputDescriptionASC captures enum value "description_ASC"
	GpuDeviceOrderByInputDescriptionASC GpuDeviceOrderByInput = "description_ASC"

	// GpuDeviceOrderByInputDescriptionDESC captures enum value "description_DESC"
	GpuDeviceOrderByInputDescriptionDESC GpuDeviceOrderByInput = "description_DESC"

	// GpuDeviceOrderByInputDriverInfoASC captures enum value "driver_info_ASC"
	GpuDeviceOrderByInputDriverInfoASC GpuDeviceOrderByInput = "driver_info_ASC"

	// GpuDeviceOrderByInputDriverInfoDESC captures enum value "driver_info_DESC"
	GpuDeviceOrderByInputDriverInfoDESC GpuDeviceOrderByInput = "driver_info_DESC"

	// GpuDeviceOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	GpuDeviceOrderByInputEntityAsyncStatusASC GpuDeviceOrderByInput = "entityAsyncStatus_ASC"

	// GpuDeviceOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	GpuDeviceOrderByInputEntityAsyncStatusDESC GpuDeviceOrderByInput = "entityAsyncStatus_DESC"

	// GpuDeviceOrderByInputIDASC captures enum value "id_ASC"
	GpuDeviceOrderByInputIDASC GpuDeviceOrderByInput = "id_ASC"

	// GpuDeviceOrderByInputIDDESC captures enum value "id_DESC"
	GpuDeviceOrderByInputIDDESC GpuDeviceOrderByInput = "id_DESC"

	// GpuDeviceOrderByInputIsNvidiaToolsReadyASC captures enum value "is_nvidia_tools_ready_ASC"
	GpuDeviceOrderByInputIsNvidiaToolsReadyASC GpuDeviceOrderByInput = "is_nvidia_tools_ready_ASC"

	// GpuDeviceOrderByInputIsNvidiaToolsReadyDESC captures enum value "is_nvidia_tools_ready_DESC"
	GpuDeviceOrderByInputIsNvidiaToolsReadyDESC GpuDeviceOrderByInput = "is_nvidia_tools_ready_DESC"

	// GpuDeviceOrderByInputIsNvidiaVfsEnabledASC captures enum value "is_nvidia_vfs_enabled_ASC"
	GpuDeviceOrderByInputIsNvidiaVfsEnabledASC GpuDeviceOrderByInput = "is_nvidia_vfs_enabled_ASC"

	// GpuDeviceOrderByInputIsNvidiaVfsEnabledDESC captures enum value "is_nvidia_vfs_enabled_DESC"
	GpuDeviceOrderByInputIsNvidiaVfsEnabledDESC GpuDeviceOrderByInput = "is_nvidia_vfs_enabled_DESC"

	// GpuDeviceOrderByInputIsNvidiaVfsSupportedASC captures enum value "is_nvidia_vfs_supported_ASC"
	GpuDeviceOrderByInputIsNvidiaVfsSupportedASC GpuDeviceOrderByInput = "is_nvidia_vfs_supported_ASC"

	// GpuDeviceOrderByInputIsNvidiaVfsSupportedDESC captures enum value "is_nvidia_vfs_supported_DESC"
	GpuDeviceOrderByInputIsNvidiaVfsSupportedDESC GpuDeviceOrderByInput = "is_nvidia_vfs_supported_DESC"

	// GpuDeviceOrderByInputLocalCreatedAtASC captures enum value "local_created_at_ASC"
	GpuDeviceOrderByInputLocalCreatedAtASC GpuDeviceOrderByInput = "local_created_at_ASC"

	// GpuDeviceOrderByInputLocalCreatedAtDESC captures enum value "local_created_at_DESC"
	GpuDeviceOrderByInputLocalCreatedAtDESC GpuDeviceOrderByInput = "local_created_at_DESC"

	// GpuDeviceOrderByInputLocalIDASC captures enum value "local_id_ASC"
	GpuDeviceOrderByInputLocalIDASC GpuDeviceOrderByInput = "local_id_ASC"

	// GpuDeviceOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	GpuDeviceOrderByInputLocalIDDESC GpuDeviceOrderByInput = "local_id_DESC"

	// GpuDeviceOrderByInputMdevSupportedTypesASC captures enum value "mdev_supported_types_ASC"
	GpuDeviceOrderByInputMdevSupportedTypesASC GpuDeviceOrderByInput = "mdev_supported_types_ASC"

	// GpuDeviceOrderByInputMdevSupportedTypesDESC captures enum value "mdev_supported_types_DESC"
	GpuDeviceOrderByInputMdevSupportedTypesDESC GpuDeviceOrderByInput = "mdev_supported_types_DESC"

	// GpuDeviceOrderByInputModelASC captures enum value "model_ASC"
	GpuDeviceOrderByInputModelASC GpuDeviceOrderByInput = "model_ASC"

	// GpuDeviceOrderByInputModelDESC captures enum value "model_DESC"
	GpuDeviceOrderByInputModelDESC GpuDeviceOrderByInput = "model_DESC"

	// GpuDeviceOrderByInputNameASC captures enum value "name_ASC"
	GpuDeviceOrderByInputNameASC GpuDeviceOrderByInput = "name_ASC"

	// GpuDeviceOrderByInputNameDESC captures enum value "name_DESC"
	GpuDeviceOrderByInputNameDESC GpuDeviceOrderByInput = "name_DESC"

	// GpuDeviceOrderByInputStatusASC captures enum value "status_ASC"
	GpuDeviceOrderByInputStatusASC GpuDeviceOrderByInput = "status_ASC"

	// GpuDeviceOrderByInputStatusDESC captures enum value "status_DESC"
	GpuDeviceOrderByInputStatusDESC GpuDeviceOrderByInput = "status_DESC"

	// GpuDeviceOrderByInputUserUsageASC captures enum value "user_usage_ASC"
	GpuDeviceOrderByInputUserUsageASC GpuDeviceOrderByInput = "user_usage_ASC"

	// GpuDeviceOrderByInputUserUsageDESC captures enum value "user_usage_DESC"
	GpuDeviceOrderByInputUserUsageDESC GpuDeviceOrderByInput = "user_usage_DESC"

	// GpuDeviceOrderByInputUserVgpuTypeIDASC captures enum value "user_vgpu_type_id_ASC"
	GpuDeviceOrderByInputUserVgpuTypeIDASC GpuDeviceOrderByInput = "user_vgpu_type_id_ASC"

	// GpuDeviceOrderByInputUserVgpuTypeIDDESC captures enum value "user_vgpu_type_id_DESC"
	GpuDeviceOrderByInputUserVgpuTypeIDDESC GpuDeviceOrderByInput = "user_vgpu_type_id_DESC"

	// GpuDeviceOrderByInputUserVgpuTypeNameASC captures enum value "user_vgpu_type_name_ASC"
	GpuDeviceOrderByInputUserVgpuTypeNameASC GpuDeviceOrderByInput = "user_vgpu_type_name_ASC"

	// GpuDeviceOrderByInputUserVgpuTypeNameDESC captures enum value "user_vgpu_type_name_DESC"
	GpuDeviceOrderByInputUserVgpuTypeNameDESC GpuDeviceOrderByInput = "user_vgpu_type_name_DESC"

	// GpuDeviceOrderByInputVgpuInstanceNumASC captures enum value "vgpu_instance_num_ASC"
	GpuDeviceOrderByInputVgpuInstanceNumASC GpuDeviceOrderByInput = "vgpu_instance_num_ASC"

	// GpuDeviceOrderByInputVgpuInstanceNumDESC captures enum value "vgpu_instance_num_DESC"
	GpuDeviceOrderByInputVgpuInstanceNumDESC GpuDeviceOrderByInput = "vgpu_instance_num_DESC"
)

// for schema
var gpuDeviceOrderByInputEnum []interface{}

func init() {
	var res []GpuDeviceOrderByInput
	if err := json.Unmarshal([]byte(`["assigned_vgpus_num_ASC","assigned_vgpus_num_DESC","available_vgpus_num_ASC","available_vgpus_num_DESC","brand_ASC","brand_DESC","bus_location_ASC","bus_location_DESC","description_ASC","description_DESC","driver_info_ASC","driver_info_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","is_nvidia_tools_ready_ASC","is_nvidia_tools_ready_DESC","is_nvidia_vfs_enabled_ASC","is_nvidia_vfs_enabled_DESC","is_nvidia_vfs_supported_ASC","is_nvidia_vfs_supported_DESC","local_created_at_ASC","local_created_at_DESC","local_id_ASC","local_id_DESC","mdev_supported_types_ASC","mdev_supported_types_DESC","model_ASC","model_DESC","name_ASC","name_DESC","status_ASC","status_DESC","user_usage_ASC","user_usage_DESC","user_vgpu_type_id_ASC","user_vgpu_type_id_DESC","user_vgpu_type_name_ASC","user_vgpu_type_name_DESC","vgpu_instance_num_ASC","vgpu_instance_num_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gpuDeviceOrderByInputEnum = append(gpuDeviceOrderByInputEnum, v)
	}
}

func (m GpuDeviceOrderByInput) validateGpuDeviceOrderByInputEnum(path, location string, value GpuDeviceOrderByInput) error {
	if err := validate.EnumCase(path, location, value, gpuDeviceOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this gpu device order by input
func (m GpuDeviceOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGpuDeviceOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this gpu device order by input based on context it is used
func (m GpuDeviceOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}