package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockauthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"
	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Role

func TestListRoles(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockListRolesV1(sim, secatest.RoleResponseV1{
		Metadata:       secatest.MetadataResponseV1{Name: secatest.AuthorizationRole1Name},
		PermissionVerb: secatest.AuthorizationPermissionVerb,
		Status:         secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	iter, err := client.AuthorizationV1.ListRoles(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	require.NotNil(t, resp[0].Metadata.Name)
	assert.Equal(t, secatest.AuthorizationRole1Name, resp[0].Metadata.Name)

	require.Len(t, resp[0].Spec.Permissions, 1)
	require.Len(t, resp[0].Spec.Permissions[0].Verb, 1)
	require.NotNil(t, resp[0].Spec.Permissions[0].Verb[0])
	assert.Equal(t, secatest.AuthorizationPermissionVerb, resp[0].Spec.Permissions[0].Verb[0])

	require.NotNil(t, resp[0].Status.State)
	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetRole(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockGetRoleV1(sim, secatest.RoleResponseV1{
		Metadata:       secatest.MetadataResponseV1{Name: secatest.AuthorizationRole1Name},
		PermissionVerb: secatest.AuthorizationPermissionVerb,
		Status:         secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.AuthorizationRole1Name}
	resp, err := client.AuthorizationV1.GetRole(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, resp.Metadata.Name)
	assert.Equal(t, secatest.AuthorizationRole1Name, resp.Metadata.Name)

	require.Len(t, resp.Spec.Permissions, 1)
	require.Len(t, resp.Spec.Permissions[0].Verb, 1)
	require.NotNil(t, resp.Spec.Permissions[0].Verb[0])
	assert.Equal(t, secatest.AuthorizationPermissionVerb, resp.Spec.Permissions[0].Verb[0])

	require.NotNil(t, resp.Status.State)
	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRole(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateRoleV1(sim, secatest.RoleResponseV1{
		Metadata:       secatest.MetadataResponseV1{Name: secatest.AuthorizationRole1Name},
		PermissionVerb: secatest.AuthorizationPermissionVerb,
		Status:         secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

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
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockDeleteRoleV1(sim)
	secatest.ConfigureAuthorizationHandler(sim, sm)

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

// Role Assignment

func TestListRoleAssignments(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockListRoleAssignmentsV1(sim, secatest.RoleAssignmentResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.AuthorizationRoleAssignment1Name},
		Subject:  secatest.AuthorizationRoleAssignment1Subject,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	iter, err := client.AuthorizationV1.ListRoleAssignments(ctx, secatest.Tenant1Name)
	require.NoError(t, err)
	require.NotNil(t, iter)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	require.NotNil(t, resp[0].Metadata.Name)
	assert.Equal(t, secatest.AuthorizationRoleAssignment1Name, resp[0].Metadata.Name)

	require.Len(t, resp[0].Spec.Subs, 1)
	assert.Equal(t, secatest.AuthorizationRoleAssignment1Subject, resp[0].Spec.Subs[0])

	require.NotNil(t, resp[0].Status.State)
	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetRoleAssignment(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockGetRoleAssignmentV1(sim, secatest.RoleAssignmentResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.AuthorizationRoleAssignment1Name},
		Subject:  secatest.AuthorizationRoleAssignment1Subject,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{AuthorizationV1: server.URL + secatest.ProviderAuthorizationEndpoint})
	require.NoError(t, err)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.AuthorizationRoleAssignment1Name}
	resp, err := client.AuthorizationV1.GetRoleAssignment(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, resp.Metadata.Name)
	assert.Equal(t, secatest.AuthorizationRoleAssignment1Name, resp.Metadata.Name)

	require.Len(t, resp.Spec.Subs, 1)
	assert.Equal(t, secatest.AuthorizationRoleAssignment1Subject, resp.Spec.Subs[0])

	require.NotNil(t, resp.Status.State)
	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRoleAssignment(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateRoleAssignmentV1(sim, secatest.RoleAssignmentResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.AuthorizationRoleAssignment1Name},
		Subject:  secatest.AuthorizationRoleAssignment1Subject,
		Status:   secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

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
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockDeleteRoleAssignmentV1(sim)
	secatest.ConfigureAuthorizationHandler(sim, sm)

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
