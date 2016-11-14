package backoff

import (
	"fmt"
	"time"
)

type maxDelay struct {
	s    Strategy
	max  time.Duration
	done bool
}

//MaxDelay will iterate the underlying Strategy, stopping after
//Delay() >= max or the underlying Strategy is exhausted.
//The returned Delay will never be greater than max.
//
//If strategy == nil, the returned Strategy will be equivalent to OneShot(max).
//
//The max must be positive.
func MaxDelay(max time.Duration, strategy Strategy) Strategy {
	if max < 0 {
		panic("max must be positive")
	}
	if strategy == nil {
		return OneShot(max)
	}
	return &maxDelay{
		s:   strategy,
		max: max,
	}
}

func (m *maxDelay) Next() bool {
	if m.done {
		return false
	}
	if !m.s.Next() {
		m.done = true
		return false
	}
	return true
}

func (m *maxDelay) Delay() time.Duration {
	if m.done {
		panic(errNoNext)
	}
	ret := m.s.Delay()
	if ret >= m.max {
		ret, m.done = m.max, true
	}
	return ret
}

func (m *maxDelay) String() string {
	return fmt.Sprintf("MaxDelay(%s, %s)", m.max, m.s)
}
