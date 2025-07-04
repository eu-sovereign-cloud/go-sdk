package secatest

import (
	"net/http"
	"text/template"

	mockStorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"

	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
	"github.com/stretchr/testify/mock"
)

// Storage
func MockStorageListSkusV1(sim *mockStorage.MockServerInterface, resp ListStorageSkusResponseV1) {
	json := template.Must(template.New("response").Parse(ListStorageSkusResponseTemplateV1))
	sim.EXPECT().ListSkus(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, params storage.ListSkusParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockGetStorageSkusV1(sim *mockStorage.MockServerInterface, resp NameResponseV1) {
	json := template.Must(template.New("response").Parse(GetStorageSkuResponseTemplateV1))
	sim.EXPECT().GetSku(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockListBlockStoragesV1(sim *mockStorage.MockServerInterface, resp NameResponseV1) {
	json := template.Must(template.New("response").Parse(ListBlockStorageResponseTemplateV1))
	sim.EXPECT().ListBlockStorages(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, params storage.ListBlockStoragesParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockGetBlockStorageV1(sim *mockStorage.MockServerInterface, resp NameResponseV1) {
	json := template.Must(template.New("response").Parse(GetBlockStorageResponseTemplateV1))
	sim.EXPECT().GetBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, name storage.ResourceName) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockCreateOrUpdateBlockStorageV1(sim *mockStorage.MockServerInterface, resp NameResponseV1) {
	json := template.Must(template.New("response").Parse(CreateOrUpdateBlockStorageResponseTemplateV1))
	sim.EXPECT().CreateOrUpdateBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, name storage.ResourceName, params storage.CreateOrUpdateBlockStorageParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockDeleteBlockStorageV1(sim *mockStorage.MockServerInterface) {
	sim.EXPECT().DeleteBlockStorage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, workspace storage.Workspace, name storage.ResourceName, params storage.DeleteBlockStorageParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockListStorageImagesV1(sim *mockStorage.MockServerInterface, resp NameResponseV1) {
	json := template.Must(template.New("response").Parse(ListBlockStorageResponseTemplateV1))
	sim.EXPECT().ListImages(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, params storage.ListImagesParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockGetStorageImageV1(sim *mockStorage.MockServerInterface, resp NameResponseV1) {
	json := template.Must(template.New("response").Parse(GetStorageImageResponseTemplateV1))
	sim.EXPECT().GetImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockCreateOrUpdateImageV1(sim *mockStorage.MockServerInterface, resp NameResponseV1) {
	json := template.Must(template.New("response").Parse(CreateOrUpdateImageResponseTemplateV1))
	sim.EXPECT().CreateOrUpdateImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName, params storage.CreateOrUpdateImageParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		writeTemplateResponse(w, json, resp)
	})
}

func MockDeleteImageV1(sim *mockStorage.MockServerInterface) {
	sim.EXPECT().DeleteImage(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, tenant storage.Tenant, name storage.ResourceName, params storage.DeleteImageParams) {
		w.Header().Set(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusAccepted)
	})
}
