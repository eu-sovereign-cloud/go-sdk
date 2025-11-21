package secapi

import (
	"context"
	"net/http"
	"time"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

type API struct {
	authToken string
}

type ResourceObserverConfig[T any] struct {
	ExpectedValue T
	Delay         time.Duration
	Interval      time.Duration
	MaxAttempts   int
}

func (api *API) loadRequestHeaders(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+api.authToken)
	req.Header.Set("Accept", "application/json")
	return nil
}

func (api *API) validateGlobalMetadata(metadata *schema.GlobalTenantResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	return nil
}

func (api *API) validateRegionalMetadata(metadata *schema.RegionalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	return nil
}

func (api *API) validateWorkspaceMetadata(metadata *schema.RegionalWorkspaceResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetatada
	}

	if metadata.Tenant == "" {
		return ErrNoMetatadaTenant
	}

	if metadata.Workspace == "" {
		return ErrNoMetatadaWorkspace
	}

	return nil
}
