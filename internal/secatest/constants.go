package secatest

const (
	// Http Headers
	ContentTypeHeader = "Content-Type"
	ContentTypeJSON   = "application/json"

	// Providers
	ProviderRegionName     = "seca.region/v1"
	ProviderRegionEndpoint = "/providers/seca.regions"

	ProviderWorkspaceName     = "seca.workspace"
	ProviderWorkspaceEndpoint = "/providers/seca.workspace"

	ProviderNetworkName     = "seca.network"
	ProviderNetworkEndpoint = "/providers/seca.network"

	ProviderComputeName     = "seca.compute"
	ProviderComputeEndpoint = "/providers/seca.compute"

	ProviderStorageName     = "seca.storage"
	ProviderStorageEndpoint = "/providers/seca.storage"

	ProviderAuthorizationName     = "seca.authorization"
	ProviderAuthorizationEndpoint = "/providers/seca.authorization"

	// Test Data

	// Metadata
	Tenant1Name = "tenant-1"

	Workspace1Name = "woskpace-1"

	RegionName  = "eu-central-1"
	Region1Name = "region-1"

	ZoneA = "a"

	// Network
	NetworkSku1Name = "sku-1"
	NetworkSku1Ref  = "skus/sku-1"

	Network1Name = "network-1"
	Network1Ref  = "networks/network-1"

	Subnet1Name = "subnet-1"

	RouteTable1Name = "route-table-1"
	RouteTable1Ref  = "route-tables/route-table-1"

	InternetGateway1Name = "internet-gateway-1"

	SecurityGroup1Name = "security-group-1"

	Nic1Name = "nic-1"

	PublicIp1Name = "public-ip-1"

	CidrIpv4 = "0.0.0.0/16"

	// Compute
	Instance1Ref  = "instances/instance-1"
	Instance1Name = "instance-1"

	// Authorization
	AuthorizationName               = "authorization-1"
	AuthorizationRoleName           = "role-1"
	AuthorizationRoleAssignmentName = "role-assignment-1"
)
