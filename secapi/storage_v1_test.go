package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

// Storage Sku

func TestListStorageSkus(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListStorageSkusV1(sim, secatest.StorageSkuResponseV1{})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	sgIter, err := regionalClient.StorageV1.ListSkus(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	sg, err := sgIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, sg, 1)

	require.NotEmpty(t, sg[0].Labels)
	require.NotEmpty(t, sg[0].Spec.Iops)
	require.NotEmpty(t, sg[0].Spec.MinVolumeSize)
}

func TestGetStorageSku(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetStorageSkusV1(sim, secatest.StorageSkuResponseV1{})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	sg, err := regionalClient.StorageV1.GetSku(ctx, TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.StorageSku1Name,
	})
	require.NoError(t, err)
	require.NotEmpty(t, sg)

	assert.Equal(t, secatest.StorageSku1Name, sg.Metadata.Name)
}

// Block Storage

func TestListBlockStorages(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListBlockStoragesV1(sim, secatest.BlockStorageResponseV1{})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	sgIter, err := regionalClient.StorageV1.ListBlockStorages(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	sg, err := sgIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, sg, 1)

	assert.Equal(t, secatest.Workspace1Name, sg[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, sg[0].Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, sg[0].Metadata.Region)
	assert.Equal(t, secatest.ZoneA, sg[0].Metadata.Zone)
}

func TestGetBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetBlockStorageV1(sim, secatest.BlockStorageResponseV1{})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.Storage1Name,
	}
	sg, err := regionalClient.StorageV1.GetBlockStorage(ctx, wref)
	require.NoError(t, err)
	require.NotEmpty(t, sg)

	assert.Equal(t, secatest.Storage1Name, sg.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, sg.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, sg.Metadata.Region)
	assert.Equal(t, secatest.ZoneA, sg.Metadata.Zone)
}

func TestCreateOrUpdateBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateBlockStorageV1(sim, secatest.BlockStorageResponseV1{})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	block := &storage.BlockStorage{
		Metadata: &storage.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Workspace1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.StorageV1.CreateOrUpdateBlockStorage(ctx, block)
	require.NoError(t, err)
}

func TestDeleteBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockDeleteBlockStorageV1(sim)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	block := &storage.BlockStorage{
		Metadata: &storage.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Workspace1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
	}
	err := regionalClient.StorageV1.DeleteBlockStorage(ctx, block)
	require.NoError(t, err)
}

// Image

func TestListImages(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListStorageImagesV1(sim, secatest.ImageResponseV1{})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	imgIter, err := regionalClient.StorageV1.ListImages(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	images, err := imgIter.All(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, images)

	assert.Equal(t, secatest.Image1Name, images[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, images[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *images[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, images[0].Metadata.Region)
}

func TestGetImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetStorageImageV1(sim, secatest.ImageResponseV1{
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	tref := TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Image1Name,
	}
	image, err := regionalClient.StorageV1.GetImage(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, image)

	assert.Equal(t, secatest.Image1Name, image.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, image.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *image.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, image.Metadata.Region)
}

func TestCreateOrUpdateImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateImageV1(sim, secatest.ImageResponseV1{
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	image := &storage.Image{
		Metadata: &storage.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Image1Name,
		},
	}
	err := regionalClient.StorageV1.CreateOrUpdateImage(ctx, image)
	require.NoError(t, err)
}

func TestDeleteImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockDeleteImageV1(sim)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := getTestRegionalClient(t, ctx, []RegionalAPI{NetworkV1API}, server)

	image := &storage.Image{
		Metadata: &storage.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Image1Name,
		},
	}
	err := regionalClient.StorageV1.DeleteImage(ctx, image)
	require.NoError(t, err)
}
