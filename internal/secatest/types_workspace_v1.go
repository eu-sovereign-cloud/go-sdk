package secatest

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
