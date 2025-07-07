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
					"tenant": "{{.Tenant}}",
					"workspace": "{{.Workspace}}",
					"region": "{{.Region}}"
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
		    "provider": "seca.compute/v1",
    		"resource": "tenants/1/workspaces/ws-1/instances/my-server",
    		"verb": "get"
		}
	}`

	WorkspaceResponseV1 = `
	{

		"metadata": {
			"name": "{{.Name}}",
			"tenant": "{{.Tenant}}",
			"workspace": "{{.Workspace}}",
			"region": "{{.Region}}",
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
