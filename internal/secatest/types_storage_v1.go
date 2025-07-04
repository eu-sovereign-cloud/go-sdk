package secatest

type ListStorageSkuMetaInfoResponseProviderV1 struct {
	Provider      string
	Tier          string
	Iops          int
	MinVolumeSize int
	Type          string
}
type ListStorageSkusResponseV1 struct {
	Name   string
	Tenant string
	Skus   []ListStorageSkuMetaInfoResponseProviderV1
}
