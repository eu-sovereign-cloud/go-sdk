package secatest

const (
	// Instance Sku
	instanceSkusResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec: {
					"tier": "{{.Tier}}"
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
		]
	}`
	instanceSkuResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"tier": "{{.Tier}}"
		},
		"status": {
			"state": "{{.Status}}"
		}
	}`

	// Instance
	instancesResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"skuRef": "{{.SkuRef}}"
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
		]
	}`
	instanceResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"skuRef": "{{.SkuRef}}"
		},
		"status": {
			"state": "{{.Status}}"
		}
	}`
)
