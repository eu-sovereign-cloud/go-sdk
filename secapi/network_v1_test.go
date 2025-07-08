package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

// Network Sku

func TestListNetworkSkusV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListNetworkSkusV1(sim, secatest.NetworkSkuResponseV1{
		Metadata:  secatest.MetadataResponseV1{Name: secatest.NetworkSku1Name},
		Tier:      secatest.NetworkSku1Tier,
		Bandwidth: secatest.NetworkSku1Bandwidth,
		Packets:   secatest.NetworkSku1Packets,
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListSkus(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.NetworkSku1Name, resp[0].Metadata.Name)

	labels := *resp[0].Labels
	require.Len(t, labels, 1)
	assert.Equal(t, secatest.NetworkSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.NetworkSku1Bandwidth, resp[0].Spec.Bandwidth)
	assert.Equal(t, secatest.NetworkSku1Packets, resp[0].Spec.Packets)
}

func TestGetNetworkSkuV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNetworkSkuV1(sim, secatest.NetworkSkuResponseV1{
		Metadata:  secatest.MetadataResponseV1{Name: secatest.NetworkSku1Name},
		Tier:      secatest.NetworkSku1Tier,
		Bandwidth: secatest.NetworkSku1Bandwidth,
		Packets:   secatest.NetworkSku1Packets,
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSku(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.NetworkSku1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.NetworkSku1Name, resp.Metadata.Name)

	labels := *resp.Labels
	require.Len(t, labels, 1)
	assert.Equal(t, secatest.NetworkSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.NetworkSku1Bandwidth, resp.Spec.Bandwidth)
	assert.Equal(t, secatest.NetworkSku1Packets, resp.Spec.Packets)
}

// Network

func TestListNetworksV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListNetworksV1(sim, secatest.NetworkResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Network1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RouteTableRef: secatest.RouteTable1Ref,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListNetworks(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.Network1Name, resp[0].Metadata.Name)

	assert.Equal(t, secatest.RouteTable1Ref, resp[0].Spec.RouteTableRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp[0].Status.State))
}

func TestGetNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNetworkV1(sim, secatest.NetworkResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Network1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RouteTableRef: secatest.RouteTable1Ref,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetNetwork(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Network1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)

	assert.Equal(t, secatest.RouteTable1Ref, resp.Spec.RouteTableRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestCreateOrUpdateOrUpdateNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateNetworkV1(sim, secatest.NetworkResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Network1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RouteTableRef: secatest.RouteTable1Ref,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	net := &network.Network{
		Metadata: &network.RegionalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
			Name:      secatest.Network1Name,
		},
		Spec: network.NetworkSpec{
			Cidr:          network.Cidr{Ipv4: ptr.To(secatest.CidrIpv4)},
			RouteTableRef: secatest.RouteTable1Ref,
			SkuRef:        secatest.NetworkSku1Ref,
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateNetwork(ctx, net)
	require.NoError(t, err)

	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)

	assert.Equal(t, secatest.RouteTable1Ref, resp.Spec.RouteTableRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNetworkV1(sim, secatest.NetworkResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Network1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RouteTableRef: secatest.RouteTable1Ref,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteNetworkV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	resp, err := regionalClient.NetworkV1.GetNetwork(ctx, wref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteNetwork(ctx, resp)
	require.NoError(t, err)
}

// Subnet
func TestListSubnetsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListSubnetsV1(sim, secatest.SubnetResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Subnet1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.NetworkSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListSubnets(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.Subnet1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.NetworkSku1Ref, *resp[0].Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSubnetV1(sim, secatest.SubnetResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.Subnet1Name},
		SkuRef:   secatest.NetworkSku1Ref,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSubnet(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Subnet1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Subnet1Name, resp.Metadata.Name)

	assert.Equal(t, secatest.NetworkSku1Ref, *resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateSubnetV1(sim, secatest.SubnetResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Subnet1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.NetworkSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	sub := &network.Subnet{
		Metadata: &network.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
			Name:      secatest.Subnet1Name,
		},
		Spec: network.SubnetSpec{
			Cidr: network.Cidr{Ipv4: ptr.To(secatest.CidrIpv4)},
			Zone: network.Zone(secatest.ZoneA),
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateSubnet(ctx, sub)
	require.NoError(t, err)

	assert.Equal(t, secatest.Subnet1Name, resp.Metadata.Name)

	assert.Equal(t, secatest.NetworkSku1Ref, *resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestDeleteSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSubnetV1(sim, secatest.SubnetResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Subnet1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.NetworkSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteSubnetV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSubnet(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Subnet1Name})
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteSubnet(ctx, resp)
	require.NoError(t, err)
}

// Route Table

func TestListRouteTablesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListRouteTablesV1(sim, secatest.RouteTableResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.RouteTable1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		LocalRef: secatest.Network1Ref,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListRouteTables(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.RouteTable1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.Network1Ref, resp[0].Spec.LocalRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetRouteTableV1(sim, secatest.RouteTableResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.RouteTable1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		LocalRef: secatest.Network1Ref,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetRouteTable(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.RouteTable1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.RouteTable1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.Network1Ref, resp.Spec.LocalRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateRouteTableV1(sim, secatest.RouteTableResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.RouteTable1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		LocalRef: secatest.Network1Ref,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	route := &network.RouteTable{
		Metadata: &network.RegionalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
			Name:      secatest.Network1Name,
		},
		Spec: network.RouteTableSpec{
			LocalRef: secatest.Network1Ref,
			Routes: []network.RouteSpec{
				{
					DestinationCidrBlock: secatest.CidrIpv4,
					TargetRef:            secatest.Instance1Ref,
				},
			},
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateRouteTable(ctx, route)
	require.NoError(t, err)

	assert.Equal(t, secatest.RouteTable1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.Network1Ref, resp.Spec.LocalRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestDeleteRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetRouteTableV1(sim, secatest.RouteTableResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.RouteTable1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		LocalRef: secatest.Network1Ref,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteRouteTableV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetRouteTable(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.RouteTable1Name})
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteRouteTable(ctx, resp)
	require.NoError(t, err)
}

// Internet Gateway

func TestListInternetGatewaysV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListInternetGatewaysV1(sim, secatest.InternetGatewayResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		EgressOnly: false,
		Status:     secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListInternetGateways(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.InternetGateway1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, false, *resp[0].Spec.EgressOnly)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetInternetGatewayV1(sim, secatest.InternetGatewayResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		EgressOnly: false,
		Status:     secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetInternetGateway(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, false, *resp.Spec.EgressOnly)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateInternetGatewayV1(sim, secatest.InternetGatewayResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		EgressOnly: false,
		Status:     secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	gtw := &network.InternetGateway{
		Metadata: &network.RegionalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
			Name:      secatest.Network1Name,
		},
		Spec: network.InternetGatewaySpec{},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateInternetGateway(ctx, gtw)
	require.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, false, *resp.Spec.EgressOnly)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestDeleteInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetInternetGatewayV1(sim, secatest.InternetGatewayResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		EgressOnly: false,
		Status:     secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteInternetGatewayV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetInternetGateway(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name})
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteInternetGateway(ctx, resp)
	require.NoError(t, err)
}

// Security Group

func TestListSecurityGroupsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListSecurityGroupsV1(sim, secatest.SecurityGroupResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RuleDirection: secatest.SecurityGroupRuleDirectionIngress,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListSecurityGroups(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.InternetGateway1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp[0].Spec.Rules[0].Direction))

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSecurityGroupV1(sim, secatest.SecurityGroupResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RuleDirection: secatest.SecurityGroupRuleDirectionIngress,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSecurityGroup(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp.Spec.Rules[0].Direction))

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateSecurityGroupV1(sim, secatest.SecurityGroupResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RuleDirection: secatest.SecurityGroupRuleDirectionIngress,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	group := &network.SecurityGroup{
		Metadata: &network.RegionalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
			Name:      secatest.Network1Name,
		},
		Spec: network.SecurityGroupSpec{
			Rules: []network.SecurityGroupRuleSpec{
				{
					Direction: network.Ingress,
					Version:   ptr.To(network.IPv4),
					Protocol:  ptr.To(network.Tcp),
					Ports: &network.Ports{
						From: ptr.To(network.Port(secatest.SecurityGroup1PortFrom)),
						To:   ptr.To(network.Port(secatest.SecurityGroup1PortTo)),
					},
				},
			},
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateSecurityGroup(ctx, group)
	require.NoError(t, err)

	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.SecurityGroupRuleDirectionIngress, string(resp.Spec.Rules[0].Direction))

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestDeleteSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSecurityGroupV1(sim, secatest.SecurityGroupResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		RuleDirection: secatest.SecurityGroupRuleDirectionIngress,
		Status:        secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteSecurityGroupV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSecurityGroup(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name})
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteSecurityGroup(ctx, resp)
	require.NoError(t, err)
}

// Nic

func TestListNicsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListNicsV1(sim, secatest.NicResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Nic1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SubnetRef: secatest.Subnet1Ref,
		Status:    secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListNics(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.Nic1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.Subnet1Ref, *&resp[0].Spec.SubnetRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNicV1(sim, secatest.NicResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Nic1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SubnetRef: secatest.Subnet1Ref,
		Status:    secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetNic(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Nic1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.Subnet1Ref, *&resp.Spec.SubnetRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateNicV1(sim, secatest.NicResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Nic1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SubnetRef: secatest.Subnet1Ref,
		Status:    secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	nic := &network.Nic{
		Metadata: &network.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
			Name:      secatest.Network1Name,
		},
		Spec: network.NicSpec{},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateNic(ctx, nic)
	require.NoError(t, err)

	assert.Equal(t, secatest.Nic1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.Subnet1Ref, *&resp.Spec.SubnetRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestDeleteNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNicV1(sim, secatest.NicResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.InternetGateway1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SubnetRef: secatest.Subnet1Ref,
		Status:    secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteNicV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetNic(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name})
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeleteNic(ctx, resp)
	require.NoError(t, err)
}

// Public Ip

func TestListPublicIpsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListPublicIpsV1(sim, secatest.PublicIpResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.PublicIp1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		Address: secatest.Address1,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListPublicIps(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.PublicIp1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.Address1, *resp[0].Spec.Address)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetPublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetPublicIpV1(sim, secatest.PublicIpResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.PublicIp1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		Address: secatest.Address1,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetPublicIp(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.PublicIp1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.Address1, *resp.Spec.Address)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdatePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdatePublicIpV1(sim, secatest.PublicIpResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.PublicIp1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		Address: secatest.Address1,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	ip := &network.PublicIp{
		Metadata: &network.RegionalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
			Name:      secatest.Network1Name,
		},
		Spec: network.PublicIpSpec{},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdatePublicIp(ctx, ip)
	require.NoError(t, err)

	assert.Equal(t, secatest.PublicIp1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.Address1, *resp.Spec.Address)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestDeletePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetPublicIpV1(sim, secatest.PublicIpResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.PublicIp1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		Address: secatest.Address1,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeletePublicIpV1(sim)
	secatest.ConfigureNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetPublicIp(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name})
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeletePublicIp(ctx, resp)
	require.NoError(t, err)
}
