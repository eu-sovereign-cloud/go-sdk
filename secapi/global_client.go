package secapi

import (
	"context"
	"fmt"
)

type GlobalConfig struct {
	AuthToken string
	Endpoints GlobalEndpoints
}

type GlobalEndpoints struct {
	RegionV1        string
	AuthorizationV1 string

	WellknownV1 string
}

type GlobalClient struct {
	authToken string

	RegionV1        RegionV1
	AuthorizationV1 AuthorizationV1

	WellknownV1 WellknownV1
}

func NewGlobalClient(config *GlobalConfig) (*GlobalClient, error) {
	if config == nil {
		return nil, fmt.Errorf("GlobalConfig is required to create a global client")
	}
	if config.AuthToken == "" {
		return nil, fmt.Errorf("AuthToken is required to create a global client")
	}

	client := &GlobalClient{
		authToken: config.AuthToken,
	}

	// Initializes regionsV1 API client
	if config.Endpoints.RegionV1 != "" {
		if err := initGlobalAPIImpl(client, config.Endpoints.RegionV1, newRegionV1Impl, client.setRegionV1); err != nil {
			return nil, err
		}
	} else {
		initGlobalAPIDummy(newRegionV1Dummy, client.setRegionV1)
	}

	// Initializes authorizationV1 API client
	if config.Endpoints.AuthorizationV1 != "" {
		if err := initGlobalAPIImpl(client, config.Endpoints.AuthorizationV1, newAuthorizationV1Impl, client.setAuthorizationV1); err != nil {
			return nil, err
		}
	} else {
		initGlobalAPIDummy(newAuthorizationV1Dummy, client.setAuthorizationV1)
	}

	// Initializes wellknownV1 API client
	if config.Endpoints.WellknownV1 != "" {
		if err := initGlobalAPIImpl(client, config.Endpoints.WellknownV1, newWellknownV1Impl, client.setWellknownV1); err != nil {
			return nil, err
		}
	} else {
		initGlobalAPIDummy(newWellknownV1Dummy, client.setWellknownV1)
	}

	return client, nil
}

func (client *GlobalClient) NewRegionalClient(ctx context.Context, name string) (*RegionalClient, error) {
	region, err := client.RegionV1.GetRegion(ctx, name)
	if err != nil {
		return nil, err
	}
	if region == nil {
		return nil, fmt.Errorf("region %s not found in the regions provider", name)
	}

	return newRegionalClient(client.authToken, region)
}

func initGlobalAPIImpl[T any](client *GlobalClient, endpoint string, newFunc func(client *GlobalClient, url string) (T, error), setFunc func(T)) error {
	api, err := newFunc(client, endpoint)
	if err != nil {
		return err
	}

	setFunc(api)
	return nil
}

func initGlobalAPIDummy[T any](newFunc func() T, setFunc func(T)) {
	api := newFunc()
	setFunc(api)
}

func (client *GlobalClient) setRegionV1(region RegionV1) {
	client.RegionV1 = region
}

func (client *GlobalClient) setAuthorizationV1(authorization AuthorizationV1) {
	client.AuthorizationV1 = authorization
}

func (client *GlobalClient) setWellknownV1(wellknown WellknownV1) {
	client.WellknownV1 = wellknown
}
