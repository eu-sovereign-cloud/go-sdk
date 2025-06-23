package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
)

type StorageV1 struct {
	storage storage.ClientWithResponsesInterface
}

func newStorageV1(storageUrl string) (*StorageV1, error) {
	storage, err := storage.NewClientWithResponses(storageUrl)
	if err != nil {
		return nil, err
	}

	return &StorageV1{storage: storage}, nil
}

func validateStorageMetadataV1(metadata *storage.ZonalResourceMetadata) {
	if metadata == nil {
		panic(ErrNoMetatada)
	}

	if metadata.Workspace == nil {
		panic(ErrNoMetatadaWorkspace)
	}

	if metadata.Tenant == "" {
		panic(ErrNoMetatadaTenant)
	}
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
	validateStorageMetadataV1(block.Metadata)

	resp, err := api.storage.CreateOrUpdateBlockStorageWithResponse(ctx, block.Metadata.Tenant, *block.Metadata.Workspace, block.Metadata.Name,
		&storage.CreateOrUpdateBlockStorageParams{
			IfUnmodifiedSince: &block.Metadata.ResourceVersion,
		}, *block)
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 200, 201)
	if err != nil {
		return err
	}

	return nil
}

func (api *StorageV1) DeleteBlockStorage(ctx context.Context, block *storage.BlockStorage) error {
	validateStorageMetadataV1(block.Metadata)

	resp, err := api.storage.DeleteBlockStorageWithResponse(ctx, block.Metadata.Tenant, *block.Metadata.Workspace, block.Metadata.Name, &storage.DeleteBlockStorageParams{
		IfUnmodifiedSince: &block.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 204, 404)
	if err != nil {
		return err
	}

	return nil
}
