package secapi

import (
	"fmt"
	"slices"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type RegionalAPI int

const (
	ComputeV1API RegionalAPI = iota
	NetworkV1API
	StorageV1API
	WorkspaceV1API
)

type RegionalClient struct {
	ComputeV1   *ComputeV1
	NetworkV1   *NetworkV1
	StorageV1   *StorageV1
	WorkspaceV1 *WorkspaceV1
}

func (client *RegionalClient) setComputeV1(compute *ComputeV1) {
	client.ComputeV1 = compute
}

func (client *RegionalClient) setNetworkV1(network *NetworkV1) {
	client.NetworkV1 = network
}

func (client *RegionalClient) setStorageV1(storage *StorageV1) {
	client.StorageV1 = storage
}

func (client *RegionalClient) setWorkspaceV1(workspace *WorkspaceV1) {
	client.WorkspaceV1 = workspace
}

func initRegionalAPI[T any](name string, region *region.Region, newFn func(url string) (*T, error), setFn func(*T)) error {
	provider, err := findRegionalProvider(name, region)
	if err != nil {
		return err
	}

	client, err := newFn(provider.Url)
	if err != nil {
		return err
	}

	setFn(client)
	return nil
}

func findRegionalProvider(name string, region *region.Region) (*region.Provider, error) {
	for _, provider := range region.Spec.Providers {
		if provider.Name == name {
			return &provider, nil
		}
	}

	return nil, fmt.Errorf("provider %s not found in region", name)
}

func NewRegionalClient(region *region.Region, regionalAPIs []RegionalAPI) (*RegionalClient, error) {
	regionalClient := &RegionalClient{}

	// Initializes computeV1 API client
	if found := slices.Contains(regionalAPIs, ComputeV1API); found {
		err := initRegionalAPI("seca.compute", region, newComputeV1, regionalClient.setComputeV1)
		if err != nil {
			return nil, err
		}
	}

	// Initializes networkV1 API client
	if found := slices.Contains(regionalAPIs, NetworkV1API); found {
		err := initRegionalAPI("seca.network", region, newNetworkV1, regionalClient.setNetworkV1)
		if err != nil {
			return nil, err
		}
	}

	// Initializes storageV1 API client
	if found := slices.Contains(regionalAPIs, StorageV1API); found {
		err := initRegionalAPI("seca.storage", region, newStorageV1, regionalClient.setStorageV1)
		if err != nil {
			return nil, err
		}
	}

	// Initializes workspaceV1 API client
	if found := slices.Contains(regionalAPIs, WorkspaceV1API); found {
		err := initRegionalAPI("seca.workspace", region, newWorkspaceV1, regionalClient.setWorkspaceV1)
		if err != nil {
			return nil, err
		}
	}

	return regionalClient, nil
}
