package generators

import (
	"fmt"
)

func GenerateRoleURL(provider, tenant, role string) string {
	return fmt.Sprintf(roleURL, provider, tenant, role)
}

func GenerateRoleListURL(provider, tenant string) string {
	return fmt.Sprintf(roleListURL, provider, tenant)
}

func GenerateRoleAssignmentURL(provider, tenant string, roleAssignment string) string {
	return fmt.Sprintf(roleAssignmentURL, provider, tenant, roleAssignment)
}

func GenerateRoleAssignmentListURL(provider, tenant string) string {
	return fmt.Sprintf(roleAssignmentListURL, provider, tenant)
}

func GenerateRegionListURL(provider string) string {
	return fmt.Sprintf(regionListURL, provider)
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

func GenerateWorkspaceListURL(provider, tenant string) string {
	return fmt.Sprintf(workspaceListURL, provider, tenant)
}

func GenerateStorageSkuURL(provider, tenant, sku string) string {
	return fmt.Sprintf(storageSkuURL, provider, tenant, sku)
}

func GenerateStorageSkuListURL(provider, tenant string) string {
	return fmt.Sprintf(storageSkuListURL, provider, tenant)
}

func GenerateBlockStorageURL(provider, tenant, workspace, blockStorage string) string {
	return fmt.Sprintf(blockStorageURL, provider, tenant, workspace, blockStorage)
}

func GenerateBlockStorageListURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(blockStorageListURL, provider, tenant, workspace)
}

func GenerateImageURL(provider, tenant, image string) string {
	return fmt.Sprintf(imageURL, provider, tenant, image)
}

func GenerateImageListURL(provider, tenant string) string {
	return fmt.Sprintf(imageListURL, provider, tenant)
}

func GenerateInstanceSkuURL(provider, tenant, sku string) string {
	return fmt.Sprintf(instanceSkuURL, provider, tenant, sku)
}

func GenerateInstanceSkuListURL(provider, tenant string) string {
	return fmt.Sprintf(instanceSkuListURL, provider, tenant)
}

func GenerateInstanceURL(provider, tenant, workspace, instance string) string {
	return fmt.Sprintf(instanceURL, provider, tenant, workspace, instance)
}

func GenerateInstanceListURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(instanceListURL, provider, tenant, workspace)
}

func GenerateNetworkURL(provider, tenant, workspace, network string) string {
	return fmt.Sprintf(networkURL, provider, tenant, workspace, network)
}

func GenerateNetworkListURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(networkListURL, provider, tenant, workspace)
}

func GenerateNetworkSkuURL(provider, tenant, sku string) string {
	return fmt.Sprintf(networkSkuURL, provider, tenant, sku)
}

func GenerateNetworkSkuListURL(provider, tenant string) string {
	return fmt.Sprintf(networkSkuListURL, provider, tenant)
}

func GenerateInternetGatewayURL(provider, tenant, workspace, internetGateway string) string {
	return fmt.Sprintf(internetGatewayURL, provider, tenant, workspace, internetGateway)
}

func GenerateInternetGatewayListURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(internetGatewayListURL, provider, tenant, workspace)
}

func GenerateNicURL(provider, tenant, workspace, nic string) string {
	return fmt.Sprintf(nicURL, provider, tenant, workspace, nic)
}

func GenerateNicListURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(nicListURL, provider, tenant, workspace)
}

func GeneratePublicIpURL(provider, tenant, workspace, publicIp string) string {
	return fmt.Sprintf(publicIpURL, provider, tenant, workspace, publicIp)
}

func GeneratePublicIpListURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(publicIpListURL, provider, tenant, workspace)
}

func GenerateRouteTableURL(provider, tenant, workspace, network, routeTable string) string {
	return fmt.Sprintf(routeTableURL, provider, tenant, workspace, network, routeTable)
}

func GenerateRouteTableListURL(provider, tenant, workspace, network string) string {
	return fmt.Sprintf(routeTableListURL, provider, tenant, workspace, network)
}

func GenerateSubnetURL(provider, tenant, workspace, network, subnet string) string {
	return fmt.Sprintf(subnetURL, provider, tenant, workspace, network, subnet)
}

func GenerateSubnetListURL(provider, tenant, workspace, network string) string {
	return fmt.Sprintf(subnetListURL, provider, tenant, workspace, network)
}

func GenerateSecurityGroupURL(provider, tenant, workspace, securityGroup string) string {
	return fmt.Sprintf(securityGroupURL, provider, tenant, workspace, securityGroup)
}

func GenerateSecurityGroupListURL(provider, tenant, workspace string) string {
	return fmt.Sprintf(securityGroupListURL, provider, tenant, workspace)
}
