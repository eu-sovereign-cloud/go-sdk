package secapi

import (
	"context"
)

type GlobalEndpoints struct {
	RegionV1        string
	AuthorizationV1 string
}

type GlobalClient struct {
	RegionV1        *RegionV1
	AuthorizationV1 *AuthorizationV1
}

func NewGlobalClient(endpoints *GlobalEndpoints) (*GlobalClient, error) {
	client := &GlobalClient{}

	if endpoints == nil {
		return client, nil
	}

	// Initializes regionsV1 API client
	if endpoints.RegionV1 != "" {
		if err := initGlobalAPI(endpoints.RegionV1, newRegionV1, client.setRegionV1); err != nil {
			return nil, err
		}
	}

	// Initializes authorizationV1 API client
	if endpoints.AuthorizationV1 != "" {
		if err := initGlobalAPI(endpoints.AuthorizationV1, newAuthorizationV1, client.setAuthorizationV1); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func (client *GlobalClient) NewRegionalClient(ctx context.Context, name string, regionalAPIs []RegionalAPI) (*RegionalClient, error) {
	if client.RegionV1 == nil {
		return nil, ErrRegionRequiredToRegionalClient
	}

	region, err := client.RegionV1.GetRegion(ctx, name)
	if err != nil {
		return nil, err
	}

	return NewRegionalClient(region, regionalAPIs)
}

func initGlobalAPI[T any](endpoint string, newFunc func(url string) (*T, error), setFunc func(*T)) error {
	client, err := newFunc(endpoint)
	if err != nil {
		return err
	}

	setFunc(client)
	return nil
}

func (client *GlobalClient) setRegionV1(region *RegionV1) {
	client.RegionV1 = region
}

func (client *GlobalClient) setAuthorizationV1(authorization *AuthorizationV1) {
	client.AuthorizationV1 = authorization
}
