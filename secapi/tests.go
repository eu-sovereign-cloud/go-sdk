package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	"github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"

	"github.com/stretchr/testify/require"
)

func getTestRegionalClient(t *testing.T, ctx context.Context, apis []RegionalAPI, server *httptest.Server) *RegionalClient {

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, secatest.Region1Name, apis)
	require.NoError(t, err)
	return regionalClient
}

func initTestRegionV1Handler(t *testing.T, sm *http.ServeMux) *mockregion.MockServerInterface {

	sim := mockregion.NewMockServerInterface(t)

	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.Region1Name,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderNetworkName,
				URL:  secatest.ProviderNetworkEndpoint,
			},
		},
	})
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})

	return sim
}
