package secapi

import (
	"fmt"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

func BuildReferenceFromURN(urn string) (*schema.Reference, error) {
	urnRef := schema.ReferenceURN(urn)

	ref := &schema.Reference{}
	if err := ref.FromReferenceURN(urnRef); err != nil {
		return nil, fmt.Errorf("error building reference from URN %s: %s", urn, err)
	}

	return ref, nil
}
