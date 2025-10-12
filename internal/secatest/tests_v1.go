package secatest

import (
	"net/http"
	"testing"

	mockauthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"
	mockcompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secalib"
)

func ConfigureRegionV1Handler(t *testing.T, sm *http.ServeMux) *mockregion.MockServerInterface {
	sim := mockregion.NewMockServerInterface(t)

	MockGetRegionV1(sim, &schema.Region{
		Metadata: secalib.BuildResponseGlobalResourceMetadata(InstanceSku1Name, Tenant1Name),
		Spec: schema.RegionSpec{
			Providers: []schema.Provider{
				{
					Name:    ProviderWorkspaceName,
					Url:     ProviderWorkspaceEndpoint,
					Version: ProviderVersion1,
				},
				{
					Name:    ProviderNetworkName,
					Url:     ProviderNetworkEndpoint,
					Version: ProviderVersion1,
				},
				{
					Name:    ProviderComputeName,
					Url:     ProviderComputeEndpoint,
					Version: ProviderVersion1,
				},
				{
					Name:    ProviderStorageName,
					Url:     ProviderStorageEndpoint,
					Version: ProviderVersion1,
				},
			},
		},
	})
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    ProviderRegionEndpoint,
		BaseRouter: sm,
	})

	return sim
}

func ConfigureAuthorizationHandler(sim *mockauthorization.MockServerInterface, sm *http.ServeMux) {
	authorization.HandlerWithOptions(sim, authorization.StdHTTPServerOptions{
		BaseURL:    ProviderAuthorizationEndpoint,
		BaseRouter: sm,
	})
}

func ConfigureComputeHandler(sim *mockcompute.MockServerInterface, sm *http.ServeMux) {
	compute.HandlerWithOptions(sim, compute.StdHTTPServerOptions{
		BaseURL:    ProviderComputeEndpoint,
		BaseRouter: sm,
	})
}

func ConfigureNetworkHandler(sim *mocknetwork.MockServerInterface, sm *http.ServeMux) {
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    ProviderNetworkEndpoint,
		BaseRouter: sm,
	})
}

func ConfigureRegionHandler(sim *mockregion.MockServerInterface, sm *http.ServeMux) {
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    ProviderRegionEndpoint,
		BaseRouter: sm,
	})
}

func ConfigureStorageHandler(sim *mockstorage.MockServerInterface, sm *http.ServeMux) {
	storage.HandlerWithOptions(sim, storage.StdHTTPServerOptions{
		BaseURL:    ProviderStorageEndpoint,
		BaseRouter: sm,
	})
}

func ConfigureWorkspaceHandler(sim *mockworkspace.MockServerInterface, sm *http.ServeMux) {
	workspace.HandlerWithOptions(sim, workspace.StdHTTPServerOptions{
		BaseURL:    ProviderWorkspaceEndpoint,
		BaseRouter: sm,
	})
}
