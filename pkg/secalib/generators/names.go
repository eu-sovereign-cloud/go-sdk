package generators

import (
	"fmt"
	"math"
	"math/rand"
)

func GenerateRoleName() string {
	return fmt.Sprintf("role-%d", rand.Intn(math.MaxInt32))
}

func GenerateRoleAssignmentName() string {
	return fmt.Sprintf("role-assignment-%d", rand.Intn(math.MaxInt32))
}

func GenerateRegionName() string {
	return fmt.Sprintf("region-%d", rand.Intn(math.MaxInt32))
}

func GenerateWorkspaceName() string {
	return fmt.Sprintf("workspace-%d", rand.Intn(math.MaxInt32))
}

func GenerateBlockStorageName() string {
	return fmt.Sprintf("disk-%d", rand.Intn(math.MaxInt32))
}

func GenerateImageName() string {
	return fmt.Sprintf("image-%d", rand.Intn(math.MaxInt32))
}

func GenerateInstanceName() string {
	return fmt.Sprintf("instance-%d", rand.Intn(math.MaxInt32))
}

func GenerateNetworkName() string {
	return fmt.Sprintf("network-%d", rand.Intn(math.MaxInt32))
}

func GenerateInternetGatewayName() string {
	return fmt.Sprintf("internet-gateway-%d", rand.Intn(math.MaxInt32))
}

func GenerateRouteTableName() string {
	return fmt.Sprintf("route-table-%d", rand.Intn(math.MaxInt32))
}

func GenerateSubnetName() string {
	return fmt.Sprintf("subnet-%d", rand.Intn(math.MaxInt32))
}

func GeneratePublicIpName() string {
	return fmt.Sprintf("public-ip-%d", rand.Intn(math.MaxInt32))
}

func GenerateNicName() string {
	return fmt.Sprintf("nic-%d", rand.Intn(math.MaxInt32))
}

func GenerateSecurityGroupName() string {
	return fmt.Sprintf("security-group-%d", rand.Intn(math.MaxInt32))
}
