package fake

import (
	"net/http"
	"net/http/httptest"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
)

type Server struct {
	Wellknown  *wellknown.Wellknown
	Regions    map[string]*region.Region
	Workspaces map[string]*workspace.Workspace
}

func NewServer(regionNames ...string) *Server {
	wellknown := &wellknown.Wellknown{
		Version: "v1",
		Endpoints: []wellknown.WellknownEndpoint{
			{
				Provider: "seca.region/v1",
				Url:      "/providers/seca.regions",
			},
		},
	}

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
		Wellknown:  wellknown,
		Regions:    regionsMap,
		Workspaces: make(map[string]*workspace.Workspace),
	}
}

func (s *Server) Start() *httptest.Server {
	mux := http.NewServeMux()

	wellknown.HandlerWithOptions(s, wellknown.StdHTTPServerOptions{
		BaseURL:    "/.wellknown/secapi",
		BaseRouter: mux,
	})

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
