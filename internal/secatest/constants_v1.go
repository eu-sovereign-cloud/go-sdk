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
)
