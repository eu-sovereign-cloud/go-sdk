package secapi

import (
	"errors"
	"testing"
	"time"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

func TestErrorFromResourceState(t *testing.T) {
	conditions := []schema.StatusCondition{{State: schema.ResourceStateError, Reason: "ProviderError", Message: "401 Unauthorized", LastTransitionAt: time.Now()}}

	if err := errorFromResourceState(schema.ResourceStateCreating, conditions); err != nil {
		t.Fatalf("non-error state: got %v, want nil", err)
	}
	if err := errorFromResourceState(schema.ResourceStateError, nil); err != ErrResourceInErrorState {
		t.Fatalf("no conditions: got %v, want %v", err, ErrResourceInErrorState)
	}
	err := errorFromResourceState(schema.ResourceStateError, conditions)
	if !errors.Is(err, ErrResourceInErrorState) || err.Error() != "resource entered error state: ProviderError: 401 Unauthorized" {
		t.Fatalf("with conditions: got %v", err)
	}
}

func TestLatestConditionMessage(t *testing.T) {
	now := time.Now()
	older := schema.StatusCondition{Message: "accepted", LastTransitionAt: now.Add(-time.Minute)}
	newer := schema.StatusCondition{Message: "401 Unauthorized", LastTransitionAt: now}

	failed := schema.StatusCondition{State: schema.ResourceStateError, Reason: "ProviderError", Message: "401 Unauthorized", LastTransitionAt: now.Add(-time.Minute)}
	retrying := schema.StatusCondition{State: schema.ResourceStateCreating, Message: "retrying", LastTransitionAt: now}

	tests := []struct {
		name       string
		conditions []schema.StatusCondition
		want       string
	}{
		{"none", nil, ""},
		{"newest first", []schema.StatusCondition{newer, older}, "401 Unauthorized"},
		{"oldest first", []schema.StatusCondition{older, newer}, "401 Unauthorized"},
		{"error condition wins over newer non-error", []schema.StatusCondition{retrying, failed}, "ProviderError: 401 Unauthorized"},
		{"reason only", []schema.StatusCondition{{Reason: "ProviderError", LastTransitionAt: now}}, "ProviderError"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := latestConditionMessage(tt.conditions); got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
