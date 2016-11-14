package backoff

import (
	"fmt"
	"time"
)

type maxTries struct {
	s     Strategy
	tries int
}

//MaxTries iterates the underlying Strategy at most tries times.
//
//If strategy == nil, None() will be used.
//
//The tries parameter must be positive.
func MaxTries(tries int, strategy Strategy) Strategy {
	if tries < 0 {
		panic("tries must be positive")
	}
	if strategy == nil {
		strategy = None()
	}
	return &maxTries{
		s:     strategy,
		tries: tries,
	}
}

func (m *maxTries) Next() bool {
	if m.tries == 0 {
		return false
	}
	if !m.s.Next() {
		m.tries = 0
		return false
	}
	return true
}

func (m *maxTries) Delay() time.Duration {
	if m.tries == 0 {
		panic(errNoNext)
	}
	m.tries--
	return m.s.Delay()
}

func (m *maxTries) String() string {
	return fmt.Sprintf("MaxTries(%d, %s)", m.tries, m.s)
}
