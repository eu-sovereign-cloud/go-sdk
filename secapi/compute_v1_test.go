package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockcompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/assert"
)

// Instance Sku

func TestListInstancesSku(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	labels := schema.Labels{secatest.LabelKeyTier: secatest.InstanceSku1Tier}
	spec := buildResponseInstanceSkuSpec(secatest.InstanceSku1VCPU, secatest.InstanceSku1RAM)
	secatest.MockListInstanceSkusV1(sim, []schema.InstanceSku{
		*buildResponseInstanceSku(secatest.InstanceSku1Name, secatest.Tenant1Name, labels, spec),
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

	assert.Len(t, resp[0].Labels, 1)
	assert.Equal(t, secatest.InstanceSku1Tier, resp[0].Labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.InstanceSku1VCPU, resp[0].Spec.VCPU)
	assert.Equal(t, secatest.InstanceSku1RAM, resp[0].Spec.Ram)
}

func TestGetInstanceSkU(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	labels := schema.Labels{secatest.LabelKeyTier: secatest.InstanceSku1Tier}
	spec := buildResponseInstanceSkuSpec(secatest.InstanceSku1VCPU, secatest.InstanceSku1RAM)
	secatest.MockGetInstanceSkuV1(sim, buildResponseInstanceSku(secatest.InstanceSku1Name, secatest.Tenant1Name, labels, spec))
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

	assert.Len(t, resp.Labels, 1)
	assert.Equal(t, secatest.InstanceSku1Tier, resp.Labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.InstanceSku1VCPU, resp.Spec.VCPU)
	assert.Equal(t, secatest.InstanceSku1RAM, resp.Spec.Ram)
}

// Instance

func TestListInstances(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	spec := buildResponseInstanceSpec(t, secatest.InstanceSku1Ref, secatest.ZoneA)
	secatest.MockListInstancesV1(sim, []schema.Instance{
		*buildResponseInstance(secatest.Instance1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	instanceSkuRef, err := BuildReferenceFromURN(secatest.InstanceSku1Ref)
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
}

func TestGetInstance(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockcompute.NewMockServerInterface(t)
	spec := buildResponseInstanceSpec(t, secatest.InstanceSku1Ref, secatest.ZoneA)
	secatest.MockGetInstanceV1(sim, buildResponseInstance(secatest.Instance1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	instanceSkuRef, err := BuildReferenceFromURN(secatest.InstanceSku1Ref)
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
	spec := buildResponseInstanceSpec(t, secatest.InstanceSku1Ref, secatest.ZoneA)
	secatest.MockCreateOrUpdateInstanceV1(sim, buildResponseInstance(secatest.Instance1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateCreating))
	secatest.ConfigureComputeHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	instanceSkuRef, err := BuildReferenceFromURN(secatest.InstanceSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	inst := &schema.Instance{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
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
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
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
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
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
			Name:      secatest.Instance1Name,
			Tenant:    secatest.Tenant1Name,
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
	spec := buildResponseInstanceSpec(t, secatest.InstanceSku1Ref, secatest.ZoneA)
	secatest.MockGetInstanceV1(sim, buildResponseInstance(secatest.Instance1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateActive))

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

// Builders

func buildResponseInstanceSku(name string, tenant string, labels schema.Labels, spec *schema.InstanceSkuSpec) *schema.InstanceSku {
	return &schema.InstanceSku{
		Metadata: secatest.NewSkuResourceMetadata(name, tenant),
		Labels:   labels,
		Spec:     spec,
	}
}

func buildResponseInstanceSkuSpec(vCPU int, ram int) *schema.InstanceSkuSpec {
	return &schema.InstanceSkuSpec{
		VCPU: vCPU,
		Ram:  ram,
	}
}

func buildResponseInstance(name string, tenant string, workspace string, spec *schema.InstanceSpec, state string) *schema.Instance {
	return &schema.Instance{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace),
		Spec:     *spec,
		Status:   secatest.NewInstanceStatus(state),
	}
}

func buildResponseInstanceSpec(t *testing.T, skuRef string, zone string) *schema.InstanceSpec {
	urnRef, err := BuildReferenceFromURN(skuRef)
	if err != nil {
		t.Fatal(err)
	}

	return &schema.InstanceSpec{
		SkuRef: *urnRef,
		Zone:   zone,
	}
}
