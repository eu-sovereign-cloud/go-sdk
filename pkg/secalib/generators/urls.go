package generators

import (
	"fmt"
)

func GenerateRoleURL(provider, tenant, role string) string {
	return fmt.Sprintf(roleURL, provider, tenant, role)
}

func GenerateRoleAssignmentURL(provider, tenant string, roleAssignment string) string {
	return fmt.Sprintf(roleAssignmentURL, provider, tenant, roleAssignment)
}

func GenerateRegionsURL(provider string) string {
	return fmt.Sprintf(regionsURL, provider)
}

func GenerateRegionURL(provider, region string) string {
	return fmt.Sprintf(regionURL, provider, region)
}

func GenerateRegionProviderUrl(provider string) string {
	return fmt.Sprintf("{{request.scheme}}://{{request.host}}:{{request.port}}%s%s", urlProvidersPrefix, provider)
}

func GenerateWorkspaceURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(workspaceURL, provider, tenant, workspace)
}

func GenerateStorageSkuURL(provider, tenant, sku string) string {
	return fmt.Sprintf(storageSkuURL, provider, tenant, sku)
}

func GenerateBlockStorageURL(provider, tenant, workspace, blockStorage string) string {
	return fmt.Sprintf(blockStorageURL, provider, tenant, workspace, blockStorage)
}

func GenerateImageURL(provider, tenant, image string) string {
	return fmt.Sprintf(imageURL, provider, tenant, image)
}

func GenerateInstanceSkuURL(provider, tenant, sku string) string {
	return fmt.Sprintf(instanceSkuURL, provider, tenant, sku)
}

func GenerateInstanceURL(provider, tenant, workspace, instance string) string {
	return fmt.Sprintf(instanceURL, provider, tenant, workspace, instance)
}

func GenerateNetworkURL(provider, tenant, workspace, network string) string {
	return fmt.Sprintf(networkURL, provider, tenant, workspace, network)
}

func GenerateNetworkSkuURL(provider, tenant, sku string) string {
	return fmt.Sprintf(networkSkuURL, provider, tenant, sku)
}

func GenerateInternetGatewayURL(provider, tenant, workspace, internetGateway string) string {
	return fmt.Sprintf(internetGatewayURL, provider, tenant, workspace, internetGateway)
}

func GenerateNicURL(provider, tenant, workspace, nic string) string {
	return fmt.Sprintf(nicURL, provider, tenant, workspace, nic)
}

func GeneratePublicIpURL(provider, tenant, workspace, publicIp string) string {
	return fmt.Sprintf(publicIpURL, provider, tenant, workspace, publicIp)
}

func GenerateRouteTableURL(provider, tenant, workspace, network, routeTable string) string {
	return fmt.Sprintf(routeTableURL, provider, tenant, workspace, network, routeTable)
}

func GenerateSubnetURL(provider, tenant, workspace, network, subnet string) string {
	return fmt.Sprintf(subnetURL, provider, tenant, workspace, network, subnet)
}

func GenerateSecurityGroupURL(provider, tenant, workspace, securityGroup string) string {
	return fmt.Sprintf(securityGroupURL, provider, tenant, workspace, securityGroup)
}
