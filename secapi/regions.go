package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/foundation.region.v1"
)

func (c *Client) Regions(ctx context.Context) (*Iterator[region.Region], error) {
	iter := Iterator[region.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]region.Region, *string, error) {
			resp, err := c.regions.ListRegionsWithResponse(ctx, &region.ListRegionsParams{
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

func (c *Client) Region(ctx context.Context, name string) (*region.Region, error) {
	resp, err := c.regions.GetRegionWithResponse(ctx, name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (c *Client) RegionClient(ctx context.Context, name string) (*RegionClient, error) {
	region, err := c.Region(ctx, name)
	if err != nil {
		return nil, err
	}

	return NewRegionClient(region), nil
}
