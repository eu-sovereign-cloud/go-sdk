package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Workspace
func TestListWorkspacesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockListWorkspaceV1(sim, secatest.WorkspaceTypeResponseV1{
	})
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wsIter, err := regionalClient.WorkspaceV1.ListWorkspaces(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	ws, err := wsIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, ws, 1)

	assert.Equal(t, secatest.Workspace1Name, ws[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, ws[0].Metadata.Tenant)
	assert.EqualValues(t, secatest.StatusStateActive, *ws[0].Status.State)
}
func TestGetWorkspaces(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockGetWorkspaceV1(sim, secatest.WorkspaceTypeResponseV1{
	})
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

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
	assert.EqualValues(t, secatest.StatusStateActive, *ws.Status.State)
}
func TestCreateOrUpdateWorkspace(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateWorkspaceV1(sim, secatest.WorkspaceTypeResponseV1{
	})
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	ws := &workspace.Workspace{
		Metadata: &workspace.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Workspace1Name,
		},
		Spec: workspace.WorkspaceSpec{},
	}
	err := regionalClient.WorkspaceV1.CreateOrUpdateWorkspace(ctx, ws)
	require.NoError(t, err)
}

func TestDeleteWorkspace(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockDeleteWorkspaceV1(sim)
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	ws := &workspace.Workspace{
		Metadata: &workspace.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Workspace1Name,
		},
		Spec: workspace.WorkspaceSpec{},
	}
	err := regionalClient.WorkspaceV1.DeleteWorkspace(ctx, ws)
	require.NoError(t, err)
}
