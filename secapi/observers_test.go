package secapi

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

type dummyResource struct {
	state schema.ResourceState
}

func TestResourceStateObserverWaitUntil_SingleTry(t *testing.T) {
	attempts := 0
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 1,
		actFunc: func() (schema.ResourceState, *dummyResource, error) {
			attempts++
			return schema.ResourceStateActive, &dummyResource{state: schema.ResourceStateActive}, nil
		},
	}

	resp, err := observer.WaitUntil(schema.ResourceStateActive)
	assert.NoError(t, err)
	assert.Equal(t, 1, attempts)

	assert.NotNil(t, resp)
	assert.Equal(t, resp.state, schema.ResourceStateActive)
}

func TestResourceStateObserverWaitUntil_MultipleTries(t *testing.T) {
	attempts := 0
	maxAttempts := 3
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: maxAttempts,
		actFunc: func() (schema.ResourceState, *dummyResource, error) {
			attempts++
			if attempts < maxAttempts {
				return schema.ResourceStateCreating, &dummyResource{state: schema.ResourceStateCreating}, nil
			} else {
				return schema.ResourceStateActive, &dummyResource{state: schema.ResourceStateActive}, nil
			}
		},
	}

	resp, err := observer.WaitUntil(schema.ResourceStateActive)
	assert.NoError(t, err)
	assert.Equal(t, 3, attempts)

	assert.NotNil(t, resp)
	assert.Equal(t, resp.state, schema.ResourceStateActive)
}

func TestResourceStateObserverWaitUntil_MaxAttempts(t *testing.T) {
	attempts := 1
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 3,
		actFunc: func() (schema.ResourceState, *dummyResource, error) {
			attempts++
			return schema.ResourceStateCreating, &dummyResource{state: schema.ResourceStateCreating}, nil
		},
	}

	resp, err := observer.WaitUntil(schema.ResourceStateActive)
	assert.Error(t, err)
	assert.Equal(t, err, ErrRetryMaxAttemptsReached)
	assert.Equal(t, 4, attempts)

	assert.Nil(t, resp)
}

func TestResourceStateObserverWaitUntil_FuncError(t *testing.T) {
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 1,
		actFunc: func() (schema.ResourceState, *dummyResource, error) {
			return "", nil, assert.AnError
		},
	}

	resp, err := observer.WaitUntil(schema.ResourceStateActive)
	assert.Error(t, err)
	assert.Equal(t, err, assert.AnError)

	assert.Nil(t, resp)
}
