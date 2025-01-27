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

	sim.EXPECT().ListRegions(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, lrp region.ListRegionsParams) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"items": [
					{
						"apiVersion": "v1",
						"kind": "region",
						"metadata": {
							"name": "primary-load-balancer",
							"deletionTimestamp": "2019-08-24T14:15:22Z",
							"lastModifiedTimestamp": "2019-08-24T14:15:22Z",
							"description": "string",
							"labels": {
								"language": "en",
								"billing.cost-center": "platform-eng",
								"env": "production"
							}
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
							"conditions": [
								{
									"type": "power-mgmt",
									"status": "True, false, unknown",
									"lastTransitionTime": "2019-08-24T14:15:22Z",
									"reason": "^A(A)?$",
									"message": "string"
								}
							]
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

	ctx = WithTenantID(ctx, "test")

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

	sim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, name string) {
		assert.Equal(t, "test", name)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"apiVersion": "v1",
				"kind": "region",
				"metadata": {
					"name": "primary-load-balancer",
					"deletionTimestamp": "2019-08-24T14:15:22Z",
					"lastModifiedTimestamp": "2019-08-24T14:15:22Z",
					"description": "string",
					"labels": {
						"language": "en",
						"billing.cost-center": "platform-eng",
						"env": "production"
					}
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
					"conditions": [
						{
							"type": "power-mgmt",
							"status": "True, false, unknown",
							"lastTransitionTime": "2019-08-24T14:15:22Z",
							"reason": "^A(A)?$",
							"message": "string"
						}
					]
				}
			}
		`))
	})

	server := httptest.NewServer(region.HandlerWithOptions(sim, region.StdHTTPServerOptions{}))
	defer server.Close()

	client, err := NewClient(server.URL)
	require.NoError(t, err)

	ctx = WithTenantID(ctx, "test")

	region, err := client.Region(ctx, "test")
	require.NoError(t, err)

	assert.Len(t, region.Spec.Providers, 1)
	assert.Equal(t, "seca.network", region.Spec.Providers[0].Name)
}
