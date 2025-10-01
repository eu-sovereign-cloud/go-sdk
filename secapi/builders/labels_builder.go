package builders

import (
	"fmt"
	"strings"
)

type LabelsBuilder struct {
	Items []string
}

func NewLabelsBuilder() *LabelsBuilder {
	return &LabelsBuilder{}
}

func (b *LabelsBuilder) Equals(key, value string) *LabelsBuilder {
	b.Items = append(b.Items, fmt.Sprintf("%s=%s", key, value))
	return b
}

func (b *LabelsBuilder) NsEquals(namespace, key, value string) *LabelsBuilder {
	b.Items = append(b.Items, fmt.Sprintf("%s:%s=%s", namespace, key, value))
	return b
}

func (b *LabelsBuilder) Neq(key, value string) *LabelsBuilder {
	b.Items = append(b.Items, fmt.Sprintf("%s!=%s", key, value))
	return b
}

func (b *LabelsBuilder) Gt(key string, value int) *LabelsBuilder {
	b.Items = append(b.Items, fmt.Sprintf("%s>%d", key, value))
	return b
}

func (b *LabelsBuilder) Lt(key string, value int) *LabelsBuilder {
	b.Items = append(b.Items, fmt.Sprintf("%s<%d", key, value))
	return b
}

func (b *LabelsBuilder) Gte(key string, value int) *LabelsBuilder {
	b.Items = append(b.Items, fmt.Sprintf("%s>=%d", key, value))
	return b
}

func (b *LabelsBuilder) Lte(key string, value int) *LabelsBuilder {
	b.Items = append(b.Items, fmt.Sprintf("%s<=%d", key, value))
	return b
}

func (b *LabelsBuilder) Build() string {
	return strings.Join(b.Items, ",")
}
