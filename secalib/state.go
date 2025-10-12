package secalib

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"k8s.io/utils/ptr"
)

func BuildResponseResourceState(state string) *schema.ResourceState {
	return ptr.To(schema.ResourceState(state))
}
