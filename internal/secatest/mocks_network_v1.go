package secatest

import (
	"net/http"

	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"

	"github.com/stretchr/testify/mock"
)

// Network Sku
func MockListNetworkSkusV1(sim *mocknetwork.MockServerInterface, resp NetworkSkuResponseV1) {
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, params network.ListSkusParams) {
			configGetHttpResponse(w, networkSkusResponseTemplateV1, resp)
		})
}
func MockGetNetworkSkuV1(sim *mocknetwork.MockServerInterface, resp NetworkSkuResponseV1) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
			configGetHttpResponse(w, networkSkuResponseTemplateV1, resp)
		})
}

// Network
func MockListNetworksV1(sim *mocknetwork.MockServerInterface, resp NetworkResponseV1) {
	sim.EXPECT().ListNetworks(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListNetworksParams) {
			configGetHttpResponse(w, networksResponseTemplateV1, resp)
		})
}
func MockGetNetworkV1(sim *mocknetwork.MockServerInterface, resp NetworkResponseV1) {
	sim.EXPECT().GetNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configGetHttpResponse(w, networkResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateNetworkV1(sim *mocknetwork.MockServerInterface, resp NetworkResponseV1) {
	sim.EXPECT().CreateOrUpdateNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateNetworkParams) {
			configPutHttpResponse(w, networkResponseTemplateV1, resp)
		})
}
func MockDeleteNetworkV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteNetworkParams) {
			configDeleteHttpResponse(w)
		})
}

// Subnet
func MockListSubnetsV1(sim *mocknetwork.MockServerInterface, resp SubnetResponseV1) {
	sim.EXPECT().ListSubnets(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListSubnetsParams) {
			configGetHttpResponse(w, subnetsResponseTemplateV1, resp)
		})
}
func MockGetSubnetV1(sim *mocknetwork.MockServerInterface, resp SubnetResponseV1) {
	sim.EXPECT().GetSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configGetHttpResponse(w, subnetResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateSubnetV1(sim *mocknetwork.MockServerInterface, resp SubnetResponseV1) {
	sim.EXPECT().CreateOrUpdateSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateSubnetParams) {
			configPutHttpResponse(w, subnetResponseTemplateV1, resp)
		})
}
func MockDeleteSubnetV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteSubnetParams) {
			configDeleteHttpResponse(w)
		})
}

// Route Table
func MockListRouteTablesV1(sim *mocknetwork.MockServerInterface, resp RouteTableResponseV1) {
	sim.EXPECT().ListRouteTables(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListRouteTablesParams) {
			configGetHttpResponse(w, routeTablesResponseTemplateV1, resp)
		})
}
func MockGetRouteTableV1(sim *mocknetwork.MockServerInterface, resp RouteTableResponseV1) {
	sim.EXPECT().GetRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configGetHttpResponse(w, routeTableResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateRouteTableV1(sim *mocknetwork.MockServerInterface, resp RouteTableResponseV1) {
	sim.EXPECT().CreateOrUpdateRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateRouteTableParams) {
			configPutHttpResponse(w, routeTableResponseTemplateV1, resp)
		})
}
func MockDeleteRouteTableV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteRouteTableParams) {
			configDeleteHttpResponse(w)
		})
}

// Internet Gateway
func MockListInternetGatewaysV1(sim *mocknetwork.MockServerInterface, resp InternetGatewayResponseV1) {
	sim.EXPECT().ListInternetGateways(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListInternetGatewaysParams) {
			configGetHttpResponse(w, internetGatewaysResponseTemplateV1, resp)
		})
}
func MockGetInternetGatewayV1(sim *mocknetwork.MockServerInterface, resp InternetGatewayResponseV1) {
	sim.EXPECT().GetInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configGetHttpResponse(w, internetGatewayResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateInternetGatewayV1(sim *mocknetwork.MockServerInterface, resp InternetGatewayResponseV1) {
	sim.EXPECT().CreateOrUpdateInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateInternetGatewayParams) {
			configPutHttpResponse(w, internetGatewayResponseTemplateV1, resp)
		})
}
func MockDeleteInternetGatewayV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteInternetGatewayParams) {
			configDeleteHttpResponse(w)
		})
}

// Security Group
func MockListSecurityGroupsV1(sim *mocknetwork.MockServerInterface, resp SecurityGroupResponseV1) {
	sim.EXPECT().ListSecurityGroups(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListSecurityGroupsParams) {
			configGetHttpResponse(w, securityGroupsResponseTemplateV1, resp)
		})
}
func MockGetSecurityGroupV1(sim *mocknetwork.MockServerInterface, resp SecurityGroupResponseV1) {
	sim.EXPECT().GetSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configGetHttpResponse(w, securityGroupResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateSecurityGroupV1(sim *mocknetwork.MockServerInterface, resp SecurityGroupResponseV1) {
	sim.EXPECT().CreateOrUpdateSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateSecurityGroupParams) {
			configPutHttpResponse(w, securityGroupResponseTemplateV1, resp)
		})
}
func MockDeleteSecurityGroupV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteSecurityGroupParams) {
			configDeleteHttpResponse(w)
		})
}

// Nic
func MockListNicsV1(sim *mocknetwork.MockServerInterface, resp NicResponseV1) {
	sim.EXPECT().ListNics(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListNicsParams) {
			configGetHttpResponse(w, nicsResponseTemplateV1, resp)
		})
}
func MockGetNicV1(sim *mocknetwork.MockServerInterface, resp NicResponseV1) {
	sim.EXPECT().GetNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configGetHttpResponse(w, nicResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateNicV1(sim *mocknetwork.MockServerInterface, resp NicResponseV1) {
	sim.EXPECT().CreateOrUpdateNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateNicParams) {
			configPutHttpResponse(w, nicResponseTemplateV1, resp)
		})
}
func MockDeleteNicV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteNicParams) {
			configDeleteHttpResponse(w)
		})
}

// Public Ip
func MockListPublicIpsV1(sim *mocknetwork.MockServerInterface, resp PublicIpResponseV1) {
	sim.EXPECT().ListPublicIps(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListPublicIpsParams) {
			configGetHttpResponse(w, publicIpsResponseTemplateV1, resp)
		})
}
func MockGetPublicIpV1(sim *mocknetwork.MockServerInterface, resp PublicIpResponseV1) {
	sim.EXPECT().GetPublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			configGetHttpResponse(w, publicIpResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdatePublicIpV1(sim *mocknetwork.MockServerInterface, resp PublicIpResponseV1) {
	sim.EXPECT().CreateOrUpdatePublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace, name string, params network.CreateOrUpdatePublicIpParams) {
			configPutHttpResponse(w, publicIpResponseTemplateV1, resp)
		})
}
func MockDeletePublicIpV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeletePublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace, name string, params network.DeletePublicIpParams) {
			configDeleteHttpResponse(w)
		})
}
