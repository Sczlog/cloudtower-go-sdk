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

// GetIscsiTargetsRequestBody get iscsi targets request body
// Example: {"after":"iscsiTargets-id-string","before":"iscsiTargets-id-string","first":0,"last":0,"orderBy":"bps_ASC","skip":0,"where":{"AND":"IscsiTargetWhereInput[]","NOT":"IscsiTargetWhereInput[]","OR":"IscsiTargetWhereInput[]","bps":0,"bps_gt":0,"bps_gte":0,"bps_in":[0],"bps_lt":0,"bps_lte":0,"bps_max":0,"bps_max_gt":0,"bps_max_gte":0,"bps_max_in":[0],"bps_max_length":0,"bps_max_length_gt":0,"bps_max_length_gte":0,"bps_max_length_in":[0],"bps_max_length_lt":0,"bps_max_length_lte":0,"bps_max_length_not":0,"bps_max_length_not_in":[0],"bps_max_lt":0,"bps_max_lte":0,"bps_max_not":0,"bps_max_not_in":[0],"bps_not":0,"bps_not_in":[0],"bps_rd":0,"bps_rd_gt":0,"bps_rd_gte":0,"bps_rd_in":[0],"bps_rd_lt":0,"bps_rd_lte":0,"bps_rd_max":0,"bps_rd_max_gt":0,"bps_rd_max_gte":0,"bps_rd_max_in":[0],"bps_rd_max_length":0,"bps_rd_max_length_gt":0,"bps_rd_max_length_gte":0,"bps_rd_max_length_in":[0],"bps_rd_max_length_lt":0,"bps_rd_max_length_lte":0,"bps_rd_max_length_not":0,"bps_rd_max_length_not_in":[0],"bps_rd_max_lt":0,"bps_rd_max_lte":0,"bps_rd_max_not":0,"bps_rd_max_not_in":[0],"bps_rd_not":0,"bps_rd_not_in":[0],"bps_wr":0,"bps_wr_gt":0,"bps_wr_gte":0,"bps_wr_in":[0],"bps_wr_lt":0,"bps_wr_lte":0,"bps_wr_max":0,"bps_wr_max_gt":0,"bps_wr_max_gte":0,"bps_wr_max_in":[0],"bps_wr_max_length":0,"bps_wr_max_length_gt":0,"bps_wr_max_length_gte":0,"bps_wr_max_length_in":[0],"bps_wr_max_length_lt":0,"bps_wr_max_length_lte":0,"bps_wr_max_length_not":0,"bps_wr_max_length_not_in":[0],"bps_wr_max_lt":0,"bps_wr_max_lte":0,"bps_wr_max_not":0,"bps_wr_max_not_in":[0],"bps_wr_not":0,"bps_wr_not_in":[0],"chap_enabled":false,"chap_enabled_not":false,"chap_name":"string","chap_name_contains":"string","chap_name_ends_with":"string","chap_name_gt":"string","chap_name_gte":"string","chap_name_in":["string"],"chap_name_lt":"string","chap_name_lte":"string","chap_name_not":"string","chap_name_not_contains":"string","chap_name_not_ends_with":"string","chap_name_not_in":["string"],"chap_name_not_starts_with":"string","chap_name_starts_with":"string","chap_secret":"string","chap_secret_contains":"string","chap_secret_ends_with":"string","chap_secret_gt":"string","chap_secret_gte":"string","chap_secret_in":["string"],"chap_secret_lt":"string","chap_secret_lte":"string","chap_secret_not":"string","chap_secret_not_contains":"string","chap_secret_not_ends_with":"string","chap_secret_not_in":["string"],"chap_secret_not_starts_with":"string","chap_secret_starts_with":"string","cluster":"ClusterWhereInput","description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"external_use":false,"external_use_not":false,"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"io_size":0,"io_size_gt":0,"io_size_gte":0,"io_size_in":[0],"io_size_lt":0,"io_size_lte":0,"io_size_not":0,"io_size_not_in":[0],"iops":0,"iops_gt":0,"iops_gte":0,"iops_in":[0],"iops_lt":0,"iops_lte":0,"iops_max":0,"iops_max_gt":0,"iops_max_gte":0,"iops_max_in":[0],"iops_max_length":0,"iops_max_length_gt":0,"iops_max_length_gte":0,"iops_max_length_in":[0],"iops_max_length_lt":0,"iops_max_length_lte":0,"iops_max_length_not":0,"iops_max_length_not_in":[0],"iops_max_lt":0,"iops_max_lte":0,"iops_max_not":0,"iops_max_not_in":[0],"iops_not":0,"iops_not_in":[0],"iops_rd":0,"iops_rd_gt":0,"iops_rd_gte":0,"iops_rd_in":[0],"iops_rd_lt":0,"iops_rd_lte":0,"iops_rd_max":0,"iops_rd_max_gt":0,"iops_rd_max_gte":0,"iops_rd_max_in":[0],"iops_rd_max_length":0,"iops_rd_max_length_gt":0,"iops_rd_max_length_gte":0,"iops_rd_max_length_in":[0],"iops_rd_max_length_lt":0,"iops_rd_max_length_lte":0,"iops_rd_max_length_not":0,"iops_rd_max_length_not_in":[0],"iops_rd_max_lt":0,"iops_rd_max_lte":0,"iops_rd_max_not":0,"iops_rd_max_not_in":[0],"iops_rd_not":0,"iops_rd_not_in":[0],"iops_wr":0,"iops_wr_gt":0,"iops_wr_gte":0,"iops_wr_in":[0],"iops_wr_lt":0,"iops_wr_lte":0,"iops_wr_max":0,"iops_wr_max_gt":0,"iops_wr_max_gte":0,"iops_wr_max_in":[0],"iops_wr_max_length":0,"iops_wr_max_length_gt":0,"iops_wr_max_length_gte":0,"iops_wr_max_length_in":[0],"iops_wr_max_length_lt":0,"iops_wr_max_length_lte":0,"iops_wr_max_length_not":0,"iops_wr_max_length_not_in":[0],"iops_wr_max_lt":0,"iops_wr_max_lte":0,"iops_wr_max_not":0,"iops_wr_max_not_in":[0],"iops_wr_not":0,"iops_wr_not_in":[0],"ip_whitelist":"string","ip_whitelist_contains":"string","ip_whitelist_ends_with":"string","ip_whitelist_gt":"string","ip_whitelist_gte":"string","ip_whitelist_in":["string"],"ip_whitelist_lt":"string","ip_whitelist_lte":"string","ip_whitelist_not":"string","ip_whitelist_not_contains":"string","ip_whitelist_not_ends_with":"string","ip_whitelist_not_in":["string"],"ip_whitelist_not_starts_with":"string","ip_whitelist_starts_with":"string","iqn_name":"string","iqn_name_contains":"string","iqn_name_ends_with":"string","iqn_name_gt":"string","iqn_name_gte":"string","iqn_name_in":["string"],"iqn_name_lt":"string","iqn_name_lte":"string","iqn_name_not":"string","iqn_name_not_contains":"string","iqn_name_not_ends_with":"string","iqn_name_not_in":["string"],"iqn_name_not_starts_with":"string","iqn_name_starts_with":"string","iqn_whitelist":"string","iqn_whitelist_contains":"string","iqn_whitelist_ends_with":"string","iqn_whitelist_gt":"string","iqn_whitelist_gte":"string","iqn_whitelist_in":["string"],"iqn_whitelist_lt":"string","iqn_whitelist_lte":"string","iqn_whitelist_not":"string","iqn_whitelist_not_contains":"string","iqn_whitelist_not_ends_with":"string","iqn_whitelist_not_in":["string"],"iqn_whitelist_not_starts_with":"string","iqn_whitelist_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","luns_every":"IscsiLunWhereInput","luns_none":"IscsiLunWhereInput","luns_some":"IscsiLunWhereInput","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","replica_num":0,"replica_num_gt":0,"replica_num_gte":0,"replica_num_in":[0],"replica_num_lt":0,"replica_num_lte":0,"replica_num_not":0,"replica_num_not_in":[0],"stripe_num":0,"stripe_num_gt":0,"stripe_num_gte":0,"stripe_num_in":[0],"stripe_num_lt":0,"stripe_num_lte":0,"stripe_num_not":0,"stripe_num_not_in":[0],"stripe_size":0,"stripe_size_gt":0,"stripe_size_gte":0,"stripe_size_in":[0],"stripe_size_lt":0,"stripe_size_lte":0,"stripe_size_not":0,"stripe_size_not_in":[0],"thin_provision":false,"thin_provision_not":false}}
//
// swagger:model GetIscsiTargetsRequestBody
type GetIscsiTargetsRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *IscsiTargetOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *IscsiTargetWhereInput `json:"where,omitempty"`
}

// Validate validates this get iscsi targets request body
func (m *GetIscsiTargetsRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetIscsiTargetsRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetIscsiTargetsRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get iscsi targets request body based on the context it is used
func (m *GetIscsiTargetsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetIscsiTargetsRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetIscsiTargetsRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetIscsiTargetsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIscsiTargetsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetIscsiTargetsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
