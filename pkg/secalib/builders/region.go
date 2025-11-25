package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Region

type RegionBuilder struct {
	*resourceBuilder[RegionBuilder, schema.RegionSpec]
	metadata *GlobalResourceMetadataBuilder
	spec     *schema.RegionSpec
}

func NewRegionBuilder() *RegionBuilder {
	builder := &RegionBuilder{
		metadata: NewGlobalResourceMetadataBuilder(),
		spec:     &schema.RegionSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[RegionBuilder, schema.RegionSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.RegionSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *RegionBuilder) BuildResponse() (*schema.Region, error) {
	medatata, err := builder.metadata.Kind(schema.GlobalResourceMetadataKindResourceKindRegion).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.AvailableZones,
		builder.spec.Providers,
	); err != nil {
		return nil, err
	}
	// Validate each provider
	for _, provider := range builder.spec.Providers {
		if err := validateRequired(builder.validator,
			provider.Name,
			provider.Version,
			provider.Url,
		); err != nil {
			return nil, err
		}
	}

	return &schema.Region{
		Metadata: medatata,
		Spec:     *builder.spec,
	}, nil
}
