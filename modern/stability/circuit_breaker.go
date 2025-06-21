package stability

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	currentState   State
	failureCount   int
	failureTimeOut time.Duration
	lastFailure    time.Time
	failureLimit   int
}

func NewCircuitBreaker(failureTimeOut time.Duration, failureLimit int) *CircuitBreaker {
	return &CircuitBreaker{
		currentState:   Closed,
		failureCount:   0,
		failureTimeOut: failureTimeOut,
		failureLimit:   failureLimit,
	}
}

func (cb *CircuitBreaker) SendRequest(fn func() error) error {
	switch cb.currentState {
	case Open:
		if time.Now().Sub(cb.lastFailure) > cb.failureTimeOut {
			cb.currentState = HalfOpen
			fmt.Println("State changed to half-open")
		} else {
			fmt.Println("Current state is open. Requests are blocked")
			return errors.New("Request is blocked as the current state is open ")
		}
	case HalfOpen:
		fmt.Println("Current state is half-open")
	}

	err := fn()
	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()
		if cb.failureCount >= cb.failureLimit {
			cb.currentState = Open
			fmt.Println("Current state changed to open as failure limit exceeds")
		}

		return err
	}

	if cb.currentState == HalfOpen {
		fmt.Println("Request is successful. State changes to Closed from Half Open")
	}

	cb.failureCount = 0
	cb.currentState = Closed

	return nil
}

// RunAPI mock the unstable API
func RunAPI() error {
	if rand.Float32() < 0.6 {
		return errors.New("API Failure")
	}

	return nil
}
