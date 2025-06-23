package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
)

type ComputeV1 struct {
	compute compute.ClientWithResponsesInterface
}

func newComputeV1(computeUrl string) (*ComputeV1, error) {
	compute, err := compute.NewClientWithResponses(computeUrl)
	if err != nil {
		return nil, err
	}

	return &ComputeV1{compute: compute}, nil
}

func validateComputeMetadataV1(metadata *compute.ZonalResourceMetadata) {
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

func (api *ComputeV1) ListInstances(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[compute.Instance], error) {
	iter := Iterator[compute.Instance]{
		fn: func(ctx context.Context, skipToken *string) ([]compute.Instance, *string, error) {
			resp, err := api.compute.ListInstancesWithResponse(ctx, compute.Tenant(tid), compute.Workspace(wid), &compute.ListInstancesParams{
				Accept: ptr.To(compute.Applicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *ComputeV1) GetInstance(ctx context.Context, wref WorkspaceReference) (*compute.Instance, error) {
	resp, err := api.compute.GetInstanceWithResponse(ctx, compute.Tenant(wref.Tenant), compute.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *ComputeV1) CreateOrUpdateInstance(ctx context.Context, inst *compute.Instance) error {
	validateComputeMetadataV1(inst.Metadata)

	resp, err := api.compute.CreateOrUpdateInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name,
		&compute.CreateOrUpdateInstanceParams{
			IfUnmodifiedSince: &inst.Metadata.ResourceVersion,
		}, *inst)
	if err != nil {
		return err
	}

	err = checkStatusCode(resp, 200, 201)
	if err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) DeleteInstance(ctx context.Context, inst *compute.Instance) error {
	validateComputeMetadataV1(inst.Metadata)

	resp, err := api.compute.DeleteInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, &compute.DeleteInstanceParams{
		IfUnmodifiedSince: &inst.Metadata.ResourceVersion,
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
