package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.network.v1"
)

type NetworkV1 struct {
	network network.ClientWithResponsesInterface
}

func newNetworkV1(networkUrl string) (*NetworkV1, error) {
	network, err := network.NewClientWithResponses(networkUrl)
	if err != nil {
		return nil, err
	}

	return &NetworkV1{network: network}, nil
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

func (api *NetworkV1) ListNetworks(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[network.Network], error) {
	iter := Iterator[network.Network]{
		fn: func(ctx context.Context, skipToken *string) ([]network.Network, *string, error) {
			resp, err := api.network.ListNetworksWithResponse(ctx, network.Tenant(tid), network.Workspace(wid), &network.ListNetworksParams{
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

func (api *NetworkV1) GetNetwork(ctx context.Context, wref WorkspaceReference) (*network.Network, error) {
	resp, err := api.network.GetNetworkWithResponse(ctx, network.Tenant(wref.Tenant), network.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *NetworkV1) CreateOrUpdateInstance(ctx context.Context, net *network.Network) error {
	validateNetworkMetadataV1(net.Metadata)

	resp, err := api.network.CreateOrUpdateNetworkWithResponse(ctx, net.Metadata.Tenant, *net.Metadata.Workspace, net.Metadata.Name,
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

func (api *NetworkV1) DeleteInstance(ctx context.Context, net *network.Network) error {
	validateNetworkMetadataV1(net.Metadata)

	resp, err := api.network.DeleteNetworkWithResponse(ctx, net.Metadata.Tenant, *net.Metadata.Workspace, net.Metadata.Name, &network.DeleteNetworkParams{
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
