package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type RegionV1 struct {
	region region.ClientWithResponsesInterface
}

func newRegionV1(regionsUrl string) (*RegionV1, error) {
	region, err := region.NewClientWithResponses(regionsUrl)
	if err != nil {
		return nil, err
	}

	return &RegionV1{region: region}, nil
}

func (api *RegionV1) ListRegions(ctx context.Context) (*Iterator[region.Region], error) {
	iter := Iterator[region.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]region.Region, *string, error) {
			resp, err := api.region.ListRegionsWithResponse(ctx, &region.ListRegionsParams{
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

func (api *RegionV1) GetRegion(ctx context.Context, name string) (*region.Region, error) {
	resp, err := api.region.GetRegionWithResponse(ctx, name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}
