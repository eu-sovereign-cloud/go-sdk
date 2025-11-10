package secapi

import (
	"time"
)

type resourceStateObserver[V comparable, R any] struct {
	delay       time.Duration
	interval    time.Duration
	maxAttempts int
	actFunc     func() (V, *R, error)
}

func (retry *resourceStateObserver[V, R]) WaitUntil(expectedValue V) (*R, error) {
	time.Sleep(retry.delay)

	for attempt := 1; attempt <= retry.maxAttempts; attempt++ {
		value, resp, err := retry.actFunc()
		if err != nil {
			return resp, err
		}

		if value == expectedValue {
			return resp, nil
		}

		if attempt >= retry.maxAttempts {
			return nil, ErrRetryMaxAttemptsReached
		}

		time.Sleep(retry.interval)
	}
	return nil, ErrRetryMaxAttemptsReached
}
