package secapi

import (
	"slices"
	"time"

	"github.com/cenkalti/backoff/v4"
)

const RETRY_MULTIPLIER = 0.5

type resourceStateObserver[V comparable, R any] struct {
	delay        time.Duration
	interval     time.Duration
	maxAttempts  int
	getValueFunc func() (V, *R, error)
	getErrorFunc func() error
}

func (retry *resourceStateObserver[V, R]) WaitUntilValue(expectedValues []V) (*R, error) {
	be := backoff.NewExponentialBackOff()
	be.InitialInterval = retry.interval
	be.Multiplier = RETRY_MULTIPLIER

	attempt := 0
	operation := func() (*R, error) {
		attempt++

		value, resp, err := retry.getValueFunc()
		if err != nil {
			// Stop to try if it returns an error
			return nil, backoff.Permanent(err)
		}

		if slices.Contains(expectedValues, value) {
			// Stop to try and returns the response
			return resp, nil
		}

		if attempt >= retry.maxAttempts {
			return nil, backoff.Permanent(ErrRetryMaxAttemptsReached)
		}

		// Try again
		return nil, ErrRetryNotFoundExpectedValue
	}

	// Wait to start to try
	time.Sleep(retry.delay)

	resp, err := backoff.RetryWithData(operation, be)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (retry *resourceStateObserver[V, R]) WaitUntilError(expectedError error) (error, error) {
	be := backoff.NewExponentialBackOff()
	be.InitialInterval = retry.interval
	be.Multiplier = RETRY_MULTIPLIER

	attempt := 0
	operation := func() (error, error) {
		attempt++

		err := retry.getErrorFunc()
		if err != nil {
			if err == expectedError {
				// Stop to try and returns the expected error
				return err, nil
			} else {
				// Stop to try if it returns an unexpected error
				return nil, backoff.Permanent(err)
			}
		}

		if attempt >= retry.maxAttempts {
			return nil, backoff.Permanent(ErrRetryMaxAttemptsReached)
		}

		// Try again
		return nil, ErrRetryNotFoundExpectedError
	}

	// Wait to start to try
	time.Sleep(retry.delay)

	resp, err := backoff.RetryWithData(operation, be)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
