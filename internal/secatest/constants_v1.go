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
)
