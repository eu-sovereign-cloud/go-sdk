package secalib

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// TODO Add all tags returned in conformane tests
func BuildResponseGlobalResourceMetadata(name string, tenant string) *schema.GlobalResourceMetadata {
	return &schema.GlobalResourceMetadata{
		Name:   name,
		Tenant: tenant,
	}
}

func BuildResponseSkuResourceMetadata(name string) *schema.SkuResourceMetadata {
	return &schema.SkuResourceMetadata{
		Name: name,
	}
}

func BuildResponseRegionalResourceMetadata(name string, tenant string) *schema.RegionalResourceMetadata {
	return &schema.RegionalResourceMetadata{
		Name:   name,
		Tenant: tenant,
	}
}

func BuildResponseRegionalWorkspaceResourceMetadata(name string, tenant string, workspace string) *schema.RegionalWorkspaceResourceMetadata {
	return &schema.RegionalWorkspaceResourceMetadata{
		Name:      name,
		Tenant:    tenant,
		Workspace: workspace,
	}
}
