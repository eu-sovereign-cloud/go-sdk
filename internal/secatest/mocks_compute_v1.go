package secatest

import (
	"net/http"

	mockcompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/mock"
)

// Instance Sku
func MockListInstanceSkusV1(sim *mockcompute.MockServerInterface, resp []schema.InstanceSku) {
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, params compute.ListSkusParams) {
			iter := compute.SkuIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetInstanceSkuV1(sim *mockcompute.MockServerInterface, resp *schema.InstanceSku) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, name schema.ResourcePathParam) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

// Instance
func MockListInstancesV1(sim *mockcompute.MockServerInterface, resp []schema.Instance) {
	sim.EXPECT().ListInstances(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, params compute.ListInstancesParams) {
			iter := compute.InstanceIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetInstanceV1(sim *mockcompute.MockServerInterface, resp *schema.Instance, times int) {
	sim.EXPECT().GetInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}).Times(times)
}

func MockCreateOrUpdateInstanceV1(sim *mockcompute.MockServerInterface, resp *schema.Instance) {
	sim.EXPECT().CreateOrUpdateInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.TenantPathParam, lwp compute.CreateOrUpdateInstanceParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().DeleteInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam, params compute.DeleteInstanceParams) {
			configDeleteHttpResponse(w)
		})
}

func MockStartInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().StartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam, params compute.StartInstanceParams) {
			configPostHttpResponse(w)
		})
}

func MockRestartInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().RestartInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam, params compute.RestartInstanceParams) {
			configPostHttpResponse(w)
		})
}

func MockStopInstanceV1(sim *mockcompute.MockServerInterface) {
	sim.EXPECT().StopInstance(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam, params compute.StopInstanceParams) {
			configPostHttpResponse(w)
		})
}
