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

type TenantResponseV1 struct {
	Tenant string
}
type NameResponseV1 struct {
	Name string
}
type NameAndTenantResponseV1 struct {
	Name   string
	Tenant string
}
