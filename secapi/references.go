package secapi

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

type Reference interface {
	TenantReference | WorkspaceReference | NetworkReference
}

type ReferencedResource[Ref Reference, Data any] struct {
	Ref  Ref
	Data Data
}

type TenantReference struct {
	Tenant TenantID
	Name   string
}

type WorkspaceReference struct {
	Tenant    TenantID
	Workspace WorkspaceID
	Name      string
}

type NetworkReference struct {
	Tenant    TenantID
	Workspace WorkspaceID
	Network   NetworkID
	Name      string
}

// Validators

func (tref *TenantReference) validate() error {
	if tref.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if tref.Name == "" {
		return ErrNoMetatadaName
	}

	return nil
}

func (wref *WorkspaceReference) validate() error {
	if wref.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if wref.Workspace == "" {
		return ErrNoMetatadaWorkspace
	}

	if wref.Name == "" {
		return ErrNoMetatadaName
	}

	return nil
}

func (nref *NetworkReference) validate() error {
	if nref.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if nref.Workspace == "" {
		return ErrNoMetatadaWorkspace
	}

	if nref.Network == "" {
		return ErrNoPathMetadata
	}

	if nref.Name == "" {
		return ErrNoMetatadaName
	}

	return nil
}

// Converters

func BuildReferenceObj(provider, region, resourceName, tenant, workspace string) *schema.Reference {
	ref := &schema.Reference{
		Provider:  &provider,
		Region:    &region,
		Resource:  resourceName,
		Tenant:    &tenant,
		Workspace: &workspace,
	}

	return ref
}
