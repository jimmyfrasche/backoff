package backoff

import (
	"strings"
	"time"
)

type fixed struct {
	delays []time.Duration
}

//Fixed returns a Strategy which returns each of the passed delays in turn.
//
//If no delays are passed, the Strategy will always return Next() == false.
func Fixed(delays ...time.Duration) Strategy {
	for _, d := range delays {
		if d < 0 {
			panic("delay must be positive")
		}
	}
	return &fixed{
		delays: delays,
	}
}

func (f *fixed) Next() bool {
	return len(f.delays) > 0
}

func (f *fixed) Delay() time.Duration {
	if len(f.delays) == 0 {
		panic(errNoNext)
	}
	ret := f.delays[0]
	f.delays = f.delays[1:]
	return ret
}

func (f *fixed) String() string {
	var items []string
	for _, v := range f.delays {
		items = append(items, v.String())
	}
	return "Fixed(" + strings.Join(items, ", ") + ")"
}
