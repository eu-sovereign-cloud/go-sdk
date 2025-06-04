package secapi

import "errors"

var ErrNoTenantID = errors.New("tenant ID not found in context")

var ErrNoMetatada = errors.New("metadata is empty")

var ErrNoMetatadaWorkspace = errors.New("metadata workspace is empty")

var ErrNoMetatadaTenant = errors.New("metadata tenant is empty")
