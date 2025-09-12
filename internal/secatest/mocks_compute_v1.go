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
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, params compute.ListSkusParams) {
			if err := configGetHttpResponse(w, instanceSkusResponseTemplateV1, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetInstanceSkuV1(sim *mockcompute.MockServerInterface, resp InstanceSkuResponseV1) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, name compute.ResourcePathParam) {
			if err := configGetHttpResponse(w, instanceSkuResponseTemplateV1, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

// Instance
func MockListInstancesV1(sim *mockcompute.MockServerInterface, resp InstanceResponseV1) {
	sim.EXPECT().ListInstances(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, workspace compute.WorkspacePathParam, params compute.ListInstancesParams) {
			if err := configGetHttpResponse(w, instancesResponseTemplateV1, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetInstanceV1(sim *mockcompute.MockServerInterface, resp InstanceResponseV1) {
	sim.EXPECT().GetInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, workspace compute.WorkspacePathParam, name compute.ResourcePathParam) {
			if err := configGetHttpResponse(w, instanceResponseTemplateV1, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockCreateOrUpdateInstanceV1(sim *mockcompute.MockServerInterface, resp InstanceResponseV1) {
	sim.EXPECT().CreateOrUpdateInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, workspace compute.WorkspacePathParam, name compute.TenantPathParam, lwp compute.CreateOrUpdateInstanceParams) {
			if err := configPutHttpResponse(w, instanceResponseTemplateV1, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().DeleteInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, workspace compute.WorkspacePathParam, name compute.ResourcePathParam, params compute.DeleteInstanceParams) {
			configDeleteHttpResponse(w)
		})
}

func MockStartInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().StartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, workspace compute.WorkspacePathParam, name compute.ResourcePathParam, params compute.StartInstanceParams) {
			configPostHttpResponse(w)
		})
}

func MockRestartInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().RestartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, workspace compute.WorkspacePathParam, name compute.ResourcePathParam, params compute.RestartInstanceParams) {
			configPostHttpResponse(w)
		})
}

func MockStopInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().StopInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant compute.TenantPathParam, workspace compute.WorkspacePathParam, name compute.ResourcePathParam, params compute.StopInstanceParams) {
			configPostHttpResponse(w)
		})
}
