package secapi

import (
	"fmt"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// URN

func BuildReferenceFromURN(urn string) (*schema.Reference, error) {
	ref := &schema.Reference{}
	if err := ref.FromReferenceURN(urn); err != nil {
		return nil, fmt.Errorf("error building reference from URN %s: %s", urn, err)
	}

	return ref, nil
}

func AsReferenceURN(ref schema.Reference) (string, error) {
	urn, err := ref.AsReferenceURN()
	if err != nil {
		return "", fmt.Errorf("error extracting URN from reference: %w", err)
	}
	return urn, nil
}
