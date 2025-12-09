package generators

import (
	"fmt"
)

func GenerateSkuResource(tenant, sku string) string {
	return fmt.Sprintf(skuResource, tenant, sku)
}

func GenerateRoleResource(tenant, role string) string {
	return fmt.Sprintf(roleResource, tenant, role)
}

func GenerateRoleAssignmentResource(tenant, roleAssignment string) string {
	return fmt.Sprintf(roleAssignmentResource, tenant, roleAssignment)
}

func GenerateRegionResource(region string) string {
	return fmt.Sprintf(regionResource, region)
}

func GenerateWorkspaceResource(tenant, workspace string) string {
	return fmt.Sprintf(workspaceResource, tenant, workspace)
}

func GenerateBlockStorageResource(tenant, workspace, blockStorage string) string {
	return fmt.Sprintf(blockStorageResource, tenant, workspace, blockStorage)
}

func GenerateImageResource(tenant, image string) string {
	return fmt.Sprintf(imageResource, tenant, image)
}

func GenerateInstanceResource(tenant, workspace, instance string) string {
	return fmt.Sprintf(instanceResource, tenant, workspace, instance)
}

func GenerateNetworkResource(tenant, workspace, network string) string {
	return fmt.Sprintf(networkResource, tenant, workspace, network)
}

func GenerateInternetGatewayResource(tenant, workspace, internetGateway string) string {
	return fmt.Sprintf(internetGatewayResource, tenant, workspace, internetGateway)
}

func GenerateNicResource(tenant, workspace, nic string) string {
	return fmt.Sprintf(nicResource, tenant, workspace, nic)
}

func GeneratePublicIpResource(tenant, workspace, publicIp string) string {
	return fmt.Sprintf(publicIpResource, tenant, workspace, publicIp)
}

func GenerateRouteTableResource(tenant, workspace, network, routeTable string) string {
	return fmt.Sprintf(routeTableResource, tenant, workspace, network, routeTable)
}

func GenerateSubnetResource(tenant, workspace, network, subnet string) string {
	return fmt.Sprintf(subnetResource, tenant, workspace, network, subnet)
}

func GenerateSecurityGroupResource(tenant, workspace, securityGroup string) string {
	return fmt.Sprintf(securityGroupResource, tenant, workspace, securityGroup)
}
