package secatest

import (
	"net/http"

	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"

	"github.com/stretchr/testify/mock"
)

// Sku
func MockListSkusV1(sim *mocknetwork.MockServerInterface, resp NetworkSkusResponseV1) {
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, params network.ListSkusParams) {
			configTemplateMockResponse(w, http.StatusOK, networkSkusResponseV1, resp)
		})
}
func MockGetSkuV1(sim *mocknetwork.MockServerInterface, resp NetworkSkuResponseV1) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
			configTemplateMockResponse(w, http.StatusOK, networkSkuResponseV1, resp)
		})
}

// Network
func MockListNetworksV1(sim *mocknetwork.MockServerInterface, resp NetworksResponseV1) {
	sim.EXPECT().ListNetworks(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListNetworksParams) {
			configTemplateMockResponse(w, http.StatusOK, networksResponseV1, resp)
		})
}
func MockGetNetworkV1(sim *mocknetwork.MockServerInterface, resp NetworkResponseV1) {
	sim.EXPECT().GetNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configTemplateMockResponse(w, http.StatusOK, networkResponseV1, resp)
		})
}
func MockCreateOrUpdateNetworkV1(sim *mocknetwork.MockServerInterface, resp NetworkResponseV1) {
	sim.EXPECT().CreateOrUpdateNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateNetworkParams) {
			configTemplateMockResponse(w, http.StatusCreated, networkResponseV1, resp)
		})
}
func MockDeleteNetworkV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteNetworkParams) {
			configMockResponse(w, http.StatusAccepted)
		})
}

// Subnet
func MockListSubnetsV1(sim *mocknetwork.MockServerInterface, resp SubnetsResponseV1) {
	sim.EXPECT().ListSubnets(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListSubnetsParams) {
			configTemplateMockResponse(w, http.StatusOK, subnetsResponseV1, resp)
		})
}
func MockGetSubnetV1(sim *mocknetwork.MockServerInterface, resp SubnetResponseV1) {
	sim.EXPECT().GetSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configTemplateMockResponse(w, http.StatusOK, subnetResponseV1, resp)
		})
}
func MockCreateOrUpdateSubnetV1(sim *mocknetwork.MockServerInterface, resp SubnetResponseV1) {
	sim.EXPECT().CreateOrUpdateSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateSubnetParams) {
			configTemplateMockResponse(w, http.StatusCreated, subnetResponseV1, resp)
		})
}
func MockDeleteSubnetV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteSubnetParams) {
			configMockResponse(w, http.StatusAccepted)
		})
}

// Route Table
func MockListRouteTablesV1(sim *mocknetwork.MockServerInterface, resp RouteTablesResponseV1) {
	sim.EXPECT().ListRouteTables(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListRouteTablesParams) {
			configTemplateMockResponse(w, http.StatusOK, routeTablesResponseV1, resp)
		})
}
func MockGetRouteTableV1(sim *mocknetwork.MockServerInterface, resp RouteTableResponseV1) {
	sim.EXPECT().GetRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configTemplateMockResponse(w, http.StatusOK, routeTableResponseV1, resp)
		})
}
func MockCreateOrUpdateRouteTableV1(sim *mocknetwork.MockServerInterface, resp RouteTableResponseV1) {
	sim.EXPECT().CreateOrUpdateRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateRouteTableParams) {
			configTemplateMockResponse(w, http.StatusCreated, routeTableResponseV1, resp)
		})
}
func MockDeleteRouteTableV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteRouteTableParams) {
			configMockResponse(w, http.StatusAccepted)
		})
}

// Internet Gateway
func MockListInternetGatewaysV1(sim *mocknetwork.MockServerInterface, resp InternetGatewaysResponseV1) {
	sim.EXPECT().ListInternetGateways(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListInternetGatewaysParams) {
			configTemplateMockResponse(w, http.StatusOK, internetGatewaysResponseV1, resp)
		})
}
func MockGetInternetGatewayV1(sim *mocknetwork.MockServerInterface, resp InternetGatewayResponseV1) {
	sim.EXPECT().GetInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configTemplateMockResponse(w, http.StatusOK, internetGatewayResponseV1, resp)
		})
}
func MockCreateOrUpdateInternetGatewayV1(sim *mocknetwork.MockServerInterface, resp InternetGatewayResponseV1) {
	sim.EXPECT().CreateOrUpdateInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateInternetGatewayParams) {
			configTemplateMockResponse(w, http.StatusCreated, internetGatewayResponseV1, resp)
		})
}
func MockDeleteInternetGatewayV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteInternetGatewayParams) {
			configMockResponse(w, http.StatusAccepted)
		})
}

// Security Group
func MockListSecurityGroupsV1(sim *mocknetwork.MockServerInterface, resp SecurityGroupsResponseV1) {
	sim.EXPECT().ListSecurityGroups(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListSecurityGroupsParams) {
			configTemplateMockResponse(w, http.StatusOK, securityGroupsResponseV1, resp)
		})
}
func MockGetSecurityGroupV1(sim *mocknetwork.MockServerInterface, resp SecurityGroupResponseV1) {
	sim.EXPECT().GetSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configTemplateMockResponse(w, http.StatusOK, securityGroupResponseV1, resp)
		})
}
func MockCreateOrUpdateSecurityGroupV1(sim *mocknetwork.MockServerInterface, resp SecurityGroupResponseV1) {
	sim.EXPECT().CreateOrUpdateSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateSecurityGroupParams) {
			configTemplateMockResponse(w, http.StatusCreated, securityGroupResponseV1, resp)
		})
}
func MockDeleteSecurityGroupV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteSecurityGroupParams) {
			configMockResponse(w, http.StatusAccepted)
		})
}

// Nic
func MockListNicsV1(sim *mocknetwork.MockServerInterface, resp NicsResponseV1) {
	sim.EXPECT().ListNics(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListNicsParams) {
			configTemplateMockResponse(w, http.StatusOK, nicsResponseV1, resp)
		})
}
func MockGetNicV1(sim *mocknetwork.MockServerInterface, resp NicResponseV1) {
	sim.EXPECT().GetNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configTemplateMockResponse(w, http.StatusOK, nicResponseV1, resp)
		})
}
func MockCreateOrUpdateNicV1(sim *mocknetwork.MockServerInterface, resp NicResponseV1) {
	sim.EXPECT().CreateOrUpdateNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateNicParams) {
			configTemplateMockResponse(w, http.StatusCreated, nicResponseV1, resp)
		})
}
func MockDeleteNicV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteNicParams) {
			configMockResponse(w, http.StatusAccepted)
		})
}

// Public Ip
func MockListPublicIpsV1(sim *mocknetwork.MockServerInterface, resp PublicIpsResponseV1) {
	sim.EXPECT().ListPublicIps(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListPublicIpsParams) {
			configTemplateMockResponse(w, http.StatusOK, publicIpsResponseV1, resp)
		})
}
func MockGetPublicIpV1(sim *mocknetwork.MockServerInterface, resp PublicIpResponseV1) {
	sim.EXPECT().GetPublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configTemplateMockResponse(w, http.StatusOK, publicIpResponseV1, resp)
		})
}
func MockCreateOrUpdatePublicIpV1(sim *mocknetwork.MockServerInterface, resp PublicIpResponseV1) {
	sim.EXPECT().CreateOrUpdatePublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace, name string, params network.CreateOrUpdatePublicIpParams) {
			configTemplateMockResponse(w, http.StatusCreated, publicIpResponseV1, resp)
		})
}
func MockDeletePublicIpV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeletePublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace, name string, params network.DeletePublicIpParams) {
			configMockResponse(w, http.StatusAccepted)
		})
}
