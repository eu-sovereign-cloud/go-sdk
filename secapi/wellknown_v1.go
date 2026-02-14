package secapi

import (
	"context"
	"net/http"

	wellknown "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

type WellknownV1 struct {
	API
	wellknown wellknown.ClientWithResponsesInterface
}

// Wellknown

func (api *WellknownV1) GetWellknown(ctx context.Context) (*schema.Wellknown, error) {
	resp, err := api.wellknown.GetWellknownWithResponse(ctx, api.loadRequestHeaders)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.JSON200, nil
	} else {
		return nil, mapStatusCodeToError(resp.StatusCode())
	}
}

func newWellknownV1(client *GlobalClient, wellknownUrl string) (*WellknownV1, error) {
	wellknown, err := wellknown.NewClientWithResponses(wellknownUrl)
	if err != nil {
		return nil, err
	}

	return &WellknownV1{API: API{authToken: client.authToken}, wellknown: wellknown}, nil
}
