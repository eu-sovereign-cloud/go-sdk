package secatest

import (
	"net/http"

	mockauthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"
	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/mock"
)

// Role
func MockListRolesV1(sim *mockauthorization.MockServerInterface, resp []schema.Role) {
	sim.EXPECT().ListRoles(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, params authorization.ListRolesParams) {
			iter := &authorization.RoleIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetRoleV1(sim *mockauthorization.MockServerInterface, resp *schema.Role) {
	sim.EXPECT().GetRole(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateRoleV1(sim *mockauthorization.MockServerInterface, resp *schema.Role) {
	sim.EXPECT().CreateOrUpdateRole(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.CreateOrUpdateRoleParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteRoleV1(sim *mockauthorization.MockServerInterface) {
	sim.EXPECT().DeleteRole(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.DeleteRoleParams) {
			configDeleteHttpResponse(w)
		})
}

// Role Assignment
func MockListRoleAssignmentsV1(sim *mockauthorization.MockServerInterface, resp []schema.RoleAssignment) {
	sim.EXPECT().ListRoleAssignments(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, params authorization.ListRoleAssignmentsParams) {
			iter := &authorization.RoleAssignmentIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetRoleAssignmentV1(sim *mockauthorization.MockServerInterface, resp *schema.RoleAssignment) {
	sim.EXPECT().GetRoleAssignment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateRoleAssignmentV1(sim *mockauthorization.MockServerInterface, resp *schema.RoleAssignment) {
	sim.EXPECT().CreateOrUpdateRoleAssignment(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.CreateOrUpdateRoleAssignmentParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteRoleAssignmentV1(sim *mockauthorization.MockServerInterface) {
	sim.EXPECT().DeleteRoleAssignment(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.DeleteRoleAssignmentParams) {
			configDeleteHttpResponse(w)
		})
}
