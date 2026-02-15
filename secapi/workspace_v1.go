package secapi

import (
	"context"
	"net/http"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"k8s.io/utils/ptr"
)

// Interface

type WorkspaceV1 interface {
	// Workspace
	ListWorkspaces(ctx context.Context, tid TenantID) (*Iterator[schema.Workspace], error)
	ListWorkspacesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Workspace], error)

	GetWorkspace(ctx context.Context, tref TenantReference) (*schema.Workspace, error)
	GetWorkspaceUntilState(ctx context.Context, tref TenantReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Workspace, error)

	CreateOrUpdateWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.CreateOrUpdateWorkspaceParams) (*schema.Workspace, error)
	CreateOrUpdateWorkspace(ctx context.Context, ws *schema.Workspace) (*schema.Workspace, error)

	DeleteWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.DeleteWorkspaceParams) error
	DeleteWorkspace(ctx context.Context, ws *schema.Workspace) error
}

// Dummy

type WorkspaceV1Dummy struct{}

func newWorkspaceV1Dummy() WorkspaceV1 {
	return &WorkspaceV1Dummy{}
}

/// Workspace

func (api *WorkspaceV1Dummy) ListWorkspaces(ctx context.Context, tid TenantID) (*Iterator[schema.Workspace], error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Dummy) ListWorkspacesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Workspace], error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Dummy) GetWorkspace(ctx context.Context, tref TenantReference) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Dummy) GetWorkspaceUntilState(ctx context.Context, tref TenantReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Dummy) CreateOrUpdateWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.CreateOrUpdateWorkspaceParams) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Dummy) CreateOrUpdateWorkspace(ctx context.Context, ws *schema.Workspace) (*schema.Workspace, error) {
	return nil, ErrProviderNotAvailable
}

func (api *WorkspaceV1Dummy) DeleteWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.DeleteWorkspaceParams) error {
	return ErrProviderNotAvailable
}

func (api *WorkspaceV1Dummy) DeleteWorkspace(ctx context.Context, ws *schema.Workspace) error {
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

func (api *WorkspaceV1Impl) ListWorkspaces(ctx context.Context, tid TenantID) (*Iterator[schema.Workspace], error) {
	iter := Iterator[schema.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Workspace, *string, error) {
			resp, err := api.workspace.ListWorkspacesWithResponse(ctx, schema.TenantPathParam(tid), &workspace.ListWorkspacesParams{
				Accept:    ptr.To(workspace.ListWorkspacesParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *WorkspaceV1Impl) ListWorkspacesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Workspace], error) {
	iter := Iterator[schema.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Workspace, *string, error) {
			resp, err := api.workspace.ListWorkspacesWithResponse(ctx, schema.TenantPathParam(tid), &workspace.ListWorkspacesParams{
				Accept:    ptr.To(workspace.ListWorkspacesParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
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

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *WorkspaceV1Impl) GetWorkspaceUntilState(ctx context.Context, tref TenantReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Workspace, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Workspace]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.Workspace, error) {
			resp, err := api.workspace.GetWorkspaceWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *WorkspaceV1Impl) CreateOrUpdateWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.CreateOrUpdateWorkspaceParams) (*schema.Workspace, error) {
	if err := api.validateRegionalMetadata(ws.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.workspace.CreateOrUpdateWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, params, *ws, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
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

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *WorkspaceV1Impl) DeleteWorkspace(ctx context.Context, ws *schema.Workspace) error {
	return api.DeleteWorkspaceWithParams(ctx, ws, nil)
}
