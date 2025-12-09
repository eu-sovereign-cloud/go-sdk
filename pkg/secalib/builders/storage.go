package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/secalib/generators"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// BlockStorage

type BlockStorageMetadataBuilder struct {
	*regionalWorkspaceResourceMetadataBuilder[BlockStorageMetadataBuilder]
}

func NewBlockStorageMetadataBuilder() *BlockStorageMetadataBuilder {
	builder := &BlockStorageMetadataBuilder{}
	builder.regionalWorkspaceResourceMetadataBuilder = newRegionalWorkspaceResourceMetadataBuilder(builder)
	return builder
}

func (builder *BlockStorageMetadataBuilder) BuildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {
	medatata, err := builder.kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindBlockStorage).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateBlockStorageResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type BlockStorageBuilder struct {
	*regionalWorkspaceResourceBuilder[BlockStorageBuilder, schema.BlockStorageSpec]
	metadata *BlockStorageMetadataBuilder
	labels   schema.Labels
	spec     *schema.BlockStorageSpec
}

func NewBlockStorageBuilder() *BlockStorageBuilder {
	builder := &BlockStorageBuilder{
		metadata: NewBlockStorageMetadataBuilder(),
		spec:     &schema.BlockStorageSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[BlockStorageBuilder, schema.BlockStorageSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[BlockStorageBuilder, schema.BlockStorageSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.BlockStorageSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *BlockStorageBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.SkuRef,
		builder.spec.SizeGB,
	); err != nil {
		return err
	}

	return nil
}

func (builder *BlockStorageBuilder) BuildRequest() (*schema.BlockStorage, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.BlockStorage{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *BlockStorageBuilder) BuildResponse() (*schema.BlockStorage, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	medatata, err := builder.metadata.buildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.BlockStorage{
		Metadata: medatata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.BlockStorageStatus{},
	}, nil
}

// Image

type ImageMetadataBuilder struct {
	*regionalResourceMetadataBuilder[ImageMetadataBuilder]
}

func NewImageMetadataBuilder() *ImageMetadataBuilder {
	builder := &ImageMetadataBuilder{}
	builder.regionalResourceMetadataBuilder = newRegionalResourceMetadataBuilder(builder)
	return builder
}

func (builder *ImageMetadataBuilder) BuildResponse() (*schema.RegionalResourceMetadata, error) {
	medatata, err := builder.kind(schema.RegionalResourceMetadataKindResourceKindImage).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateImageResource(builder.metadata.Tenant, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type ImageBuilder struct {
	*regionalResourceBuilder[ImageBuilder, schema.ImageSpec]
	metadata *ImageMetadataBuilder
	labels   schema.Labels
	spec     *schema.ImageSpec
}

func NewImageBuilder() *ImageBuilder {
	builder := &ImageBuilder{
		metadata: NewImageMetadataBuilder(),
		spec:     &schema.ImageSpec{},
	}

	builder.regionalResourceBuilder = newRegionalResourceBuilder(newRegionalResourceBuilderParams[ImageBuilder, schema.ImageSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[ImageBuilder, schema.ImageSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.ImageSpec) { builder.spec = spec },
		},
		setTenant: func(tenant string) { builder.metadata.Tenant(tenant) },
		setRegion: func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *ImageBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.BlockStorageRef,
		builder.spec.CpuArchitecture,
	); err != nil {
		return err
	}

	return nil
}

func (builder *ImageBuilder) BuildRequest() (*schema.Image, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.Image{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *ImageBuilder) BuildResponse() (*schema.Image, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	medatata, err := builder.metadata.buildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.Image{
		Metadata: medatata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.ImageStatus{},
	}, nil
}
