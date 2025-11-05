package secapi

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Resource State

func TestResourceStateObserverWaitUntil_SingleTry(t *testing.T) {
	attempts := 0
	observer := NewResourceStateObserver(
		0,
		0,
		1,
		func() (schema.ResourceState, error) {
			attempts++
			return schema.ResourceStateActive, nil
		},
	)

	err := observer.WaitUntil(schema.ResourceStateActive)
	assert.NoError(t, err)
	assert.Equal(t, 1, attempts)
}

func TestResourceStateObserverWaitUntil_MultipleTries(t *testing.T) {
	attempts := 0
	maxAttempts := 3
	observer := NewResourceStateObserver(
		0,
		0,
		maxAttempts,
		func() (schema.ResourceState, error) {
			attempts++
			if attempts < maxAttempts {
				return schema.ResourceStateCreating, nil
			} else {
				return schema.ResourceStateActive, nil
			}
		},
	)

	err := observer.WaitUntil(schema.ResourceStateActive)
	assert.NoError(t, err)
	assert.Equal(t, 3, attempts)
}

func TestResourceStateObserverWaitUntil_MaxAttempts(t *testing.T) {
	attempts := 1
	observer := NewResourceStateObserver(
		0,
		0,
		3,
		func() (schema.ResourceState, error) {
			attempts++
			return schema.ResourceStateCreating, nil
		},
	)

	err := observer.WaitUntil(schema.ResourceStateActive)
	assert.Error(t, err)
	assert.Equal(t, err, ErrRetryMaxAttemptsReached)
	assert.Equal(t, 4, attempts)
}

func TestResourceStateObserverWaitUntil_FuncError(t *testing.T) {
	observer := NewResourceStateObserver(
		0,
		0,
		1,
		func() (schema.ResourceState, error) {
			return "", assert.AnError
		},
	)

	err := observer.WaitUntil(schema.ResourceStateActive)
	assert.Error(t, err)
	assert.Equal(t, err, assert.AnError)
}

// Instance Power State

func TestInstancePowerStateObserverWaitUntil_SingleTry(t *testing.T) {
	attempts := 0
	observer := NewInstancePowerStateObserver(
		0,
		0,
		1,
		func() (schema.InstanceStatusPowerState, error) {
			attempts++
			return schema.InstanceStatusPowerStateOn, nil
		},
	)

	err := observer.WaitUntil(schema.InstanceStatusPowerStateOn)
	assert.NoError(t, err)
	assert.Equal(t, 1, attempts)
}

func TestInstancePowerStateObserverWaitUntil_MultipleTries(t *testing.T) {
	attempts := 0
	maxAttempts := 3
	observer := NewInstancePowerStateObserver(
		0,
		0,
		maxAttempts,
		func() (schema.InstanceStatusPowerState, error) {
			attempts++
			if attempts < maxAttempts {
				return schema.InstanceStatusPowerStateOff, nil
			} else {
				return schema.InstanceStatusPowerStateOn, nil
			}
		},
	)

	err := observer.WaitUntil(schema.InstanceStatusPowerStateOn)
	assert.NoError(t, err)
	assert.Equal(t, 3, attempts)
}

func TestInstancePowerStateObserverWaitUntil_MaxAttempts(t *testing.T) {
	attempts := 1
	observer := NewInstancePowerStateObserver(
		0,
		0,
		3,
		func() (schema.InstanceStatusPowerState, error) {
			attempts++
			return schema.InstanceStatusPowerStateOff, nil
		},
	)

	err := observer.WaitUntil(schema.InstanceStatusPowerStateOn)
	assert.Error(t, err)
	assert.Equal(t, err, ErrRetryMaxAttemptsReached)
	assert.Equal(t, 4, attempts)
}

func TestInstancePowerStateObserverWaitUntil_FuncError(t *testing.T) {
	observer := NewInstancePowerStateObserver(
		0,
		0,
		1,
		func() (schema.InstanceStatusPowerState, error) {
			return "", assert.AnError
		},
	)

	err := observer.WaitUntil(schema.InstanceStatusPowerStateOn)
	assert.Error(t, err)
	assert.Equal(t, err, assert.AnError)
}
