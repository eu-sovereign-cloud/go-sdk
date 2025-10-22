package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi/builders"

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
		*buildResponseRegion(secatest.Region2Name, spec),
	})
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	iter, err := client.RegionV1.ListRegions(ctx)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 2)

	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Name)

	assert.Len(t, resp[0].Spec.Providers, 1)
	assert.Contains(t, resp[0].Spec.Providers[0].Name, secatest.ProviderNetworkName)
	assert.Contains(t, resp[0].Spec.Providers[0].Url, secatest.ProviderNetworkEndpoint)
	assert.Equal(t, secatest.ProviderVersion1, resp[0].Spec.Providers[0].Version)
}

func TestListRegionsWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockListRegionsV1(sim, []schema.Region{
		{
			Metadata: &schema.GlobalResourceMetadata{
				Name: secatest.Region1Name,
			},
			Spec: schema.RegionSpec{
				Providers: []schema.Provider{
					{
						Name:    secatest.ProviderNetworkName,
						Url:     secatest.ProviderNetworkEndpoint,
						Version: secatest.ProviderVersion1,
					},
				},
			},
		},
		{
			Metadata: &schema.GlobalResourceMetadata{
				Name: secatest.Region2Name,
			},
			Spec: schema.RegionSpec{
				Providers: []schema.Provider{
					{
						Name:    secatest.ProviderNetworkName,
						Url:     secatest.ProviderNetworkEndpoint,
						Version: secatest.ProviderVersion1,
					},
				},
			},
		},
	})
	secatest.ConfigureRegionHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)
	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	ListOptions := builders.NewListOptions().WithLimit(10).WithLabels(labelsParams)
	iter, err := client.RegionV1.ListRegionsWithFilters(ctx, ListOptions)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)

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
