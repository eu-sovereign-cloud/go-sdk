package secapi

import (
	"context"

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
	GetInstanceUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Instance, error)

	WatchInstanceUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error

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

// Unavailable

type ComputeV1Unavailable struct{}

func newComputeV1Unavailable() ComputeV1 {
	return &ComputeV1Unavailable{}
}

/// Instance Sku

func (api *ComputeV1Unavailable) ListSkus(ctx context.Context, tid TenantID) (*Iterator[schema.InstanceSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) ListSkusWithFilters(ctx context.Context, tid TenantID, opts *ListOptions) (*Iterator[schema.InstanceSku], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) GetSku(ctx context.Context, tref TenantReference) (*schema.InstanceSku, error) {
	return nil, ErrProviderNotAvailable
}

/// Instance

func (api *ComputeV1Unavailable) ListInstances(ctx context.Context, tid TenantID, wid WorkspaceID) (*Iterator[schema.Instance], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) ListInstancesWithFilters(ctx context.Context, tid TenantID, wid WorkspaceID, opts *ListOptions) (*Iterator[schema.Instance], error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) GetInstance(ctx context.Context, wref WorkspaceReference) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) GetInstanceUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) WatchInstanceUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) CreateOrUpdateInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.CreateOrUpdateInstanceParams) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) CreateOrUpdateInstance(ctx context.Context, inst *schema.Instance) (*schema.Instance, error) {
	return nil, ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) DeleteInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.DeleteInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) DeleteInstance(ctx context.Context, inst *schema.Instance) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) StartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StartInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) StartInstance(ctx context.Context, inst *schema.Instance) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) StopInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.StopInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) StopInstance(ctx context.Context, inst *schema.Instance) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) RestartInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.RestartInstanceParams) error {
	return ErrProviderNotAvailable
}

func (api *ComputeV1Unavailable) RestartInstance(ctx context.Context, inst *schema.Instance) error {
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

			if checkSuccessGetStatusCode(resp.StatusCode()) {
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

			if checkSuccessGetStatusCode(resp.StatusCode()) {
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

	if checkSuccessGetStatusCode(resp.StatusCode()) {
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

			if checkSuccessGetStatusCode(resp.StatusCode()) {
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

			if checkSuccessGetStatusCode(resp.StatusCode()) {
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

	if checkSuccessGetStatusCode(resp.StatusCode()) {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) GetInstanceUntilState(ctx context.Context, wref WorkspaceReference, config ResourceObserverUntilValueConfig[schema.ResourceState]) (*schema.Instance, error) {
	if err := wref.validate(); err != nil {
		return nil, err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Instance]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getValueFunc: func() (schema.ResourceState, *schema.Instance, error) {
			resp, err := api.compute.GetInstanceWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return "", nil, err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return *resp.JSON200.Status.State, resp.JSON200, nil
			} else {
				return "", nil, mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	resp, err := observer.WaitUntilValue(config.ExpectedValues)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func (api *ComputeV1Impl) WatchInstanceUntilDeleted(ctx context.Context, wref WorkspaceReference, config ResourceObserverConfig) error {
	if err := wref.validate(); err != nil {
		return err
	}

	observer := resourceStateObserver[schema.ResourceState, schema.Instance]{
		delay:       config.Delay,
		interval:    config.Interval,
		maxAttempts: config.MaxAttempts,
		getErrorFunc: func() error {
			resp, err := api.compute.GetInstanceWithResponse(ctx, schema.TenantPathParam(wref.Tenant), schema.WorkspacePathParam(wref.Workspace), wref.Name, api.loadRequestHeaders)
			if err != nil {
				return err
			}

			if checkSuccessGetStatusCode(resp.StatusCode()) {
				return nil
			} else {
				return mapStatusCodeToError(resp.StatusCode())
			}
		},
	}

	_, err := observer.WaitUntilError(ErrResourceNotFound)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (api *ComputeV1Impl) CreateOrUpdateInstanceWithParams(ctx context.Context, inst *schema.Instance, params *compute.CreateOrUpdateInstanceParams) (*schema.Instance, error) {
	if err := api.validateWorkspaceMetadata(inst.Metadata); err != nil {
		return nil, err
	}

	resp, err := api.compute.CreateOrUpdateInstanceWithResponse(ctx, inst.Metadata.Tenant, inst.Metadata.Workspace, inst.Metadata.Name, params, *inst, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if valid, json := checkSuccessPutStatusCode(resp.StatusCode(), resp.JSON201, resp.JSON200); valid {
		return json, nil
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

	if checkSuccessDeleteStatusCode(resp.StatusCode()) {
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

	if checkSuccessPostStatusCode(resp.StatusCode()) {
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

	if checkSuccessPostStatusCode(resp.StatusCode()) {
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

	if checkSuccessPostStatusCode(resp.StatusCode()) {
		return nil
	} else {
		return mapStatusCodeToError(resp.StatusCode())
	}
}

func (api *ComputeV1Impl) RestartInstance(ctx context.Context, inst *schema.Instance) error {
	return api.RestartInstanceWithParams(ctx, inst, nil)
}
