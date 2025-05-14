package secapi

import (
	region "github.com/eu-sovereign-cloud/go-sdk/pkg/foundation.region.v1"
)

type Client struct {
	regions region.ClientWithResponsesInterface
}

func NewClient(regionsUrl string) (*Client, error) {
	client := &Client{}

	regionsClient, err := region.NewClientWithResponses(regionsUrl)
	if err != nil {
		return nil, err
	}
	client.regions = regionsClient

	return client, nil
}
