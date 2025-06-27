package fake

import (
	"encoding/json"
	"net/http"

	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.workspace.v1"
)

var _ workspace.ServerInterface = (*Server)(nil)

// DeleteWorkspace implements workspace.ServerInterface.
func (s *Server) DeleteWorkspace(w http.ResponseWriter, r *http.Request, id workspace.Tenant, name string, params workspace.DeleteWorkspaceParams) {
	panic("unimplemented")
}

// ListWorkspaces implements workspace.ServerInterface.
func (s *Server) ListWorkspaces(w http.ResponseWriter, r *http.Request, id workspace.Tenant, params workspace.ListWorkspacesParams) {
	var resp workspace.ListWorkspacesResponse

	resp.JSON200 = &workspace.WorkspaceIterator{
		Items: make([]workspace.Workspace, 0, len(s.Workspaces)),
	}

	for _, workspace := range s.Workspaces {
		resp.JSON200.Items = append(resp.JSON200.Items, *workspace)
	}

	http.Header.Add(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.JSON200) // nolint:errcheck
}

// GetWorkspace implements workspace.ServerInterface.
func (s *Server) GetWorkspace(w http.ResponseWriter, r *http.Request, id workspace.Tenant, name string) {
	var resp workspace.GetWorkspaceResponse

	resp.JSON200 = &workspace.Workspace{}

	http.Header.Add(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.JSON200)
}

// CreateOrUpdateWorkspace implements workspace.ServerInterface.
func (s *Server) CreateOrUpdateWorkspace(w http.ResponseWriter, r *http.Request, id workspace.Tenant, name workspace.ResourceName, ws workspace.CreateOrUpdateWorkspaceParams) {
	var resp workspace.CreateOrUpdateWorkspaceResponse

	resp.JSON200 = &workspace.Workspace{}

	http.Header.Add(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.JSON200)
}
