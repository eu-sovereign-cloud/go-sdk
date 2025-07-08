package secatest

const (
	// Network Sku
	networkSkusResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"bandwidth": {{.Bandwidth}}
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	networkSkuResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"bandwidth": {{.Bandwidth}}
		},
		status": {
			"state": "{{.Status}}"
		}
	}`

	// Network
	networksResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
        			"routeTableRef": "{{.RouteTableRef}}"
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	networkResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"routeTableRef": "{{.RouteTableRef}}"
		},
		status": {
			"state": "{{.Status}}"
		}
	}`

	// Subnet
	subnetsResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"skuRef": "{{.SkuRef}}"
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	subnetResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"skuRef": "{{.SkuRef}}"
		},
		status": {
			"state": "{{.Status}}"
		}
	}`

	// Route Table
	routeTablesResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"localRef": "{{.LocalRef}}"
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	routeTableResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"localRef": "{{.LocalRef}}"
		},
		status": {
			"state": "{{.Status}}"
		}
	}`

	// Internet Gateway
	internetGatewaysResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"egressOnly": {{.EgressOnly}}
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	internetGatewayResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"egressOnly": {{.EgressOnly}}
		},
		status": {
			"state": "{{.Status}}"
		}
	}`

	// Security Group
	securityGroupsResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"rules": [
          				{
							"direction": "{{.RuleDirection}}"
          				}
        			]
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	securityGroupResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"rules": [
				{
					"direction": "{{.RuleDirection}}"
				}
			]
		},
		status": {
			"state": "{{.Status}}"
		}
	}`

	// Nic
	nicsResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"subnetRef": "{{.SubnetRef}}"
				},
				status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	nicResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"subnetRef": "{{.SubnetRef}}"
		},
		status": {
			"state": "{{.Status}}"
		}
	}`

	// Public Ip
	publicIpsResponseTemplateV1 = `
	{
		"items": [
        	{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"address": "{{.Address}}"
				},
				status": {
					"state": "{{.Status}}"
				}
			}
    	]
	}`
	publicIpResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"address": "{{.Address}}"
		},
		status": {
			"state": "{{.Status}}"
		}
	}`
)
