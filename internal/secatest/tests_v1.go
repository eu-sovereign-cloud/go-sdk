package secatest

import (
	"net/http"
	"testing"

	mockwellknown "github.com/eu-sovereign-cloud/go-sdk/mock/spec/extensions.wellknown.v1"
	mockauthorization "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.authorization.v1"
	mockcompute "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.compute.v1"
	mocknetwork "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.network.v1"
	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	mockstorage "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.storage.v1"
	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/constants"
	wellknown "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
	authorization "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.authorization.v1"
	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	network "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	storage "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.storage.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

func ConfigureWellknownHandler(sim *mockwellknown.MockServerInterface, sm *http.ServeMux) {
	wellknown.HandlerWithOptions(sim, wellknown.StdHTTPServerOptions{
		BaseURL:    WellknownV1Endpoint,
		BaseRouter: sm,
	})
}

func ConfigureRegionV1Handler(t *testing.T, sm *http.ServeMux) *mockregion.MockServerInterface {
	sim := mockregion.NewMockServerInterface(t)

	MockGetRegionV1(sim, &schema.Region{
		Metadata: NewGlobalResourceMetadata(Region1Name),
		Spec: schema.RegionSpec{
			Providers: []schema.Provider{
				{
					Name:    constants.WorkspaceProviderName,
					Url:     ProviderWorkspaceV1Endpoint,
					Version: constants.ApiVersion1,
				},
				{
					Name:    constants.NetworkProviderName,
					Url:     ProviderNetworkV1Endpoint,
					Version: constants.ApiVersion1,
				},
				{
					Name:    constants.ComputeProviderName,
					Url:     ProviderComputeV1Endpoint,
					Version: constants.ApiVersion1,
				},
				{
					Name:    constants.StorageProviderName,
					Url:     ProviderStorageV1Endpoint,
					Version: constants.ApiVersion1,
				},
			},
		},
	})
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    ProviderRegionV1Endpoint,
		BaseRouter: sm,
	})

	return sim
}

func ConfigureRegionHandler(sim *mockregion.MockServerInterface, sm *http.ServeMux) {
	region.HandlerWithOptions(sim, region.StdHTTPServerOptions{
		BaseURL:    ProviderRegionV1Endpoint,
		BaseRouter: sm,
	})
}

func ConfigureAuthorizationHandler(sim *mockauthorization.MockServerInterface, sm *http.ServeMux) {
	authorization.HandlerWithOptions(sim, authorization.StdHTTPServerOptions{
		BaseURL:    ProviderAuthorizationV1Endpoint,
		BaseRouter: sm,
	})
}

func ConfigureWorkspaceHandler(sim *mockworkspace.MockServerInterface, sm *http.ServeMux) {
	workspace.HandlerWithOptions(sim, workspace.StdHTTPServerOptions{
		BaseURL:    ProviderWorkspaceV1Endpoint,
		BaseRouter: sm,
	})
}

func ConfigureComputeHandler(sim *mockcompute.MockServerInterface, sm *http.ServeMux) {
	compute.HandlerWithOptions(sim, compute.StdHTTPServerOptions{
		BaseURL:    ProviderComputeV1Endpoint,
		BaseRouter: sm,
	})
}

func ConfigureNetworkHandler(sim *mocknetwork.MockServerInterface, sm *http.ServeMux) {
	network.HandlerWithOptions(sim, network.StdHTTPServerOptions{
		BaseURL:    ProviderNetworkV1Endpoint,
		BaseRouter: sm,
	})
}

func ConfigureStorageHandler(sim *mockstorage.MockServerInterface, sm *http.ServeMux) {
	storage.HandlerWithOptions(sim, storage.StdHTTPServerOptions{
		BaseURL:    ProviderStorageV1Endpoint,
		BaseRouter: sm,
	})
}
