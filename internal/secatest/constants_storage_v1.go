package secatest

const (
	// Response Templates
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
)
