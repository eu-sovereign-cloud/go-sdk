package secapi

import (
	"context"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Interface

type WorkspaceV1 interface {
	// Workspace
	ListWorkspaces(ctx context.Context, filter TenantFilter) (*Iterator[schema.Workspace], error)

	GetWorkspace(ctx context.Context, tref TenantReference) (*schema.Workspace, error)
	GetWorkspaceUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Workspace, error)

	WatchWorkspaceUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error

	CreateOrUpdateWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.CreateOrUpdateWorkspaceParams) (*schema.Workspace, error)
	CreateOrUpdateWorkspace(ctx context.Context, ws *schema.Workspace) (*schema.Workspace, error)

	DeleteWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.DeleteWorkspaceParams) error
	DeleteWorkspace(ctx context.Context, ws *schema.Workspace) error
}

// Unavailable

type WorkspaceV1Unavailable struct{}

func newWorkspaceV1Unavailable() WorkspaceV1 {
	return &WorkspaceV1Unavailable{}
}

/// Workspace

func (api *WorkspaceV1Unavailable) ListWorkspaces(ctx context.Context, filter TenantFilter) (*Iterator[schema.Workspace], error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Unavailable) GetWorkspace(ctx context.Context, tref TenantReference) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Unavailable) GetWorkspaceUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Unavailable) WatchWorkspaceUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *WorkspaceV1Unavailable) CreateOrUpdateWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.CreateOrUpdateWorkspaceParams) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Unavailable) CreateOrUpdateWorkspace(ctx context.Context, ws *schema.Workspace) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Unavailable) DeleteWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.DeleteWorkspaceParams) error {
	return ErrProviderNotAvailable
}

func (api *WorkspaceV1Unavailable) DeleteWorkspace(ctx context.Context, ws *schema.Workspace) error {
	return ErrProviderNotAvailable
}

// Impl

type WorkspaceV1Impl struct {
	API
	workspace workspace.ClientWithResponsesInterface
}

func newWorkspaceV1Impl(client *RegionalClient, workspaceUrl string) (WorkspaceV1, error) {
	workspace, err := workspace.NewClientWithResponses(workspaceUrl)
	if err != nil {
		return nil, err
	}

	return &WorkspaceV1Impl{API: API{authToken: client.authToken}, workspace: workspace}, nil
}

// Workspace

func (api *WorkspaceV1Impl) ListWorkspaces(ctx context.Context, filter TenantFilter) (*Iterator[schema.Workspace], error) {
	if err := filter.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Workspace, *string, error) {
			var params *workspace.ListWorkspacesParams
			if filter.Options == nil {
				params = &workspace.ListWorkspacesParams{
					Accept:    AcceptHeaderJson[workspace.ListWorkspacesParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &workspace.ListWorkspacesParams{
					Accept:    AcceptHeaderJson[workspace.ListWorkspacesParamsAccept](),
					Labels:    filter.Options.Labels.BuildPtr(),
					Limit:     filter.Options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.workspace.ListWorkspacesWithResponse(ctx, schema.TenantPathParam(filter.Tenant), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *WorkspaceV1Impl) GetWorkspace(ctx context.Context, tref TenantReference) (*schema.Workspace, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.workspace.GetWorkspaceWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *WorkspaceV1Impl) GetWorkspaceUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Workspace, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Workspace]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.Workspace, error) {
			resp, err := api.workspace.GetWorkspaceWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *WorkspaceV1Impl) WatchWorkspaceUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error {
	if err := tref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Workspace]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.workspace.GetWorkspaceWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *WorkspaceV1Impl) CreateOrUpdateWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.CreateOrUpdateWorkspaceParams) (*schema.Workspace, error) {
	if err := api.validateRegionalMetadata(ws.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.workspace.CreateOrUpdateWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, params, *ws, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *WorkspaceV1Impl) CreateOrUpdateWorkspace(ctx context.Context, ws *schema.Workspace) (*schema.Workspace, error) {
	return api.CreateOrUpdateWorkspaceWithParams(ctx, ws, nil)
}

func (api *WorkspaceV1Impl) DeleteWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.DeleteWorkspaceParams) error {
	if err := api.validateRegionalMetadata(ws.Metadata); err != nil {
		return err
	}

	resp, err := api.workspace.DeleteWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *WorkspaceV1Impl) DeleteWorkspace(ctx context.Context, ws *schema.Workspace) error {
	return api.DeleteWorkspaceWithParams(ctx, ws, nil)
}
