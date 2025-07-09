package secapi

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"

	"github.com/stretchr/testify/require"
)

func getTestRegionalClient(t *testing.T, ctx context.Context, server *httptest.Server) *RegionalClient {
	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, secatest.Region1Name)
	require.NoError(t, err)
	return regionalClient
}
