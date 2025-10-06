package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockcompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi/builders"
	"k8s.io/utils/ptr"

	"github.com/stretchr/testify/assert"
)

// Instance Sku

func TestListInstancesSku(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockListInstanceSkusV1(sim, secatest.InstanceSkuResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.InstanceSku1Name,
		},
		Tier: secatest.InstanceSku1Tier,
		VCPU: secatest.InstanceSku1VCPU,
		Ram:  secatest.InstanceSku1RAM,
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.ComputeV1.ListSkus(ctx, secatest.Tenant1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)

	assert.Equal(t, secatest.InstanceSku1Name, resp[0].Metadata.Name)

	respLabels := resp[0].Labels
	assert.Len(t, respLabels, 1)
	assert.Equal(t, secatest.InstanceSku1Tier, respLabels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.InstanceSku1VCPU, resp[0].Spec.VCPU)
	assert.Equal(t, secatest.InstanceSku1RAM, resp[0].Spec.Ram)

	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75).
		Build()

	iter, err = regionalClient.ComputeV1.ListSkusWithFilters(ctx, secatest.Tenant1Name, ptr.To(1), ptr.To(labelsParams))
	assert.NoError(t, err)

	resp, err = iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetInstanceSkU(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	secatest.MockGetInstanceSkuV1(sim, secatest.InstanceSkuResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.InstanceSku1Name,
		},
		Tier: secatest.InstanceSku1Tier,
		VCPU: secatest.InstanceSku1VCPU,
		Ram:  secatest.InstanceSku1RAM,
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.ComputeV1.GetSku(ctx, TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.InstanceSku1Name,
	})
	assert.NoError(t, err)

	assert.Equal(t, secatest.InstanceSku1Name, resp.Metadata.Name)

	labels := resp.Labels
	assert.Len(t, labels, 1)
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
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	instanceSkuRef, err := regionalClient.ComputeV1.BuildReferenceURN(secatest.InstanceSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := regionalClient.ComputeV1.ListInstances(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Instance1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)

	assert.Equal(t, *instanceSkuRef, resp[0].Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))

	labelsParams := builders.NewLabelsBuilder().
		Equals("env", "test").
		Equals("*env*", "*prod*").
		NsEquals("monitoring", "alert-level", "high").
		Neq("tier", "frontend").
		Gt("version", 1).
		Lt("version", 3).
		Gte("uptime", 99).
		Lte("load", 75).
		Build()

	iter, err = regionalClient.ComputeV1.ListInstancesWithFilters(ctx, secatest.Tenant1Name, secatest.Workspace1Name, ptr.To(1), ptr.To(labelsParams))
	assert.NoError(t, err)

	resp, err = iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
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
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	instanceSkuRef, err := regionalClient.ComputeV1.BuildReferenceURN(secatest.InstanceSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Instance1Name,
	}
	resp, err := regionalClient.ComputeV1.GetInstance(ctx, wref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Instance1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)

	assert.Equal(t, *instanceSkuRef, resp.Spec.SkuRef)

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
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	instanceSkuRef, err := regionalClient.ComputeV1.BuildReferenceURN(secatest.InstanceSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	inst := &schema.Instance{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.Instance1Name,
		},
		Spec: schema.InstanceSpec{
			SkuRef: *instanceSkuRef,
			Zone:   secatest.ZoneA,
		},
	}
	resp, err := regionalClient.ComputeV1.CreateOrUpdateInstance(ctx, inst)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Instance1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)

	assert.Equal(t, *instanceSkuRef, resp.Spec.SkuRef)

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

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	inst := &schema.Instance{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: secatest.Workspace1Name,
		},
	}
	err := regionalClient.ComputeV1.StartInstance(ctx, inst)
	assert.NoError(t, err)
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

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	inst := &schema.Instance{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: secatest.Workspace1Name,
		},
	}
	err := regionalClient.ComputeV1.RestartInstance(ctx, inst)
	assert.NoError(t, err)
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

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	inst := &schema.Instance{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Instance1Name,
			Workspace: secatest.Workspace1Name,
		},
	}
	err := regionalClient.ComputeV1.StopInstance(ctx, inst)
	assert.NoError(t, err)
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
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.InstanceSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})

	secatest.MockDeleteInstanceV1(sim)
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)
	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Instance1Name,
	}
	resp, err := regionalClient.ComputeV1.GetInstance(ctx, wref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.ComputeV1.DeleteInstance(ctx, resp)
	assert.NoError(t, err)
}
