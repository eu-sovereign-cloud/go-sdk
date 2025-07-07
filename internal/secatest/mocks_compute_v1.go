package secatest

import (
	"net/http"
	"text/template"

	mockCompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"

	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	"github.com/stretchr/testify/mock"
)

// Compute
func MockInstanceListSkusV1(sim *mockCompute.MockServerInterface, resp ListInstancesSkusResponseV1) {
	json := template.Must(template.New("response").Parse(ListInstancesSkusResponseTemplateV1))

	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, params compute.ListSkusParams) {
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

func MockListInstancesV1(sim *mockCompute.MockServerInterface, resp InstanceResponseV1) {
	json := template.Must(template.New("response").Parse(ListInstancesResponseV1))

	sim.EXPECT().ListInstances(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, params compute.ListInstancesParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockGetInstanceV1(sim *mockCompute.MockServerInterface, resp InstanceResponseV1) {
	json := template.Must(template.New("response").Parse(InstancesResponseV1))

	sim.EXPECT().GetInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)

		writeTemplateResponse(w, json, resp)
	})
}

func MockCreateOrUpdateInstanceV1(sim *mockCompute.MockServerInterface, resp InstanceResponseV1) {
	json := template.Must(template.New("response").Parse(InstancesResponseV1))

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

func MockStartInstanceV1(sim *mockCompute.MockServerInterface) {
	sim.EXPECT().StartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.StartInstanceParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockRestartInstanceV1(sim *mockCompute.MockServerInterface) {
	sim.EXPECT().RestartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.RestartInstanceParams) {
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
