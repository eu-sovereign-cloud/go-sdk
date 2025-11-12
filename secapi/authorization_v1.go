package secapi

import (
	"context"
	"net/http"

	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	. "github.com/eu-sovereign-cloud/go-sdk/secapi/builders"
	"k8s.io/utils/ptr"
)

type AuthorizationV1 struct {
	API
	authorization authorization.ClientWithResponsesInterface
}

// Role

func (api *AuthorizationV1) ListRoles(ctx context.Context, tid TenantID) (*Iterator[schema.Role], error) {
	iter := Iterator[schema.Role]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Role, *string, error) {
			resp, err := api.authorization.ListRolesWithResponse(ctx, schema.TenantPathParam(tid), &authorization.ListRolesParams{
				Accept: ptr.To(authorization.ListRolesParamsAccept(schema.AcceptHeaderJson)),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *AuthorizationV1) ListRolesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Role], error) {
	iter := Iterator[schema.Role]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Role, *string, error) {
			resp, err := api.authorization.ListRolesWithResponse(ctx, schema.TenantPathParam(tid), &authorization.ListRolesParams{
				Accept:    ptr.To(authorization.ListRolesParamsAccept(schema.AcceptHeaderJson)),
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

func (api *AuthorizationV1) GetRole(ctx context.Context, tref TenantReference) (*schema.Role, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.authorization.GetRoleWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *AuthorizationV1) GetRoleUntilState(ctx context.Context, tref TenantReference, config ResourceStateObserverConfig) (*schema.Role, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Role]{
		delay:       config.delay,
		interval:    config.interval,
		maxAttempts: config.maxAttempts,
		actFunc: func() (schema.ResourceState, *schema.Role, error) {
			resp, err := api.authorization.GetRoleWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
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

	resp, err := observer.WaitUntil(config.expectedState)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *AuthorizationV1) CreateOrUpdateRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.CreateOrUpdateRoleParams) (*schema.Role, error) {
	if err := api.validateGlobalMetadata(role.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.authorization.CreateOrUpdateRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name, params, *role, api.loadRequestHeaders)
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

func (api *AuthorizationV1) CreateOrUpdateRole(ctx context.Context, role *schema.Role) (*schema.Role, error) {
	return api.CreateOrUpdateRoleWithParams(ctx, role, nil)
}

func (api *AuthorizationV1) DeleteRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.DeleteRoleParams) error {
	if err := api.validateGlobalMetadata(role.Metadata); err != nil {
		return err
	}

	resp, err := api.authorization.DeleteRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *AuthorizationV1) DeleteRole(ctx context.Context, role *schema.Role) error {
	return api.DeleteRoleWithParams(ctx, role, nil)
}

// Role Assignment

func (api *AuthorizationV1) ListRoleAssignments(ctx context.Context, tid TenantID) (*Iterator[schema.RoleAssignment], error) {
	iter := Iterator[schema.RoleAssignment]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.RoleAssignment, *string, error) {
			resp, err := api.authorization.ListRoleAssignmentsWithResponse(ctx, schema.TenantPathParam(tid), &authorization.ListRoleAssignmentsParams{
				Accept: ptr.To(authorization.ListRoleAssignmentsParamsAccept(schema.AcceptHeaderJson)),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *AuthorizationV1) ListRoleAssignmentsWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.RoleAssignment], error) {
	iter := Iterator[schema.RoleAssignment]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.RoleAssignment, *string, error) {
			resp, err := api.authorization.ListRoleAssignmentsWithResponse(ctx, schema.TenantPathParam(tid), &authorization.ListRoleAssignmentsParams{
				Accept:    ptr.To(authorization.ListRoleAssignmentsParamsAccept(schema.AcceptHeaderJson)),
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

func (api *AuthorizationV1) GetRoleAssignment(ctx context.Context, tref TenantReference) (*schema.RoleAssignment, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.authorization.GetRoleAssignmentWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *AuthorizationV1) GetRoleAssignmentUntilState(ctx context.Context, tref TenantReference, config ResourceStateObserverConfig) (*schema.RoleAssignment, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.RoleAssignment]{
		delay:       config.delay,
		interval:    config.interval,
		maxAttempts: config.maxAttempts,
		actFunc: func() (schema.ResourceState, *schema.RoleAssignment, error) {
			resp, err := api.authorization.GetRoleAssignmentWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
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

	resp, err := observer.WaitUntil(config.expectedState)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *AuthorizationV1) CreateOrUpdateRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.CreateOrUpdateRoleAssignmentParams) (*schema.RoleAssignment, error) {
	if err := api.validateGlobalMetadata(assign.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.authorization.CreateOrUpdateRoleAssignmentWithResponse(ctx, assign.Metadata.Tenant, assign.Metadata.Name, params, *assign, api.loadRequestHeaders)
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

func (api *AuthorizationV1) CreateOrUpdateRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) (*schema.RoleAssignment, error) {
	return api.CreateOrUpdateRoleAssignmentWithParams(ctx, assign, nil)
}

func (api *AuthorizationV1) DeleteRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.DeleteRoleAssignmentParams) error {
	if err := api.validateGlobalMetadata(assign.Metadata); err != nil {
		return err
	}

	resp, err := api.authorization.DeleteRoleAssignmentWithResponse(ctx, assign.Metadata.Tenant, assign.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *AuthorizationV1) DeleteRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) error {
	return api.DeleteRoleAssignmentWithParams(ctx, assign, nil)
}

func newAuthorizationV1(client *GlobalClient, authorizationsUrl string) (*AuthorizationV1, error) {
	authorization, err := authorization.NewClientWithResponses(authorizationsUrl)
	if err != nil {
		return nil, err
	}

	return &AuthorizationV1{API: API{authToken: client.authToken}, authorization: authorization}, nil
}
