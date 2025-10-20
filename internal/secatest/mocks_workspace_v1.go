package secatest

import (
	"net/http"

	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/mock"
)

// Workspace
func MockListWorkspaceV1(sim *mockworkspace.MockServerInterface, resp []schema.Workspace) {
	sim.EXPECT().ListWorkspaces(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, lwp workspace.ListWorkspacesParams) {
			iter := workspace.WorkspaceIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetWorkspaceV1(sim *mockworkspace.MockServerInterface, resp *schema.Workspace) {
	sim.EXPECT().GetWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateWorkspaceV1(sim *mockworkspace.MockServerInterface, resp *schema.Workspace) {
	sim.EXPECT().CreateOrUpdateWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params workspace.CreateOrUpdateWorkspaceParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteWorkspaceV1(sim *mockworkspace.MockServerInterface) {
	sim.EXPECT().DeleteWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params workspace.DeleteWorkspaceParams) {
			configDeleteHttpResponse(w)
		})
}
