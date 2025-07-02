package secatest

type ListRegionsResponseV1 struct {
	Name      string
	Providers []ListRegionsResponseProviderV1
}

type ListRegionsResponseProviderV1 struct {
	Name string
	URL  string
}

type GetRegionResponseV1 struct {
	Name      string
	Providers []GetRegionResponseProviderV1
}

type GetRegionResponseProviderV1 struct {
	Name string
	URL  string
}
type ListWorkspaceResponseV1 struct {
	Name   string
	Tenant string
	State  string
}

type GetWorkspaceResponseV1 struct {
	Name   string
	Tenant string
	State  string
}
type CreateOrUpdateWorkspaceResponseV1 struct {
	Name   string
	Tenant string
}

type CreateOrUpdateInstanceResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
}

type GetInstanceResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
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
type ListInstancesSkusResponseV1 struct {
	Name   string
	Tenant string
	Skus   []ListInstanceSkuMetaInfoResponseProviderV1
}
type ListStorageSkusResponseV1 struct {
	Name   string
	Tenant string
	Skus   []ListStorageSkuMetaInfoResponseProviderV1
}

type ListInstanceSkuMetaInfoResponseProviderV1 struct {
	Provider     string
	Tier         string
	Ram          int
	VCPU         int
	Architecture string
}

type ListStorageSkuMetaInfoResponseProviderV1 struct {
	Provider      string
	Tier          string
	Iops          int
	MinVolumeSize int
	Type          string
}
type GenericTenantResponseV1 struct {
	Tenant string
}
type GenericNameResponseV1 struct {
	Name string
}
type GenericNameAndTenantResponseV1 struct {
	Name   string
	Tenant string
}
