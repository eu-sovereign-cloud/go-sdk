package secapi

import (
	"context"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secapi"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
)

type WellknownV1 struct {
	secapi.GlobalAPI[wellknown.ClientWithResponsesInterface]
}

func newWellknownV1() *WellknownV1 {
	return &WellknownV1{
		GlobalAPI: secapi.GlobalAPI[wellknown.ClientWithResponsesInterface]{},
	}
}

func (api *WellknownV1) getClient() (wellknown.ClientWithResponsesInterface, error) {
	fn := func(url string) (wellknown.ClientWithResponsesInterface, error) {
		return wellknown.NewClientWithResponses(url)
	}

	client, err := api.GetClient("seca.workspace", fn)
	if err != nil {
		return nil, err
	}
	return *client, nil
}

func (api *WellknownV1) GetWellknown(ctx context.Context) (*wellknown.Wellknown, error) {
	client, err := api.getClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GetWellknownWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}
