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

type ResourceObserverConfig struct {
	Delay       time.Duration
	Interval    time.Duration
	MaxAttempts int
}

type ResourceObserverUntilValueConfig[T any] struct {
	ExpectedValues []T
	Delay          time.Duration
	Interval       time.Duration
	MaxAttempts    int
}

func (api *API) loadRequestHeaders(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+api.authToken)
	req.Header.Set("Accept", "application/json")
	return nil
}

func (api *API) validateGlobalMetadata(metadata *schema.GlobalTenantResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetadata
	}

	if metadata.Tenant == "" {
		return ErrNoMetadataTenant
	}

	return nil
}

func (api *API) validateRegionalMetadata(metadata *schema.RegionalResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetadata
	}

	if metadata.Tenant == "" {
		return ErrNoMetadataTenant
	}

	return nil
}

func (api *API) validateWorkspaceMetadata(metadata *schema.RegionalWorkspaceResourceMetadata) error {
	if metadata == nil {
		return ErrNoMetadata
	}

	if metadata.Tenant == "" {
		return ErrNoMetadataTenant
	}

	if metadata.Workspace == "" {
		return ErrNoMetadataWorkspace
	}

	return nil
}
