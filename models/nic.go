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
	"github.com/go-openapi/validate"
)

// Nic nic
//
// swagger:model Nic
type Nic struct {

	// driver
	Driver *string `json:"driver,omitempty"`

	// driver state
	DriverState *NicDriverState `json:"driver_state,omitempty"`

	// gateway ip
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// host
	// Required: true
	Host *NestedHost `json:"host"`

	// ibdev
	Ibdev *string `json:"ibdev,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// is sriov
	IsSriov *bool `json:"is_sriov,omitempty"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mac address
	// Required: true
	MacAddress *string `json:"mac_address"`

	// max vf num
	MaxVfNum *int32 `json:"max_vf_num,omitempty"`

	// model
	Model *string `json:"model,omitempty"`

	// mtu
	// Required: true
	Mtu *int32 `json:"mtu"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nic uuid
	NicUUID *string `json:"nic_uuid,omitempty"`

	// physical
	// Required: true
	Physical *bool `json:"physical"`

	// rdma enabled
	RdmaEnabled *bool `json:"rdma_enabled,omitempty"`

	// running
	// Required: true
	Running *bool `json:"running"`

	// speed
	Speed *float64 `json:"speed,omitempty"`

	// subnet mask
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// total vf num
	TotalVfNum *int32 `json:"total_vf_num,omitempty"`

	// type
	Type *NetworkType `json:"type,omitempty"`

	// up
	// Required: true
	Up *bool `json:"up"`

	// used vf num
	UsedVfNum *int32 `json:"used_vf_num,omitempty"`

	// vds
	Vds *NestedVds `json:"vds,omitempty"`
}

// Validate validates this nic
func (m *Nic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriverState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Nic) validateDriverState(formats strfmt.Registry) error {
	if swag.IsZero(m.DriverState) { // not required
		return nil
	}

	if m.DriverState != nil {
		if err := m.DriverState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver_state")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Nic) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateMacAddress(formats strfmt.Registry) error {

	if err := validate.Required("mac_address", "body", m.MacAddress); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateMtu(formats strfmt.Registry) error {

	if err := validate.Required("mtu", "body", m.Mtu); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validatePhysical(formats strfmt.Registry) error {

	if err := validate.Required("physical", "body", m.Physical); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateRunning(formats strfmt.Registry) error {

	if err := validate.Required("running", "body", m.Running); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateUp(formats strfmt.Registry) error {

	if err := validate.Required("up", "body", m.Up); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateVds(formats strfmt.Registry) error {
	if swag.IsZero(m.Vds) { // not required
		return nil
	}

	if m.Vds != nil {
		if err := m.Vds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nic based on the context it is used
func (m *Nic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDriverState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Nic) contextValidateDriverState(ctx context.Context, formats strfmt.Registry) error {

	if m.DriverState != nil {
		if err := m.DriverState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver_state")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Nic) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateVds(ctx context.Context, formats strfmt.Registry) error {

	if m.Vds != nil {
		if err := m.Vds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Nic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Nic) UnmarshalBinary(b []byte) error {
	var res Nic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
