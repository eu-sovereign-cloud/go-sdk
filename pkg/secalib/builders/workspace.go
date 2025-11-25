package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// workspace

type WorkspaceBuilder struct {
	*regionalResourceBuilder[WorkspaceBuilder, schema.WorkspaceSpec]
	metadata *RegionalResourceMetadataBuilder
	labels   schema.Labels
}

func NewWorkspaceBuilder() *WorkspaceBuilder {
	builder := &WorkspaceBuilder{
		metadata: NewRegionalResourceMetadataBuilder(),
	}

	builder.regionalResourceBuilder = newRegionalResourceBuilder(newRegionalResourceBuilderParams[WorkspaceBuilder, schema.WorkspaceSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[WorkspaceBuilder, schema.WorkspaceSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
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
	medatata, err := builder.metadata.Kind(schema.RegionalResourceMetadataKindResourceKindWorkspace).BuildResponse()
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
