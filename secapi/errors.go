package secapi

import "errors"

var ErrNoTenantID = errors.New("tenant ID not found in context")
