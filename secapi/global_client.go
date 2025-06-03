package secapi

import (
	"context"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type GlobalClient struct {
	region region.ClientWithResponsesInterface
}

func NewGlobalClient(regionsUrl string) (*GlobalClient, error) {
	client := &GlobalClient{}

	regionClient, err := region.NewClientWithResponses(regionsUrl)
	if err != nil {
		return nil, err
	}
	client.region = regionClient

	return client, nil
}

func (client *GlobalClient) NewRegionalClient(ctx context.Context, name string) (*RegionalClient, error) {
	region, err := client.GetRegion(ctx, name)
	if err != nil {
		return nil, err
	}

	return NewRegionalClient(region), nil
}
