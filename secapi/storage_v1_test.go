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

	"github.com/stretchr/testify/assert"
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
		Name:   "sku1",
		Tenant: secatest.Tenant1Name,
		Skus: []secatest.ListStorageSkuMetaInfoResponseProviderV1{
			{
				Provider:      "seca",
				Tier:          "RD500",
				Iops:          100,
				MinVolumeSize: 50,
				Type:          "remote-durable",
			},
			{
				Provider:      "seca",
				Tier:          "DXS",
				Iops:          200,
				MinVolumeSize: 50,
				Type:          "remote-durable",
			},
		},
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	sgIter, err := regionalClient.StorageV1.ListSkus(ctx, secatest.Tenant1Name)
	require.NoError(t, err)

	sg, err := sgIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, sg, 2)

	for _, sku := range sg {

		require.NotEmpty(t, sku.Labels)
		require.NotEmpty(t, sku.Spec.Iops)
		require.NotEmpty(t, sku.Spec.MinVolumeSize)
	}
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
	secatest.MockGetStorageSkusV1(sgSim, secatest.NameResponseV1{
		Name: secatest.Workspace1Name,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	sg, err := regionalClient.StorageV1.GetSku(ctx, TenantReference{
		Tenant: secatest.Tenant1Name,
		Name:   secatest.Workspace1Name,
	})
	require.NoError(t, err)
	require.NotEmpty(t, sg)

	assert.Equal(t, secatest.Workspace1Name, sg.Metadata.Name)
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
	secatest.MockListBlockStoragesV1(sgSim, secatest.BlockStorageResponseV1{
		Name:      secatest.Workspace1Name,
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Region:    secatest.RegionName,
		Zone:      secatest.ZoneA,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	sgIter, err := regionalClient.StorageV1.ListBlockStorages(ctx, secatest.Tenant1Name, secatest.Workspace1Name)
	require.NoError(t, err)

	sg, err := sgIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, sg, 1)

	assert.Equal(t, secatest.Workspace1Name, sg[0].Metadata.Name)
	assert.Equal(t, secatest.Tenant1Name, sg[0].Metadata.Tenant)
	assert.Equal(t, secatest.RegionName, sg[0].Metadata.Region)
	assert.Equal(t, secatest.ZoneA, sg[0].Metadata.Zone)
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
	secatest.MockGetBlockStorageV1(sgSim, secatest.BlockStorageResponseV1{
		Name:      secatest.Storage1Name,
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Region:    secatest.RegionName,
		Zone:      secatest.ZoneA,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)
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
	assert.Equal(t, secatest.RegionName, sg.Metadata.Region)
	assert.Equal(t, secatest.ZoneA, sg.Metadata.Zone)
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
	secatest.MockCreateOrUpdateBlockStorageV1(sgSim, secatest.NameResponseV1{
		Name: secatest.Workspace1Name,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)
	ws := secatest.Workspace1Name
	block := &storage.BlockStorage{
		Metadata: &storage.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Workspace1Name,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)
	ws := secatest.Workspace1Name
	block := &storage.BlockStorage{
		Metadata: &storage.ZonalResourceMetadata{
			Tenant:    secatest.Tenant1Name,
			Name:      secatest.Workspace1Name,
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
	secatest.MockListStorageImagesV1(sgSim, secatest.ImageResponseV1{
		Name:      secatest.Image1Name,
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Region:    secatest.Region1Name,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)

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
	secatest.MockGetStorageImageV1(sgSim, secatest.ImageResponseV1{
		Name:      secatest.Image1Name,
		Tenant:    secatest.Tenant1Name,
		Workspace: secatest.Workspace1Name,
		Region:    secatest.Region1Name,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)

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
	secatest.MockCreateOrUpdateImageV1(sgSim, secatest.NameResponseV1{
		Name: secatest.Image1Name,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	image := &storage.Image{
		Metadata: &storage.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Image1Name,
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

	regionalClient, err := client.NewRegionalClient(ctx, secatest.RegionName, []RegionalAPI{StorageV1API})
	require.NoError(t, err)

	image := &storage.Image{
		Metadata: &storage.RegionalResourceMetadata{
			Tenant: secatest.Tenant1Name,
			Name:   secatest.Image1Name,
		},
	}

	err = regionalClient.StorageV1.DeleteImage(ctx, image)
	require.NoError(t, err)
}
