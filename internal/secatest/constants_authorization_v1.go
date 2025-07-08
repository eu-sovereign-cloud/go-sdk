package secatest

const (
	// Role
	rolesResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Metadata.Name}}"
				}
				"spec": {
					"permissions": [
						{
							"verb": ["{{.PermissionVerb}}"]
						}
					]
				},
				"status": {
					"state": "{{.Status.State}}"
				}
			}
		]
	}`
	roleResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"permissions": [
				{
					"verb": ["{{.PermissionVerb}}"]
				}
			]
		},
		"status": {
			"state": "{{.Status}}"
		}
	}`

	// Role Assignment
	roleAssignmentsResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"subs": [ "{{.Subject}}" ]
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
		]
	}`
	roleAssignmentResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"subs": [ "{{.Subject}}" ]
		},
		"status": {
			"state": "{{.Status}}"
		}
	}`
)
