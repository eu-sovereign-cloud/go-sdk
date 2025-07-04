package secatest

import (
	"net/http"
	"text/template"

	mockAuthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"

	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
	"github.com/stretchr/testify/mock"
)

func MockListRolesV1(sim *mockAuthorization.MockServerInterface, resp NameAndTenantResponseV1) {
	json := template.Must(template.New("response").Parse(RolesResponseV1))

	sim.EXPECT().ListRoles(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, params authorization.ListRolesParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetRoleV1(sim *mockAuthorization.MockServerInterface, resp NameAndTenantResponseV1) {
	json := template.Must(template.New("response").Parse(RolesResponseV1))

	sim.EXPECT().GetRole(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockCreateOrUpdateRoleV1(sim *mockAuthorization.MockServerInterface, resp NameAndTenantResponseV1) {
	json := template.Must(template.New("response").Parse(RolesResponseV1))

	sim.EXPECT().CreateOrUpdateRole(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.CreateOrUpdateRoleParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockDeleteRoleV1(sim *mockAuthorization.MockServerInterface) {

	sim.EXPECT().DeleteRole(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.DeleteRoleParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockListRoleAssignmentsV1(sim *mockAuthorization.MockServerInterface, resp RoleAssignmentResponseV1) {
	json := template.Must(template.New("response").Parse(RolesAssignmentsResponseV1))

	sim.EXPECT().ListRoleAssignments(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, params authorization.ListRoleAssignmentsParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetRoleAssignmentV1(sim *mockAuthorization.MockServerInterface, resp RoleAssignmentResponseV1) {
	json := template.Must(template.New("response").Parse(RolesAssignmentsResponseV1))

	sim.EXPECT().GetRoleAssignment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockCreateOrUpdateRoleAssignmentV1(sim *mockAuthorization.MockServerInterface, resp RoleAssignmentResponseV1) {
	json := template.Must(template.New("response").Parse(RolesAssignmentsResponseV1))

	sim.EXPECT().CreateOrUpdateRoleAssignment(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.CreateOrUpdateRoleAssignmentParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockDeleteRoleAssignmentV1(sim *mockAuthorization.MockServerInterface) {

	sim.EXPECT().DeleteRoleAssignment(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params authorization.DeleteRoleAssignmentParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}
