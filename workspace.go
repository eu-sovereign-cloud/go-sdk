package gosdk

import (
	"context"
	"fmt"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/workspace.v1"
	"k8s.io/utils/ptr"
)

func (client *RegionClient) Workspaces(ctx context.Context) (*Iterator[workspace.Workspace], error) {
	tid := workspace.TenantID(MustTenantIDFromContext(ctx))

	wsClient, err := client.workspaceClient()
	if err != nil {
		return nil, err
	}

	iter := Iterator[workspace.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]workspace.Workspace, *string, error) {
			resp, err := wsClient.ListWorkspacesWithResponse(ctx, tid, &workspace.ListWorkspacesParams{
				Accept: ptr.To(workspace.ListWorkspacesParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return *resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (client *RegionClient) Workspace(ctx context.Context, name string) (*workspace.Workspace, error) {
	tid := workspace.TenantID(MustTenantIDFromContext(ctx))

	wsClient, err := client.workspaceClient()
	if err != nil {
		return nil, err
	}

	resp, err := wsClient.GetWorkspaceWithResponse(ctx, tid, name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (client *RegionClient) DeleteWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	tid := workspace.TenantID(MustTenantIDFromContext(ctx))

	wsClient, err := client.workspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.DeleteWorkspaceWithResponse(ctx, tid, ws.Spec.Name, &workspace.DeleteWorkspaceParams{
		IfUnmodifiedSince: ws.Metadata.LastModifiedTimestamp,
	})
	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 || resp.StatusCode() != 404 {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode())
	}

	return nil
}

func (client *RegionClient) SaveWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	tid := workspace.TenantID(MustTenantIDFromContext(ctx))

	wsClient, err := client.workspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.CreateOrUpdateWorkspaceWithResponse(ctx, tid, ws.Spec.Name,
		&workspace.CreateOrUpdateWorkspaceParams{
			IfUnmodifiedSince: ws.Metadata.LastModifiedTimestamp,
		}, *ws)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode())
	}

	return nil
}
