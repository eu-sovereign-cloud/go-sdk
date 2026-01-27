package secatest

import (
	"net/http"

	mockwellknown "github.com/eu-sovereign-cloud/go-sdk/mock/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"github.com/stretchr/testify/mock"
)

// Wellknown

func MockGetWellknownV1(sim *mockwellknown.MockServerInterface, resp *schema.Wellknown) {
	sim.EXPECT().GetWellknown(mock.Anything, mock.Anything).
		RunAndReturn(func(w http.ResponseWriter, r *http.Request) {
			configHttpResponse(w, http.StatusOK)

			for i := range resp.Endpoints {
				resp.Endpoints[i].Url = "http://" + r.Host + resp.Endpoints[i].Url
			}

			if err := encodeResponseBody(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
}
