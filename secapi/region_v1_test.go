package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/mock/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListRegions(t *testing.T) {
	ctx := context.Background()

	wkSim := mockwellknown.NewMockServerInterface(t)
	wkSim.EXPECT().GetWellknown(mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
			{
				"version": "v1",
				"endpoints": [
					{
						"provider": "seca.region/v1",
						"url": "http://` + r.Host + `/providers/seca.regions"
					}
				]
			}
		`))
	})

	reSim := mockregion.NewMockServerInterface(t)
	reSim.EXPECT().ListRegions(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, lrp region.ListRegionsParams) {
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

	sm := http.NewServeMux()
	wellknown.HandlerWithOptions(wkSim, wellknown.StdHTTPServerOptions{
		BaseURL:    "/.wellknown/secapi",
		BaseRouter: sm,
	})
	region.HandlerWithOptions(reSim, region.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.regions",
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(server.URL+"/.wellknown/secapi", nil)
	require.NoError(t, err)

	regionIter, err := client.RegionV1.ListRegions(ctx)
	require.NoError(t, err)

	region, err := regionIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, region, 1)

	assert.Len(t, region[0].Spec.Providers, 1)
	assert.Equal(t, "seca.network", region[0].Spec.Providers[0].Name)
}

func TestGetRegion(t *testing.T) {	
	ctx := context.Background()

	wkSim := mockwellknown.NewMockServerInterface(t)
	wkSim.EXPECT().GetWellknown(mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
			{
				"version": "v1",
				"endpoints": [
					{
						"provider": "seca.region/v1",
						"url": "http://` + r.Host + `/providers/seca.regions"
					}
				]
			}
		`))
	})

	reSim := mockregion.NewMockServerInterface(t)
	reSim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, name string) {
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

	sm := http.NewServeMux()
	wellknown.HandlerWithOptions(wkSim, wellknown.StdHTTPServerOptions{
		BaseURL:    "/.wellknown/secapi",
		BaseRouter: sm,
	})
	region.HandlerWithOptions(reSim, region.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.regions",
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(server.URL+"/.wellknown/secapi", nil)
	require.NoError(t, err)

	region, err := client.RegionV1.GetRegion(ctx, "test")
	require.NoError(t, err)

	assert.Len(t, region.Spec.Providers, 1)
	assert.Equal(t, "seca.network", region.Spec.Providers[0].Name)
}
