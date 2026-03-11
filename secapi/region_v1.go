package secapi

import (
	"context"

	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Interface

type RegionV1 interface {
	ListRegions(ctx context.Context, filter *GlobalFilter) (*Iterator[schema.Region], error)

	GetRegion(ctx context.Context, name string) (*schema.Region, error)
}

// Unavailable

type RegionV1Unavailable struct{}

func newRegionV1Unavailable() RegionV1 {
	return &RegionV1Unavailable{}
}

func (api *RegionV1Unavailable) ListRegions(ctx context.Context, filter *GlobalFilter) (*Iterator[schema.Region], error) {
	return nil, ErrProviderNotAvailable
}

func (api *RegionV1Unavailable) GetRegion(ctx context.Context, name string) (*schema.Region, error) {
	return nil, ErrProviderNotAvailable
}

// Impl

type RegionV1Impl struct {
	API
	region region.ClientWithResponsesInterface
}

func newRegionV1Impl(client *GlobalClient, regionsUrl string) (RegionV1, error) {
	region, err := region.NewClientWithResponses(regionsUrl)
	if err != nil {
		return nil, err
	}

	return &RegionV1Impl{API: API{authToken: client.authToken}, region: region}, nil
}

func (api *RegionV1Impl) ListRegions(ctx context.Context, filter *GlobalFilter) (*Iterator[schema.Region], error) {
	iter := Iterator[schema.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *string, error) {
			var params *region.ListRegionsParams
			if filter != nil && filter.Options != nil {
				params = &region.ListRegionsParams{
					Accept:    AcceptHeaderJson[region.ListRegionsParamsAccept](),
					Labels:    filter.Options.Labels.BuildPtr(),
					Limit:     filter.Options.Limit,
					SkipToken: skipToken,
				}
			} else {
				params = &region.ListRegionsParams{
					Accept:    AcceptHeaderJson[region.ListRegionsParamsAccept](),
					SkipToken: skipToken,
				}
			}

			resp, err := api.region.ListRegionsWithResponse(ctx, params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}
	return &iter, nil
}

func (api *RegionV1Impl) GetRegion(ctx context.Context, name string) (*schema.Region, error) {
	resp, err := api.region.GetRegionWithResponse(ctx, name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}
