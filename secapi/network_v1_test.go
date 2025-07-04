package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	"github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"

	"k8s.io/utils/ptr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Sku
func TestListNetworkSkusV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListSkusV1(sim, secatest.NetworkSkusResponseV1{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.NetworkSku1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListSkus(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.NetworkSku1Name, resp[0].Metadata.Name)
}
func TestGetNetworkSkuV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSkuV1(sim, secatest.NetworkSkuResponseV1{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.NetworkSku1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSku(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.NetworkSku1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.NetworkSku1Name, resp.Metadata.Name)
}

// Network
func TestListNetworksV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListNetworksV1(sim, secatest.NetworksResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	configTestNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListNetworks(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp[0].Metadata.Name)
}
func TestGetNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNetworkV1(sim, secatest.NetworkResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetNetwork(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Network1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestCreateOrUpdateOrUpdateNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateNetworkV1(sim, secatest.NetworkResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	network := &network.Network{
		Spec: network.NetworkSpec{
			Cidr:          network.Cidr{Ipv4: ptr.To(secatest.CidrIpv4)},
			RouteTableRef: secatest.RouteTable1Ref,
			SkuRef:        secatest.NetworkSku1Ref,
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateNetwork(ctx, wref, network)
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestDeleteNetworkV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNetworkV1(sim, secatest.NetworkResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	secatest.MockDeleteNetworkV1(sim)
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

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

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListSubnetsV1(sim, secatest.SubnetsResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Subnet1Name,
	})
	configTestNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListSubnets(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Subnet1Name, resp[0].Metadata.Name)
}
func TestGetSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSubnetV1(sim, secatest.SubnetResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Subnet1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSubnet(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Subnet1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Subnet1Name, resp.Metadata.Name)
}
func TestCreateOrUpdateSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateSubnetV1(sim, secatest.SubnetResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	sub := &network.Subnet{
		Spec: network.SubnetSpec{
			Cidr: network.Cidr{Ipv4: ptr.To(secatest.CidrIpv4)},
			Zone: secatest.ZoneA,
		},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateSubnet(ctx, wref, sub)
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestDeleteSubnetV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSubnetV1(sim, secatest.SubnetResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Subnet1Name,
	})
	secatest.MockDeleteSubnetV1(sim)
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

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

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListRouteTablesV1(sim, secatest.RouteTablesResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.RouteTable1Name,
	})
	configTestNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListRouteTables(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.RouteTable1Name, resp[0].Metadata.Name)
}
func TestGetRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetRouteTableV1(sim, secatest.RouteTableResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.RouteTable1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetRouteTable(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.RouteTable1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.RouteTable1Name, resp.Metadata.Name)
}
func TestCreateOrUpdateRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateRouteTableV1(sim, secatest.RouteTableResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	route := &network.RouteTable{
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
	resp, err := regionalClient.NetworkV1.CreateOrUpdateRouteTable(ctx, wref, route)
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestDeleteRouteTableV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetRouteTableV1(sim, secatest.RouteTableResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.RouteTable1Name,
	})
	secatest.MockDeleteRouteTableV1(sim)
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

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

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListInternetGatewaysV1(sim, secatest.InternetGatewaysResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.InternetGateway1Name,
	})
	configTestNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListInternetGateways(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.InternetGateway1Name, resp[0].Metadata.Name)
}
func TestGetInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetInternetGatewayV1(sim, secatest.InternetGatewayResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.InternetGateway1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetInternetGateway(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.InternetGateway1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.InternetGateway1Name, resp.Metadata.Name)
}
func TestCreateOrUpdateInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateInternetGatewayV1(sim, secatest.InternetGatewayResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	gtw := &network.InternetGateway{
		Spec: network.InternetGatewaySpec{},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateInternetGateway(ctx, wref, gtw)
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestDeleteInternetGatewayV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetInternetGatewayV1(sim, secatest.InternetGatewayResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.InternetGateway1Name,
	})
	secatest.MockDeleteInternetGatewayV1(sim)
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

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

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListSecurityGroupsV1(sim, secatest.SecurityGroupsResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.SecurityGroup1Name,
	})
	configTestNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListSecurityGroups(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.SecurityGroup1Name, resp[0].Metadata.Name)
}
func TestGetSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSecurityGroupV1(sim, secatest.SecurityGroupResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.SecurityGroup1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetSecurityGroup(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.SecurityGroup1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.SecurityGroup1Name, resp.Metadata.Name)
}
func TestCreateOrUpdateSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateSecurityGroupV1(sim, secatest.SecurityGroupResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	group := &network.SecurityGroup{
		Spec: network.SecurityGroupSpec{},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateSecurityGroup(ctx, wref, group)
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestDeleteSecurityGroupV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetSecurityGroupV1(sim, secatest.SecurityGroupResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.SecurityGroup1Name,
	})
	secatest.MockDeleteSecurityGroupV1(sim)
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

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

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListNicsV1(sim, secatest.NicsResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Nic1Name,
	})
	configTestNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListNics(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Nic1Name, resp[0].Metadata.Name)
}
func TestGetNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNicV1(sim, secatest.NicResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Nic1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetNic(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.Nic1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Nic1Name, resp.Metadata.Name)
}
func TestCreateOrUpdateNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateNicV1(sim, secatest.NicResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	nic := &network.Nic{
		Spec: network.NicSpec{},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdateNic(ctx, wref, nic)
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestDeleteNicV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetNicV1(sim, secatest.NicResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Nic1Name,
	})
	secatest.MockDeleteNicV1(sim)
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

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

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockListPublicIpsV1(sim, secatest.PublicIpsResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.PublicIp1Name,
	})
	configTestNetworkHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	iter, err := regionalClient.NetworkV1.ListPublicIps(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	require.Len(t, resp, 1)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.PublicIp1Name, resp[0].Metadata.Name)
}
func TestGetPublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetPublicIpV1(sim, secatest.PublicIpResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.PublicIp1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetPublicIp(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name})
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.PublicIp1Name, resp.Metadata.Name)
}
func TestCreateOrUpdatePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockCreateOrUpdatePublicIpV1(sim, secatest.PublicIpResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	})
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Network1Name,
	}
	ip := &network.PublicIp{
		Spec: network.PublicIpSpec{},
	}
	resp, err := regionalClient.NetworkV1.CreateOrUpdatePublicIp(ctx, wref, ip)
	require.NoError(t, err)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)
	assert.Equal(t, secatest.Network1Name, resp.Metadata.Name)
}
func TestDeletePublicIpV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	initTestRegionV1Handler(t, sm)

	sim := mocknetwork.NewMockServerInterface(t)
	secatest.MockGetPublicIpV1(sim, secatest.PublicIpResponseV1{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.PublicIp1Name,
	})
	secatest.MockDeletePublicIpV1(sim)
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	resp, err := regionalClient.NetworkV1.GetPublicIp(ctx, WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.PublicIp1Name})
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.NetworkV1.DeletePublicIp(ctx, resp)
	require.NoError(t, err)
}

func configTestNetworkHandler(sim *mocknetwork.MockServerInterface, sm *http.ServeMux) {
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderNetworkEndpoint,
		BaseRouter: sm,
	})
}
