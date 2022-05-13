// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatacenterWhereInput datacenter where input
//
// swagger:model DatacenterWhereInput
type DatacenterWhereInput struct {

	// a n d
	AND []*DatacenterWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*DatacenterWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*DatacenterWhereInput `json:"OR,omitempty"`

	// cluster num
	ClusterNum *int32 `json:"cluster_num,omitempty"`

	// cluster num gt
	ClusterNumGt *int32 `json:"cluster_num_gt,omitempty"`

	// cluster num gte
	ClusterNumGte *int32 `json:"cluster_num_gte,omitempty"`

	// cluster num in
	ClusterNumIn []int32 `json:"cluster_num_in,omitempty"`

	// cluster num lt
	ClusterNumLt *int32 `json:"cluster_num_lt,omitempty"`

	// cluster num lte
	ClusterNumLte *int32 `json:"cluster_num_lte,omitempty"`

	// cluster num not
	ClusterNumNot *int32 `json:"cluster_num_not,omitempty"`

	// cluster num not in
	ClusterNumNotIn []int32 `json:"cluster_num_not_in,omitempty"`

	// clusters every
	ClustersEvery *ClusterWhereInput `json:"clusters_every,omitempty"`

	// clusters none
	ClustersNone *ClusterWhereInput `json:"clusters_none,omitempty"`

	// clusters some
	ClustersSome *ClusterWhereInput `json:"clusters_some,omitempty"`

	// failure data space
	FailureDataSpace *int64 `json:"failure_data_space,omitempty"`

	// failure data space gt
	FailureDataSpaceGt *int64 `json:"failure_data_space_gt,omitempty"`

	// failure data space gte
	FailureDataSpaceGte *int64 `json:"failure_data_space_gte,omitempty"`

	// failure data space in
	FailureDataSpaceIn []int64 `json:"failure_data_space_in,omitempty"`

	// failure data space lt
	FailureDataSpaceLt *int64 `json:"failure_data_space_lt,omitempty"`

	// failure data space lte
	FailureDataSpaceLte *int64 `json:"failure_data_space_lte,omitempty"`

	// failure data space not
	FailureDataSpaceNot *int64 `json:"failure_data_space_not,omitempty"`

	// failure data space not in
	FailureDataSpaceNotIn []int64 `json:"failure_data_space_not_in,omitempty"`

	// host num
	HostNum *int32 `json:"host_num,omitempty"`

	// host num gt
	HostNumGt *int32 `json:"host_num_gt,omitempty"`

	// host num gte
	HostNumGte *int32 `json:"host_num_gte,omitempty"`

	// host num in
	HostNumIn []int32 `json:"host_num_in,omitempty"`

	// host num lt
	HostNumLt *int32 `json:"host_num_lt,omitempty"`

	// host num lte
	HostNumLte *int32 `json:"host_num_lte,omitempty"`

	// host num not
	HostNumNot *int32 `json:"host_num_not,omitempty"`

	// host num not in
	HostNumNotIn []int32 `json:"host_num_not_in,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// labels every
	LabelsEvery *LabelWhereInput `json:"labels_every,omitempty"`

	// labels none
	LabelsNone *LabelWhereInput `json:"labels_none,omitempty"`

	// labels some
	LabelsSome *LabelWhereInput `json:"labels_some,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// organization
	Organization *OrganizationWhereInput `json:"organization,omitempty"`

	// total cpu hz
	TotalCPUHz *int64 `json:"total_cpu_hz,omitempty"`

	// total cpu hz gt
	TotalCPUHzGt *int64 `json:"total_cpu_hz_gt,omitempty"`

	// total cpu hz gte
	TotalCPUHzGte *int64 `json:"total_cpu_hz_gte,omitempty"`

	// total cpu hz in
	TotalCPUHzIn []int64 `json:"total_cpu_hz_in,omitempty"`

	// total cpu hz lt
	TotalCPUHzLt *int64 `json:"total_cpu_hz_lt,omitempty"`

	// total cpu hz lte
	TotalCPUHzLte *int64 `json:"total_cpu_hz_lte,omitempty"`

	// total cpu hz not
	TotalCPUHzNot *int64 `json:"total_cpu_hz_not,omitempty"`

	// total cpu hz not in
	TotalCPUHzNotIn []int64 `json:"total_cpu_hz_not_in,omitempty"`

	// total data capacity
	TotalDataCapacity *int64 `json:"total_data_capacity,omitempty"`

	// total data capacity gt
	TotalDataCapacityGt *int64 `json:"total_data_capacity_gt,omitempty"`

	// total data capacity gte
	TotalDataCapacityGte *int64 `json:"total_data_capacity_gte,omitempty"`

	// total data capacity in
	TotalDataCapacityIn []int64 `json:"total_data_capacity_in,omitempty"`

	// total data capacity lt
	TotalDataCapacityLt *int64 `json:"total_data_capacity_lt,omitempty"`

	// total data capacity lte
	TotalDataCapacityLte *int64 `json:"total_data_capacity_lte,omitempty"`

	// total data capacity not
	TotalDataCapacityNot *int64 `json:"total_data_capacity_not,omitempty"`

	// total data capacity not in
	TotalDataCapacityNotIn []int64 `json:"total_data_capacity_not_in,omitempty"`

	// total memory bytes
	TotalMemoryBytes *int64 `json:"total_memory_bytes,omitempty"`

	// total memory bytes gt
	TotalMemoryBytesGt *int64 `json:"total_memory_bytes_gt,omitempty"`

	// total memory bytes gte
	TotalMemoryBytesGte *int64 `json:"total_memory_bytes_gte,omitempty"`

	// total memory bytes in
	TotalMemoryBytesIn []int64 `json:"total_memory_bytes_in,omitempty"`

	// total memory bytes lt
	TotalMemoryBytesLt *int64 `json:"total_memory_bytes_lt,omitempty"`

	// total memory bytes lte
	TotalMemoryBytesLte *int64 `json:"total_memory_bytes_lte,omitempty"`

	// total memory bytes not
	TotalMemoryBytesNot *int64 `json:"total_memory_bytes_not,omitempty"`

	// total memory bytes not in
	TotalMemoryBytesNotIn []int64 `json:"total_memory_bytes_not_in,omitempty"`

	// used cpu hz
	UsedCPUHz *int64 `json:"used_cpu_hz,omitempty"`

	// used cpu hz gt
	UsedCPUHzGt *int64 `json:"used_cpu_hz_gt,omitempty"`

	// used cpu hz gte
	UsedCPUHzGte *int64 `json:"used_cpu_hz_gte,omitempty"`

	// used cpu hz in
	UsedCPUHzIn []int64 `json:"used_cpu_hz_in,omitempty"`

	// used cpu hz lt
	UsedCPUHzLt *int64 `json:"used_cpu_hz_lt,omitempty"`

	// used cpu hz lte
	UsedCPUHzLte *int64 `json:"used_cpu_hz_lte,omitempty"`

	// used cpu hz not
	UsedCPUHzNot *int64 `json:"used_cpu_hz_not,omitempty"`

	// used cpu hz not in
	UsedCPUHzNotIn []int64 `json:"used_cpu_hz_not_in,omitempty"`

	// used data space
	UsedDataSpace *int64 `json:"used_data_space,omitempty"`

	// used data space gt
	UsedDataSpaceGt *int64 `json:"used_data_space_gt,omitempty"`

	// used data space gte
	UsedDataSpaceGte *int64 `json:"used_data_space_gte,omitempty"`

	// used data space in
	UsedDataSpaceIn []int64 `json:"used_data_space_in,omitempty"`

	// used data space lt
	UsedDataSpaceLt *int64 `json:"used_data_space_lt,omitempty"`

	// used data space lte
	UsedDataSpaceLte *int64 `json:"used_data_space_lte,omitempty"`

	// used data space not
	UsedDataSpaceNot *int64 `json:"used_data_space_not,omitempty"`

	// used data space not in
	UsedDataSpaceNotIn []int64 `json:"used_data_space_not_in,omitempty"`

	// used memory bytes
	UsedMemoryBytes *int64 `json:"used_memory_bytes,omitempty"`

	// used memory bytes gt
	UsedMemoryBytesGt *int64 `json:"used_memory_bytes_gt,omitempty"`

	// used memory bytes gte
	UsedMemoryBytesGte *int64 `json:"used_memory_bytes_gte,omitempty"`

	// used memory bytes in
	UsedMemoryBytesIn []int64 `json:"used_memory_bytes_in,omitempty"`

	// used memory bytes lt
	UsedMemoryBytesLt *int64 `json:"used_memory_bytes_lt,omitempty"`

	// used memory bytes lte
	UsedMemoryBytesLte *int64 `json:"used_memory_bytes_lte,omitempty"`

	// used memory bytes not
	UsedMemoryBytesNot *int64 `json:"used_memory_bytes_not,omitempty"`

	// used memory bytes not in
	UsedMemoryBytesNotIn []int64 `json:"used_memory_bytes_not_in,omitempty"`

	// vm num
	VMNum *int32 `json:"vm_num,omitempty"`

	// vm num gt
	VMNumGt *int32 `json:"vm_num_gt,omitempty"`

	// vm num gte
	VMNumGte *int32 `json:"vm_num_gte,omitempty"`

	// vm num in
	VMNumIn []int32 `json:"vm_num_in,omitempty"`

	// vm num lt
	VMNumLt *int32 `json:"vm_num_lt,omitempty"`

	// vm num lte
	VMNumLte *int32 `json:"vm_num_lte,omitempty"`

	// vm num not
	VMNumNot *int32 `json:"vm_num_not,omitempty"`

	// vm num not in
	VMNumNotIn []int32 `json:"vm_num_not_in,omitempty"`
}

// Validate validates this datacenter where input
func (m *DatacenterWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterWhereInput) validateClustersEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersEvery) { // not required
		return nil
	}

	if m.ClustersEvery != nil {
		if err := m.ClustersEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_every")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) validateClustersNone(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersNone) { // not required
		return nil
	}

	if m.ClustersNone != nil {
		if err := m.ClustersNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_none")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) validateClustersSome(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersSome) { // not required
		return nil
	}

	if m.ClustersSome != nil {
		if err := m.ClustersSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_some")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) validateLabelsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsEvery) { // not required
		return nil
	}

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) validateLabelsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsNone) { // not required
		return nil
	}

	if m.LabelsNone != nil {
		if err := m.LabelsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) validateLabelsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsSome) { // not required
		return nil
	}

	if m.LabelsSome != nil {
		if err := m.LabelsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) validateOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter where input based on the context it is used
func (m *DatacenterWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateClustersEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersEvery != nil {
		if err := m.ClustersEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_every")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateClustersNone(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersNone != nil {
		if err := m.ClustersNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_none")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateClustersSome(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersSome != nil {
		if err := m.ClustersSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_some")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateLabelsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateLabelsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsNone != nil {
		if err := m.LabelsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateLabelsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsSome != nil {
		if err := m.LabelsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterWhereInput) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.Organization != nil {
		if err := m.Organization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterWhereInput) UnmarshalBinary(b []byte) error {
	var res DatacenterWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
