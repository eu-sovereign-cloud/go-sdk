package fake

import (
	"encoding/json"
	"net/http"

	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

var _ region.ServerInterface = (*Server)(nil)

// GetRegion implements regions.ServerInterface.
func (s *Server) GetRegion(w http.ResponseWriter, r *http.Request, name string) {
	region, ok := s.Regions[name]
	if !ok {
		http.Error(w, "region not found", http.StatusNotFound)
		return
	}
	for i, provider := range region.Spec.Providers {
		region.Spec.Providers[i].Url = "http://" + r.Host + "/providers/" + provider.Name
	}

	http.Header.Add(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(region) // nolint:errcheck
}

// ListRegions implements regions.ServerInterface.
func (s *Server) ListRegions(w http.ResponseWriter, r *http.Request, params region.ListRegionsParams) {
	var resp region.ListRegionsResponse

	// TODO: remove this once we have a proper type
	resp.JSON200 = &region.RegionIterator{
		Items: make([]region.Region, 0, len(s.Regions)),
	}

	for _, region := range s.Regions {
		for i, provider := range region.Spec.Providers {
			region.Spec.Providers[i].Url = "http://" + r.Host + "/providers/" + provider.Name
		}
		resp.JSON200.Items = append(resp.JSON200.Items, *region)
	}

	http.Header.Add(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.JSON200) // nolint:errcheck
}
