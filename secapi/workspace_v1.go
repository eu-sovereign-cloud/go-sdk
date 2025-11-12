package secapi

import (
	"context"
	"net/http"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	. "github.com/eu-sovereign-cloud/go-sdk/secapi/builders"
	"k8s.io/utils/ptr"
)

type WorkspaceV1 struct {
	API
	workspace workspace.ClientWithResponsesInterface
}

// Workspace

func (api *WorkspaceV1) ListWorkspaces(ctx context.Context, tid TenantID) (*Iterator[schema.Workspace], error) {
	iter := Iterator[schema.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Workspace, *string, error) {
			resp, err := api.workspace.ListWorkspacesWithResponse(ctx, schema.TenantPathParam(tid), &workspace.ListWorkspacesParams{
				Accept:    ptr.To(workspace.ListWorkspacesParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *WorkspaceV1) ListWorkspacesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Workspace], error) {
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

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *WorkspaceV1) GetWorkspace(ctx context.Context, tref TenantReference) (*schema.Workspace, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.workspace.GetWorkspaceWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *WorkspaceV1) GetWorkspaceUntilState(ctx context.Context, tref TenantReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Workspace, error) {
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

			if resp.StatusCode() == http.StatusNotFound {
				return "", nil, ErrResourceNotFound
			} else {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *WorkspaceV1) CreateOrUpdateWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.CreateOrUpdateWorkspaceParams) (*schema.Workspace, error) {
	if err := api.validateRegionalMetadata(ws.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.workspace.CreateOrUpdateWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, params, *ws, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *WorkspaceV1) CreateOrUpdateWorkspace(ctx context.Context, ws *schema.Workspace) (*schema.Workspace, error) {
	return api.CreateOrUpdateWorkspaceWithParams(ctx, ws, nil)
}

func (api *WorkspaceV1) DeleteWorkspaceWithParams(ctx context.Context, ws *schema.Workspace, params *workspace.DeleteWorkspaceParams) error {
	if err := api.validateRegionalMetadata(ws.Metadata); err != nil {
		return err
	}

	resp, err := api.workspace.DeleteWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *WorkspaceV1) DeleteWorkspace(ctx context.Context, ws *schema.Workspace) error {
	return api.DeleteWorkspaceWithParams(ctx, ws, nil)
}

func newWorkspaceV1(client *RegionalClient, workspaceUrl string) (*WorkspaceV1, error) {
	workspace, err := workspace.NewClientWithResponses(workspaceUrl)
	if err != nil {
		return nil, err
	}

	return &WorkspaceV1{API: API{authToken: client.authToken}, workspace: workspace}, nil
}
