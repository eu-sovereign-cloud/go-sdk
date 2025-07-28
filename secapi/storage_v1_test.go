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
	secatest.MockListStorageSkusV1(sim, secatest.StorageSkuResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.StorageSku1Name},
		Tier:     secatest.StorageSku1Tier,
		Iops:     secatest.StorageSku1Iops,
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.StorageV1.ListSkus(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)

	assert.Equal(t, secatest.StorageSku1Name, resp[0].Metadata.Name)

	labels := *resp[0].Labels
	require.Len(t, labels, 1)
	assert.Equal(t, secatest.StorageSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.StorageSku1Iops, resp[0].Spec.Iops)
}

func TestGetStorageSku(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetStorageSkusV1(sim, secatest.StorageSkuResponseV1{
		Metadata: secatest.MetadataResponseV1{Name: secatest.StorageSku1Name},
		Tier:     secatest.StorageSku1Tier,
		Iops:     secatest.StorageSku1Iops,
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.StorageV1.GetSku(ctx, TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.StorageSku1Name,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.StorageSku1Name, resp.Metadata.Name)

	labels := *resp.Labels
	require.Len(t, labels, 1)
	assert.Equal(t, secatest.StorageSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.StorageSku1Iops, resp.Spec.Iops)
}

// Block Storage

func TestListBlockStorages(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListBlockStoragesV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.StorageV1.ListBlockStorages(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.Len(t, resp, 1)

	assert.Equal(t, secatest.BlockStorage1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.StorageSku1Ref, resp[0].Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetBlockStorageV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.BlockStorage1Name,
	}
	resp, err := regionalClient.StorageV1.GetBlockStorage(ctx, wref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.BlockStorage1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.StorageSku1Ref, resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateBlockStorageV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	block := &storage.BlockStorage{
		Metadata: &storage.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Workspace1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		Spec: storage.BlockStorageSpec{
			SkuRef: secatest.StorageSku1Ref,
			SizeGB: secatest.BlockStorage1SizeGB,
		},
	}
	resp, err := regionalClient.StorageV1.CreateOrUpdateBlockStorage(ctx, block)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.BlockStorage1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.StorageSku1Ref, resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetBlockStorageV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteBlockStorageV1(sim)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.BlockStorage1Name,
	}
	resp, err := regionalClient.StorageV1.GetBlockStorage(ctx, wref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.StorageV1.DeleteBlockStorage(ctx, resp)
	require.NoError(t, err)
}

// Image

func TestListImages(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListStorageImagesV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Image1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.StorageV1.ListImages(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	resp, err := iter.All(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp[0].Metadata.Workspace)

	assert.Equal(t, secatest.BlockStorage1Ref, resp[0].Spec.BlockStorageRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetStorageImageV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Image1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	tref := TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Image1Name,
	}
	resp, err := regionalClient.StorageV1.GetImage(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.BlockStorage1Ref, resp.Spec.BlockStorageRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateImageV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Image1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	image := &storage.Image{
		Metadata: &storage.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Image1Name,
		},
		Spec: storage.ImageSpec{
			BlockStorageRef: secatest.BlockStorage1Ref,
			CpuArchitecture: secatest.Image1CpuArch,
		},
	}
	resp, err := regionalClient.StorageV1.CreateOrUpdateImage(ctx, image)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, *resp.Metadata.Workspace)

	assert.Equal(t, secatest.BlockStorage1Ref, resp.Spec.BlockStorageRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetStorageImageV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Image1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.MockDeleteImageV1(sim)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	tref := TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Image1Name,
	}
	resp, err := regionalClient.StorageV1.GetImage(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = regionalClient.StorageV1.DeleteImage(ctx, resp)
	require.NoError(t, err)
}
