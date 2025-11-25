package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Role

type RoleBuilder struct {
	*resourceBuilder[RoleBuilder, schema.RoleSpec]
	metadata *GlobalTenantResourceMetadataBuilder
	spec     *schema.RoleSpec
}

func NewRoleBuilder() *RoleBuilder {
	builder := &RoleBuilder{
		metadata: NewGlobalTenantResourceMetadataBuilder(),
		spec:     &schema.RoleSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[RoleBuilder, schema.RoleSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.RoleSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *RoleBuilder) Tenant(tenant string) *RoleBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *RoleBuilder) BuildResponse() (*schema.Role, error) {
	medatata, err := builder.metadata.Kind(schema.GlobalTenantResourceMetadataKindResourceKindRole).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Permissions,
	); err != nil {
		return nil, err
	}
	// Validate each permission
	for _, permission := range builder.spec.Permissions {
		if err := validateRequired(builder.validator,
			permission.Provider,
			permission.Resources,
			permission.Verb,
		); err != nil {
			return nil, err
		}
	}

	return &schema.Role{
		Metadata: medatata,
		Labels:   schema.Labels{},
		Spec:     *builder.spec,
		Status:   &schema.RoleStatus{},
	}, nil
}

// RoleAssignment

type RoleAssignmentBuilder struct {
	*resourceBuilder[RoleAssignmentBuilder, schema.RoleAssignmentSpec]
	metadata *GlobalTenantResourceMetadataBuilder
	spec     *schema.RoleAssignmentSpec
}

func NewRoleAssignmentBuilder() *RoleAssignmentBuilder {
	builder := &RoleAssignmentBuilder{
		metadata: NewGlobalTenantResourceMetadataBuilder(),
		spec:     &schema.RoleAssignmentSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[RoleAssignmentBuilder, schema.RoleAssignmentSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.RoleAssignmentSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *RoleAssignmentBuilder) Tenant(tenant string) *RoleAssignmentBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *RoleAssignmentBuilder) BuildResponse() (*schema.RoleAssignment, error) {
	medatata, err := builder.metadata.Kind(schema.GlobalTenantResourceMetadataKindResourceKindRoleAssignment).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Subs,
		builder.spec.Scopes,
		builder.spec.Roles,
	); err != nil {
		return nil, err
	}
	// Validate each scope
	for _, scope := range builder.spec.Scopes {
		if err := validateOneRequired(builder.validator,
			scope.Tenants,
			scope.Workspaces,
			scope.Regions,
		); err != nil {
			return nil, err
		}
	}
	// TODO Validate each scope, if all fields are nil

	return &schema.RoleAssignment{
		Metadata: medatata,
		Labels:   schema.Labels{},
		Spec:     *builder.spec,
		Status:   &schema.RoleAssignmentStatus{},
	}, nil
}
