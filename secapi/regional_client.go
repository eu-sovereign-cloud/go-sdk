package secapi

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/constants"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

type RegionalClient struct {
	authToken string

	WorkspaceV1 WorkspaceV1
	ComputeV1   ComputeV1
	StorageV1   StorageV1
	NetworkV1   NetworkV1
}

func newRegionalClient(authToken string, region *schema.Region) (*RegionalClient, error) {
	client := &RegionalClient{
		authToken: authToken,
	}

	// Initializes workspaceV1 API client
	workspaceV1provider := findRegionalProvider(constants.WorkspaceProviderName, constants.ApiVersion1, region)
	if workspaceV1provider != nil {
		if err := initRegionalAPI(client, workspaceV1provider, newWorkspaceV1Impl, client.setWorkspaceV1); err != nil {
			return nil, err
		}
	} else {
		setUnavailableRegionalAPI(newWorkspaceV1Unavailable, client.setWorkspaceV1)
	}

	// Initializes computeV1 API client
	computeV1provider := findRegionalProvider(constants.ComputeProviderName, constants.ApiVersion1, region)
	if computeV1provider != nil {
		if err := initRegionalAPI(client, computeV1provider, newComputeV1Impl, client.setComputeV1); err != nil {
			return nil, err
		}
	} else {
		setUnavailableRegionalAPI(newComputeV1Unavailable, client.setComputeV1)
	}

	// Initializes storageV1 API client
	storageV1provider := findRegionalProvider(constants.StorageProviderName, constants.ApiVersion1, region)
	if storageV1provider != nil {
		if err := initRegionalAPI(client, storageV1provider, newStorageV1Impl, client.setStorageV1); err != nil {
			return nil, err
		}
	} else {
		setUnavailableRegionalAPI(newStorageV1Unavailable, client.setStorageV1)
	}

	// Initializes networkV1 API client
	networkV1provider := findRegionalProvider(constants.NetworkProviderName, constants.ApiVersion1, region)
	if networkV1provider != nil {
		if err := initRegionalAPI(client, networkV1provider, newNetworkV1Impl, client.setNetworkV1); err != nil {
			return nil, err
		}
	} else {
		setUnavailableRegionalAPI(newNetworkV1Unavailable, client.setNetworkV1)
	}

	return client, nil
}

func initRegionalAPI[T any](client *RegionalClient, provider *schema.Provider, newFunc func(client *RegionalClient, url string) (T, error), setFunc func(T)) error {
	api, err := newFunc(client, provider.Url)
	if err != nil {
		return err
	}

	setFunc(api)
	return nil
}

func setUnavailableRegionalAPI[T any](newFunc func() T, setFunc func(T)) {
	api := newFunc()
	setFunc(api)
}

func findRegionalProvider(name, version string, region *schema.Region) *schema.Provider {
	for _, provider := range region.Spec.Providers {
		if provider.Name == name && provider.Version == version {
			return &provider
		}
	}

	return nil
}

func (client *RegionalClient) setComputeV1(compute ComputeV1) {
	client.ComputeV1 = compute
}

func (client *RegionalClient) setNetworkV1(network NetworkV1) {
	client.NetworkV1 = network
}

func (client *RegionalClient) setStorageV1(storage StorageV1) {
	client.StorageV1 = storage
}

func (client *RegionalClient) setWorkspaceV1(workspace WorkspaceV1) {
	client.WorkspaceV1 = workspace
}
