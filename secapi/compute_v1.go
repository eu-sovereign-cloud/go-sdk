package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
)

type ComputeV1 struct {
	compute compute.ClientWithResponsesInterface
}

func (api *ComputeV1) ListSkus(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[compute.InstanceSku], error) {
	iter := Iterator[compute.InstanceSku]{
		fn: func(ctx context.Context, skipToken *string) ([]compute.InstanceSku, *string, error) {
			resp, err := api.compute.ListSkusWithResponse(ctx, compute.Tenant(tid), &compute.ListSkusParams{
				Accept: ptr.To(compute.ListSkusParamsAcceptApplicationjson),
			})
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *ComputeV1) GetSku(ctx context.Context, tref TenantReference) (*compute.InstanceSku, error) {
	resp, err := api.compute.GetSkuWithResponse(ctx, compute.Tenant(tref.Tenant), tref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
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
	if err := validateComputeMetadataV1(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.CreateOrUpdateInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name,
		&compute.CreateOrUpdateInstanceParams{
			IfUnmodifiedSince: &inst.Metadata.ResourceVersion,
		}, *inst)
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 200, 201); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) DeleteInstance(ctx context.Context, inst *compute.Instance) error {
	if err := validateComputeMetadataV1(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.DeleteInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, &compute.DeleteInstanceParams{
		IfUnmodifiedSince: &inst.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 202, 404); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) StartInstance(ctx context.Context, inst *compute.Instance) error {
	if err := validateComputeMetadataV1(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.StartInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, &compute.StartInstanceParams{
		IfUnmodifiedSince: &inst.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) StopInstance(ctx context.Context, inst *compute.Instance) error {
	if err := validateComputeMetadataV1(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.StopInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, &compute.StopInstanceParams{
		IfUnmodifiedSince: &inst.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) RestartInstance(ctx context.Context, inst *compute.Instance) error {
	if err := validateComputeMetadataV1(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.RestartInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, &compute.RestartInstanceParams{
		IfUnmodifiedSince: &inst.Metadata.ResourceVersion,
	})
	if err != nil {
		return err
	}

	if err = checkStatusCode(resp, 204, 404); err != nil {
		return err
	}

	return nil
}

func newComputeV1(computeUrl string) (*ComputeV1, error) {
	compute, err := compute.NewClientWithResponses(computeUrl)
	if err != nil {
		return nil, err
	}

	return &ComputeV1{compute: compute}, nil
}

func validateComputeMetadataV1(metadata *compute.ZonalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Workspace == nil {
		return ErrNoMetatadaWorkspace
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	return nil
}
