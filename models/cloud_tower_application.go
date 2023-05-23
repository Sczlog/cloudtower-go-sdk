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

// CloudTowerApplication cloud tower application
//
// swagger:model CloudTowerApplication
type CloudTowerApplication struct {

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// instance statuses
	// Required: true
	InstanceStatuses interface{} `json:"instanceStatuses"`

	// name
	// Required: true
	Name *string `json:"name"`

	// package
	Package *NestedCloudTowerApplicationPackage `json:"package,omitempty"`

	// placement situation
	PlacementSituation *string `json:"placementSituation,omitempty"`

	// placement verb
	PlacementVerb *string `json:"placementVerb,omitempty"`

	// resource version
	// Required: true
	ResourceVersion *int32 `json:"resourceVersion"`

	// state
	State *CloudTowerApplicationState `json:"state,omitempty"`

	// target package
	// Required: true
	TargetPackage *string `json:"targetPackage"`

	// user
	User *NestedUser `json:"user,omitempty"`

	// vm spec
	// Required: true
	VMSpec interface{} `json:"vmSpec"`
}

// Validate validates this cloud tower application
func (m *CloudTowerApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetPackage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudTowerApplication) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *CloudTowerApplication) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CloudTowerApplication) validateInstanceStatuses(formats strfmt.Registry) error {

	if m.InstanceStatuses == nil {
		return errors.Required("instanceStatuses", "body", nil)
	}

	return nil
}

func (m *CloudTowerApplication) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudTowerApplication) validatePackage(formats strfmt.Registry) error {
	if swag.IsZero(m.Package) { // not required
		return nil
	}

	if m.Package != nil {
		if err := m.Package.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("package")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("package")
			}
			return err
		}
	}

	return nil
}

func (m *CloudTowerApplication) validateResourceVersion(formats strfmt.Registry) error {

	if err := validate.Required("resourceVersion", "body", m.ResourceVersion); err != nil {
		return err
	}

	return nil
}

func (m *CloudTowerApplication) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *CloudTowerApplication) validateTargetPackage(formats strfmt.Registry) error {

	if err := validate.Required("targetPackage", "body", m.TargetPackage); err != nil {
		return err
	}

	return nil
}

func (m *CloudTowerApplication) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *CloudTowerApplication) validateVMSpec(formats strfmt.Registry) error {

	if m.VMSpec == nil {
		return errors.Required("vmSpec", "body", nil)
	}

	return nil
}

// ContextValidate validate this cloud tower application based on the context it is used
func (m *CloudTowerApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudTowerApplication) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *CloudTowerApplication) contextValidatePackage(ctx context.Context, formats strfmt.Registry) error {

	if m.Package != nil {
		if err := m.Package.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("package")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("package")
			}
			return err
		}
	}

	return nil
}

func (m *CloudTowerApplication) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *CloudTowerApplication) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudTowerApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudTowerApplication) UnmarshalBinary(b []byte) error {
	var res CloudTowerApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
