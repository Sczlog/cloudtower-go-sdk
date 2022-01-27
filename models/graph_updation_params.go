// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GraphUpdationParams graph updation params
//
// swagger:model GraphUpdationParams
type GraphUpdationParams struct {

	// data
	Data *GraphUpdationParamsData `json:"data,omitempty"`

	// where
	// Required: true
	Where *GraphWhereInput `json:"where"`
}

// Validate validates this graph updation params
func (m *GraphUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
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

func (m *GraphUpdationParams) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
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

// ContextValidate validate this graph updation params based on the context it is used
func (m *GraphUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
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

func (m *GraphUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GraphUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphUpdationParams) UnmarshalBinary(b []byte) error {
	var res GraphUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GraphUpdationParamsData graph updation params data
//
// swagger:model GraphUpdationParamsData
type GraphUpdationParamsData struct {

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// connect id
	ConnectID []string `json:"connect_id,omitempty"`

	// disks
	Disks *DiskWhereInput `json:"disks,omitempty"`

	// hosts
	Hosts *HostWhereInput `json:"hosts,omitempty"`

	// instance ids
	InstanceIds []string `json:"instance_ids,omitempty"`

	// luns
	Luns *IscsiLunWhereInput `json:"luns,omitempty"`

	// metric count
	MetricCount *int32 `json:"metric_count,omitempty"`

	// metric name
	MetricName *string `json:"metric_name,omitempty"`

	// metric type
	MetricType MetricType `json:"metric_type,omitempty"`

	// network
	Network NetworkType `json:"network,omitempty"`

	// nics
	Nics *NicWhereInput `json:"nics,omitempty"`

	// resource type
	ResourceType *string `json:"resource_type,omitempty"`

	// service
	Service *string `json:"service,omitempty"`

	// title
	Title *string `json:"title,omitempty"`

	// type
	Type GraphType `json:"type,omitempty"`

	// vm nics
	VMNics *VMNicWhereInput `json:"vmNics,omitempty"`

	// vm volumes
	VMVolumes *VMVolumeWhereInput `json:"vmVolumes,omitempty"`

	// vms
	Vms *VMWhereInput `json:"vms,omitempty"`
}

// Validate validates this graph updation params data
func (m *GraphUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphUpdationParamsData) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	if m.Disks != nil {
		if err := m.Disks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "disks")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	if m.Hosts != nil {
		if err := m.Hosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "hosts")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) validateLuns(formats strfmt.Registry) error {
	if swag.IsZero(m.Luns) { // not required
		return nil
	}

	if m.Luns != nil {
		if err := m.Luns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "luns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "luns")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) validateMetricType(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricType) { // not required
		return nil
	}

	if err := m.MetricType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "metric_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data" + "." + "metric_type")
		}
		return err
	}

	return nil
}

func (m *GraphUpdationParamsData) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if err := m.Network.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "network")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data" + "." + "network")
		}
		return err
	}

	return nil
}

func (m *GraphUpdationParamsData) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	if m.Nics != nil {
		if err := m.Nics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "nics")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data" + "." + "type")
		}
		return err
	}

	return nil
}

func (m *GraphUpdationParamsData) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	if m.VMNics != nil {
		if err := m.VMNics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vmNics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vmNics")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) validateVMVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.VMVolumes) { // not required
		return nil
	}

	if m.VMVolumes != nil {
		if err := m.VMVolumes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vmVolumes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vmVolumes")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	if m.Vms != nil {
		if err := m.Vms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this graph updation params data based on the context it is used
func (m *GraphUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetricType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphUpdationParamsData) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.Disks != nil {
		if err := m.Disks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "disks")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	if m.Hosts != nil {
		if err := m.Hosts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "hosts")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateLuns(ctx context.Context, formats strfmt.Registry) error {

	if m.Luns != nil {
		if err := m.Luns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "luns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "luns")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateMetricType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MetricType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "metric_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data" + "." + "metric_type")
		}
		return err
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Network.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "network")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data" + "." + "network")
		}
		return err
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	if m.Nics != nil {
		if err := m.Nics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "nics")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data" + "." + "type")
		}
		return err
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	if m.VMNics != nil {
		if err := m.VMNics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vmNics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vmNics")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateVMVolumes(ctx context.Context, formats strfmt.Registry) error {

	if m.VMVolumes != nil {
		if err := m.VMVolumes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vmVolumes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vmVolumes")
			}
			return err
		}
	}

	return nil
}

func (m *GraphUpdationParamsData) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	if m.Vms != nil {
		if err := m.Vms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GraphUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res GraphUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
