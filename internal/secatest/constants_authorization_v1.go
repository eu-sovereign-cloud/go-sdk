package secatest

const (
	ListRolesResponseTemplateV1 = `
	{
		"items": [
			{
			"metadata": {
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
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		}
	}`

	GetRoleResponseTemplateV1 = `
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
	CreateOrUpdateRoleResponseTemplateV1 = `
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

	ListRolesAssignmentsResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"tenant": "{{.Tenant}}",
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
								"tenant-1"
							],
							"regions": [
								"region-1"
							],
							"workspaces": [
								"workspace-1"
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
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		}
	}`

	GetRoleAssignmentResponseTemplateV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
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
				"tenant-1"
				],
				"regions": [
				"region-1"
				],
				"workspaces": [
				"workspace-1"
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
	}`

	CreateOrUpdateRoleAssigmentResponseTemplateV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
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
				"tenant-1"
				],
				"regions": [
				"region-1"
				],
				"workspaces": [
				"workspace-1"
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
	}`
)
