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
				"labels": {
					"tier": "{{.Tier}}"
				},
				"spec": {
					"vCPU": {{.VCPU}}
				}
			}
		]
	}`
	instanceSkuResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Metadata.Name}}"
		},
		"labels": {
			"tier": "{{.Tier}}"
		},
		"spec": {
			"vCPU": {{.VCPU}}
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
