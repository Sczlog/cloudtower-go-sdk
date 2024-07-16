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

// VirtualPrivateCloudWhereInput virtual private cloud where input
//
// swagger:model VirtualPrivateCloudWhereInput
type VirtualPrivateCloudWhereInput struct {

	// a n d
	AND []*VirtualPrivateCloudWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VirtualPrivateCloudWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VirtualPrivateCloudWhereInput `json:"OR,omitempty"`

	// associate external subnet num
	AssociateExternalSubnetNum *int32 `json:"associate_external_subnet_num,omitempty"`

	// associate external subnet num gt
	AssociateExternalSubnetNumGt *int32 `json:"associate_external_subnet_num_gt,omitempty"`

	// associate external subnet num gte
	AssociateExternalSubnetNumGte *int32 `json:"associate_external_subnet_num_gte,omitempty"`

	// associate external subnet num in
	AssociateExternalSubnetNumIn []int32 `json:"associate_external_subnet_num_in,omitempty"`

	// associate external subnet num lt
	AssociateExternalSubnetNumLt *int32 `json:"associate_external_subnet_num_lt,omitempty"`

	// associate external subnet num lte
	AssociateExternalSubnetNumLte *int32 `json:"associate_external_subnet_num_lte,omitempty"`

	// associate external subnet num not
	AssociateExternalSubnetNumNot *int32 `json:"associate_external_subnet_num_not,omitempty"`

	// associate external subnet num not in
	AssociateExternalSubnetNumNotIn []int32 `json:"associate_external_subnet_num_not_in,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

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

	// isolation policies every
	IsolationPoliciesEvery *VirtualPrivateCloudIsolationPolicyWhereInput `json:"isolation_policies_every,omitempty"`

	// isolation policies none
	IsolationPoliciesNone *VirtualPrivateCloudIsolationPolicyWhereInput `json:"isolation_policies_none,omitempty"`

	// isolation policies some
	IsolationPoliciesSome *VirtualPrivateCloudIsolationPolicyWhereInput `json:"isolation_policies_some,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// mtu
	Mtu *int32 `json:"mtu,omitempty"`

	// mtu gt
	MtuGt *int32 `json:"mtu_gt,omitempty"`

	// mtu gte
	MtuGte *int32 `json:"mtu_gte,omitempty"`

	// mtu in
	MtuIn []int32 `json:"mtu_in,omitempty"`

	// mtu lt
	MtuLt *int32 `json:"mtu_lt,omitempty"`

	// mtu lte
	MtuLte *int32 `json:"mtu_lte,omitempty"`

	// mtu not
	MtuNot *int32 `json:"mtu_not,omitempty"`

	// mtu not in
	MtuNotIn []int32 `json:"mtu_not_in,omitempty"`

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

	// route tables every
	RouteTablesEvery *VirtualPrivateCloudRouteTableWhereInput `json:"route_tables_every,omitempty"`

	// route tables none
	RouteTablesNone *VirtualPrivateCloudRouteTableWhereInput `json:"route_tables_none,omitempty"`

	// route tables some
	RouteTablesSome *VirtualPrivateCloudRouteTableWhereInput `json:"route_tables_some,omitempty"`

	// security groups every
	SecurityGroupsEvery *VirtualPrivateCloudSecurityGroupWhereInput `json:"security_groups_every,omitempty"`

	// security groups none
	SecurityGroupsNone *VirtualPrivateCloudSecurityGroupWhereInput `json:"security_groups_none,omitempty"`

	// security groups some
	SecurityGroupsSome *VirtualPrivateCloudSecurityGroupWhereInput `json:"security_groups_some,omitempty"`

	// security policies every
	SecurityPoliciesEvery *VirtualPrivateCloudSecurityPolicyWhereInput `json:"security_policies_every,omitempty"`

	// security policies none
	SecurityPoliciesNone *VirtualPrivateCloudSecurityPolicyWhereInput `json:"security_policies_none,omitempty"`

	// security policies some
	SecurityPoliciesSome *VirtualPrivateCloudSecurityPolicyWhereInput `json:"security_policies_some,omitempty"`

	// subnets every
	SubnetsEvery *VirtualPrivateCloudSubnetWhereInput `json:"subnets_every,omitempty"`

	// subnets none
	SubnetsNone *VirtualPrivateCloudSubnetWhereInput `json:"subnets_none,omitempty"`

	// subnets some
	SubnetsSome *VirtualPrivateCloudSubnetWhereInput `json:"subnets_some,omitempty"`
}

// Validate validates this virtual private cloud where input
func (m *VirtualPrivateCloudWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsolationPoliciesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsolationPoliciesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsolationPoliciesSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteTablesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteTablesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteTablesSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityPoliciesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityPoliciesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityPoliciesSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetsSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateIsolationPoliciesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.IsolationPoliciesEvery) { // not required
		return nil
	}

	if m.IsolationPoliciesEvery != nil {
		if err := m.IsolationPoliciesEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isolation_policies_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isolation_policies_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateIsolationPoliciesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.IsolationPoliciesNone) { // not required
		return nil
	}

	if m.IsolationPoliciesNone != nil {
		if err := m.IsolationPoliciesNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isolation_policies_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isolation_policies_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateIsolationPoliciesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.IsolationPoliciesSome) { // not required
		return nil
	}

	if m.IsolationPoliciesSome != nil {
		if err := m.IsolationPoliciesSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isolation_policies_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isolation_policies_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateRouteTablesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteTablesEvery) { // not required
		return nil
	}

	if m.RouteTablesEvery != nil {
		if err := m.RouteTablesEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_tables_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_tables_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateRouteTablesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteTablesNone) { // not required
		return nil
	}

	if m.RouteTablesNone != nil {
		if err := m.RouteTablesNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_tables_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_tables_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateRouteTablesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteTablesSome) { // not required
		return nil
	}

	if m.RouteTablesSome != nil {
		if err := m.RouteTablesSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_tables_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_tables_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSecurityGroupsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupsEvery) { // not required
		return nil
	}

	if m.SecurityGroupsEvery != nil {
		if err := m.SecurityGroupsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSecurityGroupsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupsNone) { // not required
		return nil
	}

	if m.SecurityGroupsNone != nil {
		if err := m.SecurityGroupsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSecurityGroupsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupsSome) { // not required
		return nil
	}

	if m.SecurityGroupsSome != nil {
		if err := m.SecurityGroupsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSecurityPoliciesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityPoliciesEvery) { // not required
		return nil
	}

	if m.SecurityPoliciesEvery != nil {
		if err := m.SecurityPoliciesEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_policies_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_policies_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSecurityPoliciesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityPoliciesNone) { // not required
		return nil
	}

	if m.SecurityPoliciesNone != nil {
		if err := m.SecurityPoliciesNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_policies_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_policies_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSecurityPoliciesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityPoliciesSome) { // not required
		return nil
	}

	if m.SecurityPoliciesSome != nil {
		if err := m.SecurityPoliciesSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_policies_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_policies_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSubnetsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetsEvery) { // not required
		return nil
	}

	if m.SubnetsEvery != nil {
		if err := m.SubnetsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnets_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnets_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSubnetsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetsNone) { // not required
		return nil
	}

	if m.SubnetsNone != nil {
		if err := m.SubnetsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnets_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnets_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) validateSubnetsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetsSome) { // not required
		return nil
	}

	if m.SubnetsSome != nil {
		if err := m.SubnetsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnets_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnets_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual private cloud where input based on the context it is used
func (m *VirtualPrivateCloudWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsolationPoliciesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsolationPoliciesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsolationPoliciesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteTablesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteTablesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteTablesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroupsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroupsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroupsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityPoliciesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityPoliciesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityPoliciesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateIsolationPoliciesEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.IsolationPoliciesEvery != nil {
		if err := m.IsolationPoliciesEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isolation_policies_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isolation_policies_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateIsolationPoliciesNone(ctx context.Context, formats strfmt.Registry) error {

	if m.IsolationPoliciesNone != nil {
		if err := m.IsolationPoliciesNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isolation_policies_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isolation_policies_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateIsolationPoliciesSome(ctx context.Context, formats strfmt.Registry) error {

	if m.IsolationPoliciesSome != nil {
		if err := m.IsolationPoliciesSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isolation_policies_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isolation_policies_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateRouteTablesEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.RouteTablesEvery != nil {
		if err := m.RouteTablesEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_tables_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_tables_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateRouteTablesNone(ctx context.Context, formats strfmt.Registry) error {

	if m.RouteTablesNone != nil {
		if err := m.RouteTablesNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_tables_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_tables_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateRouteTablesSome(ctx context.Context, formats strfmt.Registry) error {

	if m.RouteTablesSome != nil {
		if err := m.RouteTablesSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_tables_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_tables_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSecurityGroupsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroupsEvery != nil {
		if err := m.SecurityGroupsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSecurityGroupsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroupsNone != nil {
		if err := m.SecurityGroupsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSecurityGroupsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroupsSome != nil {
		if err := m.SecurityGroupsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSecurityPoliciesEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityPoliciesEvery != nil {
		if err := m.SecurityPoliciesEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_policies_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_policies_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSecurityPoliciesNone(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityPoliciesNone != nil {
		if err := m.SecurityPoliciesNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_policies_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_policies_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSecurityPoliciesSome(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityPoliciesSome != nil {
		if err := m.SecurityPoliciesSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_policies_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_policies_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSubnetsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.SubnetsEvery != nil {
		if err := m.SubnetsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnets_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnets_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSubnetsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.SubnetsNone != nil {
		if err := m.SubnetsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnets_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnets_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudWhereInput) contextValidateSubnetsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.SubnetsSome != nil {
		if err := m.SubnetsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnets_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnets_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudWhereInput) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
