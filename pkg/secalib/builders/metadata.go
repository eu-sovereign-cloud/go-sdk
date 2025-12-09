package builders

import "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

// GlobalResourceMetadataBuilder

type globalResourceMetadataBuilder struct {
	*metadataBuilder[globalResourceMetadataBuilder, schema.GlobalResourceMetadataKind]
	metadata *schema.GlobalResourceMetadata
}

func newGlobalResourceMetadataBuilder() *globalResourceMetadataBuilder {
	builder := &globalResourceMetadataBuilder{
		metadata: &schema.GlobalResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[globalResourceMetadataBuilder, schema.GlobalResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.GlobalResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *globalResourceMetadataBuilder) buildResponse() (*schema.GlobalResourceMetadata, error) {
	if err := validateRequired(builder.validator,
		builder.metadata,
		builder.metadata.Name,
		builder.metadata.Provider,
		builder.metadata.Resource,
		builder.metadata.ApiVersion,
		builder.metadata.Kind,
	); err != nil {
		return nil, err
	}

	return builder.metadata, nil
}

// globalTenantResourceMetadataBuilder

type globalTenantResourceMetadataBuilder struct {
	*metadataBuilder[globalTenantResourceMetadataBuilder, schema.GlobalTenantResourceMetadataKind]
	metadata *schema.GlobalTenantResourceMetadata
}

func newGlobalTenantResourceMetadataBuilder() *globalTenantResourceMetadataBuilder {
	builder := &globalTenantResourceMetadataBuilder{
		metadata: &schema.GlobalTenantResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[globalTenantResourceMetadataBuilder, schema.GlobalTenantResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.GlobalTenantResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *globalTenantResourceMetadataBuilder) Tenant(tenant string) *globalTenantResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *globalTenantResourceMetadataBuilder) buildResponse() (*schema.GlobalTenantResourceMetadata, error) {
	if err := validateRequired(builder.validator,
		builder.metadata,
		builder.metadata.Name,
		builder.metadata.Provider,
		builder.metadata.Resource,
		builder.metadata.ApiVersion,
		builder.metadata.Kind,
		builder.metadata.Tenant,
	); err != nil {
		return nil, err
	}

	return builder.metadata, nil
}

// RegionalResourceMetadata

type regionalResourceMetadataBuilder struct {
	*metadataBuilder[regionalResourceMetadataBuilder, schema.RegionalResourceMetadataKind]
	metadata *schema.RegionalResourceMetadata
}

func newRegionalResourceMetadataBuilder() *regionalResourceMetadataBuilder {
	builder := &regionalResourceMetadataBuilder{
		metadata: &schema.RegionalResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[regionalResourceMetadataBuilder, schema.RegionalResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *regionalResourceMetadataBuilder) Tenant(tenant string) *regionalResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *regionalResourceMetadataBuilder) Region(region string) *regionalResourceMetadataBuilder {
	builder.metadata.Region = region
	return builder
}

func (builder *regionalResourceMetadataBuilder) buildResponse() (*schema.RegionalResourceMetadata, error) {
	if err := validateRequired(builder.validator,
		builder.metadata,
		builder.metadata.Name,
		builder.metadata.Provider,
		builder.metadata.ApiVersion,
		builder.metadata.Kind,
		builder.metadata.Tenant,
		builder.metadata.Region,
	); err != nil {
		return nil, err
	}

	return builder.metadata, nil
}

// RegionalWorkspaceResourceMetadata

type regionalWorkspaceResourceMetadataBuilder struct {
	*metadataBuilder[regionalWorkspaceResourceMetadataBuilder, schema.RegionalWorkspaceResourceMetadataKind]
	metadata *schema.RegionalWorkspaceResourceMetadata
}

func newRegionalWorkspaceResourceMetadataBuilder() *regionalWorkspaceResourceMetadataBuilder {
	builder := &regionalWorkspaceResourceMetadataBuilder{
		metadata: &schema.RegionalWorkspaceResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[regionalWorkspaceResourceMetadataBuilder, schema.RegionalWorkspaceResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalWorkspaceResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *regionalWorkspaceResourceMetadataBuilder) Tenant(tenant string) *regionalWorkspaceResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *regionalWorkspaceResourceMetadataBuilder) Workspace(workspace string) *regionalWorkspaceResourceMetadataBuilder {
	builder.metadata.Workspace = workspace
	return builder
}

func (builder *regionalWorkspaceResourceMetadataBuilder) Region(region string) *regionalWorkspaceResourceMetadataBuilder {
	builder.metadata.Region = region
	return builder
}

func (builder *regionalWorkspaceResourceMetadataBuilder) buildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {
	if err := validateRequired(builder.validator,
		builder.metadata,
		builder.metadata.Name,
		builder.metadata.Provider,
		builder.metadata.ApiVersion,
		builder.metadata.Kind,
		builder.metadata.Tenant,
		builder.metadata.Workspace,
		builder.metadata.Region,
	); err != nil {
		return nil, err
	}

	return builder.metadata, nil
}

// RegionalNetworkResourceMetadata

type regionalNetworkResourceMetadataBuilder struct {
	*metadataBuilder[regionalNetworkResourceMetadataBuilder, schema.RegionalNetworkResourceMetadataKind]
	metadata *schema.RegionalNetworkResourceMetadata
}

func newRegionalNetworkResourceMetadataBuilder() *regionalNetworkResourceMetadataBuilder {
	builder := &regionalNetworkResourceMetadataBuilder{
		metadata: &schema.RegionalNetworkResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[regionalNetworkResourceMetadataBuilder, schema.RegionalNetworkResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalNetworkResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *regionalNetworkResourceMetadataBuilder) Tenant(tenant string) *regionalNetworkResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *regionalNetworkResourceMetadataBuilder) Workspace(workspace string) *regionalNetworkResourceMetadataBuilder {
	builder.metadata.Workspace = workspace
	return builder
}

func (builder *regionalNetworkResourceMetadataBuilder) Network(network string) *regionalNetworkResourceMetadataBuilder {
	builder.metadata.Network = network
	return builder
}

func (builder *regionalNetworkResourceMetadataBuilder) Region(region string) *regionalNetworkResourceMetadataBuilder {
	builder.metadata.Region = region
	return builder
}

func (builder *regionalNetworkResourceMetadataBuilder) buildResponse() (*schema.RegionalNetworkResourceMetadata, error) {
	if err := validateRequired(builder.validator,
		builder.metadata,
		builder.metadata.Name,
		builder.metadata.Provider,
		builder.metadata.ApiVersion,
		builder.metadata.Kind,
		builder.metadata.Tenant,
		builder.metadata.Workspace,
		builder.metadata.Network,
		builder.metadata.Region,
	); err != nil {
		return nil, err
	}

	return builder.metadata, nil
}
