package secapi

import (
	"context"
	"net/http"

	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"

	"k8s.io/utils/ptr"
)

type StorageV1 struct {
	API
	storage storage.ClientWithResponsesInterface
}

// Storage Sku

func (api *StorageV1) ListSkus(ctx context.Context, tid TenantID) (*Iterator[storage.StorageSku], error) {
	iter := Iterator[storage.StorageSku]{
		fn: func(ctx context.Context, skipToken *string) ([]storage.StorageSku, *string, error) {
			resp, err := api.storage.ListSkusWithResponse(ctx, storage.Tenant(tid), &storage.ListSkusParams{
				Accept: ptr.To(storage.ListSkusParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *StorageV1) GetSku(ctx context.Context, tref TenantReference) (*storage.StorageSku, error) {
	if err := validateTenantReference(tref); err != nil {
		return nil, err
	}

	resp, err := api.storage.GetSkuWithResponse(ctx, storage.Tenant(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

// Block Storage

func (api *StorageV1) ListBlockStorages(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[storage.BlockStorage], error) {
	iter := Iterator[storage.BlockStorage]{
		fn: func(ctx context.Context, skipToken *string) ([]storage.BlockStorage, *string, error) {
			resp, err := api.storage.ListBlockStoragesWithResponse(ctx, storage.Tenant(tid), storage.Workspace(wid), &storage.ListBlockStoragesParams{
				Accept: ptr.To(storage.ListBlockStoragesParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *StorageV1) GetBlockStorage(ctx context.Context, wref WorkspaceReference) (*storage.BlockStorage, error) {
	if err := validateWorkspaceReference(wref); err != nil {
		return nil, err
	}

	resp, err := api.storage.GetBlockStorageWithResponse(ctx, storage.Tenant(wref.Tenant), storage.Workspace(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *StorageV1) CreateOrUpdateBlockStorageWithParams(ctx context.Context, wref WorkspaceReference, block *storage.BlockStorage, params *storage.CreateOrUpdateBlockStorageParams) (*storage.BlockStorage, error) {
	if err := validateWorkspaceReference(wref); err != nil {
		return nil, err
	}

	resp, err := api.storage.CreateOrUpdateBlockStorageWithResponse(ctx, storage.Tenant(wref.Tenant), storage.Workspace(wref.Workspace), wref.Name, params, *block, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *StorageV1) CreateOrUpdateBlockStorage(ctx context.Context, wref WorkspaceReference, block *storage.BlockStorage) (*storage.BlockStorage, error) {
	return api.CreateOrUpdateBlockStorageWithParams(ctx, wref, block, nil)
}

func (api *StorageV1) DeleteBlockStorageWithParams(ctx context.Context, block *storage.BlockStorage, params *storage.DeleteBlockStorageParams) error {
	if err := validateStorageZonalMetadataV1(block.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.DeleteBlockStorageWithResponse(ctx, block.Metadata.Tenant, *block.Metadata.Workspace, block.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *StorageV1) DeleteBlockStorage(ctx context.Context, block *storage.BlockStorage) error {
	return api.DeleteBlockStorageWithParams(ctx, block, nil)
}

// Image

func (api *StorageV1) ListImages(ctx context.Context, tid TenantID) (*Iterator[storage.Image], error) {
	iter := Iterator[storage.Image]{
		fn: func(ctx context.Context, skipToken *string) ([]storage.Image, *string, error) {
			resp, err := api.storage.ListImagesWithResponse(ctx, storage.Tenant(tid), &storage.ListImagesParams{
				Accept: ptr.To(storage.ListImagesParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *StorageV1) GetImage(ctx context.Context, tref TenantReference) (*storage.Image, error) {
	if err := validateTenantReference(tref); err != nil {
		return nil, err
	}

	resp, err := api.storage.GetImageWithResponse(ctx, storage.Tenant(tref.Tenant), tref.Name, api.loadRequestHeaders, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *StorageV1) CreateOrUpdateImageWithParams(ctx context.Context, tref TenantReference, image *storage.Image, params *storage.CreateOrUpdateImageParams) (*storage.Image, error) {
	if err := validateTenantReference(tref); err != nil {
		return nil, err
	}

	resp, err := api.storage.CreateOrUpdateImageWithResponse(ctx, storage.Tenant(tref.Tenant), tref.Name, params, *image, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *StorageV1) CreateOrUpdateImage(ctx context.Context, tref TenantReference, image *storage.Image) (*storage.Image, error) {
	return api.CreateOrUpdateImageWithParams(ctx, tref, image, nil)
}

func (api *StorageV1) DeleteImageWithParams(ctx context.Context, image *storage.Image, params *storage.DeleteImageParams) error {
	if err := validateStorageRegionalMetadataV1(image.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.DeleteImageWithResponse(ctx, image.Metadata.Tenant, image.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *StorageV1) DeleteImage(ctx context.Context, image *storage.Image) error {
	return api.DeleteImageWithParams(ctx, image, nil)
}

func newStorageV1(client *RegionalClient, storageUrl string) (*StorageV1, error) {
	storage, err := storage.NewClientWithResponses(storageUrl)
	if err != nil {
		return nil, err
	}

	return &StorageV1{API: API{authToken: client.authToken}, storage: storage}, nil
}

func validateStorageRegionalMetadataV1(metadata *storage.RegionalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	return nil
}

func validateStorageZonalMetadataV1(metadata *storage.ZonalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == nil {
		return ErrNoMetatadaWorkspace
	}

	return nil
}
