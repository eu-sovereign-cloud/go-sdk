package secatest

const (
	// Response Templates

	ListRegionsResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"apiVersion": "v1",
					"kind": "region",
					"name": "{{.Name}}"
				},
				"spec": {
					"availableZones": [ "A", "B" ],
					"providers": [
						{{- range $i, $p := .Providers }}
						{{if $i}},{{end}}
						{
							"name": "{{$p.Name}}",
							"url": "{{$p.URL}}"
						}
						{{- end}}
					]
				},
				"status": {
					"conditions": [ { "status": "Ready" } ]
				}
			}
		],
		"metadata": {
			"skipToken": null
		}
	}`

	GetRegionResponseTemplateV1 = `
	{
		"metadata": {
			"apiVersion": "v1",
			"kind": "region",
			"name": "{{.Name}}"
		},
		"spec": {
			"availableZones": [ "A", "B" ],
			"providers": [
				{{- range $i, $p := .Providers }}
				{{if $i}},{{end}}
				{
					"name": "{{$p.Name}}",
					"url": "{{$p.URL}}"
				}
				{{- end}}
			]
		},
		"status": {
			"conditions": [ { "status": "Ready" } ]
		}
	}`

	ListWorkspaceResponseTemplateV1 = `
	{
		"items": [
			{
				"apiVersion": "v1",
				"kind": "workspace",
				"metadata": {
					"name": "{{.Name}}",
					"tenant": "{{.Tenant}}"
				},
				"spec": {},
				"status": {
					"state": "{{.State}}",
					"conditions": [
						{
							"status": "Ready"
						}
					]
				},
				"resourcesCount": 1
			}
		],
		"metadata": {
			"skipToken": null
		}
	}`

	GetWorkspaceResponseTemplateV1 = `
	{
		"apiVersion": "v1",
		"kind": "workspace",
		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}"
		},
		"spec": {},
		"status": {
			"state": "active",
			"conditions": [ {
					"status": "active",
					"lastTransitionAt": "2025-06-23T14:15:22Z"
			 }]
		}
	}`

	CreateOrUpdateWorkspaceResponseTemplateV1 = `
	{
		"apiVersion": "v1",
		"kind": "workspace",
		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}"
		},
		"spec": {},
		"status": {
			"state": "active",
			"conditions": [ {
					"status": "active",
					"lastTransitionAt": "2025-06-23T14:15:22Z"
			 }],
				"resourcesCount": 1
		}
	}`

	CreateOrUpdateInstaceResponseTemplateV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"zone": "a",
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		},
		"spec": {
			"skuRef": "skus/seca.s",
			"zone": "a",
			"bootVolume": {
				"deviceRef": {
					"provider": "seca.storage/v1",
					"resource": "block-storages/block-123"
				},
				"type": "virtio"
			}
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

	GetInstanceResponseTemplateV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",		
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		},
		"spec": {
			"skuRef": "skus/seca.s",
			"zone": "a",
			"bootVolume": {
			"deviceRef": {
				"provider": "seca.storage/v1",
				"resource": "block-storages/block-123"
			},
			"type": "virtio"
			}
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

	GetInstanceSkuResponseTemplateV1 = `
	{

		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"vCPU": {{.VCPU}},
			"ram": {{.Ram}}
		}
	}`

	ListInstancesResponseTemplateV1 = `
	{
		"items": [
			{
				"extensions": {},
				"metadata": {
					"name": "{{.Name}}",
					"tenant": "{{.Tenant}}",			
					"provider": "seca.compute/v1",
					"resource": "tenants/1/workspaces/ws-1/instances/my-server",
					"verb": "get"
				},
				"spec": {
					"skuRef": "skus/seca.s",
					"zone": "a",
					"bootVolume": {
					"deviceRef": {
						"provider": "seca.storage/v1",
						"resource": "block-storages/block-123"
					},
					"type": "virtio"
					}
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

	ListInstancesSkusResponseTemplateV1 = `{
			"items": [
				{{- range $i, $d := .Skus }}
						{{if $i}},{{end}}
						{
							"labels": {
								"provider": "{{$d.Provider}}",
								"tier": "{{$d.Tier}}",
								"architecture": "{{$d.Architecture}}"
							},
							"metadata": {
								"name": "{{$d.Provider}}.{{$d.Tier}}"
							},
							"spec": {
								"ram": {{$d.Ram}},
								"vCPU": {{$d.VCPU}}
							}
						}
						{{- end}}
			],
			"metadata": {
				"provider": "seca.compute/v1",
				"resource": "tenants/1/skus",
				"verb": "list"
			}
	}`

	ListStorageSkusResponseTemplateV1 = `{
	"items": [
			{{- range $i, $d := .Skus }}
			{{- if $i}},{{ end }}
			{
				"labels": {
					"provider": "{{$d.Provider}}",
					"tier": "{{$d.Tier}}"
				},
				"metadata": {
					"name": "{{$d.Provider}}.{{$d.Tier}}"
				},
				"spec": {
					"iops": {{$d.Iops}},
					"minVolumeSize": {{$d.MinVolumeSize}},
					"type": "{{$d.Type}}"
				}
			}
			{{- end }}
		],
		"metadata": {
			"provider": "seca.compute/v1",
			"resource": "tenants/1/skus",
			"verb": "list"
		}
	}`

	GetStorageSkuResponseTemplateV1 = `
	{
		"extensions": {},
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"iops": 100,
			"type": "remote-durable",
			"minVolumeSize": 50
		}
	}`

	ListBlockStorageResponseTemplateV1 = `
	{
		"items": [
			{

			"metadata": {
				"provider": "seca.compute/v1",
				"resource": "tenants/1/workspaces/ws-1/instances/my-server",
				"verb": "get"
			},
			"spec": {
				"skuRef": "resources/resource-a1b2c3",
				"sizeGB": 10,
				"sourceImageRef": "resources/resource-a1b2c3"
			},
			"status": {
				"state": "active",
				"conditions": [
				{
					"state": "active",
					"lastTransitionAt": "2024-11-21T14:39:22Z"
				}
				],
				"sizeGB": 10,
				"attachedTo": "resources/resource-a1b2c3"
			}
			}
		],
		"metadata": {
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		}
	}`

	GetBlockStorageResponseTemplateV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		},
		"spec": {
			"skuRef": "resources/resource-a1b2c3",
			"sizeGB": 10,
			"sourceImageRef": "resources/resource-a1b2c3"
		},
		"status": {
			"state": "active",
			"conditions": [
			{
				"state": "active",
				"lastTransitionAt": "2024-11-21T14:39:22Z"
			}
			],
			"sizeGB": 10,
			"attachedTo": "resources/resource-a1b2c3"
		}
	}`

	CreateOrUpdateBlockStorageResponseTemplateV1 = `
	{

		"metadata": {
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		},
		"spec": {
			"skuRef": "resources/resource-a1b2c3",
			"sizeGB": 10,
			"sourceImageRef": "resources/resource-a1b2c3"
		},
		"status": {
			"state": "active",
			"conditions": [
			{
				"state": "active",
				"lastTransitionAt": "2024-11-21T14:39:22Z"
			}
			],
			"sizeGB": 10,
			"attachedTo": "resources/resource-a1b2c3"
		}
	}`

	ListImageResponseTemplateV1 = `
	{
			"items": [
				{
					"labels": {
						"os": "linux",
						"version": "13",
						"base": "debian"
					},
					"annotations": {
						"name": "Debian Container",
						"description": "The image contains the debian image base including\npreinstalled software for the use of linux containers.\n",
						"release": "2025-01-01T00:00:00Z",
						"eol": "2026-01-01T00:00:00Z",
						"recommendedCpu": "2",
						"recommendedMemory": "2",
						"recommendedNics": "2",
						"recommendedStorageSize": "100"
					},
					"spec": {
						"blockStorageRef": "block-storages/temp-a9b2bc3d81",
						"cpuArchitecture": "amd64",
						"boot": "UEFI",
						"initializer": "cloudinit-22"
					}
				}
			],
			"metadata": {
				"provider": "seca.compute/v1",
				"resource": "tenants/1/workspaces/ws-1/instances/my-server",
				"verb": "get",
				"skipToken": "false"
			}
		}
	`

	GetStorageImageResponseTemplateV1 = `
	{
		"labels": {
			"os": "linux",
			"version": "13",
			"base": "debian"
		},
		"metadata": {
			"name": "{{.Name}}"
		},
		"annotations": {
			"name": "Debian Container",
			"description": "The image contains the debian image base including\npreinstalled software for the use of linux containers.\n",
			"release": "2025-01-01T00:00:00Z",
			"eol": "2026-01-01T00:00:00Z",
			"recommendedCpu": "2",
			"recommendedMemory": "2",
			"recommendedNics": "2",
			"recommendedStorageSize": "100"
		},
		"spec": {
			"blockStorageRef": "block-storages/temp-a9b2bc3d81",
			"cpuArchitecture": "amd64",
			"boot": "UEFI",
			"initializer": "cloudinit-22"
		}
	}`

	CreateOrUpdateImageResponseTemplateV1 = `
	{
		"labels": {
			"os": "linux",
			"version": "13",
			"base": "debian"
		},
		"metadata": {
			"name": "{{.Name}}"
		},
		"annotations": {
			"name": "Debian Container",
			"description": "The image contains the debian image base including\npreinstalled software for the use of linux containers.\n",
			"release": "2025-01-01T00:00:00Z",
			"eol": "2026-01-01T00:00:00Z",
			"recommendedCpu": "2",
			"recommendedMemory": "2",
			"recommendedNics": "2",
			"recommendedStorageSize": "100"
		},
		"spec": {
			"blockStorageRef": "block-storages/temp-a9b2bc3d81",
			"cpuArchitecture": "amd64",
			"boot": "UEFI",
			"initializer": "cloudinit-22"
		}
	}`

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
