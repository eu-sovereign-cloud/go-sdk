package builders

type ListOptions struct {
	Limit  *int
	Labels *LabelsBuilder
}

const defaultLimit = 1000

func NewListOptions() *ListOptions {
	limit := defaultLimit
	return &ListOptions{
		Limit:  &limit,
		Labels: NewLabelsBuilder(),
	}
}

func (o *ListOptions) WithLimit(limit int) *ListOptions {
	o.Limit = &limit
	return o
}

func (o *ListOptions) WithLabels(labels *LabelsBuilder) *ListOptions {
	o.Labels = labels
	return o
}
