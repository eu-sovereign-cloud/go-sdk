package gosdk

import (
	regions "github.com/eu-sovereign-cloud/go-sdk/pkg/regions.v1"
)

type Client struct {
	regions regions.ClientWithResponsesInterface
}

func NewClient(regionsUrl string) (*Client, error) {
	client := &Client{}

	regionsClient, err := regions.NewClientWithResponses(regionsUrl)
	if err != nil {
		return nil, err
	}
	client.regions = regionsClient

	return client, nil
}
