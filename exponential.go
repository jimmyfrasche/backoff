package backoff

import (
	"fmt"
	"math"
	"time"
)

type exponential struct {
	delay time.Duration
	mul   float64
}

//Exponential returns an exponential back-off Strategy.
//
//An exponential strategy will always return Next() == true.
//
//The first Delay() returned will always be delay.
//Afterwards, Delay() will be the previous delay multiplied by multiplier.
//
//If allowed to continue indefinitely,
//it will eventually always return Delay() == InfiniteDelay.
//
//The delay and multiplier must be positive.
//The multiplier must not be infinite or NaN.
func Exponential(delay time.Duration, multiplier float64) Strategy {
	if delay < 0 {
		panic("delay must be positive")
	}
	if multiplier < 0 || multiplier != multiplier || math.IsInf(multiplier, 0) {
		panic("invalid multiplier")
	}
	return &exponential{
		delay: delay,
		mul:   multiplier,
	}
}

func (e *exponential) Next() bool {
	return true
}

func (e *exponential) Delay() time.Duration {
	ret := e.delay
	e.delay = time.Duration(e.mul * float64(e.delay))
	if e.delay < ret {
		e.delay = InfiniteDelay
	}
	return ret
}

func (e *exponential) String() string {
	return fmt.Sprintf("Exp(%s, %g)", e.delay, e.mul)
}
