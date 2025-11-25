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

type resourceBuilderType interface {
	RegionBuilder |
		RoleBuilder |
		RoleAssignmentBuilder |
		WorkspaceBuilder |
		BlockStorageBuilder |
		ImageBuilder |
		InstanceBuilder |
		NetworkBuilder |
		InternetGatewayBuilder |
		RouteTableBuilder |
		SubnetBuilder |
		PublicIpBuilder |
		NicBuilder |
		SecurityGroupBuilder
}

type resourceBuilder[B resourceBuilderType, S secalib.SpecType] struct {
	validator     *validator.Validate
	parent        *B
	setName       func(string)
	setProvider   func(string)
	setResource   func(string)
	setApiVersion func(string)
	setSpec       func(*S)
}

type newResourceBuilderParams[B resourceBuilderType, S secalib.SpecType] struct {
	parent        *B
	setName       func(string)
	setProvider   func(string)
	setResource   func(string)
	setApiVersion func(string)
	setSpec       func(*S)
}

func newResourceBuilder[B resourceBuilderType, S secalib.SpecType](params newResourceBuilderParams[B, S]) *resourceBuilder[B, S] {
	return &resourceBuilder[B, S]{
		validator:     validator.New(),
		parent:        params.parent,
		setName:       params.setName,
		setProvider:   params.setProvider,
		setResource:   params.setResource,
		setApiVersion: params.setApiVersion,
		setSpec:       params.setSpec,
	}
}

func (builder *resourceBuilder[B, S]) Name(name string) *B {
	builder.setName(name)
	return builder.parent
}

func (builder *resourceBuilder[B, S]) Provider(provider string) *B {
	builder.setProvider(provider)
	return builder.parent
}

func (builder *resourceBuilder[B, S]) Resource(resource string) *B {
	builder.setResource(resource)
	return builder.parent
}

func (builder *resourceBuilder[B, S]) ApiVersion(apiVersion string) *B {
	builder.setApiVersion(apiVersion)
	return builder.parent
}

func (builder *resourceBuilder[B, S]) Spec(spec *S) *B {
	builder.setSpec(spec)
	return builder.parent
}
