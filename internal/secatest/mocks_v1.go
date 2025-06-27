package secatest

import (
	"bytes"
	"net/http"
	"text/template"

	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	mockWorkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/stretchr/testify/mock"
)

func writeTemplateResponse(w http.ResponseWriter, tmpl *template.Template, data any) {
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, data)
	_, _ = w.Write(buffer.Bytes())
}

func MockListRegionsV1(sim *mockregion.MockServerInterface, resp ListRegionsResponseV1) {
	json := template.Must(template.New("response").Parse(ListRegionsResponseTemplateV1))

	sim.EXPECT().ListRegions(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, lrp region.ListRegionsParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		for i := range resp.Providers {
			resp.Providers[i].URL = "http://" + r.Host + resp.Providers[i].URL
		}

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetRegionV1(sim *mockregion.MockServerInterface, resp GetRegionResponseV1) {
	json := template.Must(template.New("response").Parse(GetRegionResponseTemplateV1))

	sim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, name string) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		for i := range resp.Providers {
			resp.Providers[i].URL = "http://" + r.Host + resp.Providers[i].URL
		}

		writeTemplateResponse(w, json, resp)
	})
}

func MockListWorkspaceV1(sim *mockWorkspace.MockServerInterface, resp ListWorkspaceResponseV1) {
	json := template.Must(template.New("response").Parse(ListWorkspaceResponseTemplateV1))

	sim.EXPECT().ListWorkspaces(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, lwp workspace.ListWorkspacesParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetWorkspaceV1(sim *mockWorkspace.MockServerInterface, resp GetWorkspaceResponseV1) {
	json := template.Must(template.New("response").Parse(GetWorkspaceResponseTemplateV1))

	sim.EXPECT().GetWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant string, name string) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockCreateOrUpdateWorkspaceV1(sim *mockWorkspace.MockServerInterface, resp CreateOrUpdateWorkspaceResponseV1) {
	json := template.Must(template.New("response").Parse(CreateOrUpdateWorkspaceResponseTemplateV1))

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
