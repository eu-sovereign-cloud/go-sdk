package builders

type ListOptions struct {
	Limit  *int
	Labels *string
}

const defaultLimit = 1000

func NewListOptions() *ListOptions {
	limit := defaultLimit
	return &ListOptions{
		Limit: &limit,
	}
}

func (o *ListOptions) WithLimit(limit int) *ListOptions {
	o.Limit = &limit
	return o
}

func (o *ListOptions) WithLabels(labels *LabelsBuilder) *ListOptions {
	labelsStr := labels.Build()
	o.Labels = &labelsStr
	return o
}
