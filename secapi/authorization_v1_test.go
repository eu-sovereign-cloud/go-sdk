package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockauthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi/builders"

	"github.com/stretchr/testify/assert"
)

// Role

func TestListRoles(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockListRolesV1(sim, secatest.RoleResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.Role1Name,
			Tenant: secatest.Tenant1Name,
		},
		PermissionVerb: secatest.Role1PermissionVerb,
		Status:         secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	iter, err := client.AuthorizationV1.ListRoles(ctx, secatest.Tenant1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.Role1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)

	assert.Len(t, resp[0].Spec.Permissions, 1)
	assert.Len(t, resp[0].Spec.Permissions[0].Verb, 1)
	assert.Equal(t, secatest.Role1PermissionVerb, resp[0].Spec.Permissions[0].Verb[0])

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))

	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := builders.NewListOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err = client.AuthorizationV1.ListRolesWithFilters(ctx, secatest.Tenant1Name, listOptions)
	assert.NoError(t, err)

	resp, err = iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetRole(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockGetRoleV1(sim, secatest.RoleResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.Role1Name,
			Tenant: secatest.Tenant1Name,
		},
		PermissionVerb: secatest.Role1PermissionVerb,
		Status:         secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Role1Name}
	resp, err := client.AuthorizationV1.GetRole(ctx, tref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Role1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)

	assert.Len(t, resp.Spec.Permissions, 1)
	assert.Len(t, resp.Spec.Permissions[0].Verb, 1)
	assert.Equal(t, secatest.Role1PermissionVerb, resp.Spec.Permissions[0].Verb[0])

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRole(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateRoleV1(sim, secatest.RoleResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.Role1Name,
			Tenant: secatest.Tenant1Name,
		},
		PermissionVerb: secatest.Role1PermissionVerb,
		Status:         secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	role := schema.Role{
		Metadata: &schema.GlobalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Role1Name,
		},
		Spec: schema.RoleSpec{
			Permissions: []schema.Permission{
				{
					Provider:  secatest.Role1PermissionProvider,
					Resources: []string{secatest.Role1PermissionResource},
					Verb:      []string{secatest.Role1PermissionVerb},
				},
			},
		},
	}
	resp, err := client.AuthorizationV1.CreateOrUpdateRole(ctx, &role)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Role1Name, resp.Metadata.Name)

	assert.Len(t, resp.Spec.Permissions, 1)
	assert.Len(t, resp.Spec.Permissions[0].Verb, 1)
	assert.Equal(t, secatest.Role1PermissionVerb, resp.Spec.Permissions[0].Verb[0])

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteRole(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockGetRoleV1(sim, secatest.RoleResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.Role1Name,
			Tenant: secatest.Tenant1Name,
		},
		PermissionVerb: secatest.Role1PermissionVerb,
		Status:         secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})

	secatest.MockDeleteRoleV1(sim)
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	resp, err := client.AuthorizationV1.GetRole(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Role1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = client.AuthorizationV1.DeleteRole(ctx, resp)
	assert.NoError(t, err)
}

// Role Assignment

func TestListRoleAssignments(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockListRoleAssignmentsV1(sim, secatest.RoleAssignmentResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.RoleAssignment1Name,
			Tenant: secatest.Tenant1Name,
		},
		Subject: secatest.RoleAssignment1Subject,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	iter, err := client.AuthorizationV1.ListRoleAssignments(ctx, secatest.Tenant1Name)
	assert.NoError(t, err)
	assert.NotNil(t, iter)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.RoleAssignment1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)

	assert.Len(t, resp[0].Spec.Subs, 1)
	assert.Equal(t, secatest.RoleAssignment1Subject, resp[0].Spec.Subs[0])

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetRoleAssignment(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockGetRoleAssignmentV1(sim, secatest.RoleAssignmentResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.RoleAssignment1Name,
			Tenant: secatest.Tenant1Name,
		},
		Subject: secatest.RoleAssignment1Subject,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.RoleAssignment1Name}
	resp, err := client.AuthorizationV1.GetRoleAssignment(ctx, tref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.RoleAssignment1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)

	assert.Len(t, resp.Spec.Subs, 1)
	assert.Equal(t, secatest.RoleAssignment1Subject, resp.Spec.Subs[0])

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRoleAssignment(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateRoleAssignmentV1(sim, secatest.RoleAssignmentResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.RoleAssignment1Name,
			Tenant: secatest.Tenant1Name,
		},
		Subject: secatest.RoleAssignment1Subject,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	assign := &schema.RoleAssignment{
		Metadata: &schema.GlobalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.RoleAssignment1Name,
		},
		Spec: schema.RoleAssignmentSpec{
			Roles: []string{secatest.Role1Name},
			Scopes: []schema.RoleAssignmentScope{
				{
					Tenants: &[]string{secatest.Tenant1Name},
				},
			},
			Subs: []string{secatest.RoleAssignment1Subject},
		},
	}
	resp, err := client.AuthorizationV1.CreateOrUpdateRoleAssignment(ctx, assign)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.RoleAssignment1Name, resp.Metadata.Name)

	assert.Len(t, resp.Spec.Subs, 1)
	assert.Equal(t, secatest.RoleAssignment1Subject, resp.Spec.Subs[0])

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteRoleAssignment(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockGetRoleAssignmentV1(sim, secatest.RoleAssignmentResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.RoleAssignment1Name,
			Tenant: secatest.Tenant1Name,
		},
		Subject: secatest.RoleAssignment1Subject,
		Status:  secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})

	secatest.MockDeleteRoleAssignmentV1(sim)
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	resp, err := client.AuthorizationV1.GetRoleAssignment(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.RoleAssignment1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = client.AuthorizationV1.DeleteRoleAssignment(ctx, resp)
	assert.NoError(t, err)
}
