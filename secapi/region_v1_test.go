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
		Metadata: secatest.MetadataResponseV1{Name: secatest.Region1Name},
		Providers: []secatest.RegionResponseProviderV1{
			{
				Name:    secatest.ProviderNetworkName,
				URL:     secatest.ProviderNetworkEndpoint,
				Version: secatest.ProviderVersion1,
			},
		},
	})
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	iter, err := client.RegionV1.ListRegions(ctx)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Name)

	require.Len(t, resp[0].Spec.Providers, 1)
	assert.Contains(t, resp[0].Spec.Providers[0].Name, secatest.ProviderNetworkName)
	assert.Contains(t, resp[0].Spec.Providers[0].Url, secatest.ProviderNetworkEndpoint)
	assert.Equal(t, secatest.ProviderVersion1, resp[0].Spec.Providers[0].Version)
}

func TestGetRegionV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.RegionResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.Region1Name},
		Providers: []secatest.RegionResponseProviderV1{
			{
				Name:    secatest.ProviderNetworkName,
				URL:     secatest.ProviderNetworkEndpoint,
				Version: secatest.ProviderVersion1,
			},
		},
	})
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	resp, err := client.RegionV1.GetRegion(ctx, secatest.Region1Name)
	require.NoError(t, err)

	assert.Equal(t, secatest.Region1Name, resp.Metadata.Name)

	assert.Contains(t, resp.Spec.Providers[0].Name, secatest.ProviderNetworkName)
	assert.Contains(t, resp.Spec.Providers[0].Url, secatest.ProviderNetworkEndpoint)
	assert.Equal(t, secatest.ProviderVersion1, resp.Spec.Providers[0].Version)
}
