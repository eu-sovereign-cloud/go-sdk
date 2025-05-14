package secapi

import (
	"context"
	"fmt"

	//     "github.com/aws/aws-sdk-go/aws"
	//     "github.com/aws/aws-sdk-go/aws/session"
	//     "github.com/aws/aws-sdk-go/service/s3"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/foundation.workspace.v1"
	"k8s.io/utils/ptr"
)

type WorkspaceID string

func (client *RegionClient) Workspaces(ctx context.Context, tid TenantID) (*Iterator[workspace.Workspace], error) {
	wsClient, err := client.workspaceClient()
	if err != nil {
		return nil, err
	}

	iter := Iterator[workspace.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]workspace.Workspace, *string, error) {
			resp, err := wsClient.ListWorkspacesWithResponse(ctx, workspace.TenantID(tid), &workspace.ListWorkspacesParams{
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

func (client *RegionClient) Workspace(ctx context.Context, tref TenantReference) (*workspace.Workspace, error) {
	wsClient, err := client.workspaceClient()
	if err != nil {
		return nil, err
	}

	resp, err := wsClient.GetWorkspaceWithResponse(ctx, workspace.TenantID(tref.Tenant), string(tref.Name))
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (client *RegionClient) DeleteWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	panicUnlessTenantExists(ws)

	wsClient, err := client.workspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.DeleteWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, &workspace.DeleteWorkspaceParams{
		IfUnmodifiedSince: &ws.Metadata.LastModifiedTimestamp,
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
	panicUnlessTenantExists(ws)

	wsClient, err := client.workspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.CreateOrUpdateWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name,
		&workspace.CreateOrUpdateWorkspaceParams{
			IfUnmodifiedSince: &ws.Metadata.LastModifiedTimestamp,
		}, *ws)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode())
	}

	return nil
}

func panicUnlessTenantExists(ws *workspace.Workspace) {
	if ws.Metadata == nil {
		panic("Metadata is nil")
	}

	if ws.Metadata.Tenant == "" {
		panic("Metadata.Tenant is empty")
	}
}
