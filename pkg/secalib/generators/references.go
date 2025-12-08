package generators

import (
	"fmt"
)

func GenerateSkuRef(name string) string {
	return fmt.Sprintf(skuRef, name)
}

func GenerateInstanceRef(instanceName string) string {
	return fmt.Sprintf(instanceRef, instanceName)
}

func GenerateBlockStorageRef(blockStorageName string) string {
	return fmt.Sprintf(blockStorageRef, blockStorageName)
}

func GenerateInternetGatewayRef(internetGatewayName string) string {
	return fmt.Sprintf(internetGatewayRef, internetGatewayName)
}

func GenerateNetworkRef(networkName string) string {
	return fmt.Sprintf(networkRef, networkName)
}

func GenerateRouteTableRef(routeTableName string) string {
	return fmt.Sprintf(routeTableRef, routeTableName)
}

func GenerateSubnetRef(subnetName string) string {
	return fmt.Sprintf(subnetRef, subnetName)
}

func GeneratePublicIpRef(publicIpName string) string {
	return fmt.Sprintf(publicIpRef, publicIpName)
}
