package secapi

import (
	"fmt"
	"sync"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/foundation.workspace.v1"
)

type RegionalClient struct {
	region *region.Region

	workspace workspace.ClientWithResponsesInterface

	mu sync.Mutex
}

func NewRegionClient(region *region.Region) *RegionalClient {
	var client RegionalClient
	client.region = region
	return &client
}

func (client *RegionalClient) workspaceClient() (workspace.ClientWithResponsesInterface, error) {
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

func (client *RegionalClient) findProvider(name string) *region.Provider {
	for _, provider := range client.region.Spec.Providers {
		if provider.Name == name {
			return &provider
		}
	}

	return nil
}
