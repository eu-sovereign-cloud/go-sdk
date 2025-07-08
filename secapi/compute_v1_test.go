package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockcompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

// Instance Sku

func TestListInstancesSku(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockListInstanceSkusV1(sim, secatest.InstanceSkuResponseV1{
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.ComputeV1.ListSkus(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.NotEmpty(t, resp[0].Labels)
	require.NotEmpty(t, resp[0].Spec.VCPU)
	require.NotEmpty(t, resp[0].Spec.Ram)
}

func TestGetInstanceSkU(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockGetInstanceSkuV1(sim, secatest.InstanceSkuResponseV1{
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.ComputeV1.GetSku(ctx, TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Workspace1Name,
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp)

	assert.Equal(t, 4, resp.Spec.VCPU)
	assert.Equal(t, 32, resp.Spec.Ram)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Name)
}

// Instance

func TestListInstances(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockListInstancesV1(sim, secatest.InstanceResponseV1{
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.ComputeV1.ListInstances(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.Instance1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)
	assert.Equal(t, secatest.ZoneA, resp[0].Metadata.Zone)
}

func TestGetInstance(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)
	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockGetInstanceV1(sim, secatest.InstanceResponseV1{
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Workspace1Name,
	}
	resp, err := regionalClient.ComputeV1.GetInstance(ctx, wref)
	require.NoError(t, err)
	require.NotEmpty(t, resp)

	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)
	assert.Equal(t, secatest.ZoneA, resp.Metadata.Zone)
}

func TestCreateOrUpdateInstance(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateInstanceV1(sim, secatest.InstanceResponseV1{
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Region:    secatest.Region1Name,
			Zone:      secatest.ZoneA,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.CreateOrUpdateInstance(ctx, cp)
	require.NoError(t, err)
}

func TestStartInstanace(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockStartInstanceV1(sim)
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.StartInstance(ctx, cp)
	require.NoError(t, err)
}

func TestRestartInstanace(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockRestartInstanceV1(sim)
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.RestartInstance(ctx, cp)
	require.NoError(t, err)
}

func TestStopInstanace(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockStopInstanceV1(sim)
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.StopInstance(ctx, cp)
	require.NoError(t, err)
}

func TestDeleteInstance(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockDeleteInstanceV1(sim)
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	cp := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.DeleteInstance(ctx, cp)
	require.NoError(t, err)
}
