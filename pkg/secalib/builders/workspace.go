package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/secalib/generators"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// workspace

type WorkspaceMetadataBuilder struct {
	*regionalResourceMetadataBuilder[WorkspaceMetadataBuilder]
}

func NewWorkspaceMetadataBuilder() *WorkspaceMetadataBuilder {
	builder := &WorkspaceMetadataBuilder{}
	builder.regionalResourceMetadataBuilder = newRegionalResourceMetadataBuilder(builder)
	return builder
}

func (builder *WorkspaceMetadataBuilder) BuildResponse() (*schema.RegionalResourceMetadata, error) {
	medatata, err := builder.kind(schema.RegionalResourceMetadataKindResourceKindWorkspace).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateWorkspaceResource(builder.metadata.Tenant, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type WorkspaceBuilder struct {
	*regionalResourceBuilder[WorkspaceBuilder, schema.WorkspaceSpec]
	metadata *WorkspaceMetadataBuilder
	labels   schema.Labels
}

func NewWorkspaceBuilder() *WorkspaceBuilder {
	builder := &WorkspaceBuilder{
		metadata: NewWorkspaceMetadataBuilder(),
	}

	builder.regionalResourceBuilder = newRegionalResourceBuilder(newRegionalResourceBuilderParams[WorkspaceBuilder, schema.WorkspaceSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[WorkspaceBuilder, schema.WorkspaceSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
		},
		setTenant: func(tenant string) { builder.metadata.Tenant(tenant) },
		setRegion: func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *WorkspaceBuilder) BuildRequest() (*schema.Workspace, error) {
	return &schema.Workspace{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     schema.WorkspaceSpec{},
		Status:   nil,
	}, nil
}

func (builder *WorkspaceBuilder) BuildResponse() (*schema.Workspace, error) {
	medatata, err := builder.metadata.buildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.Workspace{
		Metadata: medatata,
		Labels:   builder.labels,
		Spec:     schema.WorkspaceSpec{},
		Status:   &schema.WorkspaceStatus{},
	}, nil
}
