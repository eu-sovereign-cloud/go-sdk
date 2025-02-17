package gosdk

import (
	"context"
	"fmt"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/workspace.v1"
	"k8s.io/utils/ptr"
)

type WorkspaceID string

type Workspace ReferencedResource[TenantReference, *workspace.Workspace]

func (client *RegionClient) Workspaces(ctx context.Context, tid TenantID) (*Iterator[Workspace], error) {
	wsClient, err := client.workspaceClient()
	if err != nil {
		return nil, err
	}

	iter := Iterator[Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]Workspace, *string, error) {
			resp, err := wsClient.ListWorkspacesWithResponse(ctx, workspace.TenantID(tid), &workspace.ListWorkspacesParams{
				Accept: ptr.To(workspace.ListWorkspacesParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			items := make([]Workspace, len(resp.JSON200.Items))
			for i := 0; i < len(resp.JSON200.Items); i++ {
				items[i] = Workspace{
					Ref:  TenantReference{Tenant: tid, Name: resp.JSON200.Items[i].Metadata.Name},
					Data: &resp.JSON200.Items[i],
				}
			}

			return items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (client *RegionClient) Workspace(ctx context.Context, tref TenantReference) (*Workspace, error) {
	wsClient, err := client.workspaceClient()
	if err != nil {
		return nil, err
	}

	resp, err := wsClient.GetWorkspaceWithResponse(ctx, workspace.TenantID(tref.Tenant), string(tref.Name))
	if err != nil {
		return nil, err
	}

	return &Workspace{
		Ref:  tref,
		Data: resp.JSON200,
	}, nil
}

func (client *RegionClient) DeleteWorkspace(ctx context.Context, ws *Workspace) error {

	wsClient, err := client.workspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.DeleteWorkspaceWithResponse(ctx, workspace.TenantID(ws.Ref.Tenant), ws.Ref.Name, &workspace.DeleteWorkspaceParams{
		IfUnmodifiedSince: ws.Data.Metadata.LastModifiedTimestamp,
	})
	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 || resp.StatusCode() != 404 {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode())
	}

	return nil
}

func (client *RegionClient) SaveWorkspace(ctx context.Context, ws *Workspace) error {
	wsClient, err := client.workspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.CreateOrUpdateWorkspaceWithResponse(ctx, workspace.TenantID(ws.Ref.Tenant), ws.Ref.Name,
		&workspace.CreateOrUpdateWorkspaceParams{
			IfUnmodifiedSince: ws.Data.Metadata.LastModifiedTimestamp,
		}, *ws.Data)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode())
	}

	return nil
}
