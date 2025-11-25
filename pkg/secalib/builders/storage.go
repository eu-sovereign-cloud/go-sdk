package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// BlockStorage

type BlockStorageBuilder struct {
	*resourceBuilder[BlockStorageBuilder, schema.BlockStorageSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.BlockStorageSpec
}

func NewBlockStorageBuilder() *BlockStorageBuilder {
	builder := &BlockStorageBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.BlockStorageSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[BlockStorageBuilder, schema.BlockStorageSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.BlockStorageSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *BlockStorageBuilder) Tenant(tenant string) *BlockStorageBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *BlockStorageBuilder) Workspace(workspace string) *BlockStorageBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *BlockStorageBuilder) Region(region string) *BlockStorageBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *BlockStorageBuilder) BuildResponse() (*schema.BlockStorage, error) {
	medatata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindBlockStorage).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.SkuRef,
		builder.spec.SizeGB,
	); err != nil {
		return nil, err
	}

	return &schema.BlockStorage{
		Metadata: medatata,
		Labels:   schema.Labels{},
		Spec:     *builder.spec,
		Status:   &schema.BlockStorageStatus{},
	}, nil
}

// Image

type ImageBuilder struct {
	*resourceBuilder[ImageBuilder, schema.ImageSpec]
	metadata *RegionalResourceMetadataBuilder
	spec     *schema.ImageSpec
}

func NewImageBuilder() *ImageBuilder {
	builder := &ImageBuilder{
		metadata: NewRegionalResourceMetadataBuilder(),
		spec:     &schema.ImageSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[ImageBuilder, schema.ImageSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.ImageSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *ImageBuilder) Tenant(tenant string) *ImageBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *ImageBuilder) Region(region string) *ImageBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *ImageBuilder) BuildResponse() (*schema.Image, error) {
	medatata, err := builder.metadata.Kind(schema.RegionalResourceMetadataKindResourceKindImage).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.BlockStorageRef,
		builder.spec.CpuArchitecture,
	); err != nil {
		return nil, err
	}

	return &schema.Image{
		Metadata: medatata,
		Labels:   schema.Labels{},
		Spec:     *builder.spec,
		Status:   &schema.ImageStatus{},
	}, nil
}
