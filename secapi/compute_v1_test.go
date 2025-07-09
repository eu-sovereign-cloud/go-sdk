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
		Metadata: secatest.MetadataResponseV1{Name: secatest.InstanceSku1Name},
		Tier:     secatest.InstanceSku1Tier,
		VCPU:     secatest.InstanceSku1VCPU,
		Ram:      secatest.InstanceSku1RAM,
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, server)

	iter, err := regionalClient.ComputeV1.ListSkus(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	assert.Equal(t, secatest.InstanceSku1Name, resp[0].Metadata.Name)

	labels := *resp[0].Labels
	require.Len(t, labels, 1)
	assert.Equal(t, secatest.InstanceSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.InstanceSku1VCPU, resp[0].Spec.VCPU)
	assert.Equal(t, secatest.InstanceSku1RAM, resp[0].Spec.Ram)
}

func TestGetInstanceSkU(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockGetInstanceSkuV1(sim, secatest.InstanceSkuResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.InstanceSku1Name},
		Tier:     secatest.InstanceSku1Tier,
		VCPU:     secatest.InstanceSku1VCPU,
		Ram:      secatest.InstanceSku1RAM,
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, server)

	resp, err := regionalClient.ComputeV1.GetSku(ctx, TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Workspace1Name,
	})
	require.NoError(t, err)

	assert.Equal(t, secatest.InstanceSku1Name, resp.Metadata.Name)

	labels := *resp.Labels
	require.Len(t, labels, 1)
	assert.Equal(t, secatest.InstanceSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.InstanceSku1VCPU, resp.Spec.VCPU)
	assert.Equal(t, secatest.InstanceSku1RAM, resp.Spec.Ram)
}

// Instance

func TestListInstances(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockListInstancesV1(sim, secatest.InstanceResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, server)

	iter, err := regionalClient.ComputeV1.ListInstances(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.Instance1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.InstanceSku1Ref, resp[0].Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetInstance(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)
	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockGetInstanceV1(sim, secatest.InstanceResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Instance1Name,
	}
	resp, err := regionalClient.ComputeV1.GetInstance(ctx, wref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.Instance1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.InstanceSku1Ref, resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateInstance(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateInstanceV1(sim, secatest.InstanceResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, server)

	inst := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Region:    secatest.Region1Name,
			Zone:      secatest.ZoneA,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		Spec: compute.InstanceSpec{
			SkuRef: secatest.InstanceSku1Ref,
			Zone:   secatest.ZoneA,
		},
	}
	resp, err := regionalClient.ComputeV1.CreateOrUpdateInstance(ctx, inst)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.Instance1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.InstanceSku1Ref, resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
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

	regionalClient := getTestRegionalClient(t, ctx, server)

	inst := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.StartInstance(ctx, inst)
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

	regionalClient := getTestRegionalClient(t, ctx, server)

	inst := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.RestartInstance(ctx, inst)
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

	regionalClient := getTestRegionalClient(t, ctx, server)

	inst := &compute.Instance{
		Metadata: &compute.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.ComputeV1.StopInstance(ctx, inst)
	require.NoError(t, err)
}

func TestDeleteInstance(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockGetInstanceV1(sim, secatest.InstanceResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})

	secatest.MockDeleteInstanceV1(sim)
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, server)
	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Instance1Name,
	}
	resp, err := regionalClient.ComputeV1.GetInstance(ctx, wref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.ComputeV1.DeleteInstance(ctx, resp)
	require.NoError(t, err)
}
