package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	"github.com/eu-sovereign-cloud/go-sdk/mock/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRegionsV1(t *testing.T) {
	ctx := context.Background()

	wkSim := mockwellknown.NewMockServerInterface(t)
	secatest.MockGetWellknownV1(wkSim, secatest.GetWellknownResponseV1{
		Endpoints: []secatest.GetWellknownResponseEndpointV1{
			{
				Provider: secatest.ProviderRegionName,
				URL:      secatest.ProviderRegionEndpoint,
			},
		},
	})

	reSim := mockregion.NewMockServerInterface(t)
	secatest.MockListRegionsV1(reSim, secatest.ListRegionsResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.ListRegionsResponseProviderV1{
			{
				Name: secatest.ProviderNetworkName,
				URL:  secatest.ProviderNetworkEndpoint,
			},
		},
	})

	sm := http.NewServeMux()
	wellknown.HandlerWithOptions(wkSim, wellknown.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderWellknownEndpoint,
		BaseRouter: sm,
	})
	region.HandlerWithOptions(reSim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(server.URL+secatest.ProviderWellknownEndpoint, nil)
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

	wkSim := mockwellknown.NewMockServerInterface(t)
	secatest.MockGetWellknownV1(wkSim, secatest.GetWellknownResponseV1{
		Endpoints: []secatest.GetWellknownResponseEndpointV1{
			{
				Provider: secatest.ProviderRegionName,
				URL:      secatest.ProviderRegionEndpoint,
			},
		},
	})

	reSim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(reSim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderNetworkName,
				URL:  secatest.ProviderNetworkEndpoint,
			},
		},
	})

	sm := http.NewServeMux()
	wellknown.HandlerWithOptions(wkSim, wellknown.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderWellknownEndpoint,
		BaseRouter: sm,
	})
	region.HandlerWithOptions(reSim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(server.URL+secatest.ProviderWellknownEndpoint, nil)
	require.NoError(t, err)

	region, err := client.RegionV1.GetRegion(ctx, "test")
	require.NoError(t, err)

	assert.Len(t, region.Spec.Providers, 1)
	assert.Equal(t, secatest.ProviderNetworkName, region.Spec.Providers[0].Name)
}
