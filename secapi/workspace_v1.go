package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
)

type WorkspaceV1 struct {
	workspace workspace.ClientWithResponsesInterface
}

func (api *WorkspaceV1) ListWorkspaces(ctx context.Context, tid TenantID) (*Iterator[workspace.Workspace], error) {
	iter := Iterator[workspace.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]workspace.Workspace, *string, error) {
			resp, err := api.workspace.ListWorkspacesWithResponse(ctx, workspace.Tenant(tid), &workspace.ListWorkspacesParams{
				Accept: ptr.To(workspace.ListWorkspacesParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *WorkspaceV1) GetWorkspace(ctx context.Context, tref TenantReference) (*workspace.Workspace, error) {
	resp, err := api.workspace.GetWorkspaceWithResponse(ctx, workspace.Tenant(tref.Tenant), string(tref.Name))
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *WorkspaceV1) CreateOrUpdateWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	if err := validateWorkspaceMetadataV1(ws.Metadata); err != nil {
		return err
	}

	resp, err := api.workspace.CreateOrUpdateWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name,
		&workspace.CreateOrUpdateWorkspaceParams{
			IfUnmodifiedSince: &ws.Metadata.ResourceVersion,
		}, *ws)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *WorkspaceV1) DeleteWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	if err := validateWorkspaceMetadataV1(ws.Metadata); err != nil {
		return err
	}

	resp, err := api.workspace.DeleteWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, &workspace.DeleteWorkspaceParams{
		IfUnmodifiedSince: &ws.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 202); err != nil {
		return err
	}

	return nil
}

func newWorkspaceV1(workspaceUrl string) (*WorkspaceV1, error) {
	workspace, err := workspace.NewClientWithResponses(workspaceUrl)
	if err != nil {
		return nil, err
	}

	return &WorkspaceV1{workspace: workspace}, nil
}

func validateWorkspaceMetadataV1(metadata *workspace.RegionalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	return nil
}
