package secatest

import (
	"net/http"
	"text/template"

	mockWorkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/stretchr/testify/mock"
)

// Workspace
func MockListWorkspaceV1(sim *mockWorkspace.MockServerInterface, resp WorkspaceTypeResponseV1) {
	json := template.Must(template.New("response").Parse(ItemsWorkspaceResponseV1))

	sim.EXPECT().ListWorkspaces(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, lwp workspace.ListWorkspacesParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetWorkspaceV1(sim *mockWorkspace.MockServerInterface, resp WorkspaceTypeResponseV1) {
	json := template.Must(template.New("response").Parse(WorkspaceResponseV1))

	sim.EXPECT().GetWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockCreateOrUpdateWorkspaceV1(sim *mockWorkspace.MockServerInterface, resp WorkspaceTypeResponseV1) {
	json := template.Must(template.New("response").Parse(WorkspaceResponseV1))

	sim.EXPECT().CreateOrUpdateWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params workspace.CreateOrUpdateWorkspaceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockDeleteWorkspaceV1(sim *mockWorkspace.MockServerInterface) {

	sim.EXPECT().DeleteWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string, params workspace.DeleteWorkspaceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)

	})
}

/*
func writeTemplateResponse(w http.ResponseWriter, tmpl *template.Template, data any) {
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, data)
	_, _ = w.Write(buffer.Bytes())
}
*/
