package secatest

const (
	// Response Templates

	GetWellknownResponseTemplateV1 = `
	{
		"version": "v1",
		"endpoints": [
			{{- range $i, $ep := .Endpoints }}
			{{if $i}},{{end}}
			{
				"provider": "{{$ep.Provider}}",
				"url": "{{$ep.URL}}"
			}
			{{- end}}
		]
	}`

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
)
