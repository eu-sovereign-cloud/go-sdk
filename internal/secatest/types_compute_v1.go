package secatest

type ListInstanceSkuMetaInfoResponseProviderV1 struct {
	Provider     string
	Tier         string
	Ram          int
	VCPU         int
	Architecture string
}
type ListInstancesSkusResponseV1 struct {
	Name   string
	Tenant string
	Skus   []ListInstanceSkuMetaInfoResponseProviderV1
}

type GetInstanceSkuResponseV1 struct {
	Name   string
	Tenant string
	VCPU   int
	Ram    int
}

type ListInstancesResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
}

type GetInstanceResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
}

type CreateOrUpdateInstanceResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
}
