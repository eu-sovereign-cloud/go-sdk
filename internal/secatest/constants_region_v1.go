package secatest

const (
	// Region
	regionsTemplateV1 = `
	{
		"items": [
			{{- range $i, $r := . }}
			{{if $i}},{{end}}
			{
				"metadata": {
					"name": "{{$r.Metadata.Name}}"
				},
				"spec": {
					"providers": [
						{{- range $i, $p := $r.Providers }}
						{{if $i}},{{end}}
						{
							"name": "{{$p.Name}}",
							"url": "{{$p.URL}}",
							"version": "{{$p.Version}}"
						}
						{{- end}}
					]
				}
			}{{- end}}
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
					"url": "{{$p.URL}}",
					"version": "{{$p.Version}}"
				}
				{{- end}}
			]
		}
	}`
)
