package secapi

import (
	"context"
	"net/http"

	wellknown "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/extensions.wellknown.v1"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Interface

type WellknownV1 interface {
	GetWellknown(ctx context.Context) (*schema.Wellknown, error)
}

// Dummy

type WellknownV1Dummy struct{}

func newWellknownV1Dummy() WellknownV1 {
	return &WellknownV1Dummy{}
}

func (api *WellknownV1Dummy) GetWellknown(ctx context.Context) (*schema.Wellknown, error) {
	return nil, ErrProviderNotAvailable
}

// Impl

type WellknownV1Impl struct {
	API
	wellknown wellknown.ClientWithResponsesInterface
}

func newWellknownV1Impl(client *GlobalClient, wellknownUrl string) (WellknownV1, error) {
	wellknown, err := wellknown.NewClientWithResponses(wellknownUrl)
	if err != nil {
		return nil, err
	}

	return &WellknownV1Impl{API: API{authToken: client.authToken}, wellknown: wellknown}, nil
}

func (api *WellknownV1Impl) GetWellknown(ctx context.Context) (*schema.Wellknown, error) {
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
