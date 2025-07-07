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

type BlockStorageResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
	Region    string
	Zone      string
}
type ImageResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
	Region    string
}
