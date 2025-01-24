package gosdk

import (
	"context"

	"k8s.io/utils/ptr"

	regions "github.com/eu-sovereign-cloud/go-sdk/pkg/regions.v1"
)

func (c *Client) Regions(ctx context.Context) (*Iterator[regions.Region], error) {
	tid := regions.TenantID("some-tenant-id")

	iter := Iterator[regions.Region]{
		fn: func(ctx context.Context, skipToken *string) ([]regions.Region, *string, error) {
			resp, err := c.regions.ListRegionsWithResponse(ctx, tid, &regions.ListRegionsParams{
				Accept: ptr.To(regions.ListRegionsParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}
