package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockCompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	mockRegion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"

	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"

	"github.com/stretchr/testify/require"
)

func TestCreateOrUpdateInstance(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateInstanceV1(wsSim, secatest.CreateOrUpdateInstanceResponseV1{
		Name:      "some-workspace",
		Tenant:    "test",
		Workspace: "test-workspace",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)
	ws := "test-workspace"
	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    "test",
			Name:      "test-instance",
			Workspace: &ws,
		},
	}

	err = regionalClient.ComputeV1.CreateOrUpdateInstance(ctx, cp)
	require.NoError(t, err)

}

func TestGetInstance(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockGetInstanceV1(wsSim, secatest.GetInstanceResponseV1{
		Name:   "some-workspace",
		Tenant: "test",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)

	wref := WorkspaceReference{
		TenantReference: TenantReference{
			Tenant: "test-tenant",
			Name:   "some-workspace",
		},
		Workspace: "workspace_1",
	}
	cp, err := regionalClient.ComputeV1.GetInstance(ctx, wref)
	require.NoError(t, err)
	require.NotEmpty(t, cp)

}

func TestGetInstanceSkU(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockGetInstanceSkuV1(wsSim, secatest.GetInstanceSkuResponseV1{
		Name:   "some-workspace",
		Tenant: "test",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)

	cp, err := regionalClient.ComputeV1.GetSku(ctx, TenantReference{
		Tenant: "test-tenant",
		Name:   "some-workspace",
	})
	require.NoError(t, err)
	require.NotEmpty(t, cp)

}

// error
func TestListInstances(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockListInstancesV1(wsSim, secatest.ListInstancesResponseV1{
		Name:      "some-workspace",
		Tenant:    "test",
		Workspace: "test-workspace",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)

	cpIter, err := regionalClient.ComputeV1.ListInstances(ctx, "test", "some-workspace")
	cp, err := cpIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, cp, 1)

}
func TestListInstancesSku(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockInstanceListSkusV1(wsSim, secatest.ListInstancesSkusResponseV1{
		Name:      "some-workspace",
		Tenant:    "test",
		Workspace: "test-workspace",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)

	cpIter, err := regionalClient.ComputeV1.ListSkus(ctx, "test", "some-workspace")
	require.NoError(t, err)
	cp, err := cpIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, cp, 14)

}

func TestRestartInstanace(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockRestartInstanceV1(wsSim)

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)
	ws := "test-workspace"
	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    "test",
			Name:      "test-instance",
			Workspace: &ws,
		},
	}
	err = regionalClient.ComputeV1.RestartInstance(ctx, cp)
	require.NoError(t, err)

}

func TestStartInstanace(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockStartInstanceV1(wsSim)

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)
	ws := "test-workspace"
	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    "test",
			Name:      "test-instance",
			Workspace: &ws,
		},
	}
	err = regionalClient.ComputeV1.StartInstance(ctx, cp)
	require.NoError(t, err)

}

func TestStopInstanace(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderComputeName,
				URL:  secatest.ProviderComputeEndpoint,
			},
		},
	})
	wsSim := mockCompute.NewMockServerInterface(t)
	secatest.MockStopInstanceV1(wsSim)

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	compute.HandlerWithOptions(wsSim, compute.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderComputeEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{ComputeV1API})
	require.NoError(t, err)
	ws := "test-workspace"
	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    "test",
			Name:      "test-instance",
			Workspace: &ws,
		},
	}
	err = regionalClient.ComputeV1.StopInstance(ctx, cp)
	require.NoError(t, err)

}
