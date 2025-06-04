package secapi

import (
	"sync"
)

type GlobalAPI[T any] struct {
	Client *T

	mu sync.Mutex
}

func (api *GlobalAPI[T]) GetClient(name string, fn func(string) (T, error)) (*T, error) {
	if api.Client != nil {
		return api.Client, nil
	}

	api.mu.Lock()
	defer api.mu.Unlock()
	if api.Client != nil {
		return api.Client, nil
	}

	// TODO Find how validate the global provider
	providerUrl := name

	client, err := fn(providerUrl)
	if err != nil {
		return nil, err
	}
	api.Client = &client

	return &client, nil
}
