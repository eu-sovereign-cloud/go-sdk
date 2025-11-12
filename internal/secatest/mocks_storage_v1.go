package secatest

import (
	"net/http"

	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/mock"
)

// Storage Sku
func MockListStorageSkusV1(sim *mockstorage.MockServerInterface, resp []schema.StorageSku) {
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, params storage.ListSkusParams) {
			iter := storage.SkuIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetStorageSkusV1(sim *mockstorage.MockServerInterface, resp *schema.StorageSku) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, name schema.ResourcePathParam) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

// Block Storage
func MockListBlockStoragesV1(sim *mockstorage.MockServerInterface, resp []schema.BlockStorage) {
	sim.EXPECT().ListBlockStorages(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, params storage.ListBlockStoragesParams) {
			iter := storage.BlockStorageIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetBlockStorageV1(sim *mockstorage.MockServerInterface, resp *schema.BlockStorage, times int) {
	sim.EXPECT().GetBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}).Times(times)
}

func MockCreateOrUpdateBlockStorageV1(sim *mockstorage.MockServerInterface, resp *schema.BlockStorage) {
	sim.EXPECT().CreateOrUpdateBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam, params storage.CreateOrUpdateBlockStorageParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteBlockStorageV1(sim *mockstorage.MockServerInterface) {
	sim.EXPECT().DeleteBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, workspace schema.WorkspacePathParam, name schema.ResourcePathParam, params storage.DeleteBlockStorageParams) {
			configDeleteHttpResponse(w)
		})
}

// Image
func MockListStorageImagesV1(sim *mockstorage.MockServerInterface, resp []schema.Image) {
	sim.EXPECT().ListImages(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, params storage.ListImagesParams) {
			iter := storage.ImageIterator{Items: resp}
			if err := configGetHttpResponse(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetStorageImageV1(sim *mockstorage.MockServerInterface, resp *schema.Image, times int) {
	sim.EXPECT().GetImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, name schema.ResourcePathParam) {
			if err := configGetHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}).Times(times)
}

func MockCreateOrUpdateImageV1(sim *mockstorage.MockServerInterface, resp *schema.Image) {
	sim.EXPECT().CreateOrUpdateImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, name schema.ResourcePathParam, params storage.CreateOrUpdateImageParams) {
			if err := configPutHttpResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockDeleteImageV1(sim *mockstorage.MockServerInterface) {
	sim.EXPECT().DeleteImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant schema.TenantPathParam, name schema.ResourcePathParam, params storage.DeleteImageParams) {
			configDeleteHttpResponse(w)
		})
}
