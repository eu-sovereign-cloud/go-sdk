package secatest

const (
	// Workspace
	workspacesResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
				},
				"status": {
					"state": "{{.State}}",
				}
			}
		]
	}`
	workspaceResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
		},
		"status": {
			"state": "{{.State}}"
		}
	}`
)
