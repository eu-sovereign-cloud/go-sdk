package secapi

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type RegionalClient struct {
	ComputeV1   *ComputeAPIV1
	NetworkV1   *NetworkAPIV1
	StorageV1   *StorageAPIV1
	WorkspaceV1 *WorkspaceAPIV1
}

func NewRegionalClient(region *region.Region) *RegionalClient {
	return &RegionalClient{
		ComputeV1:   newComputeAPIV1(region),
		NetworkV1:   newNetworkAPIV1(region),
		StorageV1:   newStorageAPIV1(region),
		WorkspaceV1: newWorkspaceAPIV1(region),
	}
}
