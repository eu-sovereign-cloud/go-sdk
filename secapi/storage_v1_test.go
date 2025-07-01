package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"
	mockRegion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	mockStorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"

	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"

	"github.com/stretchr/testify/require"
)

func TestListSkus(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockStorageListSkusV1(sgSim, secatest.ListStorageSkusResponseV1{
		Tenant: "test",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	sgIter, err := regionalClient.StorageV1.ListSkus(ctx, "test", "some-workspace")
	require.NoError(t, err)

	sg, err := sgIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, sg, 6)
}

func TestGetSku(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockGetStorageSkusV1(sgSim, secatest.GetStorageSkuResponseV1{
		Name: "some-workspace",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	cp, err := regionalClient.StorageV1.GetSku(ctx, TenantReference{
		Tenant: "test-tenant",
		Name:   "some-workspace",
	})
	require.NoError(t, err)
	require.NotEmpty(t, cp)
}

func TestListBlockStorages(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockListBlockStoragesV1(sgSim, secatest.GenericResponseV1{
		Name: "some-workspace",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	sgIter, err := regionalClient.StorageV1.ListBlockStorages(ctx, "test-tenant", "some-workspace")
	require.NoError(t, err)

	sg, err := sgIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, sg, 1)
}

func TestGetBlockStorage(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockGetBlockStorageV1(sgSim, secatest.GenericResponseV1{
		Name: "workspace_1",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)
	wref := WorkspaceReference{
		TenantReference: TenantReference{
			Tenant: "test-tenant",
			Name:   "some-workspace",
		},
		Workspace: "workspace_1",
	}
	sg, err := regionalClient.StorageV1.GetBlockStorage(ctx, wref)
	require.NoError(t, err)
	require.NotEmpty(t, sg)
	require.Equal(t, "workspace_1", sg.Metadata.Name)
}

func TestCreateOrUpdateBlockStorage(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateBlockStorageV1(sgSim, secatest.GenericResponseV1{
		Name: "workspace_1",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)
	ws := "workspace_1"
	block := &storage.BlockStorage{
		Metadata: &storage.ZonalResourceMetadata{
			Tenant:    "test-tenant",
			Name:      "workspace_1",
			Workspace: &ws,
		},
	}

	err = regionalClient.StorageV1.CreateOrUpdateBlockStorage(ctx, block)
	require.NoError(t, err)
}

func TestDeleteBlockStorage(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockDeleteBlockStorageV1(sgSim)

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)
	ws := "workspace_1"
	block := &storage.BlockStorage{
		Metadata: &storage.ZonalResourceMetadata{
			Tenant:    "test-tenant",
			Name:      "workspace_1",
			Workspace: &ws,
		},
	}

	err = regionalClient.StorageV1.DeleteBlockStorage(ctx, block)
	require.NoError(t, err)
}

func TestListImageStorage(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockListStorageImagesV1(sgSim, secatest.GenericResponseV1{
		Name: "test",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	imgIter, err := regionalClient.StorageV1.ListImages(ctx, "test-tenant")
	require.NoError(t, err)

	images, err := imgIter.All(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, images)
}
func TestGetImage(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockGetStorageImageV1(sgSim, secatest.GenericResponseV1{
		Name: "test-image",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	tref := TenantReference{
		Tenant: "test-tenant",
		Name:   "test-image",
	}
	img, err := regionalClient.StorageV1.GetImage(ctx, tref)
	require.NoError(t, err)
	require.NotNil(t, img)
	require.Equal(t, "test-image", img.Metadata.Name)
}
func TestCreateOrUpdateImage(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockCreateOrUpdateImageV1(sgSim, secatest.GenericResponseV1{
		Name: "test-image",
	})

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	image := &storage.Image{
		Metadata: &storage.RegionalResourceMetadata{
			Tenant: "test-tenant",
			Name:   "test-image",
		},
	}

	err = regionalClient.StorageV1.CreateOrUpdateImage(ctx, image)
	require.NoError(t, err)
}
func TestDeleteImage(t *testing.T) {
	ctx := context.Background()

	sim := mockRegion.NewMockServerInterface(t)
	secatest.MockGetRegionV1(sim, secatest.GetRegionResponseV1{
		Name: secatest.RegionName,
		Providers: []secatest.GetRegionResponseProviderV1{
			{
				Name: secatest.ProviderStorageName,
				URL:  secatest.ProviderStorageEndpoint,
			},
		},
	})
	sgSim := mockStorage.NewMockServerInterface(t)
	secatest.MockDeleteImageV1(sgSim)

	sm := http.NewServeMux()
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderRegionEndpoint,
		BaseRouter: sm,
	})
	storage.HandlerWithOptions(sgSim, storage.StdHTTPServerOptions{
		BaseURL:    secatest.ProviderStorageEndpoint,
		BaseRouter: sm,
	})

	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewGlobalClient(&GlobalEndpoints{RegionV1: server.URL + secatest.ProviderRegionEndpoint})
	require.NoError(t, err)

	regionalClient, err := client.NewRegionalClient(ctx, "eu-central-1", []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	image := &storage.Image{
		Metadata: &storage.RegionalResourceMetadata{
			Tenant: "test-tenant",
			Name:   "test-image",
		},
	}

	err = regionalClient.StorageV1.DeleteImage(ctx, image)
	require.NoError(t, err)
}
