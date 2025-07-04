package secatest

const (
	// Sku
	networkSkusResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"bandwidth": 1000,
					"packets": 10000
				}
			}
    	],
    	"metadata": {
    	}
	}`
	networkSkuResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"bandwidth": 1000,
			"packets": 10000
		}
	}`

	// Network
	networksResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"skuRef": "skus/sku-1",
        			"routeTableRef": "route-tables/route-table-1",
        			"cidr": {
          				"ipv4": "0.0.0.0/16",
          				"ipv6": "::/56"
        			}
				}
			}
    	],
    	"metadata": {
    	}
	}`
	networkResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"skuRef": "skus/sku-1",
			"routeTableRef": "route-tables/route-table-1",
			"cidr": {
				"ipv4": "0.0.0.0/16",
				"ipv6": "::/56"
			}
		}	
	}`

	// Subnet
	subnetsResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"cidr": {
						"ipv4": "0.0.0.0/24",
						"ipv6": "::/64"
					},
					"zone": "a",
					"skuRef": "skus/sku-1",
					"routeTableRef": "route-tables/route-table-1"
				}
			}
    	],
    	"metadata": {
    	}
	}`
	subnetResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"cidr": {
				"ipv4": "0.0.0.0/24",
				"ipv6": "::/64"
			},
			"zone": "a",
			"skuRef": "skus/sku-1",
			"routeTableRef": "route-tables/route-table-1"
		}
	}`

	// Route Table
	routeTablesResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"localRef": "networks/network-1",
					"routes": [
						{
							"destinationCidrBlock": "0.0.0.0/0",
							"targetRef": "internet-gateways/internet-gateway-1"
						},
						{
							"destinationCidrBlock": "::/0",
							"targetRef": "internet-gateways/internet-gateway-1"
						}
					]
				}
			}
    	],
    	"metadata": {
    	}
	}`
	routeTableResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"localRef": "networks/network-1",
			"routes": [
				{
					"destinationCidrBlock": "0.0.0.0/0",
					"targetRef": "internet-gateways/internet-gateway-1"
				},
				{
					"destinationCidrBlock": "::/0",
					"targetRef": "internet-gateways/internet-gateway-1"
				}
			]
		}
	}`

	// Internet Gateway
	internetGatewaysResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"egressOnly": false
				}
			}
    	],
    	"metadata": {
    	}
	}`
	internetGatewayResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"egressOnly": false
		}
	}`

	// Security Group
	securityGroupsResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"rules": [
          				{
							"direction": "ingress",
							"protocol": "tcp",
							"ports": {
								"list": [80, 443]
							}
          				}
        			]
				}
			}
    	],
    	"metadata": {
    	}
	}`
	securityGroupResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"rules": [
				{
					"direction": "ingress",
					"protocol": "tcp",
					"ports": {
						"list": [80, 443]
					}
				}
			]
		}
	}`

	// Nic
	nicsResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"addresses": [
						"10.0.0.1",
						"0.0.0.0",
						"::"
					],
					"skuRef": "skus/sku-1",
					"subnetRef": "seca.networks/subnet-1"
				}
			}
    	],
    	"metadata": {
    	}
	}`
	nicResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"addresses": [
				"10.0.0.1",
				"0.0.0.0",
				"::"
			],
			"skuRef": "skus/sku-1",
			"subnetRef": "seca.networks/subnet-1"
		}
	}`

	// Public Ip
	publicIpsResponseV1 = `
	{
		"items": [
        	{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"name": "{{.Name}}"
				},
				"spec": {
					"version": "IPv4",
					"address": "10.0.0.1"
				}
			}
    	],
    	"metadata": {
    	}
	}`
	publicIpResponseV1 = `
	{
		"metadata": {
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"name": "{{.Name}}"
		},
		"spec": {
			"version": "IPv4",
			"address": "10.0.0.1"
		}
	}`
)
