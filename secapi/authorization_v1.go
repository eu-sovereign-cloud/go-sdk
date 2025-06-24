package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
)

type AuthorizationV1 struct {
	authorization authorization.ClientWithResponsesInterface
}

func newAuthorizationV1(authorizationsUrl string) (*AuthorizationV1, error) {
	authorization, err := authorization.NewClientWithResponses(authorizationsUrl)
	if err != nil {
		return nil, err
	}

	return &AuthorizationV1{authorization: authorization}, nil
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
	iter := Iterator[authorization.Role]{
		fn: func(ctx context.Context, skipToken *string) ([]authorization.Role, *string, error) {
			resp, err := api.authorization.ListRolesWithResponse(ctx, authorization.Tenant(tid), &authorization.ListRolesParams{
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
	resp, err := api.authorization.GetRoleWithResponse(ctx, authorization.Tenant(tref.Tenant), string(tref.Name))
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *AuthorizationV1) CreateOrUpdateRole(ctx context.Context, role *authorization.Role) error {
	validateAuthorizationMetadataV1(role.Metadata)

	resp, err := api.authorization.CreateOrUpdateRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name,
		&authorization.CreateOrUpdateRoleParams{
			IfUnmodifiedSince: &role.Metadata.ResourceVersion,
		}, *role)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *AuthorizationV1) DeleteRole(ctx context.Context, role *authorization.Role) error {
	validateAuthorizationMetadataV1(role.Metadata)

	resp, err := api.authorization.DeleteRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name, &authorization.DeleteRoleParams{
		IfUnmodifiedSince: &role.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}
