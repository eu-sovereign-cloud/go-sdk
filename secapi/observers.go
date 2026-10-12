package secapi

import (
	"slices"
	"time"

	"github.com/cenkalti/backoff/v4"
)

type resourceStateObserver[V comparable, R any] struct {
	delay        time.Duration
	interval     time.Duration
	maxAttempts  int
	getValueFunc func() (V, *R, error)
	getErrorFunc func() error

	// checkValueFunc, when set, inspects a response that did not match the expected values and
	// returns a non-nil error if the resource will never reach them (e.g. it is in an error
	// state). WaitUntilValue then stops immediately with that error, instead of retrying until
	// maxAttempts is exhausted and masking the real failure behind ErrRetryMaxAttemptsReached.
	checkValueFunc func(*R) error
}

func (retry *resourceStateObserver[V, R]) WaitUntilValue(expectedValues []V) (*R, error) {
	be := backoff.NewExponentialBackOff()
	be.InitialInterval = retry.interval
	be.MaxElapsedTime = 0

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

		if retry.checkValueFunc != nil {
			if err := retry.checkValueFunc(resp); err != nil {
				return nil, backoff.Permanent(err)
			}
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
	be.MaxElapsedTime = 0

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
