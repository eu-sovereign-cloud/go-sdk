package secapi

import "github.com/eu-sovereign-cloud/go-sdk/pkg/secalib/builders"

type ListOptions struct {
	Limit  *int
	Labels *builders.LabelsBuilder
}

const defaultLimit = 1000

func NewListOptions() *ListOptions {
	limit := defaultLimit
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
