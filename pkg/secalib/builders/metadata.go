package builders

import "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

// GlobalResourceMetadataBuilder

type GlobalResourceMetadataBuilder struct {
	*metadataBuilder[GlobalResourceMetadataBuilder, schema.GlobalResourceMetadataKind]
	metadata *schema.GlobalResourceMetadata
}

func newGlobalResourceMetadataBuilder() *GlobalResourceMetadataBuilder {
	builder := &GlobalResourceMetadataBuilder{
		metadata: &schema.GlobalResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[GlobalResourceMetadataBuilder, schema.GlobalResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.GlobalResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *GlobalResourceMetadataBuilder) buildResponse() (*schema.GlobalResourceMetadata, error) {
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

// GlobalTenantResourceMetadataBuilder

type GlobalTenantResourceMetadataBuilder struct {
	*metadataBuilder[GlobalTenantResourceMetadataBuilder, schema.GlobalTenantResourceMetadataKind]
	metadata *schema.GlobalTenantResourceMetadata
}

func newGlobalTenantResourceMetadataBuilder() *GlobalTenantResourceMetadataBuilder {
	builder := &GlobalTenantResourceMetadataBuilder{
		metadata: &schema.GlobalTenantResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[GlobalTenantResourceMetadataBuilder, schema.GlobalTenantResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.GlobalTenantResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *GlobalTenantResourceMetadataBuilder) Tenant(tenant string) *GlobalTenantResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *GlobalTenantResourceMetadataBuilder) buildResponse() (*schema.GlobalTenantResourceMetadata, error) {
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

type RegionalResourceMetadataBuilder struct {
	*metadataBuilder[RegionalResourceMetadataBuilder, schema.RegionalResourceMetadataKind]
	metadata *schema.RegionalResourceMetadata
}

func newRegionalResourceMetadataBuilder() *RegionalResourceMetadataBuilder {
	builder := &RegionalResourceMetadataBuilder{
		metadata: &schema.RegionalResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[RegionalResourceMetadataBuilder, schema.RegionalResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *RegionalResourceMetadataBuilder) Tenant(tenant string) *RegionalResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *RegionalResourceMetadataBuilder) Region(region string) *RegionalResourceMetadataBuilder {
	builder.metadata.Region = region
	return builder
}

func (builder *RegionalResourceMetadataBuilder) buildResponse() (*schema.RegionalResourceMetadata, error) {
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

type RegionalWorkspaceResourceMetadataBuilder struct {
	*metadataBuilder[RegionalWorkspaceResourceMetadataBuilder, schema.RegionalWorkspaceResourceMetadataKind]
	metadata *schema.RegionalWorkspaceResourceMetadata
}

func newRegionalWorkspaceResourceMetadataBuilder() *RegionalWorkspaceResourceMetadataBuilder {
	builder := &RegionalWorkspaceResourceMetadataBuilder{
		metadata: &schema.RegionalWorkspaceResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[RegionalWorkspaceResourceMetadataBuilder, schema.RegionalWorkspaceResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalWorkspaceResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *RegionalWorkspaceResourceMetadataBuilder) Tenant(tenant string) *RegionalWorkspaceResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *RegionalWorkspaceResourceMetadataBuilder) Workspace(workspace string) *RegionalWorkspaceResourceMetadataBuilder {
	builder.metadata.Workspace = workspace
	return builder
}

func (builder *RegionalWorkspaceResourceMetadataBuilder) Region(region string) *RegionalWorkspaceResourceMetadataBuilder {
	builder.metadata.Region = region
	return builder
}

func (builder *RegionalWorkspaceResourceMetadataBuilder) buildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {
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

type RegionalNetworkResourceMetadataBuilder struct {
	*metadataBuilder[RegionalNetworkResourceMetadataBuilder, schema.RegionalNetworkResourceMetadataKind]
	metadata *schema.RegionalNetworkResourceMetadata
}

func newRegionalNetworkResourceMetadataBuilder() *RegionalNetworkResourceMetadataBuilder {
	builder := &RegionalNetworkResourceMetadataBuilder{
		metadata: &schema.RegionalNetworkResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[RegionalNetworkResourceMetadataBuilder, schema.RegionalNetworkResourceMetadataKind]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalNetworkResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *RegionalNetworkResourceMetadataBuilder) Tenant(tenant string) *RegionalNetworkResourceMetadataBuilder {
	builder.metadata.Tenant = tenant
	return builder
}

func (builder *RegionalNetworkResourceMetadataBuilder) Workspace(workspace string) *RegionalNetworkResourceMetadataBuilder {
	builder.metadata.Workspace = workspace
	return builder
}

func (builder *RegionalNetworkResourceMetadataBuilder) Network(network string) *RegionalNetworkResourceMetadataBuilder {
	builder.metadata.Network = network
	return builder
}

func (builder *RegionalNetworkResourceMetadataBuilder) Region(region string) *RegionalNetworkResourceMetadataBuilder {
	builder.metadata.Region = region
	return builder
}

func (builder *RegionalNetworkResourceMetadataBuilder) buildResponse() (*schema.RegionalNetworkResourceMetadata, error) {
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
