package gosdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/mock/mockregion.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/region.v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegions(t *testing.T) {
	sim := mockregion.NewMockServerInterface(t)
	ctx := context.Background()

	sim.EXPECT().ListRegions(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, lrp region.ListRegionsParams) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(` 
			{
				"items": [
					{
						"metadata": {
							"apiVersion": "v1",
							"kind": "region",
							"name": "eu-central-1"
						},
						"spec": {
							"availableZones": [ "A", "B" ],
							"providers": [
								{
									"name": "seca.network",
									"version": "v1",
									"url": "https://demo.secapi.cloud/providers/seca.network"
								}
							]
						},
						"status": {
							"conditions": [ { "status": "Ready" } ]
						}
					}
				],
				"metadata": {
					"skipToken": null
				}
			}
		`))
	})

	server := httptest.NewServer(region.HandlerWithOptions(sim, region.StdHTTPServerOptions{}))
	defer server.Close()

	client, err := NewClient(server.URL)
	require.NoError(t, err)

	regionIter, err := client.Regions(ctx)
	require.NoError(t, err)

	region, err := regionIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, region, 1)

	assert.Len(t, region[0].Spec.Providers, 1)
	assert.Equal(t, "seca.network", region[0].Spec.Providers[0].Name)
}

func TestRegion(t *testing.T) {
	sim := mockregion.NewMockServerInterface(t)
	ctx := context.Background()

	sim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, name string) {
		assert.Equal(t, "test", name)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
			{
				"metadata": {
					"apiVersion": "v1",
					"kind": "region",
					"name": "eu-central-1"
				},
				"spec": {
					"availableZones": [ "A", "B" ],
					"providers": [
						{
							"name": "seca.network",
							"version": "v1",
							"url": "https://demo.secapi.cloud/providers/seca.network"
						}
					]
				},
				"status": {
					"conditions": [ { "status": "Ready" } ]
				}
			}
		`))
	})

	server := httptest.NewServer(region.HandlerWithOptions(sim, region.StdHTTPServerOptions{}))
	defer server.Close()

	client, err := NewClient(server.URL)
	require.NoError(t, err)

	region, err := client.Region(ctx, "test")
	require.NoError(t, err)

	assert.Len(t, region.Spec.Providers, 1)
	assert.Equal(t, "seca.network", region.Spec.Providers[0].Name)
}
