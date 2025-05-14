package secapi

type Reference interface {
	TenantReference | WorkspaceReference
}

type ReferencedResource[Ref Reference, Data any] struct {
	Ref  Ref
	Data Data
}

type TenantReference struct {
	Tenant TenantID
	Name   string
}

type WorkspaceReference struct {
	TenantReference

	Workspace WorkspaceID
}
