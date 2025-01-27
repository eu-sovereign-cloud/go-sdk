package gosdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/mock/mockregions.v1"
	mockworkspace "github.com/eu-sovereign-cloud/go-sdk/mock/mockworkspace.v1"
	regions "github.com/eu-sovereign-cloud/go-sdk/pkg/regions.v1"
	workspace "github.com/eu-sovereign-cloud/go-sdk/pkg/workspace.v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWorkspaces(t *testing.T) {
	ctx := context.Background()
	wsSim := mockworkspace.NewMockServerInterface(t)

	wsSim.EXPECT().ListWorkspaces(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, lwp workspace.ListWorkspacesParams) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"items": [
					{
						"apiVersion": "v1",
						"kind": "workspace",
						"metadata": {
							"name": "primary-load-balancer",
							"deletionTimestamp": "2019-08-24T14:15:22Z",
							"lastModifiedTimestamp": "2019-08-24T14:15:22Z",
							"description": "string",
							"labels": {
							"language": "en",
							"billing.cost-center": "platform-eng",
							"env": "production"
							}
						},
						"spec": {
							"name": "some-workspace",
							"description": "string"
						},
						"status": {
							"conditions": [
								{
									"type": "power-mgmt",
									"status": "True, false, unknown",
									"lastTransitionTime": "2019-08-24T14:15:22Z",
									"reason": "^A(A)?$",
									"message": "string"
								}
							],
							"phase": "Pending",
							"resourceCount": 0
						}
					}
				],
				"metadata": {
					"skipToken": null
				}
			}
		`))
	})

	reSim := mockregions.NewMockServerInterface(t)

	reSim.EXPECT().GetRegion(mock.Anything, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(w http.ResponseWriter, r *http.Request, s string, name string) {
		assert.Equal(t, "eu-central-1", name)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"apiVersion": "v1",
				"kind": "region",
				"metadata": {
					"name": "eu-central-1",
					"deletionTimestamp": "2019-08-24T14:15:22Z",
					"lastModifiedTimestamp": "2019-08-24T14:15:22Z",
					"description": "string",
					"labels": {
						"language": "en",
						"billing.cost-center": "platform-eng",
						"env": "production"
					}
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
					"conditions": [
						{
							"type": "power-mgmt",
							"status": "True, false, unknown",
							"lastTransitionTime": "2019-08-24T14:15:22Z",
							"reason": "^A(A)?$",
							"message": "string"
						}
					]
				}
			}
		`))
	})

	sm := http.NewServeMux()
	workspace.HandlerWithOptions(wsSim, workspace.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.workspace",
		BaseRouter: sm,
	})
	regions.HandlerWithOptions(reSim, regions.StdHTTPServerOptions{
		BaseURL:    "/providers/seca.regions",
		BaseRouter: sm,
	})
	server := httptest.NewServer(sm)
	defer server.Close()

	client, err := NewClient(server.URL + "/providers/seca.regions")
	require.NoError(t, err)

	ctx = WithTenantID(ctx, "test")

	regionClient, err := client.RegionClient(ctx, "eu-central-1")
	require.NoError(t, err)

	wsIter, err := regionClient.Workspaces(ctx)
	require.NoError(t, err)

	ws, err := wsIter.All(ctx)
	require.NoError(t, err)
	require.Len(t, ws, 1)

	assert.Equal(t, "some-workspace", ws[0].Spec.Name)
	assert.EqualValues(t, "Pending", *ws[0].Status.Phase)
}
