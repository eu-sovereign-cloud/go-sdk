package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
)

type StorageV1 struct {
	storage storage.ClientWithResponsesInterface
	crud    CRUD[
		storage.BlockStorage,
		storage.Tenant,
		storage.Workspace,
		storage.ListBlockStoragesParams,
		storage.RequestEditorFn,
		storage.ListBlockStoragesResponse,
		storage.ResourceName,
		storage.GetBlockStorageResponse,
	]
}

func newStorageV1(storageUrl string) (*StorageV1, error) {
	s, err := storage.NewClientWithResponses(storageUrl)
	if err != nil {
		return nil, err
	}

	return &StorageV1{
		storage: s,
		crud: CRUD[storage.BlockStorage, storage.Tenant, storage.Workspace, storage.ListBlockStoragesParams, storage.RequestEditorFn, storage.ListBlockStoragesResponse, storage.ResourceName, storage.GetBlockStorageResponse]{
			ReturnFn: func(resp *storage.ListBlockStoragesResponse) ([]storage.BlockStorage, *string, error) {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			},
			ReturnSingleFn: func(resp *storage.GetBlockStorageResponse) (*storage.BlockStorage, error) {
				return resp.JSON200, nil
			},
			ParamsFn: func() *storage.ListBlockStoragesParams {
				return &storage.ListBlockStoragesParams{
					Accept: ptr.To(storage.ListBlockStoragesParamsAcceptApplicationjson),
				}
			},
			ListFn: s.ListBlockStoragesWithResponse,
			GetFn:  s.GetBlockStorageWithResponse,
		},
	}, nil
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
	return api.crud.List(ctx, tid, wid)
}

func (api *StorageV1) GetBlockStorage(ctx context.Context, wref WorkspaceReference) (*storage.BlockStorage, error) {
	return api.crud.Get(ctx, wref)
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

	if err = checkStatusCode(resp, 200, 201); err != nil {
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

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}
