package secapi

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
		return ErrNoMetadataTenant
	}

	if tref.Name == "" {
		return ErrNoMetadataName
	}

	return nil
}

func (wref *WorkspaceReference) validate() error {
	if wref.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if wref.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	if wref.Name == "" {
		return ErrNoMetadataName
	}

	return nil
}

func (nref *NetworkReference) validate() error {
	if nref.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if nref.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	if nref.Network == "" {
		return ErrNoMetadataNetwork
	}

	if nref.Name == "" {
		return ErrNoMetadataName
	}

	return nil
}
