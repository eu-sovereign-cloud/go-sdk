package secapi

import "context"

type CRUD[T any, tenantT string, workspaceT string, paramsT any, reT any, lbsT any, nT string, rT any] struct {
	ReturnFn       func(*lbsT) ([]T, *string, error)
	ReturnSingleFn func(*rT) (*T, error)
	ParamsFn       func() *paramsT
	ListFn         func(ctx context.Context, tenant tenantT, workspace workspaceT, params *paramsT, reqEditors ...reT) (*lbsT, error)
	GetFn          func(ctx context.Context, tenant tenantT, workspace workspaceT, name nT, reqEditors ...reT) (*rT, error)
}

func (c *CRUD[T, tenantT, workspaceT, paramsT, reT, lbsT, nT, rT]) List(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[T], error) {
	iter := Iterator[T]{
		fn: func(ctx context.Context, skipToken *string) ([]T, *string, error) {
			resp, err := c.ListFn(ctx, tenantT(tid), workspaceT(wid), c.ParamsFn())
			if err != nil {
				return nil, nil, err
			}

			return c.ReturnFn(resp)
		},
	}

	return &iter, nil
}

type Get[T any, tenantT string, workspaceT string, reT any, nT string, rT any] struct {
	ReturnSingleFn func(*rT) (*T, error)
	GetFn          func(ctx context.Context, tenant tenantT, workspace workspaceT, name nT, reqEditors ...reT) (*rT, error)
}

func (c *CRUD[T, tenantT, workspaceT, paramsT, reT, lbsT, nT, rT]) Get(ctx context.Context, wref WorkspaceReference) (*T, error) {
	resp, err := c.GetFn(ctx, tenantT(wref.Tenant), workspaceT(wref.Workspace), nT(wref.Name))
	if err != nil {
		return nil, err
	}

	return c.ReturnSingleFn(resp)
}
