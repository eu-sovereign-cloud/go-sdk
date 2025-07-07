package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/internal/fake"
	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListWorkspacesV1(t *testing.T) {
	ctx := context.Background()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderWorkspaceName,
				URL:  secatest.ProviderWorkspaceEndpoint,
			},
		},
	})

	wsSim := mockworkspace.NewMockServerInterface(t)
	secatest.MockListWorkspaceV1(wsSim, secatest.WorkspaceTypeResponseV1{
		Name:      secatest.Workspace1Name,
		Tenant:    secatest.Tenant1Name,
		Region:    secatest.Region1Name,
		Workspace: secatest.Workspace1Name,
		State:     secatest.State1Active,
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	workspace.HandlerWithOptions(wsSim, workspace.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderWorkspaceEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{WorkspaceV1API})
	require.NoError(t, err)

	wsIter, err := regionalClient.WorkspaceV1.ListWorkspaces(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	ws, err := wsIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, ws, 1)
	assert.Equal(t, secatest.Workspace1Name, ws[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, ws[0].Metadata.Tenant)
	assert.EqualValues(t, secatest.State1Active, *ws[0].Status.State)

}

func TestFakedListWorkspacesV1(t *testing.T) {
	ctx := context.Background()

	fakeServer := fake.NewServer(secatest.RegionName)
	server := fakeServer.Start()
	defer server.Close()

	fakeServer.Workspaces[secatest.Workspace1Name] = &workspace.Workspace{
		Metadata: &workspace.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Workspace1Name,
		},
		Status: &workspace.WorkspaceStatus{
			State: ptr.To(workspace.ResourceStatePending),
		},
	}

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + "/providers/seca.regions"})
	require.NoError(t, err)

	regionClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{WorkspaceV1API})
	require.NoError(t, err)

	wsIter, err := regionClient.WorkspaceV1.ListWorkspaces(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	ws, err := wsIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, ws, 1)

	assert.Equal(t, secatest.Workspace1Name, ws[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, ws[0].Metadata.Tenant)
	assert.EqualValues(t, "pending", *ws[0].Status.State)
}

func TestGetWorkspaces(t *testing.T) {
	ctx := context.Background()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderWorkspaceName,
				URL:  secatest.ProviderWorkspaceEndpoint,
			},
		},
	})

	wsSim := mockworkspace.NewMockServerInterface(t)
	secatest.MockGetWorkspaceV1(wsSim, secatest.WorkspaceTypeResponseV1{
		Name:      secatest.Workspace1Name,
		Tenant:    secatest.Tenant1Name,
		Region:    secatest.Region1Name,
		Workspace: secatest.Workspace1Name,
		State:     secatest.State1Active,
	})
	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	workspace.HandlerWithOptions(wsSim, workspace.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderWorkspaceEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{WorkspaceV1API})
	require.NoError(t, err)

	tref := TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Workspace1Name,
	}

	ws, err := regionalClient.WorkspaceV1.GetWorkspace(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, ws)
	assert.Equal(t, secatest.Workspace1Name, ws.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, ws.Metadata.Tenant)

	assert.NotNil(t, *ws.Status.State)
	assert.EqualValues(t, secatest.State1Active, *ws.Status.State)

}

func TestCreateOrUpdateWorkspace(t *testing.T) {
	ctx := context.Background()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderWorkspaceName,
				URL:  secatest.ProviderWorkspaceEndpoint,
			},
		},
	})

	wsSim := mockworkspace.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateWorkspaceV1(wsSim, secatest.WorkspaceTypeResponseV1{
		Name:   secatest.Workspace1Name,
		Tenant: secatest.Tenant1Name,
		State:  secatest.State1Active,
	})
	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	workspace.HandlerWithOptions(wsSim, workspace.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderWorkspaceEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{WorkspaceV1API})
	require.NoError(t, err)

	ws := &workspace.Workspace{
		Metadata: &workspace.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Workspace1Name,
		},
		Spec: workspace.WorkspaceSpec{},
	}

	err = regionalClient.WorkspaceV1.CreateOrUpdateWorkspace(ctx, ws)

	require.NoError(t, err)

}

func TestDeleteWorkspace(t *testing.T) {
	ctx := context.Background()

	sim := mockregion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderWorkspaceName,
				URL:  secatest.ProviderWorkspaceEndpoint,
			},
		},
	})

	wsSim := mockworkspace.NewMockServerInterface(t)
	secatest.MockDeleteWorkspaceV1(wsSim)
	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	workspace.HandlerWithOptions(wsSim, workspace.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderWorkspaceEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{WorkspaceV1API})
	require.NoError(t, err)

	ws := &workspace.Workspace{
		Metadata: &workspace.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Workspace1Name,
		},
		Spec: workspace.WorkspaceSpec{},
	}

	err = regionalClient.WorkspaceV1.DeleteWorkspace(ctx, ws)

	require.NoError(t, err)

}
