package stability

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type CircuitBreakerBackOff struct {
	CurrentState          State
	ThresholdTimeout      time.Duration
	lastFailureTime       time.Time
	FailureCount          int
	FailureCountThreshold int
	mu                    sync.Mutex
}

func NewCircuitBreakerBackOff(thresholdTimeout time.Duration, failureCountThreshold int) *CircuitBreakerBackOff {
	return &CircuitBreakerBackOff{
		CurrentState:          Closed,
		ThresholdTimeout:      thresholdTimeout,
		FailureCountThreshold: failureCountThreshold,
	}
}

func (cbb *CircuitBreakerBackOff) ProcessRequest() error {
	cbb.mu.Lock()
	defer cbb.mu.Unlock()

	if cbb.CurrentState == Open {
		if time.Since(cbb.lastFailureTime) < cbb.ThresholdTimeout {
			fmt.Println("Circuit is open, Within threshold limit. Returning")
			// reject the request without retry
			return errors.New("Db failure ")
		} else {
			fmt.Println("Circuit is closing, Processing this request since threshold exceeded")
			cbb.CurrentState = Closed
		}
	}

	// try the request
	err := dbOperation()
	if err != nil {
		cbb.lastFailureTime = time.Now()
		cbb.FailureCount++

		if cbb.FailureCount >= cbb.FailureCountThreshold {
			fmt.Println("Circuit is opening, within failure threshold count ")
			cbb.CurrentState = Open
		}

		return err
	}

	// operation successful
	cbb.FailureCount = 0
	return nil
}

func ProcessRequestWithBackoff(cb *CircuitBreakerBackOff, maxRetries int, baseDelay time.Duration) error {
	for i := 1; i <= maxRetries; i++ {
		fmt.Println("Request try count = ", i)
		err := cb.ProcessRequest()
		if err == nil {
			fmt.Println("Request successful, returning...")
			return nil
		}

		delay := baseDelay * time.Duration(i)
		fmt.Println("Request failed, retrying with a delay : ", delay)
		time.Sleep(delay)
	}

	return errors.New("max retries exceeded")
}

func dbOperation() error {
	if rand.Float32() < 0.6 {
		return errors.New("Db Failure ")
	}

	return nil
}
