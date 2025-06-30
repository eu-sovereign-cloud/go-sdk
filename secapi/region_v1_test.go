package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	"github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRegionsV1(t *testing.T) {
	ctx := context.Background()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockListRegionsV1(sim, secatest.ListRegionsResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.ListRegionsResponseProviderV1{
			{
				Name: secatest.ProviderNetworkName,
				URL:  secatest.ProviderNetworkEndpoint,
			},
		},
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionIter, err := client.RegionV1.ListRegions(ctx)
	require.NoError(t, err)

	region, err := regionIter.All(ctx)
	require.NoError(t, err)

	require.Len(t, region, 1)

	assert.Equal(t, secatest.RegionName, region[0].Metadata.Name)
	assert.Len(t, region[0].Spec.Providers, 1)
	assert.Equal(t, secatest.ProviderNetworkName, region[0].Spec.Providers[0].Name)
}

func TestGetRegionV1(t *testing.T) {
	ctx := context.Background()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderNetworkName,
				URL:  secatest.ProviderNetworkEndpoint,
			},
		},
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	region, err := client.RegionV1.GetRegion(ctx, "test")
	require.NoError(t, err)

	assert.Len(t, region.Spec.Providers, 1)
	assert.Equal(t, secatest.ProviderNetworkName, region.Spec.Providers[0].Name)
}
