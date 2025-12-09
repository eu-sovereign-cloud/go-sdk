package generators

const (
	// Endpoint URLs
	urlProvidersPrefix = "/providers/"

	roleURL            = urlProvidersPrefix + "%s/" + roleResource
	roleAssignmentURL  = urlProvidersPrefix + "%s/" + roleAssignmentResource
	regionsURL         = urlProvidersPrefix + "%s/regions"
	regionURL          = urlProvidersPrefix + "%s/" + regionResource
	workspaceURL       = urlProvidersPrefix + "%s/" + workspaceResource
	instanceSkuURL     = urlProvidersPrefix + "%s/" + skuResource
	instanceURL        = urlProvidersPrefix + "%s/" + instanceResource
	storageSkuURL      = urlProvidersPrefix + "%s/" + skuResource
	blockStorageURL    = urlProvidersPrefix + "%s/" + blockStorageResource
	imageURL           = urlProvidersPrefix + "%s/" + imageResource
	networkSkuURL      = urlProvidersPrefix + "%s/" + skuResource
	networkURL         = urlProvidersPrefix + "%s/" + networkResource
	internetGatewayURL = urlProvidersPrefix + "%s/" + internetGatewayResource
	nicURL             = urlProvidersPrefix + "%s/" + nicResource
	publicIpURL        = urlProvidersPrefix + "%s/" + publicIpResource
	routeTableURL      = urlProvidersPrefix + "%s/" + routeTableResource
	subnetURL          = urlProvidersPrefix + "%s/" + subnetResource
	securityGroupURL   = urlProvidersPrefix + "%s/" + securityGroupResource

	// Resource URLs
	resourceTenantsPrefix    = "tenants/%s"
	resourceWorkspacesPrefix = resourceTenantsPrefix + "/workspaces/%s"

	regionResource          = "regions/%s"
	skuResource             = resourceTenantsPrefix + "/skus/%s"
	roleResource            = resourceTenantsPrefix + "/roles/%s"
	roleAssignmentResource  = resourceTenantsPrefix + "/role-assignments/%s"
	workspaceResource       = resourceTenantsPrefix + "/workspaces/%s"
	blockStorageResource    = resourceWorkspacesPrefix + "/block-storages/%s"
	imageResource           = resourceTenantsPrefix + "/images/%s"
	instanceResource        = resourceWorkspacesPrefix + "/instances/%s"
	networkResource         = resourceWorkspacesPrefix + "/networks/%s"
	internetGatewayResource = resourceWorkspacesPrefix + "/internet-gateways/%s"
	nicResource             = resourceWorkspacesPrefix + "/nics/%s"
	publicIpResource        = resourceWorkspacesPrefix + "/public-ips/%s"
	routeTableResource      = resourceWorkspacesPrefix + "/networks/%s/route-tables/%s"
	subnetResource          = resourceWorkspacesPrefix + "/networks/%s/subnets/%s"
	securityGroupResource   = resourceWorkspacesPrefix + "/security-groups/%s"

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
