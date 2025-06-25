package fake

import (
	"net/http"
	"net/http/httptest"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
)

type Server struct {
	Regions    map[string]*region.Region
	Workspaces map[string]*workspace.Workspace
}

func NewServer(regionNames ...string) *Server {
	regionsMap := make(map[string]*region.Region)
	for _, name := range regionNames {
		regionsMap[name] = &region.Region{
			Metadata: &region.GlobalResourceMetadata{
				Name:       name,
				ApiVersion: "v1",
				Kind:       region.GlobalResourceMetadataKindRegion,
			},
			Spec: region.RegionSpec{
				AvailableZones: []string{"A", "B"},
				Providers: []region.Provider{
					{
						Name:    "seca.workspace",
						Version: "v1",
					},
				},
			},
		}
	}

	return &Server{
		Regions:    regionsMap,
		Workspaces: make(map[string]*workspace.Workspace),
	}
}

func (s *Server) Start() *httptest.Server {
	mux := http.NewServeMux()

	region.HandlerWithOptions(s, region.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.regions",
		BaseRouter: mux,
	})

	workspace.HandlerWithOptions(s, workspace.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.workspace",
		BaseRouter: mux,
	})

	return httptest.NewServer(mux)
}
