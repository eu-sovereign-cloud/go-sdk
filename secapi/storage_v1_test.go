package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secalib"

	"github.com/stretchr/testify/assert"
)

// Storage Sku

func TestListStorageSkus(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	labels := schema.Labels{secatest.LabelKeyTier: secatest.StorageSku1Tier}
	spec := buildResponseStorageSkuSpec(secatest.StorageSku1Iops)
	secatest.MockListStorageSkusV1(sim, []schema.StorageSku{
		*buildResponseStorageSku(secatest.StorageSku1Name, labels, spec),
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.StorageV1.ListSkus(ctx, secatest.Tenant1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)

	assert.Equal(t, secatest.StorageSku1Name, resp[0].Metadata.Name)

	assert.Len(t, labels, 1)
	assert.Equal(t, secatest.StorageSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.StorageSku1Iops, resp[0].Spec.Iops)
}

func TestGetStorageSku(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	labels := schema.Labels{secatest.LabelKeyTier: secatest.StorageSku1Tier}
	spec := buildResponseStorageSkuSpec(secatest.StorageSku1Iops)
	secatest.MockGetStorageSkusV1(sim, buildResponseStorageSku(secatest.StorageSku1Name, labels, spec))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.StorageV1.GetSku(ctx, TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.StorageSku1Name,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.StorageSku1Name, resp.Metadata.Name)

	assert.Len(t, labels, 1)
	assert.Equal(t, secatest.StorageSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.StorageSku1Iops, resp.Spec.Iops)
}

// Block Storage

func TestListBlockStorages(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockListBlockStoragesV1(sim, []schema.BlockStorage{
		*buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef, err := BuildReferenceFromURN(secatest.StorageSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := regionalClient.StorageV1.ListBlockStorages(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.BlockStorage1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)

	assert.Equal(t, *storageSkuRef, resp[0].Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockGetBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef, err := BuildReferenceFromURN(secatest.StorageSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	wref := WorkspaceReference{
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Name:      secatest.BlockStorage1Name,
	}
	resp, err := regionalClient.StorageV1.GetBlockStorage(ctx, wref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.BlockStorage1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)

	assert.Equal(t, *storageSkuRef, resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockCreateOrUpdateBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateCreating))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef, err := BuildReferenceFromURN(secatest.StorageSku1Ref)
	if err != nil {
		t.Fatal(err)
	}

	block := &schema.BlockStorage{
		Metadata: &schema.RegionalWorkspaceResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Workspace: secatest.Workspace1Name,
			Name:      secatest.BlockStorage1Name,
		},
		Spec: schema.BlockStorageSpec{
			SkuRef: *storageSkuRef,
			SizeGB: secatest.BlockStorage1SizeGB,
		},
	}
	resp, err := regionalClient.StorageV1.CreateOrUpdateBlockStorage(ctx, block)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.BlockStorage1Name, resp.Metadata.Name)

	assert.Equal(t, *storageSkuRef, resp.Spec.SkuRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteBlockStorage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockGetBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, spec, secatest.StatusStateActive))
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
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.StorageV1.DeleteBlockStorage(ctx, resp)
	assert.NoError(t, err)
}

// Image

func TestListImages(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	blockStorageRef, err := BuildReferenceFromURN(secatest.BlockStorage1Ref)
	if err != nil {
		t.Fatal(err)
	}

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockListStorageImagesV1(sim, []schema.Image{
		*buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, spec, secatest.StatusStateActive),
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	iter, err := regionalClient.StorageV1.ListImages(ctx, secatest.Tenant1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)

	assert.Equal(t, *blockStorageRef, resp[0].Spec.BlockStorageRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))
}

func TestGetImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockGetStorageImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, spec, secatest.StatusStateActive))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef, err := BuildReferenceFromURN(secatest.BlockStorage1Ref)
	if err != nil {
		t.Fatal(err)
	}

	tref := TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Image1Name,
	}
	resp, err := regionalClient.StorageV1.GetImage(ctx, tref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)

	assert.Equal(t, *blockStorageRef, resp.Spec.BlockStorageRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp.Status.State))
}

func TestCreateOrUpdateImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockCreateOrUpdateImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, spec, secatest.StatusStateCreating))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef, err := BuildReferenceFromURN(secatest.BlockStorage1Ref)
	if err != nil {
		t.Fatal(err)
	}

	image := &schema.Image{
		Metadata: secalib.BuildResponseRegionalResourceMetadata(secatest.Image1Name, secatest.Tenant1Name),
		Spec: schema.ImageSpec{
			BlockStorageRef: *blockStorageRef,
			CpuArchitecture: secatest.Image1CpuArch,
		},
	}
	resp, err := regionalClient.StorageV1.CreateOrUpdateImage(ctx, image)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Image1Name, resp.Metadata.Name)

	assert.Equal(t, *blockStorageRef, resp.Spec.BlockStorageRef)

	assert.Equal(t, secatest.StatusStateCreating, string(*resp.Status.State))
}

func TestDeleteImage(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockGetStorageImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, spec, secatest.StatusStateActive))
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
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.StorageV1.DeleteImage(ctx, resp)
	assert.NoError(t, err)
}

func buildResponseStorageSku(name string, labels schema.Labels, spec *schema.StorageSkuSpec) *schema.StorageSku {
	return &schema.StorageSku{
		Metadata: secalib.BuildResponseSkuResourceMetadata(name),
		Labels:   labels,
		Spec:     spec,
	}
}

func buildResponseStorageSkuSpec(iops int) *schema.StorageSkuSpec {
	return &schema.StorageSkuSpec{
		Iops: iops,
	}
}

func buildResponseBlockStorage(name string, tenant string, workspace string, spec *schema.BlockStorageSpec, state string) *schema.BlockStorage {
	return &schema.BlockStorage{
		Metadata: secalib.BuildResponseRegionalWorkspaceResourceMetadata(name, tenant, workspace),
		Spec:     *spec,
		Status: &schema.BlockStorageStatus{
			State: secalib.BuildResponseResourceState(state),
		},
	}
}

func buildResponseBlockStorageSpec(t *testing.T, skuRef string, sizeGB int) *schema.BlockStorageSpec {
	urnRef, err := BuildReferenceFromURN(skuRef)
	if err != nil {
		t.Fatal(err)
	}

	return &schema.BlockStorageSpec{
		SkuRef: *urnRef,
		SizeGB: sizeGB,
	}
}

func buildResponseImage(name string, tenant string, spec *schema.ImageSpec, state string) *schema.Image {
	return &schema.Image{
		Metadata: secalib.BuildResponseRegionalResourceMetadata(name, tenant),
		Spec:     *spec,
		Status: &schema.ImageStatus{
			State: secalib.BuildResponseResourceState(state),
		},
	}
}

func buildResponseImageSpec(t *testing.T, blockStorageRef string) *schema.ImageSpec {
	urnRef, err := BuildReferenceFromURN(blockStorageRef)
	if err != nil {
		t.Fatal(err)
	}

	return &schema.ImageSpec{
		BlockStorageRef: *urnRef,
	}
}
