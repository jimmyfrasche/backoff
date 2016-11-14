package backoff

import (
	"fmt"
	"time"
)

type mapper struct {
	Strategy
	f func(time.Duration) time.Duration
}

//Map calls f for each Delay returned by the underlying strategy.
//
//If strategy is nil, None() is used.
//
//If f is nil, it is not applied.
//
//If f returns a duration < 0, it will be clamped to 0.
func Map(strategy Strategy, f func(time.Duration) time.Duration) Strategy {
	if strategy == nil {
		strategy = None()
	}
	if f == nil {
		return strategy
	}
	return &mapper{
		Strategy: strategy,
		f:        f,
	}
}

func (m *mapper) Delay() time.Duration {
	v := m.Strategy.Delay()
	ret := m.f(v)
	if ret <= 0 {
		return 0
	}
	return ret
}

func (m *mapper) String() string {
	return fmt.Sprintf("Map(%s)", m.Strategy)
}
