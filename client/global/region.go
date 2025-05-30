package global

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/client"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

func (gc *GlobalClient) ListRegions(ctx context.Context) (*client.Iterator[region.Region], error) {

	iter := client.Iterator[region.Region]{
		Func: func(ctx context.Context, skipToken *string) ([]region.Region, *string, error) {
			resp, err := gc.regionClient.ListRegionsWithResponse(ctx, &region.ListRegionsParams{
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

func (gc *GlobalClient) GetRegion(ctx context.Context, name string) (*region.Region, error) {
	resp, err := gc.regionClient.GetRegionWithResponse(ctx, name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}
