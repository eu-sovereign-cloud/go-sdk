package secapi

import (
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type RegionalClient struct {
	ComputeV1   *ComputeV1
	NetworkV1   *NetworkV1
	StorageV1   *StorageV1
	WorkspaceV1 *WorkspaceV1
}

func NewRegionalClient(region *region.Region) *RegionalClient {
	return &RegionalClient{
		ComputeV1:   newComputeV1(region),
		NetworkV1:   newNetworkV1(region),
		StorageV1:   newStorageV1(region),
		WorkspaceV1: newWorkspaceV1(region),
	}
}
