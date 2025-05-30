package regional

import (
	"fmt"
	"sync"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
)

type RegionalClient struct {
	region *region.Region

	workspaceClient     workspace.ClientWithResponsesInterface

	mu sync.Mutex
}

func NewRegionalClient(region *region.Region) *RegionalClient {
	var client RegionalClient
	client.region = region
	return &client
}

func (client *RegionalClient) findProvider(name string) *region.Provider {
	for _, provider := range client.region.Spec.Providers {
		if provider.Name == name {
			return &provider
		}
	}

	return nil
}

func (client *RegionalClient) getWorkspaceClient() (workspace.ClientWithResponsesInterface, error) {
	if client.workspaceClient != nil {
		return client.workspaceClient, nil
	}

	client.mu.Lock()
	defer client.mu.Unlock()
	if client.workspaceClient != nil {
		return client.workspaceClient, nil
	}

	provider := client.findProvider("seca.workspace")
	if provider == nil {
		return nil, fmt.Errorf("provider workspace not found in region")
	}

	workspaceClient, err := workspace.NewClientWithResponses(provider.Url)
	if err != nil {
		return nil, err
	}
	client.workspaceClient = workspaceClient

	return workspaceClient, nil
}
