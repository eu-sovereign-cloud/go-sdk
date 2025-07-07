package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockAuthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"

	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRoles(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)
	secatest.MockListRolesV1(authSim, secatest.NameAndTenantResponseV1{
		Name:   secatest.AuthorizationRole1Name,
		Tenant: secatest.Tenant1Name})

	sm := http.NewServeMux()

	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	authIter, err := client.AuthorizationV1.ListRoles(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	auth, err := authIter.All(ctx)
	require.NoError(t, err)

	require.Len(t, auth, 1)
	assert.Equal(t, secatest.AuthorizationRole1Name, auth[0].Metadata.Name)
}

func TestGetRole(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)

	secatest.MockGetRoleV1(authSim, secatest.NameAndTenantResponseV1{
		Name:   secatest.AuthorizationRole1Name,
		Tenant: secatest.Tenant1Name})

	sm := http.NewServeMux()
	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.AuthorizationRole1Name}
	role, err := client.AuthorizationV1.GetRole(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, role)

	assert.Equal(t, secatest.AuthorizationRole1Name, role.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, role.Metadata.Tenant)

}

func TestCreateOrUpdateRole(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)

	secatest.MockCreateOrUpdateRoleV1(authSim, secatest.NameAndTenantResponseV1{
		Name:   secatest.AuthorizationRole1Name,
		Tenant: secatest.Tenant1Name,
	})

	sm := http.NewServeMux()
	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)
	role := authorization.Role{
		Metadata: &authorization.GlobalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.AuthorizationRole1Name,
		},
	}
	err = client.AuthorizationV1.CreateOrUpdateRole(ctx, &role)
	require.NoError(t, err)
}
func TestDeleteRole(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)
	secatest.MockDeleteRoleV1(authSim)

	sm := http.NewServeMux()
	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	role := &authorization.Role{
		Metadata: &authorization.GlobalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.AuthorizationRole1Name,
		},
	}

	err = client.AuthorizationV1.DeleteRole(ctx, role)
	require.NoError(t, err)
}

func TestListRoleAssignments(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)

	secatest.MockListRoleAssignmentsV1(authSim, secatest.RoleAssignmentResponseV1{
		Tenant:    secatest.Tenant1Name,
		Region:    secatest.Region1Name,
		Workspace: secatest.Workspace1Name,
	})

	sm := http.NewServeMux()
	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	iter, err := client.AuthorizationV1.ListRoleAssignments(ctx, secatest.Tenant1Name)
	require.NoError(t, err)
	require.NotNil(t, iter)

	assignments, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, assignments, 1)

	assert.Equal(t, secatest.Tenant1Name, assignments[0].Metadata.Tenant)
	assert.Equal(t, secatest.Network1Name, assignments[0].Metadata.Name)

}

func TestGetRoleAssignment(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)

	secatest.MockGetRoleAssignmentV1(authSim, secatest.RoleAssignmentResponseV1{
		Tenant:    secatest.Tenant1Name,
		Name:      secatest.AuthorizationRoleAssignment1Name,
		Region:    secatest.Region1Name,
		Workspace: secatest.Workspace1Name,
	})

	sm := http.NewServeMux()
	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.AuthorizationRoleAssignment1Name}
	assignment, err := client.AuthorizationV1.GetRoleAssignment(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, assignment)

	assert.Equal(t, secatest.Tenant1Name, assignment.Metadata.Tenant)
	assert.Equal(t, secatest.AuthorizationRoleAssignment1Name, assignment.Metadata.Name)
}

func TestCreateOrUpdateRoleAssignment(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateRoleAssignmentV1(authSim, secatest.RoleAssignmentResponseV1{
		Tenant:    secatest.Tenant1Name,
		Name:      secatest.AuthorizationRoleAssignment1Name,
		Region:    secatest.Region1Name,
		Workspace: secatest.Workspace1Name,
	})

	sm := http.NewServeMux()
	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	roleAssignment := &authorization.RoleAssignment{
		Metadata: &authorization.GlobalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.AuthorizationRoleAssignment1Name,
		},
	}
	err = client.AuthorizationV1.CreateOrUpdateRoleAssignment(ctx, roleAssignment)
	require.NoError(t, err)
}

func TestDeleteRoleAssignment(t *testing.T) {
	ctx := context.Background()

	authSim := mockAuthorization.NewMockServerInterface(t)
	secatest.MockDeleteRoleAssignmentV1(authSim)

	sm := http.NewServeMux()
	authorization.HandlerWithOptions(authSim, authorization.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	roleAssignment := &authorization.RoleAssignment{
		Metadata: &authorization.GlobalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.AuthorizationRoleAssignment1Name,
		},
	}

	err = client.AuthorizationV1.DeleteRoleAssignment(ctx, roleAssignment)
	require.NoError(t, err)
}
