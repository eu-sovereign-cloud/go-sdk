package secatest

const (
	// Workspace
	workspacesResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Metadata.Name}}"
				},
				"spec": {
				},
				"status": {
					"state": "{{.Status.State}}"
				}
			}
		]
	}`

	workspaceResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Metadata.Name}}"
		},
		"spec": {
		},
		"status": {
			"state": "{{.Status.State}}"
		}
	}`
)
