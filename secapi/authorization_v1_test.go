package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockauthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/secalib/builders"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

// Role

func TestListRolesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleSpec(secatest.Role1PermissionProvider, []string{secatest.Role1PermissionResource}, []string{secatest.Role1PermissionVerb})
	secatest.MockListRolesV1(sim, []schema.Role{
		*buildResponseRole(secatest.Role1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive),
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

	assert.Equal(t, schema.ResourceStateActive, string(*resp[0].Status.State))
}

func TestListRolesWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockListRolesV1(sim, []schema.Role{
		{
			Metadata: &schema.GlobalTenantResourceMetadata{
				Name:   secatest.Role1Name,
				Tenant: secatest.Tenant1Name,
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
			Status: &schema.Status{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := NewListOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := client.AuthorizationV1.ListRolesWithFilters(ctx, secatest.Tenant1Name, listOptions)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetRoleV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleSpec(secatest.Role1PermissionProvider, []string{secatest.Role1PermissionResource}, []string{secatest.Role1PermissionVerb})
	secatest.MockGetRoleV1(sim, buildResponseRole(secatest.Role1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive), 1)
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

	assert.Equal(t, schema.ResourceStateActive, string(*resp.Status.State))
}

func TestGetRoleUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleSpec(secatest.Role1PermissionProvider, []string{secatest.Role1PermissionResource}, []string{secatest.Role1PermissionVerb})
	secatest.MockGetRoleV1(sim, buildResponseRole(secatest.Role1Name, secatest.Tenant1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetRoleV1(sim, buildResponseRole(secatest.Role1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Role1Name}
	config := ResourceObserverConfig[schema.ResourceState]{ExpectedValue: schema.ResourceStateActive, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := client.AuthorizationV1.GetRoleUntilState(ctx, tref, config)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Role1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)

	assert.Len(t, resp.Spec.Permissions, 1)
	assert.Len(t, resp.Spec.Permissions[0].Verb, 1)
	assert.Equal(t, secatest.Role1PermissionVerb, resp.Spec.Permissions[0].Verb[0])

	assert.Equal(t, schema.ResourceStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRoleV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleSpec(secatest.Role1PermissionProvider, []string{secatest.Role1PermissionResource}, []string{secatest.Role1PermissionVerb})
	secatest.MockCreateOrUpdateRoleV1(sim, buildResponseRole(secatest.Role1Name, secatest.Tenant1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	role := schema.Role{
		Metadata: &schema.GlobalTenantResourceMetadata{
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

	assert.Equal(t, schema.ResourceStateCreating, string(*resp.Status.State))
}

func TestDeleteRoleV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)

	spec := buildResponseRoleSpec(secatest.Role1PermissionProvider, []string{secatest.Role1PermissionResource}, []string{secatest.Role1PermissionVerb})
	secatest.MockGetRoleV1(sim, buildResponseRole(secatest.Role1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive), 1)
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

func TestListRoleAssignmentsV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleAssignmentSpec([]string{secatest.Role1Name}, []string{secatest.RoleAssignment1Subject})
	secatest.MockListRoleAssignmentsV1(sim, []schema.RoleAssignment{
		*buildResponseRoleAssignment(secatest.RoleAssignment1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive),
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

	assert.Equal(t, schema.ResourceStateActive, string(*resp[0].Status.State))
}

func TestListRoleAssignmentsWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	secatest.MockListRoleAssignmentsV1(sim, []schema.RoleAssignment{
		{
			Metadata: &schema.GlobalTenantResourceMetadata{
				Name:   secatest.RoleAssignment1Name,
				Tenant: secatest.Tenant1Name,
			},
			Spec: schema.RoleAssignmentSpec{
				Subs: []string{secatest.RoleAssignment1Subject},
			},
			Status: &schema.Status{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
	})
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)
	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := NewListOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := client.AuthorizationV1.ListRoleAssignmentsWithFilters(ctx, secatest.Tenant1Name, listOptions)
	assert.NoError(t, err)
	assert.NotNil(t, iter)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.RoleAssignment1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)

	assert.Len(t, resp[0].Spec.Subs, 1)
	assert.Equal(t, secatest.RoleAssignment1Subject, resp[0].Spec.Subs[0])

	assert.Equal(t, schema.ResourceStateActive, string(*resp[0].Status.State))
}

func TestGetRoleAssignmentV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleAssignmentSpec([]string{secatest.Role1Name}, []string{secatest.RoleAssignment1Subject})
	secatest.MockGetRoleAssignmentV1(sim, buildResponseRoleAssignment(secatest.RoleAssignment1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive), 1)
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

	assert.Equal(t, schema.ResourceStateActive, string(*resp.Status.State))
}

func TestGetRoleAssignmentUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleAssignmentSpec([]string{secatest.Role1Name}, []string{secatest.RoleAssignment1Subject})
	secatest.MockGetRoleAssignmentV1(sim, buildResponseRoleAssignment(secatest.RoleAssignment1Name, secatest.Tenant1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetRoleAssignmentV1(sim, buildResponseRoleAssignment(secatest.RoleAssignment1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.RoleAssignment1Name}
	config := ResourceObserverConfig[schema.ResourceState]{ExpectedValue: schema.ResourceStateActive, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := client.AuthorizationV1.GetRoleAssignmentUntilState(ctx, tref, config)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.RoleAssignment1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)

	assert.Len(t, resp.Spec.Subs, 1)
	assert.Equal(t, secatest.RoleAssignment1Subject, resp.Spec.Subs[0])

	assert.Equal(t, schema.ResourceStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateRoleAssignmentV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleAssignmentSpec([]string{secatest.Role1Name}, []string{secatest.RoleAssignment1Subject})
	secatest.MockCreateOrUpdateRoleAssignmentV1(sim, buildResponseRoleAssignment(secatest.RoleAssignment1Name, secatest.Tenant1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureAuthorizationHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	client := newTestGlobalClientV1(t, server)

	assign := &schema.RoleAssignment{
		Metadata: &schema.GlobalTenantResourceMetadata{
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

	assert.Equal(t, schema.ResourceStateCreating, string(*resp.Status.State))
}

func TestDeleteRoleAssignmentV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	sim := mockauthorization.NewMockServerInterface(t)
	spec := buildResponseRoleAssignmentSpec([]string{secatest.Role1Name}, []string{secatest.RoleAssignment1Subject})
	secatest.MockGetRoleAssignmentV1(sim, buildResponseRoleAssignment(secatest.RoleAssignment1Name, secatest.Tenant1Name, spec, schema.ResourceStateActive), 1)
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

// Builders

func buildResponseRole(name string, tenant string, spec *schema.RoleSpec, state schema.ResourceState) *schema.Role {
	return &schema.Role{
		Metadata: secatest.NewGlobalTenantResourceMetadata(name, tenant),
		Spec:     *spec,
		Status:   secatest.NewRoleStatus(state),
	}
}

func buildResponseRoleSpec(provider string, resources []string, verbs []string) *schema.RoleSpec {
	return &schema.RoleSpec{
		Permissions: []schema.Permission{
			{
				Provider:  provider,
				Resources: resources,
				Verb:      verbs,
			},
		},
	}
}

func buildResponseRoleAssignment(name string, tenant string, spec *schema.RoleAssignmentSpec, state schema.ResourceState) *schema.RoleAssignment {
	return &schema.RoleAssignment{
		Metadata: secatest.NewGlobalTenantResourceMetadata(name, tenant),
		Spec:     *spec,
		Status:   secatest.NewRoleAssignmentStatus(state),
	}
}

func buildResponseRoleAssignmentSpec(roles []string, subs []string) *schema.RoleAssignmentSpec {
	return &schema.RoleAssignmentSpec{
		Roles: roles,
		Subs:  subs,
	}
}
