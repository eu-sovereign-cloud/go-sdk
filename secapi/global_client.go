package secapi

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/foundation.region.v1"
)

type GlobalClient struct {
	regions region.ClientWithResponsesInterface
}

func NewGlobalClient(regionsUrl string) (*GlobalClient, error) {
	client := &GlobalClient{}

	regionsClient, err := region.NewClientWithResponses(regionsUrl)
	if err != nil {
		return nil, err
	}
	client.regions = regionsClient

	return client, nil
}
