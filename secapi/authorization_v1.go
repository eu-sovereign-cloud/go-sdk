package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secapi"
	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
)

type AuthorizationV1 struct {
	secapi.GlobalAPI[authorization.ClientWithResponsesInterface]
}

func newAuthorizationV1() *AuthorizationV1 {
	return &AuthorizationV1{
		GlobalAPI: secapi.GlobalAPI[authorization.ClientWithResponsesInterface]{},
	}
}

func (api *AuthorizationV1) getClient() (authorization.ClientWithResponsesInterface, error) {
	fn := func(url string) (authorization.ClientWithResponsesInterface, error) {
		return authorization.NewClientWithResponses(url)
	}

	client, err := api.GetClient("seca.authorization", fn)
	if err != nil {
		return nil, err
	}
	return *client, nil
}

func validateAuthorizationMetadataV1(metadata *authorization.GlobalResourceMetadata) {
	if metadata == nil {
		panic(ErrNoMetatada)
	}

	if metadata.Tenant == "" {
		panic(ErrNoMetatadaTenant)
	}
}

func (api *AuthorizationV1) ListRoles(ctx context.Context, tid TenantID) (*Iterator[authorization.Role], error) {
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	iter := Iterator[authorization.Role]{
		fn: func(ctx context.Context, skipToken *string) ([]authorization.Role, *string, error) {
			resp, err := client.ListRolesWithResponse(ctx, authorization.Tenant(tid), &authorization.ListRolesParams{
				Accept: ptr.To(authorization.Applicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *AuthorizationV1) GetRole(ctx context.Context, tref TenantReference) (*authorization.Role, error) {
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GetRoleWithResponse(ctx, authorization.Tenant(tref.Tenant), string(tref.Name))
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *AuthorizationV1) CreateOrUpdateRole(ctx context.Context, role *authorization.Role) error {
	validateAuthorizationMetadataV1(role.Metadata)

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.CreateOrUpdateRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name,
		&authorization.CreateOrUpdateRoleParams{
			IfUnmodifiedSince: &role.Metadata.ResourceVersion,
		}, *role)
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 200, 201)
	if err != nil {
		return err
	}

	return nil
}

func (api *AuthorizationV1) DeleteRole(ctx context.Context, role *authorization.Role) error {
	validateAuthorizationMetadataV1(role.Metadata)

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.DeleteRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name, &authorization.DeleteRoleParams{
		IfUnmodifiedSince: &role.Metadata.ResourceVersion,
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
