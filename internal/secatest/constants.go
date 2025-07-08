package secatest

const (
	// Http Headers
	ContentTypeHeader = "Content-Type"
	ContentTypeJSON   = "application/json"

	// Providers
	ProviderRegionName     = "seca.region/v1"
	ProviderRegionEndpoint = "/providers/seca.regions"

	ProviderAuthorizationName     = "seca.authorization"
	ProviderAuthorizationEndpoint = "/providers/seca.authorization"

	ProviderWorkspaceName     = "seca.workspace"
	ProviderWorkspaceEndpoint = "/providers/seca.workspace"

	ProviderNetworkName     = "seca.network"
	ProviderNetworkEndpoint = "/providers/seca.network"

	ProviderComputeName     = "seca.compute"
	ProviderComputeEndpoint = "/providers/seca.compute"

	ProviderStorageName     = "seca.storage"
	ProviderStorageEndpoint = "/providers/seca.storage"

	ProviderVersion1 = "v1"

	// Test Data

	/// Metadata
	Tenant1Name    = "tenant-1"
	Workspace1Name = "woskpace-1"
	Region1Name    = "region-1"
	ZoneA          = "a"

	/// Network
	NetworkSku1Name = "sku-1"
	NetworkSku1Ref  = "skus/sku-1"

	Network1Name = "network-1"
	Network1Ref  = "networks/network-1"

	Subnet1Name = "subnet-1"

	RouteTable1Name = "route-table-1"
	RouteTable1Ref  = "route-tables/route-table-1"

	InternetGateway1Name = "internet-gateway-1"

	SecurityGroup1Name     = "security-group-1"
	SecurityGroup1PortFrom = 80
	SecurityGroup1PortTo   = 80

	Nic1Name = "nic-1"

	PublicIp1Name = "public-ip-1"

	CidrIpv4 = "0.0.0.0/16"

	/// Compute
	InstanceSku1Name         = "sku-1"
	InstanceSku1Provider     = "seca"
	InstanceSku1Tier         = "D2XS"
	InstanceSku1VCPU         = 16
	InstanceSku1Ram          = 32
	InstanceSku1Architecture = "amd64"

	Instance1Ref  = "instances/instance-1"
	Instance1Name = "instance-1"

	/// Authorization
	AuthorizationRole1Name = "role-1"

	AuthorizationRoleAssignment1Name    = "role-assignment-1"
	AuthorizationRoleAssignment1Subject = "sub@secapi.com"
	AuthorizationPermissionVerb         = "get"

	/// Storage
	StorageSku1Ref           = "storage/skus-1"
	StorageSku1Name          = "sku-1"
	StorageSku1Provider      = "seca"
	StorageSku1Tier          = "DXS"
	StorageSku1Iops          = 100
	StorageSku1MinVolumeSize = 50
	StorageSku1Type          = "remote-durable"

	BlockStorage1Ref = "storage/block-storage-1"
	Storage1Name     = "storage-1"
	Image1Name       = "image-1"

	/// Status
	StatusStateActive    = "active"
	StatusStateCreating  = "creating"
	StatusStateDeleting  = "deleting"
	StatusStateError     = "error"
	StatusStatePending   = "pending"
	StatusStateSuspended = "suspended"
	StatusStateUpdating  = "updating"
)
