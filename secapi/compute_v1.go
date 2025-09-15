package secapi

import (
	"context"
	"fmt"
	"net/http"

	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"

	"k8s.io/utils/ptr"
)

type ComputeV1 struct {
	API
	compute compute.ClientWithResponsesInterface
}

// Instance Sku

func (api *ComputeV1) ListSkus(ctx context.Context, tid TenantID) (*Iterator[compute.InstanceSku], error) {
	iter := Iterator[compute.InstanceSku]{
		fn: func(ctx context.Context, skipToken *string) ([]compute.InstanceSku, *string, error) {
			resp, err := api.compute.ListSkusWithResponse(ctx, compute.TenantPathParam(tid), &compute.ListSkusParams{
				Accept: ptr.To(compute.ListSkusParamsAcceptApplicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *ComputeV1) GetSku(ctx context.Context, tref TenantReference) (*compute.InstanceSku, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.compute.GetSkuWithResponse(ctx, compute.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

// Instance

func (api *ComputeV1) ListInstances(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[compute.Instance], error) {
	iter := Iterator[compute.Instance]{
		fn: func(ctx context.Context, skipToken *string) ([]compute.Instance, *string, error) {
			resp, err := api.compute.ListInstancesWithResponse(ctx, compute.TenantPathParam(tid), compute.WorkspacePathParam(wid), &compute.ListInstancesParams{
				Accept: ptr.To(compute.Applicationjson),
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
		},
	}

	return &iter, nil
}

func (api *ComputeV1) GetInstance(ctx context.Context, wref WorkspaceReference) (*compute.Instance, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.compute.GetInstanceWithResponse(ctx, compute.TenantPathParam(wref.Tenant), compute.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, ErrResourceNotFound
	} else {
		return resp.JSON200, nil
	}
}

func (api *ComputeV1) CreateOrUpdateInstanceWithParams(ctx context.Context, wref WorkspaceReference, inst *compute.Instance, params *compute.CreateOrUpdateInstanceParams) (*compute.Instance, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.compute.CreateOrUpdateInstanceWithResponse(ctx, compute.TenantPathParam(wref.Tenant), compute.WorkspacePathParam(wref.Workspace), wref.Name, params, *inst, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if err = checkSuccessPutStatusCodes(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return resp.JSON201, nil
	}
}

func (api *ComputeV1) CreateOrUpdateInstance(ctx context.Context, wref WorkspaceReference, inst *compute.Instance) (*compute.Instance, error) {
	return api.CreateOrUpdateInstanceWithParams(ctx, wref, inst, nil)
}

func (api *ComputeV1) DeleteInstanceWithParams(ctx context.Context, inst *compute.Instance, params *compute.DeleteInstanceParams) error {
	if err := api.validateMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.DeleteInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessDeleteStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) DeleteInstance(ctx context.Context, inst *compute.Instance) error {
	return api.DeleteInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1) StartInstanceWithParams(ctx context.Context, inst *compute.Instance, params *compute.StartInstanceParams) error {
	if err := api.validateMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.StartInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessPostStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) StartInstance(ctx context.Context, inst *compute.Instance) error {
	return api.StartInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1) StopInstanceWithParams(ctx context.Context, inst *compute.Instance, params *compute.StopInstanceParams) error {
	if err := api.validateMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.StopInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessPostStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) StopInstance(ctx context.Context, inst *compute.Instance) error {
	return api.StopInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1) RestartInstanceWithParams(ctx context.Context, inst *compute.Instance, params *compute.RestartInstanceParams) error {
	if err := api.validateMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.RestartInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if err = checkSuccessPostStatusCodes(resp); err != nil {
		return err
	}

	return nil
}

func (api *ComputeV1) RestartInstance(ctx context.Context, inst *compute.Instance) error {
	return api.RestartInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1) BuildReferenceURN(urn string) (*compute.Reference, error) {
	urnRef := compute.ReferenceURN(urn)

	ref := &compute.Reference{}
	if err := ref.FromReferenceURN(urnRef); err != nil {
		return nil, fmt.Errorf("error building referenceURN from URN %s: %s", urn, err)
	}

	return ref, nil
}

func (api *ComputeV1) validateMetadata(metadata *compute.ZonalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == nil {
		return ErrNoMetatadaWorkspace
	}

	return nil
}

func newComputeV1(client *RegionalClient, computeUrl string) (*ComputeV1, error) {
	compute, err := compute.NewClientWithResponses(computeUrl)
	if err != nil {
		return nil, err
	}

	return &ComputeV1{API: API{authToken: client.authToken}, compute: compute}, nil
}
