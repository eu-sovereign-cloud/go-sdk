package secapi

import (
	"errors"
	"time"

	"github.com/cenkalti/backoff/v4"
)

const RETRY_MULTIPLIER = 0.5

type resourceStateObserver[V comparable, R any] struct {
	delay       time.Duration
	interval    time.Duration
	maxAttempts int
	actFunc     func() (V, *R, error)
}

func (retry *resourceStateObserver[V, R]) WaitUntil(expectedValue V) (*R, error) {
	be := backoff.NewExponentialBackOff()
	be.InitialInterval = retry.interval
	be.Multiplier = RETRY_MULTIPLIER

	attempt := 0
	operation := func() (*R, error) {
		attempt++

		value, resp, err := retry.actFunc()
		if err != nil {
			// Stop to try if it returns an error
			return nil, backoff.Permanent(err)
		}

		if value == expectedValue {
			// Stop to try and returns the response
			return resp, nil
		}

		if attempt >= retry.maxAttempts {
			return nil, backoff.Permanent(ErrRetryMaxAttemptsReached)
		}

		// Try again
		return nil, errors.New("Not found the expected value")
	}

	// Wait to start to try
	time.Sleep(retry.delay)

	resp, err := backoff.RetryWithData(operation, be)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
