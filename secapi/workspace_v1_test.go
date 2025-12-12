package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi/builders"

	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

// Workspace

func TestListWorkspacesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockListWorkspaceV1(sim, []schema.Workspace{
		*buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, schema.ResourceStateActive),
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

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListWorkspacesWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockListWorkspaceV1(sim, []schema.Workspace{
		{
			Metadata: &schema.RegionalResourceMetadata{
				Name:   secatest.Workspace1Name,
				Tenant: secatest.Tenant1Name,
			},
			Status: &schema.WorkspaceStatus{State: ptr.To(schema.ResourceStateActive)},
		},
	})
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := NewListOptions().WithLimit(10).WithLabels(labelsParams)
	iter, err := regionalClient.WorkspaceV1.ListWorkspacesWithFilters(ctx, secatest.Tenant1Name, listOptions)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetWorkspaceV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockGetWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, schema.ResourceStateActive), 1)
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Workspace1Name}
	resp, err := regionalClient.WorkspaceV1.GetWorkspace(ctx, tref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetWorkspaceUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockGetWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, schema.ResourceStateCreating), 2)
	secatest.MockGetWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, schema.ResourceStateActive), 1)
	secatest.ConfigureWorkspaceHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Workspace1Name}
	config := ResourceObserverConfig[schema.ResourceState]{ExpectedValue: schema.ResourceStateActive, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.WorkspaceV1.GetWorkspaceUntilState(ctx, tref, config)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestCreateOrUpdateWorkspaceV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, schema.ResourceStateCreating))
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

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteWorkspaceV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockworkspace.NewMockServerInterface(t)
	secatest.MockGetWorkspaceV1(sim, buildResponseWorkspace(secatest.Workspace1Name, secatest.Tenant1Name, secatest.Region1Name, schema.ResourceStateActive), 1)
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

func buildResponseWorkspace(name string, tenant string, region string, state schema.ResourceState) *schema.Workspace {
	return &schema.Workspace{
		Metadata: secatest.NewRegionalResourceMetadata(name, tenant, region),
		Spec:     schema.WorkspaceSpec{},
		Status:   secatest.NewWorkspaceStatus(state),
	}
}
