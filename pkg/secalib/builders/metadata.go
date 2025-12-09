package builders

import "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

// GlobalResourceMetadataBuilder

type globalResourceMetadataBuilder[B any] struct {
	*metadataBuilder[B, schema.GlobalResourceMetadataKind]
	metadata *schema.GlobalResourceMetadata
}

func newGlobalResourceMetadataBuilder[B any](parent *B) *globalResourceMetadataBuilder[B] {
	builder := &globalResourceMetadataBuilder[B]{
		metadata: &schema.GlobalResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[B, schema.GlobalResourceMetadataKind]{
		parent:        parent,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.GlobalResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *globalResourceMetadataBuilder[B]) buildResponse() (*schema.GlobalResourceMetadata, error) {
	if err := validateRequired(builder.validator,
		builder.metadata,
		builder.metadata.Name,
		builder.metadata.Provider,
		builder.metadata.ApiVersion,
		builder.metadata.Kind,
	); err != nil {
		return nil, err
	}

	return builder.metadata, nil
}

// GlobalTenantResourceMetadataBuilder

type globalTenantResourceMetadataBuilder[B any] struct {
	*metadataBuilder[B, schema.GlobalTenantResourceMetadataKind]
	metadata *schema.GlobalTenantResourceMetadata
}

func newGlobalTenantResourceMetadataBuilder[B any](parent *B) *globalTenantResourceMetadataBuilder[B] {
	builder := &globalTenantResourceMetadataBuilder[B]{
		metadata: &schema.GlobalTenantResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[B, schema.GlobalTenantResourceMetadataKind]{
		parent:        parent,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.GlobalTenantResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *globalTenantResourceMetadataBuilder[B]) Tenant(tenant string) *B {
	builder.metadata.Tenant = tenant
	return builder.metadataBuilder.parent
}

func (builder *globalTenantResourceMetadataBuilder[B]) buildResponse() (*schema.GlobalTenantResourceMetadata, error) {
	if err := validateRequired(builder.validator,
		builder.metadata,
		builder.metadata.Name,
		builder.metadata.Provider,
		builder.metadata.ApiVersion,
		builder.metadata.Kind,
		builder.metadata.Tenant,
	); err != nil {
		return nil, err
	}

	return builder.metadata, nil
}

// RegionalResourceMetadata

type regionalResourceMetadataBuilder[B any] struct {
	*metadataBuilder[B, schema.RegionalResourceMetadataKind]
	metadata *schema.RegionalResourceMetadata
}

func newRegionalResourceMetadataBuilder[B any](parent *B) *regionalResourceMetadataBuilder[B] {
	builder := &regionalResourceMetadataBuilder[B]{
		metadata: &schema.RegionalResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[B, schema.RegionalResourceMetadataKind]{
		parent:        parent,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *regionalResourceMetadataBuilder[B]) Tenant(tenant string) *B {
	builder.metadata.Tenant = tenant
	return builder.metadataBuilder.parent
}

func (builder *regionalResourceMetadataBuilder[B]) Region(region string) *B {
	builder.metadata.Region = region
	return builder.metadataBuilder.parent
}

func (builder *regionalResourceMetadataBuilder[B]) buildResponse() (*schema.RegionalResourceMetadata, error) {
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

type regionalWorkspaceResourceMetadataBuilder[B any] struct {
	*metadataBuilder[B, schema.RegionalWorkspaceResourceMetadataKind]
	metadata *schema.RegionalWorkspaceResourceMetadata
}

func newRegionalWorkspaceResourceMetadataBuilder[B any](parent *B) *regionalWorkspaceResourceMetadataBuilder[B] {
	builder := &regionalWorkspaceResourceMetadataBuilder[B]{
		metadata: &schema.RegionalWorkspaceResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[B, schema.RegionalWorkspaceResourceMetadataKind]{
		parent:        parent,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalWorkspaceResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *regionalWorkspaceResourceMetadataBuilder[B]) Tenant(tenant string) *B {
	builder.metadata.Tenant = tenant
	return builder.metadataBuilder.parent
}

func (builder *regionalWorkspaceResourceMetadataBuilder[B]) Workspace(workspace string) *B {
	builder.metadata.Workspace = workspace
	return builder.metadataBuilder.parent
}

func (builder *regionalWorkspaceResourceMetadataBuilder[B]) Region(region string) *B {
	builder.metadata.Region = region
	return builder.metadataBuilder.parent
}

func (builder *regionalWorkspaceResourceMetadataBuilder[B]) buildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {
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

type regionalNetworkResourceMetadataBuilder[B any] struct {
	*metadataBuilder[B, schema.RegionalNetworkResourceMetadataKind]
	metadata *schema.RegionalNetworkResourceMetadata
}

func newRegionalNetworkResourceMetadataBuilder[B any](parent *B) *regionalNetworkResourceMetadataBuilder[B] {
	builder := &regionalNetworkResourceMetadataBuilder[B]{
		metadata: &schema.RegionalNetworkResourceMetadata{},
	}

	builder.metadataBuilder = newMetadataBuilder(newMetadataBuilderParams[B, schema.RegionalNetworkResourceMetadataKind]{
		parent:        parent,
		setName:       func(name string) { builder.metadata.Name = name },
		setProvider:   func(provider string) { builder.metadata.Provider = provider },
		setApiVersion: func(apiVersion string) { builder.metadata.ApiVersion = apiVersion },
		setKind:       func(kind schema.RegionalNetworkResourceMetadataKind) { builder.metadata.Kind = kind },
	})

	return builder
}

func (builder *regionalNetworkResourceMetadataBuilder[B]) Tenant(tenant string) *B {
	builder.metadata.Tenant = tenant
	return builder.metadataBuilder.parent
}

func (builder *regionalNetworkResourceMetadataBuilder[B]) Workspace(workspace string) *B {
	builder.metadata.Workspace = workspace
	return builder.metadataBuilder.parent
}

func (builder *regionalNetworkResourceMetadataBuilder[B]) Network(network string) *B {
	builder.metadata.Network = network
	return builder.metadataBuilder.parent
}

func (builder *regionalNetworkResourceMetadataBuilder[B]) Region(region string) *B {
	builder.metadata.Region = region
	return builder.metadataBuilder.parent
}

func (builder *regionalNetworkResourceMetadataBuilder[B]) buildResponse() (*schema.RegionalNetworkResourceMetadata, error) {
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
