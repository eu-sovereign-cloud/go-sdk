package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/secalib/generators"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Region

type RegionMetadataBuilder struct {
	*globalResourceMetadataBuilder[RegionMetadataBuilder]
}

func NewRegionMetadataBuilder() *RegionMetadataBuilder {
	builder := &RegionMetadataBuilder{}
	builder.globalResourceMetadataBuilder = newGlobalResourceMetadataBuilder(builder)
	return builder
}

func (builder *RegionMetadataBuilder) BuildResponse() (*schema.GlobalResourceMetadata, error) {
	medatata, err := builder.kind(schema.GlobalResourceMetadataKindResourceKindRegion).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateRegionResource(builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type RegionBuilder struct {
	*globalResourceBuilder[RegionBuilder, schema.RegionSpec]
	metadata *RegionMetadataBuilder
	spec     *schema.RegionSpec
}

func NewRegionBuilder() *RegionBuilder {
	builder := &RegionBuilder{
		metadata: NewRegionMetadataBuilder(),
		spec:     &schema.RegionSpec{},
	}

	builder.globalResourceBuilder = newGlobalResourceBuilder(newGlobalResourceBuilderParams[RegionBuilder, schema.RegionSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.RegionSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *RegionBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.AvailableZones,
		builder.spec.Providers,
	); err != nil {
		return err
	}

	// Validate each provider
	for _, provider := range builder.spec.Providers {
		if err := validateRequired(builder.validator,
			provider.Name,
			provider.Version,
			provider.Url,
		); err != nil {
			return err
		}
	}

	return nil
}

func (builder *RegionBuilder) BuildRequest() (*schema.Region, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.Region{
		Metadata: nil,
		Spec:     *builder.spec,
	}, nil
}

func (builder *RegionBuilder) BuildResponse() (*schema.Region, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	medatata, err := builder.metadata.BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.Region{
		Metadata: medatata,
		Spec:     *builder.spec,
	}, nil
}
