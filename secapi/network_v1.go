package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secapi"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type NetworkAPIV1 struct {
	secapi.RegionalAPI[network.ClientWithResponsesInterface]
}

func newNetworkAPIV1(region *region.Region) *NetworkAPIV1 {
	return &NetworkAPIV1{
		RegionalAPI: secapi.RegionalAPI[network.ClientWithResponsesInterface]{Region: region},
	}
}

func (api *NetworkAPIV1) getClient() (network.ClientWithResponsesInterface, error) {
	fn := func(url string) (network.ClientWithResponsesInterface, error) {
		return network.NewClientWithResponses(url)
	}

	client, err := api.GetClient("seca.network", fn)
	if err != nil {
		return nil, err
	}
	return *client, nil
}

func validateNetworkMetadataV1(metadata *network.RegionalResourceMetadata) {
	if metadata == nil {
		panic(ErrNoMetatada)
	}

	if metadata.Workspace == nil {
		panic(ErrNoMetatadaWorkspace)
	}

	if metadata.Tenant == "" {
		panic(ErrNoMetatadaTenant)
	}
}

func (api *NetworkAPIV1) ListNetworks(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.Network], error) {
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	iter := Iterator[network.Network]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Network, *string, error) {
			resp, err := client.ListNetworksWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListNetworksParams{
				Accept: ptr.To(network.ListNetworksParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *NetworkAPIV1) GetNetwork(ctx context.Context, wref WorkspaceReference) (*network.Network, error) {
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GetNetworkWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkAPIV1) CreateOrUpdateInstance(ctx context.Context, net *network.Network) error {
	validateNetworkMetadataV1(net.Metadata)

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.CreateOrUpdateNetworkWithResponse(ctx, net.Metadata.Tenant, *net.Metadata.Workspace, net.Metadata.Name,
		&network.CreateOrUpdateNetworkParams{
			IfUnmodifiedSince: &net.Metadata.ResourceVersion,
		}, *net)
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 200, 201)
	if err != nil {
		return err
	}

	return nil
}

func (api *NetworkAPIV1) DeleteInstance(ctx context.Context, net *network.Network) error {
	validateNetworkMetadataV1(net.Metadata)

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.DeleteNetworkWithResponse(ctx, net.Metadata.Tenant, *net.Metadata.Workspace, net.Metadata.Name, &network.DeleteNetworkParams{
		IfUnmodifiedSince: &net.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 204, 404)
	if err != nil {
		return err
	}

	return nil
}
