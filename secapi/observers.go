package secapi

import (
	"time"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Resource State

type ResourceStateObserver struct {
	delay       time.Duration
	interval    time.Duration
	maxAttempts int
	actFunc     func() (schema.ResourceState, error)
}

func NewResourceStateObserver(
	delay time.Duration,
	interval time.Duration,
	maxAttempts int,
	actFunc func() (schema.ResourceState, error),
) *ResourceStateObserver {
	return &ResourceStateObserver{
		delay:       delay,
		interval:    interval,
		maxAttempts: maxAttempts,
		actFunc:     actFunc,
	}
}

func (retry *ResourceStateObserver) WaitUntil(expectedState schema.ResourceState) error {
	time.Sleep(retry.delay)

	for attempt := 1; attempt <= retry.maxAttempts; attempt++ {
		state, err := retry.actFunc()
		if err != nil {
			return err
		}

		if state == expectedState {
			break
		}

		if attempt >= retry.maxAttempts {
			return ErrRetryMaxAttemptsReached
		}

		time.Sleep(retry.interval)
	}
	return nil
}

// Instance Power State

type InstancePowerStateObserver struct {
	delay       time.Duration
	interval    time.Duration
	maxAttempts int
	actFunc     func() (schema.InstanceStatusPowerState, error)
}

func NewInstancePowerStateObserver(
	delay time.Duration,
	interval time.Duration,
	maxAttempts int,
	actFunc func() (schema.InstanceStatusPowerState, error),
) *InstancePowerStateObserver {
	return &InstancePowerStateObserver{
		delay:       delay,
		interval:    interval,
		maxAttempts: maxAttempts,
		actFunc:     actFunc,
	}
}

func (retry *InstancePowerStateObserver) WaitUntil(expectedState schema.InstanceStatusPowerState) error {
	time.Sleep(retry.delay)

	for attempt := 1; attempt <= retry.maxAttempts; attempt++ {
		state, err := retry.actFunc()
		if err != nil {
			return err
		}

		if state == expectedState {
			break
		}

		if attempt >= retry.maxAttempts {
			return ErrRetryMaxAttemptsReached
		}

		time.Sleep(retry.interval)
	}
	return nil
}
