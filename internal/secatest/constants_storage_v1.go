package secatest

const (
	// Storage Sku
	storageSkusResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"type": "{{.Type}}"
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
		]
	}`
	storageSkuResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"type": "{{.Type}}"
		},
		"status": {
			"state": "{{.Status}}"
		}
	}`

	// Block Storage
	blockStoragesResponseTemplateV1 = `
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
	blockStorageResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		"spec": {
			"skuRef": "{{.SkuRef}}"
		},
		"status": {
			"state": "{{.Status}}"
		}
	}`

	// Image
	imagesResponseTemplateV1 = `
	{
		"items": [
			{
				"metadata": {
					"name": "{{.Name}}"
				},
				"spec": {
					"blockStorageRef": "{{.BlockStorageRef}}"
				},
				"status": {
					"state": "{{.Status}}"
				}
			}
		]
	}`
	imageResponseTemplateV1 = `
	{
		"metadata": {
			"name": "{{.Name}}"
		},
		"spec": {
			"blockStorageRef": "{{.BlockStorageRef}}"
		},
		"status": {
			"state": "{{.Status}}"
		}
	}`
)
