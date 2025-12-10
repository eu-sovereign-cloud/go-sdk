package generators

const (
	// Endpoint URLs
	urlProvidersPrefix = "/providers/"

	roleURL                = urlProvidersPrefix + "%s/" + roleResource
	roleListURL            = urlProvidersPrefix + "%s/" + roleListResource
	roleAssignmentURL      = urlProvidersPrefix + "%s/" + roleAssignmentResource
	roleAssignmentListURL  = urlProvidersPrefix + "%s/" + roleAssignmentListResource
	regionURL              = urlProvidersPrefix + "%s/" + regionResource
	regionListURL          = urlProvidersPrefix + "%s/" + regionListResource
	workspaceURL           = urlProvidersPrefix + "%s/" + workspaceResource
	workspaceListURL       = urlProvidersPrefix + "%s/" + workspaceListResource
	instanceSkuURL         = urlProvidersPrefix + "%s/" + skuResource
	instanceSkuListURL     = urlProvidersPrefix + "%s/" + skuListResource
	instanceURL            = urlProvidersPrefix + "%s/" + instanceResource
	instanceListURL        = urlProvidersPrefix + "%s/" + instanceListResource
	storageSkuURL          = urlProvidersPrefix + "%s/" + skuResource
	storageSkuListURL      = urlProvidersPrefix + "%s/" + skuListResource
	blockStorageURL        = urlProvidersPrefix + "%s/" + blockStorageResource
	blockStorageListURL    = urlProvidersPrefix + "%s/" + blockStorageListResource
	imageURL               = urlProvidersPrefix + "%s/" + imageResource
	imageListURL           = urlProvidersPrefix + "%s/" + imageListResource
	networkSkuURL          = urlProvidersPrefix + "%s/" + skuResource
	networkSkuListURL      = urlProvidersPrefix + "%s/" + skuListResource
	networkURL             = urlProvidersPrefix + "%s/" + networkResource
	networkListURL         = urlProvidersPrefix + "%s/" + networkListResource
	internetGatewayURL     = urlProvidersPrefix + "%s/" + internetGatewayResource
	internetGatewayListURL = urlProvidersPrefix + "%s/" + internetGatewayListResource
	nicURL                 = urlProvidersPrefix + "%s/" + nicResource
	nicListURL             = urlProvidersPrefix + "%s/" + nicListResource
	publicIpURL            = urlProvidersPrefix + "%s/" + publicIpResource
	publicIpListURL        = urlProvidersPrefix + "%s/" + publicIpListResource
	routeTableURL          = urlProvidersPrefix + "%s/" + routeTableResource
	routeTableListURL      = urlProvidersPrefix + "%s/" + routeTableListResource
	subnetURL              = urlProvidersPrefix + "%s/" + subnetResource
	subnetListURL          = urlProvidersPrefix + "%s/" + subnetListResource
	securityGroupURL       = urlProvidersPrefix + "%s/" + securityGroupResource
	securityGroupListURL   = urlProvidersPrefix + "%s/" + securityGroupListResource

	// Resource URLs
	resourceTenantsPrefix    = "tenants/%s"
	resourceWorkspacesPrefix = resourceTenantsPrefix + "/workspaces/%s"

	regionResource              = "regions/%s"
	regionListResource          = "regions"
	skuResource                 = resourceTenantsPrefix + "/skus/%s"
	skuListResource             = resourceTenantsPrefix + "/skus"
	roleResource                = resourceTenantsPrefix + "/roles/%s"
	roleListResource            = resourceTenantsPrefix + "/roles"
	roleAssignmentResource      = resourceTenantsPrefix + "/role-assignments/%s"
	roleAssignmentListResource  = resourceTenantsPrefix + "/role-assignments"
	workspaceResource           = resourceTenantsPrefix + "/workspaces/%s"
	workspaceListResource       = resourceTenantsPrefix + "/workspaces"
	blockStorageResource        = resourceWorkspacesPrefix + "/block-storages/%s"
	blockStorageListResource    = resourceWorkspacesPrefix + "/block-storages"
	imageResource               = resourceTenantsPrefix + "/images/%s"
	imageListResource           = resourceTenantsPrefix + "/images"
	instanceResource            = resourceWorkspacesPrefix + "/instances/%s"
	instanceListResource        = resourceWorkspacesPrefix + "/instances"
	networkResource             = resourceWorkspacesPrefix + "/networks/%s"
	networkListResource         = resourceWorkspacesPrefix + "/networks"
	internetGatewayResource     = resourceWorkspacesPrefix + "/internet-gateways/%s"
	internetGatewayListResource = resourceWorkspacesPrefix + "/internet-gateways"
	nicResource                 = resourceWorkspacesPrefix + "/nics/%s"
	nicListResource             = resourceWorkspacesPrefix + "/nics"
	publicIpResource            = resourceWorkspacesPrefix + "/public-ips/%s"
	publicIpListResource        = resourceWorkspacesPrefix + "/public-ips"
	routeTableResource          = resourceWorkspacesPrefix + "/networks/%s/route-tables/%s"
	routeTableListResource      = resourceWorkspacesPrefix + "/networks/%s/route-tables"
	subnetResource              = resourceWorkspacesPrefix + "/networks/%s/subnets/%s"
	subnetListResource          = resourceWorkspacesPrefix + "/networks/%s/subnets"
	securityGroupResource       = resourceWorkspacesPrefix + "/security-groups/%s"
	securityGroupListResource   = resourceWorkspacesPrefix + "/security-groups"

	// References
	skuRef             = "skus/%s"
	instanceRef        = "instances/%s"
	blockStorageRef    = "block-storages/%s"
	internetGatewayRef = "internet-gateways/%s"
	networkRef         = "networks/%s"
	routeTableRef      = "route-tables/%s"
	subnetRef          = "subnets/%s"
	publicIpRef        = "public-ips/%s"

	// Generators
	maxBlockStorageSize = 1000000 // GB
)
