package secatest

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

func NewGlobalResourceMetadata(name string) *schema.GlobalResourceMetadata {
	return &schema.GlobalResourceMetadata{
		Name: name,
	}
}

func NewGlobalTenantResourceMetadata(name string, tenant string) *schema.GlobalTenantResourceMetadata {
	return &schema.GlobalTenantResourceMetadata{
		Name:   name,
		Tenant: tenant,
	}
}

func NewSkuResourceMetadata(name string, tenant string) *schema.SkuResourceMetadata {
	return &schema.SkuResourceMetadata{
		Name:   name,
		Tenant: tenant,
	}
}

func NewRegionalResourceMetadata(name string, tenant string, region string) *schema.RegionalResourceMetadata {
	return &schema.RegionalResourceMetadata{
		Name:   name,
		Tenant: tenant,
		Region: region,
	}
}

func NewRegionalWorkspaceResourceMetadata(name string, tenant string, workspace string, region string) *schema.RegionalWorkspaceResourceMetadata {
	return &schema.RegionalWorkspaceResourceMetadata{
		Name:      name,
		Tenant:    tenant,
		Workspace: workspace,
		Region:    region,
	}
}

func NewRegionalNetworkResourceMetadata(name string, tenant string, workspace string, network string, region string) *schema.RegionalNetworkResourceMetadata {
	return &schema.RegionalNetworkResourceMetadata{
		Name:      name,
		Tenant:    tenant,
		Workspace: workspace,
		Network:   network,
		Region:    region,
	}
}
