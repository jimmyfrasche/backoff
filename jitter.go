package backoff

import (
	"math/rand"
	"time"
)

//Jitter introduces random jitter into a Strategy.
//
//The scale must be within the closed interval [0, 1].
//
//The jitter introduced for a Delay d will be within
//the closed interval [-1*scale*d, 1+scale*d].
//
//If strategy is nil, None() will be used.
func Jitter(scale float64, strategy Strategy) Strategy {
	if scale < 0 || scale > 1 {
		panic("invalid jitter")
	}
	return Map(strategy, func(v time.Duration) time.Duration {
		r := 2*rand.Float64() - 1
		jitter := r * scale * float64(v)
		return v + time.Duration(jitter)
	})
}
