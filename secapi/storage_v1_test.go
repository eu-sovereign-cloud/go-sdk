package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi/builders"
	"k8s.io/utils/ptr"

	"github.com/stretchr/testify/assert"
)

// Storage Sku

func TestListStorageSkusV1(t *testing.T) {
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
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)

	assert.Equal(t, secatest.StorageSku1Name, resp[0].Metadata.Name)

	labels := resp[0].Labels
	assert.Len(t, labels, 1)
	assert.Equal(t, secatest.StorageSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.StorageSku1Iops, resp[0].Spec.Iops)
}

func TestListStorageSkusWithFiltersV1(t *testing.T) {
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
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)

	assert.Equal(t, secatest.StorageSku1Name, resp[0].Metadata.Name)

	labels := resp[0].Labels
	assert.Len(t, labels, 1)
	assert.Equal(t, secatest.StorageSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.StorageSku1Iops, resp[0].Spec.Iops)
}

func TestGetStorageSkuV1(t *testing.T) {
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
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.StorageSku1Name, resp.Metadata.Name)

	labels := resp.Labels
	assert.Len(t, labels, 1)
	assert.Equal(t, secatest.StorageSku1Tier, labels[secatest.LabelKeyTier])

	assert.Equal(t, secatest.StorageSku1Iops, resp.Spec.Iops)
}

// Block Storage

func TestListBlockStoragesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListBlockStoragesV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef, err := regionalClient.StorageV1.BuildReferenceURN(secatest.StorageSku1Ref)
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

func TestListBlockStoragesWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListBlockStoragesV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := builders.NewListOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := regionalClient.StorageV1.ListBlockStoragesWithFilters(ctx, secatest.Tenant1Name, secatest.Workspace1Name, listOptions)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetBlockStorageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetBlockStorageV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef, err := regionalClient.StorageV1.BuildReferenceURN(secatest.StorageSku1Ref)
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

func TestCreateOrUpdateBlockStorageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateBlockStorageV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		SkuRef: secatest.StorageSku1Ref,
		Status: secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef, err := regionalClient.StorageV1.BuildReferenceURN(secatest.StorageSku1Ref)
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

func TestDeleteBlockStorageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetBlockStorageV1(sim, secatest.BlockStorageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.BlockStorage1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
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
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.StorageV1.DeleteBlockStorage(ctx, resp)
	assert.NoError(t, err)
}

// Image

func TestListImagesV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListStorageImagesV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.Image1Name,
			Tenant: secatest.Tenant1Name,
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef, err := regionalClient.StorageV1.BuildReferenceURN(secatest.BlockStorage1Ref)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := regionalClient.StorageV1.ListImages(ctx, secatest.Tenant1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)

	assert.Equal(t, *blockStorageRef, resp[0].Spec.BlockStorageRef)

	assert.Equal(t, secatest.StatusStateActive, string(*resp[0].Status.State))

	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := builders.NewListOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err = regionalClient.StorageV1.ListImagesWithFilters(ctx, secatest.Tenant1Name, listOptions)
	assert.NoError(t, err)

	resp, err = iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestListImagesWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListStorageImagesV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.Image1Name,
			Tenant: secatest.Tenant1Name,
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)
	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := builders.NewListOptions().WithLimit(10).WithLabels(labelsParams)

	iter, err := regionalClient.StorageV1.ListImagesWithFilters(ctx, secatest.Tenant1Name, listOptions)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetImageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetStorageImageV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:   secatest.Image1Name,
			Tenant: secatest.Tenant1Name,
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateActive},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef, err := regionalClient.StorageV1.BuildReferenceURN(secatest.BlockStorage1Ref)
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

func TestCreateOrUpdateImageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateImageV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Image1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
		},
		BlockStorageRef: secatest.BlockStorage1Ref,
		Status:          secatest.StatusResponseV1{State: secatest.StatusStateCreating},
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef, err := regionalClient.StorageV1.BuildReferenceURN(secatest.BlockStorage1Ref)
	if err != nil {
		t.Fatal(err)
	}

	image := &schema.Image{
		Metadata: &schema.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Image1Name,
		},
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

func TestDeleteImageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockGetStorageImageV1(sim, secatest.ImageResponseV1{
		Metadata: secatest.MetadataResponseV1{
			Name:      secatest.Image1Name,
			Tenant:    secatest.Tenant1Name,
			Workspace: ptr.To(secatest.Workspace1Name),
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
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	err = regionalClient.StorageV1.DeleteImage(ctx, resp)
	assert.NoError(t, err)
}
