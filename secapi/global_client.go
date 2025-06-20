package secapi

import (
	"context"
)

type GlobalClient struct {
	AuthorizationV1 *AuthorizationV1
	RegionV1        *RegionV1
	WellknownV1     *WellknownV1
}

func NewGlobalClient(regionsUrl string) (*GlobalClient, error) {
	regionV1, err := newRegionV1(regionsUrl)
	if err != nil {
		return nil, err
	}

	return &GlobalClient{
		AuthorizationV1: newAuthorizationV1(),
		RegionV1:        regionV1,
		WellknownV1:     newWellknownV1(),
	}, nil
}

func (client *GlobalClient) NewRegionalClient(ctx context.Context, name string) (*RegionalClient, error) {
	region, err := client.RegionV1.GetRegion(ctx, name)
	if err != nil {
		return nil, err
	}

	return NewRegionalClient(region), nil
}
