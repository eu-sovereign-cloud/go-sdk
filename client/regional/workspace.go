package regional

import (
	"context"
	"fmt"

	"k8s.io/utils/ptr"
	
	"github.com/eu-sovereign-cloud/go-sdk/client"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
)

func (rc *RegionalClient) ListWorkspaces(ctx context.Context, tid client.TenantID) (*client.Iterator[workspace.Workspace], error) {
	wsClient, err := rc.getWorkspaceClient()
	if err != nil {
		return nil, err
	}

	iter := client.Iterator[workspace.Workspace]{
		Func: func(ctx context.Context, skipToken *string) ([]workspace.Workspace, *string, error) {
			resp, err := wsClient.ListWorkspacesWithResponse(ctx, workspace.Tenant(tid), &workspace.ListWorkspacesParams{
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

func (rc *RegionalClient) GetWorkspace(ctx context.Context, tref client.TenantReference) (*workspace.Workspace, error) {
	wsClient, err := rc.getWorkspaceClient()
	if err != nil {
		return nil, err
	}

	resp, err := wsClient.GetWorkspaceWithResponse(ctx, workspace.Tenant(tref.Tenant), string(tref.Name))
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (rc *RegionalClient) CreateOrUpdateWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	panicUnlessTenantExists(ws)

	wsClient, err := rc.getWorkspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.CreateOrUpdateWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name,
		&workspace.CreateOrUpdateWorkspaceParams{
			IfUnmodifiedSince: &ws.Metadata.ResourceVersion,
		}, *ws)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode())
	}

	return nil
}

func (rc *RegionalClient) DeleteWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	panicUnlessTenantExists(ws)

	wsClient, err := rc.getWorkspaceClient()
	if err != nil {
		return err
	}

	resp, err := wsClient.DeleteWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, &workspace.DeleteWorkspaceParams{
		IfUnmodifiedSince: &ws.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 || resp.StatusCode() != 404 {
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
