package secapi

import "errors"

var (
	ErrNoTenantID          = errors.New("tenant ID not found in context")
	ErrNoMetatada          = errors.New("metadata is empty")
	ErrNoMetatadaWorkspace = errors.New("metadata workspace is empty")
	ErrNoMetatadaTenant    = errors.New("metadata tenant is empty")

	ErrRegionRequiredToRegionalClient = errors.New("region provider is required to create a regional client")
)
