package secapi

import (
	"context"
	"fmt"
	"slices"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
)

type GlobalAPI int

const (
	AuthorizationV1API GlobalAPI = iota
)

type GlobalClient struct {
	WellknownV1 *WellknownV1
	RegionV1    *RegionV1

	AuthorizationV1 *AuthorizationV1
}

func (client *GlobalClient) setAuthorizationV1(authorization *AuthorizationV1) {
	client.AuthorizationV1 = authorization
}

func initGlobalAPI[T any](provider string, endpoints []wellknown.WellknownEndpoint, newFn func(url string) (*T, error), setFn func(*T)) error {
	url, err := findGlobalProviderUrl(provider, endpoints)
	if err != nil {
		return err
	}

	client, err := newFn(url)
	if err != nil {
		return err
	}

	setFn(client)
	return nil
}

func findGlobalProviderUrl(name string, endpoints []wellknown.WellknownEndpoint) (string, error) {
	for _, endpoint := range endpoints {
		if endpoint.Provider == name {
			return endpoint.Url, nil
		}
	}
	return "", fmt.Errorf("provider endpoint %s not found", name)
}

func NewGlobalClient(wellknownUrl string, globalAPIs []GlobalAPI) (*GlobalClient, error) {
	globalClient := &GlobalClient{}

	// Initializes wellknownV1 API client
	wellknownV1, err := newWellknownV1(wellknownUrl)
	if err != nil {
		return nil, err
	}
	globalClient.WellknownV1 = wellknownV1

	wellknown, err := wellknownV1.GetWellknown(context.Background())
	if err != nil {
		return nil, err
	}

	// Initializes regionsV1 API client
	regionsUrl, err := findGlobalProviderUrl("seca.region/v1", wellknown.Endpoints)
	if err != nil {
		return nil, err
	}
	regionV1, err := newRegionV1(regionsUrl)
	if err != nil {
		return nil, err
	}
	globalClient.RegionV1 = regionV1

	// Initializes authorizationV1 API client
	if found := slices.Contains(globalAPIs, AuthorizationV1API); found {
		err := initGlobalAPI("seca.authorization/v1", wellknown.Endpoints, newAuthorizationV1, globalClient.setAuthorizationV1)
		if err != nil {
			return nil, err
		}
	}

	return globalClient, nil
}

func (client *GlobalClient) NewRegionalClient(ctx context.Context, name string, regionalAPIs []RegionalAPI) (*RegionalClient, error) {
	region, err := client.RegionV1.GetRegion(ctx, name)
	if err != nil {
		return nil, err
	}

	return NewRegionalClient(region, regionalAPIs)
}
