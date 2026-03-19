package secapi

type TenantPath struct {
	Tenant TenantID
}

type WorkspacePath struct {
	Tenant    TenantID
	Workspace WorkspaceID
}

type NetworkPath struct {
	Tenant    TenantID
	Workspace WorkspaceID
	Network   NetworkID
}

// Validators

func (path *TenantPath) validate() error {
	if path.Tenant == "" {
		return ErrNoMetadataTenant
	}

	return nil
}

func (path *WorkspacePath) validate() error {
	if path.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if path.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	return nil
}

func (path *NetworkPath) validate() error {
	if path.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if path.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	if path.Network == "" {
		return ErrNoMetadataNetwork
	}

	return nil
}
