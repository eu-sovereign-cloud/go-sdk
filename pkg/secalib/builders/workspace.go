package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// workspace

type WorkspaceBuilder struct {
	*resourceBuilder[WorkspaceBuilder, schema.WorkspaceSpec]
	labels   schema.Labels
	metadata *RegionalResourceMetadataBuilder
	spec     *schema.WorkspaceSpec
}

func NewWorkspaceBuilder() *WorkspaceBuilder {
	builder := &WorkspaceBuilder{
		metadata: NewRegionalResourceMetadataBuilder(),
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[WorkspaceBuilder, schema.WorkspaceSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.WorkspaceSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *WorkspaceBuilder) Labels(labels schema.Labels) *WorkspaceBuilder {
	builder.labels = labels
	return builder
}

func (builder *WorkspaceBuilder) Tenant(tenant string) *WorkspaceBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *WorkspaceBuilder) Region(region string) *WorkspaceBuilder {
	builder.metadata.Region(region)
	return builder
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
