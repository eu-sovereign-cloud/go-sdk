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
}

type GetWorkspaceResponseV1 struct {
	Name   string
	Tenant string
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
	Name      string
	Tenant    string
	Workspace string
}

type ListInstancesResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
}
type ListInstancesSkusResponseV1 struct {
	Name      string
	Tenant    string
	Workspace string
}
