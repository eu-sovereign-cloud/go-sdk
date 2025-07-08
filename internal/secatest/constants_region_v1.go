package secatest

const (
	// Region
	regionsTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Name}}"
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
					"state": "{{.Status}}"
				}
			}
		]
	}`
	regionResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
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
			"state": "{{.Status}}"
		}
	}`
)
