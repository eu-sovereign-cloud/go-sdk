package secapi

import "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

type TenantID string

type WorkspaceID string

type NetworkID string

type ResourceType interface {
	schema.Role |
		schema.RoleAssignment |
		schema.Workspace |
		schema.BlockStorage |
		schema.Image |
		schema.Instance |
		schema.Network |
		schema.InternetGateway |
		schema.RouteTable |
		schema.Subnet |
		schema.PublicIp |
		schema.Nic |
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
	schema.RoleSpec |
		schema.RoleAssignmentSpec |
		schema.WorkspaceSpec |
		schema.BlockStorageSpec |
		schema.ImageSpec |
		schema.InstanceSpec |
		schema.NetworkSpec |
		schema.InternetGatewaySpec |
		schema.RouteTableSpec |
		schema.SubnetSpec |
		schema.PublicIpSpec |
		schema.NicSpec |
		schema.SecurityGroupSpec
}
