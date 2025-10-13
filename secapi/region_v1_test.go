package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/assert"
)

// Region

func TestListRegionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockregion.NewMockServerInterface(t)
	spec := buildResponseRegionSpec(secatest.ProviderNetworkName, secatest.ProviderNetworkEndpoint, secatest.ProviderVersion1)
	secatest.MockListRegionsV1(sim, []schema.Region{
		*buildResponseRegion(secatest.Region1Name, spec),
	})
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	iter, err := client.RegionV1.ListRegions(ctx)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Name)

	assert.Len(t, resp[0].Spec.Providers, 1)
	assert.Contains(t, resp[0].Spec.Providers[0].Name, secatest.ProviderNetworkName)
	assert.Contains(t, resp[0].Spec.Providers[0].Url, secatest.ProviderNetworkEndpoint)
	assert.Equal(t, secatest.ProviderVersion1, resp[0].Spec.Providers[0].Version)
}

func TestGetRegionV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockregion.NewMockServerInterface(t)
	spec := buildResponseRegionSpec(secatest.ProviderNetworkName, secatest.ProviderNetworkEndpoint, secatest.ProviderVersion1)
	secatest.MockGetRegionV1(sim, buildResponseRegion(secatest.Region1Name, spec))
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	resp, err := client.RegionV1.GetRegion(ctx, secatest.Region1Name)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Region1Name, resp.Metadata.Name)

	assert.Contains(t, resp.Spec.Providers[0].Name, secatest.ProviderNetworkName)
	assert.Contains(t, resp.Spec.Providers[0].Url, secatest.ProviderNetworkEndpoint)
	assert.Equal(t, secatest.ProviderVersion1, resp.Spec.Providers[0].Version)
}

// Builders

func buildResponseRegion(name string, spec *schema.RegionSpec) *schema.Region {
	return &schema.Region{
		Metadata: secatest.NewGlobalResourceMetadata(name),
		Spec:     *spec,
	}
}

func buildResponseRegionSpec(providerName, providerUrl, providerVersion string) *schema.RegionSpec {
	return &schema.RegionSpec{
		Providers: []schema.Provider{
			{
				Name:    providerName,
				Url:     providerUrl,
				Version: providerVersion,
			},
		},
	}
}
