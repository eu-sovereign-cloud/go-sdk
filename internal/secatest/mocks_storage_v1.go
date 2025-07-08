package secatest

import (
	"net/http"

	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"

	"github.com/stretchr/testify/mock"
)

// Storage Sku
func MockListStorageSkusV1(sim *mockstorage.MockServerInterface, resp StorageSkuResponseV1) {
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, params storage.ListSkusParams) {
			configGetHttpResponse(w, storageSkusResponseTemplateV1, resp)
		})
}
func MockGetStorageSkusV1(sim *mockstorage.MockServerInterface, resp StorageSkuResponseV1) {
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName) {
			configGetHttpResponse(w, storageSkuResponseTemplateV1, resp)
		})
}

// Block Storage
func MockListBlockStoragesV1(sim *mockstorage.MockServerInterface, resp BlockStorageResponseV1) {
	sim.EXPECT().ListBlockStorages(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, params storage.ListBlockStoragesParams) {
			configGetHttpResponse(w, blockStoragesResponseTemplateV1, resp)
		})
}
func MockGetBlockStorageV1(sim *mockstorage.MockServerInterface, resp BlockStorageResponseV1) {
	sim.EXPECT().GetBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, name storage.ResourceName) {
			configGetHttpResponse(w, blockStorageResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateBlockStorageV1(sim *mockstorage.MockServerInterface, resp BlockStorageResponseV1) {
	sim.EXPECT().CreateOrUpdateBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, name storage.ResourceName, params storage.CreateOrUpdateBlockStorageParams) {
			configPutHttpResponse(w, blockStorageResponseTemplateV1, resp)
		})
}
func MockDeleteBlockStorageV1(sim *mockstorage.MockServerInterface) {
	sim.EXPECT().DeleteBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, name storage.ResourceName, params storage.DeleteBlockStorageParams) {
			configDeleteHttpResponse(w)
		})
}

// Image
func MockListStorageImagesV1(sim *mockstorage.MockServerInterface, resp ImageResponseV1) {
	sim.EXPECT().ListImages(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, params storage.ListImagesParams) {
			configGetHttpResponse(w, imagesResponseTemplateV1, resp)
		})
}
func MockGetStorageImageV1(sim *mockstorage.MockServerInterface, resp ImageResponseV1) {
	sim.EXPECT().GetImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName) {
			configGetHttpResponse(w, imageResponseTemplateV1, resp)
		})
}
func MockCreateOrUpdateImageV1(sim *mockstorage.MockServerInterface, resp ImageResponseV1) {
	sim.EXPECT().CreateOrUpdateImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName, params storage.CreateOrUpdateImageParams) {
			configPutHttpResponse(w, imageResponseTemplateV1, resp)
		})
}
func MockDeleteImageV1(sim *mockstorage.MockServerInterface) {
	sim.EXPECT().DeleteImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName, params storage.DeleteImageParams) {
			configDeleteHttpResponse(w)
		})
}
