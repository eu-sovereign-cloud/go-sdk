package secapi

import (
	"context"
	"net/http"

	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

	"k8s.io/utils/ptr"
)

// Interface

type ComputeV1 interface {
	// Instance Sku
	ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.InstanceSku], error)
	ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.InstanceSku], error)

	GetSku(ctx context.Context, tref TenantReference) (*schema.InstanceSku, error)

	// Instance
	ListInstances(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Instance], error)
	ListInstancesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Instance], error)

	GetInstance(ctx context.Context, wref WorkspaceReference) (*schema.Instance, error)
	GetInstanceUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Instance, error)

	CreateOrUpdateInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.CreateOrUpdateInstanceParams) (*schema.Instance, error)
	CreateOrUpdateInstance(ctx context.Context, inst *schema.Instance) (*schema.Instance, error)

	DeleteInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.DeleteInstanceParams) error
	DeleteInstance(ctx context.Context, inst *schema.Instance) error

	StartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StartInstanceParams) error
	StartInstance(ctx context.Context, inst *schema.Instance) error

	StopInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StopInstanceParams) error
	StopInstance(ctx context.Context, inst *schema.Instance) error

	RestartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.RestartInstanceParams) error
	RestartInstance(ctx context.Context, inst *schema.Instance) error
}

// Dummy

type ComputeV1Dummy struct{}

func newComputeV1Dummy() ComputeV1 {
	return &ComputeV1Dummy{}
}

/// Instance Sku

func (api *ComputeV1Dummy) ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.InstanceSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.InstanceSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) GetSku(ctx context.Context, tref TenantReference) (*schema.InstanceSku, error) {
	return nil, ErrProviderNotAvailable
}

/// Instance

func (api *ComputeV1Dummy) ListInstances(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Instance], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) ListInstancesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Instance], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) GetInstance(ctx context.Context, wref WorkspaceReference) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) GetInstanceUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) CreateOrUpdateInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.CreateOrUpdateInstanceParams) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) CreateOrUpdateInstance(ctx context.Context, inst *schema.Instance) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) DeleteInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.DeleteInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) DeleteInstance(ctx context.Context, inst *schema.Instance) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) StartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StartInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) StartInstance(ctx context.Context, inst *schema.Instance) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) StopInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StopInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) StopInstance(ctx context.Context, inst *schema.Instance) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) RestartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.RestartInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Dummy) RestartInstance(ctx context.Context, inst *schema.Instance) error {
	return ErrProviderNotAvailable
}

// Impl

type ComputeV1Impl struct {
	API
	compute compute.ClientWithResponsesInterface
}

func newComputeV1Impl(client *RegionalClient, computeUrl string) (ComputeV1, error) {
	compute, err := compute.NewClientWithResponses(computeUrl)
	if err != nil {
		return nil, err
	}

	return &ComputeV1Impl{API: API{authToken: client.authToken}, compute: compute}, nil
}

// Instance Sku

func (api *ComputeV1Impl) ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.InstanceSku], error) {
	iter := Iterator[schema.InstanceSku]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.InstanceSku, *string, error) {
			resp, err := api.compute.ListSkusWithResponse(ctx, schema.TenantPathParam(tid), &compute.ListSkusParams{
				Accept:    ptr.To(compute.ListSkusParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *ComputeV1Impl) ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.InstanceSku], error) {
	iter := Iterator[schema.InstanceSku]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.InstanceSku, *string, error) {
			resp, err := api.compute.ListSkusWithResponse(ctx, schema.TenantPathParam(tid), &compute.ListSkusParams{
				Accept:    ptr.To(compute.ListSkusParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *ComputeV1Impl) GetSku(ctx context.Context, tref TenantReference) (*schema.InstanceSku, error) {
	if err := tref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.compute.GetSkuWithResponse(ctx, schema.TenantPathParam(tref.Tenant), tref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

// Instance

func (api *ComputeV1Impl) ListInstances(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Instance], error) {
	iter := Iterator[schema.Instance]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Instance, *string, error) {
			resp, err := api.compute.ListInstancesWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &compute.ListInstancesParams{
				Accept:    ptr.To(compute.ListInstancesParamsAccept(schema.AcceptHeaderJson)),
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *ComputeV1Impl) ListInstancesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Instance], error) {
	iter := Iterator[schema.Instance]{
		fn: func(ctx context.Context, skipToken *string) ([]schema.Instance, *string, error) {
			resp, err := api.compute.ListInstancesWithResponse(ctx, schema.TenantPathParam(tid), schema.WorkspacePathParam(wid), &compute.ListInstancesParams{
				Accept:    ptr.To(compute.ListInstancesParamsAccept(schema.AcceptHeaderJson)),
				Labels:    opts.Labels.BuildPtr(),
				Limit:     opts.Limit,
				SkipToken: skipToken,
			}, api.loadRequestHeaders)
			if err != nil {
				return nil, nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return resp.JSON200.Items, resp.JSON200.Metadata.SkipToken, nil
			} else {
				return nil, nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	return &iter, nil
}

func (api *ComputeV1Impl) GetInstance(ctx context.Context, wref WorkspaceReference) (*schema.Instance, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	resp, err := api.compute.GetInstanceWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) GetInstanceUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig[schema.ResourceState]) (*schema.Instance, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Instance]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		actFunc: func() (schema.ResourceState, *schema.Instance, error) {
			resp, err := api.compute.GetInstanceWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if resp.StatusCode() == http.StatusOK {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntil(config.ExpectedValue)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *ComputeV1Impl) CreateOrUpdateInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.CreateOrUpdateInstanceParams) (*schema.Instance, error) {
	if err := api.validateWorkspaceMetadata(inst.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.compute.CreateOrUpdateInstanceWithResponse(ctx, inst.Metadata.Tenant, inst.Metadata.Workspace, inst.Metadata.Name, params, *inst, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else if resp.StatusCode() == http.StatusCreated {
		return resp.JSON201, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) CreateOrUpdateInstance(ctx context.Context, inst *schema.Instance) (*schema.Instance, error) {
	return api.CreateOrUpdateInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1Impl) DeleteInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.DeleteInstanceParams) error {
	if err := api.validateWorkspaceMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.DeleteInstanceWithResponse(ctx, inst.Metadata.Tenant, inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) DeleteInstance(ctx context.Context, inst *schema.Instance) error {
	return api.DeleteInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1Impl) StartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StartInstanceParams) error {
	if err := api.validateWorkspaceMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.StartInstanceWithResponse(ctx, inst.Metadata.Tenant, inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) StartInstance(ctx context.Context, inst *schema.Instance) error {
	return api.StartInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1Impl) StopInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StopInstanceParams) error {
	if err := api.validateWorkspaceMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.StopInstanceWithResponse(ctx, inst.Metadata.Tenant, inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) StopInstance(ctx context.Context, inst *schema.Instance) error {
	return api.StopInstanceWithParams(ctx, inst, nil)
}

func (api *ComputeV1Impl) RestartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.RestartInstanceParams) error {
	if err := api.validateWorkspaceMetadata(inst.Metadata); err != nil {
		return err
	}

	resp, err := api.compute.RestartInstanceWithResponse(ctx, inst.Metadata.Tenant, inst.Metadata.Workspace, inst.Metadata.Name, params, api.loadRequestHeaders)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusAccepted {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) RestartInstance(ctx context.Context, inst *schema.Instance) error {
	return api.RestartInstanceWithParams(ctx, inst, nil)
}
