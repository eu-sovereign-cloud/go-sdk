package secatest

const (
	// Response Templates

	ItemsWorkspaceResponseV1 = `
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
					"state": "{{.State}}",
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
					"name": "{{.Name}}",
					"tenant": "{{.Tenant}}"
				}
	}`

	WorkspaceResponseV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
			"apiVersion": "v1",
			"kind": "workspace"
		},
		"spec": {},
		"status": {
			"state": "{{.State}}",
			"conditions": [
				{
					"status": "Ready"
				}
			]
		},
		"resourcesCount": 1
	}`
)
