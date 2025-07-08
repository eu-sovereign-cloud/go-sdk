package secatest

import (
	"net/http"

	mockcompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"

	"github.com/stretchr/testify/mock"
)

// Instance Sku
func MockListInstanceSkusV1(sim *mockcompute.MockServerInterface, resp InstanceSkuResponseV1) {
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, params compute.ListSkusParams) {
			configGetHttpResponse(w, instanceSkusResponseTemplateV1, resp)
		})
}
func MockGetInstanceSkuV1(sim *mockcompute.MockServerInterface, resp InstanceSkuResponseV1) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, name compute.ResourceName) {
			configGetHttpResponse(w, instanceSkuResponseTemplateV1, resp)
		})
}

// Instance
func MockListInstancesV1(sim *mockcompute.MockServerInterface, resp InstanceResponseV1) {
	sim.EXPECT().ListInstances(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, params compute.ListInstancesParams) {
			configGetHttpResponse(w, instancesResponseTemplateV1, resp)
		})
}
func MockGetInstanceV1(sim *mockcompute.MockServerInterface, resp InstanceResponseV1) {
	sim.EXPECT().GetInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName) {
			configGetHttpResponse(w, instanceResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateInstanceV1(sim *mockcompute.MockServerInterface, resp InstanceResponseV1) {
	sim.EXPECT().CreateOrUpdateInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.Tenant, lwp compute.CreateOrUpdateInstanceParams) {
			configPutHttpResponse(w, instanceResponseTemplateV1, resp)
		})
}
func MockDeleteInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().DeleteInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.DeleteInstanceParams) {
			configDeleteHttpResponse(w)
		})
}
func MockStartInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().StartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.StartInstanceParams) {
			configPostHttpResponse(w)
		})
}
func MockRestartInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().RestartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.RestartInstanceParams) {
			configPostHttpResponse(w)
		})
}
func MockStopInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().StopInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.Tenant, workspace compute.Workspace, name compute.ResourceName, params compute.StopInstanceParams) {
			configPostHttpResponse(w)
		})
}
