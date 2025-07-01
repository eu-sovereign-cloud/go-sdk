package secatest

import (
	"bytes"
	"net/http"
	"text/template"

	mockCompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	mockRegion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	mockWorkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/stretchr/testify/mock"
)

// Region
func MockListRegionsV1(sim *mockRegion.MockServerInterface, resp ListRegionsResponseV1) {
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

func MockGetRegionV1(sim *mockRegion.MockServerInterface, resp GetRegionResponseV1) {
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

// Workspace
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

// Compute
func MockCreateOrUpdateInstanceV1(sim *mockCompute.MockServerInterface, resp CreateOrUpdateInstanceResponseV1) {
	json := template.Must(template.New("response").Parse(CreateOrUpdateInstaceResponseTemplateV1))

	sim.EXPECT().CreateOrUpdateInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.Tenant, lwp compute.CreateOrUpdateInstanceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockDeleteInstanceV1(sim *mockCompute.MockServerInterface) {

	sim.EXPECT().DeleteInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.DeleteInstanceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)

	})
}

func MockGetInstanceV1(sim *mockCompute.MockServerInterface, resp GetInstanceResponseV1) {
	json := template.Must(template.New("response").Parse(GetInstanceResponseTemplateV1))

	sim.EXPECT().GetInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetInstanceSkuV1(sim *mockCompute.MockServerInterface, resp GetInstanceSkuResponseV1) {
	json := template.Must(template.New("response").Parse(GetInstanceSkuResponseTemplateV1))

	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, name compute.ResourceName) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}
func MockListInstancesV1(sim *mockCompute.MockServerInterface, resp ListInstancesResponseV1) {
	json := template.Must(template.New("response").Parse(ListInstancesResponseTemplateV1))

	sim.EXPECT().ListInstances(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, params compute.ListInstancesParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}
func MockInstanceListSkusV1(sim *mockCompute.MockServerInterface, resp ListInstancesSkusResponseV1) {
	json := template.Must(template.New("response").Parse(ListInstancesSkusResponseTemplateV1))

	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, params compute.ListSkusParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockRestartInstanceV1(sim *mockCompute.MockServerInterface) {
	sim.EXPECT().RestartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.RestartInstanceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockStartInstanceV1(sim *mockCompute.MockServerInterface) {
	sim.EXPECT().StartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.StartInstanceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}
func MockStopInstanceV1(sim *mockCompute.MockServerInterface) {
	sim.EXPECT().StopInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.StopInstanceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}

func writeTemplateResponse(w http.ResponseWriter, tmpl *template.Template, data any) {
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, data)
	_, _ = w.Write(buffer.Bytes())
}
