package secatest

type GetWellknownResponseV1 struct {
	Endpoints []GetWellknownResponseEndpointV1
}

type GetWellknownResponseEndpointV1 struct {
	Provider string
	URL      string
}

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
