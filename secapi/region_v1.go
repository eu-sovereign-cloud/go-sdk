package secapi

import (
	"context"

	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Interface

type RegionV1 interface {
	ListRegionsWithOptions(ctx context.Context, options *ListOptions) (*Iterator[schema.Region], error)
	ListRegions(ctx context.Context) (*Iterator[schema.Region], error)

	GetRegion(ctx context.Context, name string) (*schema.Region, error)
}

// Unavailable

type RegionV1Unavailable struct{}

func newRegionV1Unavailable() RegionV1 {
	return &RegionV1Unavailable{}
}

func (api *RegionV1Unavailable) ListRegionsWithOptions(ctx context.Context, options *ListOptions) (*Iterator[schema.Region], error) {
	return nil, ErrProviderNotAvailable
}

func (api *RegionV1Unavailable) ListRegions(ctx context.Context) (*Iterator[schema.Region], error) {
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

func (api *RegionV1Impl) ListRegionsWithOptions(ctx context.Context, options *ListOptions) (*Iterator[schema.Region], error) {
	iter := Iterator[schema.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *schema.ResponseMetadata, error) {
			var params *region.ListRegionsParams
			if options != nil {
				params = &region.ListRegionsParams{
					Accept:    AcceptHeaderJson[region.ListRegionsParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
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
				return resp.JSON200.Items, &resp.JSON200.Metadata, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}
	return &iter, nil
}

func (api *RegionV1Impl) ListRegions(ctx context.Context) (*Iterator[schema.Region], error) {
	return api.ListRegionsWithOptions(ctx, nil)
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
