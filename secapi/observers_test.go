package secapi

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

type dummyResource struct {
	state schema.ResourceState
}

// WaitUntilValue

func TestResourceStateObserverWaitUntilValue_SingleTry(t *testing.T) {
	attempts := 0
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 1,
		getValueFunc: func() (schema.ResourceState, *dummyResource, error) {
			attempts++
			return schema.ResourceStateActive, &dummyResource{state: schema.ResourceStateActive}, nil
		},
	}

	resp, err := observer.WaitUntilValue([]schema.ResourceState{schema.ResourceStateActive})
	assert.NoError(t, err)
	assert.Equal(t, 1, attempts)

	assert.NotNil(t, resp)
	assert.Equal(t, resp.state, schema.ResourceStateActive)
}

func TestResourceStateObserverWaitUntilValue_MultipleTries(t *testing.T) {
	attempts := 0
	maxAttempts := 3
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: maxAttempts,
		getValueFunc: func() (schema.ResourceState, *dummyResource, error) {
			attempts++
			if attempts < maxAttempts {
				return schema.ResourceStateCreating, &dummyResource{state: schema.ResourceStateCreating}, nil
			} else {
				return schema.ResourceStateActive, &dummyResource{state: schema.ResourceStateActive}, nil
			}
		},
	}

	resp, err := observer.WaitUntilValue([]schema.ResourceState{schema.ResourceStateActive})
	assert.NoError(t, err)
	assert.Equal(t, 3, attempts)

	assert.NotNil(t, resp)
	assert.Equal(t, resp.state, schema.ResourceStateActive)
}

func TestResourceStateObserverWaitUntilValue_MultipleExpectedValues(t *testing.T) {
	attempts := 0
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 1,
		getValueFunc: func() (schema.ResourceState, *dummyResource, error) {
			attempts++

			states := []schema.ResourceState{schema.ResourceStatePending, schema.ResourceStateActive}
			randomState := states[rand.Intn(len(states))]
			return schema.ResourceStateActive, &dummyResource{state: randomState}, nil
		},
	}

	resp, err := observer.WaitUntilValue([]schema.ResourceState{schema.ResourceStatePending, schema.ResourceStateActive})
	assert.NoError(t, err)
	assert.Equal(t, 1, attempts)

	assert.NotNil(t, resp)
	assert.Contains(t, []schema.ResourceState{schema.ResourceStatePending, schema.ResourceStateActive}, resp.state)
}

func TestResourceStateObserverWaitUntilValue_MaxAttempts(t *testing.T) {
	attempts := 1
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 3,
		getValueFunc: func() (schema.ResourceState, *dummyResource, error) {
			attempts++
			return schema.ResourceStateCreating, &dummyResource{state: schema.ResourceStateCreating}, nil
		},
	}

	resp, err := observer.WaitUntilValue([]schema.ResourceState{schema.ResourceStateActive})
	assert.Error(t, err)
	assert.Equal(t, err, ErrRetryMaxAttemptsReached)
	assert.Equal(t, 4, attempts)

	assert.Nil(t, resp)
}

func TestResourceStateObserverWaitUntilValue_Error(t *testing.T) {
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 1,
		getValueFunc: func() (schema.ResourceState, *dummyResource, error) {
			return "", nil, assert.AnError
		},
	}

	resp, err := observer.WaitUntilValue([]schema.ResourceState{schema.ResourceStateActive})
	assert.Error(t, err)
	assert.Equal(t, err, assert.AnError)

	assert.Nil(t, resp)
}

// WaitUntilError

func TestResourceStateObserverWaitUntilError_SingleTry(t *testing.T) {
	attempts := 0
	observer := resourceStateObserver[any, any]{
		delay:       0,
		interval:    0,
		maxAttempts: 1,
		getErrorFunc: func() error {
			attempts++
			return ErrResourceNotFound
		},
	}

	resp, err := observer.WaitUntilError(ErrResourceNotFound)
	assert.NoError(t, err)
	assert.Equal(t, 1, attempts)

	assert.NotNil(t, resp)
	assert.Equal(t, resp, ErrResourceNotFound)
}

func TestResourceStateObserverWaitUntilError_MultipleTries(t *testing.T) {
	attempts := 0
	maxAttempts := 5
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: maxAttempts,
		getErrorFunc: func() error {
			attempts++
			if attempts < maxAttempts {
				return nil
			} else {
				return ErrResourceNotFound
			}
		},
	}

	resp, err := observer.WaitUntilError(ErrResourceNotFound)
	assert.NoError(t, err)
	assert.Equal(t, 5, attempts)

	assert.NotNil(t, resp)
	assert.Equal(t, resp, ErrResourceNotFound)
}

func TestResourceStateObserverWaitUntilError_MaxAttempts(t *testing.T) {
	attempts := 1
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 3,
		getErrorFunc: func() error {
			attempts++
			return nil
		},
	}

	resp, err := observer.WaitUntilError(ErrResourceNotFound)
	assert.Error(t, err)
	assert.Equal(t, err, ErrRetryMaxAttemptsReached)
	assert.Equal(t, 4, attempts)

	assert.Nil(t, resp)
}

func TestResourceStateObserverWaitUntilError_UnexpectedError(t *testing.T) {
	observer := resourceStateObserver[schema.ResourceState, dummyResource]{
		delay:       0,
		interval:    0,
		maxAttempts: 1,
		getErrorFunc: func() error {
			return assert.AnError
		},
	}

	resp, err := observer.WaitUntilError(ErrResourceNotFound)
	assert.Error(t, err)
	assert.Equal(t, err, assert.AnError)

	assert.Nil(t, resp)
}
