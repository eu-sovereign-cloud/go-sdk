package secapi

import "github.com/eu-sovereign-cloud/go-sdk/secapi/builders"

// Filters

type ListFilter interface {
	TenantReference | WorkspaceReference | NetworkReference
}

type GlobalFilter struct {
	Options *FilterOptions
}

type TenantFilter struct {
	Tenant  TenantID
	Options *FilterOptions
}

type WorkspaceFilter struct {
	Tenant    TenantID
	Workspace WorkspaceID
	Options   *FilterOptions
}

type NetworkFilter struct {
	Tenant    TenantID
	Workspace WorkspaceID
	Network   NetworkID
	Options   *FilterOptions
}

// Options

const defaultLimit = 1000

type FilterOptions struct {
	Limit  *int
	Labels *builders.LabelsBuilder
}

func NewFilterOptions() *FilterOptions {
	limit := defaultLimit
	return &FilterOptions{
		Limit:  &limit,
		Labels: builders.NewLabelsBuilder(),
	}
}

func (o *FilterOptions) WithLimit(limit int) *FilterOptions {
	o.Limit = &limit
	return o
}

func (o *FilterOptions) WithLabels(labels *builders.LabelsBuilder) *FilterOptions {
	o.Labels = labels
	return o
}

// Validators

func (tref *TenantFilter) validate() error {
	if tref.Tenant == "" {
		return ErrNoMetadataTenant
	}

	return nil
}

func (wref *WorkspaceFilter) validate() error {
	if wref.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if wref.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	return nil
}

func (nref *NetworkFilter) validate() error {
	if nref.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if nref.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	if nref.Network == "" {
		return ErrNoMetadataNetwork
	}

	return nil
}
