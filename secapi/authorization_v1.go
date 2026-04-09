package secapi

import (
	"context"

	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Interface

type AuthorizationV1 interface {
	// Role
	ListRolesWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.Role], error)
	ListRoles(ctx context.Context, tpath TenantPath) (*Iterator[schema.Role], error)

	GetRole(ctx context.Context, tref TenantReference) (*schema.Role, error)
	GetRoleUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Role, error)

	WatchRoleUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error

	CreateOrUpdateRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.CreateOrUpdateRoleParams) (*schema.Role, error)
	CreateOrUpdateRole(ctx context.Context, role *schema.Role) (*schema.Role, error)

	DeleteRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.DeleteRoleParams) error
	DeleteRole(ctx context.Context, role *schema.Role) error

	// Role Assignments
	ListRoleAssignmentsWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.RoleAssignment], error)
	ListRoleAssignments(ctx context.Context, tpath TenantPath) (*Iterator[schema.RoleAssignment], error)

	GetRoleAssignment(ctx context.Context, tref TenantReference) (*schema.RoleAssignment, error)
	GetRoleAssignmentUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.RoleAssignment, error)

	WatchRoleAssignmentUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error

	CreateOrUpdateRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.CreateOrUpdateRoleAssignmentParams) (*schema.RoleAssignment, error)
	CreateOrUpdateRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) (*schema.RoleAssignment, error)

	DeleteRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.DeleteRoleAssignmentParams) error
	DeleteRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) error
}

// Unavailable

type AuthorizationV1Unavailable struct{}

func newAuthorizationV1Unavailable() AuthorizationV1 {
	return &AuthorizationV1Unavailable{}
}

/// Role

func (api *AuthorizationV1Unavailable) ListRolesWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.Role], error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) ListRoles(ctx context.Context, tpath TenantPath) (*Iterator[schema.Role], error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) GetRole(ctx context.Context, tref TenantReference) (*schema.Role, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) GetRoleUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Role, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) WatchRoleUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) CreateOrUpdateRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.CreateOrUpdateRoleParams) (*schema.Role, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) CreateOrUpdateRole(ctx context.Context, role *schema.Role) (*schema.Role, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) DeleteRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.DeleteRoleParams) error {
	return ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) DeleteRole(ctx context.Context, role *schema.Role) error {
	return ErrProviderNotAvailable
}

/// Role Assignment

func (api *AuthorizationV1Unavailable) ListRoleAssignmentsWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.RoleAssignment], error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) ListRoleAssignments(ctx context.Context, tpath TenantPath) (*Iterator[schema.RoleAssignment], error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) GetRoleAssignment(ctx context.Context, tref TenantReference) (*schema.RoleAssignment, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) GetRoleAssignmentUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.RoleAssignment, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) WatchRoleAssignmentUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) CreateOrUpdateRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.CreateOrUpdateRoleAssignmentParams) (*schema.RoleAssignment, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) CreateOrUpdateRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) (*schema.RoleAssignment, error) {
	return nil, ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) DeleteRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.DeleteRoleAssignmentParams) error {
	return ErrProviderNotAvailable
}

func (api *AuthorizationV1Unavailable) DeleteRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) error {
	return ErrProviderNotAvailable
}

// Impl

type AuthorizationV1Impl struct {
	API
	authorization authorization.ClientWithResponsesInterface
}

func newAuthorizationV1Impl(client *GlobalClient, authorizationsUrl string) (AuthorizationV1, error) {
	authorization, err := authorization.NewClientWithResponses(authorizationsUrl)
	if err != nil {
		return nil, err
	}

	return &AuthorizationV1Impl{API: API{authToken: client.authToken}, authorization: authorization}, nil
}

/// Role

func (api *AuthorizationV1Impl) ListRolesWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.Role], error) {
	if err := tpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.Role]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Role, *schema.ResponseMetadata, error) {
			var params *authorization.ListRolesParams
			if options == nil {
				params = &authorization.ListRolesParams{
					Accept:    AcceptHeaderJson[authorization.ListRolesParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &authorization.ListRolesParams{
					Accept:    AcceptHeaderJson[authorization.ListRolesParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.authorization.ListRolesWithResponse(ctx, schema.TenantPathParam(tpath.Tenant), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, &resp.JSON200.Metadata, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *AuthorizationV1Impl) ListRoles(ctx context.Context, tpath TenantPath) (*Iterator[schema.Role], error) {
	return api.ListRolesWithOptions(ctx, tpath, nil)
}

func (api *AuthorizationV1Impl) GetRole(ctx context.Context, tref TenantReference) (*schema.Role, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.authorization.GetRoleWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *AuthorizationV1Impl) GetRoleUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Role, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Role]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.Role, error) {
			resp, err := api.authorization.GetRoleWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *AuthorizationV1Impl) WatchRoleUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error {
	if err := tref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Role]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.authorization.GetRoleWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *AuthorizationV1Impl) CreateOrUpdateRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.CreateOrUpdateRoleParams) (*schema.Role, error) {
	if err := api.validateGlobalMetadata(role.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.authorization.CreateOrUpdateRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name, params, *role, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *AuthorizationV1Impl) CreateOrUpdateRole(ctx context.Context, role *schema.Role) (*schema.Role, error) {
	return api.CreateOrUpdateRoleWithParams(ctx, role, nil)
}

func (api *AuthorizationV1Impl) DeleteRoleWithParams(ctx context.Context, role *schema.Role, params *authorization.DeleteRoleParams) error {
	if err := api.validateGlobalMetadata(role.Metadata); err != nil {
		return err
	}

	resp, err := api.authorization.DeleteRoleWithResponse(ctx, role.Metadata.Tenant, role.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *AuthorizationV1Impl) DeleteRole(ctx context.Context, role *schema.Role) error {
	return api.DeleteRoleWithParams(ctx, role, nil)
}

/// Role Assignment

func (api *AuthorizationV1Impl) ListRoleAssignmentsWithOptions(ctx context.Context, tpath TenantPath, options *ListOptions) (*Iterator[schema.RoleAssignment], error) {
	if err := tpath.validate(); err != nil {
		return nil, err
	}

	iter := Iterator[schema.RoleAssignment]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.RoleAssignment, *schema.ResponseMetadata, error) {
			var params *authorization.ListRoleAssignmentsParams
			if options == nil {
				params = &authorization.ListRoleAssignmentsParams{
					Accept:    AcceptHeaderJson[authorization.ListRoleAssignmentsParamsAccept](),
					SkipToken: skipToken,
				}
			} else {
				params = &authorization.ListRoleAssignmentsParams{
					Accept:    AcceptHeaderJson[authorization.ListRoleAssignmentsParamsAccept](),
					Labels:    options.Labels.BuildPtr(),
					Limit:     options.Limit,
					SkipToken: skipToken,
				}
			}

			resp, err := api.authorization.ListRoleAssignmentsWithResponse(ctx, schema.TenantPathParam(tpath.Tenant), params, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Items, &resp.JSON200.Metadata, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *AuthorizationV1Impl) ListRoleAssignments(ctx context.Context, tpath TenantPath) (*Iterator[schema.RoleAssignment], error) {
	return api.ListRoleAssignmentsWithOptions(ctx, tpath, nil)
}

func (api *AuthorizationV1Impl) GetRoleAssignment(ctx context.Context, tref TenantReference) (*schema.RoleAssignment, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.authorization.GetRoleAssignmentWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *AuthorizationV1Impl) GetRoleAssignmentUntilState(ctx context.Context, tref TenantReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.RoleAssignment, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.RoleAssignment]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.RoleAssignment, error) {
			resp, err := api.authorization.GetRoleAssignmentWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *AuthorizationV1Impl) WatchRoleAssignmentUntilDeleted(ctx context.Context, tref TenantReference, config ResourceObserverConfig) error {
	if err := tref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.RoleAssignment]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.authorization.GetRoleAssignmentWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *AuthorizationV1Impl) CreateOrUpdateRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.CreateOrUpdateRoleAssignmentParams) (*schema.RoleAssignment, error) {
	if err := api.validateGlobalMetadata(assign.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.authorization.CreateOrUpdateRoleAssignmentWithResponse(ctx, assign.Metadata.Tenant, assign.Metadata.Name, params, *assign, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *AuthorizationV1Impl) CreateOrUpdateRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) (*schema.RoleAssignment, error) {
	return api.CreateOrUpdateRoleAssignmentWithParams(ctx, assign, nil)
}

func (api *AuthorizationV1Impl) DeleteRoleAssignmentWithParams(ctx context.Context, assign *schema.RoleAssignment, params *authorization.DeleteRoleAssignmentParams) error {
	if err := api.validateGlobalMetadata(assign.Metadata); err != nil {
		return err
	}

	resp, err := api.authorization.DeleteRoleAssignmentWithResponse(ctx, assign.Metadata.Tenant, assign.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *AuthorizationV1Impl) DeleteRoleAssignment(ctx context.Context, assign *schema.RoleAssignment) error {
	return api.DeleteRoleAssignmentWithParams(ctx, assign, nil)
}
