package secatest

const (
	// Instance Sku
	instanceSkusResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Metadata.Name}}"
				},
				"spec: {
					"tier": "{{.Tier}}"
				},
				"status": {
					"state": "{{.Status.State}}"
				}
			}
		]
	}`
	instanceSkuResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Metadata.Name}}"
		},
		"spec": {
			"tier": "{{.Tier}}"
		},
		"status": {
			"state": "{{.Status.State}}"
		}
	}`

	// Instance
	instancesResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Metadata.Name}}"
				},
				"spec": {
					"skuRef": "{{.SkuRef}}"
				},
				"status": {
					"state": "{{.Status.State}}"
				}
			}
		]
	}`
	instanceResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Metadata.Name}}"
		},
		"spec": {
			"skuRef": "{{.SkuRef}}"
		},
		"status": {
			"state": "{{.Status.State}}"
		}
	}`
)
