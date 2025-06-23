package secapi

import (
	"context"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
)

type WellknownV1 struct {
	wellknown wellknown.ClientWithResponsesInterface
}

func newWellknownV1(wellknownUrl string) (*WellknownV1, error) {
	wellknown, err := wellknown.NewClientWithResponses(wellknownUrl)
	if err != nil {
		return nil, err
	}

	return &WellknownV1{wellknown: wellknown}, nil
}

func (api *WellknownV1) GetWellknown(ctx context.Context) (*wellknown.Wellknown, error) {
	resp, err := api.wellknown.GetWellknownWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	return resp.JSON200, nil
}
