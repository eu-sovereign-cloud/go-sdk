package secapi

import "github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"

type TenantID string

type WorkspaceID string

type NetworkID string

func AcceptHeaderJson[T ~string]() *T {
	v := T(schema.AcceptHeaderJson)
	return &v
}
