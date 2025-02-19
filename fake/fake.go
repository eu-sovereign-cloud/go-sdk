package fake

import (
	"net/http"
	"net/http/httptest"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/workspace.v1"
)

type Server struct {
	Regions    map[string]*region.Region
	Workspaces map[string]*workspace.Workspace
}

func NewServer(regionNames ...string) *Server {
	regionsMap := make(map[string]*region.Region)
	for _, name := range regionNames {
		regionsMap[name] = &region.Region{
			Metadata: &region.ResourceMetadata{
				Name:       name,
				ApiVersion: "v1",
				Kind:       region.ResourceMetadataKindRegion,
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
			Status: &region.Status{
				Conditions: []region.StatusCondition{
					{
						Status: "Ready", // TODO: no enum?
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

	workspace.HandlerWithOptions(s, workspace.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.workspace",
		BaseRouter: mux,
	})
	region.HandlerWithOptions(s, region.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.regions",
		BaseRouter: mux,
	})

	return httptest.NewServer(mux)
}
