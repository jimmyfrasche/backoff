package backoff

import (
	"fmt"
	"time"
)

type linear struct {
	delay, incr time.Duration
}

//Linear returns a linear back-off Strategy.
//
//A linear strategy will always return Next() == true.
//
//The first Delay() returned will always be delay.
//Afterwards, Delay() will be the previous delay plus increment.
//
//If allowed to continue indefinitely,
//it will eventually always return Delay() == InfiniteDelay.
//
//The delay and increment must be positive.
func Linear(delay, increment time.Duration) Strategy {
	if delay < 0 {
		panic("delay must be positive")
	}
	return &linear{
		delay: delay,
		incr:  increment,
	}
}

//Constant returns a Strategy that always returns delay.
//
//It is equivalent to Linear(delay, 0).
func Constant(delay time.Duration) Strategy {
	return Linear(delay, 0)
}

//None represent no back-off Strategy.
//It always returns 0 Delay.
//
//It is equivalent to Linear(0, 0) and Constant(0).
func None() Strategy {
	return &linear{}
}

func (l *linear) Next() bool {
	return true
}

func (l *linear) Delay() time.Duration {
	ret := l.delay
	l.delay += l.incr
	if l.delay < ret {
		l.delay = InfiniteDelay
	}
	return ret
}

func (l *linear) String() string {
	if l.incr == 0 {
		if l.delay == 0 {
			return "None"
		}
		return fmt.Sprintf("Constant(%s)", l.delay)
	}
	return fmt.Sprintf("Linear(%s, %s)", l.delay, l.incr)
}
