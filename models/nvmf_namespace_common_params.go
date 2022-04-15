// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NvmfNamespaceCommonParams nvmf namespace common params
//
// swagger:model NvmfNamespaceCommonParams
type NvmfNamespaceCommonParams struct {

	// bps
	Bps *int64 `json:"bps,omitempty"`

	// bps max
	BpsMax *int64 `json:"bps_max,omitempty"`

	// bps max length
	BpsMaxLength *int64 `json:"bps_max_length,omitempty"`

	// bps rd
	BpsRd *int64 `json:"bps_rd,omitempty"`

	// bps rd max
	BpsRdMax *int64 `json:"bps_rd_max,omitempty"`

	// bps rd max length
	BpsRdMaxLength *int64 `json:"bps_rd_max_length,omitempty"`

	// bps wr
	BpsWr *int64 `json:"bps_wr,omitempty"`

	// bps wr max
	BpsWrMax *int64 `json:"bps_wr_max,omitempty"`

	// bps wr max length
	BpsWrMaxLength *int64 `json:"bps_wr_max_length,omitempty"`

	// iops
	Iops *int64 `json:"iops,omitempty"`

	// iops max
	IopsMax *int64 `json:"iops_max,omitempty"`

	// iops max length
	IopsMaxLength *int64 `json:"iops_max_length,omitempty"`

	// iops rd
	IopsRd *int64 `json:"iops_rd,omitempty"`

	// iops rd max
	IopsRdMax *int64 `json:"iops_rd_max,omitempty"`

	// iops rd max length
	IopsRdMaxLength *int64 `json:"iops_rd_max_length,omitempty"`

	// iops wr
	IopsWr *int64 `json:"iops_wr,omitempty"`

	// iops wr max
	IopsWrMax *int64 `json:"iops_wr_max,omitempty"`

	// iops wr max length
	IopsWrMaxLength *int64 `json:"iops_wr_max_length,omitempty"`

	// nqn whitelist
	NqnWhitelist *string `json:"nqn_whitelist,omitempty"`
}

// Validate validates this nvmf namespace common params
func (m *NvmfNamespaceCommonParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nvmf namespace common params based on context it is used
func (m *NvmfNamespaceCommonParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NvmfNamespaceCommonParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmfNamespaceCommonParams) UnmarshalBinary(b []byte) error {
	var res NvmfNamespaceCommonParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
