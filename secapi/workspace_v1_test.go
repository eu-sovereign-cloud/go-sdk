package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/assert"
)

// Workspace

func TestListWorkspacesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockListWorkspaceV1(sim, []schema.Workspace{
		*buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, secatest.StatusStateActive),
	})
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.WorkspaceV1.ListWorkspaces(ctx, secatest.Tenant1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetWorkspaces(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockGetWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, secatest.StatusStateActive))
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.WorkspaceV1.GetWorkspace(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Workspace1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateWorkspace(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, secatest.StatusStateCreating))
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	ws := &schema.Workspace{
		Metadata: &schema.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Workspace1Name,
		},
	}
	resp, err := regionalClient.WorkspaceV1.CreateOrUpdateWorkspace(ctx, ws)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)	
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteWorkspace(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockGetWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, secatest.StatusStateActive))
	secatest.MockDeleteWorkspaceV1(sim)
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.WorkspaceV1.GetWorkspace(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Workspace1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.WorkspaceV1.DeleteWorkspace(ctx, resp)
	assert.NoError(t, err)
}

// Builders

func buildResponseWorkspace(name string, tenant string, region string, state string) *schema.Workspace {
	return &schema.Workspace{
		Metadata: secatest.NewRegionalResourceMetadata(name, tenant, region),
		Spec:     schema.WorkspaceSpec{},
		Status:   secatest.NewWorkspaceStatus(state),
	}
}
