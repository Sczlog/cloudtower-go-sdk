// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NvmfSubsystemCommonParams nvmf subsystem common params
//
// swagger:model NvmfSubsystemCommonParams
type NvmfSubsystemCommonParams struct {

	// bps
	Bps *float64 `json:"bps,omitempty"`

	// bps max
	BpsMax *float64 `json:"bps_max,omitempty"`

	// bps max length
	BpsMaxLength *float64 `json:"bps_max_length,omitempty"`

	// bps rd
	BpsRd *float64 `json:"bps_rd,omitempty"`

	// bps rd max
	BpsRdMax *float64 `json:"bps_rd_max,omitempty"`

	// bps rd max length
	BpsRdMaxLength *float64 `json:"bps_rd_max_length,omitempty"`

	// bps wr
	BpsWr *float64 `json:"bps_wr,omitempty"`

	// bps wr max
	BpsWrMax *float64 `json:"bps_wr_max,omitempty"`

	// bps wr max length
	BpsWrMaxLength *float64 `json:"bps_wr_max_length,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// iops
	Iops *float64 `json:"iops,omitempty"`

	// iops max
	IopsMax *float64 `json:"iops_max,omitempty"`

	// iops max length
	IopsMaxLength *float64 `json:"iops_max_length,omitempty"`

	// iops rd
	IopsRd *float64 `json:"iops_rd,omitempty"`

	// iops rd max
	IopsRdMax *float64 `json:"iops_rd_max,omitempty"`

	// iops rd max length
	IopsRdMaxLength *float64 `json:"iops_rd_max_length,omitempty"`

	// iops wr
	IopsWr *float64 `json:"iops_wr,omitempty"`

	// iops wr max
	IopsWrMax *float64 `json:"iops_wr_max,omitempty"`

	// iops wr max length
	IopsWrMaxLength *float64 `json:"iops_wr_max_length,omitempty"`

	// ip whitelist
	IPWhitelist *string `json:"ip_whitelist,omitempty"`

	// nqn whitelist
	NqnWhitelist *string `json:"nqn_whitelist,omitempty"`
}

// Validate validates this nvmf subsystem common params
func (m *NvmfSubsystemCommonParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nvmf subsystem common params based on context it is used
func (m *NvmfSubsystemCommonParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NvmfSubsystemCommonParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmfSubsystemCommonParams) UnmarshalBinary(b []byte) error {
	var res NvmfSubsystemCommonParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
