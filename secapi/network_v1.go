package secapi

import (
	"context"

	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Interface

type NetworkV1 interface {
	// Network Sku
	ListSkusWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.NetworkSku], error)
	ListSkus(ctx context.Context, tpath TenantPath) (*Iterator[schema.NetworkSku], error)

	GetSku(ctx context.Context, tref TenantReference) (*schema.NetworkSku, error)

	// Network
	ListNetworksWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.Network], error)
	ListNetworks(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.Network], error)

	GetNetwork(ctx context.Context, wref WorkspaceReference) (*schema.Network, error)
	GetNetworkUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Network, error)

	WatchNetworkUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error

	CreateOrUpdateNetworkWithParams(ctx context.Context, net *schema.Network, params *network.CreateOrUpdateNetworkParams) (*schema.Network, error)
	CreateOrUpdateNetwork(ctx context.Context, net *schema.Network) (*schema.Network, error)

	DeleteNetworkWithParams(ctx context.Context, net *schema.Network, params *network.DeleteNetworkParams) error
	DeleteNetwork(ctx context.Context, net *schema.Network) error

	// Subnet
	ListSubnetsWithOptions(ctx context.Context, npath NetworkPath, options *ListOptions) (*Iterator[schema.Subnet], error)
	ListSubnets(ctx context.Context, npath NetworkPath) (*Iterator[schema.Subnet], error)

	GetSubnet(ctx context.Context, nref NetworkReference) (*schema.Subnet, error)
	GetSubnetUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Subnet, error)

	WatchSubnetUntilDeleted(ctx context.Context, nref NetworkReference, config ResourceObserverConfig) error

	CreateOrUpdateSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.CreateOrUpdateSubnetParams) (*schema.Subnet, error)
	CreateOrUpdateSubnet(ctx context.Context, sub *schema.Subnet) (*schema.Subnet, error)

	DeleteSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.DeleteSubnetParams) error
	DeleteSubnet(ctx context.Context, sub *schema.Subnet) error

	// Route Table
	ListRouteTablesWithOptions(ctx context.Context, npath NetworkPath, options *ListOptions) (*Iterator[schema.RouteTable], error)
	ListRouteTables(ctx context.Context, npath NetworkPath) (*Iterator[schema.RouteTable], error)

	GetRouteTable(ctx context.Context, nref NetworkReference) (*schema.RouteTable, error)
	GetRouteTableUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.RouteTable, error)

	WatchRouteTableUntilDeleted(ctx context.Context, nref NetworkReference, config ResourceObserverConfig) error

	CreateOrUpdateRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.CreateOrUpdateRouteTableParams) (*schema.RouteTable, error)
	CreateOrUpdateRouteTable(ctx context.Context, route *schema.RouteTable) (*schema.RouteTable, error)

	DeleteRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.DeleteRouteTableParams) error
	DeleteRouteTable(ctx context.Context, route *schema.RouteTable) error

	// Internet Gateway
	ListInternetGatewaysWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.InternetGateway], error)
	ListInternetGateways(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.InternetGateway], error)

	GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*schema.InternetGateway, error)
	GetInternetGatewayUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.InternetGateway, error)

	WatchInternetGatewayUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error

	CreateOrUpdateInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.CreateOrUpdateInternetGatewayParams) (*schema.InternetGateway, error)
	CreateOrUpdateInternetGateway(ctx context.Context, gtw *schema.InternetGateway) (*schema.InternetGateway, error)

	DeleteInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.DeleteInternetGatewayParams) error
	DeleteInternetGateway(ctx context.Context, gtw *schema.InternetGateway) error

	// Security Group Rule
	ListSecurityGroupRulesWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.SecurityGroupRule], error)
	ListSecurityGroupRules(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.SecurityGroupRule], error)

	GetSecurityGroupRule(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroupRule, error)
	GetSecurityGroupRuleUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.SecurityGroupRule, error)

	WatchSecurityGroupRuleUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error

	CreateOrUpdateSecurityGroupRuleWithParams(ctx context.Context, group *schema.SecurityGroupRule, params *network.CreateOrUpdateSecurityGroupRuleParams) (*schema.SecurityGroupRule, error)
	CreateOrUpdateSecurityGroupRule(ctx context.Context, group *schema.SecurityGroupRule) (*schema.SecurityGroupRule, error)

	DeleteSecurityGroupRuleWithParams(ctx context.Context, rule *schema.SecurityGroupRule, params *network.DeleteSecurityGroupRuleParams) error
	DeleteSecurityGroupRule(ctx context.Context, rule *schema.SecurityGroupRule) error

	// Security Group
	ListSecurityGroupsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.SecurityGroup], error)
	ListSecurityGroups(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.SecurityGroup], error)

	GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroup, error)
	GetSecurityGroupUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.SecurityGroup, error)

	WatchSecurityGroupUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error

	CreateOrUpdateSecurityGroupWithParams(ctx context.Context, group *schema.SecurityGroup, params *network.CreateOrUpdateSecurityGroupParams) (*schema.SecurityGroup, error)
	CreateOrUpdateSecurityGroup(ctx context.Context, group *schema.SecurityGroup) (*schema.SecurityGroup, error)

	DeleteSecurityGroupWithParams(ctx context.Context, route *schema.SecurityGroup, params *network.DeleteSecurityGroupParams) error
	DeleteSecurityGroup(ctx context.Context, route *schema.SecurityGroup) error

	// Nic
	ListNicsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.Nic], error)
	ListNics(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.Nic], error)

	GetNic(ctx context.Context, wref WorkspaceReference) (*schema.Nic, error)
	GetNicUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Nic, error)

	WatchNicUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error

	CreateOrUpdateNicWithParams(ctx context.Context, nic *schema.Nic, params *network.CreateOrUpdateNicParams) (*schema.Nic, error)
	CreateOrUpdateNic(ctx context.Context, nic *schema.Nic) (*schema.Nic, error)

	DeleteNicWithParams(ctx context.Context, nic *schema.Nic, params *network.DeleteNicParams) error
	DeleteNic(ctx context.Context, nic *schema.Nic) error

	// Public Ip
	ListPublicIpsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.PublicIp], error)
	ListPublicIps(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.PublicIp], error)

	GetPublicIp(ctx context.Context, wref WorkspaceReference) (*schema.PublicIp, error)
	GetPublicIpUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.PublicIp, error)

	WatchPublicIpUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error

	CreateOrUpdatePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.CreateOrUpdatePublicIpParams) (*schema.PublicIp, error)
	CreateOrUpdatePublicIp(ctx context.Context, ip *schema.PublicIp) (*schema.PublicIp, error)

	DeletePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.DeletePublicIpParams) error
	DeletePublicIp(ctx context.Context, ip *schema.PublicIp) error
}

// Unavailable

type NetworkV1Unavailable struct{}

func newNetworkV1Unavailable() NetworkV1 {
	return &NetworkV1Unavailable{}
}

/// Network Sku

func (api *NetworkV1Unavailable) ListSkusWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.NetworkSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListSkus(ctx context.Context, tpath TenantPath) (*Iterator[schema.NetworkSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetSku(ctx context.Context, tref TenantReference) (*schema.NetworkSku, error) {
	return nil, ErrProviderNotAvailable
}

/// Network

func (api *NetworkV1Unavailable) ListNetworksWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.Network], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListNetworks(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.Network], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetNetwork(ctx context.Context, wref WorkspaceReference) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetNetworkUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchNetworkUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateNetworkWithParams(ctx context.Context, net *schema.Network, params *network.CreateOrUpdateNetworkParams) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateNetwork(ctx context.Context, net *schema.Network) (*schema.Network, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteNetworkWithParams(ctx context.Context, net *schema.Network, params *network.DeleteNetworkParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteNetwork(ctx context.Context, net *schema.Network) error {
	return ErrProviderNotAvailable
}

/// Subnet

func (api *NetworkV1Unavailable) ListSubnetsWithOptions(ctx context.Context, npath NetworkPath, options *ListOptions) (*Iterator[schema.Subnet], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListSubnets(ctx context.Context, npath NetworkPath) (*Iterator[schema.Subnet], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetSubnet(ctx context.Context, nref NetworkReference) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetSubnetUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchSubnetUntilDeleted(ctx context.Context, nref NetworkReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.CreateOrUpdateSubnetParams) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateSubnet(ctx context.Context, sub *schema.Subnet) (*schema.Subnet, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.DeleteSubnetParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteSubnet(ctx context.Context, sub *schema.Subnet) error {
	return ErrProviderNotAvailable
}

/// Route Table

func (api *NetworkV1Unavailable) ListRouteTablesWithOptions(ctx context.Context, npath NetworkPath, options *ListOptions) (*Iterator[schema.RouteTable], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListRouteTables(ctx context.Context, npath NetworkPath) (*Iterator[schema.RouteTable], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetRouteTable(ctx context.Context, nref NetworkReference) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetRouteTableUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchRouteTableUntilDeleted(ctx context.Context, nref NetworkReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.CreateOrUpdateRouteTableParams) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateRouteTable(ctx context.Context, route *schema.RouteTable) (*schema.RouteTable, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.DeleteRouteTableParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteRouteTable(ctx context.Context, route *schema.RouteTable) error {
	return ErrProviderNotAvailable
}

/// Internet Gateway

func (api *NetworkV1Unavailable) ListInternetGatewaysWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.InternetGateway], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListInternetGateways(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.InternetGateway], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetInternetGatewayUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchInternetGatewayUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.CreateOrUpdateInternetGatewayParams) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateInternetGateway(ctx context.Context, gtw *schema.InternetGateway) (*schema.InternetGateway, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.DeleteInternetGatewayParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteInternetGateway(ctx context.Context, gtw *schema.InternetGateway) error {
	return ErrProviderNotAvailable
}

/// Security Group Rule

func (api *NetworkV1Unavailable) ListSecurityGroupRulesWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.SecurityGroupRule], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListSecurityGroupRules(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.SecurityGroupRule], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetSecurityGroupRule(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroupRule, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetSecurityGroupRuleUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.SecurityGroupRule, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchSecurityGroupRuleUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateSecurityGroupRuleWithParams(ctx context.Context, group *schema.SecurityGroupRule, params *network.CreateOrUpdateSecurityGroupRuleParams) (*schema.SecurityGroupRule, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateSecurityGroupRule(ctx context.Context, group *schema.SecurityGroupRule) (*schema.SecurityGroupRule, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteSecurityGroupRuleWithParams(ctx context.Context, rule *schema.SecurityGroupRule, params *network.DeleteSecurityGroupRuleParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteSecurityGroupRule(ctx context.Context, rule *schema.SecurityGroupRule) error {
	return ErrProviderNotAvailable
}

/// Security Group

func (api *NetworkV1Unavailable) ListSecurityGroupsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.SecurityGroup], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListSecurityGroups(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.SecurityGroup], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetSecurityGroupUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchSecurityGroupUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateSecurityGroupWithParams(ctx context.Context, group *schema.SecurityGroup, params *network.CreateOrUpdateSecurityGroupParams) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateSecurityGroup(ctx context.Context, group *schema.SecurityGroup) (*schema.SecurityGroup, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteSecurityGroupWithParams(ctx context.Context, route *schema.SecurityGroup, params *network.DeleteSecurityGroupParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteSecurityGroup(ctx context.Context, route *schema.SecurityGroup) error {
	return ErrProviderNotAvailable
}

/// Nic

func (api *NetworkV1Unavailable) ListNicsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.Nic], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListNics(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.Nic], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetNic(ctx context.Context, wref WorkspaceReference) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetNicUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchNicUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateNicWithParams(ctx context.Context, nic *schema.Nic, params *network.CreateOrUpdateNicParams) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdateNic(ctx context.Context, nic *schema.Nic) (*schema.Nic, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteNicWithParams(ctx context.Context, nic *schema.Nic, params *network.DeleteNicParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeleteNic(ctx context.Context, nic *schema.Nic) error {
	return ErrProviderNotAvailable
}

/// Public Ip

func (api *NetworkV1Unavailable) ListPublicIpsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.PublicIp], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) ListPublicIps(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.PublicIp], error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetPublicIp(ctx context.Context, wref WorkspaceReference) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) GetPublicIpUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) WatchPublicIpUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdatePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.CreateOrUpdatePublicIpParams) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) CreateOrUpdatePublicIp(ctx context.Context, ip *schema.PublicIp) (*schema.PublicIp, error) {
	return nil, ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeletePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.DeletePublicIpParams) error {
	return ErrProviderNotAvailable
}

func (api *NetworkV1Unavailable) DeletePublicIp(ctx context.Context, ip *schema.PublicIp) error {
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

func (api *NetworkV1Impl) ListSkusWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.NetworkSku], error) {
	if err := tpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.NetworkSku]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.NetworkSku, *string, error) {
			var params *network.ListSkusParams
			if options == nil {
				params = &network.ListSkusParams{
					Accept:    AcceptHeaderJson[network.ListSkusParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListSkusParams{
					Accept:    AcceptHeaderJson[network.ListSkusParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListSkusWithResponse(ctx, schema.TenantPathParam(tpath.Tenant), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListSkus(ctx context.Context, tpath TenantPath) (*Iterator[schema.NetworkSku], error) {
	return api.ListSkusWithOptions(ctx, tpath, nil)
}

func (api *NetworkV1Impl) GetSku(ctx context.Context, tref TenantReference) (*schema.NetworkSku, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSkuWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

// Network

func (api *NetworkV1Impl) ListNetworksWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.Network], error) {
	if err := wpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.Network]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Network, *string, error) {
			var params *network.ListNetworksParams
			if options == nil {
				params = &network.ListNetworksParams{
					Accept:    AcceptHeaderJson[network.ListNetworksParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListNetworksParams{
					Accept:    AcceptHeaderJson[network.ListNetworksParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListNetworksWithResponse(ctx, schema.TenantPathParam(wpath.Tenant), schema.WorkspacePathParam(wpath.Workspace), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListNetworks(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.Network], error) {
	return api.ListNetworksWithOptions(ctx, wpath, nil)
}

func (api *NetworkV1Impl) GetNetwork(ctx context.Context, wref WorkspaceReference) (*schema.Network, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetNetworkWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetNetworkUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Network, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Network]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.Network, error) {
			resp, err := api.network.GetNetworkWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchNetworkUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	if err := wref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Network]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetNetworkWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdateNetworkWithParams(ctx context.Context, net *schema.Network, params *network.CreateOrUpdateNetworkParams) (*schema.Network, error) {
	if err := api.validateRegionalMetadata(net.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateNetworkWithResponse(ctx, net.Metadata.Tenant, net.Metadata.Workspace, net.Metadata.Name, params, *net, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
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

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteNetwork(ctx context.Context, net *schema.Network) error {
	return api.DeleteNetworkWithParams(ctx, net, nil)
}

// Subnet

func (api *NetworkV1Impl) ListSubnetsWithOptions(ctx context.Context, npath NetworkPath, options *ListOptions) (*Iterator[schema.Subnet], error) {
	if err := npath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.Subnet]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Subnet, *string, error) {
			var params *network.ListSubnetsParams
			if options == nil {
				params = &network.ListSubnetsParams{
					Accept:    AcceptHeaderJson[network.ListSubnetsParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListSubnetsParams{
					Accept:    AcceptHeaderJson[network.ListSubnetsParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListSubnetsWithResponse(ctx, schema.TenantPathParam(npath.Tenant), schema.WorkspacePathParam(npath.Workspace), schema.NetworkPathParam(npath.Network), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListSubnets(ctx context.Context, npath NetworkPath) (*Iterator[schema.Subnet], error) {
	return api.ListSubnetsWithOptions(ctx, npath, nil)
}

func (api *NetworkV1Impl) GetSubnet(ctx context.Context, nref NetworkReference) (*schema.Subnet, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSubnetWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetSubnetUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Subnet, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Subnet]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.Subnet, error) {
			resp, err := api.network.GetSubnetWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchSubnetUntilDeleted(ctx context.Context, nref NetworkReference, config ResourceObserverConfig) error {
	if err := nref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Subnet]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetSubnetWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdateSubnetWithParams(ctx context.Context, sub *schema.Subnet, params *network.CreateOrUpdateSubnetParams) (*schema.Subnet, error) {
	if err := api.validateNetworkMetadata(sub.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateSubnetWithResponse(ctx, sub.Metadata.Tenant, sub.Metadata.Workspace, sub.Metadata.Network, sub.Metadata.Name, params, *sub, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
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

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteSubnet(ctx context.Context, sub *schema.Subnet) error {
	return api.DeleteSubnetWithParams(ctx, sub, nil)
}

// Route Table

func (api *NetworkV1Impl) ListRouteTablesWithOptions(ctx context.Context, npath NetworkPath, options *ListOptions) (*Iterator[schema.RouteTable], error) {
	if err := npath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.RouteTable]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.RouteTable, *string, error) {
			var params *network.ListRouteTablesParams
			if options == nil {
				params = &network.ListRouteTablesParams{
					Accept:    AcceptHeaderJson[network.ListRouteTablesParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListRouteTablesParams{
					Accept:    AcceptHeaderJson[network.ListRouteTablesParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListRouteTablesWithResponse(ctx, schema.TenantPathParam(npath.Tenant), schema.WorkspacePathParam(npath.Workspace), schema.NetworkPathParam(npath.Network), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListRouteTables(ctx context.Context, npath NetworkPath) (*Iterator[schema.RouteTable], error) {
	return api.ListRouteTablesWithOptions(ctx, npath, nil)
}

func (api *NetworkV1Impl) GetRouteTable(ctx context.Context, nref NetworkReference) (*schema.RouteTable, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetRouteTableWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetRouteTableUntilState(ctx context.Context, nref NetworkReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.RouteTable, error) {
	if err := nref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.RouteTable]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.RouteTable, error) {
			resp, err := api.network.GetRouteTableWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchRouteTableUntilDeleted(ctx context.Context, nref NetworkReference, config ResourceObserverConfig) error {
	if err := nref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.RouteTable]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetRouteTableWithResponse(ctx, schema.TenantPathParam(nref.Tenant), schema.WorkspacePathParam(nref.Workspace), schema.NetworkPathParam(nref.Network), nref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdateRouteTableWithParams(ctx context.Context, route *schema.RouteTable, params *network.CreateOrUpdateRouteTableParams) (*schema.RouteTable, error) {
	if err := api.validateNetworkMetadata(route.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateRouteTableWithResponse(ctx, route.Metadata.Tenant, route.Metadata.Workspace, route.Metadata.Network, route.Metadata.Name, params, *route, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
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

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteRouteTable(ctx context.Context, route *schema.RouteTable) error {
	return api.DeleteRouteTableWithParams(ctx, route, nil)
}

// Internet Gateway

func (api *NetworkV1Impl) ListInternetGatewaysWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.InternetGateway], error) {
	if err := wpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.InternetGateway]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.InternetGateway, *string, error) {
			var params *network.ListInternetGatewaysParams
			if options == nil {
				params = &network.ListInternetGatewaysParams{
					Accept:    AcceptHeaderJson[network.ListInternetGatewaysParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListInternetGatewaysParams{
					Accept:    AcceptHeaderJson[network.ListInternetGatewaysParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListInternetGatewaysWithResponse(ctx, schema.TenantPathParam(wpath.Tenant), schema.WorkspacePathParam(wpath.Workspace), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListInternetGateways(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.InternetGateway], error) {
	return api.ListInternetGatewaysWithOptions(ctx, wpath, nil)
}

func (api *NetworkV1Impl) GetInternetGateway(ctx context.Context, wref WorkspaceReference) (*schema.InternetGateway, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetInternetGatewayWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetInternetGatewayUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.InternetGateway, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.InternetGateway]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.InternetGateway, error) {
			resp, err := api.network.GetInternetGatewayWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchInternetGatewayUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	if err := wref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.InternetGateway]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetInternetGatewayWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdateInternetGatewayWithParams(ctx context.Context, gtw *schema.InternetGateway, params *network.CreateOrUpdateInternetGatewayParams) (*schema.InternetGateway, error) {
	if err := api.validateRegionalMetadata(gtw.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateInternetGatewayWithResponse(ctx, gtw.Metadata.Tenant, gtw.Metadata.Workspace, gtw.Metadata.Name, params, *gtw, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
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

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteInternetGateway(ctx context.Context, gtw *schema.InternetGateway) error {
	return api.DeleteInternetGatewayWithParams(ctx, gtw, nil)
}

// Security Group Rules

func (api *NetworkV1Impl) ListSecurityGroupRulesWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.SecurityGroupRule], error) {
	if err := wpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.SecurityGroupRule]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.SecurityGroupRule, *string, error) {
			var params *network.ListSecurityGroupRulesParams
			if options == nil {
				params = &network.ListSecurityGroupRulesParams{
					Accept:    AcceptHeaderJson[network.ListSecurityGroupRulesParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListSecurityGroupRulesParams{
					Accept:    AcceptHeaderJson[network.ListSecurityGroupRulesParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListSecurityGroupRulesWithResponse(ctx, schema.TenantPathParam(wpath.Tenant), schema.WorkspacePathParam(wpath.Workspace), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListSecurityGroupRules(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.SecurityGroupRule], error) {
	return api.ListSecurityGroupRulesWithOptions(ctx, wpath, nil)
}

func (api *NetworkV1Impl) GetSecurityGroupRule(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroupRule, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSecurityGroupRuleWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetSecurityGroupRuleUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.SecurityGroupRule, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.SecurityGroupRule]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.SecurityGroupRule, error) {
			resp, err := api.network.GetSecurityGroupRuleWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchSecurityGroupRuleUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	if err := wref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.SecurityGroupRule]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetSecurityGroupRuleWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdateSecurityGroupRuleWithParams(ctx context.Context, group *schema.SecurityGroupRule, params *network.CreateOrUpdateSecurityGroupRuleParams) (*schema.SecurityGroupRule, error) {
	if err := api.validateRegionalMetadata(group.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateSecurityGroupRuleWithResponse(ctx, group.Metadata.Tenant, group.Metadata.Workspace, group.Metadata.Name, params, *group, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateSecurityGroupRule(ctx context.Context, group *schema.SecurityGroupRule) (*schema.SecurityGroupRule, error) {
	return api.CreateOrUpdateSecurityGroupRuleWithParams(ctx, group, nil)
}

func (api *NetworkV1Impl) DeleteSecurityGroupRuleWithParams(ctx context.Context, rule *schema.SecurityGroupRule, params *network.DeleteSecurityGroupRuleParams) error {
	if err := api.validateRegionalMetadata(rule.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteSecurityGroupRuleWithResponse(ctx, rule.Metadata.Tenant, rule.Metadata.Workspace, rule.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteSecurityGroupRule(ctx context.Context, rule *schema.SecurityGroupRule) error {
	return api.DeleteSecurityGroupRuleWithParams(ctx, rule, nil)
}

// Security Group

func (api *NetworkV1Impl) ListSecurityGroupsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.SecurityGroup], error) {
	if err := wpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.SecurityGroup]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.SecurityGroup, *string, error) {
			var params *network.ListSecurityGroupsParams
			if options == nil {
				params = &network.ListSecurityGroupsParams{
					Accept:    AcceptHeaderJson[network.ListSecurityGroupsParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListSecurityGroupsParams{
					Accept:    AcceptHeaderJson[network.ListSecurityGroupsParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListSecurityGroupsWithResponse(ctx, schema.TenantPathParam(wpath.Tenant), schema.WorkspacePathParam(wpath.Workspace), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListSecurityGroups(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.SecurityGroup], error) {
	return api.ListSecurityGroupsWithOptions(ctx, wpath, nil)
}

func (api *NetworkV1Impl) GetSecurityGroup(ctx context.Context, wref WorkspaceReference) (*schema.SecurityGroup, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetSecurityGroupWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetSecurityGroupUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.SecurityGroup, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.SecurityGroup]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.SecurityGroup, error) {
			resp, err := api.network.GetSecurityGroupWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchSecurityGroupUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	if err := wref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.SecurityGroup]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetSecurityGroupWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdateSecurityGroupWithParams(ctx context.Context, group *schema.SecurityGroup, params *network.CreateOrUpdateSecurityGroupParams) (*schema.SecurityGroup, error) {
	if err := api.validateRegionalMetadata(group.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateSecurityGroupWithResponse(ctx, group.Metadata.Tenant, group.Metadata.Workspace, group.Metadata.Name, params, *group, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) CreateOrUpdateSecurityGroup(ctx context.Context, group *schema.SecurityGroup) (*schema.SecurityGroup, error) {
	return api.CreateOrUpdateSecurityGroupWithParams(ctx, group, nil)
}

func (api *NetworkV1Impl) DeleteSecurityGroupWithParams(ctx context.Context, group *schema.SecurityGroup, params *network.DeleteSecurityGroupParams) error {
	if err := api.validateRegionalMetadata(group.Metadata); err != nil {
		return err
	}

	resp, err := api.network.DeleteSecurityGroupWithResponse(ctx, group.Metadata.Tenant, group.Metadata.Workspace, group.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteSecurityGroup(ctx context.Context, group *schema.SecurityGroup) error {
	return api.DeleteSecurityGroupWithParams(ctx, group, nil)
}

// Nic

func (api *NetworkV1Impl) ListNicsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.Nic], error) {
	if err := wpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.Nic]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Nic, *string, error) {
			var params *network.ListNicsParams
			if options == nil {
				params = &network.ListNicsParams{
					Accept:    AcceptHeaderJson[network.ListNicsParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListNicsParams{
					Accept:    AcceptHeaderJson[network.ListNicsParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListNicsWithResponse(ctx, schema.TenantPathParam(wpath.Tenant), schema.WorkspacePathParam(wpath.Workspace), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListNics(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.Nic], error) {
	return api.ListNicsWithOptions(ctx, wpath, nil)
}

func (api *NetworkV1Impl) GetNic(ctx context.Context, wref WorkspaceReference) (*schema.Nic, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetNicWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetNicUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Nic, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Nic]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.Nic, error) {
			resp, err := api.network.GetNicWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchNicUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	if err := wref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Nic]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetNicWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdateNicWithParams(ctx context.Context, nic *schema.Nic, params *network.CreateOrUpdateNicParams) (*schema.Nic, error) {
	if err := api.validateRegionalMetadata(nic.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdateNicWithResponse(ctx, nic.Metadata.Tenant, nic.Metadata.Workspace, nic.Metadata.Name, params, *nic, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
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

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) DeleteNic(ctx context.Context, nic *schema.Nic) error {
	return api.DeleteNicWithParams(ctx, nic, nil)
}

// Public Ip

func (api *NetworkV1Impl) ListPublicIpsWithOptions(ctx context.Context, wpath WorkspacePath, options *ListOptions) (*Iterator[schema.PublicIp], error) {
	if err := wpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.PublicIp]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.PublicIp, *string, error) {
			var params *network.ListPublicIpsParams
			if options == nil {
				params = &network.ListPublicIpsParams{
					Accept:    AcceptHeaderJson[network.ListPublicIpsParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &network.ListPublicIpsParams{
					Accept:    AcceptHeaderJson[network.ListPublicIpsParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.network.ListPublicIpsWithResponse(ctx, schema.TenantPathParam(wpath.Tenant), schema.WorkspacePathParam(wpath.Workspace), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *NetworkV1Impl) ListPublicIps(ctx context.Context, wpath WorkspacePath) (*Iterator[schema.PublicIp], error) {
	return api.ListPublicIpsWithOptions(ctx, wpath, nil)
}

func (api *NetworkV1Impl) GetPublicIp(ctx context.Context, wref WorkspaceReference) (*schema.PublicIp, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.network.GetPublicIpWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *NetworkV1Impl) GetPublicIpUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.PublicIp, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.PublicIp]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.PublicIp, error) {
			resp, err := api.network.GetPublicIpWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *NetworkV1Impl) WatchPublicIpUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	if err := wref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.PublicIp]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.network.GetPublicIpWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *NetworkV1Impl) CreateOrUpdatePublicIpWithParams(ctx context.Context, ip *schema.PublicIp, params *network.CreateOrUpdatePublicIpParams) (*schema.PublicIp, error) {
	if err := api.validateRegionalMetadata(ip.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.network.CreateOrUpdatePublicIpWithResponse(ctx, ip.Metadata.Tenant, ip.Metadata.Workspace, ip.Metadata.Name, params, *ip, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
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

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
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
		return ErrNoMetadata
	}

	if metadata.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if metadata.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	return nil
}

func (api *NetworkV1Impl) validateNetworkMetadata(metadata *schema.RegionalNetworkResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetadata
	}

	if metadata.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if metadata.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	if metadata.Network == "" {
		return ErrNoMetadataNetwork
	}

	return nil
}
