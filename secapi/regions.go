package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

func (client *GlobalClient) ListRegions(ctx context.Context) (*Iterator[region.Region], error) {
	iter := Iterator[region.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]region.Region, *string, error) {
			resp, err := client.regions.ListRegionsWithResponse(ctx, &region.ListRegionsParams{
				Accept: ptr.To(region.ListRegionsParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (client *GlobalClient) GetRegion(ctx context.Context, name string) (*region.Region, error) {
	resp, err := client.regions.GetRegionWithResponse(ctx, name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (client *GlobalClient) RegionClient(ctx context.Context, name string) (*RegionalClient, error) {
	region, err := client.GetRegion(ctx, name)
	if err != nil {
		return nil, err
	}

	return NewRegionalClient(region), nil
}
