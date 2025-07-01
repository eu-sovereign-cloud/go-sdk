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
					"state": "active",
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
		"labels": {
			"env": "production"
		},
		"annotations": {
			"description": "Human readable description"
		},
		"extensions": {},
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
		"labels": {
			"env": "production"
		},
		"annotations": {
			"description": "Human readable description"
		},
		"extensions": {},
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
		"labels": {
			"env": "production"
		},
		"annotations": {
			"description": "Human readable description"
		},
		"extensions": {},
		"metadata": {
			"name": "resource-name"
		},
		"spec": {
			"vCPU": 2,
			"ram": 230
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
			"verb": "get",
			"skipToken": "null"
		}
	}`

	ListInstancesSkusResponseTemplateV1 = `{
			"items": [
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "D2XS"
				},
				"metadata": {
					"name": "seca.d2xs"
				},
				"spec": {
					"ram": 1,
					"vCPU": 1
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "DXS"
				},
				"metadata": {
					"name": "seca.dxs"
				},
				"spec": {
					"ram": 2,
					"vCPU": 1
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "DS"
				},
				"metadata": {
					"name": "seca.ds"
				},
				"spec": {
					"ram": 4,
					"vCPU": 2
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "DM"
				},
				"metadata": {
					"name": "seca.dm"
				},
				"spec": {
					"ram": 8,
					"vCPU": 4
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "DL"
				},
				"metadata": {
					"name": "seca.dl"
				},
				"spec": {
					"ram": 16,
					"vCPU": 8
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "DXL"
				},
				"metadata": {
					"name": "seca.dxl"
				},
				"spec": {
					"ram": 32,
					"vCPU": 16
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "D2XL"
				},
				"metadata": {
					"name": "seca.d2xl"
				},
				"spec": {
					"ram": 64,
					"vCPU": 32
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "S2XS"
				},
				"metadata": {
					"name": "seca.s2xs"
				},
				"spec": {
					"ram": 1,
					"vCPU": 1
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "SXS"
				},
				"metadata": {
					"name": "seca.sxs"
				},
				"spec": {
					"ram": 2,
					"vCPU": 1
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "SS"
				},
				"metadata": {
					"name": "seca.ss"
				},
				"spec": {
					"ram": 4,
					"vCPU": 2
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "SM"
				},
				"metadata": {
					"name": "seca.sm"
				},
				"spec": {
					"ram": 8,
					"vCPU": 4
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "SL"
				},
				"metadata": {
					"name": "seca.sl"
				},
				"spec": {
					"ram": 16,
					"vCPU": 8
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "SXL"
				},
				"metadata": {
					"name": "seca.sxl"
				},
				"spec": {
					"ram": 32,
					"vCPU": 16
				}
				},
				{
				"labels": {
					"architecture": "amd64",
					"provider": "seca",
					"tier": "S2XL"
				},
				"metadata": {
					"name": "seca.s2xl"
				},
				"spec": {
					"ram": 64,
					"vCPU": 32
				}
				}
			],
			"metadata": {
				"provider": "seca.compute/v1",
				"resource": "tenants/1/skus",
				"verb": "list"
			}
	}`

	ListStorageSkusResponseTemplateV1 = `{
		"items": [
			{
				"labels": {
					"provider": "seca",
					"tier": "RD100"
				},
				"metadata": {
					"name": "seca.rd100"
				},
				"spec": {
					"iops": 100,
					"minVolumeSize": 50,
					"type": "remote-durable"
				}
			},
			{
				"labels": {
					"provider": "seca",
					"tier": "RD500"
				},
				"metadata": {
					"name": "seca.rd500"
				},
				"spec": {
					"iops": 500,
					"minVolumeSize": 50,
					"type": "remote-durable"
				}
			},
			{
				"labels": {
					"provider": "seca",
					"tier": "RD2K"
				},
				"metadata": {
					"name": "seca.rd2k"
				},
				"spec": {
					"iops": 2000,
					"minVolumeSize": 50,
					"type": "remote-durable"
				}
			},
			{
				"labels": {
					"provider": "seca",
					"tier": "RD10K"
				},
				"metadata": {
					"name": "seca.rd10k"
				},
				"spec": {
					"iops": 10000,
					"minVolumeSize": 50,
					"type": "remote-durable"
				}
			},
			{
				"labels": {
					"provider": "seca",
					"tier": "RD20K"
				},
				"metadata": {
					"name": "seca.rd20k"
				},
				"spec": {
					"iops": 20000,
					"minVolumeSize": 50,
					"type": "remote-durable"
				}
			},
			{
				"labels": {
					"provider": "seca",
					"tier": "LD100"
				},
				"metadata": {
					"name": "seca.ld100"
				},
				"spec": {
					"iops": 100,
					"minVolumeSize": 50,
					"type": "local-durable"
				}
			}
		],
		"metadata": {
			"provider": "seca.compute/v1",
			"resource": "tenants/1/skus",
			"verb": "list"
		}
	}`

	GetStorageSkuResponseTemplateV1 = `
	{
		"labels": {
			"env": "production"
		},
		"annotations": {
			"description": "Human readable description"
		},
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
			"labels": {
				"env": "production"
			},
			"annotations": {
				"description": "Human readable description"
			},
			"extensions": {},
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
		"labels": {
			"env": "production"
		},
		"annotations": {
			"description": "Human readable description"
		},
		"extensions": {},
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
		"labels": {
			"env": "production"
		},
		"annotations": {
			"description": "Human readable description"
		},
		"extensions": {},
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
)
