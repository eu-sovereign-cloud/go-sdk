package generators

import (
	"math/rand"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
)

// Random

func GenerateBlockStorageSize() int {
	return rand.Intn(maxBlockStorageSize)
}

// Network

func GenerateSubnetCidr(networkCidr string, size int, netNum int) (string, error) {
	_, network, err := net.ParseCIDR(networkCidr)
	if err != nil {
		return "", err
	}

	subnet, err := cidr.Subnet(network, size, netNum)
	if err != nil {
		return "", err
	}

	return subnet.String(), nil
}

func GenerateNicAddress(subnetCidr string, hostNum int) (string, error) {
	_, network, err := net.ParseCIDR(subnetCidr)
	if err != nil {
		return "", err
	}

	ip, err := cidr.Host(network, hostNum)
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}

func GeneratePublicIp(publicIpRange string, hostNum int) (string, error) {
	_, network, err := net.ParseCIDR(publicIpRange)
	if err != nil {
		return "", err
	}

	ip, err := cidr.Host(network, hostNum)
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}
