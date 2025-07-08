package secatest

const (
	// Storage Sku
	storageSkusResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Metadata.Name}}"
				},
				"labels": {
					"iops": "{{.Iops}}"
				},
				"spec": {
					"vCPU": {{.VCPU}}
				}
			}
		]
	}`
	storageSkuResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Metadata.Name}}"
		},
		"spec": {
			"type": "{{.Type}}"
		},
		"status": {
			"state": "{{.Status.State}}"
		}
	}`

	// Block Storage
	blockStoragesResponseTemplateV1 = `
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
	blockStorageResponseTemplateV1 = `
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

	// Image
	imagesResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Metadata.Name}}"
				},
				"spec": {
					"blockStorageRef": "{{.BlockStorageRef}}"
				},
				"status": {
					"state": "{{.Status.State}}"
				}
			}
		]
	}`
	imageResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Metadata.Name}}"
		},
		"spec": {
			"blockStorageRef": "{{.BlockStorageRef}}"
		},
		"status": {
			"state": "{{.Status.State}}"
		}
	}`
)
