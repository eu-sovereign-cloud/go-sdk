package secapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/fake"
	"github.com/eu-sovereign-cloud/go-sdk/mock/mockregion.v1"
	"github.com/eu-sovereign-cloud/go-sdk/mock/mockworkspace.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/foundation.region.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/foundation.workspace.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMockedWorkspaces(t *testing.T) {
	ctx := context.Background()
	wsSim := mockworkspace.NewMockServerInterface(t)

	wsSim.EXPECT().ListWorkspaces(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, lwp workspace.ListWorkspacesParams) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
			{
				"items": [
					{
						"apiVersion": "v1",
						"kind": "workspace",
						"metadata": {
							"name": "some-workspace",
							"tenant": "test",
							"deletionTimestamp": "2019-08-24T14:15:22Z",
							"lastModifiedTimestamp": "2019-08-24T14:15:22Z",
							"description": "string",
							"labels": {
								"language": "en",
								"billing.cost-center": "platform-eng",
								"env": "production"
							}
						},
						"spec": {},
						"status": {
							"conditions": [ { "status": "Ready" } ]
						}
					}
				],
				"metadata": {
					"skipToken": null
				}
			}
		`))
	})

	reSim := mockregion.NewMockServerInterface(t)

	reSim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, name string) {
		assert.Equal(t, "eu-central-1", name)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
			{
				"apiVersion": "v1",
				"kind": "region",
				"metadata": {
					"name": "eu-central-1"
				},
				"spec": {
					"availableZones": [ "A", "B" ],
					"providers": [
						{
							"name": "seca.workspace",
							"version": "v1",
							"url": "http://` + r.Host + `/providers/seca.workspace"
						}
					]
				},
				"status": {
					"conditions": [ { "status": "Ready" } ]
				}
			}
		`))
	})

	sm := http.NewServeMux()
	workspace.HandlerWithOptions(wsSim, workspace.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.workspace",
		BaseRouter: sm,
	})
	region.HandlerWithOptions(reSim, region.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.regions",
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewClient(server.URL + "/providers/seca.regions")
	require.NoError(t, err)

	regionClient, err := client.RegionClient(ctx, "eu-central-1")
	require.NoError(t, err)

	wsIter, err := regionClient.Workspaces(ctx, "test")
	require.NoError(t, err)

	ws, err := wsIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, ws, 1)

	assert.Equal(t, "some-workspace", ws[0].Metadata.Name)
}

func TestFakedWorkspaces(t *testing.T) {
	ctx := context.Background()

	fakeServer := fake.NewServer("eu-central-1")
	server := fakeServer.Start()
	defer server.Close()

	fakeServer.Workspaces["some-workspace"] = &workspace.Workspace{
		Metadata: &workspace.RegionalResourceMetadata{
			Tenant: "test",
			Name:   "some-workspace",
		},
		Status: &workspace.WorkspaceStatus{
			State: ptr.To(workspace.ResourceStatePending),
		},
	}

	client, err := NewClient(server.URL + "/providers/seca.regions")
	require.NoError(t, err)

	regionClient, err := client.RegionClient(ctx, "eu-central-1")
	require.NoError(t, err)

	wsIter, err := regionClient.Workspaces(ctx, "test")
	require.NoError(t, err)

	ws, err := wsIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, ws, 1)

	assert.Equal(t, "some-workspace", ws[0].Metadata.Name)
	assert.Equal(t, "test", ws[0].Metadata.Tenant)
	assert.EqualValues(t, "pending", *ws[0].Status.State)
}
