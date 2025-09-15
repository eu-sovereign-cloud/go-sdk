package secapi

import (
	"context"
	"fmt"
	"net/http"

	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"

	"k8s.io/utils/ptr"
)

type NetworkV1 struct {
	API
	network network.ClientWithResponsesInterface
}

// Network Sku

func (api *NetworkV1) ListSkus(ctx context.Context, tid TenantID) (*Iterator[network.NetworkSku], error) {
	iter := Iterator[network.NetworkSku]{
		fn: func(ctx context.Context, skipToken *string) ([]network.NetworkSku, *string, error) {
			resp, err := api.network.ListSkusWithResponse(ctx, network.TenantPathParam(tid), &network.ListSkusParams{
				Accept: ptr.To(network.ListSkusParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetSku(ctx context.Context, tref TenantReference) (*network.NetworkSku, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSkuWithResponse(ctx, network.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

// Network

func (api *NetworkV1) ListNetworks(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.Network], error) {
	iter := Iterator[network.Network]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Network, *string, error) {
			resp, err := api.network.ListNetworksWithResponse(ctx, network.TenantPathParam(tid), network.WorkspacePathParam(wid), &network.ListNetworksParams{
				Accept: ptr.To(network.ListNetworksParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetNetwork(ctx context.Context, wref WorkspaceReference) (*network.Network, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetNetworkWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *NetworkV1) CreateOrUpdateNetworkWithParams(ctx context.Context, wref WorkspaceReference, net *network.Network, params *network.CreateOrUpdateNetworkParams) (*network.Network, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateNetworkWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, params, *net, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *NetworkV1) CreateOrUpdateNetwork(ctx context.Context, wref WorkspaceReference, net *network.Network) (*network.Network, error) {
	return api.CreateOrUpdateNetworkWithParams(ctx, wref, net, nil)
}

func (api *NetworkV1) DeleteNetworkWithParams(ctx context.Context, net *network.Network, params *network.DeleteNetworkParams) error {
	if err := api.validateRegionalMetadata(net.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteNetworkWithResponse(ctx, net.Metadata.Tenant, *net.Metadata.Workspace, net.Metadata.Name, params, api.loadRequestHeaders, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteNetwork(ctx context.Context, net *network.Network) error {
	return api.DeleteNetworkWithParams(ctx, net, nil)
}

// Subnet

func (api *NetworkV1) ListSubnets(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[network.Subnet], error) {
	iter := Iterator[network.Subnet]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Subnet, *string, error) {
			resp, err := api.network.ListSubnetsWithResponse(ctx, network.TenantPathParam(tid), network.WorkspacePathParam(wid), network.NetworkPathParam(nid), &network.ListSubnetsParams{
				Accept: ptr.To(network.ListSubnetsParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetSubnet(ctx context.Context, nref NetworkReference) (*network.Subnet, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSubnetWithResponse(ctx, network.TenantPathParam(nref.Tenant), network.WorkspacePathParam(nref.Workspace), network.NetworkPathParam(nref.Network), network.ResourcePathParam(nref.Name), api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *NetworkV1) CreateOrUpdateSubnetWithParams(ctx context.Context, nref NetworkReference, sub *network.Subnet, params *network.CreateOrUpdateSubnetParams) (*network.Subnet, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateSubnetWithResponse(ctx, network.TenantPathParam(nref.Tenant), network.WorkspacePathParam(nref.Workspace), network.NetworkPathParam(nref.Network), network.ResourcePathParam(nref.Name), params, *sub, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *NetworkV1) CreateOrUpdateSubnet(ctx context.Context, nref NetworkReference, sub *network.Subnet) (*network.Subnet, error) {
	return api.CreateOrUpdateSubnetWithParams(ctx, nref, sub, nil)
}

func (api *NetworkV1) DeleteSubnetWithParams(ctx context.Context, sub *network.Subnet, params *network.DeleteSubnetParams) error {
	if err := api.validateNetworkZonalMetadata(sub.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteSubnetWithResponse(ctx, sub.Metadata.Tenant, *sub.Metadata.Workspace, sub.Metadata.Network, sub.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteSubnet(ctx context.Context, sub *network.Subnet) error {
	return api.DeleteSubnetWithParams(ctx, sub, nil)
}

// Route Table

func (api *NetworkV1) ListRouteTables(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[network.RouteTable], error) {
	iter := Iterator[network.RouteTable]{
		fn: func(ctx context.Context, skipToken *string) ([]network.RouteTable, *string, error) {
			resp, err := api.network.ListRouteTablesWithResponse(ctx, network.TenantPathParam(tid), network.WorkspacePathParam(wid), network.NetworkPathParam(nid), &network.ListRouteTablesParams{
				Accept: ptr.To(network.ListRouteTablesParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetRouteTable(ctx context.Context, nref NetworkReference) (*network.RouteTable, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetRouteTableWithResponse(ctx, network.TenantPathParam(nref.Tenant), network.WorkspacePathParam(nref.Workspace), network.NetworkPathParam(nref.Network), network.ResourcePathParam(nref.Name), api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *NetworkV1) CreateOrUpdateRouteTableWithParams(ctx context.Context, nref NetworkReference, route *network.RouteTable, params *network.CreateOrUpdateRouteTableParams) (*network.RouteTable, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateRouteTableWithResponse(ctx, network.TenantPathParam(nref.Tenant), network.WorkspacePathParam(nref.Workspace), network.NetworkPathParam(nref.Network), network.ResourcePathParam(nref.Name), params, *route, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *NetworkV1) CreateOrUpdateRouteTable(ctx context.Context, nref NetworkReference, route *network.RouteTable) (*network.RouteTable, error) {
	return api.CreateOrUpdateRouteTableWithParams(ctx, nref, route, nil)
}

func (api *NetworkV1) DeleteRouteTableWithParams(ctx context.Context, route *network.RouteTable, params *network.DeleteRouteTableParams) error {
	if err := api.validateNetworkRegionalMetadata(route.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteRouteTableWithResponse(ctx, route.Metadata.Tenant, *route.Metadata.Workspace, route.Metadata.Network, route.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteRouteTable(ctx context.Context, route *network.RouteTable) error {
	return api.DeleteRouteTableWithParams(ctx, route, nil)
}

// Internet Gateway

func (api *NetworkV1) ListInternetGateways(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.InternetGateway], error) {
	iter := Iterator[network.InternetGateway]{
		fn: func(ctx context.Context, skipToken *string) ([]network.InternetGateway, *string, error) {
			resp, err := api.network.ListInternetGatewaysWithResponse(ctx, network.TenantPathParam(tid), network.WorkspacePathParam(wid), &network.ListInternetGatewaysParams{
				Accept: ptr.To(network.ListInternetGatewaysParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*network.InternetGateway, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetInternetGatewayWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *NetworkV1) CreateOrUpdateInternetGatewayWithParams(ctx context.Context, wref WorkspaceReference, gtw *network.InternetGateway, params *network.CreateOrUpdateInternetGatewayParams) (*network.InternetGateway, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateInternetGatewayWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, params, *gtw, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *NetworkV1) CreateOrUpdateInternetGateway(ctx context.Context, wref WorkspaceReference, gtw *network.InternetGateway) (*network.InternetGateway, error) {
	return api.CreateOrUpdateInternetGatewayWithParams(ctx, wref, gtw, nil)
}

func (api *NetworkV1) DeleteInternetGatewayWithParams(ctx context.Context, gtw *network.InternetGateway, params *network.DeleteInternetGatewayParams) error {
	if err := api.validateRegionalMetadata(gtw.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteInternetGatewayWithResponse(ctx, gtw.Metadata.Tenant, *gtw.Metadata.Workspace, gtw.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteInternetGateway(ctx context.Context, gtw *network.InternetGateway) error {
	return api.DeleteInternetGatewayWithParams(ctx, gtw, nil)
}

// Security Group

func (api *NetworkV1) ListSecurityGroups(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.SecurityGroup], error) {
	iter := Iterator[network.SecurityGroup]{
		fn: func(ctx context.Context, skipToken *string) ([]network.SecurityGroup, *string, error) {
			resp, err := api.network.ListSecurityGroupsWithResponse(ctx, network.TenantPathParam(tid), network.WorkspacePathParam(wid), &network.ListSecurityGroupsParams{
				Accept: ptr.To(network.Applicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*network.SecurityGroup, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSecurityGroupWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *NetworkV1) CreateOrUpdateSecurityGroupWithParams(ctx context.Context, wref WorkspaceReference, group *network.SecurityGroup, params *network.CreateOrUpdateSecurityGroupParams) (*network.SecurityGroup, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateSecurityGroupWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, params, *group, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *NetworkV1) CreateOrUpdateSecurityGroup(ctx context.Context, wref WorkspaceReference, group *network.SecurityGroup) (*network.SecurityGroup, error) {
	return api.CreateOrUpdateSecurityGroupWithParams(ctx, wref, group, nil)
}

func (api *NetworkV1) DeleteSecurityGroupWithParams(ctx context.Context, route *network.SecurityGroup, params *network.DeleteSecurityGroupParams) error {
	if err := api.validateRegionalMetadata(route.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteSecurityGroupWithResponse(ctx, route.Metadata.Tenant, *route.Metadata.Workspace, route.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteSecurityGroup(ctx context.Context, route *network.SecurityGroup) error {
	return api.DeleteSecurityGroupWithParams(ctx, route, nil)
}

// Nic

func (api *NetworkV1) ListNics(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.Nic], error) {
	iter := Iterator[network.Nic]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Nic, *string, error) {
			resp, err := api.network.ListNicsWithResponse(ctx, network.TenantPathParam(tid), network.WorkspacePathParam(wid), &network.ListNicsParams{
				Accept: ptr.To(network.ListNicsParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetNic(ctx context.Context, wref WorkspaceReference) (*network.Nic, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetNicWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *NetworkV1) CreateOrUpdateNicWithParams(ctx context.Context, wref WorkspaceReference, nic *network.Nic, params *network.CreateOrUpdateNicParams) (*network.Nic, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateNicWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, params, *nic, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *NetworkV1) CreateOrUpdateNic(ctx context.Context, wref WorkspaceReference, nic *network.Nic) (*network.Nic, error) {
	return api.CreateOrUpdateNicWithParams(ctx, wref, nic, nil)
}

func (api *NetworkV1) DeleteNicWithParams(ctx context.Context, nic *network.Nic, params *network.DeleteNicParams) error {
	if err := api.validateZonalMetadata(nic.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteNicWithResponse(ctx, nic.Metadata.Tenant, *nic.Metadata.Workspace, nic.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteNic(ctx context.Context, nic *network.Nic) error {
	return api.DeleteNicWithParams(ctx, nic, nil)
}

// Public Ip

func (api *NetworkV1) ListPublicIps(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.PublicIp], error) {
	iter := Iterator[network.PublicIp]{
		fn: func(ctx context.Context, skipToken *string) ([]network.PublicIp, *string, error) {
			resp, err := api.network.ListPublicIpsWithResponse(ctx, network.TenantPathParam(tid), network.WorkspacePathParam(wid), &network.ListPublicIpsParams{
				Accept: ptr.To(network.ListPublicIpsParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetPublicIp(ctx context.Context, wref WorkspaceReference) (*network.PublicIp, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetPublicIpWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *NetworkV1) CreateOrUpdatePublicIpWithParams(ctx context.Context, wref WorkspaceReference, ip *network.PublicIp, params *network.CreateOrUpdatePublicIpParams) (*network.PublicIp, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdatePublicIpWithResponse(ctx, network.TenantPathParam(wref.Tenant), network.WorkspacePathParam(wref.Workspace), wref.Name, params, *ip, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *NetworkV1) CreateOrUpdatePublicIp(ctx context.Context, wref WorkspaceReference, ip *network.PublicIp) (*network.PublicIp, error) {
	return api.CreateOrUpdatePublicIpWithParams(ctx, wref, ip, nil)
}

func (api *NetworkV1) DeletePublicIpWithParams(ctx context.Context, ip *network.PublicIp, params *network.DeletePublicIpParams) error {
	if err := api.validateRegionalMetadata(ip.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeletePublicIpWithResponse(ctx, ip.Metadata.Tenant, *ip.Metadata.Workspace, ip.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeletePublicIp(ctx context.Context, ip *network.PublicIp) error {
	return api.DeletePublicIpWithParams(ctx, ip, nil)
}

func (api *NetworkV1) BuildReferenceURN(urn string) (*network.Reference, error) {
	urnRef := network.ReferenceURN(urn)

	ref := &network.Reference{}
	if err := ref.FromReferenceURN(urnRef); err != nil {
		return nil, fmt.Errorf("error building referenceURN from URN %s: %s", urn, err)
	}

	return ref, nil
}

func (api *NetworkV1) validateZonalMetadata(metadata *network.ZonalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == nil {
		return ErrNoMetatadaWorkspace
	}

	return nil
}

func (api *NetworkV1) validateRegionalMetadata(metadata *network.RegionalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == nil {
		return ErrNoMetatadaWorkspace
	}

	return nil
}

func (api *NetworkV1) validateNetworkRegionalMetadata(metadata *network.NetworkRegionalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == nil {
		return ErrNoMetatadaWorkspace
	}

	return nil
}

func (api *NetworkV1) validateNetworkZonalMetadata(metadata *network.NetworkZonalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == nil {
		return ErrNoMetatadaWorkspace
	}

	return nil
}

func newNetworkV1(client *RegionalClient, networkUrl string) (*NetworkV1, error) {
	network, err := network.NewClientWithResponses(networkUrl)
	if err != nil {
		return nil, err
	}

	return &NetworkV1{API: API{authToken: client.authToken}, network: network}, nil
}
