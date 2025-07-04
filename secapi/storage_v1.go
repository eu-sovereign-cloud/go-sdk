package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
)

type StorageV1 struct {
	storage storage.ClientWithResponsesInterface
}

func (api *StorageV1) ListSkus(ctx context.Context, tid TenantID) (*Iterator[storage.StorageSku], error) {
	iter := Iterator[storage.StorageSku]{
		fn: func(ctx context.Context, skipToken *string) ([]storage.StorageSku, *string, error) {
			resp, err := api.storage.ListSkusWithResponse(ctx, storage.Tenant(tid), &storage.ListSkusParams{
				Accept: ptr.To(storage.ListSkusParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *StorageV1) GetSku(ctx context.Context, tref TenantReference) (*storage.StorageSku, error) {
	resp, err := api.storage.GetSkuWithResponse(ctx, storage.Tenant(tref.Tenant), tref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *StorageV1) ListBlockStorages(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[storage.BlockStorage], error) {
	iter := Iterator[storage.BlockStorage]{
		fn: func(ctx context.Context, skipToken *string) ([]storage.BlockStorage, *string, error) {
			resp, err := api.storage.ListBlockStoragesWithResponse(ctx, storage.Tenant(tid), storage.Workspace(wid), &storage.ListBlockStoragesParams{
				Accept: ptr.To(storage.ListBlockStoragesParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *StorageV1) GetBlockStorage(ctx context.Context, wref WorkspaceReference) (*storage.BlockStorage, error) {
	resp, err := api.storage.GetBlockStorageWithResponse(ctx, storage.Tenant(wref.Tenant), storage.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *StorageV1) CreateOrUpdateBlockStorage(ctx context.Context, block *storage.BlockStorage) error {
	if err := validateStorageZonalMetadataV1(block.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.CreateOrUpdateBlockStorageWithResponse(ctx, block.Metadata.Tenant, *block.Metadata.Workspace, block.Metadata.Name,
		&storage.CreateOrUpdateBlockStorageParams{
			IfUnmodifiedSince: &block.Metadata.ResourceVersion,
		}, *block)
	if err != nil {
		return err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *StorageV1) DeleteBlockStorage(ctx context.Context, block *storage.BlockStorage) error {
	if err := validateStorageZonalMetadataV1(block.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.DeleteBlockStorageWithResponse(ctx, block.Metadata.Tenant, *block.Metadata.Workspace, block.Metadata.Name, &storage.DeleteBlockStorageParams{
		IfUnmodifiedSince: &block.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *StorageV1) ListImages(ctx context.Context, tid TenantID) (*Iterator[storage.Image], error) {
	iter := Iterator[storage.Image]{
		fn: func(ctx context.Context, skipToken *string) ([]storage.Image, *string, error) {
			resp, err := api.storage.ListImagesWithResponse(ctx, storage.Tenant(tid), &storage.ListImagesParams{
				Accept: ptr.To(storage.ListImagesParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *StorageV1) GetImage(ctx context.Context, tref TenantReference) (*storage.Image, error) {
	resp, err := api.storage.GetImageWithResponse(ctx, storage.Tenant(tref.Tenant), tref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *StorageV1) CreateOrUpdateImage(ctx context.Context, image *storage.Image) error {
	if err := validateStorageRegionalMetadataV1(image.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.CreateOrUpdateImageWithResponse(ctx, image.Metadata.Tenant, image.Metadata.Name,
		&storage.CreateOrUpdateImageParams{
			IfUnmodifiedSince: &image.Metadata.ResourceVersion,
		}, *image)
	if err != nil {
		return err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *StorageV1) DeleteImage(ctx context.Context, image *storage.Image) error {
	if err := validateStorageRegionalMetadataV1(image.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.DeleteImageWithResponse(ctx, image.Metadata.Tenant, image.Metadata.Name, &storage.DeleteImageParams{
		IfUnmodifiedSince: &image.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func newStorageV1(storageUrl string) (*StorageV1, error) {
	storage, err := storage.NewClientWithResponses(storageUrl)
	if err != nil {
		return nil, err
	}

	return &StorageV1{storage: storage}, nil
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
