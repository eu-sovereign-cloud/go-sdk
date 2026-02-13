package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/constants"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi/builders"

	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

// Storage Sku

func TestListStorageSkusV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	labels := schema.Labels{secatest.LabelKeyTier: secatest.StorageSku1Tier}
	spec := buildResponseStorageSkuSpec(secatest.StorageSku1Iops)
	secatest.MockListStorageSkusV1(sim, []schema.StorageSku{
		*buildResponseStorageSku(secatest.StorageSku1Name, secatest.Tenant1Name, labels, spec),
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

func TestListStorageSkusWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	secatest.MockListStorageSkusV1(sim, []schema.StorageSku{
		{
			Metadata: &schema.SkuResourceMetadata{
				Name: secatest.StorageSku1Name,
			},
			Labels: schema.Labels{
				secatest.LabelKeyTier: secatest.StorageSku1Tier,
			},
			Spec: &schema.StorageSkuSpec{
				Iops: secatest.StorageSku1Iops,
				Type: secatest.StorageSku1Tier,
			},
		},
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
	labels := schema.Labels{secatest.LabelKeyTier: secatest.StorageSku1Tier}
	spec := buildResponseStorageSkuSpec(secatest.StorageSku1Iops)
	secatest.MockGetStorageSkusV1(sim, buildResponseStorageSku(secatest.StorageSku1Name, secatest.Tenant1Name, labels, spec))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	resp, err := regionalClient.StorageV1.GetSku(ctx, TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.StorageSku1Name})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.StorageSku1Name, resp.Metadata.Name)

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
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockListBlockStoragesV1(sim, []schema.BlockStorage{
		*buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
	})
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.StorageSku1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	iter, err := regionalClient.StorageV1.ListBlockStorages(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	assert.NoError(t, err)

	resp, err := iter.All(ctx)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, secatest.BlockStorage1Name, resp[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp[0].Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp[0].Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, *storageSkuRef, resp[0].Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)
}

func TestListBlockStoragesWithFiltersV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	ref := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.StorageSku1Ref, secatest.Tenant1Name, secatest.Workspace1Name)
	secatest.MockListBlockStoragesV1(sim, []schema.BlockStorage{
		{
			Metadata: &schema.RegionalWorkspaceResourceMetadata{
				Name:      secatest.BlockStorage1Name,
				Tenant:    secatest.Tenant1Name,
				Workspace: secatest.Workspace1Name,
			},
			Spec: schema.BlockStorageSpec{
				SkuRef: *ref,
			},
			Status: &schema.BlockStorageStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
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

	listOptions := NewListOptions().WithLimit(10).WithLabels(labelsParams)

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
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockGetBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)
	storageSkuRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.StorageSku1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.BlockStorage1Name}
	resp, err := regionalClient.StorageV1.GetBlockStorage(ctx, wref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.BlockStorage1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *storageSkuRef, resp.Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetBlockStorageUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockGetBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.StorageSku1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	wref := WorkspaceReference{Tenant: secatest.Tenant1Name, Workspace: secatest.Workspace1Name, Name: secatest.BlockStorage1Name}
	config := ResourceObserverConfig[schema.ResourceState]{ExpectedValue: schema.ResourceStateActive, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.StorageV1.GetBlockStorageUntilState(ctx, wref, config)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.BlockStorage1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *storageSkuRef, resp.Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestCreateOrUpdateBlockStorageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockCreateOrUpdateBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	storageSkuRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.StorageSku1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

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

	assert.Equal(t, secatest.BlockStorage1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Workspace1Name, resp.Metadata.Workspace)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *storageSkuRef, resp.Spec.SkuRef)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteBlockStorageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseBlockStorageSpec(t, secatest.StorageSku1Ref, secatest.BlockStorage1SizeGB)
	secatest.MockGetBlockStorageV1(sim, buildResponseBlockStorage(secatest.BlockStorage1Name, secatest.Tenant1Name, secatest.Workspace1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
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

	blockStorageRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.BlockStorage1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockListStorageImagesV1(sim, []schema.Image{
		*buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, secatest.Region1Name, spec, schema.ResourceStateActive),
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
	assert.Equal(t, secatest.Region1Name, resp[0].Metadata.Region)

	assert.Equal(t, *blockStorageRef, resp[0].Spec.BlockStorageRef)

	assert.Equal(t, schema.ResourceStateActive, *resp[0].Status.State)

	labelsParams := builders.NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue).
		Equals(secatest.LabelEnvKey, secatest.LabelEnvValue+"*").
		NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue).
		Neq(secatest.LabelTierKey, secatest.LabelTierValue).
		Gt(secatest.LabelVersion, 1).
		Lt(secatest.LabelVersion, 3).
		Gte(secatest.LabelUptime, 99).
		Lte(secatest.LabelLoad, 75)

	listOptions := NewListOptions().WithLimit(10).WithLabels(labelsParams)

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

	ref := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.BlockStorage1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	secatest.MockListStorageImagesV1(sim, []schema.Image{
		{
			Metadata: &schema.RegionalResourceMetadata{
				Name:   secatest.Image1Name,
				Tenant: secatest.Tenant1Name,
			},
			Spec: schema.ImageSpec{
				BlockStorageRef: *ref,
			},
			Status: &schema.ImageStatus{
				State: ptr.To(schema.ResourceStateActive),
			},
		},
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

	listOptions := NewListOptions().WithLimit(10).WithLabels(labelsParams)

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
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockGetStorageImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.BlockStorage1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Image1Name}
	resp, err := regionalClient.StorageV1.GetImage(ctx, tref)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *blockStorageRef, resp.Spec.BlockStorageRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestGetImageUntilStateV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockGetStorageImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, secatest.Region1Name, spec, schema.ResourceStateCreating), 2)
	secatest.MockGetStorageImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.BlockStorage1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	tref := TenantReference{Tenant: secatest.Tenant1Name, Name: secatest.Image1Name}
	config := ResourceObserverConfig[schema.ResourceState]{ExpectedValue: schema.ResourceStateActive, Delay: 0, Interval: 0, MaxAttempts: 5}
	resp, err := regionalClient.StorageV1.GetImageUntilState(ctx, tref, config)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *blockStorageRef, resp.Spec.BlockStorageRef)

	assert.Equal(t, schema.ResourceStateActive, *resp.Status.State)
}

func TestCreateOrUpdateImageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockCreateOrUpdateImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, secatest.Region1Name, spec, schema.ResourceStateCreating))
	secatest.ConfigureStorageHandler(sim, sm)

	server := httptest.NewServer(sm)
	defer server.Close()

	regionalClient := newTestRegionalClientV1(t, ctx, server)

	blockStorageRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, secatest.BlockStorage1Ref, secatest.Tenant1Name, secatest.Workspace1Name)

	image := &schema.Image{
		Metadata: secatest.NewRegionalResourceMetadata(secatest.Image1Name, secatest.Tenant1Name, secatest.Region1Name),
		Spec: schema.ImageSpec{
			BlockStorageRef: *blockStorageRef,
			CpuArchitecture: secatest.Image1CpuArch,
		},
	}
	resp, err := regionalClient.StorageV1.CreateOrUpdateImage(ctx, image)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, secatest.Image1Name, resp.Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, resp.Metadata.Tenant)
	assert.Equal(t, secatest.Region1Name, resp.Metadata.Region)

	assert.Equal(t, *blockStorageRef, resp.Spec.BlockStorageRef)

	assert.Equal(t, schema.ResourceStateCreating, *resp.Status.State)
}

func TestDeleteImageV1(t *testing.T) {
	ctx := context.Background()
	sm := http.NewServeMux()

	secatest.ConfigureRegionV1Handler(t, sm)

	sim := mockstorage.NewMockServerInterface(t)
	spec := buildResponseImageSpec(t, secatest.BlockStorage1Ref)
	secatest.MockGetStorageImageV1(sim, buildResponseImage(secatest.Image1Name, secatest.Tenant1Name, secatest.Region1Name, spec, schema.ResourceStateActive), 1)
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

// Builders

func buildResponseStorageSku(name string, tenant string, labels schema.Labels, spec *schema.StorageSkuSpec) *schema.StorageSku {
	return &schema.StorageSku{
		Metadata: secatest.NewSkuResourceMetadata(name, tenant),
		Labels:   labels,
		Spec:     spec,
	}
}

func buildResponseStorageSkuSpec(iops int) *schema.StorageSkuSpec {
	return &schema.StorageSkuSpec{
		Iops: iops,
	}
}

func buildResponseBlockStorage(name string, tenant string, workspace string, region string, spec *schema.BlockStorageSpec, state schema.ResourceState) *schema.BlockStorage {
	return &schema.BlockStorage{
		Metadata: secatest.NewRegionalWorkspaceResourceMetadata(name, tenant, workspace, region),
		Spec:     *spec,
		Status:   secatest.NewBlockStorageStatus(state),
	}
}

func buildResponseBlockStorageSpec(t *testing.T, skuRefName string, sizeGB int) *schema.BlockStorageSpec {
	skuRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, skuRefName, secatest.Tenant1Name, secatest.Workspace1Name)

	return &schema.BlockStorageSpec{
		SkuRef: *skuRef,
		SizeGB: sizeGB,
	}
}

func buildResponseImage(name string, tenant string, region string, spec *schema.ImageSpec, state schema.ResourceState) *schema.Image {
	return &schema.Image{
		Metadata: secatest.NewRegionalResourceMetadata(name, tenant, region),
		Spec:     *spec,
		Status:   secatest.NewImageStatus(state),
	}
}

func buildResponseImageSpec(t *testing.T, blockStorageRef string) *schema.ImageSpec {
	objRef := BuildReferenceObj(constants.StorageProviderV1Name, secatest.Region1Name, blockStorageRef, secatest.Tenant1Name, secatest.Workspace1Name)

	return &schema.ImageSpec{
		BlockStorageRef: *objRef,
	}
}
