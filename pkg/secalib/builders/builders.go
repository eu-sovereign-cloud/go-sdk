package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/secalib"
	"github.com/go-playground/validator/v10"
)

// metadata

type metadataBuilderType interface {
	GlobalResourceMetadataBuilder | GlobalTenantResourceMetadataBuilder | RegionalResourceMetadataBuilder | RegionalWorkspaceResourceMetadataBuilder | RegionalNetworkResourceMetadataBuilder
}

type metadataBuilder[B metadataBuilderType, K any] struct {
	validator     *validator.Validate
	parent        *B
	setName       func(string)
	setProvider   func(string)
	setResource   func(string)
	setApiVersion func(string)
	setKind       func(K)
}

type newMetadataBuilderParams[B metadataBuilderType, K any] struct {
	parent        *B
	setName       func(string)
	setProvider   func(string)
	setResource   func(string)
	setApiVersion func(string)
	setKind       func(K)
}

func newMetadataBuilder[B metadataBuilderType, K any](params newMetadataBuilderParams[B, K]) *metadataBuilder[B, K] {
	return &metadataBuilder[B, K]{
		validator:     validator.New(),
		parent:        params.parent,
		setName:       params.setName,
		setProvider:   params.setProvider,
		setResource:   params.setResource,
		setApiVersion: params.setApiVersion,
		setKind:       params.setKind,
	}
}

func (builder *metadataBuilder[B, K]) Name(name string) *B {
	builder.setName(name)
	return builder.parent
}

func (builder *metadataBuilder[B, K]) Provider(provider string) *B {
	builder.setProvider(provider)
	return builder.parent
}

func (builder *metadataBuilder[B, K]) Resource(resource string) *B {
	builder.setResource(resource)
	return builder.parent
}

func (builder *metadataBuilder[B, K]) ApiVersion(apiVersion string) *B {
	builder.setApiVersion(apiVersion)
	return builder.parent
}

func (builder *metadataBuilder[B, K]) Kind(kind K) *B {
	builder.setKind(kind)
	return builder.parent
}

// resource

/// globalResourceBuilder

type globalResourceBuilder[B any, S secalib.SpecType] struct {
	validator     *validator.Validate
	parent        *B
	setName       func(string)
	setProvider   func(string)
	setResource   func(string)
	setApiVersion func(string)
	setSpec       func(*S)
}

type newGlobalResourceBuilderParams[B any, S secalib.SpecType] struct {
	parent        *B
	setName       func(string)
	setProvider   func(string)
	setResource   func(string)
	setApiVersion func(string)
	setSpec       func(*S)
}

func newGlobalResourceBuilder[B any, S secalib.SpecType](params newGlobalResourceBuilderParams[B, S]) *globalResourceBuilder[B, S] {
	return &globalResourceBuilder[B, S]{
		validator:     validator.New(),
		parent:        params.parent,
		setName:       params.setName,
		setProvider:   params.setProvider,
		setResource:   params.setResource,
		setApiVersion: params.setApiVersion,
		setSpec:       params.setSpec,
	}
}

func (builder *globalResourceBuilder[B, S]) Name(name string) *B {
	builder.setName(name)
	return builder.parent
}

func (builder *globalResourceBuilder[B, S]) Provider(provider string) *B {
	builder.setProvider(provider)
	return builder.parent
}

func (builder *globalResourceBuilder[B, S]) Resource(resource string) *B {
	builder.setResource(resource)
	return builder.parent
}

func (builder *globalResourceBuilder[B, S]) ApiVersion(apiVersion string) *B {
	builder.setApiVersion(apiVersion)
	return builder.parent
}

func (builder *globalResourceBuilder[B, S]) Spec(spec *S) *B {
	builder.setSpec(spec)
	return builder.parent
}

/// globalTenantResourceBuilder

type globalTenantResourceBuilder[B any, S secalib.SpecType] struct {
	*globalResourceBuilder[B, S]

	setTenant func(string)
}

type newGlobalTenantResourceBuilderParams[B any, S secalib.SpecType] struct {
	*newGlobalResourceBuilderParams[B, S]

	setTenant func(string)
}

func newGlobalTenantResourceBuilder[B any, S secalib.SpecType](params newGlobalTenantResourceBuilderParams[B, S]) *globalTenantResourceBuilder[B, S] {
	return &globalTenantResourceBuilder[B, S]{
		globalResourceBuilder: &globalResourceBuilder[B, S]{
			validator:     validator.New(),
			parent:        params.parent,
			setName:       params.setName,
			setProvider:   params.setProvider,
			setResource:   params.setResource,
			setApiVersion: params.setApiVersion,
			setSpec:       params.setSpec,
		},
		setTenant: params.setTenant,
	}
}

func (builder *globalTenantResourceBuilder[B, S]) Tenant(name string) *B {
	builder.setTenant(name)
	return builder.parent
}

// regionalResource

type regionalResourceBuilder[B any, S secalib.SpecType] struct {
	*globalResourceBuilder[B, S]

	setTenant func(string)
	setRegion func(string)
}

type newRegionalResourceBuilderParams[B any, S secalib.SpecType] struct {
	*newGlobalResourceBuilderParams[B, S]

	setTenant func(string)
	setRegion func(string)
}

func newRegionalResourceBuilder[B any, S secalib.SpecType](params newRegionalResourceBuilderParams[B, S]) *regionalResourceBuilder[B, S] {
	return &regionalResourceBuilder[B, S]{
		globalResourceBuilder: &globalResourceBuilder[B, S]{
			validator:     validator.New(),
			parent:        params.parent,
			setName:       params.setName,
			setProvider:   params.setProvider,
			setResource:   params.setResource,
			setApiVersion: params.setApiVersion,
			setSpec:       params.setSpec,
		},
		setTenant: params.setTenant,
		setRegion: params.setRegion,
	}
}

func (builder *regionalResourceBuilder[B, S]) Tenant(name string) *B {
	builder.setTenant(name)
	return builder.parent
}

func (builder *regionalResourceBuilder[B, S]) Region(region string) *B {
	builder.setRegion(region)
	return builder.parent
}

// regionalWorkspaceResource

type regionalWorkspaceResourceBuilder[B any, S secalib.SpecType] struct {
	*globalResourceBuilder[B, S]

	setTenant    func(string)
	setWorkspace func(string)
	setRegion    func(string)
}

type newRegionalWorkspaceResourceBuilderParams[B any, S secalib.SpecType] struct {
	*newGlobalResourceBuilderParams[B, S]

	setTenant    func(string)
	setWorkspace func(string)
	setRegion    func(string)
}

func newRegionalWorkspaceResourceBuilder[B any, S secalib.SpecType](params newRegionalWorkspaceResourceBuilderParams[B, S]) *regionalWorkspaceResourceBuilder[B, S] {
	return &regionalWorkspaceResourceBuilder[B, S]{
		globalResourceBuilder: &globalResourceBuilder[B, S]{
			validator:     validator.New(),
			parent:        params.parent,
			setName:       params.setName,
			setProvider:   params.setProvider,
			setResource:   params.setResource,
			setApiVersion: params.setApiVersion,
			setSpec:       params.setSpec,
		},
		setTenant:    params.setTenant,
		setWorkspace: params.setWorkspace,
		setRegion:    params.setRegion,
	}
}

func (builder *regionalWorkspaceResourceBuilder[B, S]) Tenant(name string) *B {
	builder.setTenant(name)
	return builder.parent
}

func (builder *regionalWorkspaceResourceBuilder[B, S]) Workspace(workspace string) *B {
	builder.setWorkspace(workspace)
	return builder.parent
}

func (builder *regionalWorkspaceResourceBuilder[B, S]) Region(region string) *B {
	builder.setRegion(region)
	return builder.parent
}

// regionalNetworkResource

type regionalNetworkResourceBuilder[B any, S secalib.SpecType] struct {
	*globalResourceBuilder[B, S]

	setTenant    func(string)
	setWorkspace func(string)
	setNetwork   func(string)
	setRegion    func(string)
}

type newRegionalNetworkResourceBuilderParams[B any, S secalib.SpecType] struct {
	*newGlobalResourceBuilderParams[B, S]

	setTenant    func(string)
	setWorkspace func(string)
	setNetwork   func(string)
	setRegion    func(string)
}

func newRegionalNetworkResourceBuilder[B any, S secalib.SpecType](params newRegionalNetworkResourceBuilderParams[B, S]) *regionalNetworkResourceBuilder[B, S] {
	return &regionalNetworkResourceBuilder[B, S]{
		globalResourceBuilder: &globalResourceBuilder[B, S]{
			validator:     validator.New(),
			parent:        params.parent,
			setName:       params.setName,
			setProvider:   params.setProvider,
			setResource:   params.setResource,
			setApiVersion: params.setApiVersion,
			setSpec:       params.setSpec,
		},
		setTenant:    params.setTenant,
		setWorkspace: params.setWorkspace,
		setNetwork:   params.setNetwork,
		setRegion:    params.setRegion,
	}
}

func (builder *regionalNetworkResourceBuilder[B, S]) Tenant(name string) *B {
	builder.setTenant(name)
	return builder.parent
}

func (builder *regionalNetworkResourceBuilder[B, S]) Workspace(workspace string) *B {
	builder.setWorkspace(workspace)
	return builder.parent
}

func (builder *regionalNetworkResourceBuilder[B, S]) Network(network string) *B {
	builder.setNetwork(network)
	return builder.parent
}

func (builder *regionalNetworkResourceBuilder[B, S]) Region(region string) *B {
	builder.setRegion(region)
	return builder.parent
}
