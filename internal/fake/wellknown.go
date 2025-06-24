package fake

import (
	"encoding/json"
	"net/http"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
)

var _ wellknown.ServerInterface = (*Server)(nil)

// GetWellknown implements wellknown.ServerInterface.
func (s *Server) GetWellknown(w http.ResponseWriter, r *http.Request) {
	var resp wellknown.GetWellknownResponse

	resp.JSON200 = &wellknown.Wellknown{
		Version: s.Wellknown.Version,
	}

	for _, endpoint := range s.Wellknown.Endpoints {
		resp.JSON200.Endpoints = append(resp.JSON200.Endpoints, wellknown.WellknownEndpoint{
			Provider: endpoint.Provider,
			Url:      "http://" + r.Host + endpoint.Url,
		})
	}

	http.Header.Add(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.JSON200) // nolint:errcheck
}
