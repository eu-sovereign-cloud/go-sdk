package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Instance

type InstanceBuilder struct {
	*resourceBuilder[InstanceBuilder, schema.InstanceSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.InstanceSpec
}

func NewInstanceBuilder() *InstanceBuilder {
	builder := &InstanceBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.InstanceSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[InstanceBuilder, schema.InstanceSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.InstanceSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *InstanceBuilder) Tenant(tenant string) *InstanceBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *InstanceBuilder) Workspace(workspace string) *InstanceBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *InstanceBuilder) Region(region string) *InstanceBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *InstanceBuilder) BuildResponse() (*schema.Instance, error) {
	medatata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindInstance).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.SkuRef,
		builder.spec.Zone,
		builder.spec.BootVolume,
		builder.spec.BootVolume.DeviceRef,
	); err != nil {
		return nil, err
	}

	return &schema.Instance{
		Metadata: medatata,
		Labels:   schema.Labels{},
		Spec:     *builder.spec,
		Status:   &schema.InstanceStatus{},
	}, nil
}
