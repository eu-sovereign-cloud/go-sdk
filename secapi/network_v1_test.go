package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

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

	iter, err := regionalClient.NetworkV1.ListSkus(ctx, secatest.Tenant1Name)
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
	spec := buildResponseNetworkSpec(t, secatest.RouteTable1Ref)
	secatest.MockListNetworksV1(sim, []schema.Network{
		*buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	routeTableRef, err := BuildReferenceFromURN(secatest.RouteTable1Ref)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := regionalClient.NetworkV1.ListNetworks(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Network1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, *routeTableRef, resp[0].Spec.RouteTableRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(t, secatest.RouteTable1Ref)
	secatest.MockGetNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	routeTableRef, err := BuildReferenceFromURN(secatest.RouteTable1Ref)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := regionalClient.NetworkV1.GetNetwork(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Network1Name})
	assert.NoError(t, err)

	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *routeTableRef, resp.Spec.RouteTableRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateOrUpdateNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(t, secatest.RouteTable1Ref)
	secatest.MockCreateOrUpdateNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	routeTableRef, err := BuildReferenceFromURN(secatest.RouteTable1Ref)
	if err != nil {
		t.Fatal(err)
	}

	networkSkuRef, err := BuildReferenceFromURN(secatest.NetworkSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

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

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNetworkSpec(t, secatest.RouteTable1Ref)
	secatest.MockGetNetworkV1(sim, buildResponseNetwork(secatest.Network1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
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
	spec := buildResponseSubnetSpec(t, secatest.NetworkSku1Ref)
	secatest.MockListSubnetsV1(sim, []schema.Subnet{
		*buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	networkSkuRef, err := BuildReferenceFromURN(secatest.NetworkSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := regionalClient.NetworkV1.ListSubnets(ctx, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name)
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

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(t, secatest.NetworkSku1Ref)
	secatest.MockGetSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	networkSkuRef, err := BuildReferenceFromURN(secatest.NetworkSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := regionalClient.NetworkV1.GetSubnet(ctx, NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.Subnet1Name})
	assert.NoError(t, err)

	assert.Equal(t, secatest.Subnet1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Network)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *networkSkuRef, *resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(t, secatest.NetworkSku1Ref)
	secatest.MockCreateOrUpdateSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	networkSkuRef, err := BuildReferenceFromURN(secatest.NetworkSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

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

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSubnetSpec(t, secatest.NetworkSku1Ref)
	secatest.MockGetSubnetV1(sim, buildResponseSubnet(secatest.Subnet1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
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
	spec := buildResponseRouteTableSpec(t, secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockListRouteTablesV1(sim, []schema.RouteTable{
		*buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	targetRef, err := BuildReferenceFromURN(secatest.Instance1Ref)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := regionalClient.NetworkV1.ListRouteTables(ctx, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name)
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

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(t, secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockGetRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	targetRef, err := BuildReferenceFromURN(secatest.Instance1Ref)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := regionalClient.NetworkV1.GetRouteTable(ctx, NetworkReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Network: secatest.Network1Name, Name: secatest.RouteTable1Name})
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

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(t, secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockCreateOrUpdateRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	targetRef, err := BuildReferenceFromURN(secatest.Instance1Ref)
	if err != nil {
		t.Fatal(err)
	}

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

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseRouteTableSpec(t, secatest.CidrIpv4, secatest.Instance1Ref)
	secatest.MockGetRouteTableV1(sim, buildResponseRouteTable(secatest.RouteTable1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Network1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
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
		*buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListInternetGateways(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.InternetGateway1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, false, *resp[0].Spec.EgressOnly)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockGetInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetInternetGateway(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name})
	assert.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, false, *resp.Spec.EgressOnly)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockCreateOrUpdateInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateCreating))
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

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseInternetGatewaySpec(false)
	secatest.MockGetInternetGatewayV1(sim, buildResponseInternetGateway(secatest.InternetGateway1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
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

// Security Group

func TestListSecurityGroupsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockListSecurityGroupsV1(sim, []schema.SecurityGroup{
		*buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListSecurityGroups(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.SecurityGroup1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp[0].Spec.Rules[0].Direction))

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetSecurityGroup(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name})
	assert.NoError(t, err)

	assert.Equal(t, secatest.SecurityGroup1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp.Spec.Rules[0].Direction))

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockCreateOrUpdateSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateCreating))
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
			Rules: []schema.SecurityGroupRuleSpec{
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

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp.Spec.Rules[0].Direction))

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseSecurityGroupSpec(secatest.SecurityGroupRuleDirectionIngress)
	secatest.MockGetSecurityGroupV1(sim, buildResponseSecurityGroup(secatest.SecurityGroup1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
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
	spec := buildResponseNicSpec(t, secatest.Subnet1Ref)
	secatest.MockListNicsV1(sim, []schema.Nic{
		*buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	subnetRef, err := BuildReferenceFromURN(secatest.Subnet1Ref)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := regionalClient.NetworkV1.ListNics(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Nic1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, *subnetRef, resp[0].Spec.SubnetRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(t, secatest.Subnet1Ref)
	secatest.MockGetNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	subnetRef, err := BuildReferenceFromURN(secatest.Subnet1Ref)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := regionalClient.NetworkV1.GetNic(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name})
	assert.NoError(t, err)

	assert.Equal(t, secatest.Nic1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *subnetRef, resp.Spec.SubnetRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(t, secatest.Subnet1Ref)
	secatest.MockCreateOrUpdateNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateCreating))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	subnetRef, err := BuildReferenceFromURN(secatest.Subnet1Ref)
	if err != nil {
		t.Fatal(err)
	}

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

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponseNicSpec(t, secatest.Subnet1Ref)
	secatest.MockGetNicV1(sim, buildResponseNic(secatest.Nic1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
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
		*buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.NetworkV1.ListPublicIps(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.PublicIp1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, secatest.Address1, *resp[0].Spec.Address)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetPublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockGetPublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.NetworkV1.GetPublicIp(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name})
	assert.NoError(t, err)

	assert.Equal(t, secatest.PublicIp1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, secatest.Address1, *resp.Spec.Address)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdatePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockCreateOrUpdatePublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateCreating))
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

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeletePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	spec := buildResponsePublicIpSpec(secatest.Address1)
	secatest.MockGetPublicIpV1(sim, buildResponsePublicIp(secatest.PublicIp1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, secatest.StatusStateActive))
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

func buildResponseNetwork(name string, tenant string, workspace string, region string, spec *schema.NetworkSpec, state string) *schema.Network {
	return &schema.Network{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewNetworkStatus(state),
	}
}

func buildResponseNetworkSpec(t *testing.T, routeTableRef string) *schema.NetworkSpec {
	urnRef, err := BuildReferenceFromURN(routeTableRef)
	if err != nil {
		t.Fatal(err)
	}

	return &schema.NetworkSpec{
		RouteTableRef: *urnRef,
	}
}

func buildResponseSubnet(name string, tenant string, workspace string, network string, region string, spec *schema.SubnetSpec, state string) *schema.Subnet {
	return &schema.Subnet{
		Metadata: secatest.NewRegionalNetworkResourceMetadata(name, tenant, workspace, network, region),
		Spec:     *spec,
		Status:   secatest.NewSubnetStatus(state),
	}
}

func buildResponseSubnetSpec(t *testing.T, skuRef string) *schema.SubnetSpec {
	urnRef, err := BuildReferenceFromURN(skuRef)
	if err != nil {
		t.Fatal(err)
	}

	return &schema.SubnetSpec{
		SkuRef: urnRef,
	}
}

func buildResponseRouteTable(name string, tenant string, workspace string, network string, region string, spec *schema.RouteTableSpec, state string) *schema.RouteTable {
	return &schema.RouteTable{
		Metadata: secatest.NewRegionalNetworkResourceMetadata(name, tenant, workspace, network, region),
		Spec:     *spec,
		Status:   secatest.NewRouteTableStatus(state),
	}
}

func buildResponseRouteTableSpec(t *testing.T, routeCidrBlock string, routeTargetRef string) *schema.RouteTableSpec {
	urnRef, err := BuildReferenceFromURN(routeTargetRef)
	if err != nil {
		t.Fatal(err)
	}

	return &schema.RouteTableSpec{
		Routes: []schema.RouteSpec{
			{
				DestinationCidrBlock: routeCidrBlock,
				TargetRef:            *urnRef,
			},
		},
	}
}

func buildResponseInternetGateway(name string, tenant string, workspace string, region string, spec *schema.InternetGatewaySpec, state string) *schema.InternetGateway {
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

func buildResponseSecurityGroup(name string, tenant string, workspace string, region string, spec *schema.SecurityGroupSpec, state string) *schema.SecurityGroup {
	return &schema.SecurityGroup{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewSecurityGroupStatus(state),
	}
}

func buildResponseSecurityGroupSpec(ruleDirection schema.SecurityGroupRuleSpecDirection) *schema.SecurityGroupSpec {
	return &schema.SecurityGroupSpec{
		Rules: []schema.SecurityGroupRuleSpec{
			{
				Direction: ruleDirection,
			},
		},
	}
}

func buildResponseNic(name string, tenant string, workspace string, region string, spec *schema.NicSpec, state string) *schema.Nic {
	return &schema.Nic{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewNicStatus(state),
	}
}

func buildResponseNicSpec(t *testing.T, subnetRef string) *schema.NicSpec {
	urnRef, err := BuildReferenceFromURN(subnetRef)
	if err != nil {
		t.Fatal(err)
	}

	return &schema.NicSpec{
		SubnetRef: *urnRef,
	}
}

func buildResponsePublicIp(name string, tenant string, workspace string, region string, spec *schema.PublicIpSpec, state string) *schema.PublicIp {
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
