package secapi

import (
	"context"
	"fmt"
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

// errorFromResourceState returns nil unless state is schema.ResourceStateError, in which case
// it returns ErrResourceInErrorState, annotated with the reason from conditions when there is one.
func errorFromResourceState(state schema.ResourceState, conditions []schema.StatusCondition) error {
	if state != schema.ResourceStateError {
		return nil
	}
	if msg := latestConditionMessage(conditions); msg != "" {
		return fmt.Errorf("%w: %s", ErrResourceInErrorState, msg)
	}
	return ErrResourceInErrorState
}

// latestConditionMessage describes why a resource is in an error state, using its status
// conditions. It prefers the newest condition whose State is error, so a later non-error
// condition (e.g. a retry) can't hide the actual failure, and falls back to the newest
// condition overall. The result is "<Reason>: <Message>", or whichever of the two is set,
// or "" if there are no conditions.
//
// The newest condition is chosen by LastTransitionAt rather than by position, because the
// order the server returns conditions in is not part of the contract (ECP returns them
// newest-first). Ties keep the earliest entry in the slice.
func latestConditionMessage(conditions []schema.StatusCondition) string {
	var latest, latestErr *schema.StatusCondition
	for i := range conditions {
		c := &conditions[i]
		if latest == nil || c.LastTransitionAt.After(latest.LastTransitionAt) {
			latest = c
		}
		if c.State == schema.ResourceStateError && (latestErr == nil || c.LastTransitionAt.After(latestErr.LastTransitionAt)) {
			latestErr = c
		}
	}

	chosen := latestErr
	if chosen == nil {
		chosen = latest
	}
	if chosen == nil {
		return ""
	}

	switch {
	case chosen.Reason != "" && chosen.Message != "":
		return chosen.Reason + ": " + chosen.Message
	case chosen.Reason != "":
		return chosen.Reason
	default:
		return chosen.Message
	}
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
