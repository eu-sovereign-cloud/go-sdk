package secalib

import "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

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
		schema.SecurityGroupSpec
}
