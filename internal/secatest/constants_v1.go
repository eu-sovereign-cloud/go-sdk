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
			"tenant": "{{.Tenant}}"
			"region": "{{.Region}}",
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

	GetInstanceResponseTemplateV1 = ``

	GetInstanceSkuResponseTemplateV1 = ``

	ListInstancesResponseTemplateV1 = ``

	ListInstancesSkusResponseTemplateV1 = ``
)
