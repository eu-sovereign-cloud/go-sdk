package secatest

const (
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

	InstancesResponseV1 = `
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
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
			"provider": "seca.compute/v1",
			"resource": "tenants/1/workspaces/ws-1/instances/my-server",
			"verb": "get"
		}
	}`
)
