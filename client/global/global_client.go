package global

import (
	"context"

	"github.com/eu-sovereign-cloud/go-sdk/client/regional"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type GlobalClient struct {
	regionClient        region.ClientWithResponsesInterface
}

func NewGlobalClient(regionsUrl string) (*GlobalClient, error) {
	client := &GlobalClient{}

	regionClient, err := region.NewClientWithResponses(regionsUrl)
	if err != nil {
		return nil, err
	}
	client.regionClient = regionClient

	return client, nil
}

func (gc *GlobalClient) NewRegionalClient(ctx context.Context, name string) (*regional.RegionalClient, error) {
	region, err := gc.GetRegion(ctx, name)
	if err != nil {
		return nil, err
	}

	return regional.NewRegionalClient(region), nil
}
