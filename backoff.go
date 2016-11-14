//Package backoff contains combinators for creating and managing back-off strategies.
package backoff

import (
	"errors"
	"time"
)

//InfiniteDelay is the largest representable time.Duration.
const InfiniteDelay = time.Duration(^uint64(0) >> 1)

var errNoNext = errors.New("cannot call Delay if Next returns false")

//A Strategy models a stream of time.Durations, which can be finite or infinite.
//A Strategy is not thread-safe and it is not reusable.
type Strategy interface {
	//Next reports whether there is another delay to return.
	Next() bool
	//Delay returns the next duration in the stream.
	//Delay panics if Next reported false.
	Delay() time.Duration
}

//Next returns (s.Delay(), true) if s.Next() returns true.
//If s.Next() returns false, Next returns (0, false).
//
//If s == nil, Next always returns (0, true).
func Next(s Strategy) (d time.Duration, ok bool) {
	if s == nil {
		return 0, true
	}
	if s.Next() {
		return s.Delay(), true
	}
	return 0, false
}
