package secapi

import (
	"context"
	"net/http"

	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"k8s.io/utils/ptr"
)

// Interface

type StorageV1 interface {
	// Storage Sku
	ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.StorageSku], error)
	ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.StorageSku], error)
	GetSku(ctx context.Context, tref TenantReference) (*schema.StorageSku, error)

	// Block Storage
	ListBlockStorages(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.BlockStorage], error)
	ListBlockStoragesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.BlockStorage], error)

	GetBlockStorage(ctx context.Context, wref WorkspaceReference) (*schema.BlockStorage, error)
	GetBlockStorageUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.BlockStorage, error)

	CreateOrUpdateBlockStorageWithParams(ctx context.Context, block *schema.BlockStorage, params *storage.CreateOrUpdateBlockStorageParams) (*schema.BlockStorage, error)
	CreateOrUpdateBlockStorage(ctx context.Context, block *schema.BlockStorage) (*schema.BlockStorage, error)

	DeleteBlockStorageWithParams(ctx context.Context, block *schema.BlockStorage, params *storage.DeleteBlockStorageParams) error
	DeleteBlockStorage(ctx context.Context, block *schema.BlockStorage) error

	// Image
	ListImages(ctx context.Context, tid TenantID) (*Iterator[schema.Image], error)
	ListImagesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Image], error)

	GetImage(ctx context.Context, tref TenantReference) (*schema.Image, error)
	GetImageUntilState(ctx context.Context, tref TenantReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Image, error)

	CreateOrUpdateImageWithParams(ctx context.Context, image *schema.Image, params *storage.CreateOrUpdateImageParams) (*schema.Image, error)
	CreateOrUpdateImage(ctx context.Context, image *schema.Image) (*schema.Image, error)

	DeleteImageWithParams(ctx context.Context, image *schema.Image, params *storage.DeleteImageParams) error
	DeleteImage(ctx context.Context, image *schema.Image) error
}

// Dummy

type StorageV1Dummy struct{}

func newStorageV1Dummy() StorageV1 {
	return &StorageV1Dummy{}
}

/// Storage Sku

func (api *StorageV1Dummy) ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.StorageSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.StorageSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) GetSku(ctx context.Context, tref TenantReference) (*schema.StorageSku, error) {
	return nil, ErrProviderNotAvailable
}

/// Block Storage

func (api *StorageV1Dummy) ListBlockStorages(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.BlockStorage], error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) ListBlockStoragesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.BlockStorage], error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) GetBlockStorage(ctx context.Context, wref WorkspaceReference) (*schema.BlockStorage, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) GetBlockStorageUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.BlockStorage, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) CreateOrUpdateBlockStorageWithParams(ctx context.Context, block *schema.BlockStorage, params *storage.CreateOrUpdateBlockStorageParams) (*schema.BlockStorage, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) CreateOrUpdateBlockStorage(ctx context.Context, block *schema.BlockStorage) (*schema.BlockStorage, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) DeleteBlockStorageWithParams(ctx context.Context, block *schema.BlockStorage, params *storage.DeleteBlockStorageParams) error {
	return ErrProviderNotAvailable
}

func (api *StorageV1Dummy) DeleteBlockStorage(ctx context.Context, block *schema.BlockStorage) error {
	return ErrProviderNotAvailable
}

/// Image

func (api *StorageV1Dummy) ListImages(ctx context.Context, tid TenantID) (*Iterator[schema.Image], error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) ListImagesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Image], error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) GetImage(ctx context.Context, tref TenantReference) (*schema.Image, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) GetImageUntilState(ctx context.Context, tref TenantReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Image, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) CreateOrUpdateImageWithParams(ctx context.Context, image *schema.Image, params *storage.CreateOrUpdateImageParams) (*schema.Image, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) CreateOrUpdateImage(ctx context.Context, image *schema.Image) (*schema.Image, error) {
	return nil, ErrProviderNotAvailable
}

func (api *StorageV1Dummy) DeleteImageWithParams(ctx context.Context, image *schema.Image, params *storage.DeleteImageParams) error {
	return ErrProviderNotAvailable
}

func (api *StorageV1Dummy) DeleteImage(ctx context.Context, image *schema.Image) error {
	return ErrProviderNotAvailable
}

// Impl

type StorageV1Impl struct {
	API
	storage storage.ClientWithResponsesInterface
}

func newStorageV1Impl(client *RegionalClient, storageUrl string) (StorageV1, error) {
	storage, err := storage.NewClientWithResponses(storageUrl)
	if err != nil {
		return nil, err
	}

	return &StorageV1Impl{API: API{authToken: client.authToken}, storage: storage}, nil
}

// Storage Sku

func (api *StorageV1Impl) ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.StorageSku], error) {
	iter := Iterator[schema.StorageSku]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.StorageSku, *string, error) {
			resp, err := api.storage.ListSkusWithResponse(ctx, schema.TenantPathParam(tid), &storage.ListSkusParams{
				Accept:    ptr.To(storage.ListSkusParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *StorageV1Impl) ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.StorageSku], error) {
	iter := Iterator[schema.StorageSku]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.StorageSku, *string, error) {
			resp, err := api.storage.ListSkusWithResponse(ctx, schema.TenantPathParam(tid), &storage.ListSkusParams{
				Accept:    ptr.To(storage.ListSkusParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *StorageV1Impl) GetSku(ctx context.Context, tref TenantReference) (*schema.StorageSku, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.storage.GetSkuWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

// Block Storage

func (api *StorageV1Impl) ListBlockStorages(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.BlockStorage], error) {
	iter := Iterator[schema.BlockStorage]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.BlockStorage, *string, error) {
			resp, err := api.storage.ListBlockStoragesWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &storage.ListBlockStoragesParams{
				Accept:    ptr.To(storage.ListBlockStoragesParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *StorageV1Impl) ListBlockStoragesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.BlockStorage], error) {
	iter := Iterator[schema.BlockStorage]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.BlockStorage, *string, error) {
			resp, err := api.storage.ListBlockStoragesWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &storage.ListBlockStoragesParams{
				Accept:    ptr.To(storage.ListBlockStoragesParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *StorageV1Impl) GetBlockStorage(ctx context.Context, wref WorkspaceReference) (*schema.BlockStorage, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.storage.GetBlockStorageWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *StorageV1Impl) GetBlockStorageUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.BlockStorage, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.BlockStorage]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.BlockStorage, error) {
			resp, err := api.storage.GetBlockStorageWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *StorageV1Impl) CreateOrUpdateBlockStorageWithParams(ctx context.Context, block *schema.BlockStorage, params *storage.CreateOrUpdateBlockStorageParams) (*schema.BlockStorage, error) {
	if err := api.validateWorkspaceMetadata(block.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.storage.CreateOrUpdateBlockStorageWithResponse(ctx, block.Metadata.Tenant, block.Metadata.Workspace, block.Metadata.Name, params, *block, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *StorageV1Impl) CreateOrUpdateBlockStorage(ctx context.Context, block *schema.BlockStorage) (*schema.BlockStorage, error) {
	return api.CreateOrUpdateBlockStorageWithParams(ctx, block, nil)
}

func (api *StorageV1Impl) DeleteBlockStorageWithParams(ctx context.Context, block *schema.BlockStorage, params *storage.DeleteBlockStorageParams) error {
	if err := api.validateWorkspaceMetadata(block.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.DeleteBlockStorageWithResponse(ctx, block.Metadata.Tenant, block.Metadata.Workspace, block.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *StorageV1Impl) DeleteBlockStorage(ctx context.Context, block *schema.BlockStorage) error {
	return api.DeleteBlockStorageWithParams(ctx, block, nil)
}

// Image

func (api *StorageV1Impl) ListImages(ctx context.Context, tid TenantID) (*Iterator[schema.Image], error) {
	iter := Iterator[schema.Image]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Image, *string, error) {
			resp, err := api.storage.ListImagesWithResponse(ctx, schema.TenantPathParam(tid), &storage.ListImagesParams{
				Accept:    ptr.To(storage.ListImagesParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *StorageV1Impl) ListImagesWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.Image], error) {
	iter := Iterator[schema.Image]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Image, *string, error) {
			resp, err := api.storage.ListImagesWithResponse(ctx, schema.TenantPathParam(tid), &storage.ListImagesParams{
				Accept:    ptr.To(storage.ListImagesParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *StorageV1Impl) GetImage(ctx context.Context, tref TenantReference) (*schema.Image, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.storage.GetImageWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *StorageV1Impl) GetImageUntilState(ctx context.Context, tref TenantReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Image, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Image]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.Image, error) {
			resp, err := api.storage.GetImageWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *StorageV1Impl) CreateOrUpdateImageWithParams(ctx context.Context, image *schema.Image, params *storage.CreateOrUpdateImageParams) (*schema.Image, error) {
	if err := api.validateRegionalMetadata(image.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.storage.CreateOrUpdateImageWithResponse(ctx, image.Metadata.Tenant, image.Metadata.Name, params, *image, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *StorageV1Impl) CreateOrUpdateImage(ctx context.Context, image *schema.Image) (*schema.Image, error) {
	return api.CreateOrUpdateImageWithParams(ctx, image, nil)
}

func (api *StorageV1Impl) DeleteImageWithParams(ctx context.Context, image *schema.Image, params *storage.DeleteImageParams) error {
	if err := api.validateRegionalMetadata(image.Metadata); err != nil {
		return err
	}

	resp, err := api.storage.DeleteImageWithResponse(ctx, image.Metadata.Tenant, image.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *StorageV1Impl) DeleteImage(ctx context.Context, image *schema.Image) error {
	return api.DeleteImageWithParams(ctx, image, nil)
}

func (api *StorageV1Impl) validateRegionalMetadata(metadata *schema.RegionalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	return nil
}

func (api *StorageV1Impl) validateWorkspaceMetadata(metadata *schema.RegionalWorkspaceResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == "" {
		return ErrNoMetatadaWorkspace
	}

	return nil
}
