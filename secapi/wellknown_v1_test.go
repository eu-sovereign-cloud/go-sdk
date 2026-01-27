package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockwellknown "github.com/eu-sovereign-cloud/go-sdk/mock/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/constants"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/assert"
)

// Wellknown

func TestGetWellknownV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockwellknown.NewMockServerInterface(t)

	endpoints := []schema.WellknownEndpoint{
		buildResponseWellknownEndpoint(constants.RegionProviderV1Name, secatest.ProviderRegionV1Endpoint),
		buildResponseWellknownEndpoint(constants.AuthorizationProviderV1Name, secatest.ProviderAuthorizationV1Endpoint),
	}
	secatest.MockGetWellknownV1(sim, buildResponseWellknown(constants.ApiVersion1, endpoints))

	secatest.ConfigureWellknownHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	resp, err := client.WellknownV1.GetWellknown(ctx)
	assert.NoError(t, err)

	assert.Equal(t, constants.ApiVersion1, resp.Version)

	assert.Contains(t, resp.Endpoints[0].Provider, constants.RegionProviderV1Name)
	assert.Contains(t, resp.Endpoints[0].Url, secatest.ProviderRegionV1Endpoint)

	assert.Contains(t, resp.Endpoints[1].Provider, constants.AuthorizationProviderV1Name)
	assert.Contains(t, resp.Endpoints[1].Url, secatest.ProviderAuthorizationV1Endpoint)
}

// Builders

func buildResponseWellknown(version string, endpoints []schema.WellknownEndpoint) *schema.Wellknown {
	return &schema.Wellknown{
		Version:   version,
		Endpoints: endpoints,
	}
}

func buildResponseWellknownEndpoint(providerName, providerUrl string) schema.WellknownEndpoint {
	return schema.WellknownEndpoint{
		Provider: providerName,
		Url:      providerUrl,
	}
}
