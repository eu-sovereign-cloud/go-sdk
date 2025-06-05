package secapi

import (
	"context"

	"k8s.io/utils/ptr"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secapi"
	compute "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.compute.v1"
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type ComputeV1 struct {
	secapi.RegionalAPI[compute.ClientWithResponsesInterface]
}

func newComputeV1(region *region.Region) *ComputeV1 {
	return &ComputeV1{
		RegionalAPI: secapi.RegionalAPI[compute.ClientWithResponsesInterface]{Region: region},
	}
}

func (api *ComputeV1) getClient() (compute.ClientWithResponsesInterface, error) {
	fn := func(url string) (compute.ClientWithResponsesInterface, error) {
		return compute.NewClientWithResponses(url)
	}

	client, err := api.GetClient("seca.compute", fn)
	if err != nil {
		return nil, err
	}
	return *client, nil
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
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	iter := Iterator[compute.Instance]{
		fn: func(ctx context.Context, skipToken *string) ([]compute.Instance, *string, error) {
			resp, err := client.ListInstancesWithResponse(ctx, compute.Tenant(tid), compute.Workspace(wid), &compute.ListInstancesParams{
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
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GetInstanceWithResponse(ctx, compute.Tenant(wref.Tenant), compute.Workspace(wref.Workspace), wref.Name)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}

func (api *ComputeV1) CreateOrUpdateInstance(ctx context.Context, inst *compute.Instance) error {
	validateComputeMetadataV1(inst.Metadata)

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.CreateOrUpdateInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name,
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

	client, err := api.getClient()
	if err != nil {
		return err
	}

	resp, err := client.DeleteInstanceWithResponse(ctx, inst.Metadata.Tenant, *inst.Metadata.Workspace, inst.Metadata.Name, &compute.DeleteInstanceParams{
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
