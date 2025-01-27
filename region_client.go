package gosdk

import (
	"fmt"
	"sync"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/regions.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/workspace.v1"
)

type RegionClient struct {
	region *regions.Region

	workspace workspace.ClientWithResponsesInterface

	mu sync.Mutex
}

func NewRegionClient(region *regions.Region) *RegionClient {
	var client RegionClient
	client.region = region
	return &client
}

func (client *RegionClient) workspaceClient() (workspace.ClientWithResponsesInterface, error) {
	if client.workspace != nil {
		return client.workspace, nil
	}

	client.mu.Lock()
	defer client.mu.Unlock()
	if client.workspace != nil {
		return client.workspace, nil
	}

	provider := client.findProvider("seca.workspace")
	if provider == nil {
		return nil, fmt.Errorf("provider workspace not found in region")
	}

	workspaceClient, err := workspace.NewClientWithResponses(provider.Url)
	if err != nil {
		return nil, err
	}
	client.workspace = workspaceClient

	return workspaceClient, nil
}

func (client *RegionClient) findProvider(name string) *regions.Provider {
	for _, provider := range client.region.Spec.Providers {
		if provider.Name == name {
			return &provider
		}
	}

	return nil
}
