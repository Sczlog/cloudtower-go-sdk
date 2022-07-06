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

// HostOrderByInput host order by input
//
// swagger:model HostOrderByInput
type HostOrderByInput string

func NewHostOrderByInput(value HostOrderByInput) *HostOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HostOrderByInput.
func (m HostOrderByInput) Pointer() *HostOrderByInput {
	return &m
}

const (

	// HostOrderByInputAccessIPASC captures enum value "access_ip_ASC"
	HostOrderByInputAccessIPASC HostOrderByInput = "access_ip_ASC"

	// HostOrderByInputAccessIPDESC captures enum value "access_ip_DESC"
	HostOrderByInputAccessIPDESC HostOrderByInput = "access_ip_DESC"

	// HostOrderByInputAllocatableMemoryBytesASC captures enum value "allocatable_memory_bytes_ASC"
	HostOrderByInputAllocatableMemoryBytesASC HostOrderByInput = "allocatable_memory_bytes_ASC"

	// HostOrderByInputAllocatableMemoryBytesDESC captures enum value "allocatable_memory_bytes_DESC"
	HostOrderByInputAllocatableMemoryBytesDESC HostOrderByInput = "allocatable_memory_bytes_DESC"

	// HostOrderByInputChunkIDASC captures enum value "chunk_id_ASC"
	HostOrderByInputChunkIDASC HostOrderByInput = "chunk_id_ASC"

	// HostOrderByInputChunkIDDESC captures enum value "chunk_id_DESC"
	HostOrderByInputChunkIDDESC HostOrderByInput = "chunk_id_DESC"

	// HostOrderByInputCPUBrandASC captures enum value "cpu_brand_ASC"
	HostOrderByInputCPUBrandASC HostOrderByInput = "cpu_brand_ASC"

	// HostOrderByInputCPUBrandDESC captures enum value "cpu_brand_DESC"
	HostOrderByInputCPUBrandDESC HostOrderByInput = "cpu_brand_DESC"

	// HostOrderByInputCPUFanSpeedUnitASC captures enum value "cpu_fan_speed_unit_ASC"
	HostOrderByInputCPUFanSpeedUnitASC HostOrderByInput = "cpu_fan_speed_unit_ASC"

	// HostOrderByInputCPUFanSpeedUnitDESC captures enum value "cpu_fan_speed_unit_DESC"
	HostOrderByInputCPUFanSpeedUnitDESC HostOrderByInput = "cpu_fan_speed_unit_DESC"

	// HostOrderByInputCPUHzPerCoreASC captures enum value "cpu_hz_per_core_ASC"
	HostOrderByInputCPUHzPerCoreASC HostOrderByInput = "cpu_hz_per_core_ASC"

	// HostOrderByInputCPUHzPerCoreDESC captures enum value "cpu_hz_per_core_DESC"
	HostOrderByInputCPUHzPerCoreDESC HostOrderByInput = "cpu_hz_per_core_DESC"

	// HostOrderByInputCPUModelASC captures enum value "cpu_model_ASC"
	HostOrderByInputCPUModelASC HostOrderByInput = "cpu_model_ASC"

	// HostOrderByInputCPUModelDESC captures enum value "cpu_model_DESC"
	HostOrderByInputCPUModelDESC HostOrderByInput = "cpu_model_DESC"

	// HostOrderByInputCPUVendorASC captures enum value "cpu_vendor_ASC"
	HostOrderByInputCPUVendorASC HostOrderByInput = "cpu_vendor_ASC"

	// HostOrderByInputCPUVendorDESC captures enum value "cpu_vendor_DESC"
	HostOrderByInputCPUVendorDESC HostOrderByInput = "cpu_vendor_DESC"

	// HostOrderByInputDataIPASC captures enum value "data_ip_ASC"
	HostOrderByInputDataIPASC HostOrderByInput = "data_ip_ASC"

	// HostOrderByInputDataIPDESC captures enum value "data_ip_DESC"
	HostOrderByInputDataIPDESC HostOrderByInput = "data_ip_DESC"

	// HostOrderByInputFailureDataSpaceASC captures enum value "failure_data_space_ASC"
	HostOrderByInputFailureDataSpaceASC HostOrderByInput = "failure_data_space_ASC"

	// HostOrderByInputFailureDataSpaceDESC captures enum value "failure_data_space_DESC"
	HostOrderByInputFailureDataSpaceDESC HostOrderByInput = "failure_data_space_DESC"

	// HostOrderByInputHddDataCapacityASC captures enum value "hdd_data_capacity_ASC"
	HostOrderByInputHddDataCapacityASC HostOrderByInput = "hdd_data_capacity_ASC"

	// HostOrderByInputHddDataCapacityDESC captures enum value "hdd_data_capacity_DESC"
	HostOrderByInputHddDataCapacityDESC HostOrderByInput = "hdd_data_capacity_DESC"

	// HostOrderByInputHddDiskCountASC captures enum value "hdd_disk_count_ASC"
	HostOrderByInputHddDiskCountASC HostOrderByInput = "hdd_disk_count_ASC"

	// HostOrderByInputHddDiskCountDESC captures enum value "hdd_disk_count_DESC"
	HostOrderByInputHddDiskCountDESC HostOrderByInput = "hdd_disk_count_DESC"

	// HostOrderByInputHypervisorIPASC captures enum value "hypervisor_ip_ASC"
	HostOrderByInputHypervisorIPASC HostOrderByInput = "hypervisor_ip_ASC"

	// HostOrderByInputHypervisorIPDESC captures enum value "hypervisor_ip_DESC"
	HostOrderByInputHypervisorIPDESC HostOrderByInput = "hypervisor_ip_DESC"

	// HostOrderByInputIDASC captures enum value "id_ASC"
	HostOrderByInputIDASC HostOrderByInput = "id_ASC"

	// HostOrderByInputIDDESC captures enum value "id_DESC"
	HostOrderByInputIDDESC HostOrderByInput = "id_DESC"

	// HostOrderByInputIsOsInRaid1ASC captures enum value "is_os_in_raid1_ASC"
	HostOrderByInputIsOsInRaid1ASC HostOrderByInput = "is_os_in_raid1_ASC"

	// HostOrderByInputIsOsInRaid1DESC captures enum value "is_os_in_raid1_DESC"
	HostOrderByInputIsOsInRaid1DESC HostOrderByInput = "is_os_in_raid1_DESC"

	// HostOrderByInputLocalIDASC captures enum value "local_id_ASC"
	HostOrderByInputLocalIDASC HostOrderByInput = "local_id_ASC"

	// HostOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	HostOrderByInputLocalIDDESC HostOrderByInput = "local_id_DESC"

	// HostOrderByInputLsmCapDiskSafeUmountASC captures enum value "lsm_cap_disk_safe_umount_ASC"
	HostOrderByInputLsmCapDiskSafeUmountASC HostOrderByInput = "lsm_cap_disk_safe_umount_ASC"

	// HostOrderByInputLsmCapDiskSafeUmountDESC captures enum value "lsm_cap_disk_safe_umount_DESC"
	HostOrderByInputLsmCapDiskSafeUmountDESC HostOrderByInput = "lsm_cap_disk_safe_umount_DESC"

	// HostOrderByInputManagementIPASC captures enum value "management_ip_ASC"
	HostOrderByInputManagementIPASC HostOrderByInput = "management_ip_ASC"

	// HostOrderByInputManagementIPDESC captures enum value "management_ip_DESC"
	HostOrderByInputManagementIPDESC HostOrderByInput = "management_ip_DESC"

	// HostOrderByInputModelASC captures enum value "model_ASC"
	HostOrderByInputModelASC HostOrderByInput = "model_ASC"

	// HostOrderByInputModelDESC captures enum value "model_DESC"
	HostOrderByInputModelDESC HostOrderByInput = "model_DESC"

	// HostOrderByInputNameASC captures enum value "name_ASC"
	HostOrderByInputNameASC HostOrderByInput = "name_ASC"

	// HostOrderByInputNameDESC captures enum value "name_DESC"
	HostOrderByInputNameDESC HostOrderByInput = "name_DESC"

	// HostOrderByInputNestedVirtualizationASC captures enum value "nested_virtualization_ASC"
	HostOrderByInputNestedVirtualizationASC HostOrderByInput = "nested_virtualization_ASC"

	// HostOrderByInputNestedVirtualizationDESC captures enum value "nested_virtualization_DESC"
	HostOrderByInputNestedVirtualizationDESC HostOrderByInput = "nested_virtualization_DESC"

	// HostOrderByInputNicCountASC captures enum value "nic_count_ASC"
	HostOrderByInputNicCountASC HostOrderByInput = "nic_count_ASC"

	// HostOrderByInputNicCountDESC captures enum value "nic_count_DESC"
	HostOrderByInputNicCountDESC HostOrderByInput = "nic_count_DESC"

	// HostOrderByInputNodeTopoLocalIDASC captures enum value "node_topo_local_id_ASC"
	HostOrderByInputNodeTopoLocalIDASC HostOrderByInput = "node_topo_local_id_ASC"

	// HostOrderByInputNodeTopoLocalIDDESC captures enum value "node_topo_local_id_DESC"
	HostOrderByInputNodeTopoLocalIDDESC HostOrderByInput = "node_topo_local_id_DESC"

	// HostOrderByInputOsMemoryBytesASC captures enum value "os_memory_bytes_ASC"
	HostOrderByInputOsMemoryBytesASC HostOrderByInput = "os_memory_bytes_ASC"

	// HostOrderByInputOsMemoryBytesDESC captures enum value "os_memory_bytes_DESC"
	HostOrderByInputOsMemoryBytesDESC HostOrderByInput = "os_memory_bytes_DESC"

	// HostOrderByInputOsVersionASC captures enum value "os_version_ASC"
	HostOrderByInputOsVersionASC HostOrderByInput = "os_version_ASC"

	// HostOrderByInputOsVersionDESC captures enum value "os_version_DESC"
	HostOrderByInputOsVersionDESC HostOrderByInput = "os_version_DESC"

	// HostOrderByInputPmemDimmCapacityASC captures enum value "pmem_dimm_capacity_ASC"
	HostOrderByInputPmemDimmCapacityASC HostOrderByInput = "pmem_dimm_capacity_ASC"

	// HostOrderByInputPmemDimmCapacityDESC captures enum value "pmem_dimm_capacity_DESC"
	HostOrderByInputPmemDimmCapacityDESC HostOrderByInput = "pmem_dimm_capacity_DESC"

	// HostOrderByInputPmemDimmCountASC captures enum value "pmem_dimm_count_ASC"
	HostOrderByInputPmemDimmCountASC HostOrderByInput = "pmem_dimm_count_ASC"

	// HostOrderByInputPmemDimmCountDESC captures enum value "pmem_dimm_count_DESC"
	HostOrderByInputPmemDimmCountDESC HostOrderByInput = "pmem_dimm_count_DESC"

	// HostOrderByInputPmemDiskCountASC captures enum value "pmem_disk_count_ASC"
	HostOrderByInputPmemDiskCountASC HostOrderByInput = "pmem_disk_count_ASC"

	// HostOrderByInputPmemDiskCountDESC captures enum value "pmem_disk_count_DESC"
	HostOrderByInputPmemDiskCountDESC HostOrderByInput = "pmem_disk_count_DESC"

	// HostOrderByInputProvisionedCPUCoresASC captures enum value "provisioned_cpu_cores_ASC"
	HostOrderByInputProvisionedCPUCoresASC HostOrderByInput = "provisioned_cpu_cores_ASC"

	// HostOrderByInputProvisionedCPUCoresDESC captures enum value "provisioned_cpu_cores_DESC"
	HostOrderByInputProvisionedCPUCoresDESC HostOrderByInput = "provisioned_cpu_cores_DESC"

	// HostOrderByInputProvisionedMemoryBytesASC captures enum value "provisioned_memory_bytes_ASC"
	HostOrderByInputProvisionedMemoryBytesASC HostOrderByInput = "provisioned_memory_bytes_ASC"

	// HostOrderByInputProvisionedMemoryBytesDESC captures enum value "provisioned_memory_bytes_DESC"
	HostOrderByInputProvisionedMemoryBytesDESC HostOrderByInput = "provisioned_memory_bytes_DESC"

	// HostOrderByInputRunningPauseVMMemoryBytesASC captures enum value "running_pause_vm_memory_bytes_ASC"
	HostOrderByInputRunningPauseVMMemoryBytesASC HostOrderByInput = "running_pause_vm_memory_bytes_ASC"

	// HostOrderByInputRunningPauseVMMemoryBytesDESC captures enum value "running_pause_vm_memory_bytes_DESC"
	HostOrderByInputRunningPauseVMMemoryBytesDESC HostOrderByInput = "running_pause_vm_memory_bytes_DESC"

	// HostOrderByInputRunningVMNumASC captures enum value "running_vm_num_ASC"
	HostOrderByInputRunningVMNumASC HostOrderByInput = "running_vm_num_ASC"

	// HostOrderByInputRunningVMNumDESC captures enum value "running_vm_num_DESC"
	HostOrderByInputRunningVMNumDESC HostOrderByInput = "running_vm_num_DESC"

	// HostOrderByInputScvmCPUASC captures enum value "scvm_cpu_ASC"
	HostOrderByInputScvmCPUASC HostOrderByInput = "scvm_cpu_ASC"

	// HostOrderByInputScvmCPUDESC captures enum value "scvm_cpu_DESC"
	HostOrderByInputScvmCPUDESC HostOrderByInput = "scvm_cpu_DESC"

	// HostOrderByInputScvmMemoryASC captures enum value "scvm_memory_ASC"
	HostOrderByInputScvmMemoryASC HostOrderByInput = "scvm_memory_ASC"

	// HostOrderByInputScvmMemoryDESC captures enum value "scvm_memory_DESC"
	HostOrderByInputScvmMemoryDESC HostOrderByInput = "scvm_memory_DESC"

	// HostOrderByInputScvmNameASC captures enum value "scvm_name_ASC"
	HostOrderByInputScvmNameASC HostOrderByInput = "scvm_name_ASC"

	// HostOrderByInputScvmNameDESC captures enum value "scvm_name_DESC"
	HostOrderByInputScvmNameDESC HostOrderByInput = "scvm_name_DESC"

	// HostOrderByInputSerialASC captures enum value "serial_ASC"
	HostOrderByInputSerialASC HostOrderByInput = "serial_ASC"

	// HostOrderByInputSerialDESC captures enum value "serial_DESC"
	HostOrderByInputSerialDESC HostOrderByInput = "serial_DESC"

	// HostOrderByInputSsdDataCapacityASC captures enum value "ssd_data_capacity_ASC"
	HostOrderByInputSsdDataCapacityASC HostOrderByInput = "ssd_data_capacity_ASC"

	// HostOrderByInputSsdDataCapacityDESC captures enum value "ssd_data_capacity_DESC"
	HostOrderByInputSsdDataCapacityDESC HostOrderByInput = "ssd_data_capacity_DESC"

	// HostOrderByInputSsdDiskCountASC captures enum value "ssd_disk_count_ASC"
	HostOrderByInputSsdDiskCountASC HostOrderByInput = "ssd_disk_count_ASC"

	// HostOrderByInputSsdDiskCountDESC captures enum value "ssd_disk_count_DESC"
	HostOrderByInputSsdDiskCountDESC HostOrderByInput = "ssd_disk_count_DESC"

	// HostOrderByInputStateASC captures enum value "state_ASC"
	HostOrderByInputStateASC HostOrderByInput = "state_ASC"

	// HostOrderByInputStateDESC captures enum value "state_DESC"
	HostOrderByInputStateDESC HostOrderByInput = "state_DESC"

	// HostOrderByInputStatusASC captures enum value "status_ASC"
	HostOrderByInputStatusASC HostOrderByInput = "status_ASC"

	// HostOrderByInputStatusDESC captures enum value "status_DESC"
	HostOrderByInputStatusDESC HostOrderByInput = "status_DESC"

	// HostOrderByInputStoppedVMNumASC captures enum value "stopped_vm_num_ASC"
	HostOrderByInputStoppedVMNumASC HostOrderByInput = "stopped_vm_num_ASC"

	// HostOrderByInputStoppedVMNumDESC captures enum value "stopped_vm_num_DESC"
	HostOrderByInputStoppedVMNumDESC HostOrderByInput = "stopped_vm_num_DESC"

	// HostOrderByInputSuspendedVMNumASC captures enum value "suspended_vm_num_ASC"
	HostOrderByInputSuspendedVMNumASC HostOrderByInput = "suspended_vm_num_ASC"

	// HostOrderByInputSuspendedVMNumDESC captures enum value "suspended_vm_num_DESC"
	HostOrderByInputSuspendedVMNumDESC HostOrderByInput = "suspended_vm_num_DESC"

	// HostOrderByInputTotalCacheCapacityASC captures enum value "total_cache_capacity_ASC"
	HostOrderByInputTotalCacheCapacityASC HostOrderByInput = "total_cache_capacity_ASC"

	// HostOrderByInputTotalCacheCapacityDESC captures enum value "total_cache_capacity_DESC"
	HostOrderByInputTotalCacheCapacityDESC HostOrderByInput = "total_cache_capacity_DESC"

	// HostOrderByInputTotalCPUCoresASC captures enum value "total_cpu_cores_ASC"
	HostOrderByInputTotalCPUCoresASC HostOrderByInput = "total_cpu_cores_ASC"

	// HostOrderByInputTotalCPUCoresDESC captures enum value "total_cpu_cores_DESC"
	HostOrderByInputTotalCPUCoresDESC HostOrderByInput = "total_cpu_cores_DESC"

	// HostOrderByInputTotalCPUHzASC captures enum value "total_cpu_hz_ASC"
	HostOrderByInputTotalCPUHzASC HostOrderByInput = "total_cpu_hz_ASC"

	// HostOrderByInputTotalCPUHzDESC captures enum value "total_cpu_hz_DESC"
	HostOrderByInputTotalCPUHzDESC HostOrderByInput = "total_cpu_hz_DESC"

	// HostOrderByInputTotalCPUSocketsASC captures enum value "total_cpu_sockets_ASC"
	HostOrderByInputTotalCPUSocketsASC HostOrderByInput = "total_cpu_sockets_ASC"

	// HostOrderByInputTotalCPUSocketsDESC captures enum value "total_cpu_sockets_DESC"
	HostOrderByInputTotalCPUSocketsDESC HostOrderByInput = "total_cpu_sockets_DESC"

	// HostOrderByInputTotalDataCapacityASC captures enum value "total_data_capacity_ASC"
	HostOrderByInputTotalDataCapacityASC HostOrderByInput = "total_data_capacity_ASC"

	// HostOrderByInputTotalDataCapacityDESC captures enum value "total_data_capacity_DESC"
	HostOrderByInputTotalDataCapacityDESC HostOrderByInput = "total_data_capacity_DESC"

	// HostOrderByInputTotalMemoryBytesASC captures enum value "total_memory_bytes_ASC"
	HostOrderByInputTotalMemoryBytesASC HostOrderByInput = "total_memory_bytes_ASC"

	// HostOrderByInputTotalMemoryBytesDESC captures enum value "total_memory_bytes_DESC"
	HostOrderByInputTotalMemoryBytesDESC HostOrderByInput = "total_memory_bytes_DESC"

	// HostOrderByInputUsedCPUHzASC captures enum value "used_cpu_hz_ASC"
	HostOrderByInputUsedCPUHzASC HostOrderByInput = "used_cpu_hz_ASC"

	// HostOrderByInputUsedCPUHzDESC captures enum value "used_cpu_hz_DESC"
	HostOrderByInputUsedCPUHzDESC HostOrderByInput = "used_cpu_hz_DESC"

	// HostOrderByInputUsedDataSpaceASC captures enum value "used_data_space_ASC"
	HostOrderByInputUsedDataSpaceASC HostOrderByInput = "used_data_space_ASC"

	// HostOrderByInputUsedDataSpaceDESC captures enum value "used_data_space_DESC"
	HostOrderByInputUsedDataSpaceDESC HostOrderByInput = "used_data_space_DESC"

	// HostOrderByInputUsedMemoryBytesASC captures enum value "used_memory_bytes_ASC"
	HostOrderByInputUsedMemoryBytesASC HostOrderByInput = "used_memory_bytes_ASC"

	// HostOrderByInputUsedMemoryBytesDESC captures enum value "used_memory_bytes_DESC"
	HostOrderByInputUsedMemoryBytesDESC HostOrderByInput = "used_memory_bytes_DESC"

	// HostOrderByInputVMNumASC captures enum value "vm_num_ASC"
	HostOrderByInputVMNumASC HostOrderByInput = "vm_num_ASC"

	// HostOrderByInputVMNumDESC captures enum value "vm_num_DESC"
	HostOrderByInputVMNumDESC HostOrderByInput = "vm_num_DESC"

	// HostOrderByInputVmotionIPASC captures enum value "vmotion_ip_ASC"
	HostOrderByInputVmotionIPASC HostOrderByInput = "vmotion_ip_ASC"

	// HostOrderByInputVmotionIPDESC captures enum value "vmotion_ip_DESC"
	HostOrderByInputVmotionIPDESC HostOrderByInput = "vmotion_ip_DESC"

	// HostOrderByInputWithFasterSsdAsCacheASC captures enum value "with_faster_ssd_as_cache_ASC"
	HostOrderByInputWithFasterSsdAsCacheASC HostOrderByInput = "with_faster_ssd_as_cache_ASC"

	// HostOrderByInputWithFasterSsdAsCacheDESC captures enum value "with_faster_ssd_as_cache_DESC"
	HostOrderByInputWithFasterSsdAsCacheDESC HostOrderByInput = "with_faster_ssd_as_cache_DESC"
)

// for schema
var hostOrderByInputEnum []interface{}

func init() {
	var res []HostOrderByInput
	if err := json.Unmarshal([]byte(`["access_ip_ASC","access_ip_DESC","allocatable_memory_bytes_ASC","allocatable_memory_bytes_DESC","chunk_id_ASC","chunk_id_DESC","cpu_brand_ASC","cpu_brand_DESC","cpu_fan_speed_unit_ASC","cpu_fan_speed_unit_DESC","cpu_hz_per_core_ASC","cpu_hz_per_core_DESC","cpu_model_ASC","cpu_model_DESC","cpu_vendor_ASC","cpu_vendor_DESC","data_ip_ASC","data_ip_DESC","failure_data_space_ASC","failure_data_space_DESC","hdd_data_capacity_ASC","hdd_data_capacity_DESC","hdd_disk_count_ASC","hdd_disk_count_DESC","hypervisor_ip_ASC","hypervisor_ip_DESC","id_ASC","id_DESC","is_os_in_raid1_ASC","is_os_in_raid1_DESC","local_id_ASC","local_id_DESC","lsm_cap_disk_safe_umount_ASC","lsm_cap_disk_safe_umount_DESC","management_ip_ASC","management_ip_DESC","model_ASC","model_DESC","name_ASC","name_DESC","nested_virtualization_ASC","nested_virtualization_DESC","nic_count_ASC","nic_count_DESC","node_topo_local_id_ASC","node_topo_local_id_DESC","os_memory_bytes_ASC","os_memory_bytes_DESC","os_version_ASC","os_version_DESC","pmem_dimm_capacity_ASC","pmem_dimm_capacity_DESC","pmem_dimm_count_ASC","pmem_dimm_count_DESC","pmem_disk_count_ASC","pmem_disk_count_DESC","provisioned_cpu_cores_ASC","provisioned_cpu_cores_DESC","provisioned_memory_bytes_ASC","provisioned_memory_bytes_DESC","running_pause_vm_memory_bytes_ASC","running_pause_vm_memory_bytes_DESC","running_vm_num_ASC","running_vm_num_DESC","scvm_cpu_ASC","scvm_cpu_DESC","scvm_memory_ASC","scvm_memory_DESC","scvm_name_ASC","scvm_name_DESC","serial_ASC","serial_DESC","ssd_data_capacity_ASC","ssd_data_capacity_DESC","ssd_disk_count_ASC","ssd_disk_count_DESC","state_ASC","state_DESC","status_ASC","status_DESC","stopped_vm_num_ASC","stopped_vm_num_DESC","suspended_vm_num_ASC","suspended_vm_num_DESC","total_cache_capacity_ASC","total_cache_capacity_DESC","total_cpu_cores_ASC","total_cpu_cores_DESC","total_cpu_hz_ASC","total_cpu_hz_DESC","total_cpu_sockets_ASC","total_cpu_sockets_DESC","total_data_capacity_ASC","total_data_capacity_DESC","total_memory_bytes_ASC","total_memory_bytes_DESC","used_cpu_hz_ASC","used_cpu_hz_DESC","used_data_space_ASC","used_data_space_DESC","used_memory_bytes_ASC","used_memory_bytes_DESC","vm_num_ASC","vm_num_DESC","vmotion_ip_ASC","vmotion_ip_DESC","with_faster_ssd_as_cache_ASC","with_faster_ssd_as_cache_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostOrderByInputEnum = append(hostOrderByInputEnum, v)
	}
}

func (m HostOrderByInput) validateHostOrderByInputEnum(path, location string, value HostOrderByInput) error {
	if err := validate.EnumCase(path, location, value, hostOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this host order by input
func (m HostOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHostOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this host order by input based on context it is used
func (m HostOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
