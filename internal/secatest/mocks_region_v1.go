package secatest

import (
	"net/http"

	mockregion "github.com/eu-sovereign-cloud/go-sdk/mock/spec/foundation.region.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/mock"
)

// Region
func MockListRegionsV1(sim *mockregion.MockServerInterface, resp []schema.Region) {
	sim.EXPECT().ListRegions(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, lrp region.ListRegionsParams) {
			configHttpResponse(w, http.StatusOK)

			for _, region := range resp {
				for i := range region.Spec.Providers {
					region.Spec.Providers[i].Url = "http://" + r.Host + region.Spec.Providers[i].Url
				}
			}

			iter := region.RegionIterator{Items: resp}
			if err := encodeResponseBody(w, iter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}

func MockGetRegionV1(sim *mockregion.MockServerInterface, resp *schema.Region) {
	sim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request, name string) {
			configHttpResponse(w, http.StatusOK)

			for i := range resp.Spec.Providers {
				resp.Spec.Providers[i].Url = "http://" + r.Host + resp.Spec.Providers[i].Url
			}

			if err := encodeResponseBody(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}
