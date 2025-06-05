package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secapi"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
)

type WorkspaceV1 struct {
	secapi.RegionalAPI[workspace.ClientWithResponsesInterface]
}

func newWorkspaceV1(region *region.Region) *WorkspaceV1 {
	return &WorkspaceV1{
		RegionalAPI: secapi.RegionalAPI[workspace.ClientWithResponsesInterface]{Region: region},
	}
}

func (api *WorkspaceV1) getClient() (workspace.ClientWithResponsesInterface, error) {
	fn := func(url string) (workspace.ClientWithResponsesInterface, error) {
		return workspace.NewClientWithResponses(url)
	}

	client, err := api.GetClient("seca.workspace", fn)
	if err != nil {
		return nil, err
	}
	return *client, nil
}

func validateWorkspaceMetadataV1(metadata *workspace.RegionalResourceMetadata) {
	if metadata == nil {
		panic(ErrNoMetatada)
	}

	if metadata.Tenant == "" {
		panic(ErrNoMetatadaTenant)
	}
}

func (api *WorkspaceV1) ListWorkspaces(ctx context.Context, tid TenantID) (*Iterator[workspace.Workspace], error) {
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	iter := Iterator[workspace.Workspace]{
		fn: func(ctx context.Context, skipToken *string) ([]workspace.Workspace, *string, error) {
			resp, err := client.ListWorkspacesWithResponse(ctx, workspace.Tenant(tid), &workspace.ListWorkspacesParams{
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
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GetWorkspaceWithResponse(ctx, workspace.Tenant(tref.Tenant), string(tref.Name))
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *WorkspaceV1) CreateOrUpdateWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	validateWorkspaceMetadataV1(ws.Metadata)

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.CreateOrUpdateWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name,
		&workspace.CreateOrUpdateWorkspaceParams{
			IfUnmodifiedSince: &ws.Metadata.ResourceVersion,
		}, *ws)
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 200, 201)
	if err != nil {
		return err
	}

	return nil
}

func (api *WorkspaceV1) DeleteWorkspace(ctx context.Context, ws *workspace.Workspace) error {
	validateWorkspaceMetadataV1(ws.Metadata)

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.DeleteWorkspaceWithResponse(ctx, ws.Metadata.Tenant, ws.Metadata.Name, &workspace.DeleteWorkspaceParams{
		IfUnmodifiedSince: &ws.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 204, 404)
	if err != nil {
		return err
	}

	return nil
}
