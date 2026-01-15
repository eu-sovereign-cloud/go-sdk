package secatest

import "github.com/eu-sovereign-cloud/go-sdk/pkg/constants"

const (
	// Providers
	ProviderRegionV1Endpoint        = providersEndpointPrefix + "/" + constants.RegionProviderV1Name
	ProviderAuthorizationV1Endpoint = providersEndpointPrefix + "/" + constants.AuthorizationProviderV1Name
	ProviderWorkspaceV1Endpoint     = providersEndpointPrefix + "/" + constants.WorkspaceProviderV1Name
	ProviderNetworkV1Endpoint       = providersEndpointPrefix + "/" + constants.NetworkProviderV1Name
	ProviderComputeV1Endpoint       = providersEndpointPrefix + "/" + constants.ComputeProviderV1Name
	ProviderStorageV1Endpoint       = providersEndpointPrefix + "/" + constants.StorageProviderV1Name
)
