package secatest

import (
	"net/http"

	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/mock"
)

// Network Sku
func MockListNetworkSkusV1(sim *mocknetwork.MockServerInterface, resp []schema.NetworkSku) {
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, params network.ListSkusParams) {
			iter := &network.SkuIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetNetworkSkuV1(sim *mocknetwork.MockServerInterface, resp *schema.NetworkSku) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

// Network
func MockListNetworksV1(sim *mocknetwork.MockServerInterface, resp []schema.Network) {
	sim.EXPECT().ListNetworks(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListNetworksParams) {
			iter := &network.NetworkIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetNetworkV1(sim *mocknetwork.MockServerInterface, resp *schema.Network) {
	sim.EXPECT().GetNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateNetworkV1(sim *mocknetwork.MockServerInterface, resp *schema.Network) {
	sim.EXPECT().CreateOrUpdateNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateNetworkParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteNetworkV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteNetwork(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteNetworkParams) {
			configDeleteHttpResponse(w)
		})
}

// Subnet
func MockListSubnetsV1(sim *mocknetwork.MockServerInterface, resp []schema.Subnet) {
	sim.EXPECT().ListSubnets(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, params network.ListSubnetsParams) {
			iter := &network.SubnetIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetSubnetV1(sim *mocknetwork.MockServerInterface, resp *schema.Subnet) {
	sim.EXPECT().GetSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, params schema.ResourcePathParam) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateSubnetV1(sim *mocknetwork.MockServerInterface, resp *schema.Subnet) {
	sim.EXPECT().CreateOrUpdateSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, name schema.ResourcePathParam, params network.CreateOrUpdateSubnetParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteSubnetV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteSubnet(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, name schema.ResourcePathParam, params network.DeleteSubnetParams) {
			configDeleteHttpResponse(w)
		})
}

// Route Table
func MockListRouteTablesV1(sim *mocknetwork.MockServerInterface, resp []schema.RouteTable) {
	sim.EXPECT().ListRouteTables(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, params network.ListRouteTablesParams) {
			iter := &network.RouteTableIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetRouteTableV1(sim *mocknetwork.MockServerInterface, resp *schema.RouteTable) {
	sim.EXPECT().GetRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, name schema.ResourcePathParam) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateRouteTableV1(sim *mocknetwork.MockServerInterface, resp *schema.RouteTable) {
	sim.EXPECT().CreateOrUpdateRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, name schema.ResourcePathParam, params network.CreateOrUpdateRouteTableParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteRouteTableV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteRouteTable(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, networkPath schema.NetworkPathParam, name schema.ResourcePathParam, params network.DeleteRouteTableParams) {
			configDeleteHttpResponse(w)
		})
}

// Internet Gateway
func MockListInternetGatewaysV1(sim *mocknetwork.MockServerInterface, resp []schema.InternetGateway) {
	sim.EXPECT().ListInternetGateways(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListInternetGatewaysParams) {
			iter := &network.InternetGatewayIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetInternetGatewayV1(sim *mocknetwork.MockServerInterface, resp *schema.InternetGateway) {
	sim.EXPECT().GetInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateInternetGatewayV1(sim *mocknetwork.MockServerInterface, resp *schema.InternetGateway) {
	sim.EXPECT().CreateOrUpdateInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateInternetGatewayParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteInternetGatewayV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteInternetGateway(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteInternetGatewayParams) {
			configDeleteHttpResponse(w)
		})
}

// Security Group
func MockListSecurityGroupsV1(sim *mocknetwork.MockServerInterface, resp []schema.SecurityGroup) {
	sim.EXPECT().ListSecurityGroups(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListSecurityGroupsParams) {
			iter := &network.SecurityGroupIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetSecurityGroupV1(sim *mocknetwork.MockServerInterface, resp *schema.SecurityGroup) {
	sim.EXPECT().GetSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateSecurityGroupV1(sim *mocknetwork.MockServerInterface, resp *schema.SecurityGroup) {
	sim.EXPECT().CreateOrUpdateSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateSecurityGroupParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteSecurityGroupV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteSecurityGroup(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteSecurityGroupParams) {
			configDeleteHttpResponse(w)
		})
}

// Nic
func MockListNicsV1(sim *mocknetwork.MockServerInterface, resp []schema.Nic) {
	sim.EXPECT().ListNics(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListNicsParams) {
			iter := &network.NicIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetNicV1(sim *mocknetwork.MockServerInterface, resp *schema.Nic) {
	sim.EXPECT().GetNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateNicV1(sim *mocknetwork.MockServerInterface, resp *schema.Nic) {
	sim.EXPECT().CreateOrUpdateNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.CreateOrUpdateNicParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteNicV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeleteNic(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string, params network.DeleteNicParams) {
			configDeleteHttpResponse(w)
		})
}

// Public Ip
func MockListPublicIpsV1(sim *mocknetwork.MockServerInterface, resp []schema.PublicIp) {
	sim.EXPECT().ListPublicIps(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, params network.ListPublicIpsParams) {
			iter := &network.PublicIpIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetPublicIpV1(sim *mocknetwork.MockServerInterface, resp *schema.PublicIp) {
	sim.EXPECT().GetPublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdatePublicIpV1(sim *mocknetwork.MockServerInterface, resp *schema.PublicIp) {
	sim.EXPECT().CreateOrUpdatePublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace, name string, params network.CreateOrUpdatePublicIpParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeletePublicIpV1(sim *mocknetwork.MockServerInterface) {
	sim.EXPECT().DeletePublicIp(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, workspace, name string, params network.DeletePublicIpParams) {
			configDeleteHttpResponse(w)
		})
}
