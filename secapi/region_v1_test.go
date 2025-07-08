package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Region

func TestListRegionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockListRegionsV1(sim, secatest.RegionResponseV1{
		Name: secatest.Region1Name,
		Providers: []secatest.RegionResponseProviderV1{
			{
				Name: secatest.ProviderNetworkName,
				URL:  secatest.ProviderNetworkEndpoint,
			},
		},
	})
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionIter, err := client.RegionV1.ListRegions(ctx)
	require.NoError(t, err)

	region, err := regionIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, region, 1)

	assert.Equal(t, secatest.Region1Name, region[0].Metadata.Name)
	assert.Len(t, region[0].Spec.Providers, 1)
	assert.Equal(t, secatest.ProviderNetworkName, region[0].Spec.Providers[0].Name)
}

func TestGetRegionV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.RegionResponseV1{
		Name: secatest.Region1Name,
		Providers: []secatest.RegionResponseProviderV1{
			{
				Name: secatest.ProviderNetworkName,
				URL:  secatest.ProviderNetworkEndpoint,
			},
		},
	})
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	region, err := client.RegionV1.GetRegion(ctx, secatest.Region1Name)
	require.NoError(t, err)

	assert.Len(t, region.Spec.Providers, 1)
	assert.Equal(t, secatest.ProviderNetworkName, region.Spec.Providers[0].Name)
}
