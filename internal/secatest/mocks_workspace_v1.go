package secatest

import (
	"net/http"

	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"

	"github.com/stretchr/testify/mock"
)

// Workspace
func MockListWorkspaceV1(sim *mockworkspace.MockServerInterface, resp WorkspaceTypeResponseV1) {
	sim.EXPECT().ListWorkspaces(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, lwp workspace.ListWorkspacesParams) {
			configGetHttpResponse(w, workspacesResponseTemplateV1, resp)
		})
}
func MockGetWorkspaceV1(sim *mockworkspace.MockServerInterface, resp WorkspaceTypeResponseV1) {
	sim.EXPECT().GetWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
			configGetHttpResponse(w, workspaceResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateWorkspaceV1(sim *mockworkspace.MockServerInterface, resp WorkspaceTypeResponseV1) {
	sim.EXPECT().CreateOrUpdateWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params workspace.CreateOrUpdateWorkspaceParams) {
			configPutHttpResponse(w, workspaceResponseTemplateV1, resp)
		})
}
func MockDeleteWorkspaceV1(sim *mockworkspace.MockServerInterface) {
	sim.EXPECT().DeleteWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params workspace.DeleteWorkspaceParams) {
			configDeleteHttpResponse(w)
		})
}
