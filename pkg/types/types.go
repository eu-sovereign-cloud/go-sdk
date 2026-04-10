package types

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

type ResourceType interface {
	schema.Region |
		schema.Role |
		schema.RoleAssignment |
		schema.Workspace |
		schema.BlockStorage |
		schema.StorageSku |
		schema.Image |
		schema.Instance |
		schema.InstanceSku |
		schema.Network |
		schema.NetworkSku |
		schema.InternetGateway |
		schema.RouteTable |
		schema.Subnet |
		schema.PublicIp |
		schema.Nic |
		schema.SecurityGroupRule |
		schema.SecurityGroup
}

type MetadataType interface {
	schema.GlobalResourceMetadata |
		schema.GlobalTenantResourceMetadata |
		schema.SkuResourceMetadata |
		schema.RegionalResourceMetadata |
		schema.RegionalWorkspaceResourceMetadata |
		schema.RegionalNetworkResourceMetadata
}

type SpecType interface {
	schema.RegionSpec |
		schema.RoleSpec |
		schema.RoleAssignmentSpec |
		schema.WorkspaceSpec |
		schema.BlockStorageSpec |
		schema.StorageSkuSpec |
		schema.ImageSpec |
		schema.InstanceSpec |
		schema.InstanceSkuSpec |
		schema.NetworkSpec |
		schema.NetworkSkuSpec |
		schema.InternetGatewaySpec |
		schema.RouteTableSpec |
		schema.SubnetSpec |
		schema.PublicIpSpec |
		schema.NicSpec |
		schema.SecurityGroupRuleSpec |
		schema.SecurityGroupSpec
}

type StatusType interface {
	schema.Status |
		schema.WorkspaceStatus |
		schema.BlockStorageStatus |
		schema.ImageStatus |
		schema.InstanceStatus |
		schema.NetworkStatus |
		schema.SubnetStatus |
		schema.RouteTableStatus |
		schema.NicStatus |
		schema.PublicIpStatus |
		schema.SecurityGroupStatus
}

func GetStatusState[S StatusType](status *S) schema.ResourceState {
	if status == nil {
		return ""
	}

	switch v := any(*status).(type) {
	case schema.Status:
		return v.State
	case schema.WorkspaceStatus:
		return v.State
	case schema.BlockStorageStatus:
		return v.State
	case schema.ImageStatus:
		return v.State
	case schema.InstanceStatus:
		return v.State
	case schema.NetworkStatus:
		return v.State
	case schema.SubnetStatus:
		return v.State
	case schema.RouteTableStatus:
		return v.State
	case schema.NicStatus:
		return v.State
	case schema.PublicIpStatus:
		return v.State
	case schema.SecurityGroupStatus:
		return v.State
	default:
		return ""
	}
}

func GetStatusConditions[S StatusType](status *S) []schema.StatusCondition {
	if status == nil {
		return nil
	}

	switch v := any(*status).(type) {
	case schema.Status:
		return v.Conditions
	case schema.WorkspaceStatus:
		return v.Conditions
	case schema.BlockStorageStatus:
		return v.Conditions
	case schema.ImageStatus:
		return v.Conditions
	case schema.InstanceStatus:
		return v.Conditions
	case schema.NetworkStatus:
		return v.Conditions
	case schema.SubnetStatus:
		return v.Conditions
	case schema.RouteTableStatus:
		return v.Conditions
	case schema.NicStatus:
		return v.Conditions
	case schema.PublicIpStatus:
		return v.Conditions
	case schema.SecurityGroupStatus:
		return v.Conditions
	default:
		return nil
	}
}

func GetStatusPowerState[S StatusType](status *S) schema.InstanceStatusPowerState {
	if status == nil {
		return ""
	}

	switch v := any(*status).(type) {
	case schema.InstanceStatus:
		return v.PowerState
	default:
		return ""
	}
}
