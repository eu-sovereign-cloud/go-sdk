package secapi

import (
	"context"
	"net/http"

	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"k8s.io/utils/ptr"
)

// Interface

type RegionV1 interface {
	ListRegions(ctx context.Context) (*Iterator[schema.Region], error)
	ListRegionsWithFilters(ctx context.Context, opts *ListOptions) (*Iterator[schema.Region], error)

	GetRegion(ctx context.Context, name string) (*schema.Region, error)
}

// Dummy

type RegionV1Dummy struct{}

func newRegionV1Dummy() RegionV1 {
	return &RegionV1Dummy{}
}

func (api *RegionV1Dummy) ListRegions(ctx context.Context) (*Iterator[schema.Region], error) {
	return nil, ErrProviderNotAvailable
}

func (api *RegionV1Dummy) ListRegionsWithFilters(ctx context.Context, opts *ListOptions) (*Iterator[schema.Region], error) {
	return nil, ErrProviderNotAvailable
}

func (api *RegionV1Dummy) GetRegion(ctx context.Context, name string) (*schema.Region, error) {
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

func (api *RegionV1Impl) ListRegions(ctx context.Context) (*Iterator[schema.Region], error) {
	iter := Iterator[schema.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *string, error) {
			resp, err := api.region.ListRegionsWithResponse(ctx, &region.ListRegionsParams{
				Accept:    ptr.To(region.ListRegionsParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *RegionV1Impl) ListRegionsWithFilters(ctx context.Context, opts *ListOptions) (*Iterator[schema.Region], error) {
	iter := Iterator[schema.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *string, error) {
			resp, err := api.region.ListRegionsWithResponse(ctx, &region.ListRegionsParams{
				Accept:    ptr.To(region.ListRegionsParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
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

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}
