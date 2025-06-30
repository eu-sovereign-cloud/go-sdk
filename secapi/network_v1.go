package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
)

type NetworkV1 struct {
	network network.ClientWithResponsesInterface
}

func newNetworkV1(networkUrl string) (*NetworkV1, error) {
	network, err := network.NewClientWithResponses(networkUrl)
	if err != nil {
		return nil, err
	}

	return &NetworkV1{network: network}, nil
}

func validateNetworkZonalMetadataV1(metadata *network.ZonalResourceMetadata) {
	if metadata == nil {
		panic(ErrNoMetatada)
	}

	if metadata.Workspace == nil {
		panic(ErrNoMetatadaWorkspace)
	}

	if metadata.Tenant == "" {
		panic(ErrNoMetatadaTenant)
	}
}

func validateNetworkRegionalMetadataV1(metadata *network.RegionalResourceMetadata) {
	if metadata == nil {
		panic(ErrNoMetatada)
	}

	if metadata.Workspace == nil {
		panic(ErrNoMetatadaWorkspace)
	}

	if metadata.Tenant == "" {
		panic(ErrNoMetatadaTenant)
	}
}

func (api *NetworkV1) ListSkus(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.NetworkSku], error) {
	iter := Iterator[network.NetworkSku]{
		fn: func(ctx context.Context, skipToken *string) ([]network.NetworkSku, *string, error) {
			resp, err := api.network.ListSkusWithResponse(ctx, network.Tenant(tid), &network.ListSkusParams{
				Accept: ptr.To(network.ListSkusParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetSku(ctx context.Context, tref TenantReference) (*network.NetworkSku, error) {
	resp, err := api.network.GetSkuWithResponse(ctx, network.Tenant(tref.Tenant), tref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) ListNetworks(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.Network], error) {
	iter := Iterator[network.Network]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Network, *string, error) {
			resp, err := api.network.ListNetworksWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListNetworksParams{
				Accept: ptr.To(network.ListNetworksParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetNetwork(ctx context.Context, wref WorkspaceReference) (*network.Network, error) {
	resp, err := api.network.GetNetworkWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdateNetwork(ctx context.Context, net *network.Network) error {
	validateNetworkRegionalMetadataV1(net.Metadata)

	resp, err := api.network.CreateOrUpdateNetworkWithResponse(ctx, net.Metadata.Tenant, *net.Metadata.Workspace, net.Metadata.Name,
		&network.CreateOrUpdateNetworkParams{
			IfUnmodifiedSince: &net.Metadata.ResourceVersion,
		}, *net)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteNetwork(ctx context.Context, net *network.Network) error {
	validateNetworkRegionalMetadataV1(net.Metadata)

	resp, err := api.network.DeleteNetworkWithResponse(ctx, net.Metadata.Tenant, *net.Metadata.Workspace, net.Metadata.Name, &network.DeleteNetworkParams{
		IfUnmodifiedSince: &net.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) ListSubnets(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.Subnet], error) {
	iter := Iterator[network.Subnet]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Subnet, *string, error) {
			resp, err := api.network.ListSubnetsWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListSubnetsParams{
				Accept: ptr.To(network.Applicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetSubnet(ctx context.Context, wref WorkspaceReference) (*network.Subnet, error) {
	resp, err := api.network.GetSubnetWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdateSubnet(ctx context.Context, sub *network.Subnet) error {
	validateNetworkZonalMetadataV1(sub.Metadata)

	resp, err := api.network.CreateOrUpdateSubnetWithResponse(ctx, sub.Metadata.Tenant, *sub.Metadata.Workspace, sub.Metadata.Name,
		&network.CreateOrUpdateSubnetParams{
			IfUnmodifiedSince: &sub.Metadata.ResourceVersion,
		}, *sub)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteSubnet(ctx context.Context, sub *network.Subnet) error {
	validateNetworkZonalMetadataV1(sub.Metadata)

	resp, err := api.network.DeleteSubnetWithResponse(ctx, sub.Metadata.Tenant, *sub.Metadata.Workspace, sub.Metadata.Name, &network.DeleteSubnetParams{
		IfUnmodifiedSince: &sub.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) ListRouteTables(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.RouteTable], error) {
	iter := Iterator[network.RouteTable]{
		fn: func(ctx context.Context, skipToken *string) ([]network.RouteTable, *string, error) {
			resp, err := api.network.ListRouteTablesWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListRouteTablesParams{
				Accept: ptr.To(network.ListRouteTablesParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetRouteTable(ctx context.Context, wref WorkspaceReference) (*network.RouteTable, error) {
	resp, err := api.network.GetRouteTableWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdateRouteTable(ctx context.Context, route *network.RouteTable) error {
	validateNetworkRegionalMetadataV1(route.Metadata)

	resp, err := api.network.CreateOrUpdateRouteTableWithResponse(ctx, route.Metadata.Tenant, *route.Metadata.Workspace, route.Metadata.Name,
		&network.CreateOrUpdateRouteTableParams{
			IfUnmodifiedSince: &route.Metadata.ResourceVersion,
		}, *route)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteRouteTable(ctx context.Context, route *network.RouteTable) error {
	validateNetworkRegionalMetadataV1(route.Metadata)

	resp, err := api.network.DeleteRouteTableWithResponse(ctx, route.Metadata.Tenant, *route.Metadata.Workspace, route.Metadata.Name, &network.DeleteRouteTableParams{
		IfUnmodifiedSince: &route.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) ListInternetGateways(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.InternetGateway], error) {
	iter := Iterator[network.InternetGateway]{
		fn: func(ctx context.Context, skipToken *string) ([]network.InternetGateway, *string, error) {
			resp, err := api.network.ListInternetGatewaysWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListInternetGatewaysParams{
				Accept: ptr.To(network.ListInternetGatewaysParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*network.InternetGateway, error) {
	resp, err := api.network.GetInternetGatewayWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdateInternetGateway(ctx context.Context, gtw *network.InternetGateway) error {
	validateNetworkRegionalMetadataV1(gtw.Metadata)

	resp, err := api.network.CreateOrUpdateInternetGatewayWithResponse(ctx, gtw.Metadata.Tenant, *gtw.Metadata.Workspace, gtw.Metadata.Name,
		&network.CreateOrUpdateInternetGatewayParams{
			IfUnmodifiedSince: &gtw.Metadata.ResourceVersion,
		}, *gtw)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteInternetGateway(ctx context.Context, gtw *network.InternetGateway) error {
	validateNetworkRegionalMetadataV1(gtw.Metadata)

	resp, err := api.network.DeleteInternetGatewayWithResponse(ctx, gtw.Metadata.Tenant, *gtw.Metadata.Workspace, gtw.Metadata.Name, &network.DeleteInternetGatewayParams{
		IfUnmodifiedSince: &gtw.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) ListSecurityGroups(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.SecurityGroup], error) {
	iter := Iterator[network.SecurityGroup]{
		fn: func(ctx context.Context, skipToken *string) ([]network.SecurityGroup, *string, error) {
			resp, err := api.network.ListSecurityGroupsWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListSecurityGroupsParams{
				Accept: ptr.To(network.ListSecurityGroupsParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*network.SecurityGroup, error) {
	resp, err := api.network.GetSecurityGroupWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdateSecurityGroup(ctx context.Context, route *network.SecurityGroup) error {
	validateNetworkRegionalMetadataV1(route.Metadata)

	resp, err := api.network.CreateOrUpdateSecurityGroupWithResponse(ctx, route.Metadata.Tenant, *route.Metadata.Workspace, route.Metadata.Name,
		&network.CreateOrUpdateSecurityGroupParams{
			IfUnmodifiedSince: &route.Metadata.ResourceVersion,
		}, *route)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteSecurityGroup(ctx context.Context, route *network.SecurityGroup) error {
	validateNetworkRegionalMetadataV1(route.Metadata)

	resp, err := api.network.DeleteSecurityGroupWithResponse(ctx, route.Metadata.Tenant, *route.Metadata.Workspace, route.Metadata.Name, &network.DeleteSecurityGroupParams{
		IfUnmodifiedSince: &route.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) ListNics(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.Nic], error) {
	iter := Iterator[network.Nic]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Nic, *string, error) {
			resp, err := api.network.ListNicsWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListNicsParams{
				Accept: ptr.To(network.ListNicsParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetNic(ctx context.Context, wref WorkspaceReference) (*network.Nic, error) {
	resp, err := api.network.GetNicWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdateNic(ctx context.Context, nic *network.Nic) error {
	validateNetworkZonalMetadataV1(nic.Metadata)

	resp, err := api.network.CreateOrUpdateNicWithResponse(ctx, nic.Metadata.Tenant, *nic.Metadata.Workspace, nic.Metadata.Name,
		&network.CreateOrUpdateNicParams{
			IfUnmodifiedSince: &nic.Metadata.ResourceVersion,
		}, *nic)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeleteNic(ctx context.Context, nic *network.Nic) error {
	validateNetworkZonalMetadataV1(nic.Metadata)

	resp, err := api.network.DeleteNicWithResponse(ctx, nic.Metadata.Tenant, *nic.Metadata.Workspace, nic.Metadata.Name, &network.DeleteNicParams{
		IfUnmodifiedSince: &nic.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) ListPublicIps(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.PublicIp], error) {
	iter := Iterator[network.PublicIp]{
		fn: func(ctx context.Context, skipToken *string) ([]network.PublicIp, *string, error) {
			resp, err := api.network.ListPublicIpsWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListPublicIpsParams{
				Accept: ptr.To(network.ListPublicIpsParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1) GetPublicIp(ctx context.Context, wref WorkspaceReference) (*network.PublicIp, error) {
	resp, err := api.network.GetPublicIpWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdatePublicIp(ctx context.Context, ip *network.PublicIp) error {
	validateNetworkRegionalMetadataV1(ip.Metadata)

	resp, err := api.network.CreateOrUpdatePublicIpWithResponse(ctx, ip.Metadata.Tenant, *ip.Metadata.Workspace, ip.Metadata.Name,
		&network.CreateOrUpdatePublicIpParams{
			IfUnmodifiedSince: &ip.Metadata.ResourceVersion,
		}, *ip)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *NetworkV1) DeletePublicIp(ctx context.Context, ip *network.PublicIp) error {
	validateNetworkRegionalMetadataV1(ip.Metadata)

	resp, err := api.network.DeletePublicIpWithResponse(ctx, ip.Metadata.Tenant, *ip.Metadata.Workspace, ip.Metadata.Name, &network.DeletePublicIpParams{
		IfUnmodifiedSince: &ip.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}
