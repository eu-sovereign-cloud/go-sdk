package secapi

import "github.com/eu-sovereign-cloud/go-sdk/secapi/builders"

// Options

const defaultListLimit = 1000

type ListOptions struct {
	Limit  *int
	Labels *builders.LabelsBuilder
}

func NewListOptions() *ListOptions {
	limit := defaultListLimit
	return &ListOptions{
		Limit:  &limit,
		Labels: builders.NewLabelsBuilder(),
	}
}

func (o *ListOptions) WithLimit(limit int) *ListOptions {
	o.Limit = &limit
	return o
}

func (o *ListOptions) WithLabels(labels *builders.LabelsBuilder) *ListOptions {
	o.Labels = labels
	return o
}
