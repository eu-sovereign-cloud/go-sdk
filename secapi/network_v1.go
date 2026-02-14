package secapi

import (
	"context"
	"net/http"

	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"k8s.io/utils/ptr"
)

// Interface

type NetworkV1 interface {
	// Network Sku
	ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.NetworkSku], error)
	ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.NetworkSku], error)

	GetSku(ctx context.Context, tref TenantReference) (*schema.NetworkSku, error)

	// Network
	ListNetworks(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Network], error)
	ListNetworksWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Network], error)

	GetNetwork(ctx context.Context, wref WorkspaceReference) (*schema.Network, error)
	GetNetworkUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Network, error)

	CreateOrUpdateNetworkWithParams(ctx context.Context, net *schema.Network, params *network.CreateOrUpdateNetworkParams) (*schema.Network, error)
	CreateOrUpdateNetwork(ctx context.Context, net *schema.Network) (*schema.Network, error)

	DeleteNetworkWithParams(ctx context.Context, net *schema.Network, params *network.DeleteNetworkParams) error
	DeleteNetwork(ctx context.Context, net *schema.Network) error

	// Subnet
	ListSubnets(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[schema.Subnet], error)
	ListSubnetsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID, opts *ListOptions) (*Iterator[schema.Subnet], error)

	GetSubnet(ctx context.Context, nref NetworkReference) (*schema.Subnet, error)
	GetSubnetUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Subnet, error)

	CreateOrUpdateSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.CreateOrUpdateSubnetParams) (*schema.Subnet, error)
	CreateOrUpdateSubnet(ctx context.Context, sub *schema.Subnet) (*schema.Subnet, error)

	DeleteSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.DeleteSubnetParams) error
	DeleteSubnet(ctx context.Context, sub *schema.Subnet) error

	// Route Table
	ListRouteTables(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[schema.RouteTable], error)
	ListRouteTablesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID, opts *ListOptions) (*Iterator[schema.RouteTable], error)

	GetRouteTable(ctx context.Context, nref NetworkReference) (*schema.RouteTable, error)
	GetRouteTableUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.RouteTable, error)

	CreateOrUpdateRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.CreateOrUpdateRouteTableParams) (*schema.RouteTable, error)
	CreateOrUpdateRouteTable(ctx context.Context, route *schema.RouteTable) (*schema.RouteTable, error)

	DeleteRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.DeleteRouteTableParams) error
	DeleteRouteTable(ctx context.Context, route *schema.RouteTable) error

	// Internet Gateway
	ListInternetGateways(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.InternetGateway], error)
	ListInternetGatewaysWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.InternetGateway], error)

	GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*schema.InternetGateway, error)
	GetInternetGatewayUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.InternetGateway, error)

	CreateOrUpdateInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.CreateOrUpdateInternetGatewayParams) (*schema.InternetGateway, error)
	CreateOrUpdateInternetGateway(ctx context.Context, gtw *schema.InternetGateway) (*schema.InternetGateway, error)

	DeleteInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.DeleteInternetGatewayParams) error
	DeleteInternetGateway(ctx context.Context, gtw *schema.InternetGateway) error

	// Security Group
	ListSecurityGroups(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.SecurityGroup], error)
	ListSecurityGroupsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.SecurityGroup], error)

	GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroup, error)
	GetSecurityGroupUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.SecurityGroup, error)

	CreateOrUpdateSecurityGroupWithParams(ctx context.Context, group *schema.SecurityGroup, params *network.CreateOrUpdateSecurityGroupParams) (*schema.SecurityGroup, error)
	CreateOrUpdateSecurityGroup(ctx context.Context, group *schema.SecurityGroup) (*schema.SecurityGroup, error)

	DeleteSecurityGroupWithParams(ctx context.Context, route *schema.SecurityGroup, params *network.DeleteSecurityGroupParams) error
	DeleteSecurityGroup(ctx context.Context, route *schema.SecurityGroup) error

	// Nic
	ListNics(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Nic], error)
	ListNicsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Nic], error)

	GetNic(ctx context.Context, wref WorkspaceReference) (*schema.Nic, error)
	GetNicUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Nic, error)

	CreateOrUpdateNicWithParams(ctx context.Context, nic *schema.Nic, params *network.CreateOrUpdateNicParams) (*schema.Nic, error)
	CreateOrUpdateNic(ctx context.Context, nic *schema.Nic) (*schema.Nic, error)

	DeleteNicWithParams(ctx context.Context, nic *schema.Nic, params *network.DeleteNicParams) error
	DeleteNic(ctx context.Context, nic *schema.Nic) error

	// Public Ip
	ListPublicIps(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.PublicIp], error)
	ListPublicIpsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.PublicIp], error)

	GetPublicIp(ctx context.Context, wref WorkspaceReference) (*schema.PublicIp, error)
	GetPublicIpUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.PublicIp, error)

	CreateOrUpdatePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.CreateOrUpdatePublicIpParams) (*schema.PublicIp, error)
	CreateOrUpdatePublicIp(ctx context.Context, ip *schema.PublicIp) (*schema.PublicIp, error)

	DeletePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.DeletePublicIpParams) error
	DeletePublicIp(ctx context.Context, ip *schema.PublicIp) error
}

// Dummy

type NetworkV1Dummy struct{}

func newNetworkV1Dummy() NetworkV1 {
	return &NetworkV1Dummy{}
}

/// Network Sku

func (api *NetworkV1Dummy) ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.NetworkSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.NetworkSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetSku(ctx context.Context, tref TenantReference) (*schema.NetworkSku, error) {
	return nil, ErrProviderNotAvailable
}

/// Network

func (api *NetworkV1Dummy) ListNetworks(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Network], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListNetworksWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Network], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetNetwork(ctx context.Context, wref WorkspaceReference) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetNetworkUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateNetworkWithParams(ctx context.Context, net *schema.Network, params *network.CreateOrUpdateNetworkParams) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateNetwork(ctx context.Context, net *schema.Network) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteNetworkWithParams(ctx context.Context, net *schema.Network, params *network.DeleteNetworkParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteNetwork(ctx context.Context, net *schema.Network) error {
	return ErrProviderNotAvailable
}

/// Subnet

func (api *NetworkV1Dummy) ListSubnets(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[schema.Subnet], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListSubnetsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID, opts *ListOptions) (*Iterator[schema.Subnet], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetSubnet(ctx context.Context, nref NetworkReference) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetSubnetUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.CreateOrUpdateSubnetParams) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateSubnet(ctx context.Context, sub *schema.Subnet) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.DeleteSubnetParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteSubnet(ctx context.Context, sub *schema.Subnet) error {
	return ErrProviderNotAvailable
}

/// Route Table

func (api *NetworkV1Dummy) ListRouteTables(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[schema.RouteTable], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListRouteTablesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID, opts *ListOptions) (*Iterator[schema.RouteTable], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetRouteTable(ctx context.Context, nref NetworkReference) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetRouteTableUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.CreateOrUpdateRouteTableParams) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateRouteTable(ctx context.Context, route *schema.RouteTable) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.DeleteRouteTableParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteRouteTable(ctx context.Context, route *schema.RouteTable) error {
	return ErrProviderNotAvailable
}

/// Internet Gateway

func (api *NetworkV1Dummy) ListInternetGateways(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.InternetGateway], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListInternetGatewaysWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.InternetGateway], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetInternetGatewayUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.CreateOrUpdateInternetGatewayParams) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateInternetGateway(ctx context.Context, gtw *schema.InternetGateway) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.DeleteInternetGatewayParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteInternetGateway(ctx context.Context, gtw *schema.InternetGateway) error {
	return ErrProviderNotAvailable
}

/// Security Group

func (api *NetworkV1Dummy) ListSecurityGroups(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.SecurityGroup], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListSecurityGroupsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.SecurityGroup], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetSecurityGroupUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateSecurityGroupWithParams(ctx context.Context, group *schema.SecurityGroup, params *network.CreateOrUpdateSecurityGroupParams) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateSecurityGroup(ctx context.Context, group *schema.SecurityGroup) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteSecurityGroupWithParams(ctx context.Context, route *schema.SecurityGroup, params *network.DeleteSecurityGroupParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteSecurityGroup(ctx context.Context, route *schema.SecurityGroup) error {
	return ErrProviderNotAvailable
}

/// Nic

func (api *NetworkV1Dummy) ListNics(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Nic], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListNicsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Nic], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetNic(ctx context.Context, wref WorkspaceReference) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetNicUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateNicWithParams(ctx context.Context, nic *schema.Nic, params *network.CreateOrUpdateNicParams) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdateNic(ctx context.Context, nic *schema.Nic) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteNicWithParams(ctx context.Context, nic *schema.Nic, params *network.DeleteNicParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeleteNic(ctx context.Context, nic *schema.Nic) error {
	return ErrProviderNotAvailable
}

/// Public Ip

func (api *NetworkV1Dummy) ListPublicIps(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.PublicIp], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) ListPublicIpsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.PublicIp], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetPublicIp(ctx context.Context, wref WorkspaceReference) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) GetPublicIpUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdatePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.CreateOrUpdatePublicIpParams) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) CreateOrUpdatePublicIp(ctx context.Context, ip *schema.PublicIp) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeletePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.DeletePublicIpParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Dummy) DeletePublicIp(ctx context.Context, ip *schema.PublicIp) error {
	return ErrProviderNotAvailable
}

// Impl

type NetworkV1Impl struct {
	API
	network network.ClientWithResponsesInterface
}

func newNetworkV1Impl(client *RegionalClient, networkUrl string) (NetworkV1, error) {
	network, err := network.NewClientWithResponses(networkUrl)
	if err != nil {
		return nil, err
	}

	return &NetworkV1Impl{API: API{authToken: client.authToken}, network: network}, nil
}

// Network Sku

func (api *NetworkV1Impl) ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.NetworkSku], error) {
	iter := Iterator[schema.NetworkSku]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.NetworkSku, *string, error) {
			resp, err := api.network.ListSkusWithResponse(ctx, schema.TenantPathParam(tid), &network.ListSkusParams{
				Accept:    ptr.To(network.ListSkusParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.NetworkSku], error) {
	iter := Iterator[schema.NetworkSku]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.NetworkSku, *string, error) {
			resp, err := api.network.ListSkusWithResponse(ctx, schema.TenantPathParam(tid), &network.ListSkusParams{
				Accept:    ptr.To(network.ListSkusParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetSku(ctx context.Context, tref TenantReference) (*schema.NetworkSku, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSkuWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

// Network

func (api *NetworkV1Impl) ListNetworks(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Network], error) {
	iter := Iterator[schema.Network]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Network, *string, error) {
			resp, err := api.network.ListNetworksWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListNetworksParams{
				Accept:    ptr.To(network.ListNetworksParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListNetworksWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Network], error) {
	iter := Iterator[schema.Network]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Network, *string, error) {
			resp, err := api.network.ListNetworksWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListNetworksParams{
				Accept:    ptr.To(network.ListNetworksParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetNetwork(ctx context.Context, wref WorkspaceReference) (*schema.Network, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetNetworkWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetNetworkUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Network, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Network]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.Network, error) {
			resp, err := api.network.GetNetworkWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *NetworkV1Impl) CreateOrUpdateNetworkWithParams(ctx context.Context, net *schema.Network, params *network.CreateOrUpdateNetworkParams) (*schema.Network, error) {
	if err := api.validateRegionalMetadata(net.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateNetworkWithResponse(ctx, net.Metadata.Tenant, net.Metadata.Workspace, net.Metadata.Name, params, *net, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateNetwork(ctx context.Context, net *schema.Network) (*schema.Network, error) {
	return api.CreateOrUpdateNetworkWithParams(ctx, net, nil)
}

func (api *NetworkV1Impl) DeleteNetworkWithParams(ctx context.Context, net *schema.Network, params *network.DeleteNetworkParams) error {
	if err := api.validateRegionalMetadata(net.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteNetworkWithResponse(ctx, net.Metadata.Tenant, net.Metadata.Workspace, net.Metadata.Name, params, api.loadRequestHeaders, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteNetwork(ctx context.Context, net *schema.Network) error {
	return api.DeleteNetworkWithParams(ctx, net, nil)
}

// Subnet

func (api *NetworkV1Impl) ListSubnets(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[schema.Subnet], error) {
	iter := Iterator[schema.Subnet]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Subnet, *string, error) {
			resp, err := api.network.ListSubnetsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), schema.NetworkPathParam(nid), &network.ListSubnetsParams{
				Accept:    ptr.To(network.ListSubnetsParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListSubnetsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID, opts *ListOptions) (*Iterator[schema.Subnet], error) {
	iter := Iterator[schema.Subnet]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Subnet, *string, error) {
			resp, err := api.network.ListSubnetsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), schema.NetworkPathParam(nid), &network.ListSubnetsParams{
				Accept:    ptr.To(network.ListSubnetsParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetSubnet(ctx context.Context, nref NetworkReference) (*schema.Subnet, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSubnetWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetSubnetUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Subnet, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Subnet]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.Subnet, error) {
			resp, err := api.network.GetSubnetWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *NetworkV1Impl) CreateOrUpdateSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.CreateOrUpdateSubnetParams) (*schema.Subnet, error) {
	if err := api.validateNetworkMetadata(sub.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateSubnetWithResponse(ctx, sub.Metadata.Tenant, sub.Metadata.Workspace, sub.Metadata.Network, sub.Metadata.Name, params, *sub, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateSubnet(ctx context.Context, sub *schema.Subnet) (*schema.Subnet, error) {
	return api.CreateOrUpdateSubnetWithParams(ctx, sub, nil)
}

func (api *NetworkV1Impl) DeleteSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.DeleteSubnetParams) error {
	if err := api.validateNetworkMetadata(sub.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteSubnetWithResponse(ctx, sub.Metadata.Tenant, sub.Metadata.Workspace, sub.Metadata.Network, sub.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteSubnet(ctx context.Context, sub *schema.Subnet) error {
	return api.DeleteSubnetWithParams(ctx, sub, nil)
}

// Route Table

func (api *NetworkV1Impl) ListRouteTables(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID) (*Iterator[schema.RouteTable], error) {
	iter := Iterator[schema.RouteTable]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.RouteTable, *string, error) {
			resp, err := api.network.ListRouteTablesWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), schema.NetworkPathParam(nid), &network.ListRouteTablesParams{
				Accept:    ptr.To(network.ListRouteTablesParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListRouteTablesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, nid NetworkID, opts *ListOptions) (*Iterator[schema.RouteTable], error) {
	iter := Iterator[schema.RouteTable]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.RouteTable, *string, error) {
			resp, err := api.network.ListRouteTablesWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), schema.NetworkPathParam(nid), &network.ListRouteTablesParams{
				Accept:    ptr.To(network.ListRouteTablesParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetRouteTable(ctx context.Context, nref NetworkReference) (*schema.RouteTable, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetRouteTableWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetRouteTableUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.RouteTable, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.RouteTable]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.RouteTable, error) {
			resp, err := api.network.GetRouteTableWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *NetworkV1Impl) CreateOrUpdateRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.CreateOrUpdateRouteTableParams) (*schema.RouteTable, error) {
	if err := api.validateNetworkMetadata(route.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateRouteTableWithResponse(ctx, route.Metadata.Tenant, route.Metadata.Workspace, route.Metadata.Network, route.Metadata.Name, params, *route, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateRouteTable(ctx context.Context, route *schema.RouteTable) (*schema.RouteTable, error) {
	return api.CreateOrUpdateRouteTableWithParams(ctx, route, nil)
}

func (api *NetworkV1Impl) DeleteRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.DeleteRouteTableParams) error {
	if err := api.validateNetworkMetadata(route.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteRouteTableWithResponse(ctx, route.Metadata.Tenant, route.Metadata.Workspace, route.Metadata.Network, route.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteRouteTable(ctx context.Context, route *schema.RouteTable) error {
	return api.DeleteRouteTableWithParams(ctx, route, nil)
}

// Internet Gateway

func (api *NetworkV1Impl) ListInternetGateways(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.InternetGateway], error) {
	iter := Iterator[schema.InternetGateway]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.InternetGateway, *string, error) {
			resp, err := api.network.ListInternetGatewaysWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListInternetGatewaysParams{
				Accept:    ptr.To(network.ListInternetGatewaysParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListInternetGatewaysWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.InternetGateway], error) {
	iter := Iterator[schema.InternetGateway]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.InternetGateway, *string, error) {
			resp, err := api.network.ListInternetGatewaysWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListInternetGatewaysParams{
				Accept:    ptr.To(network.ListInternetGatewaysParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*schema.InternetGateway, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetInternetGatewayWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetInternetGatewayUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.InternetGateway, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.InternetGateway]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.InternetGateway, error) {
			resp, err := api.network.GetInternetGatewayWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *NetworkV1Impl) CreateOrUpdateInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.CreateOrUpdateInternetGatewayParams) (*schema.InternetGateway, error) {
	if err := api.validateRegionalMetadata(gtw.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateInternetGatewayWithResponse(ctx, gtw.Metadata.Tenant, gtw.Metadata.Workspace, gtw.Metadata.Name, params, *gtw, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateInternetGateway(ctx context.Context, gtw *schema.InternetGateway) (*schema.InternetGateway, error) {
	return api.CreateOrUpdateInternetGatewayWithParams(ctx, gtw, nil)
}

func (api *NetworkV1Impl) DeleteInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.DeleteInternetGatewayParams) error {
	if err := api.validateRegionalMetadata(gtw.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteInternetGatewayWithResponse(ctx, gtw.Metadata.Tenant, gtw.Metadata.Workspace, gtw.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteInternetGateway(ctx context.Context, gtw *schema.InternetGateway) error {
	return api.DeleteInternetGatewayWithParams(ctx, gtw, nil)
}

// Security Group

func (api *NetworkV1Impl) ListSecurityGroups(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.SecurityGroup], error) {
	iter := Iterator[schema.SecurityGroup]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.SecurityGroup, *string, error) {
			resp, err := api.network.ListSecurityGroupsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListSecurityGroupsParams{
				Accept:    ptr.To(network.ListSecurityGroupsParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListSecurityGroupsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.SecurityGroup], error) {
	iter := Iterator[schema.SecurityGroup]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.SecurityGroup, *string, error) {
			resp, err := api.network.ListSecurityGroupsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListSecurityGroupsParams{
				Accept:    ptr.To(network.ListSecurityGroupsParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroup, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSecurityGroupWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetSecurityGroupUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.SecurityGroup, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.SecurityGroup]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.SecurityGroup, error) {
			resp, err := api.network.GetSecurityGroupWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *NetworkV1Impl) CreateOrUpdateSecurityGroupWithParams(ctx context.Context, group *schema.SecurityGroup, params *network.CreateOrUpdateSecurityGroupParams) (*schema.SecurityGroup, error) {
	if err := api.validateRegionalMetadata(group.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateSecurityGroupWithResponse(ctx, group.Metadata.Tenant, group.Metadata.Workspace, group.Metadata.Name, params, *group, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateSecurityGroup(ctx context.Context, group *schema.SecurityGroup) (*schema.SecurityGroup, error) {
	return api.CreateOrUpdateSecurityGroupWithParams(ctx, group, nil)
}

func (api *NetworkV1Impl) DeleteSecurityGroupWithParams(ctx context.Context, route *schema.SecurityGroup, params *network.DeleteSecurityGroupParams) error {
	if err := api.validateRegionalMetadata(route.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteSecurityGroupWithResponse(ctx, route.Metadata.Tenant, route.Metadata.Workspace, route.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteSecurityGroup(ctx context.Context, route *schema.SecurityGroup) error {
	return api.DeleteSecurityGroupWithParams(ctx, route, nil)
}

// Nic

func (api *NetworkV1Impl) ListNics(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Nic], error) {
	iter := Iterator[schema.Nic]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Nic, *string, error) {
			resp, err := api.network.ListNicsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListNicsParams{
				Accept:    ptr.To(network.ListNicsParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListNicsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Nic], error) {
	iter := Iterator[schema.Nic]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Nic, *string, error) {
			resp, err := api.network.ListNicsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListNicsParams{
				Accept:    ptr.To(network.ListNicsParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetNic(ctx context.Context, wref WorkspaceReference) (*schema.Nic, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetNicWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetNicUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Nic, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Nic]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.Nic, error) {
			resp, err := api.network.GetNicWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *NetworkV1Impl) CreateOrUpdateNicWithParams(ctx context.Context, nic *schema.Nic, params *network.CreateOrUpdateNicParams) (*schema.Nic, error) {
	if err := api.validateRegionalMetadata(nic.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateNicWithResponse(ctx, nic.Metadata.Tenant, nic.Metadata.Workspace, nic.Metadata.Name, params, *nic, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateNic(ctx context.Context, nic *schema.Nic) (*schema.Nic, error) {
	return api.CreateOrUpdateNicWithParams(ctx, nic, nil)
}

func (api *NetworkV1Impl) DeleteNicWithParams(ctx context.Context, nic *schema.Nic, params *network.DeleteNicParams) error {
	if err := api.validateRegionalMetadata(nic.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteNicWithResponse(ctx, nic.Metadata.Tenant, nic.Metadata.Workspace, nic.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteNic(ctx context.Context, nic *schema.Nic) error {
	return api.DeleteNicWithParams(ctx, nic, nil)
}

// Public Ip

func (api *NetworkV1Impl) ListPublicIps(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.PublicIp], error) {
	iter := Iterator[schema.PublicIp]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.PublicIp, *string, error) {
			resp, err := api.network.ListPublicIpsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListPublicIpsParams{
				Accept:    ptr.To(network.ListPublicIpsParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListPublicIpsWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.PublicIp], error) {
	iter := Iterator[schema.PublicIp]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.PublicIp, *string, error) {
			resp, err := api.network.ListPublicIpsWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &network.ListPublicIpsParams{
				Accept:    ptr.To(network.ListPublicIpsParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) GetPublicIp(ctx context.Context, wref WorkspaceReference) (*schema.PublicIp, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetPublicIpWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetPublicIpUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.PublicIp, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.PublicIp]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.PublicIp, error) {
			resp, err := api.network.GetPublicIpWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *NetworkV1Impl) CreateOrUpdatePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.CreateOrUpdatePublicIpParams) (*schema.PublicIp, error) {
	if err := api.validateRegionalMetadata(ip.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdatePublicIpWithResponse(ctx, ip.Metadata.Tenant, ip.Metadata.Workspace, ip.Metadata.Name, params, *ip, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdatePublicIp(ctx context.Context, ip *schema.PublicIp) (*schema.PublicIp, error) {
	return api.CreateOrUpdatePublicIpWithParams(ctx, ip, nil)
}

func (api *NetworkV1Impl) DeletePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.DeletePublicIpParams) error {
	if err := api.validateRegionalMetadata(ip.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeletePublicIpWithResponse(ctx, ip.Metadata.Tenant, ip.Metadata.Workspace, ip.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeletePublicIp(ctx context.Context, ip *schema.PublicIp) error {
	return api.DeletePublicIpWithParams(ctx, ip, nil)
}

func (api *NetworkV1Impl) validateRegionalMetadata(metadata *schema.RegionalWorkspaceResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == "" {
		return ErrNoMetatadaWorkspace
	}

	return nil
}

func (api *NetworkV1Impl) validateNetworkMetadata(metadata *schema.RegionalNetworkResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == "" {
		return ErrNoMetatadaWorkspace
	}

	if metadata.Network == "" {
		return ErrNoMetatadaNetwork
	}

	return nil
}
