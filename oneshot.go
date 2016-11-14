package backoff

import (
	"fmt"
	"time"
)

type oneShot struct {
	fired bool
	delay time.Duration
}

//OneShot is equivalent to Fixed(delay).
func OneShot(delay time.Duration) Strategy {
	if delay < 0 {
		panic("delay must be positive")
	}
	return &oneShot{
		delay: delay,
	}
}

//Zero is equivalent to Fixed(0) and OneShot(0).
func Zero() Strategy {
	return &oneShot{}
}

func (o *oneShot) Next() bool {
	return !o.fired
}

func (o *oneShot) Delay() time.Duration {
	if o.fired {
		panic(errNoNext)
	}
	o.fired = true
	return o.delay
}

func (o *oneShot) String() string {
	if o.delay == 0 {
		return "Zero"
	}
	return fmt.Sprintf("OneShot(%s)", o.delay)
}
