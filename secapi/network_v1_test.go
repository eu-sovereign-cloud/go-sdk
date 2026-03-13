package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi/builders"

	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

// Network Sku

func TestListNetworkSkusV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	labels := schema.Labels{secatest.LabelKeyTier: secatest.NetworkSku1Tier}
	spec := buildResponseNetworkSkuSpec(secatest.NetworkSku1Bandwidth, secatest.NetworkSku1Packets)
	secatest.MockListNetworkSkusV1(sim, []schema.NetworkSku{
		*buildResponseNetworkSku(secatest.NetworkSku1Name, secatest.Tenant1Name, labels, spec),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListSkus(ctx, TenantFilter{Tenant: secatest.Tenant1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.NetworkSku1Name, resp[0].Metadata.Name)

	assert.Len(t, resp[0].Labels, 1)
	assert.Equal(t, secatest.NetworkSku1Tier, resp[0].Labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.NetworkSku1Bandwidth, resp[0].Spec.Bandwidth)
	assert.Equal(t, secatest.NetworkSku1Packets, resp[0].Spec.Packets)
}

func TestGetNetworkSkuV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	labels := schema.Labels{secatest.LabelKeyTier: secatest.NetworkSku1Tier}
	spec := buildResponseNetworkSkuSpec(secatest.NetworkSku1Bandwidth, secatest.NetworkSku1Packets)
	secatest.MockGetNetworkSkuV1(sim, buildResponseNetworkSku(secatest.NetworkSku1Name, secatest.Tenant1Name, labels, spec))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetSku(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.NetworkSku1Name})
	assert.NoError(t, err)

	assert.Equal(t, secatest.NetworkSku1Name, resp.Metadata.Name)

	assert.Len(t, resp.Labels, 1)
	assert.Equal(t, secatest.NetworkSku1Tier, resp.Labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.NetworkSku1Bandwidth, resp.Spec.Bandwidth)
	assert.Equal(t, secatest.NetworkSku1Packets, resp.Spec.Packets)
}

// Network

func TestListNetworksV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(secatest.RouteTable1Ref)
	secatest.MockListNetworksV1(sim, []schema.Network{
		*buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	routeTableRef := &schema.Reference{Resource: secatest.RouteTable1Ref}

	iter, err := regionalClient.NetworkV1.ListNetworks(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Network1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, *routeTableRef, resp[0].Spec.RouteTableRef)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListNetworksWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	routeTableRef := &schema.Reference{Resource: secatest.RouteTable1Ref}

	secatest.MockListNetworksV1(sim, []schema.Network{
		{
			Metadata: &schema.RegionalWorkspaceResourceMetadata{
				Name:      secatest.Network1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
			},
			Spec: schema.NetworkSpec{
				RouteTableRef: *routeTableRef,
			},
			Status: &schema.NetworkStatus{
				State: ptr.To(schema.ResourceStateCreating),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)
	iter, err := regionalClient.NetworkV1.ListNetworks(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Network1Name, resp[0].Metadata.Name)

	assert.Equal(t, *routeTableRef, resp[0].Spec.RouteTableRef)

	assert.Equal(t, schema.ResourceStateCreating, *resp[0].Status.State)
}

func TestGetNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(secatest.RouteTable1Ref)
	secatest.MockGetNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	routeTableRef := &schema.Reference{Resource: secatest.RouteTable1Ref}

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Network1Name}
	resp, err := regionalClient.NetworkV1.GetNetwork(ctx, wref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *routeTableRef, resp.Spec.RouteTableRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetNetworkUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(secatest.RouteTable1Ref)
	secatest.MockGetNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	routeTableRef := &schema.Reference{Resource: secatest.RouteTable1Ref}

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Network1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetNetworkUntilState(ctx, wref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *routeTableRef, resp.Spec.RouteTableRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchNetworkUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(secatest.RouteTable1Ref)
	secatest.MockGetNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundNetworkV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Network1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchNetworkUntilDeleted(ctx, wref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdateOrUpdateNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(secatest.RouteTable1Ref)
	secatest.MockCreateOrUpdateNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	routeTableRef := &schema.Reference{Resource: secatest.RouteTable1Ref}

	networkSkuRef := &schema.Reference{Resource: secatest.NetworkSku1Ref}

	net := &schema.Network{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.Network1Name,
		},
		Spec: schema.NetworkSpec{
			Cidr:          schema.Cidr{Ipv4: ptr.To(secatest.CidrIpv4)},
			RouteTableRef: *routeTableRef,
			SkuRef:        *networkSkuRef,
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateNetwork(ctx, net)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *routeTableRef, resp.Spec.RouteTableRef)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(secatest.RouteTable1Ref)
	secatest.MockGetNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeleteNetworkV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	resp, err := regionalClient.NetworkV1.GetNetwork(ctx, wref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteNetwork(ctx, resp)
	assert.NoError(t, err)
}

// Subnet
func TestListSubnetsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(secatest.NetworkSku1Ref)
	secatest.MockListSubnetsV1(sim, []schema.Subnet{
		*buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	networkSkuRef := &schema.Reference{Resource: secatest.NetworkSku1Ref}

	iter, err := regionalClient.NetworkV1.ListSubnets(ctx, NetworkFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Subnet1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp[0].Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, *networkSkuRef, *resp[0].Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListSubnetsWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)

	skuRef := &schema.Reference{Resource: secatest.NetworkSku1Ref}
	secatest.MockListSubnetsV1(sim, []schema.Subnet{
		{
			Metadata: &schema.RegionalNetworkResourceMetadata{
				Name:      secatest.Subnet1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
				Network:   secatest.Network1Name,
			},
			Spec: schema.SubnetSpec{
				SkuRef: skuRef,
			},
			Status: &schema.SubnetStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := regionalClient.NetworkV1.ListSubnets(ctx, NetworkFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(secatest.NetworkSku1Ref)
	secatest.MockGetSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	networkSkuRef := &schema.Reference{Resource: secatest.NetworkSku1Ref}

	nref := NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.Subnet1Name}
	resp, err := regionalClient.NetworkV1.GetSubnet(ctx, nref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Subnet1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *networkSkuRef, *resp.Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetSubnetUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(secatest.NetworkSku1Ref)
	secatest.MockGetSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	networkSkuRef := &schema.Reference{Resource: secatest.NetworkSku1Ref}

	nref := NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.Subnet1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetSubnetUntilState(ctx, nref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Subnet1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *networkSkuRef, *resp.Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchSubnetUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(secatest.NetworkSku1Ref)
	secatest.MockGetSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundSubnetV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	nref := NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.Subnet1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchSubnetUntilDeleted(ctx, nref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdateSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(secatest.NetworkSku1Ref)
	secatest.MockCreateOrUpdateSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	networkSkuRef := &schema.Reference{Resource: secatest.NetworkSku1Ref}

	sub := &schema.Subnet{
		Metadata: &schema.RegionalNetworkResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Network:   secatest.Network1Name,
			Name:      secatest.Subnet1Name,
		},
		Spec: schema.SubnetSpec{
			Cidr: schema.Cidr{Ipv4: ptr.To(secatest.CidrIpv4)},
			Zone: schema.Zone(secatest.ZoneA),
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateSubnet(ctx, sub)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Subnet1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *networkSkuRef, *resp.Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(secatest.NetworkSku1Ref)
	secatest.MockGetSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeleteSubnetV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetSubnet(ctx, NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.Subnet1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteSubnet(ctx, resp)
	assert.NoError(t, err)
}

// Route Table

func TestListRouteTablesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockListRouteTablesV1(sim, []schema.RouteTable{
		*buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	targetRef := &schema.Reference{Resource: secatest.Instance1Ref}

	iter, err := regionalClient.NetworkV1.ListRouteTables(ctx, NetworkFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.GreaterOrEqual(t, 1, len(resp))

	assert.Equal(t, secatest.RouteTable1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp[0].Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.GreaterOrEqual(t, 1, len(resp[0].Spec.Routes))
	assert.Equal(t, secatest.CidrIpv4, resp[0].Spec.Routes[0].DestinationCidrBlock)
	assert.Equal(t, *targetRef, resp[0].Spec.Routes[0].TargetRef)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListRouteTablesWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	instanceRef := &schema.Reference{Resource: secatest.Instance1Ref}

	secatest.MockListRouteTablesV1(sim, []schema.RouteTable{
		{
			Metadata: &schema.RegionalNetworkResourceMetadata{
				Name:      secatest.RouteTable1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
				Network:   secatest.Network1Name,
			},
			Spec: schema.RouteTableSpec{
				Routes: []schema.RouteSpec{
					{
						DestinationCidrBlock: secatest.CidrIpv4,
						TargetRef:            *instanceRef,
					},
				},
			},
			Status: &schema.RouteTableStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := regionalClient.NetworkV1.ListRouteTables(ctx, NetworkFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockGetRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	targetRef := &schema.Reference{Resource: secatest.Instance1Ref}

	nref := NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.RouteTable1Name}
	resp, err := regionalClient.NetworkV1.GetRouteTable(ctx, nref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.RouteTable1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Len(t, resp.Spec.Routes, 1)

	route := resp.Spec.Routes[0]
	assert.Equal(t, secatest.CidrIpv4, route.DestinationCidrBlock)
	assert.Equal(t, *targetRef, route.TargetRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetRouteTableUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockGetRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	targetRef := &schema.Reference{Resource: secatest.Instance1Ref}

	nref := NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.RouteTable1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetRouteTableUntilState(ctx, nref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.RouteTable1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Len(t, resp.Spec.Routes, 1)

	route := resp.Spec.Routes[0]
	assert.Equal(t, secatest.CidrIpv4, route.DestinationCidrBlock)
	assert.Equal(t, *targetRef, route.TargetRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchRouteTableUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockGetRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundRouteTableV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	nref := NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.RouteTable1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchRouteTableUntilDeleted(ctx, nref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdateRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockCreateOrUpdateRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	targetRef := &schema.Reference{Resource: secatest.Instance1Ref}

	route := &schema.RouteTable{
		Metadata: &schema.RegionalNetworkResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Network:   secatest.Network1Name,
			Name:      secatest.RouteTable1Name,
		},
		Spec: schema.RouteTableSpec{
			Routes: []schema.RouteSpec{
				{
					DestinationCidrBlock: secatest.CidrIpv4,
					TargetRef:            *targetRef,
				},
			},
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateRouteTable(ctx, route)
	assert.NoError(t, err)

	assert.Equal(t, secatest.RouteTable1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.GreaterOrEqual(t, 1, len(resp.Spec.Routes))

	respRoute := resp.Spec.Routes[0]
	assert.Equal(t, secatest.CidrIpv4, respRoute.DestinationCidrBlock)
	assert.Equal(t, *targetRef, respRoute.TargetRef)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockGetRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeleteRouteTableV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetRouteTable(ctx, NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.RouteTable1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteRouteTable(ctx, resp)
	assert.NoError(t, err)
}

// Internet Gateway

func TestListInternetGatewaysV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockListInternetGatewaysV1(sim, []schema.InternetGateway{
		*buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListInternetGateways(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.InternetGateway1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, false, *resp[0].Spec.EgressOnly)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListInternetGatewaysWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListInternetGatewaysV1(sim, []schema.InternetGateway{
		{
			Metadata: &schema.RegionalWorkspaceResourceMetadata{
				Name:      secatest.InternetGateway1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
			},
			Spec: schema.InternetGatewaySpec{},
			Status: &schema.Status{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := regionalClient.NetworkV1.ListInternetGateways(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockGetInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name}
	resp, err := regionalClient.NetworkV1.GetInternetGateway(ctx, wref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, false, *resp.Spec.EgressOnly)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetInternetGatewayUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockGetInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetInternetGatewayUntilState(ctx, wref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, false, *resp.Spec.EgressOnly)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchInternetGatewayUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockGetInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundInternetGatewayV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchInternetGatewayUntilDeleted(ctx, wref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdateInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockCreateOrUpdateInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	gtw := &schema.InternetGateway{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.InternetGateway1Name,
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateInternetGateway(ctx, gtw)
	assert.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, false, *resp.Spec.EgressOnly)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockGetInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeleteInternetGatewayV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetInternetGateway(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteInternetGateway(ctx, resp)
	assert.NoError(t, err)
}

// Security Group Rules

func TestListSecurityGroupRulesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupRuleSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockListSecurityGroupRulesV1(sim, []schema.SecurityGroupRule{
		*buildResponseSecurityGroupRule(secatest.SecurityGroupRule1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListSecurityGroupRules(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.SecurityGroupRule1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp[0].Spec.Direction))

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListSecurityGroupRulesWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListSecurityGroupRulesV1(sim, []schema.SecurityGroupRule{
		{
			Metadata: &schema.RegionalWorkspaceResourceMetadata{
				Name:      secatest.SecurityGroupRule1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
			},
			Spec: schema.SecurityGroupRuleSpec{
				Direction: schema.SecurityGroupRuleDirectionIngress,
			},
			Status: &schema.SecurityGroupRuleStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)
	iter, err := regionalClient.NetworkV1.ListSecurityGroupRules(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetSecurityGroupRuleV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupRuleSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupRuleV1(sim, buildResponseSecurityGroupRule(secatest.SecurityGroupRule1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroupRule1Name}
	resp, err := regionalClient.NetworkV1.GetSecurityGroupRule(ctx, wref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.SecurityGroupRule1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp.Spec.Direction))

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetSecurityGroupRuleUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupRuleSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupRuleV1(sim, buildResponseSecurityGroupRule(secatest.SecurityGroupRule1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetSecurityGroupRuleV1(sim, buildResponseSecurityGroupRule(secatest.SecurityGroupRule1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroupRule1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetSecurityGroupRuleUntilState(ctx, wref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.SecurityGroupRule1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp.Spec.Direction))

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchSecurityGroupRuleUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupRuleSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupRuleV1(sim, buildResponseSecurityGroupRule(secatest.SecurityGroupRule1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundSecurityGroupRuleV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroupRule1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchSecurityGroupRuleUntilDeleted(ctx, wref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdateSecurityGroupRuleV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupRuleSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockCreateOrUpdateSecurityGroupRuleV1(sim, buildResponseSecurityGroupRule(secatest.SecurityGroupRule1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	group := &schema.SecurityGroupRule{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.SecurityGroupRule1Name,
		},
		Spec: schema.SecurityGroupRuleSpec{
			Direction: schema.SecurityGroupRuleDirectionIngress,
			Version:   ptr.To(schema.IPVersionIPv4),
			Protocol:  ptr.To(schema.SecurityGroupRuleProtocolTCP),
			Ports: &schema.Ports{
				From: ptr.To(schema.Port(secatest.SecurityGroup1PortFrom)),
				To:   ptr.To(schema.Port(secatest.SecurityGroup1PortTo)),
			},
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateSecurityGroupRule(ctx, group)
	assert.NoError(t, err)

	assert.Equal(t, secatest.SecurityGroupRule1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp.Spec.Direction))

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteSecurityGroupRuleV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupRuleSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupRuleV1(sim, buildResponseSecurityGroupRule(secatest.SecurityGroupRule1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeleteSecurityGroupRuleV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetSecurityGroupRule(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroupRule1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteSecurityGroupRule(ctx, resp)
	assert.NoError(t, err)
}

// Security Group

func TestListSecurityGroupsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockListSecurityGroupsV1(sim, []schema.SecurityGroup{
		*buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListSecurityGroups(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.SecurityGroup1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string((*resp[0].Spec.Rules)[0].Direction))

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListSecurityGroupsWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListSecurityGroupsV1(sim, []schema.SecurityGroup{
		{
			Metadata: &schema.RegionalWorkspaceResourceMetadata{
				Name:      secatest.SecurityGroup1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
			},
			Spec: schema.SecurityGroupSpec{
				Rules: &[]schema.SecurityGroupRuleSpec{
					{
						Direction: schema.SecurityGroupRuleDirectionIngress,
					},
				},
			},
			Status: &schema.SecurityGroupStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)
	iter, err := regionalClient.NetworkV1.ListSecurityGroups(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name}
	resp, err := regionalClient.NetworkV1.GetSecurityGroup(ctx, wref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.SecurityGroup1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string((*resp.Spec.Rules)[0].Direction))

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetSecurityGroupUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetSecurityGroupUntilState(ctx, wref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.SecurityGroup1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string((*resp.Spec.Rules)[0].Direction))

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchSecurityGroupUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundSecurityGroupV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchSecurityGroupUntilDeleted(ctx, wref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdateSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockCreateOrUpdateSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	group := &schema.SecurityGroup{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.SecurityGroup1Name,
		},
		Spec: schema.SecurityGroupSpec{
			Rules: &[]schema.SecurityGroupRuleSpec{
				{
					Direction: schema.SecurityGroupRuleDirectionIngress,
					Version:   ptr.To(schema.IPVersionIPv4),
					Protocol:  ptr.To(schema.SecurityGroupRuleProtocolTCP),
					Ports: &schema.Ports{
						From: ptr.To(schema.Port(secatest.SecurityGroup1PortFrom)),
						To:   ptr.To(schema.Port(secatest.SecurityGroup1PortTo)),
					},
				},
			},
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateSecurityGroup(ctx, group)
	assert.NoError(t, err)

	assert.Equal(t, secatest.SecurityGroup1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string((*resp.Spec.Rules)[0].Direction))

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeleteSecurityGroupV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetSecurityGroup(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteSecurityGroup(ctx, resp)
	assert.NoError(t, err)
}

// Nic

func TestListNicsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(secatest.Subnet1Ref)
	secatest.MockListNicsV1(sim, []schema.Nic{
		*buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	subnetRef := &schema.Reference{Resource: secatest.Subnet1Ref}

	iter, err := regionalClient.NetworkV1.ListNics(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Nic1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, *subnetRef, resp[0].Spec.SubnetRef)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListNicsWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	subnetRef := &schema.Reference{Resource: secatest.Instance1Ref}

	secatest.MockListNicsV1(sim, []schema.Nic{
		{
			Metadata: &schema.RegionalWorkspaceResourceMetadata{
				Name:      secatest.Nic1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
			},
			Spec: schema.NicSpec{
				SubnetRef: *subnetRef,
			},
			Status: &schema.NicStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := regionalClient.NetworkV1.ListNics(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(secatest.Subnet1Ref)
	secatest.MockGetNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	subnetRef := &schema.Reference{Resource: secatest.Subnet1Ref}

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name}
	resp, err := regionalClient.NetworkV1.GetNic(ctx, wref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Nic1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *subnetRef, resp.Spec.SubnetRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetNicUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(secatest.Subnet1Ref)
	secatest.MockGetNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	subnetRef := &schema.Reference{Resource: secatest.Subnet1Ref}

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetNicUntilState(ctx, wref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Nic1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *subnetRef, resp.Spec.SubnetRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchNicUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(secatest.Subnet1Ref)
	secatest.MockGetNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundNicV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchNicUntilDeleted(ctx, wref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdateNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(secatest.Subnet1Ref)
	secatest.MockCreateOrUpdateNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	subnetRef := &schema.Reference{Resource: secatest.Subnet1Ref}
	nic := &schema.Nic{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.Nic1Name,
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateNic(ctx, nic)
	assert.NoError(t, err)

	assert.Equal(t, secatest.Nic1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *subnetRef, resp.Spec.SubnetRef)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(secatest.Subnet1Ref)
	secatest.MockGetNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeleteNicV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetNic(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteNic(ctx, resp)
	assert.NoError(t, err)
}

// Public Ip

func TestListPublicIpsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockListPublicIpsV1(sim, []schema.PublicIp{
		*buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListPublicIps(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.PublicIp1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, secatest.Address1, *resp[0].Spec.Address)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListPublicIpsWithOptionsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListPublicIpsV1(sim, []schema.PublicIp{
		{
			Metadata: &schema.RegionalWorkspaceResourceMetadata{
				Name:      secatest.PublicIp1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
			},
			Spec: schema.PublicIpSpec{
				Address: ptr.To(secatest.Address1),
			},
			Status: &schema.PublicIpStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

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

	filterOptions := NewFilterOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := regionalClient.NetworkV1.ListPublicIps(ctx, WorkspaceFilter{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Options: filterOptions})
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetPublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockGetPublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name}
	resp, err := regionalClient.NetworkV1.GetPublicIp(ctx, wref)
	assert.NoError(t, err)

	assert.Equal(t, secatest.PublicIp1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.Address1, *resp.Spec.Address)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetPublicIpUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockGetPublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetPublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name}
	config := ResourceObserverUntilValueConfig[schema.ResourceState]{ExpectedValues: []schema.ResourceState{schema.ResourceStateActive}, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.NetworkV1.GetPublicIpUntilState(ctx, wref, config)
	assert.NoError(t, err)

	assert.Equal(t, secatest.PublicIp1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.Address1, *resp.Spec.Address)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestWatchPublicIpUntilDeletedV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockGetPublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateDeleting), 2)
	secatest.MockNotFoundPublicIpV1(sim, nil, 1)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name}
	config := ResourceObserverConfig{Delay: 0, Interval: 0, MaxAttempts: 5}
	err := regionalClient.NetworkV1.WatchPublicIpUntilDeleted(ctx, wref, config)
	assert.NoError(t, err)
}

func TestCreateOrUpdatePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockCreateOrUpdatePublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	ip := &schema.PublicIp{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.PublicIp1Name,
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdatePublicIp(ctx, ip)
	assert.NoError(t, err)

	assert.Equal(t, secatest.PublicIp1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.Address1, *resp.Spec.Address)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeletePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockGetPublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.MockDeletePublicIpV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetPublicIp(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeletePublicIp(ctx, resp)
	assert.NoError(t, err)
}

// Builders

func buildResponseNetworkSku(name string, tenant string, labels schema.Labels, spec *schema.NetworkSkuSpec) *schema.NetworkSku {
	return &schema.NetworkSku{
		Metadata: secatest.NewSkuResourceMetadata(name, tenant),
		Labels:   labels,
		Spec:     spec,
	}
}

func buildResponseNetworkSkuSpec(bandwidth int, packets int) *schema.NetworkSkuSpec {
	return &schema.NetworkSkuSpec{
		Bandwidth: bandwidth,
		Packets:   packets,
	}
}

func buildResponseNetwork(name string, tenant string, workspace string, region string, spec *schema.NetworkSpec, state schema.ResourceState) *schema.Network {
	return &schema.Network{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewNetworkStatus(state),
	}
}

func buildResponseNetworkSpec(routeTableRef string) *schema.NetworkSpec {
	ref := &schema.Reference{Resource: routeTableRef}

	return &schema.NetworkSpec{
		RouteTableRef: *ref,
	}
}

func buildResponseSubnet(name string, tenant string, workspace string, network string, region string, spec *schema.SubnetSpec, state schema.ResourceState) *schema.Subnet {
	return &schema.Subnet{
		Metadata: secatest.NewRegionalNetworkResourceMetadata(name, tenant, workspace, network, region),
		Spec:     *spec,
		Status:   secatest.NewSubnetStatus(state),
	}
}

func buildResponseSubnetSpec(skuRef string) *schema.SubnetSpec {
	ref := &schema.Reference{Resource: skuRef}

	return &schema.SubnetSpec{
		SkuRef: ref,
	}
}

func buildResponseRouteTable(name string, tenant string, workspace string, network string, region string, spec *schema.RouteTableSpec, state schema.ResourceState) *schema.RouteTable {
	return &schema.RouteTable{
		Metadata: secatest.NewRegionalNetworkResourceMetadata(name, tenant, workspace, network, region),
		Spec:     *spec,
		Status:   secatest.NewRouteTableStatus(state),
	}
}

func buildResponseRouteTableSpec(routeCidrBlock string, routeTargetRef string) *schema.RouteTableSpec {
	ref := &schema.Reference{Resource: routeTargetRef}

	return &schema.RouteTableSpec{
		Routes: []schema.RouteSpec{
			{
				DestinationCidrBlock: routeCidrBlock,
				TargetRef:            *ref,
			},
		},
	}
}

func buildResponseInternetGateway(name string, tenant string, workspace string, region string, spec *schema.InternetGatewaySpec, state schema.ResourceState) *schema.InternetGateway {
	return &schema.InternetGateway{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewInternetGatewayStatus(state),
	}
}

func buildResponseInternetGatewaySpec(egressOnly bool) *schema.InternetGatewaySpec {
	return &schema.InternetGatewaySpec{
		EgressOnly: &egressOnly,
	}
}

func buildResponseSecurityGroupRule(name string, tenant string, workspace string, region string, spec *schema.SecurityGroupRuleSpec, state schema.ResourceState) *schema.SecurityGroupRule {
	return &schema.SecurityGroupRule{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewSecurityGroupRuleStatus(state),
	}
}

func buildResponseSecurityGroupRuleSpec(ruleDirection schema.SecurityGroupRuleSpecDirection) *schema.SecurityGroupRuleSpec {
	return &schema.SecurityGroupRuleSpec{
		Direction: ruleDirection,
	}
}

func buildResponseSecurityGroup(name string, tenant string, workspace string, region string, spec *schema.SecurityGroupSpec, state schema.ResourceState) *schema.SecurityGroup {
	return &schema.SecurityGroup{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewSecurityGroupStatus(state),
	}
}

func buildResponseSecurityGroupSpec(ruleDirection schema.SecurityGroupRuleSpecDirection) *schema.SecurityGroupSpec {
	return &schema.SecurityGroupSpec{
		Rules: &[]schema.SecurityGroupRuleSpec{
			{
				Direction: ruleDirection,
			},
		},
	}
}

func buildResponseNic(name string, tenant string, workspace string, region string, spec *schema.NicSpec, state schema.ResourceState) *schema.Nic {
	return &schema.Nic{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewNicStatus(state),
	}
}

func buildResponseNicSpec(subnetRef string) *schema.NicSpec {
	ref := &schema.Reference{Resource: subnetRef}

	return &schema.NicSpec{
		SubnetRef: *ref,
	}
}

func buildResponsePublicIp(name string, tenant string, workspace string, region string, spec *schema.PublicIpSpec, state schema.ResourceState) *schema.PublicIp {
	return &schema.PublicIp{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewPublicIpStatus(state),
	}
}

func buildResponsePublicIpSpec(address string) *schema.PublicIpSpec {
	return &schema.PublicIpSpec{
		Address: &address,
	}
}
