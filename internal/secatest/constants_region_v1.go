package secatest

const (
	// Region
	regionsTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Metadata.Name}}"
				},
				"spec": {
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
					"state": "{{.Status.State}}"
				}
			}
		]
	}`
	regionResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Metadata.Name}}"
		},
		"spec": {
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
			"state": "{{.Status.State}}"
		}
	}`
)
