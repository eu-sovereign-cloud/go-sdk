package secapi

import (
	"fmt"
	"sync"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type RegionalAPI[T any] struct {
	Region *region.Region

	Client *T

	mu sync.Mutex
}

func (api *RegionalAPI[T]) GetClient(name string, fn func(string) (T, error)) (*T, error) {
	if api.Client != nil {
		return api.Client, nil
	}

	api.mu.Lock()
	defer api.mu.Unlock()
	if api.Client != nil {
		return api.Client, nil
	}

	provider, err := findProvider(name, api.Region)
	if err != nil {
		return nil, err
	}

	client, err := fn(provider.Url)
	if err != nil {
		return nil, err
	}
	api.Client = &client

	return &client, nil
}

func findProvider(name string, region *region.Region) (*region.Provider, error) {
	for _, provider := range region.Spec.Providers {
		if provider.Name == name {
			return &provider, nil
		}
	}

	return nil, fmt.Errorf("provider %s not found in region", name)
}
