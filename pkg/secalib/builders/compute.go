package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/secalib/generators"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Instance

type InstanceMetadataBuilder struct {
	*RegionalWorkspaceResourceMetadataBuilder
}

func NewInstanceMetadataBuilder() *InstanceMetadataBuilder {
	builder := &InstanceMetadataBuilder{
		RegionalWorkspaceResourceMetadataBuilder: newRegionalWorkspaceResourceMetadataBuilder(),
	}

	return builder
}

func (builder *InstanceMetadataBuilder) BuildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindInstance).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateInstanceResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type InstanceBuilder struct {
	*regionalWorkspaceResourceBuilder[InstanceBuilder, schema.InstanceSpec]
	metadata *InstanceMetadataBuilder
	labels   schema.Labels
	spec     *schema.InstanceSpec
}

func NewInstanceBuilder() *InstanceBuilder {
	builder := &InstanceBuilder{
		metadata: NewInstanceMetadataBuilder(),
		spec:     &schema.InstanceSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[InstanceBuilder, schema.InstanceSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[InstanceBuilder, schema.InstanceSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.InstanceSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *InstanceBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.SkuRef,
		builder.spec.Zone,
		builder.spec.BootVolume,
		builder.spec.BootVolume.DeviceRef,
	); err != nil {
		return err
	}

	return nil
}

func (builder *InstanceBuilder) BuildRequest() (*schema.Instance, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.Instance{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *InstanceBuilder) BuildResponse() (*schema.Instance, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	medatata, err := builder.metadata.BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.Instance{
		Metadata: medatata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.InstanceStatus{},
	}, nil
}
