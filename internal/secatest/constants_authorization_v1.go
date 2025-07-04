package secatest

const (
	RolesResponseV1 = `
	{
		"items": [
			{
			"metadata": {
				"tenant": "{{.Tenant}}",
				"name": "{{.Name}}",
				"provider": "seca.compute/v1",
				"resource": "tenants/1/workspaces/ws-1/instances/my-server",
				"verb": "get"
			},
			"spec": {
				"permissions": [
				{
					"provider": "seca.storage/v1",
					"resources": [
					"images/*",
					"block-storages/*"
					],
					"verb": [
					"get",
					"list"
					]
				},
				{
					"provider": "seca.compute/v1",
					"resources": [
					"instances/*"
					],
					"verb": [
					"get",
					"list"
					]
				},
				{
					"provider": "seca.network/v1",
					"resources": [
					"networks/*",
					"subnets/*",
					"route-tables/*",
					"nics/*",
					"internet-gateways/*",
					"security-groups/*",
					"public-ips/*"
					],
					"verb": [
					"get",
					"list"
					]
				}
				]
			},
			"status": {
				"state": "active",
				"conditions": [
				{
					"state": "active",
					"lastTransitionAt": "2024-11-21T14:39:22Z"
				}
				]
			}
			}
		],
		"metadata": {
			"tenant": "{{.Tenant}}",
			"name": "{{.Name}}",
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		}
	}`

	CreateOrUpdateRoleResponseV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		},
		"spec": {
			"permissions": [
				{
					"provider": "seca.storage/v1",
					"resources": [
						"images/*",
						"block-storages/*"
					],
					"verb": [
						"get",
						"list"
					]
				},
				{
					"provider": "seca.compute/v1",
					"resources": [
						"instances/*"
					],
					"verb": [
						"get",
						"list"
					]
				},
				{
					"provider": "seca.network/v1",
					"resources": [
						"networks/*",
						"subnets/*",
						"route-tables/*",
						"nics/*",
						"internet-gateways/*",
						"security-groups/*",
						"public-ips/*"
					],
					"verb": [
						"get",
						"list"
					]
				}
			]
		},
		"status": {
			"state": "active",
			"conditions": [
				{
					"state": "active",
					"lastTransitionAt": "2024-11-21T14:39:22Z"
				}
			]
		}
	}`

	RolesAssignmentsResponseV1 = `
	{
		"items": [
			{
				"metadata": {
					"tenant": "{{.Tenant}}",
					"name": "{{.Name}}",
					"provider": "seca.compute/v1",
					"resource": "tenants/1/workspaces/ws-1/instances/my-server",
					"verb": "get"
				},
				"spec": {
					"subs": [
						"user1@example.com",
						"service-account-1"
					],
					"scopes": [
						{
							"tenants": [
								"{{.Tenant}}"
							],
							"regions": [
								"{{.Region}}"
							],
							"workspaces": [
								"{{.Workspace}}"
							]
						}
					],
					"roles": [
						"project-manager",
						"workspace-viewer"
					]
				},
				"status": {
					"state": "active",
					"conditions": [
						{
							"state": "active",
							"lastTransitionAt": "2024-11-21T14:39:22Z"
						}
					]
				}
			}
		],
		"metadata": {
			"tenant": "{{.Tenant}}",
			"name": "{{.Name}}",
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		}
	}`
)
