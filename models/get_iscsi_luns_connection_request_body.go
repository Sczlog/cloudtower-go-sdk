// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetIscsiLunsConnectionRequestBody get iscsi luns connection request body
// Example: {"after":"iscsiLunsConnection-id-string","before":"iscsiLunsConnection-id-string","first":0,"last":0,"orderBy":"allowed_initiators_ASC","skip":0,"where":{"AND":"IscsiLunWhereInput[]","NOT":"IscsiLunWhereInput[]","OR":"IscsiLunWhereInput[]","allowed_initiators":"string","allowed_initiators_contains":"string","allowed_initiators_ends_with":"string","allowed_initiators_gt":"string","allowed_initiators_gte":"string","allowed_initiators_in":["string"],"allowed_initiators_lt":"string","allowed_initiators_lte":"string","allowed_initiators_not":"string","allowed_initiators_not_contains":"string","allowed_initiators_not_ends_with":"string","allowed_initiators_not_in":["string"],"allowed_initiators_not_starts_with":"string","allowed_initiators_starts_with":"string","assigned_size":0,"assigned_size_gt":0,"assigned_size_gte":0,"assigned_size_in":[0],"assigned_size_lt":0,"assigned_size_lte":0,"assigned_size_not":0,"assigned_size_not_in":[0],"bps":0,"bps_gt":0,"bps_gte":0,"bps_in":[0],"bps_lt":0,"bps_lte":0,"bps_max":0,"bps_max_gt":0,"bps_max_gte":0,"bps_max_in":[0],"bps_max_length":0,"bps_max_length_gt":0,"bps_max_length_gte":0,"bps_max_length_in":[0],"bps_max_length_lt":0,"bps_max_length_lte":0,"bps_max_length_not":0,"bps_max_length_not_in":[0],"bps_max_lt":0,"bps_max_lte":0,"bps_max_not":0,"bps_max_not_in":[0],"bps_not":0,"bps_not_in":[0],"bps_rd":0,"bps_rd_gt":0,"bps_rd_gte":0,"bps_rd_in":[0],"bps_rd_lt":0,"bps_rd_lte":0,"bps_rd_max":0,"bps_rd_max_gt":0,"bps_rd_max_gte":0,"bps_rd_max_in":[0],"bps_rd_max_length":0,"bps_rd_max_length_gt":0,"bps_rd_max_length_gte":0,"bps_rd_max_length_in":[0],"bps_rd_max_length_lt":0,"bps_rd_max_length_lte":0,"bps_rd_max_length_not":0,"bps_rd_max_length_not_in":[0],"bps_rd_max_lt":0,"bps_rd_max_lte":0,"bps_rd_max_not":0,"bps_rd_max_not_in":[0],"bps_rd_not":0,"bps_rd_not_in":[0],"bps_wr":0,"bps_wr_gt":0,"bps_wr_gte":0,"bps_wr_in":[0],"bps_wr_lt":0,"bps_wr_lte":0,"bps_wr_max":0,"bps_wr_max_gt":0,"bps_wr_max_gte":0,"bps_wr_max_in":[0],"bps_wr_max_length":0,"bps_wr_max_length_gt":0,"bps_wr_max_length_gte":0,"bps_wr_max_length_in":[0],"bps_wr_max_length_lt":0,"bps_wr_max_length_lte":0,"bps_wr_max_length_not":0,"bps_wr_max_length_not_in":[0],"bps_wr_max_lt":0,"bps_wr_max_lte":0,"bps_wr_max_not":0,"bps_wr_max_not_in":[0],"bps_wr_not":0,"bps_wr_not_in":[0],"consistency_group":"ConsistencyGroupWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","io_size":0,"io_size_gt":0,"io_size_gte":0,"io_size_in":[0],"io_size_lt":0,"io_size_lte":0,"io_size_not":0,"io_size_not_in":[0],"iops":0,"iops_gt":0,"iops_gte":0,"iops_in":[0],"iops_lt":0,"iops_lte":0,"iops_max":0,"iops_max_gt":0,"iops_max_gte":0,"iops_max_in":[0],"iops_max_length":0,"iops_max_length_gt":0,"iops_max_length_gte":0,"iops_max_length_in":[0],"iops_max_length_lt":0,"iops_max_length_lte":0,"iops_max_length_not":0,"iops_max_length_not_in":[0],"iops_max_lt":0,"iops_max_lte":0,"iops_max_not":0,"iops_max_not_in":[0],"iops_not":0,"iops_not_in":[0],"iops_rd":0,"iops_rd_gt":0,"iops_rd_gte":0,"iops_rd_in":[0],"iops_rd_lt":0,"iops_rd_lte":0,"iops_rd_max":0,"iops_rd_max_gt":0,"iops_rd_max_gte":0,"iops_rd_max_in":[0],"iops_rd_max_length":0,"iops_rd_max_length_gt":0,"iops_rd_max_length_gte":0,"iops_rd_max_length_in":[0],"iops_rd_max_length_lt":0,"iops_rd_max_length_lte":0,"iops_rd_max_length_not":0,"iops_rd_max_length_not_in":[0],"iops_rd_max_lt":0,"iops_rd_max_lte":0,"iops_rd_max_not":0,"iops_rd_max_not_in":[0],"iops_rd_not":0,"iops_rd_not_in":[0],"iops_wr":0,"iops_wr_gt":0,"iops_wr_gte":0,"iops_wr_in":[0],"iops_wr_lt":0,"iops_wr_lte":0,"iops_wr_max":0,"iops_wr_max_gt":0,"iops_wr_max_gte":0,"iops_wr_max_in":[0],"iops_wr_max_length":0,"iops_wr_max_length_gt":0,"iops_wr_max_length_gte":0,"iops_wr_max_length_in":[0],"iops_wr_max_length_lt":0,"iops_wr_max_length_lte":0,"iops_wr_max_length_not":0,"iops_wr_max_length_not_in":[0],"iops_wr_max_lt":0,"iops_wr_max_lte":0,"iops_wr_max_not":0,"iops_wr_max_not_in":[0],"iops_wr_not":0,"iops_wr_not_in":[0],"iscsi_target":"IscsiTargetWhereInput","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","lun_id":0,"lun_id_gt":0,"lun_id_gte":0,"lun_id_in":[0],"lun_id_lt":0,"lun_id_lte":0,"lun_id_not":0,"lun_id_not_in":[0],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","replica_num":0,"replica_num_gt":0,"replica_num_gte":0,"replica_num_in":[0],"replica_num_lt":0,"replica_num_lte":0,"replica_num_not":0,"replica_num_not_in":[0],"shared_size":0,"shared_size_gt":0,"shared_size_gte":0,"shared_size_in":[0],"shared_size_lt":0,"shared_size_lte":0,"shared_size_not":0,"shared_size_not_in":[0],"snapshot_num":0,"snapshot_num_gt":0,"snapshot_num_gte":0,"snapshot_num_in":[0],"snapshot_num_lt":0,"snapshot_num_lte":0,"snapshot_num_not":0,"snapshot_num_not_in":[0],"stripe_num":0,"stripe_num_gt":0,"stripe_num_gte":0,"stripe_num_in":[0],"stripe_num_lt":0,"stripe_num_lte":0,"stripe_num_not":0,"stripe_num_not_in":[0],"stripe_size":0,"stripe_size_gt":0,"stripe_size_gte":0,"stripe_size_in":[0],"stripe_size_lt":0,"stripe_size_lte":0,"stripe_size_not":0,"stripe_size_not_in":[0],"thin_provision":false,"thin_provision_not":false,"unique_size":0,"unique_size_gt":0,"unique_size_gte":0,"unique_size_in":[0],"unique_size_lt":0,"unique_size_lte":0,"unique_size_not":0,"unique_size_not_in":[0],"zbs_volume_id":"string","zbs_volume_id_contains":"string","zbs_volume_id_ends_with":"string","zbs_volume_id_gt":"string","zbs_volume_id_gte":"string","zbs_volume_id_in":["string"],"zbs_volume_id_lt":"string","zbs_volume_id_lte":"string","zbs_volume_id_not":"string","zbs_volume_id_not_contains":"string","zbs_volume_id_not_ends_with":"string","zbs_volume_id_not_in":["string"],"zbs_volume_id_not_starts_with":"string","zbs_volume_id_starts_with":"string"}}
//
// swagger:model GetIscsiLunsConnectionRequestBody
type GetIscsiLunsConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *IscsiLunOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *IscsiLunWhereInput `json:"where,omitempty"`
}

// Validate validates this get iscsi luns connection request body
func (m *GetIscsiLunsConnectionRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIscsiLunsConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	if m.OrderBy != nil {
		if err := m.OrderBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetIscsiLunsConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get iscsi luns connection request body based on the context it is used
func (m *GetIscsiLunsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIscsiLunsConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderBy != nil {
		if err := m.OrderBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetIscsiLunsConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetIscsiLunsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIscsiLunsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetIscsiLunsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
